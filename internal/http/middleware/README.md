# HTTP Middleware

This directory contains HTTP interceptors that handle request and response pre-processing.

## Directory Guidelines

- You may create subdirectories for domain-specific interceptors (e.g., `auth/`, `logging/`).
- If there are only a few interceptors, place the `.go` files directly in this directory.

### Example Structure

```bash
middleware
├── sub-directory                   # when hit a complexity
│   └── middleware_<group_name>_<some_action>.go
└── middleware_some_action.go       # when simple implementation
```