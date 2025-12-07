package gothememe

import (
	"testing"
)

func TestAnalyzeTheme(t *testing.T) {
	t.Parallel()

	theme := NewThemeBuilder("test", "Test").
		WithIsDark(true).
		WithBackground(Hex("#000000")).
		WithTextPrimary(Hex("#ffffff")).
		WithTextSecondary(Hex("#cccccc")).
		WithAccent(Hex("#ff0000")).
		Build()

	stats := AnalyzeTheme(theme)

	if stats.ColorCount == 0 {
		t.Error("ColorCount should be greater than 0")
	}

	if stats.UniqueColors == 0 {
		t.Error("UniqueColors should be greater than 0")
	}

	if !stats.IsDark {
		t.Error("IsDark should be true")
	}

	if stats.TotalPairs == 0 {
		t.Error("TotalPairs should be greater than 0")
	}

	// High contrast theme should have good accessibility
	if stats.ContrastScore < 4.5 {
		t.Logf("ContrastScore: %.2f", stats.ContrastScore)
	}
}

func TestAnalyzeThemeLuminance(t *testing.T) {
	t.Parallel()

	dark := NewThemeBuilder("dark", "Dark").
		WithIsDark(true).
		WithBackground(Hex("#000000")).
		WithTextPrimary(Hex("#ffffff")).
		Build()

	stats := AnalyzeTheme(dark)

	// Black background should have low luminance
	if stats.BackgroundLuminance > 0.1 {
		t.Errorf("BackgroundLuminance = %.3f, want < 0.1", stats.BackgroundLuminance)
	}

	// White text should have high luminance
	if stats.AverageTextLuminance < 0.9 {
		t.Errorf("AverageTextLuminance = %.3f, want > 0.9", stats.AverageTextLuminance)
	}
}

func TestCompareThemes(t *testing.T) {
	t.Parallel()

	const (
		highID = "high"
		lowID  = "low"
	)

	highContrastSemantic := SemanticColor{
		Background: Hex("#000000"),
		Border:     Hex("#ffffff"),
		Text:       Hex("#ffffff"),
	}

	highContrast := NewThemeBuilder(highID, "High Contrast").
		WithIsDark(true).
		WithBackground(Hex("#000000")).
		WithTextPrimary(Hex("#ffffff")).
		WithTextSecondary(Hex("#ffffff")).
		WithTextMuted(Hex("#ffffff")).
		WithAccent(Hex("#ffffff")).
		WithCodeBackground(Hex("#000000")).
		WithCodeText(Hex("#ffffff")).
		WithCodeComment(Hex("#ffffff")).
		WithCodeKeyword(Hex("#ffffff")).
		WithCodeString(Hex("#ffffff")).
		WithSuccess(highContrastSemantic).
		WithWarning(highContrastSemantic).
		WithError(highContrastSemantic).
		WithInfo(highContrastSemantic).
		Build()

	lowContrast := NewThemeBuilder(lowID, "Low Contrast").
		WithIsDark(true).
		WithBackground(Hex("#444444")).
		WithTextPrimary(Hex("#666666")).
		Build()

	comparison := CompareThemes(lowContrast, highContrast)

	if comparison.ThemeA != lowID {
		t.Errorf("ThemeA = %q, want %q", comparison.ThemeA, lowID)
	}

	if comparison.ThemeB != highID {
		t.Errorf("ThemeB = %q, want %q", comparison.ThemeB, highID)
	}

	// High contrast should have better accessibility
	if comparison.AccessDiff <= 0 {
		t.Errorf("AccessDiff = %.2f, want > 0 (high contrast should be better)", comparison.AccessDiff)
	}

	if !comparison.SameDarkMode {
		t.Error("SameDarkMode should be true")
	}
}

func TestFilterAccessible(t *testing.T) {
	t.Parallel()

	const highID = "accessible"

	highContrastSemantic := SemanticColor{
		Background: Hex("#000000"),
		Border:     Hex("#ffffff"),
		Text:       Hex("#ffffff"),
	}

	themes := []Theme{
		NewThemeBuilder(highID, "Accessible").
			WithIsDark(true).
			WithBackground(Hex("#000000")).
			WithTextPrimary(Hex("#ffffff")).
			WithTextSecondary(Hex("#ffffff")).
			WithTextMuted(Hex("#ffffff")).
			WithAccent(Hex("#ffffff")).
			WithCodeBackground(Hex("#000000")).
			WithCodeText(Hex("#ffffff")).
			WithCodeComment(Hex("#ffffff")).
			WithCodeKeyword(Hex("#ffffff")).
			WithCodeString(Hex("#ffffff")).
			WithSuccess(highContrastSemantic).
			WithWarning(highContrastSemantic).
			WithError(highContrastSemantic).
			WithInfo(highContrastSemantic).
			Build(),
		NewThemeBuilder("inaccessible", "Inaccessible").
			WithIsDark(true).
			WithBackground(Hex("#555555")).
			WithTextPrimary(Hex("#666666")).
			Build(),
	}

	accessible := FilterAccessible(themes, 80.0)

	// Only high contrast should pass
	if len(accessible) != 1 {
		t.Errorf("FilterAccessible returned %d themes, want 1", len(accessible))
	}

	if len(accessible) > 0 && accessible[0].ID() != highID {
		t.Errorf("Filtered theme ID = %q, want %q", accessible[0].ID(), highID)
	}
}

func TestSortByAccessibility(t *testing.T) {
	t.Parallel()

	const bestID = "best"

	highContrastSemantic := SemanticColor{
		Background: Hex("#000000"),
		Border:     Hex("#ffffff"),
		Text:       Hex("#ffffff"),
	}

	themes := []Theme{
		NewThemeBuilder("worst", "Worst").
			WithIsDark(true).
			WithBackground(Hex("#555555")).
			WithTextPrimary(Hex("#666666")).
			Build(),
		NewThemeBuilder(bestID, "Best").
			WithIsDark(true).
			WithBackground(Hex("#000000")).
			WithTextPrimary(Hex("#ffffff")).
			WithTextSecondary(Hex("#ffffff")).
			WithTextMuted(Hex("#ffffff")).
			WithAccent(Hex("#ffffff")).
			WithCodeBackground(Hex("#000000")).
			WithCodeText(Hex("#ffffff")).
			WithCodeComment(Hex("#ffffff")).
			WithCodeKeyword(Hex("#ffffff")).
			WithCodeString(Hex("#ffffff")).
			WithSuccess(highContrastSemantic).
			WithWarning(highContrastSemantic).
			WithError(highContrastSemantic).
			WithInfo(highContrastSemantic).
			Build(),
	}

	sorted := SortByAccessibility(themes)

	if len(sorted) != 2 {
		t.Fatalf("SortByAccessibility returned %d themes, want 2", len(sorted))
	}

	// High contrast should be first
	if sorted[0].ID() != bestID {
		t.Errorf("First sorted theme = %q, want %q", sorted[0].ID(), bestID)
	}
}

func TestAnalyzeAll(t *testing.T) {
	t.Parallel()

	themes := []Theme{
		NewThemeBuilder("one", "One").
			WithIsDark(true).
			WithBackground(Hex("#000000")).
			WithTextPrimary(Hex("#ffffff")).
			Build(),
		NewThemeBuilder("two", "Two").
			WithIsDark(false).
			WithBackground(Hex("#ffffff")).
			WithTextPrimary(Hex("#000000")).
			Build(),
	}

	stats := AnalyzeAll(themes)

	if len(stats) != 2 {
		t.Errorf("AnalyzeAll returned %d stats, want 2", len(stats))
	}

	if stats[0].IsDark != true {
		t.Error("First theme should be dark")
	}

	if stats[1].IsDark != false {
		t.Error("Second theme should be light")
	}
}

func TestCountUnique(t *testing.T) {
	t.Parallel()

	colors := []Color{
		Hex("#ffffff"),
		Hex("#ffffff"), // duplicate
		Hex("#000000"),
		{}, // empty
	}

	count := countUnique(colors)
	if count != 2 {
		t.Errorf("countUnique = %d, want 2", count)
	}
}

func TestAverageLuminance(t *testing.T) {
	t.Parallel()

	colors := []Color{
		Hex("#ffffff"), // luminance ~1.0
		Hex("#000000"), // luminance ~0.0
	}

	avg := averageLuminance(colors)
	if avg < 0.4 || avg > 0.6 {
		t.Errorf("averageLuminance = %.3f, want ~0.5", avg)
	}
}

func TestAverageLuminanceEmpty(t *testing.T) {
	t.Parallel()

	colors := []Color{{}, {}}

	avg := averageLuminance(colors)
	if avg != 0 {
		t.Errorf("averageLuminance of empty colors = %.3f, want 0", avg)
	}
}
