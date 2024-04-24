include .env
DATABASE = restapi_dev

.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

docker_run:
	docker run --name=db -e $(DOCKER_SETTING)

migrate_up:
	migrate -path ./migrations $(MIGRATE) up

migrate_down:
	migrate -path ./migrations $(MIGRATE) down

.DEFAULT_GOAL := build