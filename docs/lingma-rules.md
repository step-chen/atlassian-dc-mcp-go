# Using Lingma Rules

This project includes predefined Lingma rules that demonstrate how to use the Atlassian Data Center MCP service for code review tasks.

## Rule Files

The `.lingma/rules` directory in the project contains the following rule files:

- [code_review.md](../.lingma/rules/code_review.md) - An example configuration showing how to use the Atlassian Data Center MCP service for automated Bitbucket PR code reviews

## How to Use

To use these rules, you need to copy them to your user's personal path. The specific steps are as follows:

### 1. Copy Rule Files

Copy the rule files from the `.lingma/rules` directory to the `.lingma/rules` folder in your user directory:

```bash
# Create target directory
mkdir -p ~/.lingma/rules

# Copy rule files
cp .lingma/rules/*.md ~/.lingma/rules/
```

:warning: **Important Note**: After copying or modifying any rule files, you must restart VS Code for the changes to take effect. The Lingma extension caches rule configurations at startup, so any modifications require a restart to be recognized.

### 2. Using Rules in Tongyi Lingma

After completing the above steps, you can use these rules in Tongyi Lingma through trigger words:

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

1. Copy the [code_review.md](../.lingma/rules/code_review.md) file as a template
2. Modify the rule content as needed for your specific use cases
3. Save the new rule file to the `~/.lingma/rules/` directory
4. **Important**: Restart VS Code for the changes to take effect
5. Use the new rule in Tongyi Lingma

By using this example rule, you can see how to integrate the MCP service with Lingma to automate code review processes, ensuring consistency in code quality and security standards.