.PHONY: dev prod stop restart logs migrate

include .env
HOST_DIR ?= $(shell pwd)

COMPOSE = docker-compose -f docker-compose.yml

dev:
	$(COMPOSE) up --build

prod:
	ENV=production $(COMPOSE) up --build -d

stop:
	$(COMPOSE) down -v

restart:
	$(MAKE) stop
	$(MAKE) dev

logs:
	$(COMPOSE) logs -f --tail=100
