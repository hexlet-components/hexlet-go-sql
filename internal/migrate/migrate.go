package migrate

import (
    "database/sql"
    "embed"

    "github.com/pressly/goose/v3"
)

//go:embed ../../migrations/*.sql
var migrations embed.FS

// Up применяет все доступные миграции.
func Up(db *sql.DB, dialect string) error {
    goose.SetBaseFS(migrations)
    goose.SetLogger(goose.NopLogger())
    if err := goose.SetDialect(dialect); err != nil {
        return err
    }
    return goose.Up(db, "migrations")
}
