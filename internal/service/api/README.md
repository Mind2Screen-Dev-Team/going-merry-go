# Service API

This directory holds service interfaces that define the business logic layer.

## Directory Guidelines

- Create domain-specific subdirectories for complex implementations (e.g., `order/`).
- Directly add service interfaces as `.go` files if not creating subdirectories.

### Example Structure

```bash
api
├── sub-directory                           # when hit a complexity
│   └── api_<group_name>_<some_action>.go
└── api_some_action.go                      # when simple implementation
```