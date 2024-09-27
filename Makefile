include .env

# Command to build config commands

pkl-gen-cfg:
	@pkl-gen-go ./pkl/AppConfig.pkl --base-path $(APP_BASE_MODULE)

# Command to run golang commands

go-tidy:
	@go mod tidy

go-run:
	@go run ./cmd/$(app) -cfg=$(cfg)

go-build:
	@go build -o ./bin/$(app) ./cmd/$(app)

go-help:
	@echo "Application Golang Commands"
	@echo ""
	@echo "Usage: make [commands] [OPTIONS]"
	@echo ""
	@echo "Commands:"
	@echo "  go-tidy                    Go module tidy"
	@echo "  go-run APPLICATION         Go run application"
	@echo "  go-build APPLICATION       Go build application"
	@echo ""
	@echo "Examples:"
	@echo "  make go-tidy"
	@echo "  make go-run app=restapi cfg=<your_pkl_config_file>"
	@echo "  make go-build app=restapi"

# Command to run goose with the specified options
GOOSE_MIGRATE_CMD = GOOSE_DRIVER=$(DB_GOOSE_DRIVER) GOOSE_DBSTRING=$(DB_GOOSE_DBSTRING) GOOSE_MIGRATION_DIR=$(DB_GOOSE_MIGRATION_DIR) goose -table $(DB_GOOSE_MIGRATION_TABLE)

# Migration commands
migrate-up:
	@$(GOOSE_MIGRATE_CMD) up

migrate-up-by-one:
	@$(GOOSE_MIGRATE_CMD) up-by-one

migrate-up-to:
	@$(GOOSE_MIGRATE_CMD) up-to $(v)

migrate-down:
	@$(GOOSE_MIGRATE_CMD) down

migrate-down-to:
	@$(GOOSE_MIGRATE_CMD) down-to $(v)

migrate-create:
	@$(GOOSE_MIGRATE_CMD) create $(n) $(t)

migrate-redo:
	@$(GOOSE_MIGRATE_CMD) redo

migrate-reset:
	@$(GOOSE_MIGRATE_CMD) reset

migrate-status:
	@$(GOOSE_MIGRATE_CMD) status

migrate-version:
	@$(GOOSE_MIGRATE_CMD) version

migrate-fix:
	@$(GOOSE_MIGRATE_CMD) fix

migrate-validate:
	@$(GOOSE_MIGRATE_CMD) validate

# Command to run goose with the specified options
GOOSE_SEEDER_CMD = GOOSE_DRIVER=$(DB_GOOSE_DRIVER) GOOSE_DBSTRING=$(DB_GOOSE_DBSTRING) GOOSE_MIGRATION_DIR=$(DB_GOOSE_MIGRATION_SEEDER_DIR) goose -table $(DB_GOOSE_MIGRATION_SEEDER_TABLE)

# Seeders commands
seeder-up:
	@$(GOOSE_SEEDER_CMD) up

seeder-up-by-one:
	@$(GOOSE_SEEDER_CMD) up-by-one

seeder-up-to:
	@$(GOOSE_SEEDER_CMD) up-to $(v)

seeder-down:
	@$(GOOSE_SEEDER_CMD) down

seeder-down-to:
	@$(GOOSE_SEEDER_CMD) down-to $(v)

seeder-create:
	@$(GOOSE_SEEDER_CMD) create $(n) $(t)

seeder-redo:
	@$(GOOSE_SEEDER_CMD) redo

seeder-reset:
	@$(GOOSE_SEEDER_CMD) reset

seeder-status:
	@$(GOOSE_SEEDER_CMD) status

seeder-version:
	@$(GOOSE_SEEDER_CMD) version

seeder-fix:
	@$(GOOSE_SEEDER_CMD) fix

seeder-validate:
	@$(GOOSE_SEEDER_CMD) validate

# Help messages
seeder-help:
	@make migrate-help

migrate-help:
	@echo "Migrations and Seeders Docs"
	@echo ""
	@echo "Usage: make [migrate | seeder]-[command] [OPTIONS]"
	@echo ""
	@echo "Commands:"
	@echo "  ( migrate | seeder )-up                    Migrate the DB to the most recent version available"
	@echo "  ( migrate | seeder )-up-by-one             Migrate the DB up by 1"
	@echo "  ( migrate | seeder )-up-to VERSION         Migrate the DB to a specific VERSION"
	@echo "  ( migrate | seeder )-down                  Roll back the version by 1"
	@echo "  ( migrate | seeder )-down-to VERSION       Roll back to a specific VERSION"
	@echo "  ( migrate | seeder )-create NAME [sql|go]  Creates new migration file with the current timestamp"
	@echo "  ( migrate | seeder )-redo                  Re-run the latest migration"
	@echo "  ( migrate | seeder )-reset                 Roll back all migrations"
	@echo "  ( migrate | seeder )-status                Dump the migration status for the current DB"
	@echo "  ( migrate | seeder )-version               Print the current version of the database"
	@echo "  ( migrate | seeder )-fix                   Apply sequential ordering to migrations"
	@echo "  ( migrate | seeder )-validate              Check migration files without running them"
	@echo ""
	@echo "Options by env file:"
	@echo "  DB_GOOSE_DRIVER         Database driver (postgres, mysql, sqlite3, mssql, redshift, tidb, clickhouse, vertica, ydb, turso)"
	@echo "  DB_GOOSE_DBSTRING       Connection string for the database"
	@echo "  DB_GOOSE_MIGRATION_DIR  Directory for migration files (default: current directory)"
	@echo ""
	@echo "Examples:"
	@echo "  make ( migrate | seeder )-up" 
	@echo "  make ( migrate | seeder )-up-by-one"
	@echo "  make ( migrate | seeder )-up-to v=20240922160357"
	@echo "  make ( migrate | seeder )-down"
	@echo "  make ( migrate | seeder )-down-to v=20240922160357"
	@echo "  make ( migrate | seeder )-status"
	@echo "  make ( migrate | seeder )-version"
	@echo "  make ( migrate | seeder )-create n=<migration_name> t=<sql|go>"
	@echo "  make ( migrate | seeder )-validate"