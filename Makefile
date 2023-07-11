SHELL := /bin/bash
CONTAINER_NAME_QUEUE := rabbitmq
CONTAINER_NAME_PG := postgres
CONTAINER_NAME_ZOO := confluentinc/cp-zookeeper
CONTAINER_NAME_KAFKA := confluentinc/cp-kafka
PATHS = ./application/... ./cmd/... ./ports/... ./adapters/...

server-up:
	@docker compose up -d
	## currently running in docker
	# @cd cmd && go run main.go

server-down:
	@docker compose down
	@docker stop $(CONTAINER_NAME_QUEUE) $(CONTAINER_NAME_PG) $(CONTAINER_NAME_ZOO) $(CONTAINER_NAME_KAFKA)
	@docker rm $(CONTAINER_NAME_QUEUE) $(CONTAINER_NAME_PG) $(CONTAINER_NAME_ZOO) $(CONTAINER_NAME_KAFKA)
	@docker volume rm $(CONTAINER_NAME_PG)_data $(CONTAINER_NAME_KAFKA)_data $(CONTAINER_NAME_ZOO)_data
	@docker volume prune -f


run-load-tests:
	@echo "Running load tests"
	@cd resources && k6	run load.js

lint:
	@echo "Linting..."
	@golangci-lint run 

vet:
	@echo "Vetting..."
	@go vet $(PATHS)

