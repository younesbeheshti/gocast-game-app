# GoCast Game

A Go backend for a real-time, category-based trivia/quiz matching game. Users register, log in, get matched into a game within a category (e.g. football), and answer questions. The service also exposes a backoffice API with role/permission-based access control for admin operations.

## Features

- **User accounts** — registration and login with phone number + password, JWT-based access/refresh tokens
- **Matching** — players join a waiting list for a category and get matched into a game
- **Games & questions** — categories, questions with multiple-choice answers, difficulty levels, and per-player scoring
- **Access control** — role- and user-based permission system (RBAC) guarding backoffice endpoints
- **Backoffice API** — admin endpoint to list users, protected by auth + permission middleware
- **Config via YAML + env vars** — powered by [koanf](https://github.com/knadh/koanf), overridable with `GAMEAPP_*` environment variables
- **PostgreSQL** for persistent data, **Redis** for the matching waiting list
- **SQL migrations** managed with [sql-migrate](https://github.com/rubenv/sql-migrate), run automatically on startup

## Tech Stack

| Layer            | Technology                                   |
|-------------------|-----------------------------------------------|
| HTTP framework    | [Echo v5](https://github.com/labstack/echo)   |
| Database          | PostgreSQL (via `lib/pq`)                     |
| Cache / queue     | Redis (`redis/go-redis/v9`)                   |
| Config            | `knadh/koanf`                                 |
| Auth              | JWT (`golang-jwt/jwt/v5`, `labstack/echo-jwt`) |
| Validation        | `go-ozzo/ozzo-validation`                     |
| Migrations        | `rubenv/sql-migrate`                          |

## Project Structure

```
.
├── main.go                     # entry point: loads config, runs migrations, wires services, starts server
├── config.yml                  # default local configuration
├── docker-compose.yml          # Postgres + Redis for local development
├── config/                     # config loading, defaults, constants
├── entity/                     # core domain types (User, Game, Question, Role, Permission, ...)
├── param/                      # request/response DTOs
├── adapter/redis/              # low-level Redis client adapter
├── repository/
│   ├── postgres/               # Postgres repositories (user, access control) + migrations
│   ├── redis/redismatching/    # Redis-backed matching waiting list
│   └── migrator/               # runs SQL migrations on startup
├── service/
│   ├── authservice/            # JWT issuing/parsing
│   ├── userservice/            # register, login, profile
│   ├── backofficeuserservice/  # admin user operations
│   ├── authorizationservice/   # permission checks
│   └── matchingservice/        # waiting list / matching logic
├── validator/                  # request validation (user, matching)
├── delivery/httpserver/        # Echo server, routes, handlers, middleware
└── pkg/                        # shared helpers (claims, http messages, rich errors, slices)
```

## Getting Started

### Prerequisites

- Go 1.25+
- Docker (for Postgres and Redis) or your own local instances

### 1. Start dependencies

```bash
docker-compose up -d
```

This starts:
- **Postgres** on `localhost:5432` (db `gocast_game`, user/pass `postgres`/`postgres`)
- **Redis** on `localhost:6379`

### 2. Configure

Edit `config.yml` if needed, or override any value with environment variables prefixed `GAMEAPP_` (e.g. `GAMEAPP_POSTGRES_HOST=db`).

### 3. Run the app

```bash
go run main.go
```

On startup the app loads config, runs pending Postgres migrations automatically, wires up services, and starts the HTTP server (default port `8080`).

## API Endpoints

| Method | Path                          | Auth required        | Description                     |
|--------|-------------------------------|-----------------------|----------------------------------|
| GET    | `/health_check`               | –                     | Health check                    |
| POST   | `/users/register`             | –                     | Register a new user             |
| POST   | `/users/login`                | –                     | Log in, returns tokens          |
| GET    | `/users/userprofile`          | JWT                   | Get current user's profile      |
| POST   | `/matching/add-to-waiting-list` | JWT                 | Join matchmaking for a category |
| GET    | `/backoffice/users/`          | JWT + `user-list` permission | List users (admin)       |

## Configuration Reference

Configuration is loaded from `config.yml` and can be overridden via `GAMEAPP_*` environment variables (dots become underscores, e.g. `postgres.host` → `GAMEAPP_POSTGRES_HOST`).

```yaml
auth:
  sign_key: secret

http_server:
  port: 8080

postgres:
  port: 5432
  host: localhost
  db_name: postgres
  username: postgres
  password: postgres
  sslmode: disable

redis:
  host: localhost
  port: 6379
  password: ""
  db: 0

matching_service:
  waiting_timeout: "2m"
```

## Database Migrations

Migrations live in `repository/postgres/migrations/` and are applied automatically at startup via the migrator. They currently cover:
- User table creation
- Password column
- Role column
- Permission table
- Access control list

## Status

This project is under active development — expect incomplete pieces (e.g. game creation from matched players, full question/answer flow) and some TODOs in the code.

## License

No license specified yet.