.PHONY: dev prod

# Default environment-configurable values
HOST_PORT ?= 8080
CONTAINER_PORT ?= 8080
HOST_DIR ?= $(shell pwd)

dev:
	docker-compose up --build

prod:
	docker-compose up --build
