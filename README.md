# FleetDB

FleetDB is responsible for providing a store for physical server and component information.

## Quickstart to running locally

### Running FleetDB

To run the service locally you can bring it up with docker-compose. This will run with a localy built fleetdb.

```bash
make dev-env
```

### Adding/Changing database schema

Add a new migration file under `db/migrations/` with the schema change

```bash
make test-database
make gen-db-models
```

### Run individual integration tests

Export the DB URI required for integration tests.

```bash
export FLEETDB_PGDB_URI="host=localhost port=5432 user=root sslmode=disable dbname=fleetdb_test"
```

Run test.

```bash
go test -timeout 30s -tags testtools -run ^TestIntegrationServerListComponents$ github.com/metal-automata/fleetdb/pkg/api/v1 -v
```

### Dump requests and responses in the client for debugging

Setting `DEBUG_CLIENT` to a value will cause the client to write requests and responses to stdout.
```
export DEBUG_CLIENT=1 /usr/local/bin/go test -timeout 10s -tags testtools -run ^TestIntegrationServerComponentFirmwareUpdate$ github.com/metal-automata/fleetdb/pkg/api/v1  -v
```
