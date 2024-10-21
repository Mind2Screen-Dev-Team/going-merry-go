# Provider API

This directory contains the service provider interfaces or APIs.

## Directory Guidelines

- Organize your code by creating subdirectories for each domain feature (e.g., `payment/`).
- Directly place `.go` files here if no subdirectories are needed.

### Example Structure

```bash
api
├── sub-directory                           # when hit a complexity
│   └── api_<group_name>_<some_action>.go
└── api_some_action.go                      # when simple implementation
``` 