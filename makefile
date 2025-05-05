include .env

.PHONY: default run services-up services-down migrate-up migrate-down

default: run

run:
	@go run cmd/api/main.go

services-up:
	@docker compose up -d

services-down:
	@docker compose down

migrate-up:
	@DATABASE_HOST="${DATABASE_HOST}" \
	DATABASE_PORT="${DATABASE_PORT}" \
	DATABASE_NAME="${DATABASE_NAME}" \
	DATABASE_USER="${DATABASE_USER}" \
	DATABASE_PWD="${DATABASE_PWD}" \
	tern migrate --migrations internal/store/postgres/migrations \
	--config internal/store/postgres/migrations/tern.conf

migrate-down:
	DATABASE_HOST="${DATABASE_HOST}" \
	DATABASE_PORT="${DATABASE_PORT}" \
	DATABASE_NAME="${DATABASE_NAME}" \
	DATABASE_USER="${DATABASE_USER}" \
	DATABASE_PWD="${DATABASE_PWD}" \
	tern migrate --migrations internal/store/postgres/migrations \
	--config internal/store/postgres/migrations/tern.conf -d -1

