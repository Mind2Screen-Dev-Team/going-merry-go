# HTTP Handler

This directory contains handlers that process incoming HTTP requests.

## Directory Guidelines

- Organize handlers by creating subdirectories for each domain feature (e.g., `user/`, `product/`).
- If a domain is not complex, place the handler `.go` files directly here.

### Example Structure

```bash
handler
├── sub-directory               # when hit a complexity
│   └── handler_<group_name>_<some_action>.go
└── handler_some_action.go      # when simple implementation
```