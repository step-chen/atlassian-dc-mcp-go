# Use the official Golang image as the base image for building
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mcp .

# Use a minimal alpine image for the final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests and wget for health checks
RUN apk --no-cache add ca-certificates wget

# Create a non-root user
RUN adduser -D -s /bin/sh mcp

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/mcp .

# Copy the actual config file instead of the example
COPY --from=builder /app/config.yaml ./config.yaml

# Change ownership of files to the mcp user
RUN chown -R mcp:mcp /app

# Switch to the non-root user
USER mcp

# Expose the default port
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["./mcp"]