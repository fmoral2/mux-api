SHELL := /bin/bash
CONTAINER_NAME_QUEUE := rabbitmq
CONTAINER_NAME_PG := employees
PATHS = ./application/... ./cmd/... ./ports/... ./adapters/...

server-up:
	@docker compose up -d
	@cd cmd && go run main.go

server-down:
	@echo Stopping ${CONTAINER_NAME_QUEUE} && echo "OK"
	@docker stop ${CONTAINER_NAME_QUEUE}
	@echo Removing ${CONTAINER_NAME_PG} && echo "OK"
	@docker rm ${CONTAINER_NAME_PG}

run-load-tests:
	@echo "Running load tests"
	@cd resources && k6	run load.js

lint:
	@echo "Linting..."
	@golangci-lint run 

vet:
	@echo "Vetting..."
	@go vet $(PATHS)

