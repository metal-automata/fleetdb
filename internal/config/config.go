// Package config provides a struct to store the applications config
package config

import (
	"net/url"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/metal-automata/rivets/ginjwt"
	"go.infratographer.com/x/otelx"
)

// AppConfig stores all the config values for our application
var AppConfig struct {
	PGDB    DBConfig
	Logging LoggingConfig
	Tracing otelx.Config
	// APIServerJWTAuth sets the JWT verification configuration for the conditionorc API service.
	APIServerJWTAuth []ginjwt.AuthConfig `mapstructure:"ginjwt_auth"`
}

type LoggingConfig struct {
	Debug             bool `mapstructure:"debug"`
	Pretty            bool `mapstructure:"pretty"`
	DisableStacktrace bool `mapstructure:"disable_stacktrace"`
}

const (
	defaultMaxOpenConns    int           = 25
	defaultMaxIdleConns    int           = 25
	defaultMaxConnLifetime time.Duration = 5 * 60 * time.Second
)

// DBConfig is used to configure a new DB connection
type DBConfig struct {
	Name        string `mapstructure:"name"`
	Host        string `mapstructure:"host"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Params      string `mapstructure:"params"`
	URI         string `mapstructure:"uri"`
	Connections struct {
		MaxOpen     int           `mapstructure:"max_open"`
		MaxIdle     int           `mapstructure:"max_idle"`
		MaxLifetime time.Duration `mapstructure:"max_lifetime"`
	}
}

// GetURI returns the connection URI, if a config URI is provided that will be
// returned, otherwise the host, user, password, and params will be put together
// to make a URI that is returned.
func (c DBConfig) GetURI() string {
	if c.URI != "" {
		return c.URI
	}

	u := url.URL{
		Scheme:   "postgresql",
		User:     url.UserPassword(c.User, c.Password),
		Host:     c.Host,
		Path:     c.Name,
		RawQuery: c.Params,
	}

	return u.String()
}

// MustPGDBViperFlags returns the cobra flags and viper config to prevent code duplication
// and help provide consistent flags across the applications
func MustPGDBViperFlags(v *viper.Viper, _ *pflag.FlagSet) {
	v.MustBindEnv("pgdb.host")
	v.MustBindEnv("pgdb.params")
	v.MustBindEnv("pgdb.user")
	v.MustBindEnv("pgdb.password")
	v.MustBindEnv("pgdb.uri")
	v.MustBindEnv("pgdb.connections.max_open")
	v.MustBindEnv("pgdb.connections.max_idle")
	v.MustBindEnv("pgdb.connections.max_lifetime")

	v.SetDefault("pgdb.host", "localhost:5432")
	v.SetDefault("pgdb.connections.max_open", defaultMaxOpenConns)
	v.SetDefault("pgdb.connections.max_idle", defaultMaxIdleConns)
	v.SetDefault("pgdb.connections.max_lifetime", defaultMaxConnLifetime)
}
