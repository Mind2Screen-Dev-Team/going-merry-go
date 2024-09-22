# Golang Skeleton
A Go project starter kit with essential tools and structures.

## Plan

### App
- Run All APP via Command Line - (Not Yet)
- APP Scheduler Cron Job - (Not Yet)
- APP Worker Pub / Sub - (Not Yet)
- HTTP Rest API - (Done)

### Todo
- Logging
- Makefile Runner
- Docker Integration
- Open Telemetry Integration
- CMD Generator handler, middleware, router, repo and service
- Setup Advance Generator Configuration: yaml -> pkl -> yaml, environment change: dev | prod
- Unit Test
- Open API Generator Docs
- Cron Job Scheduler APP
- Worker Pub/Sub APP
- GRPC API APP

### Done
- Base Structural Directory
- Setup Basic Generator Configuration
- Registry Dependency, Repository and Service
- HTTP Handler and Router Loader
- DTO Validation
- DB migrations and seeders

## Tips

### Go Wrk for API Load Test

- go install github.com/tsliwowicz/go-wrk@latest
- Go Wrk Running
```bash
go-wrk -c 5000 -d 5 -M "POST" -body "{\"email\": \"example@gmail.com\", \"password\": \"secret\"}" http://localhost:8081/api/v1/auth/login
```

### Golang Build

- Go Build Rest Api App

    ```bash
    go build -o ./bin/restapi ./cmd/restapi/main.go
    ```

### DB Migration Tools
- Goose DB Migrations, Source: https://github.com/pressly/goose

### Align Struct Tools

- Aligo Article, Source: https://www.freedium.cfd/https://medium.com/@sddkal/use-betteralign-to-optimize-go-memory-consumption-3736a3172860

- Install aligo for align size struct
    ```bash
    go install github.com/essentialkaos/aligo/v2@latest
    ```

### PKL Config Generator Tools

- https://pkl-lang.org/go/current/quickstart.html
- Edit `.zshrc`, Add Golang Bin into PATH `export PATH=$PATH:"$HOME/go/bin"` for Mac OS and when using brew to install `go`.

#### PKL Command-Line
- Generate Application Config Code
    ```bash
    pkl-gen-go ./pkl/AppConfig.pkl --base-path github.com/Mind2Screen-Dev-Team/go-skeleton
    ```

- Generate `.pkl` file configuration
    ```bash
    pkl eval pkl/AppConfig.pkl
    ```

- Generate ouput `.pkl` file configuration
    ```bash
    pkl eval -o ./pkl/config/example.pkl pkl/AppConfig.pkl
    ```

- Generate ouput `.yaml` file configuration
    ```bash
    pkl eval -f yaml -o ./application.yaml pkl/AppConfig.pkl
    ```