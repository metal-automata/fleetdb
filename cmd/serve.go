package cmd

import (
	"context"
	"database/sql"

	"github.com/XSAM/otelsql"
	"github.com/jmoiron/sqlx"
	"github.com/metal-automata/rivets/ginjwt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/volatiletech/sqlboiler/boil"
	"go.infratographer.com/x/otelx"
	"go.infratographer.com/x/viperx"
	"go.uber.org/zap"
	"gocloud.dev/secrets"

	// import gocdk secret drivers
	_ "gocloud.dev/secrets/localsecrets"

	"github.com/metal-automata/fleetdb/internal/config"
	"github.com/metal-automata/fleetdb/internal/dbtools"
	"github.com/metal-automata/fleetdb/internal/httpsrv"

	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var (
	apiDefaultListen = "0.0.0.0:8000"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts the fleetdbapi server",
	Run: func(cmd *cobra.Command, _ []string) {
		serve(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().String("listen", apiDefaultListen, "address to listen on")
	viperx.MustBindFlag(viper.GetViper(), "listen", serveCmd.Flags().Lookup("listen"))

	otelx.MustViperFlags(viper.GetViper(), serveCmd.Flags())
	config.MustPGDBViperFlags(viper.GetViper(), serveCmd.Flags())

	// OIDC Flags
	serveCmd.Flags().Bool("oidc", true, "use oidc auth")
	ginjwt.BindFlagFromViperInst(viper.GetViper(), "oidc.enabled", serveCmd.Flags().Lookup("oidc"))

	// DB Flags
	serveCmd.Flags().String("db-encryption-driver", "", "encryption driver uri; 32 byte base64 encoded string, (example: base64key://your-encoded-secret-key)")
	viperx.MustBindFlag(viper.GetViper(), "db.encryption_driver", serveCmd.Flags().Lookup("db-encryption-driver"))
}

func serve(ctx context.Context) {
	err := otelx.InitTracer(config.AppConfig.Tracing, appName, logger)
	if err != nil {
		logger.Fatalw("unable to initialize tracing system", "error", err)
	}

	db := initDB()

	dbtools.RegisterHooks()

	if errSetup := dbtools.SetupComponentTypes(ctx, db); errSetup != nil {
		logger.With(
			zap.Error(errSetup),
		).Fatal("set up component types")
	}

	keeper, err := secrets.OpenKeeper(ctx, viper.GetString("db.encryption_driver"))
	if err != nil {
		logger.Fatalw("failed to open secrets keeper", "error", err)
	}
	defer keeper.Close()

	logger.Infow("starting server",
		"address", viper.GetString("listen"),
	)

	var oidcEnabled bool
	if viper.GetViper().GetBool("oidc.enabled") {
		logger.Infow("OIDC enabled")

		if len(config.AppConfig.APIServerJWTAuth) == 0 {
			logger.Fatal("OIDC enabled without configuration")
		}
		oidcEnabled = true
	} else {
		logger.Infow("OIDC disabled")
	}

	hs := &httpsrv.Server{
		Logger:        logger.Desugar(),
		Listen:        viper.GetString("listen"),
		Debug:         config.AppConfig.Logging.Debug,
		DB:            db,
		OIDCEnabled:   oidcEnabled,
		SecretsKeeper: keeper,
		AuthConfigs:   config.AppConfig.APIServerJWTAuth,
	}

	if err := hs.Run(); err != nil {
		logger.Fatalw("failed starting server", "error", err)
	}
}

func initDB() *sqlx.DB {
	var err error
	dbDriverName := "postgres"

	if config.AppConfig.Tracing.Enabled {
		// Register an OTel SQL driver
		dbDriverName, err = otelsql.Register(dbDriverName,
			otelsql.WithAttributes(semconv.DBSystemPostgreSQL))
		if err != nil {
			logger.Fatalw("failed creating sql tracer: %w", err)
		}
	}

	db, err := sql.Open(dbDriverName, config.AppConfig.PGDB.GetURI())
	if err != nil {
		logger.Fatalw("failed to initialize database connection", "error", err)
	}

	if err := db.Ping(); err != nil {
		logger.Fatalw("failed verifying database connection: %w", err)
	}

	db.SetMaxOpenConns(config.AppConfig.PGDB.Connections.MaxOpen)
	db.SetMaxIdleConns(config.AppConfig.PGDB.Connections.MaxIdle)
	db.SetConnMaxIdleTime(config.AppConfig.PGDB.Connections.MaxLifetime)

	boil.SetDB(db)

	return sqlx.NewDb(db, dbDriverName)
}
