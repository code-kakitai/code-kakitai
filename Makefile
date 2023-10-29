.PHONY: test

help: # コマンド確認
	@echo "\033[32mAvailable targets:\033[0m"
	@grep "^[a-zA-Z\-]*:" Makefile | grep -v "grep" | sed -e 's/^/make /' | sed -e 's/://'

# app-containerが起動しているかどうかのチェック、起動していればtrueを返す
check_app_container = $(shell docker ps --format '{{.Names}}' | grep -q code-kakitai-app && echo "true" || echo "false")

# goサーバーの操作
test: lint
ifeq ($(check_app_container),true)
	docker compose exec app sh -c "go test ./..."
else
	cd app && go test ./...
endif

run:
	docker compose exec app sh -c "go run ./cmd/main.go"

hot-reload:
	docker compose exec app air

gen:
	docker compose exec app sh -c "go generate ./..."

lint:
ifeq ($(check_app_container),true)
	docker compose exec app sh -c "go vet ./..."
else
	cd app && go vet ./...
endif

tidy:
	docker compose exec app sh -c "go mod tidy"

# コンテナの操作

init:up
	go work init ./app ./pkg
	make migrate-apply
	make create-test-data

up:
	docker compose up -d

down:
	docker compose down

destroy:
	docker compose down --rmi all --volumes --remove-orphans

restart:
	docker compose restart

logs:
	docker compose logs -f

app-container:
	docker compose exec app bash

gen-swagger:
	swag init -g app/cmd/main.go  --output app/docs/swagger

swagger-up:
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

# test dataの作成
create-test-data:
	docker compose cp ./ops/test_data/ db:/tmp/test_data
	docker compose exec db sh -c "mysql -u root code_kakitai < /tmp/test_data/create_test_users.sql"
	docker compose exec db sh -c "mysql -u root code_kakitai < /tmp/test_data/create_test_owners.sql"
	docker compose exec db sh -c "rm -rf /tmp/test_data"