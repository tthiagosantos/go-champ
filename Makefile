MIGRATE_DIR=internal/infrastructure/database/migrations

.PHONY: migrate-up migrate-down migrate-create

include .env
export $(shell sed 's/=.*//' .env)

migrate-up:
	goose -dir $(MIGRATE_DIR) postgres "$(DATABASE_URL)" up

migrate-down:
	goose -dir $(MIGRATE_DIR) postgres "$(DATABASE_URL)" down

migrate-create:
	@read -p "Nome da migration: " name; \
	goose -dir $(MIGRATE_DIR) create "$$name" sql
