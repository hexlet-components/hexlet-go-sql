package config

import (
    "os"
    "time"
)

// Config описывает подключения и таймауты.
type Config struct {
    Driver        string
    DSN           string
    Command       string
    Email         string
    Name          string
    CourseTitle   string
    UserID        int64
    CourseID      int64
    Amount        int64
    Timeout       time.Duration
}

// Load читает переменные окружения и возвращает конфигурацию по умолчанию.
func Load() Config {
    driver := os.Getenv("DB_DRIVER")
    if driver == "" {
        driver = "sqlite"
    }

    dsn := os.Getenv("DB_DSN")
    if dsn == "" {
        dsn = "file:data.db?_foreign_keys=on&_busy_timeout=5000"
    }

    return Config{
        Driver:  driver,
        DSN:     dsn,
        Timeout: 5 * time.Second,
    }
}
