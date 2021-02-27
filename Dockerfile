FROM golang:1.16-alpine as deps

ADD go.mod /app/go.mod
WORKDIR /app

RUN ["go", "mod", "download"]
RUN apk add alpine-sdk




FROM deps as build

ARG VERSION

ADD . /app
WORKDIR /app

RUN time go build -ldflags "-X main.AppVersion=$VERSION" -o "/tmp/oea-go" . && \
    chmod a+x /tmp/oea-go && \
    echo -n "BIN SIZE: " && du -k /tmp/oea-go




FROM alpine

RUN addgroup -g 9999 -S user && \
    adduser -u 9999 -G user -S -H user

COPY --from=build /tmp/oea-go /
ENTRYPOINT ["/oea-go"]

USER user
