package gothememe

import (
	"strings"
	"testing"
)

func TestDefaultRegistryNil(t *testing.T) { //nolint:paralleltest // modifies global state
	// Note: These tests modify the global defaultRegistry
	// Reset at end
	oldRegistry := defaultRegistry
	defer func() { defaultRegistry = oldRegistry }()

	// Test with nil registry
	defaultRegistry = nil

	if GetDefaultRegistry() != nil {
		t.Error("GetDefaultRegistry() should return nil before NewDefaultRegistry")
	}
	if SetTheme(nil) {
		t.Error("SetTheme() should return false with nil registry")
	}
	if SetThemeID("any") {
		t.Error("SetThemeID() should return false with nil registry")
	}
	if theme, ok := GetTheme("any"); theme != nil || ok {
		t.Error("GetTheme() should return nil, false with nil registry")
	}
	if GetCurrentTheme() != nil {
		t.Error("GetCurrentTheme() should return nil with nil registry")
	}
	if Themes() != nil {
		t.Error("Themes() should return nil with nil registry")
	}
	if ThemeIDs() != nil {
		t.Error("ThemeIDs() should return nil with nil registry")
	}
	if ID() != "" {
		t.Error("ID() should return empty string with nil registry")
	}
	if DisplayName() != "" {
		t.Error("DisplayName() should return empty string with nil registry")
	}
	if IsDark() {
		t.Error("IsDark() should return false with nil registry")
	}
	if !Background().IsEmpty() {
		t.Error("Background() should return empty color with nil registry")
	}
	if !TextPrimary().IsEmpty() {
		t.Error("TextPrimary() should return empty color with nil registry")
	}
	if !Accent().IsEmpty() {
		t.Error("Accent() should return empty color with nil registry")
	}
	if CSS(DefaultCSSOptions()) != "" {
		t.Error("CSS() should return empty string with nil registry")
	}
	if AllThemesCSS(DefaultCSSOptions()) != "" {
		t.Error("AllThemesCSS() should return empty string with nil registry")
	}

	// These should not panic with nil registry
	NextTheme()
	PreviousTheme()
	Register()
	Unregister()
}

func TestDefaultRegistryWithThemes(t *testing.T) { //nolint:paralleltest // modifies global state
	// Note: These tests modify the global defaultRegistry
	// Reset at end
	oldRegistry := defaultRegistry
	defer func() { defaultRegistry = oldRegistry }()

	// Create a mock theme and set up the registry manually
	customTheme := NewThemeBuilder("test-theme", "Test Theme").
		WithIsDark(true).
		WithBackground(Hex("#282a36")).
		WithTextPrimary(Hex("#f8f8f2")).
		WithAccent(Hex("#bd93f9")).
		Build()

	// Set up the default registry with our theme
	defaultRegistry = NewRegistry(customTheme)

	if GetDefaultRegistry() == nil {
		t.Error("GetDefaultRegistry() should not be nil")
	}

	// Should have the theme we added
	themes := Themes()
	if len(themes) != 1 {
		t.Errorf("Themes() length = %d, want 1", len(themes))
	}

	ids := ThemeIDs()
	if len(ids) != 1 {
		t.Errorf("ThemeIDs() length = %d, want 1", len(ids))
	}

	// Current theme should be our test theme
	if ID() != "test-theme" {
		t.Errorf("ID() = %q, want 'test-theme'", ID())
	}
	if DisplayName() != "Test Theme" {
		t.Errorf("DisplayName() = %q, want 'Test Theme'", DisplayName())
	}
	if !IsDark() {
		t.Error("IsDark() should be true")
	}
	if Background().Hex() != "#282a36" {
		t.Errorf("Background() = %q, want '#282a36'", Background().Hex())
	}
	if TextPrimary().Hex() != "#f8f8f2" {
		t.Errorf("TextPrimary() = %q, want '#f8f8f2'", TextPrimary().Hex())
	}
	if Accent().Hex() != "#bd93f9" {
		t.Errorf("Accent() = %q, want '#bd93f9'", Accent().Hex())
	}

	// Test GetTheme
	theme, ok := GetTheme("test-theme")
	if !ok || theme == nil {
		t.Error("GetTheme('test-theme') should return the theme")
	}
	if theme.ID() != "test-theme" {
		t.Errorf("GetTheme('test-theme').ID() = %q, want 'test-theme'", theme.ID())
	}

	// Test GetCurrentTheme
	current := GetCurrentTheme()
	if current == nil || current.ID() != "test-theme" {
		t.Error("GetCurrentTheme() should return our test theme")
	}

	// Test SetTheme
	if !SetTheme(theme) {
		t.Error("SetTheme() should return true")
	}

	// Test navigation (with single theme, should stay the same)
	origID := ID()
	NextTheme()
	if ID() != origID {
		t.Error("NextTheme() with single theme should stay the same")
	}
	PreviousTheme()
	if ID() != origID {
		t.Error("PreviousTheme() with single theme should stay the same")
	}

	// Test CSS generation
	css := CSS(DefaultCSSOptions())
	if css == "" {
		t.Error("CSS() should not be empty with valid theme")
	}
	if !strings.Contains(css, "--theme-background") {
		t.Error("CSS() should contain theme variables")
	}

	allCSS := AllThemesCSS(CSSOptions{UseDataAttribute: true})
	if allCSS == "" {
		t.Error("AllThemesCSS() should not be empty")
	}
	if !strings.Contains(allCSS, `[data-theme="test-theme"]`) {
		t.Error("AllThemesCSS() should contain data-theme selector")
	}

	// Test Register
	anotherTheme := NewThemeBuilder("another", "Another").
		WithBackground(Hex("#123456")).
		Build()
	Register(anotherTheme)

	if _, ok := GetTheme("another"); !ok {
		t.Error("Registered theme should be found")
	}

	// Test Unregister
	Unregister(anotherTheme)
	if _, ok := GetTheme("another"); ok {
		t.Error("Unregistered theme should not be found")
	}
}

func TestNewDefaultRegistry(t *testing.T) { //nolint:paralleltest // modifies global state
	// Note: This tests the actual NewDefaultRegistry function behavior
	oldRegistry := defaultRegistry
	defer func() { defaultRegistry = oldRegistry }()

	// NewDefaultRegistry() calls DefaultThemes() which returns nil in main package
	// This tests that it handles empty themes gracefully
	NewDefaultRegistry()

	// Registry should exist (even if empty since DefaultThemes returns nil)
	if GetDefaultRegistry() == nil {
		t.Error("GetDefaultRegistry() should not be nil after NewDefaultRegistry")
	}
}
