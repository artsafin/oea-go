GOBIN=/usr/local/go/bin/go
TARGET = oea-go
TARGET_HOST = artsaf.in
TARGET_PATH = $(TARGET)/
VERSION = $(shell git rev-parse --short HEAD)

.PHONY: all clean run

all: $(TARGET)

assets:
	test -f resources/bootstrap.min.css || curl -s https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css -o resources/bootstrap.min.css

$(TARGET): clean assets
	$(GOBIN) generate
	#CGO_ENABLED=0 GOOS=linux -i -installsuffix cgo
	$(GOBIN) build -o "$(TARGET)" -ldflags "-X main.AppVersion=$(VERSION)"

clean:
	rm -fv $(TARGET)

deploy: $(TARGET)
	ssh $(TARGET_HOST) mkdir -pv $(TARGET_PATH)
	rsync $(TARGET) config.yaml .env $(TARGET_HOST):$(TARGET_PATH)
