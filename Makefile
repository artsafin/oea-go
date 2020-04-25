GOBIN=$(shell go env GOROOT)/bin/go
TARGET = oea-go
TARGET_LOCAL_PATH = docker/$(TARGET)
VERSION = $(shell git rev-parse --short HEAD)

.PHONY: all

all: build

assets:
	test -f resources/bootstrap.min.css || curl -s https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css -o resources/bootstrap.min.css

build: assets
	$(shell go env GOPATH)/bin/go-bindata -o "common/bindata.go" -pkg "common" resources/ resources/partials/

	CGO_ENABLED=0 GOOS=linux \
	$(GOBIN) build -i -installsuffix cgo -o "$(TARGET_LOCAL_PATH)" -ldflags "-X main.AppVersion=$(VERSION)"

clean:
	rm -fv $(TARGET_LOCAL_PATH)
