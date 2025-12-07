# GoThemeMe

[![Go Reference](https://pkg.go.dev/badge/github.com/tj-smith47/gothememe.svg)](https://pkg.go.dev/github.com/tj-smith47/gothememe)
[![CI](https://github.com/tj-smith47/gothememe/actions/workflows/ci.yml/badge.svg)](https://github.com/tj-smith47/gothememe/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/tj-smith47/gothememe)](https://goreportcard.com/report/github.com/tj-smith47/gothememe)
[![Coverage](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/tj-smith47/gothememe/badges/coverage.json)](https://github.com/tj-smith47/gothememe/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A comprehensive Go theming package for web applications. The web counterpart to [bubbletint](https://github.com/lrstanley/bubbletint).

## Features

- **128+ Built-in Themes** - Dracula, Nord, Gruvbox, Tokyo Night, Catppuccin, and more
- **Multiple Output Formats** - CSS variables, SCSS, JSON, DTCG design tokens
- **Syntax Highlighting** - Compatible with Prism.js, Highlight.js, and Chroma
- **WCAG Accessibility** - Built-in contrast ratio validation
- **Framework Agnostic** - Works with any web framework (HTMX, React, Vue, Svelte, etc.)
- **Custom Themes** - Fluent builder API for creating custom themes
- **Type Safe** - Full Go type safety with comprehensive documentation

## Installation

```bash
go get github.com/tj-smith47/gothememe
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/tj-smith47/gothememe"
)

func main() {
    // Initialize the default registry with all themes
    gothememe.NewDefaultRegistry()

    // Generate CSS for the current theme (Dracula by default)
    css := gothememe.CSS(gothememe.DefaultCSSOptions())
    fmt.Println(css)
}
```

## Usage Patterns

### Pattern 1: Global Registry (Simplest)

For applications with a single, application-wide theme:

```go
gothememe.NewDefaultRegistry()
gothememe.SetThemeID("dracula")

// Generate CSS
css := gothememe.CSS(gothememe.DefaultCSSOptions())

// Access theme colors directly
bg := gothememe.Background()
text := gothememe.TextPrimary()
accent := gothememe.Accent()
```

### Pattern 2: Custom Registry (Dynamic Switching)

For applications that need runtime theme switching:

```go
registry := gothememe.NewRegistry(
    gothememe.ThemeDracula,
    gothememe.ThemeNord,
)

// Navigate through themes
registry.NextTheme()
registry.PreviousTheme()
registry.SetThemeID("nord")

// Generate CSS for current theme
css := gothememe.GenerateCSS(registry.GetCurrentTheme(), gothememe.DefaultCSSOptions())
```

### Pattern 3: Direct Theme Usage (Static)

For compile-time theme selection without a registry:

```go
css := gothememe.GenerateCSS(gothememe.ThemeDracula, gothememe.DefaultCSSOptions())
```

## Output Formats

### CSS Variables

```go
opts := gothememe.CSSOptions{
    Prefix:      "theme",
    IncludeRoot: true,
}
css := gothememe.GenerateCSS(theme, opts)
```

Output:
```css
:root {
    --theme-background: #282a36;
    --theme-text-primary: #f8f8f2;
    --theme-accent: #bd93f9;
    /* ... */
}
```

### Multi-Theme CSS

Generate CSS for all themes using `data-theme` attribute:

```go
css := gothememe.AllThemesCSS(gothememe.DefaultCSSOptions())
```

Output:
```css
[data-theme="dracula"] { --theme-background: #282a36; ... }
[data-theme="nord"] { --theme-background: #2e3440; ... }
```

Switch themes in HTML:
```html
<html data-theme="dracula">
```

### Design Tokens (DTCG)

Generate DTCG v1 compliant design tokens:

```go
tokens, _ := gothememe.GenerateDesignTokens(theme, gothememe.DefaultTokenOptions())
```

### Syntax Highlighting

Generate CSS for code syntax highlighting:

```go
// Prism.js compatible
syntaxCSS := gothememe.GenerateSyntaxCSS(theme, gothememe.SyntaxOptions{
    Format:       gothememe.SyntaxPrism,
    UseVariables: true,
})

// Highlight.js compatible
syntaxCSS := gothememe.GenerateSyntaxCSS(theme, gothememe.SyntaxOptions{
    Format: gothememe.SyntaxHighlightJS,
})
```

## Custom Themes

### Using the Builder

```go
theme := gothememe.NewThemeBuilder("my-theme", "My Custom Theme").
    WithBackground(gothememe.Hex("#1a1a2e")).
    WithTextPrimary(gothememe.Hex("#e4e4e4")).
    WithAccent(gothememe.Hex("#e94560")).
    WithGreen(gothememe.Hex("#22c55e")).
    WithRed(gothememe.Hex("#ef4444")).
    Build()
```

### From Minimal Palette

Generate a complete theme from just 3 colors:

```go
theme := gothememe.GenerateThemeFromPalette("minimal", "Minimal Theme", gothememe.Palette{
    Background: gothememe.Hex("#1a1a2e"),
    Foreground: gothememe.Hex("#e4e4e4"),
    Accent:     gothememe.Hex("#e94560"),
})
```

### Deriving from Existing Theme

```go
theme := gothememe.DeriveTheme(gothememe.ThemeDracula, "dracula-custom", "Dracula Custom", map[string]gothememe.Color{
    "accent": gothememe.Hex("#ff0000"),
})
```

## Color Manipulation

```go
color := gothememe.Hex("#e94560")

// Modify colors
lighter := color.Lighten(0.1)
darker := color.Darken(0.1)
transparent := color.WithAlpha(0.5)
opposite := color.Complement()
mixed := color.Mix(gothememe.Hex("#ffffff"), 0.5)

// Convert formats
hex := color.Hex()           // "#e94560"
rgb := color.CSSRGB()        // "rgb(233, 69, 96)"
hsl := color.CSSHSL()        // "hsl(350, 79%, 59%)"
```

## Accessibility

Check WCAG contrast requirements:

```go
import "github.com/tj-smith47/gothememe/pkg/contrast"

// Calculate contrast ratio
ratio := contrast.RatioHex("#ffffff", "#000000") // 21.0

// Check compliance levels
meetsAA := contrast.MeetsAAHex("#f8f8f2", "#282a36", false)   // true
meetsAAA := contrast.MeetsAAAHex("#f8f8f2", "#282a36", false) // true

// Get compliance level
level := contrast.CheckHex("#f8f8f2", "#282a36") // contrast.LevelAAA
```

## Available Themes

See [DEFAULT_THEMES.md](DEFAULT_THEMES.md) for a complete list of themes with previews.

Popular themes include:
- Dracula, Dracula Plus
- Nord
- Gruvbox (Dark/Light)
- Tokyo Night (variants)
- Catppuccin (Frappe, Latte, Macchiato, Mocha)
- One Dark
- Solarized (Dark/Light)
- And 100+ more...

## Framework Integration

GoThemeMe works with any web framework. See [docs/INTEGRATION.md](docs/INTEGRATION.md) for examples with:
- Vanilla HTML/CSS
- HTMX + templ
- React
- Vue
- Svelte
- Tailwind CSS

## Contributing

Contributions are welcome! See [docs/CONTRIBUTING.md](docs/CONTRIBUTING.md) for guidelines.

## License

MIT License - see [LICENSE](LICENSE) for details.

## Credits

- Inspired by [bubbletint](https://github.com/lrstanley/bubbletint) by Liam Stanley
- Theme colors sourced from [iTerm2-Color-Schemes](https://github.com/mbadolato/iTerm2-Color-Schemes)
- Built with [go-colorful](https://github.com/lucasb-eyer/go-colorful)
