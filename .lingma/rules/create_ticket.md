# Jira Ticket Creation Rules

## Overview

You are a helpful assistant who can create tickets in Jira.
This rule accesses Jira through MCP interface to automatically create tickets with a standardized template.
The functionality includes:
1. Creating tickets in Jira projects with a predefined template structure (IMPORTANT)
2. Ensuring tickets contain all necessary information for developers to work on the issue
3. Standardizing ticket format for better readability and processing

## Usage

```
/create-ticket [options] "Ticket title"
```

options:
    -p PROJECT: Jira project key (required)
    -t ISSUE_TYPE: Issue type (PR, CR, Task, Story, etc.) (optional, default: PR) IMPORTANT: If not specified, the default type is "PR".
    -template: Use the predefined template in this file to create ticket. (optional)
    -i: Improve - indicate it needs AI to improve the ticket content before creating (optional)
    -r REF_TICKET: Reference ticket key, if provided, then the newly created ticket format should like the ref ticket. (optional)
    -d "instructions": Provide instructions to the AI model to generate a better ticket. If this option is used, the user can only provide instructions, not provide the content. (optional)

Examples:
- `/create-ticket -p PROJECT "Implement user authentication"` - Create a ticket in PROJECT with the given title
- `/create-ticket -p PROJECT -t PR "Fix login error on mobile devices"` - Create a PR (Problem report) ticket in PROJECT
- `/create-ticket -p PROJECT -i "Performance improvement for data processing"` - Create an AI-improved ticket
- `/create-ticket -p PROJECT -r JIRA-123 "Create similar ticket for API rate limiting"` - Create a ticket in PROJECT with the same format as JIRA-123 reference ticket
- `/create-ticket -p PROJECT -t Story -r JIRA-456 -i "Enhance user profile page"` - Create an improved Story ticket in PROJECT using JIRA-456 as reference format
- `/create-ticket -p PROJECT -d "Create a ticket about improving API response times with details on how to reproduce the issue"` - Generate and create a ticket based on instructions

### Required Parameters

```
-p: Jira project key (required)
```

### Optional Parameters

```
-t: Issue type (optional, default: Task)
-i: Improve - indicate it needs AI to improve the ticket content before creating (optional)
-r: Reference ticket key to use as format template (optional)
-d: Provide instructions for AI to generate ticket content (cannot be used with explicit content)
```

### Notes

```
Use -p to specify the project key.
Ticket title should be enclosed in quotes.
The -i flag requests AI to improve the ticket content before creating it.
The -r flag allows using an existing ticket as a template for format and structure.
The -d flag provides instructions for the AI to generate content from scratch, rather than improving existing content.
Note that -d cannot be used together with explicit ticket title in quotes.
When using -r REF_TICKET, the AI will analyze the reference ticket structure and apply a similar format to the new ticket.
```

## Process

When creating a ticket, you should follow these steps:

1. Parse the command to identify the project and required parameters
2. If URL is provided, extract the project key from it
3. Validate the parameters
4. If -r REF_TICKET is provided, retrieve and analyze the reference ticket format
5. Generate the ticket description using either the standardized template or reference ticket format
6. Create the ticket in the Jira project
7. Confirm the successful creation of the ticket and provide the ticket key

## Ticket Template

All tickets created with this command will follow this standardized template when no reference ticket is provided:

h2. Issue Details

h4. Issue Description
_describe the issue_

h4. How to Reproduce
_outline steps to help the developer to reproduce the exact same issue reported here_

h4. Test/Reference data
_location of reference & test data to be used by developer to reproduce the issue_

h4. Acceptance Criteria
_outline the expected results for which the issue can be closed, such as extending the toolspec, filling spec, subscripts, parameter templates or a design is created/updated_

h4. Tests specs
_outline the test scenarios to be checked by developer during the development testing_

h4. Branch(es) to deliver to
_outline the branch(es) a fix is expected to be delivered; especially if a project fix needs to be ported to component branch_

h2. Issue Development Details

h4. Design notes
_if applicable, identify location of design notes (confluence) created/updated for this issue_

h4. Work-breakdown
_create a deliverable oriented hierarchical decomposition of the work to be executed_

h4. How to demo
_describe briefly how to demo the issue to the stakeholder(s)_

h4. Automated tests
_identify if (automated) tests needs to be updated/created for this issue_

h4. Test tooling
_identify if test tooling is available. If not, create a separate test tools issue_

h4. Issue Summary
_provide technical summary of the solution and its test results_

## Reference Ticket Usage

When the -r REF_TICKET option is used, the system will:

1. Retrieve the specified reference ticket from Jira
2. Analyze the structure and format of the reference ticket
3. Apply the same section headings, formatting style, and organizational structure to the new ticket
4. Preserve the content generation logic while using the reference format

This is particularly useful when:
- Working within projects that have specific ticket formatting requirements
- Maintaining consistency with existing tickets in a particular epic or feature area
- Following team-specific conventions that differ from the standard template

Example workflow:
1. Team has an existing well-formatted ticket JIRA-789 with custom sections
2. User wants to create a similar ticket with the same structure
3. User runs: `/create-ticket -p PROJECT -r JIRA-789 "New feature for user management"`
4. System retrieves JIRA-789 and analyzes its format
5. System creates a new ticket with the same section structure but new content based on the title

## Ticket Creation Guidelines

### Issue Description Guidelines
- Clearly describe the problem or feature request
- Include technical context and background information
- Mention any business impact if applicable

### How to Reproduce Guidelines
- Provide step-by-step instructions
- Include specific conditions or environment details
- Mention any prerequisites needed to reproduce the issue

### Test/Reference Data Guidelines
- Specify exact locations of test data
- Include version information if relevant
- Mention any special access requirements

### Acceptance Criteria Guidelines
- Define measurable outcomes
- List specific deliverables
- Include quality expectations

### Test Scenarios Guidelines
- List all relevant test cases
- Include edge cases and error conditions
- Mention integration testing requirements

## Content Format Requirements

All tickets must follow these requirements:
1. Use clear and professional language
2. Include specific details relevant to the issue
3. Follow the standardized template structure or reference ticket format
4. Provide sufficient context for developers to understand and work on the issue

Example of a good ticket description:
```
h4. Issue Description
Users are experiencing slow loading times when accessing the dashboard with large datasets. This impacts productivity and user experience.

h4. How to Reproduce
1. Log into the application
2. Navigate to the dashboard
3. Select a date range with more than 10,000 records
4. Observe the loading time exceeds 10 seconds

h4. Test/Reference data
Test data available in /data/performance/large_dataset.csv on the test server

h4. Acceptance Criteria
- Dashboard loading time should not exceed 3 seconds for 10,000+ records
- Implementation should not impact data accuracy
- Solution should be compatible with existing browsers
```

When using the -d "instructions" option, the AI will generate ticket content based on your instructions. For example:
```
/create-ticket -p PROJECT -d "Create a ticket about improving API response times. Include details about reproducing slow responses with 1000+ concurrent requests, reference data in /test/api/load_tests, and acceptance criteria for sub-200ms response times"
```

Would generate a ticket with appropriate content following the template structure.

When using the -r REF_TICKET option, the AI will follow a similar process but apply the format of the reference ticket. For example:
```
/create-ticket -p PROJECT -r JIRA-123 "Implement rate limiting for user API calls"
```

Would generate a ticket with content based on the title but using the same format and structure as JIRA-123.