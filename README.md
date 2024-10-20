
# 🚢 _Going-Merry-Go_ - Project Skeleton

Welcome aboard the **Going-Merry-Go**! Inspired by the iconic ship from the anime One Piece, this project is a robust and flexible Go project starter kit. It's designed to help you quickly set up your Go applications with the right structure and essential tools.

## 🗂 Project Structure

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
