package generator

import (
	"strings"
	"testing"
)

func TestDefaultSVGPreview(t *testing.T) {
	t.Parallel()

	preview := DefaultSVGPreview()

	if preview.Width != 400 {
		t.Errorf("Width = %d, want 400", preview.Width)
	}
	if preview.Height != 80 {
		t.Errorf("Height = %d, want 80", preview.Height)
	}
}

func TestSVGPreview_Generate(t *testing.T) {
	t.Parallel()

	preview := DefaultSVGPreview()
	theme := &WindowsTerminalTheme{
		Name:                "Dracula",
		Background:          "#282a36",
		Foreground:          "#f8f8f2",
		SelectionBackground: "#44475a",
		CursorColor:         "#f8f8f2",
		Red:                 "#ff5555",
		Green:               "#50fa7b",
		Blue:                "#bd93f9",
		Yellow:              "#f1fa8c",
	}

	svg := preview.Generate(theme)

	// Should contain SVG element
	if !strings.Contains(svg, "<svg") {
		t.Error("SVG should contain <svg element")
	}
	if !strings.Contains(svg, "</svg>") {
		t.Error("SVG should contain </svg> closing tag")
	}

	// Should contain title for accessibility
	if !strings.Contains(svg, "<title>Dracula color theme preview</title>") {
		t.Error("SVG should contain title element")
	}

	// Should contain all theme colors
	if !strings.Contains(svg, "#282a36") {
		t.Error("SVG should contain background color")
	}
	if !strings.Contains(svg, "#ff5555") {
		t.Error("SVG should contain red color")
	}

	// Should have 8 rect elements (one per color)
	rectCount := strings.Count(svg, "<rect")
	if rectCount != 9 { // 8 swatches + 1 border
		t.Errorf("SVG should have 9 rect elements (8 swatches + 1 border), got %d", rectCount)
	}
}

func TestSVGPreview_Generate_NilTheme(t *testing.T) {
	t.Parallel()

	preview := DefaultSVGPreview()
	svg := preview.Generate(nil)

	if svg != "" {
		t.Errorf("Generate(nil) should return empty string, got %q", svg)
	}
}

func TestSVGPreview_Generate_MissingColors(t *testing.T) {
	t.Parallel()

	preview := DefaultSVGPreview()
	theme := &WindowsTerminalTheme{
		Name:       "Minimal",
		Background: "#000000",
		// Missing most colors
	}

	svg := preview.Generate(theme)

	// Should still generate valid SVG
	if !strings.Contains(svg, "<svg") {
		t.Error("SVG should be generated even with missing colors")
	}

	// Should use default gray for missing colors
	if !strings.Contains(svg, "#808080") {
		t.Error("SVG should use default gray for missing colors")
	}
}

func TestSVGPreview_GenerateInline(t *testing.T) {
	t.Parallel()

	preview := DefaultSVGPreview()
	theme := &WindowsTerminalTheme{
		Name:                "Test",
		Background:          "#000000",
		Foreground:          "#ffffff",
		SelectionBackground: "#333333",
		CursorColor:         "#ffffff",
		Red:                 "#ff0000",
		Green:               "#00ff00",
		Blue:                "#0000ff",
		Yellow:              "#ffff00",
	}

	dataURI := preview.GenerateInline(theme)

	// Should start with data URI prefix
	if !strings.HasPrefix(dataURI, "data:image/svg+xml,") {
		t.Error("Inline SVG should start with data:image/svg+xml,")
	}

	// Should not contain newlines
	if strings.Contains(dataURI, "\n") {
		t.Error("Inline SVG should not contain newlines")
	}

	// Should have encoded special characters
	if strings.Contains(dataURI[20:], "#") { // Skip the "data:image/svg+xml," prefix
		t.Error("Hash characters should be encoded as %23")
	}
}

func TestSVGPreview_GenerateInline_NilTheme(t *testing.T) {
	t.Parallel()

	preview := DefaultSVGPreview()
	dataURI := preview.GenerateInline(nil)

	if dataURI != "" {
		t.Errorf("GenerateInline(nil) should return empty string, got %q", dataURI)
	}
}

func TestGenerateSVGPreview(t *testing.T) {
	t.Parallel()

	theme := &WindowsTerminalTheme{
		Name:       "Test",
		Background: "#282a36",
	}

	svg := GenerateSVGPreview(theme)

	if !strings.Contains(svg, "<svg") {
		t.Error("Convenience function should return valid SVG")
	}
}

func TestGenerateSVGPreviewInline(t *testing.T) {
	t.Parallel()

	theme := &WindowsTerminalTheme{
		Name:       "Test",
		Background: "#282a36",
	}

	dataURI := GenerateSVGPreviewInline(theme)

	if !strings.HasPrefix(dataURI, "data:image/svg+xml,") {
		t.Error("Convenience function should return data URI")
	}
}

func TestEscapeXML(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  string
	}{
		{"hello", "hello"},
		{"<script>", "&lt;script&gt;"},
		{"a & b", "a &amp; b"},
		{`"quoted"`, "&quot;quoted&quot;"},
		{"it's", "it&apos;s"},
		{"<>&\"'", "&lt;&gt;&amp;&quot;&apos;"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			got := escapeXML(tt.input)
			if got != tt.want {
				t.Errorf("escapeXML(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestSVGPreview_CustomDimensions(t *testing.T) {
	t.Parallel()

	preview := &SVGPreview{Width: 200, Height: 50}
	theme := &WindowsTerminalTheme{
		Name:       "Test",
		Background: "#000000",
	}

	svg := preview.Generate(theme)

	if !strings.Contains(svg, `width="200"`) {
		t.Error("SVG should use custom width")
	}
	if !strings.Contains(svg, `height="50"`) {
		t.Error("SVG should use custom height")
	}
}
