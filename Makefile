VERSION = $(shell git rev-parse --short HEAD)
DEV_ENV_FILE=dev-docker/.env
DEV_COMPOSE_FILE=dev-docker/docker-compose.yml
DEV_COMPOSE_PROJ=oea-go-dev
GENERATOR_IMAGE=ghcr.io/artsafin/coda-schema-generator:v1.0.2

.PHONY: all

all: build

assets:
	test -f resources/assets/bootstrap.min.css || curl -s https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css -o resources/assets/bootstrap.min.css

build: assets
	docker build -f ./Dockerfile --build-arg "VERSION=$(VERSION)" $(BUILD_ARGS) -t oea-go:local .

test:
	go test ./...

dev-env:
	@echo 'alias dcl="docker-compose --env-file $(DEV_ENV_FILE) -f $(DEV_COMPOSE_FILE) -p $(DEV_COMPOSE_PROJ)"'

run-dev: build
	docker compose --env-file $(DEV_ENV_FILE) -f $(DEV_COMPOSE_FILE) -p $(DEV_COMPOSE_PROJ) down && \
	docker compose --env-file $(DEV_ENV_FILE) -f $(DEV_COMPOSE_FILE) -p $(DEV_COMPOSE_PROJ) up -d --remove-orphans --force-recreate && \
	docker compose --env-file $(DEV_ENV_FILE) -f $(DEV_COMPOSE_FILE) -p $(DEV_COMPOSE_PROJ) logs -f --tail=50

generate-coda:
	@docker pull -q $(GENERATOR_IMAGE)
	@docker run --rm $(GENERATOR_IMAGE) $(CODA_TOKEN) Oz23vO7xol > internal/employee/codaschema/codaschema.go
	@docker run --rm $(GENERATOR_IMAGE) $(CODA_TOKEN) TAC1aAH5mf > internal/office/codaschema/codaschema.go
	@go fmt internal/employee/codaschema/codaschema.go
	@go fmt internal/office/codaschema/codaschema.go
