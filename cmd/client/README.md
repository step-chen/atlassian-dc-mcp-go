# MCP Client

This is a simple MCP client that connects to the Atlassian Data Center MCP server and calls the `jira_get_issue` tool to retrieve the content of Jira issue HAD-10228.

## Usage

To run the client:

```bash
go run cmd/client/main.go
```

## How it works

1. The client initializes the logger and loads the configuration
2. It creates an MCP client using the go-sdk
3. It connects to the server using stdio transport
4. It calls the `jira_get_issue` tool with the issue key "HAD-10228"
5. It prints the result to the console

## Notes

- This client is designed to work with the stdio transport, which means it communicates with the server through standard input/output
- The server must be running for the client to work
- The client has a 30-second timeout for the entire operation

## Debugging with Delve

To debug the server with Delve:

1. Start the server with Delve using the provided script:
   ```bash
   ./start_delve.sh
   ```

2. Connect with VSCode using the 'Connect to Delve' configuration

3. Run the client to interact with the debug server:
   ```bash
   go run cmd/client/main.go
   ```

Note: The client connects to the server via stdio transport, so it will work with both regular and Delve-debugged server instances.