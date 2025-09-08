include .envrc
MIGRATIONS_DIR ?= ./cmd/migrate/migrations

.PHONY: migrate-up migrate-down migrate-down-all migrate-version migrate-create

migrate-create:
	migrate create -seq -ext sql -dir $(MIGRATIONS_DIR) $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	migrate -path=$(MIGRATIONS_DIR) -database=$(DB_ADDR) up

migrate-down:
	migrate -path=$(MIGRATIONS_DIR) -database=$(DB_ADDR) down

migrate-down-all:
	migrate -path=$(MIGRATIONS_DIR) -database=$(DB_ADDR) down

migrate-version:
	migrate -path=$(MIGRATIONS_DIR) -database=$(DB_ADDR) version


