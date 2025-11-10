# 使用 Nginx 和 mkcert 实现 HTTPS 访问

本文档介绍如何使用 Nginx 和 mkcert 为 Atlassian DC MCP 服务配置 HTTPS 访问。

## 概述

[mkcert](https://github.com/FiloSottile/mkcert) 是一个用于本地 HTTPS 开发的简单工具，它可以创建本地受信任的开发证书。结合 Nginx 反向代理，我们可以为 MCP 服务提供 HTTPS 访问。

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

#### Ubuntu/Debian (使用 apt)

在 Ubuntu/Debian 系统上，您可以直接使用 apt 安装 mkcert：

```bash
sudo apt update
sudo apt install mkcert
```

#### 其他 Linux 发行版

在其他 Linux 发行版上，你需要安装 mkcert 的二进制文件：

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

## 配置 Nginx

项目中已经包含了 Nginx 配置文件 [nginx-config/nginx.conf](nginx-config/nginx.conf)，该配置包含：

1. HTTP 到 HTTPS 的重定向
2. SSL 证书配置
3. 到 MCP 服务的反向代理

## 启动服务

使用以下命令启动包含 Nginx 的服务：

```bash
# 构建并启动服务
docker-compose -f docker-compose.nginx.yml up -d

# 查看日志
docker-compose -f docker-compose.nginx.yml logs -f
```

## 验证配置

启动服务后，可以通过以下 URL 访问 MCP 服务：

- HTTPS: https://localhost/sse
- HTTP: http://localhost (会自动重定向到 HTTPS)

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

## 故障排除

### 证书未被信任

如果浏览器仍然提示证书不受信任，请确保：

1. 已正确运行 `mkcert -install`
2. 重启浏览器
3. 在 Linux 上，可能需要手动安装证书到系统或浏览器中

### 端口冲突

如果 80 或 443 端口已被占用，可以修改 [docker-compose.nginx.yml](docker-compose.nginx.yml) 中的端口映射：

```yaml
ports:
  - "8080:80"
  - "8443:443"
```

然后通过 https://localhost:8443 访问服务。

### 无法连接到 MCP 服务

检查以下几点：

1. MCP 服务是否正常运行：`docker-compose -f docker-compose.nginx.yml logs mcp-server`
2. Nginx 配置是否正确：`docker-compose -f docker-compose.nginx.yml logs nginx-proxy`
3. 网络连接是否正常：确保两个容器在同一个网络中

## 生产环境注意事项

此配置仅适用于开发和测试环境。在生产环境中，应该：

1. 使用有效的 SSL 证书（如 Let's Encrypt）
2. 配置适当的防火墙规则
3. 加强 Nginx 安全配置
4. 使用适当的负载均衡和高可用性配置