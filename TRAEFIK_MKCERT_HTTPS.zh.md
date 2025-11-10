# 使用 Traefik 和 mkcert 实现 HTTPS 访问

本文档介绍如何使用 Traefik 和 mkcert 为 Atlassian DC MCP 服务配置 HTTPS 访问。

## 概述

[mkcert](https://github.com/FiloSottile/mkcert) 是一个用于本地 HTTPS 开发的简单工具，它可以创建本地受信任的开发证书。结合 Traefik 反向代理，我们可以为 MCP 服务提供 HTTPS 访问。

## 前提条件

1. 安装 Docker 和 Docker Compose
2. 安装 mkcert 工具

## 安装 mkcert

### macOS (使用 Homebrew)

```bash
brew install mkcert
brew install nss # 如果使用 Firefox 浏览器
```

### Linux

在 Linux 上，你需要安装 mkcert 的二进制文件：

```bash
# 下载最新版本的 mkcert
wget https://github.com/FiloSottile/mkcert/releases/download/v1.4.4/mkcert-v1.4.4-linux-amd64
chmod +x mkcert-v1.4.4-linux-amd64
sudo mv mkcert-v1.4.4-linux-amd64 /usr/local/bin/mkcert
```

## 创建本地 CA 并生成证书

1. 安装本地 CA：

```bash
mkcert -install
```

这将在系统中安装本地 CA，使所有由 mkcert 生成的证书都被信任。

2. 创建项目目录下的 certs 文件夹：

```bash
mkdir -p ./certs
cd ./certs
```

3. 生成本地开发证书：

```bash
mkcert localhost 127.0.0.1 ::1
```

这将生成两个文件：
- `localhost+2.pem` (证书文件)
- `localhost+2-key.pem` (私钥文件)

## 配置 Traefik

项目中包含了 Traefik 配置文件，位于 [traefik-config/](traefik-config/) 目录中：

1. [traefik-config/traefik.yml](traefik-config/traefik.yml) - Traefik 主配置文件
2. [traefik-config/certs.yml](traefik-config/certs.yml) - 证书配置文件

Traefik 已配置为：
- 在 443 端口监听 HTTPS 请求
- 使用 mkcert 生成的证书
- 处理常规 HTTP 请求和 SSE 流
- 使用 HTTP/2 与 MCP 服务器通信

## 启动服务

使用以下命令启动包含 Traefik 的服务：

```bash
# 构建并启动服务
docker-compose -f docker-compose.traefik.yml up -d

# 查看日志
docker-compose -f docker-compose.traefik.yml logs -f
```

## 验证配置

启动服务后，可以通过以下 URL 访问 MCP 服务：

- HTTPS: https://localhost/sse
- Traefik Dashboard: http://localhost:8080 (仅可从宿主机访问)

## 配置 AI 助手

在 AI 助手中配置 MCP 服务时，使用以下配置：

```json
{
  "mcpServers": {
    "atlassian-dc-mcp-sse": {
      "url": "https://localhost/sse",
      "headers": {
        "Jira-Token": "your-jira-api-token",
        "Confluence-Token": "your-confluence-api-token",
        "Bitbucket-Token": "your-bitbucket-api-token"
      }
    }
  }
}
```