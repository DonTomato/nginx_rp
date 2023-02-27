# syntax=docker/dockerfile:1

FROM golang:1.20.1-alpine3.17 as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /service .

FROM alpine:3
WORKDIR /app
COPY --from=builder service .

EXPOSE 3000
ENTRYPOINT [ "/app/service" ]
