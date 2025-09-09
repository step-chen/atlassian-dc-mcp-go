# Atlassian Data Center MCP Server

This directory contains the main server implementation for the Atlassian Data Center MCP service.

## Overview

The server component exposes Atlassian Data Center products (Jira, Confluence, Bitbucket) through the Model Context Protocol (MCP) as a set of tools that can be consumed by MCP-compatible clients.

## Running the Server

To run the server directly:

```bash
go run main.go
```

Or from the project root:

```bash
go run cmd/server/main.go
```

## Building the Server

To build the server binary:

```bash
go build -o mcp-server main.go
```

Or from the project root:

```bash
go build -o mcp-server ./cmd/server
```

## Docker

The server can also be run using Docker. Please refer to the main [README.md](../../../README.md) for detailed Docker usage instructions.