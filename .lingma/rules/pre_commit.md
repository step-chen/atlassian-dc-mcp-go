---
trigger: always_on
name: pre-commit
---

# Pre-commit Analysis Rules

## Overview

This rule analyzes current code changes before committing and provides comprehensive suggestions. If a Jira Issue is associated (when PROJECT, REPO, and ISSUE_KEY parameters are provided), it will also consider issue details for additional context.

The analysis includes:
1. Code changes analysis - checking changed files, lines of code, and potential issues
2. Jira Issue integration - retrieving issue details and comments for additional context (when available)
3. Combined recommendations - providing suggestions based on both code changes and issue requirements

## Usage

```
/pre-commit [PROJECT] [REPO] [ISSUE_KEY]
```

Parameters:
- PROJECT: Bitbucket project key (optional)
- REPO: Bitbucket repository slug (optional)
- ISSUE_KEY: Jira Issue key (optional, will attempt to extract from branch name if all three parameters are provided)

## Analysis Process

When running this analysis, the system will:

1. Retrieve current workspace code changes
2. Analyze code changes for potential issues and improvements
3. If PROJECT, REPO, and ISSUE_KEY are provided:
   - Get detailed Jira Issue information
   - Retrieve comments on the Jira Issue to understand discussion context
4. Provide comprehensive suggestions for improvement, incorporating issue context when available

## Analysis Dimensions

### Code Changes Analysis
- Check number of changed files and lines of code
- Identify file types and scope of impact
- Detect potential issues (hard-coded values, security concerns, etc.)
- Provide improvement suggestions based on code changes

### Jira Issue Analysis
- Retrieve issue description, priority, type and other details
- Get issue comments to understand discussion history
- Analyze requirement compliance with code changes

### Comprehensive Recommendations
- Provide improvement suggestions based on code changes and issue requirements (when available)
- Check if tests need to be added or updated
- Verify if documentation updates are needed
- Validate commit message compliance

## Examples

```
/pre-commit MYPROJECT myrepo JIRA-123
```

This command analyzes current changes in the MYPROJECT project's myrepo repository and provides comprehensive suggestions. If JIRA-123 information is available, it will also be considered in the analysis.

## Notes

1. Ensure Atlassian Data Center access credentials are properly configured
2. Ensure sufficient permissions to access relevant projects, repositories, and issues
3. If PROJECT, REPO, and ISSUE_KEY are not provided, only local code changes will be analyzed without Jira Issue information
4. Analysis results are for reference only; final decisions should be made by developers
5. Code analysis is always performed, Jira Issue information is used to provide additional context when all parameters are provided