FROM golang:1.17-alpine as deps

RUN mkdir -pv /app
COPY go.mod go.sum /app
WORKDIR /app

RUN go mod download


FROM deps as build

ARG VERSION

ADD . /app
WORKDIR /app

RUN echo "Building version $VERSION" && \
    time go build -ldflags "-X main.AppVersion=$VERSION -v" -o "/tmp/oea-go" ./cmd/server && \
    chmod a+x /tmp/oea-go && \
    echo -n "BIN SIZE: " && du -k /tmp/oea-go




FROM alpine

ARG VERSION

RUN echo "Building version $VERSION" && \
    addgroup -g 9999 -S user && \
    adduser -u 9999 -G user -S -H user

LABEL oea_version="$VERSION"

COPY --from=build /tmp/oea-go /
ENTRYPOINT ["/oea-go"]

USER user
