---
name: Bug report
about: Create a report to help us improve
title: '[BUG] '
labels: ['bug', 'needs-triage']
assignees: ''
---

## Bug Description

A clear and concise description of what the bug is.

## Steps to Reproduce

1. Run command: `./code-quote --flag value`
2. See error: `error message here`
3. Expected behavior: `what should happen`

## Environment

- **OS**: [e.g., macOS 14.0, Ubuntu 22.04, Windows 11]
- **Go Version**: [e.g., go version go1.24.3 darwin/arm64]
- **code-quote Version**: [e.g., v1.0.0 or commit hash]
- **Terminal**: [e.g., iTerm2, VS Code Terminal, Windows Terminal]

## Expected Behavior

A clear and concise description of what you expected to happen.

## Actual Behavior

A clear and concise description of what actually happened.

## Error Messages

```
Paste any error messages or logs here
```

## Additional Context

Add any other context about the problem here, such as:
- Screenshots if applicable
- Custom quote files being used
- Specific flags or combinations that trigger the issue

## Reproduction Files

If the bug involves custom quote files, please attach them:

```json
[
  {
    "text": "Example quote",
    "author": "Example Author",
    "tags": ["example"],
    "lang": "en"
  }
]
```

## Checklist

- [ ] I have searched existing issues to avoid duplicates
- [ ] I have provided all required information
- [ ] I can reproduce this issue consistently
- [ ] This is a bug in code-quote, not a configuration issue 