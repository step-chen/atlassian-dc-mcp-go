# 通过 Docker 启动和配置 MCP 服务

本文档将指导您如何使用 Docker 快速启动和配置 Atlassian Data Center MCP 服务。

## 先决条件

- 确保您的系统中已安装 Docker 和 Docker Compose。

## 步骤

### 1. 创建并配置 `config.yaml`

首先，您需要创建一个本地的配置文件。

1.  从示例文件复制一份配置文件：
    ```bash
    cp config.yaml.example config.yaml
    ```

2.  编辑 `config.yaml` 文件，填入您自己的 Atlassian 产品（Jira, Confluence, Bitbucket）的连接信息，例如：
    -   `url`: 您的 Atlassian 产品实例地址
    -   `token`: 对应的 API Token

### 2. 修改 `docker-compose.yml`

为了让 Docker 容器能读取到您的本地配置，需要修改 `docker-compose.yml` 文件，将 `config.yaml` 文件映射到容器内部。

打开 `docker-compose.yml` 文件，在 `mcp-server` 服务的 `volumes` 部分添加以下映射：

```yaml
services:
  mcp-server:
    # ... (其他配置保持不变)
    image: ghcr.io/step-chen/atlassian-dc-mcp-go:latest
    ports:
      - "8090:8090"
    volumes:
      - ./config.yaml:/app/config.yaml # <-- 添加或修改这一行
    # ... (其他配置保持不变)
```

这会将您在项目根目录中创建的 `config.yaml` 文件链接到容器的 `/app/config.yaml` 路径，服务启动时会加载此配置。

### 3. 启动 MCP 服务

完成配置后，运行以下命令来从 GitHub Container Registry 拉取最新的 `latest` 镜像并启动服务：

```bash
docker compose up -d
```

该命令会在后台启动 MCP 服务。您可以运行 `docker compose logs -f` 来查看实时日志，确保服务正常启动。

### 4. 配置 AI 助手

最后，在您的 AI 助手或相关客户端的配置中，添加或修改 MCP Server 的地址，使其指向本地运行的 Docker 容器。

```json
{
  "mcpServers": {
    "atlassian-dc-mcp-sse": {
      "url": "http://localhost:8090/sse"
    }
  }
}
```

`http://localhost:8090` 是 `docker-compose.yml` 中配置的映射端口，`/sse` 是 MCP 服务的 SSE 端点。

## 停止服务

如果您需要停止服务，可以运行以下命令：

```bash
docker compose down
```

## 示例文档

- config.yaml的示例文档请参考 [config.yaml.example](config.yaml.example)
- docker-compose.yml的示例文档请参考 [docker-compose.yml](docker-compose.yml)