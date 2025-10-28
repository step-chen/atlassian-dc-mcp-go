---
trigger: always_on
name: pre_commit
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
/pre_commit MyLocalCodePath [options]
```

MyLocalCodePath: Local path to the code repository with changes.

options:
    -p PROJECT: Bitbucket project key (optional)
    -r REPO: Bitbucket repository slug (optional)
    -k ISSUE_KEY: Jira issue key (optional; auto-extracted from branch name if missing)
    -b BRANCH: Repository branch name (optional; uses default branch, typically controlled/XXX, if missing)
    -u URL: Bitbucket URL to extract PROJECT and REPO information (optional)
        From a path like "projects/PROJECT_KEY/repos/REPO_SLUG/browse/PATH", 
        the system can extract PROJECT_KEY and REPO_SLUG to set -p and -r options automatically.

examples:
- /pre_commit MyLocalCodePath -p MYPROJECT -r MyRepo -k JIRA-123 -b MyBranch - Analyzes current changes in MYPROJECT/MyRepo (branch MyBranch), with JIRA-123 context included.
- /pre_commit MyLocalCodePath -u projects/PROJECT_KEY/repos/REPO_SLUG/browse/PATH - Analyzes current changes with PROJECT_KEY and REPO_SLUG extracted from the URL path

## Analysis Process

When running this analysis, the system will:

1. Retrieve current workspace code changes
2. Analyze code changes for potential issues and improvements
3. If PROJECT, REPO, BRANCH and ISSUE_KEY are provided:
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
/pre_commit MyLocalCodePath -p MYPROJECT -r MyRepo -k JIRA-123 -b MyBranch
```

This command analyzes current changes in the MYPROJECT project's myrepo repository and provides comprehensive suggestions. If JIRA-123 information is available, it will also be considered in the analysis.

## Notes

1. Ensure Atlassian Data Center access credentials are properly configured
2. Ensure sufficient permissions to access relevant projects, repositories, and issues
3. If PROJECT, REPO, and ISSUE_KEY are not provided, only local code changes will be analyzed without Jira Issue information
4. Analysis results are for reference only; final decisions should be made by developers
5. Code analysis is always performed, Jira Issue information is used to provide additional context when all parameters are provided