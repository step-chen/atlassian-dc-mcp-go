# Atlassian Data Center MCP (Model Context Protocol)

作者：Stephen Chen

本项目提供了一个基于 Go 语言的 Model Context Protocol (MCP) 服务，用于管理和交互 Atlassian Data Center 产品，包括 Jira、Confluence 和 Bitbucket。它允许您通过统一接口以编程方式管理这些产品，并支持可配置的身份验证和权限。

## 功能特性

- **多产品支持**：Jira、Confluence 和 Bitbucket 的统一接口
- **Model Context Protocol**：通过 Model Context Protocol 暴露所有操作的工具
- **配置管理**：通过文件、环境变量和热重载实现灵活配置
- **身份验证**：每个服务的直接 API 令牌身份验证
- **权限控制**：每个服务的细粒度读/写权限
- **健康监控**：所有服务的内置健康检查
- **Docker 支持**：可直接用于容器化部署

## 运行应用程序

您可以通过多种方式运行应用程序：直接使用 Go、先构建二进制文件或使用 Docker。

有关运行应用程序的详细说明：
- [使用 Docker 运行](README.docker.zh.md) - 使用 Docker 和 Docker Compose 运行服务的说明
- [本地运行](README.local.zh.md) - 直接在系统上运行服务的说明

以下是用于开发目的的基本命令：

### 直接执行

```bash
# 运行服务器
go run cmd/server/main.go

# 使用自定义配置文件路径运行
go run cmd/server/main.go -c /path/to/your/config.yaml
```

### 构建和运行二进制文件

项目使用 Makefile 简化构建过程。所有二进制文件都放置在 `dist` 目录中。

```bash
# 显示所有可用命令
make help

# 为当前操作系统构建服务器二进制文件
make build

# 构建静态链接的发布二进制文件
make release
```

构建项目后，您可以直接从 `dist` 目录运行二进制文件：

```bash
# 运行服务器
./dist/atlassian-dc-mcp-server

# 使用自定义配置文件路径运行
./dist/atlassian-dc-mcp-server -c /path/to/your/config.yaml
```

## 配置

应用程序需要配置文件才能运行。默认情况下，它会在当前目录中查找 `config.yaml`，但您可以使用 `-c` 或 `--config` 标志指定不同的路径。

示例配置文件以 `config.yaml.example` 的形式提供。将此文件复制到 `config.yaml` 并使用您的设置进行编辑：

```bash
cp config.yaml.example config.yaml
```

配置文件已包含所有可用设置的示例并有详细文档。有关详细配置选项，请参阅 [config.yaml.example](config.yaml.example) 文件。

## 工具文档

### Jira 工具

用于与 Jira 交互的工具：
- 获取当前用户信息
- 获取问题
- 创建问题
- 以及更多

### Confluence 工具

用于与 Confluence 交互的工具：
- 获取当前用户信息
- 获取内容
- 搜索内容
- 以及更多

### Bitbucket 工具

用于与 Bitbucket 交互的工具：
- 获取当前用户信息
- 获取仓库
- 获取提交
- 以及更多

## Lingma 规则

本项目包含预定义的 Lingma 规则，演示了如何使用 Atlassian Data Center MCP 服务进行自动化代码审查任务。有关如何使用这些规则的详细信息，请参阅 [Lingma 规则文档](docs/lingma-rules.md)。

[code_review.md](.lingma/rules/code_review.md) 规则文件提供了一个示例配置，展示了如何利用 Atlassian Data Center MCP 服务进行自动化的 Bitbucket PR 代码审查。这些规则可以帮助您：

- 使用 MCP 服务自动化代码审查流程
- 标准化代码质量和安全检查
- 提高在代码审查期间与 AI 助手的交互效率

## 开发

### 先决条件

- Go 1.20 或更高版本
- Docker（用于容器化）

### 构建

建议使用 Makefile 构建项目，这可以确保所有构建产物都放置在统一的 `dist` 目录中：

```bash
# 构建服务器二进制文件到 dist 目录
make build
```

为特定操作系统构建二进制文件：
```bash
# 为 Linux 构建
make build-linux

# 为 Windows 构建
make build-windows

# 为 macOS 构建
make build-macos
```

使用 make 命令的好处：
- 所有构建产物都放置在 `dist` 目录中
- 自动处理跨平台构建
- 确保一致的构建参数

## 贡献

欢迎贡献！请随时提交 Pull Request。

## 许可证

本项目基于 MIT 许可证 - 有关详细信息请参阅 [LICENSE](LICENSE) 文件。

## 参考资料

- [Confluence REST API](https://developer.atlassian.com/server/confluence/rest/v1010/intro/#about)
- [Jira REST API](https://developer.atlassian.com/server/jira/platform/rest/v11000/intro/#gettingstarted)
- [Bitbucket REST API](https://developer.atlassian.com/server/bitbucket/rest/v1000/intro/#about)
- [Model Context Protocol Go SDK](https://github.com/modelcontextprotocol/go-sdk)