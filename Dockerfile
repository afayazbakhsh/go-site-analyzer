# syntax=docker/dockerfile:1

# 1️⃣ Build stage
FROM golang:1.23 AS builder

WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy all project files
COPY . .

# Build the Go binary
RUN go build -o main ./cmd/http/main.go

# 2️⃣ Run stage
FROM debian:bookworm-slim

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .

# Install CA certificates (optional, only if you use HTTPS)
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

EXPOSE 8080
CMD ["./main"]
