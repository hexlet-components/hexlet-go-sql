package config

import (
	"os"
	"time"
)

// Config описывает подключения и таймауты.
type Config struct {
	Driver      string
	Dialect     string
	DSN         string
	Command     string
	Email       string
	Name        string
	CourseTitle string
	UserID      int64
	CourseID    int64
	Timeout     time.Duration
}

// Load читает переменные окружения и возвращает конфигурацию по умолчанию.
func Load() Config {
	driver := os.Getenv("DB_DRIVER")
	if driver == "" {
		driver = "pgx"
	}

	dialect := os.Getenv("DB_DIALECT")
	if dialect == "" {
		dialect = "postgres"
	}

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "postgres://app:secret@localhost:6543/app?sslmode=disable&application_name=hexlet-go-sql"
	}

	return Config{
		Driver:  driver,
		Dialect: dialect,
		DSN:     dsn,
		Timeout: 5 * time.Second,
	}
}
