TARGET = oea-go
VERSION = $(shell git rev-parse --short HEAD)

.PHONY: all clean run

all: $(TARGET)

assets:
	test -f resources/bootstrap.min.css || curl -s https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css -o resources/bootstrap.min.css

$(TARGET): clean assets
	go generate
	go build -o "$(TARGET)" -ldflags "-X main.AppVersion=$(VERSION)"

clean:
	rm -fv $(TARGET)
