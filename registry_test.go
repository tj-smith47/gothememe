package gothememe

import (
	"sync"
	"testing"
)

// mockTheme is a minimal theme implementation for testing.
type mockTheme struct {
	id          string
	displayName string
}

func (m *mockTheme) ID() string                 { return m.id }
func (m *mockTheme) DisplayName() string        { return m.displayName }
func (m *mockTheme) Description() string        { return "Test theme" }
func (m *mockTheme) Author() string             { return "Test" }
func (m *mockTheme) License() string            { return "MIT" }
func (m *mockTheme) Source() string             { return "" }
func (m *mockTheme) IsDark() bool               { return true }
func (m *mockTheme) Background() Color          { return Hex("#282a36") }
func (m *mockTheme) BackgroundSecondary() Color { return Hex("#21222c") }
func (m *mockTheme) Surface() Color             { return Hex("#343746") }
func (m *mockTheme) SurfaceSecondary() Color    { return Hex("#3d4051") }
func (m *mockTheme) TextPrimary() Color         { return Hex("#f8f8f2") }
func (m *mockTheme) TextSecondary() Color       { return Hex("#bfbfbf") }
func (m *mockTheme) TextMuted() Color           { return Hex("#6272a4") }
func (m *mockTheme) TextInverted() Color        { return Hex("#282a36") }
func (m *mockTheme) Accent() Color              { return Hex("#bd93f9") }
func (m *mockTheme) AccentSecondary() Color     { return Hex("#ff79c6") }
func (m *mockTheme) Brand() Color               { return Hex("#bd93f9") }
func (m *mockTheme) Border() Color              { return Hex("#44475a") }
func (m *mockTheme) BorderSubtle() Color        { return Hex("#383a46") }
func (m *mockTheme) BorderStrong() Color        { return Hex("#6272a4") }
func (m *mockTheme) Success() SemanticColor     { return SemanticColor{Text: Hex("#50fa7b")} }
func (m *mockTheme) Warning() SemanticColor     { return SemanticColor{Text: Hex("#f1fa8c")} }
func (m *mockTheme) Error() SemanticColor       { return SemanticColor{Text: Hex("#ff5555")} }
func (m *mockTheme) Info() SemanticColor        { return SemanticColor{Text: Hex("#8be9fd")} }
func (m *mockTheme) Black() Color               { return Hex("#21222c") }
func (m *mockTheme) Red() Color                 { return Hex("#ff5555") }
func (m *mockTheme) Green() Color               { return Hex("#50fa7b") }
func (m *mockTheme) Yellow() Color              { return Hex("#f1fa8c") }
func (m *mockTheme) Blue() Color                { return Hex("#bd93f9") }
func (m *mockTheme) Purple() Color              { return Hex("#ff79c6") }
func (m *mockTheme) Cyan() Color                { return Hex("#8be9fd") }
func (m *mockTheme) White() Color               { return Hex("#f8f8f2") }
func (m *mockTheme) BrightBlack() Color         { return Hex("#6272a4") }
func (m *mockTheme) BrightRed() Color           { return Hex("#ff6e6e") }
func (m *mockTheme) BrightGreen() Color         { return Hex("#69ff94") }
func (m *mockTheme) BrightYellow() Color        { return Hex("#ffffa5") }
func (m *mockTheme) BrightBlue() Color          { return Hex("#d6acff") }
func (m *mockTheme) BrightPurple() Color        { return Hex("#ff92df") }
func (m *mockTheme) BrightCyan() Color          { return Hex("#a4ffff") }
func (m *mockTheme) BrightWhite() Color         { return Hex("#ffffff") }
func (m *mockTheme) CodeBackground() Color      { return Hex("#282a36") }
func (m *mockTheme) CodeText() Color            { return Hex("#f8f8f2") }
func (m *mockTheme) CodeComment() Color         { return Hex("#6272a4") }
func (m *mockTheme) CodeKeyword() Color         { return Hex("#ff79c6") }
func (m *mockTheme) CodeString() Color          { return Hex("#f1fa8c") }
func (m *mockTheme) CodeNumber() Color          { return Hex("#bd93f9") }
func (m *mockTheme) CodeFunction() Color        { return Hex("#50fa7b") }
func (m *mockTheme) CodeOperator() Color        { return Hex("#ff79c6") }
func (m *mockTheme) CodePunctuation() Color     { return Hex("#f8f8f2") }
func (m *mockTheme) CodeVariable() Color        { return Hex("#f8f8f2") }
func (m *mockTheme) CodeConstant() Color        { return Hex("#bd93f9") }
func (m *mockTheme) CodeType() Color            { return Hex("#8be9fd") }

func TestNewRegistry(t *testing.T) {
	t.Parallel()
	theme1 := &mockTheme{id: "theme1", displayName: "Theme 1"}
	theme2 := &mockTheme{id: "theme2", displayName: "Theme 2"}

	r := NewRegistry(theme1, theme2)

	if r.Count() != 2 {
		t.Errorf("Count() = %d, want 2", r.Count())
	}

	if r.GetCurrentTheme().ID() != "theme1" {
		t.Errorf("Current theme = %q, want theme1", r.GetCurrentTheme().ID())
	}
}

func TestRegistrySetTheme(t *testing.T) {
	t.Parallel()
	theme1 := &mockTheme{id: "theme1", displayName: "Theme 1"}
	theme2 := &mockTheme{id: "theme2", displayName: "Theme 2"}

	r := NewRegistry(theme1, theme2)

	if !r.SetThemeID("theme2") {
		t.Error("SetThemeID(theme2) should return true")
	}

	if r.GetCurrentTheme().ID() != "theme2" {
		t.Errorf("Current theme = %q, want theme2", r.GetCurrentTheme().ID())
	}

	if r.SetThemeID("nonexistent") {
		t.Error("SetThemeID(nonexistent) should return false")
	}
}

func TestRegistryNavigation(t *testing.T) {
	t.Parallel()
	theme1 := &mockTheme{id: "aaa", displayName: "AAA"}
	theme2 := &mockTheme{id: "bbb", displayName: "BBB"}
	theme3 := &mockTheme{id: "ccc", displayName: "CCC"}

	r := NewRegistry(theme1, theme2, theme3)

	// Should be sorted alphabetically, starting with first registered
	if r.ID() != "aaa" {
		t.Errorf("Initial theme = %q, want aaa", r.ID())
	}

	r.NextTheme()
	if r.ID() != "bbb" {
		t.Errorf("After NextTheme() = %q, want bbb", r.ID())
	}

	r.NextTheme()
	if r.ID() != "ccc" {
		t.Errorf("After NextTheme() = %q, want ccc", r.ID())
	}

	r.NextTheme() // Should wrap around
	if r.ID() != "aaa" {
		t.Errorf("After wrap NextTheme() = %q, want aaa", r.ID())
	}

	r.PreviousTheme() // Should wrap around backwards
	if r.ID() != "ccc" {
		t.Errorf("After wrap PreviousTheme() = %q, want ccc", r.ID())
	}
}

func TestRegistryRegisterUnregister(t *testing.T) {
	t.Parallel()
	theme1 := &mockTheme{id: "theme1", displayName: "Theme 1"}
	theme2 := &mockTheme{id: "theme2", displayName: "Theme 2"}

	r := NewRegistry(theme1)

	if r.Count() != 1 {
		t.Errorf("Initial count = %d, want 1", r.Count())
	}

	r.Register(theme2)
	if r.Count() != 2 {
		t.Errorf("After register count = %d, want 2", r.Count())
	}

	r.Unregister(theme2)
	if r.Count() != 1 {
		t.Errorf("After unregister count = %d, want 1", r.Count())
	}

	r.UnregisterAll()
	if r.Count() != 0 {
		t.Errorf("After unregister all count = %d, want 0", r.Count())
	}
}

func TestRegistryConcurrency(t *testing.T) {
	t.Parallel()
	theme1 := &mockTheme{id: "theme1", displayName: "Theme 1"}
	theme2 := &mockTheme{id: "theme2", displayName: "Theme 2"}

	r := NewRegistry(theme1, theme2)

	var wg sync.WaitGroup

	// Concurrent readers
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				_ = r.GetCurrentTheme()
				_ = r.Themes()
				_ = r.ThemeIDs()
				_ = r.ID()
				_ = r.Background()
			}
		}()
	}

	// Concurrent writers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				r.NextTheme()
				r.PreviousTheme()
			}
		}()
	}

	wg.Wait()

	// Verify registry is still in a valid state after concurrent access
	if r.Count() != 2 {
		t.Errorf("Count after concurrent access = %d, want 2", r.Count())
	}
}

func TestRegistryThemeIDs(t *testing.T) {
	t.Parallel()
	theme1 := &mockTheme{id: "zeta", displayName: "Zeta"}
	theme2 := &mockTheme{id: "alpha", displayName: "Alpha"}
	theme3 := &mockTheme{id: "beta", displayName: "Beta"}

	r := NewRegistry(theme1, theme2, theme3)

	ids := r.ThemeIDs()
	expected := []string{"alpha", "beta", "zeta"}

	if len(ids) != len(expected) {
		t.Fatalf("ThemeIDs() length = %d, want %d", len(ids), len(expected))
	}

	for i, id := range ids {
		if id != expected[i] {
			t.Errorf("ThemeIDs()[%d] = %q, want %q", i, id, expected[i])
		}
	}
}

func TestRegistryGetTheme(t *testing.T) {
	t.Parallel()
	theme1 := &mockTheme{id: "theme1", displayName: "Theme 1"}
	theme2 := &mockTheme{id: "theme2", displayName: "Theme 2"}

	r := NewRegistry(theme1, theme2)

	// Get existing theme
	got, ok := r.GetTheme("theme2")
	if !ok || got == nil {
		t.Fatal("GetTheme(theme2) should return theme and true")
	}
	if got.ID() != "theme2" {
		t.Errorf("GetTheme(theme2).ID() = %q, want theme2", got.ID())
	}

	// Get nonexistent theme
	notFound, ok := r.GetTheme("nonexistent")
	if ok || notFound != nil {
		t.Error("GetTheme(nonexistent) should return nil, false")
	}
}

func TestRegistryDelegatedMethods(t *testing.T) {
	t.Parallel()
	theme1 := &mockTheme{id: "theme1", displayName: "Theme 1"}

	r := NewRegistry(theme1)

	// Test delegated methods
	if r.DisplayName() != "Theme 1" {
		t.Errorf("DisplayName() = %q, want %q", r.DisplayName(), "Theme 1")
	}
	if !r.IsDark() {
		t.Error("IsDark() should return true for mockTheme")
	}
	if r.TextPrimary().Hex() != "#f8f8f2" {
		t.Errorf("TextPrimary() = %q, want #f8f8f2", r.TextPrimary().Hex())
	}
	if r.Accent().Hex() != "#bd93f9" {
		t.Errorf("Accent() = %q, want #bd93f9", r.Accent().Hex())
	}
}

func TestRegistrySetThemeDirect(t *testing.T) {
	t.Parallel()
	theme1 := &mockTheme{id: "theme1", displayName: "Theme 1"}
	theme2 := &mockTheme{id: "theme2", displayName: "Theme 2"}

	r := NewRegistry(theme1, theme2)

	// SetTheme with a registered theme
	if !r.SetTheme(theme2) {
		t.Error("SetTheme(theme2) should return true for registered theme")
	}
	if r.ID() != "theme2" {
		t.Errorf("ID() after SetTheme = %q, want theme2", r.ID())
	}

	// SetTheme with nil
	if r.SetTheme(nil) {
		t.Error("SetTheme(nil) should return false")
	}

	// SetTheme with unregistered theme
	theme3 := &mockTheme{id: "theme3", displayName: "Theme 3"}
	if r.SetTheme(theme3) {
		t.Error("SetTheme(unregistered) should return false")
	}
}

func TestRegistryUnregisterCurrent(t *testing.T) {
	t.Parallel()
	theme1 := &mockTheme{id: "theme1", displayName: "Theme 1"}
	theme2 := &mockTheme{id: "theme2", displayName: "Theme 2"}

	r := NewRegistry(theme1, theme2)
	r.SetThemeID("theme1")

	// Unregister current theme
	r.Unregister(theme1)
	if r.Count() != 1 {
		t.Errorf("Count() after unregister = %d, want 1", r.Count())
	}
	// Current theme should switch to remaining theme
	if r.ID() != "theme2" {
		t.Errorf("After unregistering current, ID() = %q, want theme2", r.ID())
	}
}

func TestRegistryEmptyNavigation(t *testing.T) {
	t.Parallel()
	theme1 := &mockTheme{id: "theme1", displayName: "Theme 1"}

	r := NewRegistry(theme1)
	r.UnregisterAll()

	// Navigation on empty registry
	r.NextTheme()     // Should not panic
	r.PreviousTheme() // Should not panic

	// ID on empty registry
	if r.ID() != "" {
		t.Errorf("Empty registry ID() = %q, want empty", r.ID())
	}
}

// Benchmarks

func BenchmarkRegistrySetTheme(b *testing.B) {
	themes := make([]*mockTheme, 100)
	for i := range themes {
		themes[i] = &mockTheme{
			id:          "theme" + string(rune('a'+i%26)) + string(rune('0'+i/26)),
			displayName: "Theme " + string(rune('A'+i%26)),
		}
	}

	// Convert to Theme interface slice
	themeSlice := make([]Theme, len(themes))
	for i, t := range themes {
		themeSlice[i] = t
	}

	r := NewRegistry(themes[0])
	for _, t := range themeSlice[1:] {
		r.Register(t)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = r.SetTheme(themes[i%len(themes)])
	}
}

func BenchmarkRegistryGetTheme(b *testing.B) {
	themes := make([]*mockTheme, 100)
	for i := range themes {
		themes[i] = &mockTheme{
			id:          "theme" + string(rune('a'+i%26)) + string(rune('0'+i/26)),
			displayName: "Theme " + string(rune('A'+i%26)),
		}
	}

	r := NewRegistry(themes[0])
	for _, t := range themes[1:] {
		r.Register(t)
	}

	ids := r.ThemeIDs()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.GetTheme(ids[i%len(ids)])
	}
}

func BenchmarkRegistryNavigation(b *testing.B) {
	themes := make([]*mockTheme, 50)
	for i := range themes {
		themes[i] = &mockTheme{
			id:          "theme" + string(rune('a'+i%26)) + string(rune('0'+i/26)),
			displayName: "Theme " + string(rune('A'+i%26)),
		}
	}

	r := NewRegistry(themes[0])
	for _, t := range themes[1:] {
		r.Register(t)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			r.NextTheme()
		} else {
			r.PreviousTheme()
		}
	}
}

// Fuzz tests

func FuzzThemeIDLookup(f *testing.F) {
	// Add seed corpus
	f.Add("dracula")
	f.Add("nord")
	f.Add("gruvbox_dark")
	f.Add("nonexistent")
	f.Add("")
	f.Add("theme-with-dashes")
	f.Add("theme_with_underscores")
	f.Add("UPPERCASE")
	f.Add("MixedCase123")
	f.Add("123numeric")
	f.Add("special!@#$%^&*()")

	// Create a registry with some themes
	themes := []*mockTheme{
		{id: "dracula", displayName: "Dracula"},
		{id: "nord", displayName: "Nord"},
		{id: "gruvbox_dark", displayName: "Gruvbox Dark"},
		{id: "tokyo_night", displayName: "Tokyo Night"},
		{id: "catppuccin_mocha", displayName: "Catppuccin Mocha"},
	}

	r := NewRegistry(themes[0])
	for _, t := range themes[1:] {
		r.Register(t)
	}

	f.Fuzz(func(t *testing.T, id string) {
		// GetTheme should never panic
		theme, found := r.GetTheme(id)

		// If found, verify the ID matches
		if found && theme != nil {
			if theme.ID() != id {
				t.Errorf("GetTheme(%q) returned theme with ID %q", id, theme.ID())
			}
		}

		// SetThemeID should never panic
		_ = r.SetThemeID(id)

		// GetCurrentTheme should always return a valid theme or nil
		current := r.GetCurrentTheme()
		if current != nil {
			_ = current.ID()
			_ = current.DisplayName()
		}
	})
}
