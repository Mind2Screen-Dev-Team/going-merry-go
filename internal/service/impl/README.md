# Service Implementation

This directory contains the implementations of the service interfaces.

## Directory Guidelines

- For clarity, separate the implementations into subdirectories by domain feature (e.g., `order/`).
- Directly place `.go` files here if the domain structure is straightforward.

### Example Structure

```bash
impl
├── sub-directory                           # when hit a complexity
│   └── impl_<group_name>_<some_action>.go
└── impl_some_action.go                     # when simple implementation
```