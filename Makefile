.PHONY: test

DB_PORT ?= 3307

help: # コマンド確認
	@echo "\033[32mAvailable targets:\033[0m"
	@grep "^[a-zA-Z\-]*:" Makefile | grep -v "grep" | sed -e 's/^/make /' | sed -e 's/://'

# goサーバーの操作
run:
	docker compose exec app go run cmd/main.go

test:
	docker compose exec app sh -c "DB_PORT=$(DB_PORT) go test ./..."

hot-reload:
	docker compose exec app air

# コンテナの操作
up:
	docker compose up -d

down:
	docker compose down

restart:
	docker compose restart

logs:
	docker compose logs -f

app-container:
	docker compose exec app bash

gen-swagger:
	swag init -g app/cmd/main.go  --output app/docs/swagger
	docker-compose -f app/docs/swagger/docker-compose.yml up -d