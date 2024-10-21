# DTO (Data Transfer Object)

This directory contains Data Transfer Objects (DTOs) for HTTP requests and responses.

## Directory Guidelines

- You can create a subdirectory for each domain feature (e.g., `user/`, `order/`) to group related DTOs.
- If the domain is simple, you may directly add the `.go` files in this directory.

### Example Structure

```bash
dto
├── sub-directory           # when hit a complexity
│   └── dto_<group_name>_<some_action>.go
└── dto_some_action.go      # when simple implementation
```