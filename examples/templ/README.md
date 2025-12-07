# Templ Example

Integration with the [templ](https://templ.guide/) template engine.

Note: This example uses `html/template` as a stand-in for templ syntax. In a real templ project, you would use `.templ` files.

## Features

- Server-side rendered theme selector
- All 451+ themes available
- Semantic color examples (success, warning, error states)
- Code syntax highlighting preview

## Running

```bash
go run main.go
```

Then open http://localhost:8080

## How It Works

1. Server loads all themes from `themes.All()`
2. Generates CSS with `data-theme` selectors for a subset (20 themes)
3. Renders dropdown with theme options
4. JavaScript handles client-side theme switching

## Key Code

```go
// Load all themes
allThemes := themes.All()

// Generate CSS for subset (performance)
css := gothememe.GenerateAllThemesCSS(subset, gothememe.CSSOptions{
    UseDataAttribute: true,
})
```

## Template Integration

In a real templ project:

```templ
// components/theme_selector.templ
templ ThemeSelector(themes []gothememe.Theme, current string) {
    <select id="theme" onchange="setTheme(this.value)">
        for _, t := range themes {
            <option value={ t.ID() } selected?={ t.ID() == current }>
                { t.DisplayName() }
            </option>
        }
    </select>
}
```

## Files

- `main.go` - Server with embedded HTML template
