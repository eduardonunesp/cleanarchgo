help:
	@echo ""
	@echo "Usage: make [command]"
	@echo ""
	@echo "Commands:"
	@echo " help                - Display this help message"
	@echo " test                - Run project tests"
	@echo " run_api             - Run API web server"
	@echo " generate            - Generate SQL Models and Interface Mocks"
	@echo " migrate             - Migrate database all migrations"
	@echo " create_migration    - Create new migration"
	@echo " install_tools       - Install golang tools for the project"
	@echo " build_scaffold      - Build project scaffold tool"
	@echo ""

.PHONY: test
test:
	@echo "Testing everything"
	@go test ./... -v --count=1 --failfast

.PHONY: run_api
run_api:
	@echo "Running API web server"
	@go run cmd/api/main.go

.PHONY: integration_test
integration_test:
	echo "Testing everything + integrations tests"
	@go test ./... -tags db_integration

.PHONY: generate, sqlc_generate, go_generate
generate: sqlc_generate go_generate
sqlc_generate:
	@echo "Generating SQL queries and models at /pkg/infra/db"
	@sqlc generate -f pkg/infra/db/sqlc.yaml
go_generate:
	@echo "Generating mock for interfaces at /pkg/test"
	@go generate ./...

.PHONY: migrate, migrate_all
migrate: migrate_all
migrate_all:
	@echo "Migrate all migrations up"
	@migrate -path ./pkg/infra/db/migrations -database "postgresql://$(PG_CONN_STR)" up

.PHONY: migrate_one_up
migrate_one_up:
	@echo "Migrate one migration up"
	@migrate -path ./pkg/infra/db/migrations -database "postgresql://$(PG_CONN_STR)" up 1

.PHONY: migrate_all_down
migrate_all_down:
	@echo "Migrate all migrations down"
	@migrate -path ./pkg/infra/db/migrations -database "postgresql://$(PG_CONN_STR)" down

.PHONY: migrate_onte_down
migrate_one_down:
	@echo "Migrate one migration down"
	@migrate -path ./pkg/infra/db/migrations -database "postgresql://$(PG_CONN_STR)" down 1

.PHONY: create_migration
create_migration:
	@echo "Creating new migration"
	@migrate create -ext sql -dir ./pkg/infra/db/migrations -seq $(name)

.PHONY: install_tools
install_tools:
	@go install github.com/vektra/mockery/v2@latest
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	@echo ""
	@echo "Don't forget to add GOPATH/bin env var to PATH"

.PHONY: build_scaffold
build_scaffold:
	@go build -o bin/scaffold cmd/scaffold/*.go
	@echo "Scaffold tool builded at ./scaffold"
	@echo "Don't forget to add GOPATH/bin env var to PATH"