VERSION = $(shell git rev-parse --short HEAD)
DEV_COMPOSE_FILE=dev-docker/docker-compose.yml

.PHONY: all

all: build

assets:
	test -f resources/bootstrap.min.css || curl -s https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css -o resources/bootstrap.min.css

build: assets
	docker build -f ./Dockerfile --build-arg "VERSION=$(VERSION)" -t oea-go:latest .

test:
	go test ./...

dev-env:
	@echo 'export COMPOSE_PROJECT_NAME=oea-go-dev; export COMPOSE_FILE=$(DEV_COMPOSE_FILE)'

run-dev: build
	cd dev-docker && \
	docker-compose -p oea-go-dev down -v && \
	docker-compose -p oea-go-dev up -d --remove-orphans && \
	docker-compose -p oea-go-dev logs -f --tail=50
