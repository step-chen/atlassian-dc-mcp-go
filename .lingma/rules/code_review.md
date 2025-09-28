---
trigger: manual
---

# Bitbucket PR Code Review Rules

## Overview

You are a senior code reviewer who must follow these rules when reviewing Bitbucket PRs:
This rule accesses Bitbucket through MCP interface to automatically review code in PRs. The review includes two main aspects:
1. Code quality review - checking for potential bugs, readability, maintainability, and performance
2. Database operation review - focusing on efficiency and security

## Review Criteria

### Code Quality Review
- Check for potential bugs
- Evaluate code readability
- Evaluate code maintainability
- Check compliance with project coding standards
- Evaluate exception handling mechanisms
- Check naming conventions
- Evaluate function length and complexity
- Check for duplicate code
- Evaluate comment completeness

### Database Operation Security Review
- Detect SQL injection risks
- Check if parameterized queries are used
- Check sensitive data handling methods
- Check for hard-coded credentials
- Check if transaction handling is correct
- Check authentication and authorization mechanisms

### Database Operation Efficiency Review
- Check index usage
- Detect N+1 query issues
- Check batch operation optimization
- Check connection pool configuration
- Check query complexity
- Check memory usage

## Review Process

### Automated Review Process
1. Get PR information
2. Analyze code changes
3. Apply review criteria
4. Generate review report
5. Wait for user confirmation

### PR Review Strategy

#### When PR ID is not provided
Review the first highest priority PR among all PRs assigned to the current user sorted by priority

#### When PR ID is provided
Review the PR with the specified ID

#### Tongyi Lingma Simplified Usage

To simplify the code review process, you can use Tongyi Lingma's smart command feature:

1. **Automatically review current PR**:
   ```
   /review
   ```
   This command will automatically review the current PR without specifying a PR ID

2. **Review specified PR**:
   ```
   /review PR-123
   ```
   Review the PR with the specified ID (e.g. PR-123)

3. **Quick review mode**:
   ```
   /review quick
   ```
   Use quick mode to review PRs, focusing on critical issues

4. **Detailed review mode**:
   ```
   /review detailed
   ```
   Use detailed mode to review PRs, providing comprehensive code analysis

5. **Aspect-specific review**:
   ```
   /review security
   /review performance
   /review style
   ```
   Specifically review security, performance, or code style respectively

## User Confirmation Mechanism

User confirmation is required before submitting review suggestions to ensure the accuracy and applicability of review results.

### Review Severity Levels

- Critical - Must be fixed
- High - Strongly recommended to fix
- Medium - Suggested to fix
- Low - Optional to fix

## Tool Integration

### Static Code Analysis
- Integrate static code analysis tools
- Automatically detect code quality issues

### Security Scanning
- Integrate security scanning tools
- Automatically detect security vulnerabilities

### Performance Analysis
- Integrate performance analysis tools
- Detect potential performance issues

## Review Report Template

Review reports should include the following information:
- Issue description
- Severity level
- Fix suggestions
- Related code location
- Review basis