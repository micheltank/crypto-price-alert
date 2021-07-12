# Stage 1
FROM golang:1.16.4-alpine AS builder

ENV GOPATH="$HOME/go"

WORKDIR $GOPATH/src/github.com/micheltank/crypto-price-alert/notification-service

COPY . $GOPATH/src/github.com/micheltank/crypto-price-alert/notification-service

RUN apk update && \
    apk add gcc librdkafka-dev openssl-libs-static zlib-static zstd-libs \
    libsasl librdkafka-static lz4-dev lz4-static zstd-static libc-dev musl-dev

RUN go mod download
RUN go build -tags musl -ldflags '-w -extldflags "-static"' -o notification-service ./cmd/main.go

RUN cd ./cmd

# Stage 2
FROM alpine:latest

ENV GOPATH="$HOME/go"

WORKDIR /app

RUN apk update && apk add ca-certificates && apk --no-cache add tzdata

COPY --from=builder $GOPATH/src/github.com/micheltank/crypto-price-alert/notification-service .

ENTRYPOINT ["./notification-service"]