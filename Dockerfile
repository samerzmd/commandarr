# Build Stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy Go modules files
COPY go.mod go.sum ./
RUN go mod download

# Copy entire project
COPY . .

# Build the Go application
RUN go build -o commandarr ./cmd/bot

# Run Stage
FROM alpine:latest

WORKDIR /app

# Install CA certificates (needed for HTTPS requests)
RUN apk --no-cache add ca-certificates

# Copy binary from builder stage
COPY --from=builder /app/commandarr .

# Expose any necessary ports (optional, only if you plan a health endpoint)
# EXPOSE 8080

# Run the binary
ENTRYPOINT ["./commandarr"]
