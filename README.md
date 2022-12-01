# Server Service

> This repository is experimental meaning that it's based on untested ideas or techniques and not yet established or finalized or involves a radically new and innovative style!
> This means that support is best effort (at best!) and we strongly encourage you to NOT use this in production.

The server service is a microservice within the Hollow eco-system. Server service is responsible for providing a store for physical server information. Support to storing the device components that make up the server is available. You are also able to create attributes and versioned-attributes for both servers and the server components.

## Quickstart to running locally

### Running server service

To run the server service locally you can bring it up with docker-compose. This will run with released images from the hollow container registry.

```bash
docker-compose -f quickstart.yml up
```
### Enable tracing

To run the server service locally with tracing enabled you just need to include the `quickstart-tracing.yml` file.

```bash
docker-compose -f quickstart.yml -f quickstart-tracing.yml up
```

### Running with local changes

The `quickstart.yml` compose file will run server service from released images and not the local code base. If you are doing development and want to run with your local code you can use the following command.

```bash
docker-compose -f quickstart.yml -f quickstart-dev.yml up --build
```

NOTE: `--build` is required to get docker-compose to rebuild the container if you have changes. You make also include the `quickstart-tracing.yml` file if you wish to have tracing support.


### Adding/Changing database schema

Add a new migration file under `db/migrations/` with the schema change

```bash
make docker-up
make test-database
sqlboiler crdb --add-soft-deletes
```

### Run individual integration tests

Export the DB URI required for integration tests.

```bash
export SERVERSERVICE_CRDB_URI="host=localhost port=26257 user=root sslmode=disable dbname=serverservice_test"
```

Run test.

```bash
go test -timeout 30s -tags testtools -run ^TestIntegrationServerListComponents$ go.hollow.sh/serverservice/pkg/api/v1 -v
```
