This repository contains the reference solution that emerges if you complete every self-study assignment in the course. It demonstrates how to combine PostgreSQL, `sqlc`, goose migrations, CLI tooling, and integration tests in a single Go application.

## Features

- CLI commands for managing users, courses, and course membership.
- Prepared statements and transactions for bulk workflows.
- Repository layer (hand-written) on top of `sqlc` queries.
- goose migrations covering users, courses, lessons, and a many-to-many membership table.
- Environment-driven configuration (`DB_DRIVER`, `DB_DSN`, etc.).

## Prerequisites

- Go 1.20+.
- `sqlc` (`go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`).
- `goose` (`go install github.com/pressly/goose/v3/cmd/goose@latest`).
- Docker + Docker Compose (for running PostgreSQL locally).
- GNU Make (optional, but all commands are provided as make targets).
- Copy `.env.example` to `.env` if you want to override defaults (the app loads `.env` automatically via `godotenv`).

## Getting Started

### PostgreSQL via Docker Compose

```bash
docker compose up -d db
cp .env.example .env                  # optional, overrides env vars
export DB_DRIVER=pgx
export DB_DSN="postgres://app:secret@localhost:6543/app?sslmode=disable&application_name=hexlet-go-sql"
make migrate
make run CMD=course-create COURSE="Graph Theory"
```

Use `make rollback` to undo the latest migration and `make sqlc` after editing SQL files.

## Project Layout

- `cmd/app` – CLI entrypoint.
- `internal/config` – configuration loader and environment variables.
- `internal/repository` – data-access layer and transaction helpers.
- `internal/migrate` – goose bootstrapper.
- `migrations` – SQL migrations consumed by goose.
- `query` & `sqlc.yaml` – `sqlc` configuration and query files.
- `docker-compose.yml` – local PostgreSQL instance for development.

## Integration Tests

```bash
make test-integration
```

Tests run against the DSN you configure (default: local PostgreSQL). Use the `integration` build tag so production data stays untouched.
