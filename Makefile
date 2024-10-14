export DOCKER_BUILDKIT=1
GOOS=linux
DEV_DB_URI := host=postgresql port=5432 user=fleetdb sslmode=disable dbname=fleetdb
TEST_DB_URI := host=localhost port=5432 user=fleetdb_test sslmode=disable dbname=fleetdb_test
DOCKER_IMAGE := "ghcr.io/metal-automata/fleetdb"
PROJECT_NAME := fleetdb
REPO := "https://github.com/metal-automata/fleetdb.git"
SQLBOILER := v4.15.0

.DEFAULT_GOAL := help

## run all tests
test: | unit-test integration-test

## run integration tests
integration-test: test-database
	@echo Running integration tests...
	@FLEETDB_CRDB_URI="${TEST_DB_URI}" go test -race -cover -tags testtools,integration \
	                                           -coverprofile=coverage.txt -covermode=atomic -p 1 -timeout 2m ./... | \
	grep -v "could not be registered in Prometheus\" error=\"duplicate metrics collector registration attempted\"" # TODO; Figure out why this message spams when tests fail

## run unit tests
unit-test: | test-database
	@echo Running unit tests...
	@FLEETDB_CRDB_URI="${TEST_DB_URI}" go test -cover -short -tags testtools ./...

## run single integration test. Example: make single-test test=TestIntegrationServerListComponents
single-test: test-database
	@FLEETDB_CRDB_URI="${TEST_DB_URI}" go test -timeout 30s -tags testtools -run ^${test}$$ github.com/metal-automata/fleetdb/pkg/api/v1 -v

## check test coverage
coverage: | test-database
	@echo Generating coverage report...
	@FLEETDB_CRDB_URI="${TEST_DB_URI}" go test ./... -race -coverprofile=coverage.out -covermode=atomic -tags testtools,integration -p 1
	@go tool cover -func=coverage.out
	@go tool cover -html=coverage.out

## lint
lint:
	@echo Linting Go files...
	@golangci-lint run

## clean docker files
clean: docker-clean test-clean
	@echo Cleaning...
	@rm -rf ./dist/
	@rm -rf coverage.out

## clean test env
test-clean:
	@go clean -testcache

## generate db models
gen-db-models: install-sqlboiler test-database
	@sqlboiler psql --add-soft-deletes

## setup fleetdb PG db container for tests and run migrations
test-database:
	@PG_DSN="${TEST_DB_URI}" docker compose -f quickstart.yml up -d postgresql
	@until pg_isready -d "${TEST_DB_URI}"; do echo "waiting for PG to be ready..."; sleep 1; done
	@psql -d "host=localhost port=5432 user=postgres sslmode=disable dbname=postgres" \
	    -c "drop database if exists fleetdb_test;" \
		-c "drop owned by fleetdb_test;" \
		-c "drop role if exists fleetdb_test;" \
		-c "create role fleetdb_test login createdb;" \
		-c "create database fleetdb_test owner fleetdb_test;" \
		-c "grant all privileges on schema public to fleetdb_test;"
	@FLEETDB_CRDB_URI="${TEST_DB_URI}" go run main.go migrate up
	# The attributes, versioned_attributes constraints are dropped to allow generated db model tests to succeed
	@psql -d "host=localhost port=5432 user=fleetdb_test sslmode=disable dbname=fleetdb_test" \
		-c "ALTER TABLE attributes DROP CONSTRAINT check_server_id_server_component_id; ALTER TABLE versioned_attributes DROP CONSTRAINT check_server_id_server_component_id;"

test-database-down:
	@PG_DSN="${TEST_DB_URI}" docker compose -f quickstart.yml down

## setup fleetdb docker dev env
dev-env-up: push-image-devel
	@PG_DSN="${DEV_DB_URI}" docker compose -f quickstart.yml up -d postgresql
	@psql -d "host=localhost port=5432 user=postgres sslmode=disable dbname=postgres" \
	       	-c "drop database if exists fleetdb;" \
		-c "drop owned by fleetdb;" \
		-c "drop role if exists fleetdb;" \
		-c "create role fleetdb login;" \
		-c "create database fleetdb owner fleetdb;" \
		-c "grant all privileges on schema public to fleetdb;"
	@PG_DSN="${DEV_DB_URI}" docker compose -f quickstart.yml up -d fleetdb-migrate
	@PG_DSN="${DEV_DB_URI}" docker compose -f quickstart.yml up -d fleetdb

## stop docker compose test env
dev-env-down:
	@PG_DSN="${DEV_DB_URI}" docker compose -f quickstart.yml down

## stop docker and clean volumes
dev-env-clean:
	@PG_DSN="${DEV_DB_URI}" docker compose -f quickstart.yml down --volumes

## install sqlboiler
install-sqlboiler:
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@${SQLBOILER}
	go install github.com/volatiletech/sqlboiler/v4@${SQLBOILER}

## log into dev database
psql-dev-db:
	@psql -d "${DEV_DB_URI}"

## log into test database
psql-test-db:
	@psql -d "${TEST_DB_URI}"

## Build linux bin
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${PROJECT_NAME}

## build docker image and tag as ghcr.io/metal-automata/fleetdb:latest
build-image: build-linux
	docker build --rm=true -f Dockerfile -t ${DOCKER_IMAGE}:latest . \
		--label org.label-schema.schema-version=1.0 \
		--label org.label-schema.vcs-ref=${GIT_COMMIT_FULL} \
		--label org.label-schema.vcs-url=${REPO}

## build and push devel docker image to KIND image repo used by the sandbox - https://github.com/metal-automata/sandbox
push-image-devel: build-image
	docker tag ${DOCKER_IMAGE}:latest localhost:5001/${PROJECT_NAME}:latest
	docker push localhost:5001/${PROJECT_NAME}:latest
	kind load docker-image localhost:5001/${PROJECT_NAME}:latest

# https://gist.github.com/prwhite/8168133
# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

TARGET_MAX_CHAR_NUM=20

## Show help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-${TARGET_MAX_CHAR_NUM}s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' ${MAKEFILE_LIST}
