
# 🚢 _Going-Merry-Go_ - Project Skeleton

Welcome aboard the **Going-Merry-Go**! Inspired by the iconic ship from the anime One Piece, this project is a robust and flexible Go project starter kit. It's designed to help you quickly set up your Go applications with the right structure and essential tools.

## 🗂 Project Structure

```bash
├── app
│   ├── bootstrap      # Contains initialization logic for starting the application
│   └── registry       # Manages dependency injection or service registration
├── bin                # Holds executable files and scripts for running the application
├── cmd
│   ├── grpc           # Entry point for starting the gRPC server
│   ├── restapi        # Entry point for starting the REST API server
│   ├── scheduler      # Entry point for the scheduler service
│   └── worker         # Entry point for worker processes or background jobs
├── config             # Configuration files for the application
├── constant
│   ├── ctxkey         # Constants for context keys used across the application
│   ├── rediskey       # Redis-related key constants
│   └── restkey        # Constants specific to the REST API
├── database
│   ├── migrations     # Database migration scripts to manage schema changes
│   └── seeders        # Seed data for initializing the database with default values
├── docs
│   └── vscode-ext     # Documentation for Visual Studio Code extensions or related setup
├── gen
│   ├── grpc
│   │   ├── greating   # Auto-generated code for the "greating" gRPC service
│   │   └── health     # Auto-generated code for the health check gRPC service
│   └── pkl
│       ├── appconfig  # Generated code for application configuration
│       ├── grpcconfig # gRPC configuration details
│       ├── httpconfig # HTTP server configuration details
│       ├── jwtconfig  # JWT (JSON Web Token) configuration settings
│       ├── logconfig
│       │   └── timeformat # Time format settings for logging
│       ├── minioconfig # MinIO (object storage) configuration
│       ├── mysqlconfig # MySQL database configuration
│       ├── natsconfig  # NATS (messaging system) configuration
│       ├── otelconfig  # OpenTelemetry configuration for distributed tracing
│       └── redisconfig # Redis database configuration
├── internal
│   ├── entity         # Data models and entities used in the application
│   ├── grpc
│   │   ├── interceptor
│   │   │   ├── stream  # Stream interceptors for gRPC requests
│   │   │   ├── unary   # Unary interceptors for gRPC requests
│   │   │   └── util    # Utility functions for gRPC interceptors
│   │   └── service    # gRPC service implementations
│   ├── http
│   │   ├── dto        # Data Transfer Objects for HTTP requests and responses
│   │   ├── handler    # Handlers for processing HTTP requests
│   │   ├── interceptor # Interceptors for HTTP requests
│   │   ├── middleware # Middleware components for HTTP processing
│   │   └── router     # Routing logic for HTTP endpoints
│   ├── provider
│   │   ├── api        # API providers or service interfaces
│   │   ├── attribute  # Attribute-related logic or utilities
│   │   └── impl       # Implementation of the providers
│   ├── repo
│   │   ├── api        # API for the repository layer
│   │   ├── attribute  # Attribute-related logic for the repository layer
│   │   └── impl       # Implementation of repositories
│   ├── scheduler      # Scheduler logic for timed or recurring tasks
│   ├── service
│   │   ├── api        # Service interfaces
│   │   ├── attribute  # Attribute-related logic for services
│   │   └── impl       # Implementation of services
│   └── worker
│       ├── pub        # Publishing logic for worker processes
│       └── sub        # Subscription logic for worker processes
├── pkg
│   ├── xfilter        # Utility for data filtering
│   ├── xhttpin        # Helper functions for HTTP input processing
│   ├── xhttputil      # Utilities for working with HTTP
│   ├── xlazy          # Lazy evaluation utilities
│   ├── xlogger        # Logging utilities and helpers
│   ├── xresponse      # Utility for handling HTTP responses
│   ├── xtracer        # Tools for distributed tracing
│   └── xvalidate      # Input validation utilities
├── pkl                # Placeholder for Pickle files for configuration
├── protos             # Protocol buffer (.proto) files for defining gRPC services
└── storage
    ├── assets        # Static assets such as images or other resources
    └── logs          # Log files generated by the application
```

The project is organized to support different modules such as REST API, gRPC, Pub/Sub, and Cron Jobs, making it easier for you to develop scalable and maintainable applications.

### Current Modules

- **APP HTTP/1.1** - REST API ✅
- **APP HTTP/2** - gRPC ✅
- **APP Nats.io** - Worker Pub/Sub ⏳ (In Progress)
- **APP CronJob** - Scheduler Cron Job ⏳ (In Progress)

## 📋 Features

Here's a quick look at what's done and what's still in progress:

### Done ✅
- 🗃️ **Base Structural Directory**: Well-organized code structure to get you started quickly.
- ⚙️ **Setup Basic Generator Configuration**: Tools to generate handlers, services, and more.
- 🔧 **Registry Dependency, Repository, and Service**: Dependency injection setup.
- 🌐 **HTTP Handler and Router Loader**: Load and manage routes effortlessly.
- 🛡️ **HTTP Interceptor Handler**: Middleware to handle requests seamlessly.
- 📜 **DTO Validation**: Validate incoming data with ease.
- 📦 **DB Migrations and Seeders**: Database migration and seeding tools.
- 📄 **Logging**: Integrated logging for better observability.
- 📑 **Makefile Runner**: Simple command runners for building and testing.
- 🌍 **Open Telemetry Integration**: Track and monitor your services.

### To Do 📝
- 🐳 **Docker Integration**: Containerize the application.
- 📚 **Open API Generator Docs**: Auto-generate API documentation.
- ⚙️ **CMD Generator**: Tool to generate handlers, middleware, routers, repos, and services.
- 🧪 **Unit Tests**: Comprehensive unit testing setup.

## 📦 Installation and Setup

To get started with Going-Merry-Go, follow these steps:

```bash
# Clone the repository
git clone https://github.com/Mind2Screen-Dev-Team/going-merry-go.git

# Navigate to the project directory
cd going-merry-go

# Install dependencies and set up the project
make setup

# Run the application
make go-run app=restapi
```

## ⚙️ Makefile Commands

The Makefile provides a set of commands to help you manage and interact with your Go project efficiently. Below is a list of the available commands:

### Setup Commands

- **`make setup`**: Sets up the project by installing necessary tools like `protoc-gen-go`, `protoc-gen-go-grpc`, `goose`, and `pkl-gen-go`.

### Go Commands

- **`make go-tidy`**: Cleans up the `go.mod` file by removing unnecessary dependencies.
- **`make go-run app=<application>`**: Runs the specified application.
- **`make go-build app=<application>`**: Builds the specified application.
- **`make go-gen-proto`**: Generates Go code from `.proto` files.

### Migration Commands

- **`make migrate-up`**: Migrates the database to the most recent version.
- **`make migrate-up-by-one`**: Migrates the database up by one version.
- **`make migrate-down`**: Rolls back the database version by one.
- **`make migrate-status`**: Displays the migration status of the database.
- **`make migrate-create n=<migration_name> t=<sql|go>`**: Creates a new migration file.

### Seeder Commands

- **`make seeder-up`**: Runs the seeders to populate the database.
- **`make seeder-down`**: Rolls back the seeders by one version.
- **`make seeder-create n=<seeder_name> t=<sql|go>`**: Creates a new seeder file.

### Utility Commands

- **`make print-path`**: Displays the current `PATH` environment variable.
- **`make migrate-help`**: Provides help on migration commands.
- **`make go-help`**: Provides help on Go commands.

### Examples

```bash
# Setup your project workspace
make setup

# Generate Go code from protobuf files
make go-gen-proto

# Run a Go application (example: restapi)
make go-run app=restapi

# Migrate the database to the latest version
make migrate-up

# Create a new migration file
make migrate-create n=create_users_table t=sql
```

These commands make it easy to manage your Go application, including its dependencies, database migrations, and proto file generation.

## 📖 Documentation

For detailed documentation and advanced usage, please refer to the [Wiki](https://github.com/Mind2Screen-Dev-Team/going-merry-go/wiki) page.

## 📜 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

We welcome contributions! Feel free to submit issues, fork the repository, and send pull requests.

## 🌟 Show Your Support

Give a ⭐️ if you like this project!

## 📧 Contact

For more information or support, you can reach out to us.
