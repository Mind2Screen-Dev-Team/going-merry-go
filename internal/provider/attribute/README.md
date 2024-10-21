# Provider Attribute

This directory includes domain-specific attributes related to providers.

## Directory Guidelines

- Create subdirectories for each domain if necessary.
- Place files directly in this directory if there are only a few attributes.

### Example Structure

```bash
attribute
├── sub-directory                                 # when hit a complexity
│   └── attribute_<group_name>_<some_action>.go
└── attribute_some_action.go                      # when simple implementation
``` 