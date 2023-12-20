.PHONY: default docs run clean build migrate_up migrate_down seed

APP_NAME=hackton-minerva
DATABASE_DRIVER=postgres
DATABASE_URL=$(shell grep -m 1 database config.toml | tr -s ' ' | tr -d '"' | tr -d "'" | cut -d' ' -f3)
MIGRATIONS_PATH=config/database/${DATABASE_DRIVER}/migration

default: run

docs:
	@swag init -d "./" -g "cmd/main.go"

run: docs
	@go run cmd/main.go

clean:
	@rm -f bin/${APP_NAME}

build: clean docs
	@go build -C cmd -o ../bin/${APP_NAME}

migrate_up:
	@docker-compose up -d
	@migrate -path ${MIGRATIONS_PATH} -database "${DATABASE_URL}" -verbose up

migrate_down:
	@docker-compose up -d
	@migrate -path ${MIGRATIONS_PATH} -database "${DATABASE_URL}" -verbose down

seed:
	@docker exec -it db_minerva \
	psql -q -U postgres \
	-d minerva \
	-f "/mnt/seed/seed.sql"