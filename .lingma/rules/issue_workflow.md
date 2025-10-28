---
trigger: manual
name: issue_workflow
---

# Issue Workflow Rules

## Overview

This rule defines the development workflow based on Jira Issues.

## Available Commands

### 1. Link PR

```
/link_pr [options]
```

options:
    -i ISSUE_KEY: Jira Issue key (required)
    -p PR_URL: Complete URL of the PR (required if -u not provided)
    -u URL: Bitbucket URL to extract PR information (required if -p not provided)
        From a path like "projects/PROJECT_KEY/repos/REPO_SLUG/pull-requests/PR_ID", 
        the system can extract PROJECT_KEY, REPO_SLUG, and PR_ID to identify the PR automatically.

Examples:
- `/link_pr -i JIRA-123 -p https://bitbucket.example.com/projects/MYPROJECT/repos/myrepo/pull-requests/456` - Link PR with JIRA-123
- `/link_pr -i JIRA-123 -u projects/MYPROJECT/repos/myrepo/pull-requests/456` - Link PR with JIRA-123 using URL path

### 2. Update Issue Status

```
/update [options]
```

options:
    -i ISSUE_KEY: Jira Issue key (required)
    -s STATUS: New status (required)
    -c COMMENT: Comment (optional)

Examples:
- `/update -i JIRA-123 -s In Progress` - Update JIRA-123 status to "In Progress"
- `/update -i JIRA-123 -s Done -c "Finished development and tests"` - Update JIRA-123 status to "Done" with a comment

## Workflow Steps

1. Use `/start_dev` command to start development based on an Issue
2. Regularly use `/pre_commit` to check code quality during development
3. Create a PR and use `/link_pr` to link it after development is complete
4. Update the Issue status using `/update` after the PR is merged

## Best Practices

1. Ensure each Issue has clear acceptance criteria
2. Read the Issue description and comments carefully before development
3. Use `/pre_commit` to check before committing code
4. Update Issue status and add relevant comments in a timely manner