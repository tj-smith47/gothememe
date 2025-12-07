# Contributing to GoThemeMe

Thank you for your interest in contributing to GoThemeMe! This document provides guidelines and instructions for contributing.

## Development Setup

### Prerequisites

- Go 1.24 or later
- Git
- Make (optional, but recommended)

### Getting Started

1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/YOUR_USERNAME/gothememe.git
   cd gothememe
   ```

3. Install dependencies:
   ```bash
   go mod download
   ```

4. Install development tools:
   ```bash
   make tools
   ```

5. Run tests to verify setup:
   ```bash
   make test
   ```

## Code Style

GoThemeMe follows standard Go conventions enforced by golangci-lint. Before submitting a PR:

```bash
make lint
make fmt
```

### Key Guidelines

- **Documentation**: All exported types, functions, and methods must have godoc comments
- **Testing**: New features must include tests; bug fixes should include regression tests
- **Naming**: Follow Go naming conventions (exported = PascalCase, unexported = camelCase)
- **Error handling**: Always handle errors appropriately; don't ignore them

## Adding New Themes

### Manual Theme Addition

1. Create a new file in `defaultthemes/`:
   ```bash
   touch defaultthemes/mytheme.go
   ```

2. Implement the Theme interface (use `dracula.go` as a template):
   ```go
   package defaultthemes

   import "github.com/tj-smith47/gothememe"

   var ThemeMyTheme gothememe.Theme = &themeMyTheme{}

   type themeMyTheme struct{}

   func (t *themeMyTheme) ID() string          { return "mytheme" }
   func (t *themeMyTheme) DisplayName() string { return "My Theme" }
   // ... implement all Theme interface methods
   ```

3. Add to `defaultthemes/themes.go`:
   ```go
   func All() []gothememe.Theme {
       return []gothememe.Theme{
           ThemeDracula,
           ThemeNord,
           ThemeMyTheme, // Add your theme here
       }
   }
   ```

4. (Optional) Export from main package in `default_themes.go`:
   ```go
   var ThemeMyTheme Theme = defaultthemes.ThemeMyTheme
   ```

5. Run tests and lint:
   ```bash
   make test lint
   ```

### Using Theme Generator

For bulk theme additions from external sources:

```bash
go run ./cmd/themegen
```

The generator fetches themes from iTerm2-Color-Schemes and other sources.

## Adding New Features

1. **Open an issue first**: Discuss the feature before implementation
2. **Create a branch**: `git checkout -b feature/my-feature`
3. **Write tests**: Include unit tests for new functionality
4. **Update documentation**: Add godoc comments and update README if needed
5. **Submit a PR**: Reference the issue in your PR description

## Testing

```bash
# Run all tests
make test

# Run tests with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run specific package tests
go test -v ./pkg/contrast/...

# Run benchmarks
make bench
```

### Test Guidelines

- Use table-driven tests where appropriate
- Include edge cases and error conditions
- Tests should be deterministic (no random values without seeding)

## Pull Request Process

1. **Ensure CI passes**: All checks must be green
2. **Update CHANGELOG**: Add entry under "Unreleased"
3. **Review checklist**:
   - [ ] Tests added/updated
   - [ ] Documentation updated
   - [ ] Lint passes
   - [ ] Commit messages follow conventional commits

### Commit Messages

Follow [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: add new Catppuccin theme variants
fix: correct contrast ratio calculation for transparent colors
docs: update installation instructions
test: add benchmarks for CSS generation
chore: update dependencies
```

## Release Process

Releases are automated via GitHub Actions:

1. Commits to `main` trigger auto-tagging based on commit messages
2. Tags trigger GoReleaser to create releases
3. Version follows semantic versioning

## Questions?

- Open an issue for bugs or feature requests
- Start a discussion for general questions
- Check existing issues/discussions before creating new ones

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
