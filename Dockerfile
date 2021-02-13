FROM golang:1.14-alpine as deps

ADD go.mod /app/go.mod
WORKDIR /app

RUN ["go", "mod", "download"]
RUN go get -u github.com/go-bindata/go-bindata/go-bindata
RUN apk add alpine-sdk




FROM deps as build

ARG VERSION

ADD . /app
WORKDIR /app

RUN go-bindata -o "internal/common/bindata.go" -pkg "common" resources/ resources/partials/ && \
    go build -race -ldflags "-X main.AppVersion=$VERSION" -o "/tmp/oea-go" . && \
    chmod a+x /tmp/oea-go




FROM alpine

RUN addgroup -g 9999 -S user && \
    adduser -u 9999 -G user -S -H user

COPY --from=build /tmp/oea-go /
ENTRYPOINT ["/oea-go"]

USER user
