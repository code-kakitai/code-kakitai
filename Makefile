.PHONY: test

help: # コマンド確認
	@echo "\033[32mAvailable targets:\033[0m"
	@grep "^[a-zA-Z\-]*:" Makefile | grep -v "grep" | sed -e 's/^/make /' | sed -e 's/://'

# goサーバーの操作
test: lint
	docker compose exec app sh -c "DB_PORT=$(DB_PORT) go test ./..."

run:
	docker compose exec app sh -c "go run ./cmd/main.go"

hot-reload:
	docker compose exec app air

gen:
	docker compose exec app sh -c "go generate ./..."

lint:
	docker compose exec app sh -c "go vet ./..."

tidy:
	docker compose exec app sh -c "go mod tidy"

# コンテナの操作
up:
	docker compose up -d

up-test-db:
	docker compose up -d test_db

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

# マイグレーション
build-cli: # cliのビルド
	cd app && go build -o ./cli/main ./cli/main.go
	
migrate-dry-run: up build-cli # migration dry-run
	shema_path=$$(find . -name "schema.sql"); \
	./app/cli/main migration $$shema_path
	cd app && rm ./cli/main

migrate-apply: up build-cli # migration apply
	shema_path=$$(find . -name "schema.sql"); \
	DB_HOST=$(DB_HOST) ./app/cli/main migration $$shema_path apply
	cd app && rm ./cli/main

migrate-local-dry-run:
	@make migrate-dry-run DB_HOST=127.0.0.1

migrate-local-apply:
	@make migrate-apply DB_HOST=127.0.0.1

# sqlc
sqlc-gen:
	docker compose exec app sh -c "sqlc generate"
