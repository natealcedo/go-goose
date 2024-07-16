APP = goose_app

POSTGRES_CONFIG = "user=postgres password=postgres dbname=goose sslmode=disable"

PHONY: build run clean test migrate install make-migration

install:
	go mod tidy

build:
	go build -o bin/$(APP)

run: build
	./bin/$(APP)

clean:
	rm -f myapp

test:
	go test ./...

migrate:
	 goose -dir migrations postgres $(POSTGRES_CONFIG) up

migrate-reset:
	 goose -dir migrations postgres $(POSTGRES_CONFIG) reset

make-migration:
	 goose -dir migrations create $(ARGS) sql
