include .env

export

MIGRATIONS_PATH=sql/migrations

BINARY=goapi

.PHONY: help run build clean fmt vet test sqlc migrate-up migrate-down migrate-create docker-up docker-down docker-build

help:
	@echo "Available commands:"
	@echo " make run"
	@echo " make build"
	@echo " make test"
	@echo " make fmt"
	@echo " make vet"
	@echo " make sqlc"
	@echo " make migrate-up"
	@echo " make migrate-down"
	@echo " make migrate-create name=create_users"
	@echo " make docker-up"
	@echo " make docker-down"

setup:
	@echo "Setting up project..."
	@go mod download
	@go mod tidy
	@sqlc generate
	@echo "Setup completed!"

run:
	@air

up:
	@docker compose up -d
	@make migrate-up

build:
	@mkdir -p bin
	@go build -o bin/$(BINARY) ./cmd/api

clean:
	@rm -rf bin
	@rm -rf tmp

fmt:
	@gofmt -w .

vet:
	@go vet ./...

test:
	@go test ./... -v

sqlc:
	@sqlc generate

migrate-up:
	@migrate \
		-path $(MIGRATIONS_PATH) \
		-database "$(DATABASE_URL)" \
		up

migrate-down:
	@migrate \
		-path $(MIGRATIONS_PATH) \
		-database "$(DATABASE_URL)" \
		down

migrate-force:
	@migrate \
		-path $(MIGRATIONS_PATH) \
		-database "$(DATABASE_URL)" \
		force $(version)

migrate-version:
	@migrate \
		-path $(MIGRATIONS_PATH) \
		-database "$(DATABASE_URL)" \
		version

migrate-create:
	@migrate create \
		-ext sql \
		-dir $(MIGRATIONS_PATH) \
		-seq \
		$(name)

docker-build:
	docker compose build

docker-up:
	docker compose up -d

docker-down:
	docker compose down

docker-logs:
	docker compose logs -f
