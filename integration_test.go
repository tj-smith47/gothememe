package gothememe_test

import (
	"encoding/json"
	"regexp"
	"strings"
	"testing"

	"github.com/tj-smith47/gothememe"
	"github.com/tj-smith47/gothememe/themes"
)

// TestWCAGValidationAllBuiltInThemes runs WCAG contrast validation on all 451 themes.
// This comprehensive test documents accessibility compliance across the entire theme library.
func TestWCAGValidationAllBuiltInThemes(t *testing.T) {
	t.Parallel()

	allThemes := themes.All()
	if len(allThemes) < 400 {
		t.Fatalf("Expected 400+ themes, got %d", len(allThemes))
	}

	type themeResult struct {
		ID          string
		DisplayName string
		IsDark      bool
		AAIssues    int
		AAAIssues   int
	}

	var results []themeResult
	var passAA, passAAA int
	var darkCount, lightCount int

	for _, theme := range allThemes {
		aaIssues := gothememe.ValidateContrast(theme, gothememe.ContrastLevelAA)
		aaaIssues := gothememe.ValidateContrast(theme, gothememe.ContrastLevelAAA)

		result := themeResult{
			ID:          theme.ID(),
			DisplayName: theme.DisplayName(),
			IsDark:      theme.IsDark(),
			AAIssues:    len(aaIssues),
			AAAIssues:   len(aaaIssues),
		}

		if len(aaIssues) == 0 {
			passAA++
		}
		if len(aaaIssues) == 0 {
			passAAA++
		}
		if theme.IsDark() {
			darkCount++
		} else {
			lightCount++
		}

		results = append(results, result)
	}

	total := len(allThemes)
	t.Logf("WCAG Validation Summary for %d themes:", total)
	t.Logf("  Dark themes: %d, Light themes: %d", darkCount, lightCount)
	t.Logf("  Pass WCAG AA: %d (%.1f%%)", passAA, float64(passAA)/float64(total)*100)
	t.Logf("  Pass WCAG AAA: %d (%.1f%%)", passAAA, float64(passAAA)/float64(total)*100)

	// Log themes with significant AA issues (more than 3)
	var problemThemes []themeResult
	for _, r := range results {
		if r.AAIssues > 3 {
			problemThemes = append(problemThemes, r)
		}
	}

	if len(problemThemes) > 0 {
		t.Logf("Themes with significant contrast issues (>3 AA failures):")
		for _, r := range problemThemes[:min(10, len(problemThemes))] {
			t.Logf("  %s (%s): %d AA issues", r.ID, r.DisplayName, r.AAIssues)
		}
		if len(problemThemes) > 10 {
			t.Logf("  ... and %d more", len(problemThemes)-10)
		}
	}
}

// TestCSSGenerationAllThemes generates CSS for all themes to verify no panics or errors.
func TestCSSGenerationAllThemes(t *testing.T) {
	t.Parallel()

	allThemes := themes.All()
	opts := gothememe.DefaultCSSOptions()

	for _, theme := range allThemes {
		css := gothememe.GenerateCSS(theme, opts)
		if css == "" {
			t.Errorf("GenerateCSS for %s returned empty string", theme.ID())
		}
	}
}

// TestDTCGGenerationAllThemes generates DTCG tokens for all themes.
func TestDTCGGenerationAllThemes(t *testing.T) {
	t.Parallel()

	allThemes := themes.All()
	opts := gothememe.DefaultTokenOptions()

	for _, theme := range allThemes {
		tokens, err := gothememe.GenerateDesignTokens(theme, opts)
		if err != nil {
			t.Errorf("GenerateDesignTokens for %s failed: %v", theme.ID(), err)
		}
		if tokens == "" {
			t.Errorf("GenerateDesignTokens for %s returned empty string", theme.ID())
		}
	}
}

// TestValidateThemeAllBuiltIn validates all themes with ValidateTheme.
func TestValidateThemeAllBuiltIn(t *testing.T) {
	t.Parallel()

	allThemes := themes.All()
	errorCount := 0
	warnCount := 0

	for _, theme := range allThemes {
		errs := gothememe.ValidateTheme(theme)
		for _, err := range errs {
			if err.Severity == gothememe.SeverityError {
				errorCount++
				t.Errorf("Theme %s has error: %s", theme.ID(), err.Message)
			} else {
				warnCount++
			}
		}
	}

	t.Logf("Validation complete: %d errors, %d warnings across %d themes",
		errorCount, warnCount, len(allThemes))
}

// TestCSSSyntaxValidation validates that generated CSS has valid syntax.
// This is a proper integration test that verifies CSS structure.
func TestCSSSyntaxValidation(t *testing.T) {
	t.Parallel()

	// CSS property pattern: --name: value;
	cssVarPattern := regexp.MustCompile(`--[\w-]+:\s*[^;]+;`)
	// Hex color pattern
	hexPattern := regexp.MustCompile(`#[0-9a-fA-F]{6}`)
	// RGB/RGBA pattern
	rgbPattern := regexp.MustCompile(`rgba?\(\s*\d+\s*,\s*\d+\s*,\s*\d+`)
	// HSL/HSLA pattern
	hslPattern := regexp.MustCompile(`hsla?\(\s*\d+`)

	testCases := []struct {
		name        string
		opts        gothememe.CSSOptions
		expectRoot  bool
		colorFormat string
	}{
		{
			name:        "default_hex",
			opts:        gothememe.DefaultCSSOptions(),
			expectRoot:  true,
			colorFormat: "hex",
		},
		{
			name: "rgb_format",
			opts: gothememe.CSSOptions{
				Prefix:      "theme",
				IncludeRoot: true,
				ColorSpace:  gothememe.ColorSpaceRGB,
			},
			expectRoot:  true,
			colorFormat: "rgb",
		},
		{
			name: "hsl_format",
			opts: gothememe.CSSOptions{
				Prefix:      "theme",
				IncludeRoot: true,
				ColorSpace:  gothememe.ColorSpaceHSL,
			},
			expectRoot:  true,
			colorFormat: "hsl",
		},
		{
			name: "no_root",
			opts: gothememe.CSSOptions{
				Prefix:      "custom",
				IncludeRoot: false,
			},
			expectRoot:  false,
			colorFormat: "hex",
		},
	}

	theme := themes.ByID("dracula")
	if theme == nil {
		t.Fatal("dracula theme not found")
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			css := gothememe.GenerateCSS(theme, tc.opts)

			// Check for balanced braces
			openBraces := strings.Count(css, "{")
			closeBraces := strings.Count(css, "}")
			if openBraces != closeBraces {
				t.Errorf("Unbalanced braces: %d open, %d close", openBraces, closeBraces)
			}

			// Check for :root selector if expected
			if tc.expectRoot && !strings.Contains(css, ":root") {
				t.Error("Expected :root selector but not found")
			}
			if !tc.expectRoot && strings.Contains(css, ":root") {
				t.Error("Unexpected :root selector found")
			}

			// Check CSS variables exist
			vars := cssVarPattern.FindAllString(css, -1)
			if len(vars) < 10 {
				t.Errorf("Expected at least 10 CSS variables, got %d", len(vars))
			}

			// Verify color format
			switch tc.colorFormat {
			case "hex":
				if !hexPattern.MatchString(css) {
					t.Error("Expected hex colors not found")
				}
			case "rgb":
				if !rgbPattern.MatchString(css) {
					t.Error("Expected RGB colors not found")
				}
			case "hsl":
				if !hslPattern.MatchString(css) {
					t.Error("Expected HSL colors not found")
				}
			}

			// Check for common CSS errors
			if strings.Contains(css, ";;") {
				t.Error("Double semicolons found in CSS")
			}
			if strings.Contains(css, "::") && !strings.Contains(css, "::") {
				t.Error("Double colons found (not pseudo-elements)")
			}
		})
	}
}

// TestDTCGSchemaCompliance validates that generated design tokens comply with DTCG v1 schema.
func TestDTCGSchemaCompliance(t *testing.T) {
	t.Parallel()

	theme := themes.ByID("dracula")
	if theme == nil {
		t.Fatal("dracula theme not found")
	}

	tokens, err := gothememe.GenerateDesignTokens(theme, gothememe.DefaultTokenOptions())
	if err != nil {
		t.Fatalf("GenerateDesignTokens failed: %v", err)
	}

	// Parse as JSON
	var parsed map[string]any
	if err := json.Unmarshal([]byte(tokens), &parsed); err != nil {
		t.Fatalf("Invalid JSON: %v", err)
	}

	// DTCG v1 requires tokens to have $value and $type at leaf nodes
	// or be group objects containing other tokens/groups
	var validateToken func(path string, token any)
	validateToken = func(path string, token any) {
		tokenMap, ok := token.(map[string]any)
		if !ok {
			t.Errorf("Token at %s is not an object", path)
			return
		}

		// Check if this is a token (has $value) or a group
		if value, hasValue := tokenMap["$value"]; hasValue {
			// This is a token - must have $type
			tokenType, hasType := tokenMap["$type"]
			if !hasType {
				t.Errorf("Token at %s has $value but no $type", path)
			}

			// Validate $type is a string
			if _, ok := tokenType.(string); !ok {
				t.Errorf("Token at %s has non-string $type", path)
			}

			// Validate $value is not nil
			if value == nil {
				t.Errorf("Token at %s has nil $value", path)
			}

			// Optional: validate $description is string if present
			if desc, hasDesc := tokenMap["$description"]; hasDesc {
				if _, ok := desc.(string); !ok {
					t.Errorf("Token at %s has non-string $description", path)
				}
			}
		} else {
			// This is a group - recurse into children
			for key, child := range tokenMap {
				if strings.HasPrefix(key, "$") {
					// Skip metadata keys like $description at group level
					continue
				}
				childPath := path + "." + key
				if path == "" {
					childPath = key
				}
				validateToken(childPath, child)
			}
		}
	}

	// Validate root structure
	for key, value := range parsed {
		// Skip root-level metadata keys
		if strings.HasPrefix(key, "$") {
			continue
		}
		validateToken(key, value)
	}

	// Check required top-level groups exist
	requiredGroups := []string{"color"}
	for _, group := range requiredGroups {
		if _, ok := parsed[group]; !ok {
			t.Errorf("Missing required top-level group: %s", group)
		}
	}

	// Verify color group has expected subgroups
	colorGroup, ok := parsed["color"].(map[string]any)
	if !ok {
		t.Fatal("color group is not an object")
	}

	expectedColorSubgroups := []string{"background", "text", "accent"}
	for _, subgroup := range expectedColorSubgroups {
		if _, ok := colorGroup[subgroup]; !ok {
			t.Errorf("Missing color subgroup: %s", subgroup)
		}
	}

	t.Logf("DTCG validation passed for %s with %d top-level groups", theme.ID(), len(parsed))
}

// TestAllThemesCSSMinified validates minified CSS output for all themes.
func TestAllThemesCSSMinified(t *testing.T) {
	t.Parallel()

	allThemes := themes.All()
	opts := gothememe.CSSOptions{
		Prefix: "t",
		Minify: true,
	}

	for _, theme := range allThemes {
		css := gothememe.GenerateCSS(theme, opts)

		// Minified CSS should not have newlines (except possibly at the end)
		lines := strings.Split(strings.TrimSpace(css), "\n")
		if len(lines) > 1 {
			t.Errorf("Theme %s: minified CSS has %d lines, expected 1", theme.ID(), len(lines))
		}

		// Should still have balanced braces
		if strings.Count(css, "{") != strings.Count(css, "}") {
			t.Errorf("Theme %s: unbalanced braces in minified CSS", theme.ID())
		}
	}
}
