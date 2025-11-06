# 通过 Docker 启动和配置 MCP 服务

本文档将指导您如何使用 Docker 快速启动和配置 Atlassian Data Center MCP 服务。

## 前提条件

- 确保您的系统已安装 Docker 和 Docker Compose。

## 步骤

### 1. 创建并配置 `config.yaml`

首先，您需要创建一个本地配置文件。

1. 从示例文件复制配置文件：
   ```bash
   cp config.yaml.example config.yaml
   ```

2. 编辑 `config.yaml` 文件，填入您自己的 Atlassian 产品（Jira、Confluence、Bitbucket）连接信息，例如：
   - `url`：您的 Atlassian 产品实例地址
   - `token`：对应的 API Token

注意：使用 Docker 时，如果您计划使用基于头部的认证，可以将 token 字段留空。

### 2. 修改 `docker-compose.yml`

为了让 Docker 容器能够读取您的本地配置，您需要修改 `docker-compose.yml` 文件，将 `config.yaml` 文件映射到容器中。

打开 `docker-compose.yml` 文件，在 `mcp-server` 服务的 `volumes` 部分添加以下映射：

```yaml
services:
  mcp-server:
    # ...（其他配置保持不变）
    image: ghcr.io/step-chen/atlassian-dc-mcp-go:latest
    ports:
      - "8090:8090"
    volumes:
      - ./config.yaml:/app/config.yaml # <-- 添加或修改此行
    # ...（其他配置保持不变）
```

要启用基于头部的认证模式，您需要添加一个命令参数：

```yaml
services:
  mcp-server:
    # ...（其他配置保持不变）
    image: ghcr.io/step-chen/atlassian-dc-mcp-go:latest
    command: ["-auth-mode=header"]
    ports:
      - "8090:8090"
    volumes:
      - ./config.yaml:/app/config.yaml
    # ...（其他配置保持不变）
```

这将您在项目根目录创建的 `config.yaml` 文件链接到容器中的 `/app/config.yaml` 路径。服务启动时将加载此配置。

### 3. 启动 MCP 服务

完成配置后，运行以下命令从 GitHub Container Registry 拉取最新的 `latest` 镜像并启动服务：

```bash
docker compose up -d
```

该命令将在后台启动 MCP 服务。您可以运行 `docker compose logs -f` 查看实时日志，确保服务正常启动。

### 4. 配置 AI 助手

最后，在您的 AI 助手或相关客户端配置中，添加或修改 MCP Server 地址，指向本地运行的 Docker 容器。

```json
{
  "mcpServers": {
    "atlassian-dc-mcp-sse": {
      "url": "http://localhost:8090/sse"
    }
  }
}
```

`http://localhost:8090` 是在 `docker-compose.yml` 中配置的映射端口，`/sse` 是 MCP 服务的 SSE 端点。

当使用基于header的认证时，您的客户端需要传递适当的头部：
- `Jira-Token`：Jira的API令牌
- `Confluence-Token`：Confluence的API令牌
- `Bitbucket-Token`：Bitbucket的API令牌

以下是基于header认证的完整配置示例：

```json
{
  "mcpServers": {
    "atlassian-dc-mcp-sse": {
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
    "atlassian-dc-mcp-sse": {
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

## 停止服务

如果您需要停止服务，可以运行以下命令：

```bash
docker compose down
```

## 示例文档

- 有关 config.yaml 的示例，请参阅 [config.yaml.example](config.yaml.example)
- 有关 docker-compose.yml 的示例，请参阅 [docker-compose.yml](docker-compose.yml)