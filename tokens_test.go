package gothememe

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestGenerateDesignTokens(t *testing.T) {
	t.Parallel()

	theme := NewThemeBuilder("test", "Test Theme").
		WithIsDark(true).
		WithBackground(Hex("#282a36")).
		WithTextPrimary(Hex("#f8f8f2")).
		WithAccent(Hex("#bd93f9")).
		Build()

	tokens, err := GenerateDesignTokens(theme, DefaultTokenOptions())
	if err != nil {
		t.Fatalf("GenerateDesignTokens() error: %v", err)
	}

	// Verify it's valid JSON
	var parsed map[string]any
	if err := json.Unmarshal([]byte(tokens), &parsed); err != nil {
		t.Fatalf("GenerateDesignTokens() produced invalid JSON: %v", err)
	}

	// Check for DTCG structure - $description at root level, color object with $type
	expectedKeys := []string{
		"$description",
		"color",
	}

	for _, key := range expectedKeys {
		if _, ok := parsed[key]; !ok {
			t.Errorf("GenerateDesignTokens() missing root key %q", key)
		}
	}

	// The $type should be inside the color group
	if colorGroup, ok := parsed["color"].(map[string]any); ok {
		if _, hasType := colorGroup["$type"]; !hasType {
			t.Error("GenerateDesignTokens() missing $type in color group")
		}
	} else {
		t.Error("GenerateDesignTokens() color is not a map")
	}
}

func TestDesignTokensStructure(t *testing.T) {
	t.Parallel()

	theme := NewThemeBuilder("dtcg", "DTCG Test").
		WithIsDark(true).
		WithBackground(Hex("#1e1e1e")).
		WithTextPrimary(Hex("#d4d4d4")).
		WithAccent(Hex("#007acc")).
		WithSuccess(SemanticColor{
			Background: Hex("#d4edda"),
			Border:     Hex("#c3e6cb"),
			Text:       Hex("#155724"),
		}).
		Build()

	tokens, err := GenerateDesignTokens(theme, DefaultTokenOptions())
	if err != nil {
		t.Fatalf("GenerateDesignTokens() error: %v", err)
	}

	// Check hierarchical structure
	expectedPaths := []string{
		`"background"`,
		`"text"`,
		`"accent"`,
		`"semantic"`,
		`"$value"`,
	}

	for _, path := range expectedPaths {
		if !strings.Contains(tokens, path) {
			t.Errorf("Design tokens missing %s", path)
		}
	}
}

func TestDesignTokensColorFormat(t *testing.T) {
	t.Parallel()

	theme := NewThemeBuilder("colors", "Colors Test").
		WithIsDark(true).
		WithBackground(Hex("#282a36")).
		Build()

	tokens, err := GenerateDesignTokens(theme, DefaultTokenOptions())
	if err != nil {
		t.Fatalf("GenerateDesignTokens() error: %v", err)
	}

	// DTCG color tokens should have $type: "color"
	if !strings.Contains(tokens, `"$type": "color"`) && !strings.Contains(tokens, `"$type":"color"`) {
		t.Error("Design tokens should contain color type declarations")
	}

	// Should contain hex color values
	if !strings.Contains(tokens, "#282a36") {
		t.Error("Design tokens should contain hex color values")
	}
}

func TestDesignTokensMinimal(t *testing.T) {
	t.Parallel()

	// Create a minimal theme
	theme := NewThemeBuilder("minimal", "Minimal").
		WithIsDark(true).
		WithBackground(Hex("#000000")).
		WithTextPrimary(Hex("#ffffff")).
		Build()

	tokens, err := GenerateDesignTokens(theme, DefaultTokenOptions())
	if err != nil {
		t.Fatalf("GenerateDesignTokens() error: %v", err)
	}

	// Should still produce valid JSON
	var parsed map[string]any
	if err := json.Unmarshal([]byte(tokens), &parsed); err != nil {
		t.Fatalf("Minimal theme produced invalid JSON: %v", err)
	}
}

func TestGenerateAllDesignTokens(t *testing.T) {
	t.Parallel()

	theme1 := NewThemeBuilder("dark", "Dark").
		WithIsDark(true).
		WithBackground(Hex("#1e1e1e")).
		WithTextPrimary(Hex("#d4d4d4")).
		Build()

	theme2 := NewThemeBuilder("light", "Light").
		WithIsDark(false).
		WithBackground(Hex("#ffffff")).
		WithTextPrimary(Hex("#333333")).
		Build()

	tokens, err := GenerateAllDesignTokens([]Theme{theme1, theme2}, DefaultTokenOptions())
	if err != nil {
		t.Fatalf("GenerateAllDesignTokens() error: %v", err)
	}

	// Verify it's valid JSON
	var parsed map[string]any
	if err := json.Unmarshal([]byte(tokens), &parsed); err != nil {
		t.Fatalf("GenerateAllDesignTokens() produced invalid JSON: %v", err)
	}

	// Should contain both theme IDs
	if !strings.Contains(tokens, "dark") {
		t.Error("GenerateAllDesignTokens() missing 'dark' theme")
	}
	if !strings.Contains(tokens, "light") {
		t.Error("GenerateAllDesignTokens() missing 'light' theme")
	}
}

func TestDesignTokensWithDescriptions(t *testing.T) {
	t.Parallel()

	theme := NewThemeBuilder("test", "Test").
		WithIsDark(true).
		WithBackground(Hex("#1e1e1e")).
		WithDescription("A test theme").
		Build()

	opts := DefaultTokenOptions()
	opts.IncludeDescriptions = true

	tokens, err := GenerateDesignTokens(theme, opts)
	if err != nil {
		t.Fatalf("GenerateDesignTokens() error: %v", err)
	}

	// Should include $description
	if !strings.Contains(tokens, "$description") {
		t.Error("GenerateDesignTokens() with descriptions should include $description")
	}
}
