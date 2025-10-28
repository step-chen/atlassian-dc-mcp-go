---
trigger: always_on
name: workflow
---

# Workflow Rules

## Overview

This rule defines standard operations and best practices in the development workflow.

## Available Commands

### 1. Branch Operations

```
/checkout [MyLocalCodePath] [options]
```

MyLocalCodePath: Local directory (optional)

options:
    -i PR_ID: Pull Request ID (required)
    -p PROJECT: Bitbucket project key (required if -u not provided)
    -r REPO: Bitbucket repository slug (required if -u not provided)
    -u URL: Bitbucket URL to extract PROJECT and REPO information (required if PROJECT and REPO not provided)
        From paths like:
        - "projects/PROJECT_KEY/repos/REPO_SLUG/pull-requests/PR_ID" for PR checkout
        - "projects/PROJECT_KEY/repos/REPO_SLUG/browse/PATH" for project/repo detection
        the system can extract PROJECT_KEY and REPO_SLUG to set PROJECT and REPO automatically.

Examples:
- `/checkout -i 123 -p MYPROJECT -r myrepo` - Check out PR #123 from MYPROJECT/myrepo to a temporary directory
- `/checkout -i 123 -p MYPROJECT -r myrepo ./mypr` - Check out PR #123 from MYPROJECT/myrepo to ./mypr directory
- `/checkout -u projects/PROJECT_KEY/repos/REPO_SLUG/pull-requests/PR_ID` - Check out PR with PROJECT_KEY, REPO_SLUG, and PR_ID extracted from the URL path
- `/checkout ./mypr -u projects/PROJECT_KEY/repos/REPO_SLUG/pull-requests/PR_ID` - Check out PR from URL path to ./mypr directory

```
/start-dev [MyLocalCodePath] [options]
```

MyLocalCodePath: Local checkout directory (optional)

options:
    -i ISSUE_KEY: Jira Issue key (required)
    -p PROJECT: Bitbucket project key (optional)
    -r REPO: Bitbucket repository slug (optional)
    -d REMOTE_DIR: Remote branch base directory (optional)
    -u URL: Bitbucket URL to extract PROJECT and REPO information (optional)
        From a path like "projects/PROJECT_KEY/repos/REPO_SLUG", 
        the system can extract PROJECT_KEY and REPO_SLUG to set PROJECT and REPO automatically.

Examples:
- `/start-dev -i JIRA-123` - Start development for JIRA-123 with auto-detected project and repository
- `/start-dev -i JIRA-123 -p MYPROJECT -r myrepo` - Start development for JIRA-123 in MYPROJECT/myrepo
- `/start-dev -i JIRA-123 -u projects/PROJECT_KEY/repos/REPO_SLUG` - Start development with PROJECT_KEY and REPO_SLUG extracted from the URL path

### 2. Code Review

See [code_review.md](code_review.md) for detailed information about the `/review` command.

### 3. Code Analysis

See [code_analysis.md](code_analysis.md) for detailed information about the `/analyze` command.

### 4. Pre-commit Analysis

See [pre_commit.md](pre_commit.md) for detailed information about the `/pre-commit` command.

## Workflow Recommendations

1. Use `/start-dev` command to begin development on a new task
2. Regularly use `/analyze` for code analysis during development
3. Use `/pre-commit` for comprehensive evaluation before committing code
4. Use `/review` for code review after creating a PR

## Best Practices

1. Ensure each feature branch is associated with a Jira Issue
2. Ensure all tests pass before committing code
3. Write clear commit messages
4. Follow team coding standards
- **Clean Architecture**: Layered design with unidirectional dependencies.
- **DRY/KISS/YAGNI**: Avoid duplicate code, keep it simple, implement only necessary features.
- **Concurrency Safety**: Reasonable use of Goroutines and Channels to avoid race conditions.
- **OWASP Security Guidelines**: Prevent SQL injection, XSS, CSRF and other attacks.
- **Code Maintainability**: Modular design with clear package structure and function naming.
- **Workflow**: You always provide a solution first, and only start modifying code after receiving explicit modification instructions from the user.

## Technical Specifications

For detailed technical specifications, see [technical_specifications.md](technical_specifications.md).