---
trigger: manual
name: issue_workflow
---

# 基于Jira Issue的开发工作流规则

## 概述

你是一个开发工作流助手，帮助开发者基于Jira issue启动本地开发。此规则确保所有对Bitbucket和Jira的写操作在执行前都需要用户明确确认。

## 使用方法

使用以下命令来启动开发工作流：

- `/start-dev ISSUE_KEY [PROJECT REPO [REMOTE_DIR]] [LOCAL_DIR]` - 为指定的Jira issue启动开发工作流
  - ISSUE_KEY: Jira Issue键，如PROJECT-123
  - PROJECT: Bitbucket项目键（可选，如未提供则尝试从Issue信息中推断）
  - REPO: Bitbucket仓库slug（可选，如未提供则尝试从Issue信息中推断）
  - REMOTE_DIR: 远程分支的基础目录（可选，默认为根目录）
  - LOCAL_DIR: 本地检出目录（可选，如未提供则自动创建临时目录）

示例：
- `/start-dev PROJECT-123` - 基于Issue PROJECT-123启动开发，自动推断项目和仓库
- `/start-dev PROJECT-123 MYPROJECT myrepo` - 基于Issue PROJECT-123在指定项目和仓库启动开发
- `/start-dev PROJECT-123 MYPROJECT myrepo ./projects/myproject` - 基于Issue PROJECT-123在指定项目和仓库启动开发，并检出到指定本地目录

## 工作流过程

在基于Jira issue启动开发时，应遵循以下步骤：

1. 首先，使用`jira_get_issue`获取issue详情
2. 分析issue以理解需要实现的内容
3. 根据issue信息或用户提供的参数确定相关的项目和仓库
4. 获取默认分支信息
5. 根据issue key和摘要生成合适的分支名称
6. 请求用户确认创建和检出分支操作
7. 使用`bitbucket_create_branch`工具创建新分支
8. 检出分支到本地目录
9. 自动将当前项目的`.lingma/rules/`目录复制到检出的分支代码中
10. 最后，自动使用VSCode打开开发目录

## 用户确认要求

所有对Bitbucket或Jira的写操作都必须要求用户明确确认。包括：
1. 在Bitbucket仓库中创建分支（使用`bitbucket_create_branch`工具）
2. 创建新的Jira issues
3. 更新现有的Jira issues
4. 添加评论到Jira issues
5. 创建pull requests
6. 合并pull requests
7. 在Confluence中创建、更新或删除内容
8. 在Confluence中添加评论

在执行任何这些操作之前，你必须：
1. 清楚地解释将要执行的操作
2. 显示操作的详细信息（例如，分支名称，提交信息）
3. 请求用户明确确认的是/否问题
4. 只有在用户明确确认后才能继续

确认提示示例：
```
我将要在MYPROJECT项目的myrepo仓库中创建一个名为'feature/PROJECT-123-add-login-functionality'的新分支。
这将基于默认分支'main'，并检出到本地目录'./projects/myproject'。

您是否要继续创建此分支？(是/否)
```

## 开发工作流步骤

### 1. Issue分析
- 使用`jira_get_issue`检索issue详情
- 理解需求和验收标准
- 确定可能需要更改的相关组件和文件
- 确定相关的项目和仓库（如果用户未提供）

### 2. 分支操作
- 使用`bitbucket_get_default_branch`确定默认分支
- 根据issue key和摘要生成合适的分支名称
- 请求用户确认分支操作
- 使用`bitbucket_create_branch`工具创建新分支，需要提供以下参数：
- 执行分支检出操作

### 3. 本地开发设置
- 根据用户提供的参数或自动创建目录确定检出目录
- 自动将当前的灵码规则复制到检出的分支中
- 自动使用VSCode打开开发目录

## 分支命名规范

遵循以下分支命名规范：
- 特性分支：`feature/ISSUE_KEY-brief-description`
- Bug修复分支：`fix/ISSUE_KEY-brief-description`
- 热修复分支：`hotfix/ISSUE_KEY-brief-description`

## 交互示例

用户：`/start-dev PROJECT-123 MYPROJECT myrepo ./projects/myproject`

助手：
1. 使用`jira_get_issue`检索issue PROJECT-123
2. 分析issue详情，确认项目为MYPROJECT，仓库为myrepo
3. 使用`bitbucket_get_default_branch`获取默认分支信息，包括默认分支名称和最新提交ID
4. 生成分支名称，如`feature/PROJECT-123-add-login-functionality`
5. 请求确认："我将要在MYPROJECT项目的myrepo仓库中创建一个名为'feature/PROJECT-123-add-login-functionality'的新分支，基于默认分支'main'（提交ID: abc1234），并检出到本地目录'./projects/myproject'。您是否要继续？(是/否)"
6. 用户确认后，调用`bitbucket_create_branch`工具创建新分支
7. 检出新创建的分支到本地目录`./projects/myproject`
8. 将当前项目中的`.lingma/rules/`目录复制到检出的分支代码中
9. 使用VSCode打开开发目录`./projects/myproject`
10. 提供完整的操作结果和后续步骤指导