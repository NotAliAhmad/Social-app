## Social App (Go)

A minimal social API built with Go, featuring a clean repository pattern, chi router, and pluggable data stores (in-memory and Postgres-ready).

### Tech Stack
- Go `1.23`
- Router: `github.com/go-chi/chi/v5`
- UUIDs: `github.com/google/uuid`

### Project Structure
```
cmd/
  api/            # HTTP server (routes, handlers, middleware)
  internal/
    env/          # Simple env helpers
    models/       # Shared domain models (User, Post, Session, DTOs)
    store/        # Repository interfaces + implementations
      errors.go   # Common store errors
      memory.go   # In-memory Store (default for local/dev)
      users.go    # PostgresStore user methods
      posts.go    # PostgresStore post methods (stubs)
      auth.go     # PostgresStore session methods (stubs)
      store.go    # Store composition + PostgresStore type
```

### Quick Start
1) Requirements: Go 1.23+

2) Configure env (optional):
```
ADDR=:8080
```
You can set this via your shell or using `.envrc` if you use `direnv`.

3) Run the server:
```
go run ./cmd
```

Server listens on `ADDR` (default `:8080`).

### API Endpoints (v1)
- Health:
  - GET `/healthcheck`
  - GET `/v1/healthcheck`

- Users:
  - GET  `/v1/users/`
  - POST `/v1/users/`
    - body: `{ "firstName": "John", "lastName": "Doe" }`
  - GET    `/v1/users/{userID}`
  - PUT    `/v1/users/{userID}`
  - DELETE `/v1/users/{userID}`

Posts and Auth endpoints are scaffolded and can be extended.

### Example Requests
Create a user:
```
curl -X POST http://localhost:8080/v1/users/ \
  -H "Content-Type: application/json" \
  -d '{"firstName":"Jane","lastName":"Doe"}'
```

List users:
```
curl http://localhost:8080/v1/users/
```

### Repository Pattern
The server depends on small, focused interfaces:
- `UsersStore` and `PostsStore` (in `cmd/internal/store/users.go` and `posts.go`)
- Combined as `Store` in `cmd/internal/store/store.go`

Implementations:
- `MemoryStore` (thread-safe, great for dev/tests)
- `PostgresStore` (methods for users implemented; posts/auth are stubs)

Compile-time guarantees (optional pattern):
```
var _ store.UsersStore = (*store.PostgresStore)(nil)
var _ store.PostsStore = (*store.PostgresStore)(nil)
```

### Switching Stores
By default, `main.go` wires the in-memory store:
```
memoryStore := store.NewMemoryStore()
app, _ := api.NewServer(cfg, memoryStore)
```

To use Postgres, create a DB, open a connection, and inject it:
```
db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
if err != nil { log.Fatal(err) }

pg := store.NewPostgresStore(db)
app, _ := api.NewServer(cfg, pg)
```
Note: Postgres methods for posts/auth are placeholders—add migrations and implement queries before enabling them.

### Development
- Lint/format with your editor’s Go tools.
- Handlers log server-side errors for easier debugging.

### Notes
- Models live in `cmd/internal/models` to avoid import cycles.
- Handlers use `context.Context` throughout and return proper HTTP codes.


