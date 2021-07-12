# Stage 1
FROM golang:1.16.4-alpine AS builder

ENV GOPATH="$HOME/go"

WORKDIR $GOPATH/src/github.com/micheltank/crypto-price-alert/price-alert-job

COPY ./price-alert-service $GOPATH/src/github.com/micheltank/crypto-price-alert/price-alert-service
COPY ./price-alert-job $GOPATH/src/github.com/micheltank/crypto-price-alert/price-alert-job

RUN apk update && \
    apk add gcc librdkafka-dev openssl-libs-static zlib-static zstd-libs \
    libsasl librdkafka-static lz4-dev lz4-static zstd-static libc-dev musl-dev

RUN go mod download
RUN go build -tags musl -ldflags '-w -extldflags "-static"' -o price-alert-job ./cmd/main.go


# Stage 2
FROM alpine:latest

ENV GOPATH="$HOME/go"

WORKDIR /app

RUN apk update && apk add ca-certificates && apk --no-cache add tzdata

COPY --from=builder $GOPATH/src/github.com/micheltank/crypto-price-alert/price-alert-job .

ENTRYPOINT ["./price-alert-job"]