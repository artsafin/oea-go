VERSION = $(shell git rev-parse --short HEAD)
DEV_ENV_FILE=dev-docker/.env
DEV_COMPOSE_FILE=dev-docker/docker-compose.yml
DEV_COMPOSE_PROJ=oea-go-dev

.PHONY: all

all: build

assets:
	test -f resources/bootstrap.min.css || curl -s https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css -o resources/bootstrap.min.css

build: assets
	docker build -f ./Dockerfile --build-arg "VERSION=$(VERSION)" -t oea-go:latest .

test:
	go test ./...

dev-env:
	@echo 'export COMPOSE_PROJECT_NAME=$(DEV_COMPOSE_PROJ); export COMPOSE_FILE=$(DEV_COMPOSE_FILE)'

run-dev: build
	docker-compose --env-file $(DEV_ENV_FILE) -f $(DEV_COMPOSE_FILE) -p $(DEV_COMPOSE_PROJ) down && \
	docker-compose --env-file $(DEV_ENV_FILE) -f $(DEV_COMPOSE_FILE) -p $(DEV_COMPOSE_PROJ) up -d --remove-orphans && \
	docker-compose --env-file $(DEV_ENV_FILE) -f $(DEV_COMPOSE_FILE) -p $(DEV_COMPOSE_PROJ) logs -f --tail=50
