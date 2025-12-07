# Vanilla HTML/CSS Example

A minimal example demonstrating gothememe with plain HTML and JavaScript.

## Features

- Dynamic theme switching using `data-theme` attribute
- CSS custom properties for all theme colors
- Theme persistence via localStorage
- No framework dependencies

## Running

```bash
go run main.go
```

Then open http://localhost:8080

## How It Works

1. The Go server generates CSS for selected themes with `data-theme` selectors
2. Static HTML uses CSS custom properties like `var(--theme-background)`
3. JavaScript toggles themes by setting `data-theme` on the `<html>` element
4. Theme preference is saved to localStorage

## Key Code

```go
css := gothememe.GenerateAllThemesCSS(themes, gothememe.CSSOptions{
    UseDataAttribute: true,  // Generates [data-theme="dracula"] selectors
})
```

```javascript
document.documentElement.setAttribute('data-theme', 'dracula');
```

## Files

- `main.go` - Server that generates and serves theme CSS
- `static/index.html` - HTML page with theme switcher
