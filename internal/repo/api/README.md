# Repository API

This directory holds repository interfaces for data access logic.

## Directory Guidelines

- Create subdirectories for each domain feature if required.
- Directly add `.go` files here if no domain-specific directories are needed.

### Example Structure

```bash
api
├── sub-directory                            # when hit a complexity
│   └── api_<group_name>_<some_action>.go
└── api_some_action.go                      # when simple implementation
```