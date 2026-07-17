# User Module

This module handles user management operations.

## Structure

```
user/
├── entity/       # User domain entity
├── repository/   # User repository interface
├── service/      # User service interface
├── dto/          # Data transfer objects
├── handler/      # HTTP handlers (skeleton)
├── routes/       # Route definitions (skeleton)
├── mapper/       # Entity to DTO mapping
├── validator/    # Request validation (skeleton)
└── docs/         # Documentation
```

## Entity

The User entity contains:
- ID (from BaseEntity)
- Name
- Email
- PasswordHash
- Status (active, inactive, banned)
- CreatedAt, UpdatedAt, DeletedAt (from BaseEntity)

## TODO

- [ ] Implement repository
- [ ] Implement service
- [ ] Add validation
- [ ] Register routes
- [ ] Add unit tests
