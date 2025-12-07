# HTMX Example

Server-side rendering with HTMX for dynamic theme switching without full page reloads.

## Features

- HTMX-powered theme preview updates
- Server-side theme analysis with `AnalyzeTheme()`
- Dynamic content swapping
- Theme persistence via localStorage

## Running

```bash
go run main.go
```

Then open http://localhost:8080

## How It Works

1. Theme buttons use HTMX to fetch preview content from the server
2. Server analyzes the selected theme and returns rendered HTML
3. HTMX swaps the preview panel content without page reload
4. JavaScript updates the `data-theme` attribute for immediate styling

## Key Code

```html
<button hx-get="/preview?theme={{.ID}}"
        hx-target="#preview"
        hx-swap="innerHTML"
        onclick="setTheme('{{.ID}}')">
    {{.DisplayName}}
</button>
```

```go
stats := gothememe.AnalyzeTheme(t)
// Returns: ColorCount, UniqueColors, ContrastScore, AccessibilityPercent
```

## API Endpoints

- `GET /` - Main page with theme selector
- `GET /themes.css` - Generated CSS for all themes
- `GET /theme-options?current=id` - Theme button partial
- `GET /preview?theme=id` - Theme preview partial with analysis

## Files

- `main.go` - Server with HTMX endpoints and embedded templates
