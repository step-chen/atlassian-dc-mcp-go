# ---- Builder Stage ----
# Use the official Golang image as the base image for building.
FROM golang:1.25-alpine AS builder

# Add image metadata following OCI standard
LABEL maintainer="https://github.com/step-chen"
LABEL org.opencontainers.image.description="Atlassian Data Center MCP (Model Context Protocol) - A Go-based Model Context Protocol service for managing Atlassian Data Center products"
LABEL org.opencontainers.image.url="https://github.com/step-chen/atlassian-dc-mcp-go"
LABEL org.opencontainers.image.source="https://github.com/step-chen/atlassian-dc-mcp-go"
LABEL org.opencontainers.image.vendor="https://github.com/step-chen"

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
    -ldflags="-s -w -extldflags '-static'" \
    -o /atlassian-dc-mcp-server \
    ./cmd/server

# ---- Final Stage ----
# Use a distroless static image for the final stage.
# It's a minimal image that contains only the bare necessities for a static binary,
# including ca-certificates. It also runs as a non-root user by default.
FROM gcr.io/distroless/static-debian12

# Set the working directory.
WORKDIR /app

# Copy the compiled binary from the builder stage.
# Note: We are not copying config.yaml. Configuration should be mounted at runtime.
COPY --from=builder /atlassian-dc-mcp-server .

# Expose the default port (only relevant for HTTP/SSE modes)
# This is for documentation; the actual port is set via config.
EXPOSE 8090

# Command to run the application
ENTRYPOINT ["./atlassian-dc-mcp-server"]