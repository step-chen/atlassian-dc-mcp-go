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

To use this code analysis feature, use the following command:

- `/analyze PROJECT REPO PATH` - Analyze code in the specified path under project and repository

Examples:
- `/analyze MYPROJECT myrepo internal/service` - Analyze code in the internal/service directory under the myrepo repository of MYPROJECT project

## Analysis Process

When analyzing code, you should follow these steps:

1. First use `bitbucket_get_repository` to get basic repository information
2. Use `bitbucket_get_files` to get the file list under the specified directory, recording every file path to be analyzed
3. For each directory encountered, recursively use `bitbucket_get_files` to get files in subdirectories, recording every file path to be analyzed
4. For each file, use `bitbucket_get_file_content` to get the file content
5. Analyze the file content and identify the following issues:
   - Code efficiency issues
   - Database operation efficiency issues
   - Hard-coded values
   - Duplicate or similar code segments
6. Use `bitbucket_get_commits` and `bitbucket_get_commit_changes` to understand the code history changes

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