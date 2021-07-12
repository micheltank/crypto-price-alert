# Stage 1
FROM golang:1.16.4-alpine AS builder

ENV GOPATH="$HOME/go"

WORKDIR $GOPATH/src/github.com/micheltank/crypto-price-alert/price-alert-service

COPY . $GOPATH/src/github.com/micheltank/crypto-price-alert/price-alert-service

RUN apk update && apk add curl git && apk add gcc libc-dev

RUN go mod download
RUN go build -o price-alert-service ./cmd/main.go

RUN cd ./cmd

# Stage 2
FROM alpine:latest

ENV GOPATH="$HOME/go"

WORKDIR /app

RUN apk update && apk add ca-certificates && apk --no-cache add tzdata

COPY --from=builder $GOPATH/src/github.com/micheltank/crypto-price-alert/price-alert-service .

ENTRYPOINT ["./price-alert-service"]