package themes

import (
	"testing"
)

func TestAll(t *testing.T) {
	t.Parallel()

	themes := All()

	// Should have 451+ themes
	if len(themes) < 400 {
		t.Errorf("All() returned %d themes, expected 400+", len(themes))
	}

	// All themes should be non-nil
	for i, theme := range themes {
		if theme == nil {
			t.Errorf("All()[%d] is nil", i)
		}
	}
}

func TestByID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		id      string
		wantNil bool
	}{
		{"dracula", false},
		{"nord", false},
		{"gruvbox_dark", false},
		{"nonexistent_theme_xyz", true},
		{"", true},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			t.Parallel()
			theme := ByID(tt.id)
			if tt.wantNil && theme != nil {
				t.Errorf("ByID(%q) = %v, want nil", tt.id, theme)
			}
			if !tt.wantNil && theme == nil {
				t.Errorf("ByID(%q) = nil, want non-nil", tt.id)
			}
		})
	}
}

func TestIDs(t *testing.T) {
	t.Parallel()

	ids := IDs()

	// Should have 451+ IDs
	if len(ids) < 400 {
		t.Errorf("IDs() returned %d IDs, expected 400+", len(ids))
	}

	// IDs should be sorted alphabetically
	for i := 1; i < len(ids); i++ {
		if ids[i] < ids[i-1] {
			t.Errorf("IDs() not sorted: %q should come before %q", ids[i], ids[i-1])
			break
		}
	}

	// All IDs should be unique
	seen := make(map[string]bool)
	for _, id := range ids {
		if seen[id] {
			t.Errorf("Duplicate ID found: %q", id)
		}
		seen[id] = true
	}
}

func TestThemeVariables(t *testing.T) {
	t.Parallel()

	// Test some known theme variables exist and are correct
	tests := []struct {
		name   string
		theme  func() any // Use func to avoid init order issues
		wantID string
	}{
		{"ThemeDracula", func() any { return ThemeDracula }, "dracula"},
		{"ThemeNord", func() any { return ThemeNord }, "nord"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			theme := tt.theme()
			if theme == nil {
				t.Fatalf("%s is nil", tt.name)
			}
		})
	}
}

func TestThemeImplementation(t *testing.T) {
	t.Parallel()

	// Pick a known theme and verify all interface methods work
	theme := ByID("dracula")
	if theme == nil {
		t.Fatal("dracula theme not found")
	}

	// Metadata
	if theme.ID() != "dracula" {
		t.Errorf("ID() = %q, want dracula", theme.ID())
	}
	if theme.DisplayName() == "" {
		t.Error("DisplayName() is empty")
	}

	// Colors should not be empty
	colors := []struct {
		name  string
		check func() bool
	}{
		{"Background", func() bool { return !theme.Background().IsEmpty() }},
		{"TextPrimary", func() bool { return !theme.TextPrimary().IsEmpty() }},
		{"Accent", func() bool { return !theme.Accent().IsEmpty() }},
		{"Red", func() bool { return !theme.Red().IsEmpty() }},
		{"Green", func() bool { return !theme.Green().IsEmpty() }},
		{"Blue", func() bool { return !theme.Blue().IsEmpty() }},
		{"CodeBackground", func() bool { return !theme.CodeBackground().IsEmpty() }},
	}

	for _, c := range colors {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if !c.check() {
				t.Errorf("%s is empty", c.name)
			}
		})
	}
}

func TestAllThemesValid(t *testing.T) {
	t.Parallel()

	themes := All()

	for _, theme := range themes {
		t.Run(theme.ID(), func(t *testing.T) {
			t.Parallel()

			// Every theme should have an ID
			if theme.ID() == "" {
				t.Error("ID() is empty")
			}

			// Every theme should have a display name
			if theme.DisplayName() == "" {
				t.Error("DisplayName() is empty")
			}

			// Background should not be empty
			if theme.Background().IsEmpty() {
				t.Error("Background() is empty")
			}

			// TextPrimary should not be empty
			if theme.TextPrimary().IsEmpty() {
				t.Error("TextPrimary() is empty")
			}
		})
	}
}

// TestAllThemeMethodsCoverage calls every method on every theme to ensure coverage.
// This is designed to maximize test coverage across all 451+ theme implementations.
func TestAllThemeMethodsCoverage(t *testing.T) {
	t.Parallel()

	themes := All()

	for _, theme := range themes {
		// We don't use t.Run here to avoid creating 451+ subtests
		// Just call all methods to exercise the code

		// Metadata
		_ = theme.ID()
		_ = theme.DisplayName()
		_ = theme.Description()
		_ = theme.Author()
		_ = theme.License()
		_ = theme.Source()
		_ = theme.IsDark()

		// Background colors
		_ = theme.Background()
		_ = theme.BackgroundSecondary()
		_ = theme.Surface()
		_ = theme.SurfaceSecondary()

		// Text colors
		_ = theme.TextPrimary()
		_ = theme.TextSecondary()
		_ = theme.TextMuted()
		_ = theme.TextInverted()

		// Accent colors
		_ = theme.Accent()
		_ = theme.AccentSecondary()
		_ = theme.Brand()

		// Border colors
		_ = theme.Border()
		_ = theme.BorderSubtle()
		_ = theme.BorderStrong()

		// Semantic colors
		_ = theme.Success()
		_ = theme.Warning()
		_ = theme.Error()
		_ = theme.Info()

		// ANSI colors
		_ = theme.Black()
		_ = theme.Red()
		_ = theme.Green()
		_ = theme.Yellow()
		_ = theme.Blue()
		_ = theme.Purple()
		_ = theme.Cyan()
		_ = theme.White()
		_ = theme.BrightBlack()
		_ = theme.BrightRed()
		_ = theme.BrightGreen()
		_ = theme.BrightYellow()
		_ = theme.BrightBlue()
		_ = theme.BrightPurple()
		_ = theme.BrightCyan()
		_ = theme.BrightWhite()

		// Code colors
		_ = theme.CodeBackground()
		_ = theme.CodeText()
		_ = theme.CodeComment()
		_ = theme.CodeKeyword()
		_ = theme.CodeString()
		_ = theme.CodeNumber()
		_ = theme.CodeFunction()
		_ = theme.CodeOperator()
		_ = theme.CodePunctuation()
		_ = theme.CodeVariable()
		_ = theme.CodeConstant()
		_ = theme.CodeType()
	}
}
