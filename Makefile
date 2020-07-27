GOBIN=$(shell go env GOROOT)/bin/go
TARGET = oea-go
TARGET_LOCAL_PATH = docker/$(TARGET)
VERSION = $(shell git rev-parse --short HEAD)
COMPOSE_FILE=dev-docker/docker-compose.yml

.PHONY: all

all: build

assets:
	test -f resources/bootstrap.min.css || curl -s https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css -o resources/bootstrap.min.css
	GO111MODULE=off $(GOBIN) get github.com/go-bindata/go-bindata/go-bindata

build: assets
	$(shell go env GOPATH)/bin/go-bindata -o "common/bindata.go" -pkg "common" resources/ resources/partials/

	CGO_ENABLED=0 GOOS=linux \
	$(GOBIN) build -i -installsuffix cgo -o "$(TARGET_LOCAL_PATH)" -ldflags "-X main.AppVersion=$(VERSION)"

clean:
	rm -fv $(TARGET_LOCAL_PATH)

test:
	go test ./...

dev-env:
	@echo 'export COMPOSE_PROJECT_NAME=oea-go-dev; export COMPOSE_FILE=$(COMPOSE_FILE)'

run-dev: build
	docker-compose -p oea-go-dev -f $(COMPOSE_FILE) up -d --build --remove-orphans --force-recreate
	docker-compose -p oea-go-dev -f $(COMPOSE_FILE) logs -f --tail=10 app
