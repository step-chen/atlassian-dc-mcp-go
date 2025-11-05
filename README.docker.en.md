# Start and Configure MCP Service via Docker

This document will guide you on how to quickly start and configure the Atlassian Data Center MCP service using Docker.

## Prerequisites

- Ensure Docker and Docker Compose are installed on your system.

## Steps

### 1. Create and Configure `config.yaml`

First, you need to create a local configuration file.

1. Copy a configuration file from the example file:
   ```bash
   cp config.yaml.example config.yaml
   ```

2. Edit the `config.yaml` file and fill in the connection information for your own Atlassian products (Jira, Confluence, Bitbucket), for example:
   - `url`: Your Atlassian product instance address
   - `token`: The corresponding API Token

Note: When using Docker, you can leave the token fields empty if you plan to use header-based authentication.

### 2. Modify `docker-compose.yml`

To allow the Docker container to read your local configuration, you need to modify the `docker-compose.yml` file to map the `config.yaml` file into the container.

Open the `docker-compose.yml` file and add the following mapping in the `volumes` section of the `mcp-server` service:

```yaml
services:
  mcp-server:
    # ... (other configurations remain unchanged)
    image: ghcr.io/step-chen/atlassian-dc-mcp-go:latest
    ports:
      - "8090:8090"
    volumes:
      - ./config.yaml:/app/config.yaml # <-- Add or modify this line
    # ... (other configurations remain unchanged)
```

To enable header-based authentication mode, you'll need to add a command parameter:

```yaml
services:
  mcp-server:
    # ... (other configurations remain unchanged)
    image: ghcr.io/step-chen/atlassian-dc-mcp-go:latest
    command: ["-auth-mode=header"]
    ports:
      - "8090:8090"
    volumes:
      - ./config.yaml:/app/config.yaml
    # ... (other configurations remain unchanged)
```

This links the `config.yaml` file you created in the project root directory to the `/app/config.yaml` path in the container. The service will load this configuration when starting.

### 3. Start the MCP Service

After completing the configuration, run the following command to pull the latest `latest` image from GitHub Container Registry and start the service:

```bash
docker compose up -d
```

This command starts the MCP service in the background. You can run `docker compose logs -f` to view real-time logs and ensure the service starts properly.

### 4. Configure the AI Assistant

Finally, in the configuration of your AI assistant or related client, add or modify the MCP Server address to point to the locally running Docker container.

```json
{
  "mcpServers": {
    "atlassian-dc-mcp-sse": {
      "url": "http://localhost:8090/sse"
    }
  }
}
```

`http://localhost:8090` is the mapped port configured in `docker-compose.yml`, and `/sse` is the SSE endpoint for the MCP service.

When using header-based authentication, your client will need to pass the appropriate headers:
- `Jira-Token`: API token for Jira
- `Confluence-Token`: API token for Confluence
- `Bitbucket-Token`: API token for Bitbucket

## Stop the Service

If you need to stop the service, you can run the following command:

```bash
docker compose down
```

## Example Documents

- For an example of config.yaml, please refer to [config.yaml.example](config.yaml.example)
- For an example of docker-compose.yml, please refer to [docker-compose.yml](docker-compose.yml)