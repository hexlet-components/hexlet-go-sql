run:
	DB_DRIVER?=sqlite3
	DB_DSN?=file:data.db?_foreign_keys=on&_busy_timeout=5000
	go run ./cmd/app --cmd=$(CMD) --email=$(EMAIL) --name=$(NAME) --course=$(COURSE) --user=$(USER) --amount=$(AMOUNT)

migrate:
	goose -dir ./migrations $(DB_DRIVER) "$(DB_DSN)" up

rollback:
	goose -dir ./migrations $(DB_DRIVER) "$(DB_DSN)" down

sqlc:
	sqlc generate

test-integration:
	go test -tags=integration ./...

setup:
	go mod download


test: test-integration
