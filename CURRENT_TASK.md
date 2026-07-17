# TASK

Version: 1.0

---

# Current Milestone

Authentication Foundation

---

# Objective

Membangun pondasi Authentication.

Belum ada Login.

Belum ada Register.

Belum ada JWT Middleware.

Belum ada endpoint.

Belum ada repository implementation.

Belum ada business logic.

---

# Requirements

1. Buat module `internal/modules/auth/` dengan struktur:
   - service/
   - repository/
   - handler/
   - routes/
   - dto/
   - mapper/
   - validator/
   - docs/

2. Repository interface:
   - FindUserByEmail
   - SaveRefreshToken
   - GetRefreshToken
   - DeleteRefreshToken

3. Service interface:
   - Register
   - Login
   - Logout
   - RefreshToken
   - ValidateToken

4. DTO (skeleton):
   - LoginRequest
   - RegisterRequest
   - RefreshTokenRequest
   - LoginResponse
   - RefreshTokenResponse

5. Handler (skeleton)

6. Routes (skeleton, belum register)

---

# Out of Scope

Jangan membuat:

- Business Logic
- Repository Implementation
- GORM
- SQL
- Migration
- Seeder
- JWT Library
- Middleware
- Route Registration
- Login/Register/Refresh Endpoint

---

# Validation

- go mod tidy
- go fmt ./...
- go vet ./...
- go build ./...

---

# Completion Report

1. File dibuat
2. Struktur module
3. Contracts
4. DTO
5. Hasil build
6. Blocker (jika ada)
