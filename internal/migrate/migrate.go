package migrate

import (
	"database/sql"

	"github.com/pressly/goose/v3"

	projectMigrations "github.com/hexlet-components/hexlet-go-sql/migrations"
)

// Up applies all migrations using goose, relying on the embedded migration files.
func Up(db *sql.DB, dialect string) error {
	goose.SetBaseFS(projectMigrations.FS)
	goose.SetLogger(goose.NopLogger())
	if dialect == "pgx" {
		dialect = "postgres"
	}
	if err := goose.SetDialect(dialect); err != nil {
		return err
	}
	return goose.Up(db, "migrations")
}
