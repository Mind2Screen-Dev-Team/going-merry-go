# Provider Implementation

This directory contains the implementations of service providers.

## Directory Guidelines

- Use subdirectories to separate implementations for different domain features (e.g., `user/`).
- Directly place `.go` files here if only a few implementations exist.

### Example Structure

```bash
impl
├── sub-directory                            # when hit a complexity
│   └── impl_<group_name>_<some_action>.go
└── impl_some_action.go                      # when simple implementation
```