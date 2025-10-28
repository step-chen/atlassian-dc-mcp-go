---
trigger: always_on
name: analyze
---

# Code Analysis Rule

## Overview

You are a senior code analyst who must follow these rules to analyze code in a specified directory (path) under a specific project and repository, and provide optimization suggestions.
This rule accesses Bitbucket through the MCP interface to automatically analyze code and provide optimization suggestions, including:

1. Code efficiency analysis - checking for potential performance issues, algorithm complexity, etc.
2. Database operation efficiency analysis - focusing on query efficiency and resource usage
3. Hard-coded detection - detecting hard-coded values in the code
4. Code duplication detection - detecting duplicate or similar code segments

## Usage

```
/analyze [MyLocalCodePath] [options]
```

MyLocalCodePath: Local path to the code repository (optional, if provided will analyze local code directly without accessing Bitbucket)

options:
    -p PROJECT: Bitbucket project key (optional)
    -r REPO: Bitbucket repository slug (optional)
    -d PATH: Path to the code directory to analyze under project and repository (optional)
    -u URL: Bitbucket URL to extract PROJECT and REPO information (optional)
        From a path like "projects/PROJECT_KEY/repos/REPO_SLUG/browse/PATH", 
        the system can extract PROJECT_KEY and REPO_SLUG to set PROJECT and REPO automatically.

Examples:
- `/analyze /path/to/local/code` - Analyze local code directly without accessing Bitbucket
- `/analyze -p MYPROJECT -r myrepo -d internal/service` - Analyze code in the internal/service directory under the myrepo repository of MYPROJECT project
- `/analyze -u projects/PROJECT_KEY/repos/REPO_SLUG/browse/PATH` - Analyze code with PROJECT_KEY, REPO_SLUG AND PATH extracted from the URL path

## Analysis Process

When analyzing code, you should follow these steps:

1. If MyLocalCodePath is provided, analyze local code directly without accessing Bitbucket
2. If MyLocalCodePath is not provided:
   - First retrieve basic repository information using PROJECT and REPO parameters or extracting from URL
   - Get the file list under the specified directory (using PATH parameter or extracting from URL), recording every file path to be analyzed
   - For each directory encountered, recursively get files in subdirectories, recording every file path to be analyzed
   - For each file, retrieve the file content
   - Analyze the file content and identify the following issues:
     - Code efficiency issues
     - Database operation efficiency issues
     - Hard-coded values
     - Duplicate or similar code segments
3. Review commit history and changes to understand the code history (for Bitbucket repository analysis)

## Analysis Focus

### Code Efficiency Analysis
- Check for code with high cyclomatic complexity
- Identify algorithms that may cause performance issues
- Check for unnecessary resource consumption

### Database Operation Efficiency Analysis
- Detect N+1 query issues
- Identify queries missing indexes
- Check if transactions are used properly

### Hard-coded Detection
- Detect hard-coded strings in the code
- Identify hard-coded numeric values
- Find hard-coded URLs that should use constants or configurations

### Code Duplication Detection
- Identify completely duplicated code segments
- Detect structurally similar but literally different code
- Suggest refactoring duplicate code into reusable functions or components

## Analysis Report Requirements

All analysis comments must include specific improvement suggestions. For each identified issue:
1. Clearly identify the problem in the current code
2. Provide concrete examples of fixes
3. Explain why the suggested changes are better
4. Reference the exact location of the relevant code

Example analysis comment:
```
Issue: Hard-coded database timeout value detected
Severity: Medium
Location: database/connection.go:45
Problem: The code has a hard-coded 30-second timeout value, which makes it difficult to adjust in different environments (development, testing, production).

Current code:
db, err := sql.Open("mysql", dsn)
db.SetConnMaxLifetime(time.Second * 30)

Suggested improvement:
const (
    DefaultDBTimeout = 30 * time.Second
)

// Or read from configuration file
db, err := sql.Open("mysql", dsn)
db.SetConnMaxLifetime(config.DBTimeout)

This change replaces hardcoded values with constants or configuration values, making the code more maintainable and configurable.
```

Analysis reports should include: issue description, severity level, fix suggestions with code examples, related code location, and analysis basis.