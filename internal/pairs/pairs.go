// Package pairs provides shared color pair definitions for contrast validation
// and accessibility analysis. This consolidates the standard color pairs checked
// across the gothememe package.
package pairs

// ColorPair represents a foreground/background color combination for contrast checking.
type ColorPair struct {
	FgName string // Name of the foreground color (e.g., "TextPrimary")
	BgName string // Name of the background color (e.g., "Background")
	FgHex  string // Hex value of the foreground color
	BgHex  string // Hex value of the background color
}

// StandardPairSpec defines a color pair by getter function names.
// This allows the pair definitions to be applied to any Theme interface.
type StandardPairSpec struct {
	FgName string
	BgName string
}

// StandardPairSpecs returns the list of standard color pairs to check for WCAG compliance.
// These pairs represent the most important foreground/background combinations
// that affect readability in a theme.
func StandardPairSpecs() []StandardPairSpec {
	return []StandardPairSpec{
		// Text on backgrounds
		{FgName: "TextPrimary", BgName: "Background"},
		{FgName: "TextSecondary", BgName: "Background"},
		{FgName: "TextMuted", BgName: "Background"},
		{FgName: "TextPrimary", BgName: "BackgroundSecondary"},
		{FgName: "TextPrimary", BgName: "Surface"},

		// Accent/interactive colors
		{FgName: "Accent", BgName: "Background"},

		// Semantic colors - text on semantic backgrounds
		{FgName: "Success.Text", BgName: "Success.Background"},
		{FgName: "Warning.Text", BgName: "Warning.Background"},
		{FgName: "Error.Text", BgName: "Error.Background"},
		{FgName: "Info.Text", BgName: "Info.Background"},

		// Code colors
		{FgName: "CodeText", BgName: "CodeBackground"},
		{FgName: "CodeComment", BgName: "CodeBackground"},
		{FgName: "CodeKeyword", BgName: "CodeBackground"},
		{FgName: "CodeString", BgName: "CodeBackground"},
	}
}
