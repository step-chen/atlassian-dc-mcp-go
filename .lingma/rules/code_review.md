---
trigger: always_on
name: review
---

# Bitbucket PR Code Review Rules

## Overview

You are a senior code reviewer who must follow these rules when reviewing Bitbucket PRs.
This rule accesses Bitbucket through MCP interface to automatically review code in PRs.
The review includes:
1. Code quality review - checking for potential bugs, readability, maintainability, and performance
2. Database operation review - focusing on efficiency and security
3. Jira Issue requirement compliance review - checking if the implementation matches the requirements
4. Confluence documentation review - checking if related documentation is updated or created

## Usage

```
/review [options]
```

options:
    -p PROJECT: Bitbucket project key (optional)
    -r REPO: Bitbucket repository slug (optional)
    -i PR_ID: Pull Request ID (optional)
    -m MODE: Review mode (quick, detailed, security, performance, style, reportIssueOnly) (optional)
    -a: Add comments directly to relevant lines, labeled: "Qwen3: [your comment]" (optional)
    -u URL: Bitbucket URL to extract PROJECT, REPO, and PR information (optional)
        From a path like "projects/PROJECT_KEY/repos/REPO_SLUG/pull-requests/PR_ID", 
        the system can extract PROJECT_KEY, REPO_SLUG, and PR_ID to identify the PR automatically.

Examples:
- `/review` - Auto-review current PR
- `/review -p PROJECT -r REPO -i PR-123` - Review the specified PR
- `/review -m quick -a -p PROJECT -r REPO -i PR-123` - Quick review of the specified PR with direct comments
- `/review -u projects/PROJECT_KEY/repos/REPO_SLUG/pull-requests/PR_ID` - Review PR with PROJECT_KEY, REPO_SLUG, and PR_ID extracted from the URL path
- `/review -m quick -a -u projects/PROJECT_KEY/repos/REPO_SLUG/pull-requests/PR_ID` - Quick review with direct comments using URL path

### Review Modes

```
quick: Focus on critical issues (default)
detailed: Comprehensive analysis
security: Check security aspects
performance: Review performance
style: Assess code style
reportIssueOnly: Add PR comments only for items needing improvement—no comments required for positive change evaluations.

```

### Additional Options

```
-a: Add comments directly to relevant lines, labeled: "Qwen3: [your comment]"

```

### Notes

```
Use -p PROJECT -r REPO -i PR-ID format (project key, repo slug, PR ID) for targeted commands.
-m specifies review focus; -a adds direct comment behavior.

```

## Review Process

When reviewing a PR, you should follow these steps:

1. First, get the basic information about the PR
2. Get existing comments and activities to understand the discussion context
3. Get the related Jira issue information and Jira issue comments to understand requirements and acceptance criteria
4. Search for related Confluence documentation to understand the context and expected behavior
5. Get a summary of changed files
6. Get the complete diff for comprehensive review
7. For issues found, check if similar comments already exist to avoid duplication
8. Add new comments directly to the PR for issues that haven't been mentioned

## Review Criteria

### Code Quality Review
- Check for potential bugs
- Evaluate code readability and maintainability
- Check compliance with project coding standards
- Evaluate exception handling mechanisms
- Check naming conventions
- Evaluate function length and complexity
- Check for duplicate code
- Evaluate comment completeness

### Database Operation Review
- Detect SQL injection risks and check for parameterized queries
- Check sensitive data handling and hard-coded credentials
- Check transaction handling and authentication mechanisms
- Check index usage and N+1 query issues
- Check batch operation optimization and query complexity
- Check memory usage

### Jira Issue Requirement Compliance Review
- Check if all requirements in the Jira issue have been implemented
- Verify that acceptance criteria are met
- Ensure the implementation aligns with the issue description
- Check if any additional changes not mentioned in the issue are introduced

### Confluence Documentation Review
- Check if related documentation is updated when code changes affect functionality
- Verify that new features have appropriate documentation
- Ensure code comments and documentation are consistent
- Check if links to relevant Confluence pages are provided where necessary

## Review Report Requirements

All review comments must include specific code improvement suggestions. For each identified issue:
1. Clearly identify the problem in the current code
2. Provide a concrete code example showing how to fix the issue
3. Explain why the suggested change is better
4. Reference relevant lines of code with precise line numbers

Example of a good review comment:
```
Issue: Potential nil pointer dereference
Severity: High
Location: file.go:25
Problem: The result variable might be nil if the operation fails, which would cause a panic when accessing result.Data.

Current code:
if result.Status == "success" {
    fmt.Println(result.Data.Name)
}

Suggested improvement:
if result != nil && result.Status == "success" {
    fmt.Println(result.Data.Name)
}

This change prevents a potential nil pointer dereference by checking if result is not nil before accessing its fields.
```

Review reports include: issue description, severity level, fix suggestions with code examples, related code location, and review basis.