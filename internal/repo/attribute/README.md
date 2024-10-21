# Repository Attribute

This directory includes attributes related to the repository layer.

## Directory Guidelines

- Use subdirectories for different domain features if applicable.
- Place `.go` files directly here for simpler setups.

### Example Structure

```bash
attribute
├── sub-directory                                 # when hit a complexity
│   └── attribute_<group_name>_<some_action>.go
└── attribute_some_action.go                      # when simple implementation
```