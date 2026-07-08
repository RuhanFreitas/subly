# Subly — Backend API

REST API for managing user accounts and their subscriptions (streaming services, SaaS tools, etc.). Users register, log in, and perform CRUD operations on their subscription list. Authentication is handled with JWT stored in an HttpOnly cookie.

---

## What it does

Subly tracks recurring subscriptions for a user. The backend exposes a JSON HTTP API that lets a client:

- Create an account and log in (JWT cookie session).
- Read, update, and delete the user account.
- Create, list, read, update, and delete subscriptions that belong to the authenticated user.

Each subscription stores a name, price, start date, next payment date, and renewal cadence (`daily`, `weekly`, `monthly`, `yearly`).

---

## Stack

| Layer | Technology |
|-------|------------|
| Language | Go `1.26.3` |
| HTTP framework | [Gin](https://github.com/gin-gonic/gin) `v1.12.0` |
| Database | PostgreSQL `16` |
| DB driver | [pgx/v5](https://github.com/jackc/pgx) (`pgxpool` connection pool), no ORM |
| Auth | [golang-jwt/jwt](https://github.com/golang-jwt/jwt) `v3` (HS256) |
| Password hashing | `golang.org/x/crypto/bcrypt` |
| Config | [joho/godotenv](https://github.com/joho/godotenv) (`.env` file) |
| Migrations | [golang-migrate](https://github.com/golang-migrate/migrate) CLI (external, run manually) |
| Local infra | Docker Compose (PostgreSQL only) |

Module name: `subly`.

---

## Architecture

The backend is a flat, two-tier design. There is **no service layer** — HTTP handlers call the repository layer directly. Business logic (validation, JWT creation, bcrypt, partial-update merging) lives inside the handlers.

```
HTTP Request
   -> Gin router (Logger + Recovery middleware)
      -> AuthMiddleware (JWT cookie) on protected routes
         -> Handler (internal/handler)
            -> Repository (internal/repository, raw SQL via pgx)
               -> PostgreSQL
```

```mermaid
flowchart TB
    Client[HTTP Client / Web Frontend]
    Gin[Gin Router - gin.Default]
    AuthMW[AuthMiddleware - JWT Cookie]
    Handlers[internal/handler]
    Repos[internal/repository]
    PG[(PostgreSQL 16)]

    Client --> Gin
    Gin -->|Logger + Recovery| Handlers
    Gin -->|protected routes| AuthMW
    AuthMW --> Handlers
    Handlers --> Repos
    Repos --> PG
```

---

## Folder structure

```
server/
├── .env.example                    # Env var template
├── .gitignore
├── compose.yml                     # Docker Compose (PostgreSQL only)
├── go.mod / go.sum                 # Module & dependencies
├── cmd/
│   └── server/
│       └── main.go                 # Entry point: config, DB, router, routes
├── internal/
│   ├── config/
│   │   └── config.go               # Loads .env into a Config struct
│   ├── database/
│   │   └── postgres.go             # pgxpool connection + ping
│   ├── handler/
│   │   ├── auth_handler.go         # Register, Login
│   │   ├── user_handler.go         # User CRUD
│   │   └── subscription_handler.go # Subscription CRUD
│   ├── middleware/
│   │   └── auth_middleware.go      # JWT cookie validation (package "midleware")
│   ├── model/
│   │   ├── user.go                 # User entity
│   │   └── subscription.go         # Subscription entity
│   └── repository/
│       ├── user_repository.go      # User SQL queries
│       └── subscription_repository.go # Subscription SQL queries
└── migrations/
    ├── 000001_create_users_table.up.sql
    ├── 000001_create_users_table.down.sql
    ├── 000002_create_subscriptions_table.up.sql
    └── 000002_create_subscriptions_table.down.sql
```

**Not present:** service layer, tests (`*_test.go`), Dockerfile for the app, Makefile, CORS, external integrations (payments, email).

### What each package does

- **`cmd/server`** — Application entry point. Loads config, opens the DB pool, builds the Gin router, registers all routes, and starts listening.
- **`internal/config`** — Reads `.env` via godotenv and exposes `Port`, `DatabaseURL`, `JWTSecret`.
- **`internal/database`** — Parses the connection string, creates a `pgxpool.Pool`, and pings the DB to confirm connectivity.
- **`internal/handler`** — Gin handler factories. Each function returns a `gin.HandlerFunc` and receives the DB pool (and config where auth is needed). Parses/validates input, applies logic, calls repositories, writes JSON responses.
- **`internal/middleware`** — `AuthMiddleware` reads the `auth_token` cookie, verifies the JWT, and stores `userID` and `email` in the Gin context for downstream handlers.
- **`internal/model`** — Plain structs mapping DB rows to JSON (`User`, `Subscription`).
- **`internal/repository`** — Data access. Raw parameterized SQL executed through pgx, each query wrapped in a 5-second context timeout.

---

## How it works

### Bootstrap (`cmd/server/main.go`)

1. `config.Load()` reads `.env`. Fatal if it fails.
2. `database.Connect(cfg.DatabaseURL)` opens a `pgxpool.Pool` and pings PostgreSQL. Fatal if it fails.
3. `gin.Default()` creates a router with Logger + Recovery middleware.
4. Routes are registered (health check, auth, user, subscription).
5. `router.Run(":" + cfg.Port)` starts the HTTP server.

Migrations are **not** run at startup — they are applied manually with the `migrate` CLI.

### Authentication flow

1. **Register / Login** validate the request and (for login) verify the password with bcrypt.
2. A JWT is signed with HS256 using `JWT_SECRET`. Claims: `user_id`, `email`, `exp` (24h).
3. The token is returned in an HttpOnly, Secure cookie named `auth_token` (MaxAge ~24 days).
4. On protected routes, `AuthMiddleware` reads the cookie, verifies the signature and validity, then puts `userID` and `email` into the request context.
5. Handlers read `userID` from the context (e.g. to scope subscriptions to the current user).

Passwords are hashed with bcrypt (`DefaultCost`) and never returned in responses (the `Password` field is cleared before serialization).

---

## API reference

Base URL: `http://localhost:<PORT>` (default `8081`).

All request/response bodies are JSON. Protected routes require the `auth_token` cookie (set automatically on register/login).

### Public routes

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/healthy` | Health check. Returns `{ "Health": "The API is healthy!" }`. |
| `POST` | `/auth/register` | Create an account, set auth cookie, return the user. |
| `POST` | `/auth/login` | Authenticate, set auth cookie, return the user. |

### Protected routes (require `auth_token` cookie)

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/user/:id` | Get a user by ID. |
| `PATCH` | `/user/:id` | Partially update a user. |
| `DELETE` | `/user/:id` | Delete a user. |
| `POST` | `/subscription` | Create a subscription for the authenticated user. |
| `GET` | `/subscription` | List all subscriptions of the authenticated user. |
| `GET` | `/subscription/:id` | Get a subscription by ID. |
| `PATCH` | `/subscription/:id` | Partially update a subscription. |
| `DELETE` | `/subscription/:id` | Delete a subscription. |

---

### Endpoint details

#### `POST /auth/register`

Request body:

```json
{
  "first_name": "Ada",
  "last_name": "Lovelace",
  "email": "ada@example.com",
  "password": "secret123"
}
```

Rules: all fields required; password must be at least 6 characters.

Responses:
- `201 Created` — `{ "user": { ...without password } }`, sets `auth_token` cookie.
- `400 Bad Request` — invalid body, password too short, or email already registered.
- `500 Internal Server Error` — hashing, DB, or token error.

#### `POST /auth/login`

Request body:

```json
{ "email": "ada@example.com", "password": "secret123" }
```

Responses:
- `200 OK` — `{ "user": { ... } }`, sets `auth_token` cookie.
- `401 Unauthorized` — invalid body or user not found.
- `400 Bad Request` — invalid credentials (password mismatch).

#### `GET /user/:id`

- `200 OK` — `{ "user": { ... } }`.
- `401` — non-numeric `:id`.
- `500` — user not found / DB error.

#### `PATCH /user/:id`

Body (all fields optional, at least one required):

```json
{ "first_name": "Ada", "last_name": "Byron", "email": "ada@new.com", "password": "newpass" }
```

- `200 OK` — `{ "user": { ...updated } }`.
- `400` — invalid body or no fields provided.
- `404` — user not found.

#### `DELETE /user/:id`

- `200 OK` — `{ "message": "Account deleted successfully" }`.
- `400` — non-numeric `:id`.
- `500` — DB error.

#### `POST /subscription`

The owner is taken from the JWT (`userID`), not the body.

Body:

```json
{
  "name": "Netflix",
  "price": 39.90,
  "starting_date": "2026-01-01T00:00:00Z",
  "payment_date": "2026-02-01T00:00:00Z",
  "subscription_renew": "monthly"
}
```

`subscription_renew` must be one of `daily`, `weekly`, `monthly`, `yearly`.

- `201 Created` — `{ "susbcription": { ... } }` (note the response key is misspelled in the code).
- `400` — invalid body.
- `500` — DB error.

#### `GET /subscription`

Returns all subscriptions for the authenticated user.

- `200 OK` — `{ "subscriptions": [ ... ] }`.

#### `GET /subscription/:id`

- `200 OK` — `{ "subscription": { ... } }`.
- `404` — no subscription found.
- `400` — invalid `:id`.

#### `PATCH /subscription/:id`

Body (all fields optional):

```json
{ "name": "Netflix Premium", "price": 55.90, "is_active": false, "subscription_renew": "yearly" }
```

- `200 OK` — `{ "subscription": { ...updated } }`.
- `400` — invalid `:id`, invalid body, or subscription not found.

#### `DELETE /subscription/:id`

- `200 OK` — `{ "message": "Subscription deleted successfully" }`.
- `400` — invalid `:id`.
- `500` — DB error.

---

## Data model

### `users` table

```sql
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(120) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### `subscriptions` table

```sql
CREATE TABLE IF NOT EXISTS subscriptions (
    id SERIAL PRIMARY KEY NOT NULL,
    user_id INTEGER NOT NULL,
    name VARCHAR(50) NOT NULL,
    price NUMERIC NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    starting_date TIMESTAMP NOT NULL,
    payment_date TIMESTAMP NOT NULL,
    subscription_renew TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_user
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,

    CONSTRAINT chk_subscription_renew
        CHECK (subscription_renew IN ('daily', 'weekly', 'monthly', 'yearly'))
);
```

Deleting a user cascades and deletes their subscriptions.

### Go entities

`model.User` and `model.Subscription` mirror these tables with `json` and `db` struct tags. Note: `id` and `user_id` are integers in the DB but typed as `string` in the Go structs.

---

## Configuration

Config comes from a `.env` file loaded at startup (see `.env.example`).

| Variable | Example | Purpose |
|----------|---------|---------|
| `PORT` | `8081` | HTTP listen port. |
| `DATABASE_URL` | `postgresql://appuser:apppassword@localhost:15432/appdb?sslmode=disable` | PostgreSQL connection string. |
| `JWT_SECRET` | `some-long-random-secret` | HMAC secret for signing/verifying JWTs. |

The `.env` file is gitignored and must exist at runtime — startup fails if it is missing. There is no validation for empty values.

---

## Running locally

### Prerequisites

- Go `1.26.3+`
- Docker (for PostgreSQL)
- `migrate` CLI ([golang-migrate](https://github.com/golang-migrate/migrate))

### Steps

1. Copy the env template and fill in values:

```bash
cp .env.example .env
```

2. Start PostgreSQL:

```bash
docker compose up -d
```

This runs `postgres:16` on host port `15432` with user `appuser`, password `apppassword`, database `appdb`.

3. Apply migrations:

```bash
migrate -path migrations -database "postgresql://appuser:apppassword@localhost:15432/appdb?sslmode=disable" up
```

4. Run the server:

```bash
go run ./cmd/server
```

The API listens on `http://localhost:8081` (or your `PORT`).

### Docker Compose service

```yaml
services:
  db:
    image: postgres:16
    container_name: subly_postgres
    restart: unless-stopped
    environment:
      POSTGRES_USER: appuser
      POSTGRES_PASSWORD: apppassword
      POSTGRES_DB: appdb
    ports:
      - "15432:5432"
    volumes:
      - subly_postgres_data:/var/lib/postgresql/data

volumes:
  subly_postgres_data:
```

Only the database runs in Compose; the Go app runs on the host.

---

## Middleware

| Middleware | Source | Scope |
|------------|--------|-------|
| Logger | Gin built-in (`gin.Default()`) | All routes |
| Recovery | Gin built-in (`gin.Default()`) | All routes |
| `AuthMiddleware(cfg)` | `internal/middleware` | All `/user/*` and `/subscription/*` routes |

`AuthMiddleware` aborts with `401` if the `auth_token` cookie is missing, the token is invalid, or the claims cannot be parsed. There is no CORS, rate limiting, or role-based authorization.

---

## Known issues / limitations

These are real inconsistencies in the current code, documented for accuracy:

1. **No resource ownership checks.** Any authenticated user can read/update/delete any user or subscription by ID. `GET /user/:id` does not compare the JWT `userID` to `:id`, and subscription-by-id routes don't verify ownership.
2. **Login route missing leading slash.** Registered as `"auth/login"` instead of `"/auth/login"`.
3. **Inverted password logic in `UpdateUserByID`.** The bcrypt hashing runs in the `else` branch (when no password is provided), which will panic/dereference a nil pointer. A provided password is stored without hashing.
4. **SQL parameter bug in `UpdateUserByID` repository.** The `WHERE` clause reuses `$4` for both password and id.
5. **SQL parameter collision in `UpdateSubscriptionByID`.** `WHERE id = $6` collides with `subscription_renew = $6`.
6. **`is_active` inconsistencies.** Some subscription SELECT/RETURNING statements omit `is_active`, and `GetAllSubscriptions` scans fewer columns than `SELECT *` returns.
7. **Type mismatch.** `id`/`user_id` are integers in the DB but `string` in the Go models and JWT claims.
8. **Response typo.** `POST /subscription` returns the key `"susbcription"`.
9. **Package name typo.** The middleware package is named `midleware`.
10. **No CORS** configured, which will block browser clients on a different origin.
11. **No automated tests.**
