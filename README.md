# Atlassian Data Center MCP (Model Context Protocol)

Author: Stephen Chen

This project provides a Go-based Model Context Protocol (MCP) service for managing and interacting with Atlassian Data Center products including Jira, Confluence, and Bitbucket. It allows you to manage these products programmatically through a unified interface with configurable authentication and permissions.

## Features

- **Multi-Product Support**: Unified interface for Jira, Confluence, and Bitbucket
- **Model Context Protocol**: Exposes tools through the Model Context Protocol for all operations
- **Configuration Management**: Flexible configuration through files, environment variables, and hot reloading
- **Authentication**: Direct API token authentication for each service
- **Permissions**: Fine-grained read/write permissions for each service
- **Health Monitoring**: Built-in health checks for all services
- **Docker Support**: Ready for containerized deployment

## Running the Application

You can run the application in multiple ways: directly using Go, by building binaries first, or using Docker.

For detailed instructions on running the application:
- [Running with Docker](README.docker.en.md) - Instructions for running the service using Docker and Docker Compose
- [Running locally](README.local.en.md) - Instructions for running the service directly on your system

Below are basic commands for development purposes:

### Direct Execution

```bash
# Run the server
go run cmd/server/main.go

# Run with custom config file path
go run cmd/server/main.go -c /path/to/your/config.yaml
```

### Building and Running Binaries

The project uses a Makefile to simplify the build process. All binaries are placed in the `dist` directory.

```bash
# Show all available commands
make help

# Build the server binary for your current OS
make build

# Build a statically linked release binary
make release
```

After building the project, you can run the binary directly from the `dist` directory:

```bash
# Run the server
./dist/atlassian-dc-mcp-server

# Run with custom config file path
./dist/atlassian-dc-mcp-server -c /path/to/your/config.yaml
```

## Configuration

The application requires a configuration file to operate. By default, it looks for `config.yaml` in the current directory, but you can specify a different path using the `-c` or `--config` flag.

A sample configuration file is provided as `config.yaml.example`. Copy this file to `config.yaml` and edit it with your settings:

```bash
cp config.yaml.example config.yaml
```

The configuration file is self-documented with examples for all available settings. Please refer to the [config.yaml.example](config.yaml.example) file for detailed configuration options.

### Authentication Modes

The service supports two authentication modes:

1. **Config Mode (default)**: API tokens are read from the configuration file
2. **Header Mode**: API tokens are passed via HTTP headers

To enable header mode, start the server with the `-auth-mode=header` flag:

```bash
./dist/atlassian-dc-mcp-server -auth-mode=header
```

In header mode, the service expects the following HTTP headers:
- `Jira-Token`: API token for Jira
- `Confluence-Token`: API token for Confluence
- `Bitbucket-Token`: API token for Bitbucket

This mode is particularly useful when deploying the service in environments where you want to avoid storing sensitive tokens in configuration files, such as when using the service behind a reverse proxy that handles authentication.

## Tools Documentation

### Jira Tools

Tools for interacting with Jira:
- Get current user information
- Get issues
- Create issues
- And many more

### Confluence Tools

Tools for interacting with Confluence:
- Get current user information
- Get content
- Search content
- And more

### Bitbucket Tools

Tools for interacting with Bitbucket:
- Get current user information
- Get repositories
- Get commits
- And more

## Lingma Rules

This project includes predefined Lingma rules that demonstrate how to use the Atlassian Data Center MCP service for automated code review tasks. For detailed information on how to use these rules, please refer to the [Lingma Rules documentation](docs/lingma-rules.md).

The [code_review.md](.lingma/rules/code_review.md) rule file provides an example configuration showing how to leverage the Atlassian Data Center MCP service for automated Bitbucket PR code reviews. These rules help you:

- Automate code review processes using the MCP service
- Standardize code quality and security checks
- Improve interaction efficiency with AI assistants during code reviews

## Development

### Prerequisites

- Go 1.20 or higher
- Docker (for containerization)

### Building

It is recommended to use the Makefile to build the project, which ensures all build artifacts are placed in the unified `dist` directory:

```bash
# Build server binary to dist directory
make build
```

To build binaries for specific operating systems:
```bash
# Build for Linux
make build-linux

# Build for Windows
make build-windows

# Build for macOS
make build-macos
```

Benefits of using make commands:
- All build artifacts are placed in the `dist` directory
- Automatically handles cross-platform builds
- Ensures consistent build parameters

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## References

- [Confluence REST API](https://developer.atlassian.com/server/confluence/rest/v1010/intro/#about)
- [Jira REST API](https://developer.atlassian.com/server/jira/platform/rest/v11000/intro/#gettingstarted)
- [Bitbucket REST API](https://developer.atlassian.com/server/bitbucket/rest/v1000/intro/#about)
- [Model Context Protocol Go SDK](https://github.com/modelcontextprotocol/go-sdk)
