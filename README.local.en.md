# Running MCP Service Locally

This document will guide you on how to run the Atlassian Data Center MCP service directly in a local environment without using Docker.

## Steps

### 1. Download Pre-compiled Binary

Visit the [GitHub Releases page](https://github.com/step-chen/atlassian-dc-mcp-go/releases) to download the latest binary for your operating system.

### 2. Create and Configure `config.yaml`

First, you need to create a local configuration file.

1. Copy a configuration file from the example:
   ```bash
   cp config.yaml.example config.yaml
   ```

2. Edit the `config.yaml` file and fill in your own Atlassian product (Jira, Confluence, Bitbucket) connection information, for example:
   - `url`: Your Atlassian product instance address
   - `token`: The corresponding API Token

Note: You can leave the token fields empty if you plan to use header-based authentication.

### 3. Run the Service

After downloading and extracting the binary, you can run the service directly:
```
./atlassian-dc-mcp-server -c /path/to/your/config.yaml
```

To enable header-based authentication mode, use the `-auth-mode=header` flag:
```
./atlassian-dc-mcp-server -c /path/to/your/config.yaml -auth-mode=header
```

### 4. Verify Service Running

After the service starts successfully, you will see log output similar to the following:

```
INFO[0000] Starting Atlassian DC MCP Server on port 8090
```

By default, the service will listen on `http://localhost:8090`.

### 5. Configure AI Assistant

In your AI assistant or related client configuration, add or modify the MCP Server address to point to the locally running service.

Depending on the transport mode enabled in your configuration file, use the corresponding endpoint:

```json
{
  "mcpServers": {
    "atlassian-dc-mcp-local": {
      "url": "http://localhost:8090/sse"
    }
  }
}
```

Or, if you have enabled HTTP mode:

```json
{
  "mcpServers": {
    "atlassian-dc-mcp-local": {
      "url": "http://localhost:8090/mcp"
    }
  }
}
```

When using header-based authentication, your client will need to pass the appropriate headers:
- `Jira-Token`: API token for Jira
- `Confluence-Token`: API token for Confluence
- `Bitbucket-Token`: API token for Bitbucket

Here's a complete example for header-based authentication:

```json
{
  "mcpServers": {
    "atlassian-dc-mcp-local": {
      "url": "http://localhost:8090/sse",
      "headers": {
        "Jira-Token": "your-jira-api-token",
        "Confluence-Token": "your-confluence-api-token",
        "Bitbucket-Token": "your-bitbucket-api-token"
      }
    }
  }
}
```

Alternatively, you can use environment variables for the tokens:

```json
{
  "mcpServers": {
    "atlassian-dc-mcp-local": {
      "url": "http://localhost:8090/sse",
      "headers": {
        "Jira-Token": "${JIRA_TOKEN}",
        "Confluence-Token": "${CONFLUENCE_TOKEN}",
        "Bitbucket-Token": "${BITBUCKET_TOKEN}"
      }
    }
  }
}
```

### 6. Stop the Service

To stop the service, simply press `Ctrl+C` in the terminal to gracefully shut down the service.

## Troubleshooting

### Service Startup Failure

1. Check if the configuration file is properly configured
2. Ensure the port is not occupied by other programs
3. Check if the Atlassian product URL and API Token are correct

### Unable to Connect to Atlassian Products

1. Ensure network connectivity is normal
2. Verify the Atlassian product URL and API Token
3. Check firewall settings

## Configure User-level Service Auto-start

### Ubuntu User-level Service Auto-start

In Ubuntu systems, you can use systemd user services to configure the service to start automatically, so the service will run as the current user rather than as a system-level service:

1. Create systemd user service directory (if it doesn't exist):
   ```bash
   mkdir -p ~/.config/systemd/user
   ```

2. Create systemd user service file:
   ```bash
   nano ~/.config/systemd/user/atlassian-dc-mcp.service
   ```

3. Add the following content to the file (please adjust according to actual paths):
   ```ini
   [Unit]
   Description=Atlassian Data Center MCP Service
   After=network.target

   [Service]
   Type=simple
   ExecStart=/path/to/your/mcp/atlassian-dc-mcp-server -c /path/to/your/mcp/config.yaml -auth-mode=header
   Restart=always
   RestartSec=10
   Environment=HOME=/home/your-username

   [Install]
   WantedBy=default.target
   ```

4. Reload systemd user configuration and enable the service:
   ```bash
   systemctl --user daemon-reload
   systemctl --user enable atlassian-dc-mcp.service
   ```

5. Start the service:
   ```bash
   systemctl --user start atlassian-dc-mcp.service
   ```

6. Check service status:
   ```bash
   systemctl --user status atlassian-dc-mcp.service
   ```

7. If you want the user service to run automatically at system startup (even when the user is not logged in), run the following command:
   ```bash
   sudo loginctl enable-linger your-username
   ```

Services configured this way will run as the current user, using the current user's permissions and environment variables, which is more suitable for single-user environments.