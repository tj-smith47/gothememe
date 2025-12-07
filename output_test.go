package gothememe

import (
	"strings"
	"testing"
)

func TestGenerateCSS(t *testing.T) {
	t.Parallel()

	theme := NewThemeBuilder("test", "Test Theme").
		WithIsDark(true).
		WithBackground(Hex("#282a36")).
		WithTextPrimary(Hex("#f8f8f2")).
		WithAccent(Hex("#bd93f9")).
		Build()

	tests := []struct {
		name     string
		opts     CSSOptions
		contains []string
	}{
		{
			name: "default options",
			opts: CSSOptions{},
			contains: []string{
				"--theme-background:",
				"--theme-text-primary:",
				"--theme-accent:",
			},
		},
		{
			name: "with root selector",
			opts: CSSOptions{IncludeRoot: true},
			contains: []string{
				":root {",
				"--theme-background:",
			},
		},
		{
			name: "with data attribute",
			opts: CSSOptions{UseDataAttribute: true},
			contains: []string{
				`[data-theme="test"]`,
			},
		},
		{
			name: "minified",
			opts: CSSOptions{Minify: true, IncludeRoot: true},
			contains: []string{
				":root{",
			},
		},
		{
			name: "RGB color space",
			opts: CSSOptions{ColorSpace: ColorSpaceRGB},
			contains: []string{
				"rgb(",
			},
		},
		{
			name: "HSL color space",
			opts: CSSOptions{ColorSpace: ColorSpaceHSL},
			contains: []string{
				"hsl(",
			},
		},
		{
			name: "with metadata",
			opts: CSSOptions{IncludeMetadata: true, IncludeRoot: true},
			contains: []string{
				"/* Theme: Test Theme (test) */",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			css := GenerateCSS(theme, tt.opts)
			for _, want := range tt.contains {
				if !strings.Contains(css, want) {
					t.Errorf("GenerateCSS() missing %q\nGot:\n%s", want, css)
				}
			}
		})
	}
}

func TestGenerateSCSS(t *testing.T) {
	t.Parallel()

	theme := NewThemeBuilder("test", "Test Theme").
		WithIsDark(true).
		WithBackground(Hex("#282a36")).
		WithTextPrimary(Hex("#f8f8f2")).
		Build()

	scss := GenerateSCSS(theme, CSSOptions{})

	expected := []string{
		"$theme-background:",
		"$theme-text-primary:",
	}

	for _, want := range expected {
		if !strings.Contains(scss, want) {
			t.Errorf("GenerateSCSS() missing %q\nGot:\n%s", want, scss)
		}
	}
}

func TestGenerateJSON(t *testing.T) {
	t.Parallel()

	theme := NewThemeBuilder("test", "Test Theme").
		WithIsDark(true).
		WithBackground(Hex("#282a36")).
		WithTextPrimary(Hex("#f8f8f2")).
		Build()

	tests := []struct {
		name     string
		opts     CSSOptions
		contains []string
	}{
		{
			name: "pretty",
			opts: CSSOptions{},
			contains: []string{
				`"background"`,
				`"text-primary"`,
			},
		},
		{
			name: "minified",
			opts: CSSOptions{Minify: true},
			contains: []string{
				`"background": "#282a36"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			jsonOut := GenerateJSON(theme, tt.opts)
			for _, want := range tt.contains {
				if !strings.Contains(jsonOut, want) {
					t.Errorf("GenerateJSON() missing %q\nGot:\n%s", want, jsonOut)
				}
			}
		})
	}
}

func TestGenerateSyntaxCSS(t *testing.T) {
	t.Parallel()

	theme := NewThemeBuilder("test", "Test Theme").
		WithIsDark(true).
		WithCodeBackground(Hex("#282a36")).
		WithCodeText(Hex("#f8f8f2")).
		WithCodeKeyword(Hex("#ff79c6")).
		WithCodeString(Hex("#f1fa8c")).
		WithCodeComment(Hex("#6272a4")).
		Build()

	tests := []struct {
		name     string
		opts     SyntaxOptions
		contains []string
	}{
		{
			name: "Prism.js",
			opts: SyntaxOptions{Format: SyntaxPrism},
			contains: []string{
				".token.keyword",
				".token.string",
				".token.comment",
			},
		},
		{
			name: "Highlight.js",
			opts: SyntaxOptions{Format: SyntaxHighlightJS},
			contains: []string{
				".hljs-keyword",
				".hljs-string",
				".hljs-comment",
			},
		},
		{
			name: "Chroma",
			opts: SyntaxOptions{Format: SyntaxChroma},
			contains: []string{
				".chroma .k", // keyword
				".chroma .s", // string
				".chroma .c", // comment
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			css := GenerateSyntaxCSS(theme, tt.opts)
			for _, want := range tt.contains {
				if !strings.Contains(css, want) {
					t.Errorf("GenerateSyntaxCSS(%v) missing %q\nGot:\n%s", tt.opts.Format, want, css)
				}
			}
		})
	}
}

func TestGenerateAllThemesCSS(t *testing.T) {
	t.Parallel()

	theme1 := NewThemeBuilder("dark", "Dark").
		WithIsDark(true).
		WithBackground(Hex("#1e1e1e")).
		Build()

	theme2 := NewThemeBuilder("light", "Light").
		WithIsDark(false).
		WithBackground(Hex("#ffffff")).
		Build()

	css := GenerateAllThemesCSS([]Theme{theme1, theme2}, CSSOptions{
		UseDataAttribute: true,
	})

	expected := []string{
		`[data-theme="dark"]`,
		`[data-theme="light"]`,
		"#1e1e1e",
		"#ffffff",
	}

	for _, want := range expected {
		if !strings.Contains(css, want) {
			t.Errorf("GenerateAllThemesCSS() missing %q", want)
		}
	}
}

func TestDefaultCSSOptions(t *testing.T) {
	t.Parallel()

	opts := DefaultCSSOptions()

	if opts.IncludeRoot != true {
		t.Error("DefaultCSSOptions().IncludeRoot should be true")
	}
	if opts.UseDataAttribute != false {
		t.Error("DefaultCSSOptions().UseDataAttribute should be false")
	}
	if opts.ColorSpace != ColorSpaceHex {
		t.Errorf("DefaultCSSOptions().ColorSpace = %v, want Hex", opts.ColorSpace)
	}
}

func TestDefaultSyntaxOptions(t *testing.T) {
	t.Parallel()

	opts := DefaultSyntaxOptions()

	if opts.Format != SyntaxPrism {
		t.Errorf("DefaultSyntaxOptions().Format = %v, want Prism", opts.Format)
	}
	if opts.UseVariables != true {
		t.Error("DefaultSyntaxOptions().UseVariables should be true")
	}
}

func TestGenerateSCSS_AllOptions(t *testing.T) {
	t.Parallel()

	theme := NewThemeBuilder("test", "Test Theme").
		WithIsDark(true).
		WithBackground(Hex("#282a36")).
		WithTextPrimary(Hex("#f8f8f2")).
		WithAccent(Hex("#bd93f9")).
		WithSuccess(SemanticColor{Background: Hex("#50fa7b"), Text: Hex("#50fa7b")}).
		Build()

	tests := []struct {
		name     string
		opts     CSSOptions
		contains []string
	}{
		{
			name: "with metadata",
			opts: CSSOptions{IncludeMetadata: true},
			contains: []string{
				"// Theme: Test Theme",
				"$theme-background:",
			},
		},
		{
			name: "with nested maps",
			opts: CSSOptions{},
			contains: []string{
				"$theme-success-background:",
				"$theme-success-text:",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			scss := GenerateSCSS(theme, tt.opts)
			for _, want := range tt.contains {
				if !strings.Contains(scss, want) {
					t.Errorf("GenerateSCSS() missing %q\nGot:\n%s", want, scss)
				}
			}
		})
	}
}

func TestGenerateSyntaxCSS_UseVariables(t *testing.T) {
	t.Parallel()

	theme := NewThemeBuilder("test", "Test Theme").
		WithCodeBackground(Hex("#282a36")).
		WithCodeText(Hex("#f8f8f2")).
		WithCodeKeyword(Hex("#ff79c6")).
		Build()

	// With UseVariables
	css := GenerateSyntaxCSS(theme, SyntaxOptions{
		Format:       SyntaxPrism,
		UseVariables: true,
		Prefix:       "theme",
	})

	if !strings.Contains(css, "var(--theme-") {
		t.Error("GenerateSyntaxCSS with UseVariables should use CSS variables")
	}

	// Without UseVariables (inline colors)
	cssInline := GenerateSyntaxCSS(theme, SyntaxOptions{
		Format:       SyntaxPrism,
		UseVariables: false,
	})

	if strings.Contains(cssInline, "var(--") {
		t.Error("GenerateSyntaxCSS without UseVariables should not use CSS variables")
	}
	if !strings.Contains(cssInline, "#") {
		t.Error("GenerateSyntaxCSS without UseVariables should contain hex colors")
	}
}

// Benchmarks

func BenchmarkGenerateCSS(b *testing.B) {
	theme := NewThemeBuilder("bench", "Benchmark Theme").
		WithIsDark(true).
		WithBackground(Hex("#282a36")).
		WithBackgroundSecondary(Hex("#21222c")).
		WithTextPrimary(Hex("#f8f8f2")).
		WithTextSecondary(Hex("#bfbfbf")).
		WithAccent(Hex("#bd93f9")).
		WithAccentSecondary(Hex("#ff79c6")).
		WithSuccess(SemanticColor{Background: Hex("#50fa7b"), Text: Hex("#50fa7b")}).
		WithWarning(SemanticColor{Background: Hex("#f1fa8c"), Text: Hex("#f1fa8c")}).
		WithError(SemanticColor{Background: Hex("#ff5555"), Text: Hex("#ff5555")}).
		Build()

	opts := CSSOptions{IncludeRoot: true}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GenerateCSS(theme, opts)
	}
}

func BenchmarkGenerateCSS_Minified(b *testing.B) {
	theme := NewThemeBuilder("bench", "Benchmark Theme").
		WithIsDark(true).
		WithBackground(Hex("#282a36")).
		WithTextPrimary(Hex("#f8f8f2")).
		WithAccent(Hex("#bd93f9")).
		Build()

	opts := CSSOptions{IncludeRoot: true, Minify: true}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GenerateCSS(theme, opts)
	}
}

func BenchmarkGenerateAllThemesCSS(b *testing.B) {
	themes := make([]Theme, 10)
	for i := range themes {
		themes[i] = NewThemeBuilder(
			"theme"+string(rune('a'+i)),
			"Theme "+string(rune('A'+i)),
		).
			WithIsDark(i%2 == 0).
			WithBackground(Hex("#282a36")).
			WithTextPrimary(Hex("#f8f8f2")).
			WithAccent(Hex("#bd93f9")).
			Build()
	}

	opts := CSSOptions{UseDataAttribute: true}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GenerateAllThemesCSS(themes, opts)
	}
}

func BenchmarkGenerateSyntaxCSS(b *testing.B) {
	theme := NewThemeBuilder("bench", "Benchmark Theme").
		WithCodeBackground(Hex("#282a36")).
		WithCodeText(Hex("#f8f8f2")).
		WithCodeKeyword(Hex("#ff79c6")).
		WithCodeString(Hex("#f1fa8c")).
		WithCodeComment(Hex("#6272a4")).
		WithCodeNumber(Hex("#bd93f9")).
		WithCodeFunction(Hex("#50fa7b")).
		Build()

	opts := SyntaxOptions{Format: SyntaxPrism, UseVariables: true}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GenerateSyntaxCSS(theme, opts)
	}
}

func BenchmarkGenerateJSON(b *testing.B) {
	theme := NewThemeBuilder("bench", "Benchmark Theme").
		WithIsDark(true).
		WithBackground(Hex("#282a36")).
		WithTextPrimary(Hex("#f8f8f2")).
		WithAccent(Hex("#bd93f9")).
		Build()

	opts := CSSOptions{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GenerateJSON(theme, opts)
	}
}
