DB_DRIVER:=pgx
DB_DIALECT:=postgres
DB_DSN?=postgres://app:secret@localhost:6543/app?sslmode=disable&application_name=hexlet-go-sql
DB_DRIVER=


run:
	go run ./cmd/app --cmd=$(CMD) --email=$(EMAIL) --name=$(NAME) --course=$(COURSE) --user=$(USER) --course-id=$(COURSE_ID)

migrate:
	goose -dir ./migrations $(DB_DIALECT) "$(DB_DSN)" up

rollback:
	goose -dir ./migrations $(DB_DIALECT) "$(DB_DSN)" down

sqlc:
	sqlc generate

test-integration:
	go test -tags=integration ./...

setup:
	go mod download
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest

compose:
	docker compose up --abort-on-container-exit

compose-down:
	docker compose down -v --remove-orphans


test: test-integration
