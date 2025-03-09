# Use the correct Go version
FROM golang:1.23 AS builder

WORKDIR /app

# Install Air separately
RUN go install github.com/air-verse/air@latest

# Copy go.mod and go.sum first (for caching dependencies)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy all project files
COPY . .  

# Build the binary (optional for Air)
RUN go mod tidy && go build -o main ./cmd/server

# Expose the application port
EXPOSE 8080

# Set environment variables
ENV REDIS_HOST=my_redis
ENV REDIS_PORT=6379

# Ensure Air is in PATH and start the app
CMD ["/go/bin/air"]
