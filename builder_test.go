package gothememe

import (
	"testing"
)

func TestThemeBuilder(t *testing.T) {
	t.Parallel()

	theme := NewThemeBuilder("custom", "Custom Theme").
		WithDescription("A custom test theme").
		WithAuthor("Test Author").
		WithLicense("MIT").
		WithSource("https://example.com").
		WithIsDark(true).
		WithBackground(Hex("#1e1e1e")).
		WithBackgroundSecondary(Hex("#252526")).
		WithSurface(Hex("#2d2d2d")).
		WithTextPrimary(Hex("#d4d4d4")).
		WithTextSecondary(Hex("#808080")).
		WithAccent(Hex("#007acc")).
		Build()

	tests := []struct {
		name string
		got  string
		want string
	}{
		{"ID", theme.ID(), "custom"},
		{"DisplayName", theme.DisplayName(), "Custom Theme"},
		{"Description", theme.Description(), "A custom test theme"},
		{"Author", theme.Author(), "Test Author"},
		{"License", theme.License(), "MIT"},
		{"Source", theme.Source(), "https://example.com"},
		{"Background", theme.Background().Hex(), "#1e1e1e"},
		{"TextPrimary", theme.TextPrimary().Hex(), "#d4d4d4"},
		{"Accent", theme.Accent().Hex(), "#007acc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.got != tt.want {
				t.Errorf("%s = %q, want %q", tt.name, tt.got, tt.want)
			}
		})
	}

	if !theme.IsDark() {
		t.Error("IsDark() = false, want true")
	}
}

func TestThemeBuilderAutoDerivation(t *testing.T) {
	t.Parallel()

	// Build theme with minimal colors - others should be derived
	theme := NewThemeBuilder("minimal", "Minimal").
		WithIsDark(true).
		WithBackground(Hex("#282a36")).
		WithTextPrimary(Hex("#f8f8f2")).
		WithAccent(Hex("#bd93f9")).
		Build()

	// Derived colors should not be empty
	derivedColors := []struct {
		name  string
		color Color
	}{
		{"BackgroundSecondary", theme.BackgroundSecondary()},
		{"Surface", theme.Surface()},
		{"TextSecondary", theme.TextSecondary()},
		{"TextMuted", theme.TextMuted()},
		{"Border", theme.Border()},
	}

	for _, tc := range derivedColors {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if tc.color.IsEmpty() {
				t.Errorf("%s should be auto-derived, but is empty", tc.name)
			}
		})
	}
}

func TestGenerateThemeFromPalette(t *testing.T) {
	t.Parallel()

	palette := Palette{
		Background: Hex("#121212"),
		Foreground: Hex("#ffffff"),
		Accent:     Hex("#6200ee"),
	}

	theme := GenerateThemeFromPalette("material", "Material Dark", palette)

	if theme.ID() != "material" {
		t.Errorf("ID() = %q, want material", theme.ID())
	}

	if theme.DisplayName() != "Material Dark" {
		t.Errorf("DisplayName() = %q, want Material Dark", theme.DisplayName())
	}

	if theme.Background().Hex() != "#121212" {
		t.Errorf("Background() = %q, want #121212", theme.Background().Hex())
	}

	if theme.Accent().Hex() != "#6200ee" {
		t.Errorf("Accent() = %q, want #6200ee", theme.Accent().Hex())
	}
}

func TestDeriveTheme(t *testing.T) {
	t.Parallel()

	base := NewThemeBuilder("base", "Base").
		WithIsDark(true).
		WithBackground(Hex("#282a36")).
		WithTextPrimary(Hex("#f8f8f2")).
		WithAccent(Hex("#bd93f9")).
		Build()

	derived := DeriveTheme(base, "derived", "Derived Theme", map[string]Color{
		"accent": Hex("#ff79c6"),
	})

	// Should inherit base properties
	if derived.Background().Hex() != base.Background().Hex() {
		t.Errorf("Derived should inherit Background, got %q want %q",
			derived.Background().Hex(), base.Background().Hex())
	}

	// Should have overridden accent
	if derived.Accent().Hex() != "#ff79c6" {
		t.Errorf("Derived Accent() = %q, want #ff79c6", derived.Accent().Hex())
	}

	// Should have new ID
	if derived.ID() != "derived" {
		t.Errorf("Derived ID() = %q, want derived", derived.ID())
	}
}

func TestThemeBuilderANSIColors(t *testing.T) {
	t.Parallel()

	theme := NewThemeBuilder("ansi", "ANSI Test").
		WithIsDark(true).
		WithBackground(Hex("#000000")).
		WithTextPrimary(Hex("#ffffff")).
		WithBlack(Hex("#000000")).
		WithRed(Hex("#ff0000")).
		WithGreen(Hex("#00ff00")).
		WithYellow(Hex("#ffff00")).
		WithBlue(Hex("#0000ff")).
		WithPurple(Hex("#ff00ff")).
		WithCyan(Hex("#00ffff")).
		WithWhite(Hex("#ffffff")).
		WithBrightBlack(Hex("#808080")).
		WithBrightRed(Hex("#ff8080")).
		WithBrightGreen(Hex("#80ff80")).
		WithBrightYellow(Hex("#ffff80")).
		WithBrightBlue(Hex("#8080ff")).
		WithBrightPurple(Hex("#ff80ff")).
		WithBrightCyan(Hex("#80ffff")).
		WithBrightWhite(Hex("#ffffff")).
		Build()

	// Verify ANSI colors are set correctly
	if theme.Red().Hex() != "#ff0000" {
		t.Errorf("Red() = %q, want #ff0000", theme.Red().Hex())
	}
	if theme.BrightRed().Hex() != "#ff8080" {
		t.Errorf("BrightRed() = %q, want #ff8080", theme.BrightRed().Hex())
	}
}

func TestThemeBuilderCodeColors(t *testing.T) {
	t.Parallel()

	theme := NewThemeBuilder("code", "Code Test").
		WithIsDark(true).
		WithBackground(Hex("#1e1e1e")).
		WithTextPrimary(Hex("#d4d4d4")).
		WithCodeBackground(Hex("#1e1e1e")).
		WithCodeText(Hex("#d4d4d4")).
		WithCodeComment(Hex("#6a9955")).
		WithCodeKeyword(Hex("#569cd6")).
		WithCodeString(Hex("#ce9178")).
		WithCodeNumber(Hex("#b5cea8")).
		WithCodeFunction(Hex("#dcdcaa")).
		Build()

	if theme.CodeKeyword().Hex() != "#569cd6" {
		t.Errorf("CodeKeyword() = %q, want #569cd6", theme.CodeKeyword().Hex())
	}
	if theme.CodeString().Hex() != "#ce9178" {
		t.Errorf("CodeString() = %q, want #ce9178", theme.CodeString().Hex())
	}
}

func TestThemeBuilderSemanticColors(t *testing.T) {
	t.Parallel()

	successColor := SemanticColor{
		Background: Hex("#d4edda"),
		Border:     Hex("#c3e6cb"),
		Text:       Hex("#155724"),
	}

	theme := NewThemeBuilder("semantic", "Semantic Test").
		WithIsDark(false).
		WithBackground(Hex("#ffffff")).
		WithTextPrimary(Hex("#212529")).
		WithSuccess(successColor).
		Build()

	if theme.Success().Text.Hex() != "#155724" {
		t.Errorf("Success().Text = %q, want #155724", theme.Success().Text.Hex())
	}
}
