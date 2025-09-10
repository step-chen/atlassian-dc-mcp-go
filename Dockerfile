# ---- Builder Stage ----
# Use the official Golang image as the base image for building.
FROM golang:1.25-alpine AS builder

# Install ca-certificates which will be copied to the final image.
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container.
WORKDIR /app

# Copy go mod and sum files to leverage Docker cache.
COPY go.mod go.sum ./

# Download dependencies.
RUN go mod download

# Copy the rest of the source code.
# This is done after downloading dependencies to ensure that the dependency
# layer is cached if only the source code changes.
COPY . .

# Build the application as a static binary.
# -s: strip debug symbols
# -w: strip DWARF table
# This significantly reduces the binary size.
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w" \
    -o /atlassian-dc-mcp-server \
    ./cmd/server

# ---- Final Stage ----
# Use a minimal 'scratch' image for the final stage.
# 'scratch' is an empty image, perfect for single static binaries.
FROM scratch

# Copy the ca-certificates from the builder stage to allow for HTTPS requests.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Set a non-root user (UID 1001 is a common choice for custom users).
USER 1001

# Set the working directory.
WORKDIR /app

# Copy the compiled binary from the builder stage and set ownership.
# Note: We are not copying config.yaml. Configuration should be mounted at runtime.
COPY --from=builder --chown=1001:1001 /atlassian-dc-mcp-server .

# Expose the default port (only relevant for HTTP/SSE modes)
# This is for documentation; the actual port is set via config.
EXPOSE 8090

# Command to run the application
ENTRYPOINT ["./atlassian-dc-mcp-server"]