// Package gothememe provides a comprehensive theming system for web applications.
//
// GoThemeMe is the web counterpart to [bubbletint], offering 450+ pre-built themes
// with full CSS variable support, syntax highlighting, and WCAG accessibility
// validation. It generates CSS custom properties, SCSS variables, JSON, and
// DTCG-compliant design tokens.
//
// # Quick Start
//
// Initialize the default registry and generate CSS:
//
//	func main() {
//	    gothememe.NewDefaultRegistry()
//	    css := gothememe.CSS(gothememe.DefaultCSSOptions())
//	    // Use the CSS in your templates or serve it directly
//	}
//
// # Three Usage Patterns
//
// Pattern 1: Global Registry (Simplest)
//
// For applications with a single, application-wide theme:
//
//	gothememe.NewDefaultRegistry()
//	gothememe.SetThemeID("dracula")
//	css := gothememe.CSS(gothememe.DefaultCSSOptions())
//
// Pattern 2: Custom Registry (Dynamic Switching)
//
// For applications that need runtime theme switching:
//
//	registry := gothememe.NewRegistry(
//	    gothememe.ThemeDracula,
//	    gothememe.ThemeNord,
//	    gothememe.ThemeGruvboxDark,
//	)
//	registry.NextTheme()     // Switch to next theme
//	registry.SetThemeID("nord")
//	css := gothememe.GenerateCSS(registry.GetCurrentTheme(), gothememe.DefaultCSSOptions())
//
// Pattern 3: Direct Theme Usage (Static)
//
// For compile-time theme selection without a registry:
//
//	css := gothememe.GenerateCSS(gothememe.ThemeDracula, gothememe.DefaultCSSOptions())
//
// # Custom Themes
//
// Create custom themes using the builder pattern:
//
//	theme := gothememe.NewThemeBuilder("my-theme", "My Custom Theme").
//	    WithBackground(gothememe.Hex("#1a1a2e")).
//	    WithTextPrimary(gothememe.Hex("#e4e4e4")).
//	    WithAccent(gothememe.Hex("#e94560")).
//	    WithGreen(gothememe.Hex("#22c55e")).
//	    WithRed(gothememe.Hex("#ef4444")).
//	    Build()
//
// Or generate a complete theme from a minimal palette:
//
//	theme := gothememe.GenerateThemeFromPalette("minimal", "Minimal Theme", gothememe.Palette{
//	    Background: gothememe.Hex("#1a1a2e"),
//	    Foreground: gothememe.Hex("#e4e4e4"),
//	    Accent:     gothememe.Hex("#e94560"),
//	})
//
// # Output Formats
//
// Generate output in multiple formats for different use cases:
//
//	// CSS custom properties
//	css := gothememe.GenerateCSS(theme, gothememe.DefaultCSSOptions())
//
//	// SCSS variables
//	scss := gothememe.GenerateSCSS(theme, gothememe.CSSOptions{Prefix: "app"})
//
//	// JSON for JavaScript consumption
//	json := gothememe.GenerateJSON(theme, gothememe.CSSOptions{})
//
//	// DTCG design tokens for Figma/Style Dictionary
//	tokens, _ := gothememe.GenerateDesignTokens(theme, gothememe.DefaultTokenOptions())
//
// # Syntax Highlighting
//
// Generate syntax highlighting CSS compatible with popular libraries:
//
//	// Prism.js compatible
//	syntaxCSS := gothememe.GenerateSyntaxCSS(theme, gothememe.SyntaxOptions{
//	    Format:       gothememe.SyntaxPrism,
//	    UseVariables: true,
//	})
//
//	// Highlight.js compatible
//	syntaxCSS := gothememe.GenerateSyntaxCSS(theme, gothememe.SyntaxOptions{
//	    Format: gothememe.SyntaxHighlightJS,
//	})
//
// # Multi-Theme CSS
//
// Generate CSS for all themes using data-theme attribute selectors:
//
//	css := gothememe.AllThemesCSS(gothememe.DefaultCSSOptions())
//
// This outputs:
//
//	[data-theme="dracula"] { --theme-background: #282a36; ... }
//	[data-theme="nord"] { --theme-background: #2e3440; ... }
//
// Then switch themes in HTML with:
//
//	<html data-theme="dracula">
//
// # Color Manipulation
//
// Colors support various manipulations:
//
//	color := gothememe.Hex("#e94560")
//	lighter := color.Lighten(0.1)
//	darker := color.Darken(0.1)
//	transparent := color.WithAlpha(0.5)
//	opposite := color.Complement()
//
// # Accessibility
//
// Check WCAG contrast requirements:
//
//	import "github.com/tj-smith47/gothememe/pkg/contrast"
//
//	ratio := contrast.Ratio(foreground, background)
//	meetsAA := contrast.MeetsAA(foreground, background, false)
//	meetsAAA := contrast.MeetsAAA(foreground, background, false)
//
// [bubbletint]: https://github.com/lrstanley/bubbletint
package gothememe
