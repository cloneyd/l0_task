# syntax=docker/dockerfile:1

## Build
FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./test ./test

WORKDIR /app/cmd/server

RUN go build -o /nats-service

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /nats-service /nats-service

USER nonroot:nonroot

ENTRYPOINT ["/nats-service"]