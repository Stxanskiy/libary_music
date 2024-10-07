FROM golang:1.23.1-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./


RUN go mod download

COPY . .

RUN go build -o main ./cmd



CMD ["/main"]
RUN go run cmd/main.go
