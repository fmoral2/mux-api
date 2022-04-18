include ./config.mk

SHELL := /bin/bash
CONTAINER_NAME := some-rabbit


mux-server-up && rabbit:
	# @docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management 
	@export POSTGRES_URL="${POSTGRES_URL}"
	@cd cmd && go run main.go

rabbit-server-removal:
	@echo Stopping ${CONTAINER_NAME} && echo "OK"
	@docker stop ${CONTAINER_NAME}
	@echo Removing ${CONTAINER_NAME} && echo "OK"
	@docker rm ${CONTAINER_NAME}

run-tests:
