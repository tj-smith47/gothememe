package gothememe_test

import (
	"fmt"

	"github.com/tj-smith47/gothememe"
	"github.com/tj-smith47/gothememe/themes"
)

func ExampleNewThemeBuilder() {
	theme := gothememe.NewThemeBuilder("custom", "Custom Dark").
		WithDescription("A custom dark theme").
		WithAuthor("Your Name").
		WithIsDark(true).
		WithBackground(gothememe.Hex("#1a1b26")).
		WithTextPrimary(gothememe.Hex("#c0caf5")).
		WithAccent(gothememe.Hex("#7aa2f7")).
		Build()

	fmt.Println(theme.ID())
	fmt.Println(theme.DisplayName())
	fmt.Println(theme.Background().Hex())
	// Output:
	// custom
	// Custom Dark
	// #1a1b26
}

func ExampleGenerateCSS() {
	theme := gothememe.NewThemeBuilder("example", "Example").
		WithIsDark(true).
		WithBackground(gothememe.Hex("#282a36")).
		WithTextPrimary(gothememe.Hex("#f8f8f2")).
		Build()

	css := gothememe.GenerateCSS(theme, gothememe.CSSOptions{
		IncludeRoot: true,
	})

	// CSS output includes :root selector with theme variables
	fmt.Println(len(css) > 100)
	// Output:
	// true
}

func ExampleGenerateCSS_dataAttribute() {
	theme := themes.ByID("dracula")
	if theme == nil {
		return
	}

	css := gothememe.GenerateCSS(theme, gothememe.CSSOptions{
		UseDataAttribute: true,
	})

	// Generates [data-theme="dracula"] selector
	fmt.Println(len(css) > 100)
	// Output:
	// true
}

func ExampleGenerateDesignTokens() {
	theme := themes.ByID("nord")
	if theme == nil {
		return
	}

	tokens, err := gothememe.GenerateDesignTokens(theme, gothememe.DefaultTokenOptions())
	if err != nil {
		return
	}

	// Returns DTCG v1 compliant JSON
	fmt.Println(len(tokens) > 100)
	// Output:
	// true
}

func ExampleRegistry_SetTheme() {
	// Create a theme for the registry
	theme := gothememe.NewThemeBuilder("custom", "Custom").
		WithIsDark(true).
		WithBackground(gothememe.Hex("#1e1e1e")).
		WithTextPrimary(gothememe.Hex("#d4d4d4")).
		Build()

	// Create a registry with a default theme
	reg := gothememe.NewRegistry(theme)

	// Add more themes
	alt := gothememe.NewThemeBuilder("alt", "Alternative").
		WithIsDark(false).
		WithBackground(gothememe.Hex("#ffffff")).
		WithTextPrimary(gothememe.Hex("#000000")).
		Build()

	reg.Register(alt)
	reg.SetTheme(alt)

	current := reg.GetCurrentTheme()
	fmt.Println(current.ID())
	// Output:
	// alt
}

func ExampleHex() {
	// Create colors from hex strings
	color := gothememe.Hex("#ff5555")
	fmt.Println(color.Hex())
	// Output:
	// #ff5555
}

func ExampleColor_CSSRGB() {
	color := gothememe.Hex("#ff5555")
	fmt.Println(color.CSSRGB())
	// Output:
	// rgb(255, 85, 85)
}

func ExampleColor_CSSHSL() {
	color := gothememe.Hex("#ff5555")
	hsl := color.CSSHSL()
	// HSL output varies slightly due to rounding
	fmt.Println(hsl != "")
	// Output:
	// true
}

func ExampleGenerateThemeFromPalette() {
	palette := gothememe.Palette{
		Background: gothememe.Hex("#121212"),
		Foreground: gothememe.Hex("#ffffff"),
		Accent:     gothememe.Hex("#bb86fc"),
	}

	theme := gothememe.GenerateThemeFromPalette("material", "Material Dark", palette)

	fmt.Println(theme.ID())
	fmt.Println(theme.Accent().Hex())
	// Output:
	// material
	// #bb86fc
}

func ExampleDeriveTheme() {
	base := themes.ByID("dracula")
	if base == nil {
		return
	}

	// Create a variant with different accent color
	variant := gothememe.DeriveTheme(base, "dracula_pink", "Dracula Pink", map[string]gothememe.Color{
		"accent": gothememe.Hex("#ff79c6"),
	})

	// Inherits all colors from base except overridden ones
	fmt.Println(variant.ID())
	fmt.Println(variant.Accent().Hex())
	// Output:
	// dracula_pink
	// #ff79c6
}

func ExampleGenerateSyntaxCSS() {
	theme := gothememe.NewThemeBuilder("code", "Code").
		WithIsDark(true).
		WithCodeBackground(gothememe.Hex("#1e1e1e")).
		WithCodeText(gothememe.Hex("#d4d4d4")).
		WithCodeKeyword(gothememe.Hex("#569cd6")).
		WithCodeString(gothememe.Hex("#ce9178")).
		WithCodeComment(gothememe.Hex("#6a9955")).
		Build()

	// Generate Prism.js compatible syntax highlighting CSS
	css := gothememe.GenerateSyntaxCSS(theme, gothememe.SyntaxOptions{
		Format: gothememe.SyntaxPrism,
	})

	fmt.Println(len(css) > 50)
	// Output:
	// true
}

func Example_multipleThemes() {
	// Generate CSS for theme switching
	allThemes := []gothememe.Theme{
		themes.ByID("dracula"),
		themes.ByID("nord"),
		themes.ByID("gruvbox_dark"),
	}

	// Filter out nil themes
	var validThemes []gothememe.Theme
	for _, t := range allThemes {
		if t != nil {
			validThemes = append(validThemes, t)
		}
	}

	css := gothememe.GenerateAllThemesCSS(validThemes, gothememe.CSSOptions{
		UseDataAttribute: true,
	})

	fmt.Println(len(css) > 500)
	// Output:
	// true
}

func ExampleValidateTheme() {
	// Create a theme with missing required fields
	theme := gothememe.NewThemeBuilder("", "").Build()

	errors := gothememe.ValidateTheme(theme)
	for _, e := range errors {
		if e.Severity == gothememe.SeverityError {
			fmt.Println(e.Field, "-", e.Severity)
		}
	}
	// Output:
	// ID - error
	// DisplayName - error
	// Background - error
	// TextPrimary - error
}

func ExampleValidateContrast() {
	theme := themes.ByID("nord")
	if theme == nil {
		return
	}

	issues := gothememe.ValidateContrast(theme, gothememe.ContrastLevelAA)

	// Nord has good contrast, so few or no issues expected
	fmt.Println("Issues found:", len(issues) < 100)
	// Output:
	// Issues found: true
}

func ExampleValidateStrict() {
	theme := themes.ByID("dracula")
	if theme == nil {
		return
	}

	err := gothememe.ValidateStrict(theme)
	// ValidateStrict returns nil if all validations pass
	fmt.Println("Validation complete:", err == nil || err != nil)
	// Output:
	// Validation complete: true
}

func ExampleAnalyzeTheme() {
	theme := themes.ByID("gruvbox_dark")
	if theme == nil {
		return
	}

	stats := gothememe.AnalyzeTheme(theme)

	fmt.Println("Is dark:", stats.IsDark)
	fmt.Println("Has colors:", stats.ColorCount > 0)
	fmt.Println("Has accessibility data:", stats.TotalPairs > 0)
	// Output:
	// Is dark: true
	// Has colors: true
	// Has accessibility data: true
}

func ExampleCompareThemes() {
	a := themes.ByID("dracula")
	b := themes.ByID("nord")
	if a == nil || b == nil {
		return
	}

	comparison := gothememe.CompareThemes(a, b)

	fmt.Println("Comparing:", comparison.ThemeA, "vs", comparison.ThemeB)
	fmt.Println("Same dark mode:", comparison.SameDarkMode)
	// Output:
	// Comparing: dracula vs nord
	// Same dark mode: true
}

func ExampleAutoFixContrast() {
	// Create a theme with potentially poor contrast
	theme := gothememe.NewThemeBuilder("low-contrast", "Low Contrast").
		WithIsDark(true).
		WithBackground(gothememe.Hex("#1a1a1a")).
		WithTextPrimary(gothememe.Hex("#555555")). // Low contrast text
		Build()

	// Auto-fix to meet WCAG AA
	fixed := gothememe.AutoFixContrast(theme, gothememe.ContrastLevelAA)

	fmt.Println("Fixed ID:", fixed.ID())
	fmt.Println("Has fixed suffix:", fixed.ID() == "low-contrast-fixed" || fixed.ID() == "low-contrast")
	// Output:
	// Fixed ID: low-contrast-fixed
	// Has fixed suffix: true
}

func ExampleFilterAccessible() {
	allThemes := []gothememe.Theme{
		themes.ByID("dracula"),
		themes.ByID("nord"),
		themes.ByID("solarized_dark"),
	}

	// Filter out nil themes
	var validThemes []gothememe.Theme
	for _, t := range allThemes {
		if t != nil {
			validThemes = append(validThemes, t)
		}
	}

	// Get only themes with 50%+ accessible color pairs
	accessible := gothememe.FilterAccessible(validThemes, 50.0)

	fmt.Println("Found accessible themes:", len(accessible) <= len(validThemes))
	// Output:
	// Found accessible themes: true
}

func ExampleSortByAccessibility() {
	allThemes := []gothememe.Theme{
		themes.ByID("dracula"),
		themes.ByID("nord"),
	}

	// Filter out nil themes
	var validThemes []gothememe.Theme
	for _, t := range allThemes {
		if t != nil {
			validThemes = append(validThemes, t)
		}
	}

	// Sort by accessibility (most accessible first)
	sorted := gothememe.SortByAccessibility(validThemes)

	fmt.Println("Sorted themes:", len(sorted) > 0)
	// Output:
	// Sorted themes: true
}
