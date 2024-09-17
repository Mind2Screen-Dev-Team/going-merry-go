# Golang Skeleton
A Go project starter kit with essential tools and structures.

## Plan

### App / Service
- Run All APP via Command Line - (Not Yet)
- APP Scheduler Cron Job - (Not Yet)
- APP Worker Pub / Sub - (Not Yet)
- HTTP Rest API - (On Progress)

### Todo
- DTO Validation
- Logging
- DB migrations and seeders
- Makefile Runner
- Docker Integration
- Open Telemetry Integration
- Example Rest API APP
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

## Tips

### Align Struct Tools
- https://www.freedium.cfd/https://medium.com/@sddkal/use-betteralign-to-optimize-go-memory-consumption-3736a3172860
- go install github.com/essentialkaos/aligo/v2@latest

### PKL Config Generator Tools
- https://pkl-lang.org/go/current/quickstart.html
- Edit `.zshrc`, Add Golang Bin into PATH `export PATH=$PATH:"$HOME/go/bin"` for Mac OS and when using brew to install `go`.
- pkl-gen-go ./pkl/AppConfig.pkl --base-path github.com/Mind2Screen-Dev-Team/go-skeleton
