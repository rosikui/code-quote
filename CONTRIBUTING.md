# Contributing to code-quote

Thank you for your interest in contributing to code-quote! This document provides guidelines and information for contributors.

## Table of Contents

- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Adding Quotes](#adding-quotes)
- [Code Contributions](#code-contributions)
- [Testing](#testing)
- [Pull Request Process](#pull-request-process)
- [Code Style](#code-style)
- [Reporting Issues](#reporting-issues)

## Getting Started

1. Fork the repository
2. Clone your fork locally
3. Create a feature branch
4. Make your changes
5. Test your changes
6. Submit a pull request

## Development Setup

### Prerequisites

- Go 1.24 or later
- Git

### Local Development

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/code-quote.git
cd code-quote

# Install dependencies
make deps

# Build the project
make build

# Run tests
make test

# Run the application
make run
```

### Project Structure

```
code-quote/
├── main.go                 # Main entry point with embed
├── cmd/
│   └── code-quote/
│       └── main.go         # CLI logic with Cobra
├── internal/
│   ├── quote/
│   │   ├── loader.go       # JSON/YAML file loading
│   │   └── picker.go       # Quote selection and filtering
│   └── term/
│       └── render.go       # Terminal output formatting
├── quotes/
│   └── quotes.en.json      # Built-in quotes
├── tests/                  # Test files
└── docs/                   # Documentation
```

## Adding Quotes

### Built-in Quotes

To add quotes to the built-in collection, edit `quotes/quotes.en.json`:

```json
[
  {
    "text": "Your new quote here",
    "author": "Author Name",
    "tags": ["tag1", "tag2", "tag3"],
    "lang": "en"
  }
]
```

### Quote Guidelines

- **Text**: Should be inspiring, educational, or humorous programming-related content
- **Author**: Use the actual author's name, or "Anonymous" if unknown
- **Tags**: Use relevant tags like `programming`, `clean-code`, `debugging`, `humor`, etc.
- **Language**: Use ISO language codes (e.g., `en`, `zh`, `ja`)

### Custom Quote Files

You can also contribute by creating custom quote files:

```bash
# Create a new quote file
cat > my-quotes.json << EOF
[
  {
    "text": "Your quote here",
    "author": "Author Name",
    "tags": ["motivation", "programming"],
    "lang": "en"
  }
]
EOF

# Test your quotes
./code-quote --file my-quotes.json
```

## Code Contributions

### Adding New Features

1. Create a feature branch: `git checkout -b feature/your-feature-name`
2. Implement your feature
3. Add tests for your feature
4. Update documentation if needed
5. Ensure all tests pass: `make test`
6. Commit your changes with a descriptive message

### Bug Fixes

1. Create a bug fix branch: `git checkout -b fix/issue-description`
2. Fix the bug
3. Add tests to prevent regression
4. Test your fix thoroughly
5. Commit with a clear description of the fix

### Code Style

- Follow Go conventions and idioms
- Use meaningful variable and function names
- Add comments for complex logic
- Keep functions small and focused
- Use proper error handling

### Commit Messages

Use conventional commit format:

```
type(scope): description

[optional body]

[optional footer]
```

Examples:
- `feat(quote): add support for YAML quote files`
- `fix(picker): resolve daily quote seed calculation`
- `docs(readme): add installation instructions`

## Testing

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
go test -cover ./...

# Run specific test
go test ./internal/quote -v
```

### Writing Tests

- Test files should be named `*_test.go`
- Use descriptive test names
- Test both success and error cases
- Mock external dependencies when appropriate

Example test:

```go
func TestQuotePicker_PickRandom(t *testing.T) {
    quotes := []Quote{
        {Text: "Test quote 1", Author: "Author 1"},
        {Text: "Test quote 2", Author: "Author 2"},
    }
    
    picker := NewPicker(quotes)
    quote, err := picker.PickRandom()
    
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    
    if quote == nil {
        t.Error("Expected a quote, got nil")
    }
}
```

## Pull Request Process

1. **Fork and Clone**: Fork the repository and clone your fork
2. **Create Branch**: Create a feature branch from `main`
3. **Make Changes**: Implement your changes
4. **Test**: Ensure all tests pass and the application works correctly
5. **Document**: Update documentation if needed
6. **Commit**: Use conventional commit messages
7. **Push**: Push your branch to your fork
8. **Submit PR**: Create a pull request with a clear description

### PR Guidelines

- Provide a clear description of the changes
- Include any relevant issue numbers
- Add screenshots for UI changes
- Ensure the PR title follows conventional commit format
- Request reviews from maintainers

### PR Template

```markdown
## Description
Brief description of the changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Documentation update
- [ ] Quote addition

## Testing
- [ ] All tests pass
- [ ] Manual testing completed
- [ ] No breaking changes

## Checklist
- [ ] Code follows style guidelines
- [ ] Self-review completed
- [ ] Documentation updated
- [ ] Tests added/updated
```

## Code Style

### Go Code Style

- Use `gofmt` for formatting
- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use meaningful package names
- Keep packages focused and cohesive

### Error Handling

```go
// Good
if err != nil {
    return fmt.Errorf("failed to load quotes: %w", err)
}

// Avoid
if err != nil {
    return err
}
```

### Documentation

- Add comments for exported functions and types
- Use complete sentences in comments
- Include examples for complex functions

## Reporting Issues

### Bug Reports

When reporting bugs, please include:

1. **Environment**: OS, Go version, code-quote version
2. **Steps to Reproduce**: Clear, step-by-step instructions
3. **Expected Behavior**: What you expected to happen
4. **Actual Behavior**: What actually happened
5. **Additional Context**: Any relevant information

### Feature Requests

For feature requests, please include:

1. **Description**: Clear description of the feature
2. **Use Case**: Why this feature would be useful
3. **Proposed Implementation**: Any ideas for implementation
4. **Alternatives**: Any alternative solutions considered

## Getting Help

- **Issues**: Use GitHub issues for bugs and feature requests
- **Discussions**: Use GitHub discussions for questions and ideas
- **Code Review**: Request reviews from maintainers

## License

By contributing to code-quote, you agree that your contributions will be licensed under the MIT License.

## Recognition

Contributors will be recognized in:
- The project README
- Release notes
- GitHub contributors page

Thank you for contributing to code-quote! 🎉 