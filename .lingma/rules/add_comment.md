---
trigger: always_on
name: add_comment
---

# Jira/Confluence Comment Addition Rules

## Overview

You are a helpful assistant who can add comments to Jira issues or Confluence pages.
This rule accesses Jira and Confluence through MCP interface to automatically add comments or content.
The functionality includes:
1. Adding comments to Jira issues - for task updates, code review feedback, or other communications
2. Adding content to Confluence pages - for documentation updates, meeting notes, or knowledge sharing

## Usage

```
/comment [options] "Your comment content here"
```

options:
    -t TARGET: Target system (jira or confluence) (required)
    -k ISSUE_KEY: Jira issue key (required when -t jira)
    -p PAGE_ID: Confluence page ID (required when -t confluence)
    -s SPACE_KEY: Confluence space key (optional, used with -t confluence)
    -m MODE: Comment mode (comment, append, prepend) (optional, default: comment)
    -i: Improve - indicate it needs AI to improve the comment content before adding it to the issue/page (optional)
    -signature: Add a signature to the comment (Qwen3 assisted) (optional)
    -d "instructions": Provide instructions to the AI model to generate a better comment. If this option is used, the user can only provide instructions, not provide the content. (optional)
    -u URL: Jira or Confluence URL to extract issue key or page information (optional)
        From paths like:
        - "browse/ISSUE_KEY" for Jira issue key extraction
        - "pages/viewpage.action?pageId=PAGE_ID" for Confluence page ID extraction
        - "display/SPACE_KEY/PAGE_TITLE" for Confluence space and page title extraction

Examples:
- `/comment -t jira -k JIRA-123 "This issue has been reviewed and looks good."` - Add comment to JIRA-123
- `/comment -t confluence -p 12345 "Adding important note about this documentation."` - Add content to Confluence page with ID 12345
- `/comment -t confluence -p 12345 -m append "Append this information to the end of the page."` - Append content to Confluence page with ID 12345
- `/comment -t jira -u browse/JIRA-123 "This is an important update."` - Add comment to JIRA-123 with key extracted from URL
- `/comment -t confluence -u pages/viewpage.action?pageId=12345 "Please review this updated documentation."` - Add comment to Confluence page with ID extracted from URL
- `/comment -t jira -k JIRA-123 -i "Please review this code."` - Add an AI-improved comment to JIRA-123
- `/comment -t confluence -p 12345 -signature "Document updated by Qwen3"` - Add content to Confluence page with signature
- `/comment -t jira -k JIRA-123 -d "Summarize the changes in pull request PR-567 and add a comment about potential issues"` - Generate and add a comment based on instructions
- `/comment -t confluence -p 12345 -d "Create a project status report for the mobile app development with sections for timeline, blockers, and next steps"` - Generate and add content based on instructions

### Target Systems

```
jira: Add comment to a Jira issue
confluence: Add content to a Confluence page
```

### Comment Modes

```
comment: Add as a new comment (default)
append: Add content to the end of the page/content
prepend: Add content to the beginning of the page/content
```

### Additional Options

```
-k: Jira issue key (required for Jira)
-p: Confluence page ID (required for Confluence)
-s: Confluence space key (optional for Confluence)
-m: Comment mode (optional, default: comment)
-i: Improve - indicate it needs AI to improve the comment content before adding it to the issue/page (optional)
-signature: Add a signature to the comment (Qwen3 assisted) (optional)
-d: Provide instructions for AI to generate content (cannot be used with explicit content)
-u: URL to extract parameters (optional)
```

### Notes

```
Use -t to specify target system (jira or confluence).
Content to be added should be enclosed in quotes.
When using Confluence, if -s is not provided, the system will try to determine the space automatically.
The -i flag requests AI to improve the comment content before adding it to the target.
The -d flag provides instructions for the AI to generate content from scratch, rather than improving existing content.
The -signature flag adds a "Qwen3 assisted" signature to the comment.
Note that -d cannot be used together with explicit content in quotes.
```

## Process

When adding a comment, you should follow these steps:

1. Parse the command to identify the target system and required parameters
2. If URL is provided, extract the required parameters from it
3. Validate the parameters based on the target system
4. Get the current content of the target (for append/prepend modes)
5. Format the content according to the mode
6. Add the comment/content to the target system
7. Confirm the successful addition of the comment/content

## Comment Guidelines

### Jira Comment Guidelines
- Comments should be clear and concise
- Use appropriate formatting (bold, italics, lists) for better readability
- Reference other issues or users when relevant using @mentions or [~user] format
- Add labels or status updates if needed

### Confluence Content Guidelines
- Content should follow the existing page structure and formatting
- Use appropriate headings and subheadings to organize information
- Include links to related pages, issues, or external resources
- Maintain consistent style with the rest of the documentation

## Content Format Requirements

All comments and content must follow these requirements:
1. Use clear and professional language
2. Include specific details relevant to the target
3. Follow the formatting standards of the target system
4. Provide context for the comment/content when necessary

Example of a good Jira comment:
```
I've reviewed the implementation and it looks good.
A few suggestions:
- Consider adding more error handling in the data processing function
- The documentation could be improved with examples
- Please update the unit tests to cover edge cases

Overall, good work on this feature!
```

Example of a good Confluence addition:
```
h2. Performance Improvements

The following changes were implemented to improve system performance:
* Database query optimization reduced response time by 30%
* Caching mechanism implemented for frequently accessed data
* Removed redundant API calls in the frontend components

These improvements were tested under load conditions and show significant performance gains.
```

When using the -d "instructions" option, the AI will generate content based on your instructions. For example:
```
/comment -t jira -k JIRA-456 -d "Create a comment summarizing the code review findings for PR-789. Include sections for critical issues, suggestions for improvement, and general feedback. Use a professional but constructive tone."
```

Would generate a comment like:
```
Code Review Summary for PR-789

Critical Issues:
- Potential null pointer exception in UserService.java line 45
- Security vulnerability with unvalidated input in AuthController.java

Suggestions for Improvement:
- Consider adding more comprehensive error handling in the data processing functions
- Documentation could be enhanced with usage examples
- Unit tests should cover additional edge cases

General Feedback:
Overall, the implementation is solid and follows our coding standards. With the above adjustments, this would be ready for production.
```

Comments and content should include: clear context, specific details, and follow the target system's formatting standards.