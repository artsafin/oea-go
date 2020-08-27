TARGET = oea-go
TARGET_LOCAL_PATH = docker/$(TARGET)
VERSION = $(shell git rev-parse --short HEAD)
COMPOSE_FILE=dev-docker/docker-compose.yml

.PHONY: all

all: build-image build

assets:
	test -f resources/bootstrap.min.css || curl -s https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css -o resources/bootstrap.min.css

build-image:
	echo "package common; func MustAsset(a string) []byte { return []byte{} }" > common/bindata.go
	docker build -f docker/Dockerfile-build -t oea-go-builder .

build: assets
	docker run --rm -v $(PWD):/app -w /app oea-go-builder \
	sh -x -c '\
		pwd && \
		go-bindata -o "common/bindata.go" -pkg "common" resources/ resources/partials/ && \
		go build -v -o "$(TARGET_LOCAL_PATH)" -ldflags "-X main.AppVersion=$(VERSION)" \
	'

clean:
	rm -fv $(TARGET_LOCAL_PATH)

test:
	go test ./...

dev-env:
	@echo 'export COMPOSE_PROJECT_NAME=oea-go-dev; export COMPOSE_FILE=$(COMPOSE_FILE)'

run-dev: build
	docker-compose -p oea-go-dev -f $(COMPOSE_FILE) down
	docker-compose -p oea-go-dev -f $(COMPOSE_FILE) up -d --force-recreate --remove-orphans --build
	docker-compose -p oea-go-dev -f $(COMPOSE_FILE) logs -f --tail=50
