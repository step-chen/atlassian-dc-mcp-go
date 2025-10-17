# 使用灵码规则

本项目包含预定义的灵码规则，演示了如何使用 Atlassian Data Center MCP 服务进行代码审查任务。

## 规则文件

项目中的 `.lingma/rules` 目录包含以下规则文件：

- [code_review.md](../.lingma/rules/code_review.md) - 演示如何使用 Atlassian Data Center MCP 服务进行自动化 Bitbucket PR 代码审查的示例配置

## 使用方法

要使用这些规则，必须在项目级别进行配置。灵码规则仅支持项目级别配置，不支持用户级别配置。

### 1. 项目级规则配置

规则已在此项目的 `.lingma/rules` 目录中配置。使用此项目时，灵码将自动加载这些规则。

:warning: **重要提示**：灵码规则仅支持项目级别配置。您不能将这些规则复制到用户级目录，因为它们是专门为本项目的结构和 MCP 服务设计的。

### 2. 在通义灵码中使用规则

打开项目后，您可以通过触发词在通义灵码中使用这些规则：

1. **代码审查规则**：
   - 在对 Bitbucket PR 进行代码审查时，可以使用 `/review` 命令
   - [code_review.md](../.lingma/rules/code_review.md) 规则文件提供了如何利用 Atlassian Data Center MCP 服务进行自动化代码审查的示例
   - 支持多种审查模式：
     - `/review` - 自动审查当前 PR
     - `/review PR-123` - 审查指定 ID 的 PR
     - `/review quick` - 快速审查模式
     - `/review detailed` - 详细审查模式
     - `/review security` - 安全专项审查
     - `/review performance` - 性能专项审查
     - `/review style` - 代码风格审查

## 规则说明

### 代码审查规则

[code_review.md](../.lingma/rules/code_review.md) 文件包含一个示例配置，展示了如何使用 Atlassian Data Center MCP 服务进行自动化 Bitbucket PR 代码审查：

- 代码质量审查标准
- 数据库操作安全审查标准
- 数据库操作效率审查标准
- 自动化审查流程
- 用户确认机制

此示例演示了如何将 MCP 服务与灵码集成，创建自动化的代码审查工作流。该规则利用 MCP 服务通过 API 接口访问 Bitbucket，自动审查 PR 中的代码变更。

## 自定义规则

您也可以基于提供的示例创建自己的自定义规则：

1. 在项目的 `.lingma/rules/` 目录中创建新规则文件
2. 遵循现有规则文件中显示的格式
3. **重要**：重启 VS Code 以使更改生效
4. 在通义灵码中使用新规则

通过使用此示例规则，您可以看到如何将 MCP 服务与灵码集成，以自动化代码审查流程，确保代码质量和安全标准的一致性。