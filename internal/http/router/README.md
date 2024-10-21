# HTTP Router

This directory handles routing logic and endpoint definitions.

## Directory Guidelines

- For complex domain-specific routers, create subdirectories.
- For simpler setups, place the routing `.go` files directly in this directory.

### Example Structure

```bash
router
├── sub-directory               # when hit a complexity
│   └── router_<group_name>_<some_action>.go
└── router_some_action.go       # when simple implementation
``` 