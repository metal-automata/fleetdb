//go:build testtools
// +build testtools

package dbtools

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Register the Postgres driver.
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"gocloud.dev/secrets"

	// import gocdk secret drivers
	_ "gocloud.dev/secrets/localsecrets"

	"github.com/metal-automata/fleetdb/internal/models"
)

// TestDBURI is the URI for the test database
var TestDBURI = os.Getenv("FLEETDB_PGDB_URI")
var testDB *sqlx.DB
var testKeeper *secrets.Keeper

func TestDatastore(t *testing.T) *sqlx.DB {
	t.Helper()
	assert.NotNil(t, testDB)
	assert.NoError(t, testDB.Ping())

	return testDB
}

func testDatastore(t *testing.T) error {
	// don't setup the datastore if we already have one
	if testDB != nil {
		return nil
	}

	// Uncomment when you are having database issues with your tests and need to see the db logs
	// Hidden by default because it can be noisy and make it harder to read normal failures.
	// You can also enable at the beginning of your test and then disable it again at the end
	// boil.DebugMode = true

	db, err := sqlx.Open("postgres", TestDBURI)
	if err != nil {
		return err
	}

	boil.SetDB(db)

	testDB = db

	cleanDB(t)

	if err := addFixtures(t); err != nil {
		return errors.Wrap(err, "error adding fixtures")
	}

	return nil
}

// TestSecretKeeper will return the secret keeper we are using for this test run. This allows
// use to use the same one for the entire test run so the secrets are able to be decrypted.
func TestSecretKeeper(t *testing.T) *secrets.Keeper {
	if testKeeper != nil {
		return testKeeper
	}

	keeper, err := secrets.OpenKeeper(context.TODO(), "base64key://")
	require.NoError(t, err)

	testKeeper = keeper

	return keeper
}

// DatabaseTest allows you to run tests that interact with the database
func DatabaseTest(t *testing.T) *sqlx.DB {
	RegisterHooks()

	if testing.Short() {
		t.Skip("skipping database test in short mode")
	}

	t.Cleanup(func() {
		cleanDB(t)
		err := addFixtures(t)
		require.NoError(t, err, "Unexpected error setting up fixture data")
	})

	err := testDatastore(t)
	require.NoError(t, err, "Unexpected error in setting up test datastore")

	return testDB
}

type deleteable interface {
	DeleteAll(context.Context, boil.ContextExecutor) (int64, error)
}

// nolint
func cleanDB(t *testing.T) {
	t.Helper()

	ctx := context.TODO()
	// Make sure the deletion goes in order so you don't break the databases foreign key constraints
	deleteFixture(ctx, t, models.Attributes())
	deleteFixture(ctx, t, models.VersionedAttributes())
	deleteFixture(ctx, t, models.ServerComponents())
	deleteFixture(ctx, t, models.ServerComponentTypes())
	deleteFixture(ctx, t, models.ServerCredentials())
	if _, err := models.Servers(qm.WithDeleted()).DeleteAll(ctx, boil.GetContextDB(), true); err != nil {
		t.Error(errors.Wrap(err, "table: model.Servers"))
	}

	deleteFixture(ctx, t, models.AttributesFirmwareSets())
	deleteFixture(ctx, t, models.ComponentFirmwareSets())
	deleteFixture(ctx, t, models.ComponentFirmwareSetMaps())
	deleteFixture(ctx, t, models.ComponentFirmwareVersions())

	// don't delete the builtin ServerCredentialTypes. Those are expected to exist for the application to work
	deleteFixture(ctx, t, models.ServerCredentialTypes(models.ServerCredentialTypeWhere.Builtin.EQ(false)))

	deleteFixture(ctx, t, models.BiosConfigSettings())
	deleteFixture(ctx, t, models.BiosConfigComponents())
	deleteFixture(ctx, t, models.BiosConfigSets())
	deleteFixture(ctx, t, models.HardwareVendors())
	deleteFixture(ctx, t, models.HardwareModels())
	deleteFixture(ctx, t, models.ServerBMCS())
	deleteFixture(ctx, t, models.InstalledFirmwares())
	deleteFixture(ctx, t, models.ComponentStatuses())
	deleteFixture(ctx, t, models.ServerStatuses())
	deleteFixture(ctx, t, models.ComponentCapabilities())
	deleteFixture(ctx, t, models.ComponentMetadata())
	deleteFixture(ctx, t, models.ComponentChangeReports())
}

func deleteFixture(ctx context.Context, t *testing.T, fixture deleteable) {
	t.Helper()

	if _, err := fixture.DeleteAll(ctx, testDB); err != nil {
		t.Error(errors.Wrap(err, fmt.Sprintf("table: %s", reflect.TypeOf(fixture))))
	}
}
