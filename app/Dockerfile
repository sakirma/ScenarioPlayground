# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main /app/cmd/service

# Final stage
FROM alpine:3.20

WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/main .

# Create a non-root user
RUN adduser -D appuser
USER appuser

# Expose the port the app runs on
EXPOSE 8080

# Run the binary
CMD ["./main"]
