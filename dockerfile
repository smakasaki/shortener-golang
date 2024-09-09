# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder

WORKDIR /build

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o shortener ./cmd

# Stage 2: Copy the Go application to a smaller image
FROM scratch

WORKDIR /app

COPY --from=builder /build/shortener .
COPY --from=builder /build/migrations ./migrations

EXPOSE 8080

ENTRYPOINT ["./shortener"]

