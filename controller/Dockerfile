# syntax=docker/dockerfile:1

# --- Builder stage ---
FROM golang:1.24.4-alpine AS builder

ARG TARGETOS=linux
ARG TARGETARCH

WORKDIR /app

# Install git for Go modules in private repos (if needed)
RUN apk add --no-cache git

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the controller binary
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -a -o uptime-kuma-controller main.go

# --- Final image ---
FROM alpine:latest

ARG TARGETARCH

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/uptime-kuma-controller .

# Use a non-root user for security (optional, can be omitted if not needed)
# RUN adduser -D appuser
# USER appuser

# Entrypoint
ENTRYPOINT ["./uptime-kuma-controller"]