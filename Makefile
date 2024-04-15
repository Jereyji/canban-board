include .env
DATABASE = restapi_dev

.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

migrate_up:
	migrate -path ./migrations -database 'postgres://postgres:$(DB_PASSWORD)@localhost:5436/postgres?sslmode=disable' up

migrate_down:
	migrate -path ./migrations -database 'postgres://postgres:$(DB_PASSWORD)@localhost:5436/postgres?sslmode=disable' down

.DEFAULT_GOAL := build