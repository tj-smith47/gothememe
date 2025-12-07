package generator

import (
	"fmt"
	"strings"
)

// SVGPreview generates an 8-color swatch SVG for a theme.
// The swatch displays: background, foreground, cursor, selection,
// and 4 ANSI colors (red, green, blue, yellow).
type SVGPreview struct {
	Width  int
	Height int
}

// DefaultSVGPreview returns the default SVG preview dimensions.
func DefaultSVGPreview() *SVGPreview {
	return &SVGPreview{
		Width:  400,
		Height: 80,
	}
}

// Generate creates an SVG preview for a theme.
// Returns the SVG content as a string.
func (s *SVGPreview) Generate(t *WindowsTerminalTheme) string {
	if t == nil {
		return ""
	}

	// 8 color swatches: bg, fg, selection, cursor, red, green, blue, yellow
	colors := []struct {
		color string
		label string
	}{
		{t.Background, "bg"},
		{t.Foreground, "fg"},
		{t.SelectionBackground, "sel"},
		{t.CursorColor, "cur"},
		{t.Red, "red"},
		{t.Green, "grn"},
		{t.Blue, "blu"},
		{t.Yellow, "yel"},
	}

	swatchWidth := s.Width / len(colors)
	swatchHeight := s.Height

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d">`,
		s.Width, s.Height, s.Width, s.Height))
	sb.WriteString("\n")

	// Add title for accessibility
	sb.WriteString(fmt.Sprintf("  <title>%s color theme preview</title>\n", escapeXML(t.Name)))

	// Generate color swatches
	for i, c := range colors {
		x := i * swatchWidth
		color := c.color
		if color == "" {
			color = "#808080" // Default gray for missing colors
		}

		// Swatch rectangle
		sb.WriteString(fmt.Sprintf(`  <rect x="%d" y="0" width="%d" height="%d" fill=%q/>`,
			x, swatchWidth, swatchHeight, escapeXML(color)))
		sb.WriteString("\n")

		// Add subtle separator lines between swatches
		if i > 0 {
			sb.WriteString(fmt.Sprintf(`  <line x1="%d" y1="0" x2="%d" y2="%d" stroke="rgba(0,0,0,0.1)" stroke-width="1"/>`,
				x, x, swatchHeight))
			sb.WriteString("\n")
		}
	}

	// Add border
	sb.WriteString(fmt.Sprintf(`  <rect x="0" y="0" width="%d" height="%d" fill="none" stroke="rgba(0,0,0,0.2)" stroke-width="1"/>`,
		s.Width, s.Height))
	sb.WriteString("\n")

	sb.WriteString("</svg>")
	return sb.String()
}

// GenerateInline creates an inline data URI for embedding in markdown/HTML.
func (s *SVGPreview) GenerateInline(t *WindowsTerminalTheme) string {
	svg := s.Generate(t)
	if svg == "" {
		return ""
	}
	// Simple URL encoding for data URI (replace problematic characters)
	encoded := strings.ReplaceAll(svg, "\n", "")
	encoded = strings.ReplaceAll(encoded, "\"", "'")
	encoded = strings.ReplaceAll(encoded, "#", "%23")
	encoded = strings.ReplaceAll(encoded, "<", "%3C")
	encoded = strings.ReplaceAll(encoded, ">", "%3E")
	return "data:image/svg+xml," + encoded
}

// escapeXML escapes special XML characters.
func escapeXML(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, "\"", "&quot;")
	s = strings.ReplaceAll(s, "'", "&apos;")
	return s
}

// GenerateSVGPreview is a convenience function that generates an SVG preview
// with default dimensions.
func GenerateSVGPreview(t *WindowsTerminalTheme) string {
	return DefaultSVGPreview().Generate(t)
}

// GenerateSVGPreviewInline is a convenience function that generates an inline
// data URI SVG preview with default dimensions.
func GenerateSVGPreviewInline(t *WindowsTerminalTheme) string {
	return DefaultSVGPreview().GenerateInline(t)
}
