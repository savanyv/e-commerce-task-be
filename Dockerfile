# Build stage
FROM golang:1.24-alpine as builder

WORKDIR /app

# Copy and download dependency using go mod
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/e-commerce-task-be ./cmd/main.go

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/e-commerce-task-be .

# Copy environment files (if any)
COPY --from=builder /app/.env.sample .env

# Create necessary directories
RUN mkdir -p /app/logs

# Make the binary executable
RUN chmod +x /app/e-commerce-task-be

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./e-commerce-task-be"]
