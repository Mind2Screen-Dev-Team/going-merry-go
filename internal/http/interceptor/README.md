# HTTP Interceptor

This directory contains HTTP interceptors that handle response post-processing.

## Directory Guidelines

- You may create subdirectories for domain-specific interceptors (e.g., `auth/`, `logging/`).
- If there are only a few interceptors, place the `.go` files directly in this directory.

### Example Structure

```bash
interceptor
├── sub-directory                   # when hit a complexity
│   └── interceptor_<group_name>_<some_action>.go
└── interceptor_some_action.go      # when simple implementation
```