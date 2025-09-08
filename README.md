## Social App (Go)

A minimal social API built with Go, featuring a clean repository pattern, chi router, and pluggable Postgres-backed stores.

### Tech Stack
- Go `1.23`
- Router: `github.com/go-chi/chi/v5`
- UUIDs: `github.com/google/uuid`
- Migrations: `golang-migrate`

### Project Structure
```
cmd/
  api/            # HTTP server (routes, handlers, middleware)
  internal/
    env/          # Simple env helpers
    models/       # Domain models + DTOs (User, Post, Session)
    store/        # Repository interfaces + implementations
      errors.go   # Common store errors
      users.go    # UserPostgresStore (pointer-based)
      posts.go    # PostPostgresStore (queries implemented)
      auth.go     # AuthPostgresStore (stubs)
      store.go    # Store interface + PostgresStore embedding sub-stores
migrate/
  migrations/     # SQL migrations (golang-migrate)
docker-compose.yaml
makefile          # migration helpers
```

### Quick Start
1) Requirements: Go 1.23+

2) Environment (example with direnv in `.envrc`):
```
export ADDR=:8080
export USE_DB=true
export DB_ADDR=postgres://admin:adminpassword@localhost:5432/social?sslmode=disable
export DB_MAX_OPEN_CONNS=30
export DB_MAX_IDLE_CONNS=30
export DB_MAX_IDLE_TIME_MIN=15
```

3) Start Postgres (optional via Docker):
```
docker-compose up -d db
```

4) Run the server:
```
go run ./cmd
```
Server listens on `ADDR`.

### API Endpoints (v1)
- Health:
  - GET `/healthcheck`
  - GET `/v1/healthcheck`

- Users:
  - GET  `/v1/users`
  - POST `/v1/users`
    - body: `{ "username": "alice", "email": "alice@example.com", "password": "secret" }`
  - GET    `/v1/users/{userID}`
  - PUT    `/v1/users/{userID}` (update username/email)
  - DELETE `/v1/users/{userID}`

### Example Requests
Create a user:
```
curl -X POST http://localhost:8080/v1/users \
  -H "Content-Type: application/json" \
  -d '{"username":"alice","email":"alice@example.com","password":"secret"}'
```

List users:
```
curl http://localhost:8080/v1/users
```

### Repository Pattern
- Interfaces: `UserStore`, `PostStore`, `AuthStore` (in `cmd/internal/store`)
- Combined interface: `Store` (in `store.go`)
- Implementation: `PostgresStore` embeds `*UserPostgresStore`, `*PostPostgresStore`, `*AuthPostgresStore` and satisfies `Store`.
- Methods use pointer args where needed to support `RETURNING`.

Optional compile-time checks:
```
var _ store.UserStore = (*store.PostgresStore)(nil)
var _ store.PostStore = (*store.PostgresStore)(nil)
```

### Migrations (golang-migrate)
- Files in `cmd/migrate/migrations` with `NNNNNN_name.up.sql` / `.down.sql`
- Makefile helpers:
```
make migrate-create <name>         # create new pair
make migrate-up                    # apply up
make migrate-down                  # rollback 1
make migrate-down-all              # rollback all
make migrate-version               # current version
```
Dirty DB fix (example):
```
migrate -path=cmd/migrate/migrations -database="$DB_ADDR" force <version>
```

### Docker
- Compose file starts Postgres 16 with:
  - DB name: `social`, user: `admin`, password: `adminpassword`
  - Port: `5432` on localhost
  - Persistent volume: `db-data`

Common commands:
```
# start db
docker-compose up -d db

# logs
docker-compose logs -f db

# stop/remove
docker-compose down

# remove container + volume (DANGEROUS)
docker-compose down -v
```

psql connection:
```
psql "postgres://admin:adminpassword@localhost:5432/social?sslmode=disable" -c "\dt"
```

Troubleshooting:
- If migrations say "Dirty database version N": fix SQL, then `force <N-1>` and rerun `make migrate-up`.
- Ensure `DB_ADDR` matches container settings and includes `:5432`.
- If schema changed, re-run migrations or drop and recreate with `docker-compose down -v`.

### TODO: Auth Implementation
- Store layer
  - Implement `AuthPostgresStore` methods: `CreateSession`, `GetSession`, `DeleteSession`
  - Add migrations for `sessions` table (token, user_id FK, expires_at, created_at)
- API layer
  - Routes: `/v1/auth/login`, `/v1/auth/register`, `/v1/auth/logout`
  - Handlers: issue/revoke sessions; return tokens (e.g., HTTP-only cookie or bearer token)
- Security
  - Hash passwords with `bcrypt` on registration; compare on login
  - Use secure, random session tokens (e.g., 32+ bytes base64)
  - Set cookie flags: HttpOnly, Secure (in prod), SameSite
- Middleware
  - `authMiddleware` to extract/validate session and set user context
  - Optional role/permission checks
- Testing
  - Unit tests for auth store + handlers
  - Integration tests: login -> access protected route -> logout


