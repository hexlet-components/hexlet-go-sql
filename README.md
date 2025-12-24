This repository contains the reference solution that emerges if you complete every self-study assignment in the course. It demonstrates how to combine SQLite/PostgreSQL, `sqlc`, goose migrations, CLI tooling, and integration tests in a single Go application.

## Features

- CLI commands for managing users, courses, and course membership.
- Prepared statements and transactions for bulk workflows.
- Repository layer (hand-written) on top of `sqlc` queries.
- goose migrations covering users, courses, lessons, reviews, and many-to-many relations.
- Environment-driven configuration (`DB_DRIVER`, `DB_DSN`, etc.).

## Prerequisites

- Go 1.24+.
- `sqlc` (`go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`).
- `goose` (`go install github.com/pressly/goose/v3/cmd/goose@latest`).
- Docker + Docker Compose (for running PostgreSQL locally).
- GNU Make (optional, but all commands are provided as make targets).

## Getting Started

### Local SQLite

```bash
make migrate                          # apply goose migrations to SQLite
make run CMD=user-create EMAIL=test@example.com NAME=Demo
```

### PostgreSQL via Docker Compose

```bash
docker compose up -d db
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

Tests run against an in-memory SQLite database (or your configured DSN) with the `integration` build tag, so the production database stays untouched.
