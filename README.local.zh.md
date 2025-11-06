# 本地运行 MCP 服务

本文档将指导您如何在本地环境中直接运行 Atlassian Data Center MCP 服务，而无需使用 Docker。

## 步骤

### 1. 下载预编译二进制文件

访问 [GitHub Releases 页面](https://github.com/step-chen/atlassian-dc-mcp-go/releases) 下载适用于您操作系统的最新二进制文件。

### 2. 创建并配置 `config.yaml`

首先，您需要创建一个本地配置文件。

1. 从示例复制配置文件：
   ```bash
   cp config.yaml.example config.yaml
   ```

2. 编辑 `config.yaml` 文件，填入您自己的 Atlassian 产品（Jira、Confluence、Bitbucket）连接信息，例如：
   - `url`：您的 Atlassian 产品实例地址
   - `token`：对应的 API Token

注意：如果您计划使用基于头部的认证，可以将 token 字段留空。

### 3. 运行服务

下载并解压二进制文件后，您可以直接运行服务：
```
./atlassian-dc-mcp-server -c /path/to/your/config.yaml
```

要启用基于头部的认证模式，请使用 `-auth-mode=header` 标志：
```
./atlassian-dc-mcp-server -c /path/to/your/config.yaml -auth-mode=header
```

### 4. 验证服务运行

服务成功启动后，您将看到类似以下的日志输出：

```
INFO[0000] Starting Atlassian DC MCP Server on port 8090
```

默认情况下，服务将监听 `http://localhost:8090`。

### 5. 配置 AI 助手

在您的 AI 助手或相关客户端配置中，添加或修改 MCP Server 地址，指向本地运行的服务。

根据配置文件中启用的传输模式，使用相应的端点：

```json
{
  "mcpServers": {
    "atlassian-dc-mcp-local": {
      "url": "http://localhost:8090/sse"
    }
  }
}
```

或者，如果您启用了 HTTP 模式：

```json
{
  "mcpServers": {
    "atlassian-dc-mcp-local": {
      "url": "http://localhost:8090/mcp"
    }
  }
}
```

当使用基于header的认证时，您的客户端需要传递适当的头部：
- `Jira-Token`：Jira的API令牌
- `Confluence-Token`：Confluence的API令牌
- `Bitbucket-Token`：Bitbucket的API令牌

以下是基于header认证的完整配置示例：

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

或者，您也可以使用环境变量来配置令牌：

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

### 6. 停止服务

要停止服务，只需在终端中按 `Ctrl+C` 即可优雅地关闭服务。

## 故障排除

### 服务启动失败

1. 检查配置文件是否正确配置
2. 确保端口未被其他程序占用
3. 检查 Atlassian 产品 URL 和 API Token 是否正确

### 无法连接到 Atlassian 产品

1. 确保网络连接正常
2. 验证 Atlassian 产品 URL 和 API Token
3. 检查防火墙设置

## 配置用户级服务自启动

### Ubuntu 用户级服务自启动

在 Ubuntu 系统中，您可以使用 systemd 用户服务来配置服务自动启动，这样服务将以当前用户而不是系统级服务运行：

1. 创建 systemd 用户服务目录（如果不存在）：
   ```bash
   mkdir -p ~/.config/systemd/user
   ```

2. 创建 systemd 用户服务文件：
   ```bash
   nano ~/.config/systemd/user/atlassian-dc-mcp.service
   ```

3. 将以下内容添加到文件中（请根据实际路径进行调整）：
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

4. 重新加载 systemd 用户配置并启用服务：
   ```bash
   systemctl --user daemon-reload
   systemctl --user enable atlassian-dc-mcp.service
   ```

5. 启动服务：
   ```bash
   systemctl --user start atlassian-dc-mcp.service
   ```

6. 检查服务状态：
   ```bash
   systemctl --user status atlassian-dc-mcp.service
   ```

7. 如果您希望用户服务在系统启动时自动运行（即使用户未登录），请运行以下命令：
   ```bash
   sudo loginctl enable-linger your-username
   ```

通过这种方式配置的服务将以当前用户身份运行，使用当前用户的权限和环境变量，更适合单用户环境。