# Using Lingma Rules

This project includes predefined Lingma rules that demonstrate how to use the Atlassian Data Center MCP service for code review tasks.

## Rule Files

The `.lingma/rules` directory in the project contains the following rule files:

- [code_review.md](../.lingma/rules/code_review.md) - An example configuration showing how to use the Atlassian Data Center MCP service for automated Bitbucket PR code reviews

## How to Use

To use these rules, they must be configured at the project level. Lingma rules only support project-level configuration, not user-level configuration.

### 1. Project-Level Rule Configuration

The rules are already configured in this project's `.lingma/rules` directory. When using this project, Lingma will automatically load these rules.

:warning: **Important Note**: Lingma rules only support project-level configuration. You cannot copy these rules to a user-level directory as they are specifically designed to work with this project's structure and MCP services.

### 2. Using Rules in Tongyi Lingma

Once you have the project open, you can use these rules in Tongyi Lingma through trigger words:

1. **Code Review Rules**:
   - When conducting code reviews on Bitbucket PRs, you can use the `/review` command
   - The [code_review.md](../.lingma/rules/code_review.md) rule file provides an example of how to leverage the Atlassian Data Center MCP service for automated code reviews
   - Supports multiple review modes:
     - `/review` - Automatically review the current PR
     - `/review PR-123` - Review the PR with the specified ID
     - `/review quick` - Quick review mode
     - `/review detailed` - Detailed review mode
     - `/review security` - Security-specific review
     - `/review performance` - Performance-specific review
     - `/review style` - Code style review

## Rule Description

### Code Review Rule

The [code_review.md](../.lingma/rules/code_review.md) file contains an example configuration showing how to use the Atlassian Data Center MCP service for automated Bitbucket PR code reviews:

- Code quality review criteria
- Database operation security review criteria
- Database operation efficiency review criteria
- Automated review process
- User confirmation mechanism

This example demonstrates how to integrate the MCP service with Lingma to create automated code review workflows. The rule leverages the MCP service to access Bitbucket through its API interface, automatically reviewing code changes in pull requests.

## Custom Rules

You can also create your own custom rules based on the provided example:

1. Create new rule files in the project's `.lingma/rules/` directory
2. Follow the format shown in the existing rule files
3. **Important**: Restart VS Code for the changes to take effect
4. Use the new rule in Tongyi Lingma

By using this example rule, you can see how to integrate the MCP service with Lingma to automate code review processes, ensuring consistency in code quality and security standards.