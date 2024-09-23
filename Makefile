include .env

# Command to run golang commands
tidy:
	@go mod tidy

run:
	@go run ./cmd/$(api)

build:
	@go build -o ./bin/$(api) ./cmd/$(api)

# Command to run goose with the specified options
GOOSE_MIGRATE_CMD = GOOSE_DRIVER=$(DB_GOOSE_DRIVER) GOOSE_DBSTRING=$(DB_GOOSE_DBSTRING) GOOSE_MIGRATION_DIR=$(DB_GOOSE_MIGRATION_DIR) goose -table $(DB_GOOSE_MIGRATION_TABLE)

# Migration commands
migrate-up:
	@$(GOOSE_MIGRATE_CMD) up

migrate-down:
	@$(GOOSE_MIGRATE_CMD) down

migrate-status:
	@$(GOOSE_MIGRATE_CMD) status

migrate-create:
	@$(GOOSE_MIGRATE_CMD) create $(n) $(t)

migrate-reset:
	@$(GOOSE_MIGRATE_CMD) reset

migrate-fix:
	@$(GOOSE_MIGRATE_CMD) fix

migrate-validate:
	@$(GOOSE_MIGRATE_CMD) validate

# Command to run goose with the specified options
GOOSE_SEEDER_CMD = GOOSE_DRIVER=$(DB_GOOSE_DRIVER) GOOSE_DBSTRING=$(DB_GOOSE_DBSTRING) GOOSE_MIGRATION_DIR=$(DB_GOOSE_MIGRATION_SEEDER_DIR) goose -table $(DB_GOOSE_MIGRATION_SEEDER_TABLE)

# Seeders commands
seeder-up:
	@$(GOOSE_SEEDER_CMD) up

seeder-down:
	@$(GOOSE_SEEDER_CMD) down

seeder-status:
	@$(GOOSE_SEEDER_CMD) status

seeder-create:
	@$(GOOSE_SEEDER_CMD) create $(n) $(t)

seeder-reset:
	@$(GOOSE_SEEDER_CMD) reset

seeder-fix:
	@$(GOOSE_SEEDER_CMD) fix

seeder-validate:
	@$(GOOSE_SEEDER_CMD) validate

# Help messages
migrate-help:
	@echo "Usage: make [command] [OPTIONS]"
	@echo ""
	@echo "Commands:"
	@echo "  (migrate / seeder)-up                    Migrate the DB to the most recent version available"
	@echo "  (migrate / seeder)-down                  Roll back the version by 1"
	@echo "  (migrate / seeder)-status                Dump the migration status for the current DB"
	@echo "  (migrate / seeder)-create NAME [sql|go]  Creates new migration file with the current timestamp"
	@echo "  (migrate / seeder)-reset                 Roll back all migrations"
	@echo "  (migrate / seeder)-fix                   Apply sequential ordering to migrations"
	@echo "  (migrate / seeder)-validate              Check migration files without running them"
	@echo ""
	@echo "Options by env file:"
	@echo "  DB_GOOSE_DRIVER         Database driver (e.g., postgres, mysql, sqlite3, etc.)"
	@echo "  DB_GOOSE_DBSTRING       Connection string for the database"
	@echo "  DB_GOOSE_MIGRATION_DIR  Directory for migration files (default: current directory)"
	@echo ""
	@echo "Examples:"
	@echo "  make (migrate / seeder)-up"
	@echo "  make (migrate / seeder)-down"
	@echo "  make (migrate / seeder)-status"
	@echo "  make (migrate / seeder)-create n=create_<table_name>_table t=sql"
	@echo "  make (migrate / seeder)-validate"