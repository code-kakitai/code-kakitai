.PHONY: test

DB_PORT ?= 3307

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