# Service Attribute

This directory includes attributes used across the service layer.

## Directory Guidelines

- Organize attributes in subdirectories if necessary.
- Add `.go` files directly to this folder for simpler use cases.

### Example Structure

```bash
attribute
├── sub-directory                                # when hit a complexity
│   └── attribute_<group_name>_<some_action>.go
└── attribute_some_action.go                     # when simple implementation
```