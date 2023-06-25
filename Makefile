.PHONY: test

DB_PORT ?= 3307

# goサーバーの操作
run:
	docker compose exec app go run cmd/main.go

test:
	docker compose exec app sh -c "DB_PORT=$(DB_PORT) go test ./..."

# コンテナの操作
app-container:
	docker compose exec app bash