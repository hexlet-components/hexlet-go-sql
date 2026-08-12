# USER занят окружением: make подхватывает имя пользователя из шелла, и флаг
# --user, объявленный как int64, падает на parse error. Числовые флаги идут со
# своими значениями по умолчанию, иначе пустая строка тоже не разбирается.
USER_ID ?= 0
COURSE_ID ?= 0

run:
	go run ./cmd/app --cmd=$(CMD) --email=$(EMAIL) --name=$(NAME) --course=$(COURSE) --user=$(USER_ID) --course-id=$(COURSE_ID)

# sqlc и goose объявлены в go.mod директивой tool, поэтому зовутся через
# `go tool` и приезжают вместе с зависимостями. Отдельная установка бинарников
# не нужна, а версия одна и та же у всех и на CI.
migrate:
	go tool goose -dir ./migrations postgres "$(DB_DSN)" up

rollback:
	go tool goose -dir ./migrations postgres "$(DB_DSN)" down

sqlc:
	go tool sqlc generate

install:
	go mod download

build:
	go build ./...

lint:
	@files=$$(gofmt -l .); if [ -n "$$files" ]; then echo "gofmt is needed for:"; echo "$$files"; exit 1; fi
	go vet ./...

test-integration:
	go test -tags=integration ./...

setup:
	cp -n .env.example .env
	go mod download

compose:
	docker compose up --abort-on-container-exit

compose-down:
	docker compose down -v --remove-orphans

.PHONY: install build lint test test-integration
test: test-integration
