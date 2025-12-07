# Echo Framework Example

Full-featured theme API using the [Echo](https://echo.labstack.com/) web framework.

## Features

- REST API for theme management
- Multiple output formats (CSS, JSON, Design Tokens)
- Theme analysis and validation endpoints
- Interactive theme explorer UI

## Running

```bash
go run main.go
```

Then open http://localhost:8080

## API Endpoints

### Theme List
```bash
GET /api/themes
# Returns: [{"id":"dracula","displayName":"Dracula","isDark":true,"author":"..."}]
```

### Theme CSS
```bash
GET /api/themes/:id/css
# Returns: CSS with :root selector
```

### Theme JSON
```bash
GET /api/themes/:id/json
# Returns: Color values as JSON
```

### Design Tokens (DTCG Format)
```bash
GET /api/themes/:id/tokens
# Returns: Design Tokens Community Group format
```

### Theme Statistics
```bash
GET /api/themes/:id/stats
# Returns: {colorCount, uniqueColors, contrastScore, accessibilityPercent}
```

### Theme Validation
```bash
GET /api/themes/:id/validate
# Returns: {valid, errors, contrastIssues}
```

## Key Code

```go
// Theme validation
validationErrors := gothememe.ValidateTheme(t)
contrastIssues := gothememe.ValidateContrast(t, gothememe.ContrastLevelAA)

// Design tokens
tokens, _ := gothememe.GenerateDesignTokens(t, gothememe.DefaultTokenOptions())

// Theme analysis
stats := gothememe.AnalyzeTheme(t)
```

## Dependencies

```bash
go get github.com/labstack/echo/v4
```

## Files

- `main.go` - Echo server with API routes and UI
- `go.mod` - Module definition with Echo dependency
