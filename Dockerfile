# Base image
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first (for dependency caching)
COPY go.mod ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build main app binary
RUN go build -o bin/app ./cmd/app

# Final stage
FROM alpine:latest

# Create app directory
WORKDIR /app

# Copy the compiled binaries from builder
COPY --from=builder /app/bin/app /app/app

# Copy file configs and migration files
COPY --from=builder /app/file /app/file

# Expose application port
EXPOSE 3000

# Command to run the application
CMD ["./app"]
