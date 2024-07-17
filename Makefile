APP = goose_app

POSTGRES_CONFIG = "user=postgres password=postgres dbname=goose sslmode=disable"

PHONY: run clean test migrate install make-migration

install:
	go mod tidy

build:
	go build -o bin/$(APP)

run: build
	@export $$(cat .env | xargs) && ./bin/$(APP)

clean:
	rm -f myapp

test:
	go test ./...

migrate:
	 goose -dir migrations postgres $(POSTGRES_CONFIG) up

migrate-reset:
	 goose -dir migrations postgres $(POSTGRES_CONFIG) reset

make-migration:
	 goose -dir migrations create $(filter-out $@,$(MAKECMDGOALS)) sql

%:
	@: