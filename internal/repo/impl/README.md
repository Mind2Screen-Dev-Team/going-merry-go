# Repository Implementation

This directory contains the implementations of repository interfaces.

## Directory Guidelines

- Separate implementations into subdirectories by domain features (e.g., `product/`).
- For simpler domains, add the `.go` files directly in this directory.

### Example Structure

```bash
impl
├── sub-directory                           # when hit a complexity
│   └── impl_<group_name>_<some_action>.go
└── impl_some_action.go                     # when simple implementation
```