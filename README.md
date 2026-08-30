# Hexlet Go SQL

[![CI](https://github.com/hexlet-components/hexlet-go-sql/actions/workflows/ci.yml/badge.svg)](https://github.com/hexlet-components/hexlet-go-sql/actions/workflows/ci.yml)

## Зачем это нужно

Эталонное решение, которое получается, если выполнить все задания из
самостоятельной работы курса по SQL на Go. Показывает, как в одном приложении
сочетаются PostgreSQL, [sqlc](https://sqlc.dev/), миграции
[goose](https://github.com/pressly/goose), CLI и слой репозитория.

Что здесь есть:

- команды CLI для пользователей, курсов и записи на курс;
- подготовленные выражения и транзакции для пакетных сценариев;
- слой репозитория, написанный руками поверх сгенерированных sqlc запросов;
- миграции goose на пользователей, курсы, уроки и связь многие-ко-многим;
- конфигурация через окружение (`DB_DRIVER`, `DB_DSN` и остальные).

## Requirements

- Go 1.26
- Docker и Docker Compose, чтобы поднять PostgreSQL локально

`sqlc` и `goose` ставить отдельно не нужно: они объявлены в `go.mod` директивой
`tool` и зовутся через `go tool`. Версия инструмента лежит там же, где версии
библиотек, поэтому у всех и на CI она одна.

## Getting Started

```bash
docker compose up -d db
make setup                            # .env из .env.example и загрузка зависимостей
export DB_DRIVER=pgx
export DB_DSN="postgres://app:secret@localhost:6543/app?sslmode=disable&application_name=hexlet-go-sql"
make migrate
make run CMD=course-create COURSE="Graph Theory"
```

## Commands

```bash
make migrate                 # goose up
make rollback                # goose down
make sqlc                    # перегенерировать код sqlc
make lint                    # проверка gofmt и go vet
make build                   # сборка всех пакетов
make run CMD=user-create EMAIL=test@example.com NAME="Demo User"
make run CMD=course-create COURSE="Graph Theory"
make run CMD=join-course USER_ID=1 COURSE_ID=1
```

Идентификаторы передаются как `USER_ID`, а не `USER`: имя `USER` занято
окружением шелла, и make подставил бы в числовой флаг имя пользователя.

## Тесты

Тестов в репозитории пока нет. Цель `make test-integration` заведена под будущий
набор с тегом сборки `integration`, чтобы тесты ходили в отдельную базу, а не в
рабочую. CI проверяет сборку и статические проверки: на пустом наборе `go test`
вернул бы успех, ничего не проверив.

## Project Layout

- `cmd/app`: точка входа CLI.
- `internal/config`: загрузка конфигурации из окружения.
- `internal/repository`: доступ к данным и помощники для транзакций.
- `internal/db`: код, сгенерированный sqlc.
- `internal/migrate`: запуск goose из приложения.
- `migrations`: SQL-миграции.
- `query` и `sqlc.yaml`: запросы и конфигурация sqlc.
- `docker-compose.yml`: PostgreSQL для локальной разработки.

---

[![Hexlet Ltd. logo](https://raw.githubusercontent.com/Hexlet/assets/master/images/hexlet_logo128.png)](https://hexlet.io?utm_source=github&utm_medium=link&utm_campaign=hexlet-go-sql)

This repository is created and maintained by the team and the community of Hexlet, an educational project. [Read more about Hexlet](https://hexlet.io?utm_source=github&utm_medium=link&utm_campaign=hexlet-go-sql).

See most active contributors on [hexlet-friends](https://friends.hexlet.io/).
