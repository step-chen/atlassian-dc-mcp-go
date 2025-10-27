---
trigger: manual
name: issue_workflow
---

# Issue Workflow Rules

## Overview

This rule defines the development workflow based on Jira Issues.

## Available Commands

### 1. Start Development

```
/start-dev ISSUE_KEY [PROJECT REPO [REMOTE_DIR]] [LOCAL_DIR]
```

Start development work based on the specified Jira Issue.

Parameters:
- ISSUE_KEY: Jira Issue key (required)
- PROJECT: Bitbucket project key (optional)
- REPO: Bitbucket repository slug (optional)
- REMOTE_DIR: Remote branch base directory (optional)
- LOCAL_DIR: Local checkout directory (optional)

### 2. Link PR

```
/link-pr ISSUE_KEY PR_URL
```

Link the specified PR with a Jira Issue.

Parameters:
- ISSUE_KEY: Jira Issue key
- PR_URL: Complete URL of the PR

### 3. Update Issue Status

```
/update-issue-status ISSUE_KEY STATUS [COMMENT]
```

Update the status of a Jira Issue.

Parameters:
- ISSUE_KEY: Jira Issue key
- STATUS: New status
- COMMENT: Comment (optional)

### 4. Pre-commit Analysis

```
/pre-commit-analysis [PROJECT] [REPO] [ISSUE_KEY]
```

Analyze current changes before committing and provide comprehensive suggestions based on associated Jira Issue information.

Parameters:
- PROJECT: Bitbucket project key (optional)
- REPO: Bitbucket repository slug (optional)
- ISSUE_KEY: Jira Issue key (optional)

## Workflow Steps

1. Use `/start-dev` command to start development based on an Issue
2. Regularly use `/pre-commit-analysis` to check code quality during development
3. Create a PR and use `/link-pr` to link it after development is complete
4. Update the Issue status using `/update-issue-status` after the PR is merged

## Best Practices

1. Ensure each Issue has clear acceptance criteria
2. Read the Issue description and comments carefully before development
3. Use `/pre-commit-analysis` to check before committing code
4. Update Issue status and add relevant comments in a timely manner