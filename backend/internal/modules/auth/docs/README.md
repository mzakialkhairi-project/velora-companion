# Auth Module

This module handles authentication operations.

## Structure

```
auth/
├── repository/   # AuthRepository interface
├── service/      # AuthService interface
├── dto/          # Data transfer objects
├── handler/      # HTTP handlers (skeleton)
├── routes/       # Route definitions (skeleton)
├── mapper/       # Mapping functions
├── validator/    # Validation functions (skeleton)
└── docs/         # Documentation
```

## Contracts

### Repository Interface

```go
type AuthRepository interface {
    FindUserByEmail(ctx, email)
    SaveRefreshToken(ctx, userID, token, expiresAt)
    GetRefreshToken(ctx, token)
    DeleteRefreshToken(ctx, token)
    DeleteUserRefreshTokens(ctx, userID)
}
```

### Service Interface

```go
type AuthService interface {
    Register(ctx, req)
    Login(ctx, req)
    Logout(ctx, refreshToken)
    RefreshToken(ctx, req)
    ValidateToken(ctx, token)
}
```

## TODO

- [ ] Implement repository
- [ ] Implement service
- [ ] Add password hashing (bcrypt/argon2)
- [ ] Add JWT token generation
- [ ] Add validation
- [ ] Register routes
- [ ] Add unit tests
