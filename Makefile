include .env

DB_URL=postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_NAME}?sslmode=disable
MIGRATION_DIR=internal/database/migrations

create_migration:
	@if [ -z "$(name)" ]; then echo "Error: 'name' parameter is required. Usage: make create_migration name=migration_name"; exit 1; fi
	migrate create -ext=sql -dir=$(MIGRATION_DIR) -seq $(name)

migrate_up:
	migrate -path=$(MIGRATION_DIR) -database "$(DB_URL)" -verbose up

migrate_down:
	migrate -path=$(MIGRATION_DIR) -database "$(DB_URL)" -verbose down

migrate_step:
	@if [ -z "$(step)" ]; then echo "Error: 'step' parameter is required. Usage: make migrate_step step=<N>"; exit 1; fi
	migrate -path=$(MIGRATION_DIR) -database "$(DB_URL)" -verbose up $(step)

.PHONY: create_migration migrate_up migrate_down migrate_step