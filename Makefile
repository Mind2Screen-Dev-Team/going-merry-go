# Check and create .env file if not present using a shell command
.env:
	@echo ""
	@echo "Read .env file not found. Copying from .example.env..."
	@echo ""
	@cp .example.env .env

include .env

# Set GOPATH to your desired Go path, or it defaults to the system GOPATH if not set
GOPATH ?= $(shell go env GOPATH)
GOPATH_BIN := $(GOPATH)/bin

# Prepend GOPATH/bin to PATH if not already present
ifneq ($(filter $(GOPATH_BIN), $(PATH)),)
    export PATH := $(PATH)
else
    export PATH := $(GOPATH_BIN):$(PATH)
endif

# List all proto files in the proto directory
PROTO_FILES := $(wildcard $(APP_PROTO_DIR)/*.proto)

help:
	@echo "Application Available Commands"
	@echo ""
	@echo "Usage: make [commands]"
	@echo ""
	@echo "Commands:"
	@echo "  setup                      Make setup your project workspace"
	@echo "  migrate-help               Make migarte help command"
	@echo "  pkl-gen                    Make pkl help you to create configuration"
	@echo "  go-help                    Make go help command"
	@echo "  print-path                 Make print current path variable"
	@echo ""
	@echo "Examples:"
	@echo "  make setup"
	@echo "  make pkl-gen"
	@echo "  make go-help"
	@echo "  make migrate-help"

print-path:
	@echo "Current PATH: $(PATH)"

# Command Setup
setup:
	@echo "- Install Protoc Generate GO Tool"
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest || { echo 'Error: protoc-gen-go installation failed.'; exit 1; }
	@echo "- Install Protoc Generate GO GRPC Tool"
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest || { echo 'Error: protoc-gen-go-grpc installation failed.'; exit 1; }
	@echo "- Install Goose For Migration & Seeder Tool"
	@go install github.com/pressly/goose/v3/cmd/goose@latest || { echo 'Error: goose installation failed.'; exit 1; }
	@echo "- Install PKL GO Configuration Generator Tool"
	@go install github.com/apple/pkl-go/cmd/pkl-gen-go@v0.8.0 || { echo 'Error: goose installation failed.'; exit 1; }
	@echo "- Checking installed tools in GOPATH/bin..."
	@echo ""
	@ls $(shell go env GOPATH)/bin || { echo 'Error: check installation tools in GOPATH/bin.'; exit 1; }
	@echo ""
	@echo "Setup Installation Tools Is Complete!"

# Command to build config commands

pkl-gen:
	@pkl-gen-go ./pkl/AppConfig.pkl --base-path $(APP_BASE_MODULE)

# Command to run golang commands

go-tidy:
	@go mod tidy

go-run:
	@if [ -z "$(cfg)" ]; then \
		go run ./cmd/$(app); \
	else \
		go run ./cmd/$(app) -cfg=$(cfg); \
	fi

go-build:
	@go build -o ./bin/$(app) ./cmd/$(app)

go-gen-proto:
	@echo "Generating Go code from .proto files..."
	@for proto_file in $(PROTO_FILES); do \
		echo "Processing: $$proto_file"; \
		DIR_NAME=$$(basename $$proto_file .proto); \
		mkdir -p $(APP_PROTO_GEN_DIR)/$$DIR_NAME; \
		protoc --go_out=$(APP_PROTO_GEN_DIR)/$$DIR_NAME --go_opt=paths=source_relative \
		       --go-grpc_out=$(APP_PROTO_GEN_DIR)/$$DIR_NAME --go-grpc_opt=paths=source_relative \
		       -I $(APP_PROTO_DIR) $$proto_file; \
	done
	@echo "Go code generation complete!"

go-help:
	@echo "Application Golang Commands"
	@echo ""
	@echo "Usage: make [commands] [OPTIONS]"
	@echo ""
	@echo "Commands:"
	@echo "  go-tidy                    Go module tidy"
	@echo "  go-run APPLICATION         Go run application"
	@echo "  go-build APPLICATION       Go build application"
	@echo "  go-gen-proto               Go generate grpc protobuff"
	@echo ""
	@echo "Examples:"
	@echo "  make go-tidy"
	@echo "  make go-run app=restapi"
	@echo "  make go-run app=restapi cfg=<your_pkl_config_file>"
	@echo "  make go-build app=restapi"
	@echo "  make go-gen-proto"

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