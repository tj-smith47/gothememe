package gothememe

import (
	"maps"
	"slices"
	"sync"
)

// Registry manages a collection of themes and tracks the currently active theme.
// It is safe for concurrent use from multiple goroutines.
type Registry struct {
	mu      sync.RWMutex
	themes  map[string]Theme
	current Theme
	sorted  []string // sorted theme IDs for navigation
}

// NewRegistry creates a new Registry with the specified default theme and
// additional themes. The first theme becomes the current (active) theme.
//
// Example:
//
//	registry := NewRegistry(ThemeDracula, ThemeNord, ThemeGruvboxDark)
func NewRegistry(defaultTheme Theme, themes ...Theme) *Registry {
	r := &Registry{
		themes:  make(map[string]Theme),
		current: defaultTheme,
	}

	// Register the default theme
	if defaultTheme != nil {
		r.themes[defaultTheme.ID()] = defaultTheme
	}

	// Register additional themes
	for _, t := range themes {
		if t != nil {
			r.themes[t.ID()] = t
		}
	}

	r.updateSorted()
	return r
}

// Register adds one or more themes to the registry.
// If a theme with the same ID already exists, it will be replaced.
func (r *Registry) Register(themes ...Theme) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, t := range themes {
		if t != nil {
			r.themes[t.ID()] = t
		}
	}
	r.updateSorted()
}

// Unregister removes one or more themes from the registry.
// If the current theme is removed, the registry falls back to the first
// available theme in alphabetical order.
func (r *Registry) Unregister(themes ...Theme) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, t := range themes {
		if t != nil {
			delete(r.themes, t.ID())
		}
	}
	r.updateSorted()

	// Check if current theme was removed
	if r.current != nil {
		if _, exists := r.themes[r.current.ID()]; !exists {
			// Fall back to first available theme
			if len(r.sorted) > 0 {
				r.current = r.themes[r.sorted[0]]
			} else {
				r.current = nil
			}
		}
	}
}

// UnregisterAll removes all themes from the registry.
func (r *Registry) UnregisterAll() {
	r.mu.Lock()
	defer r.mu.Unlock()

	clear(r.themes)
	r.sorted = nil
	r.current = nil
}

// SetTheme sets the current theme to the specified theme.
// The theme must already be registered in the registry.
// Returns false if the theme is not found in the registry.
func (r *Registry) SetTheme(theme Theme) bool {
	if theme == nil {
		return false
	}
	return r.SetThemeID(theme.ID())
}

// SetThemeID sets the current theme by its ID.
// Returns false if no theme with the given ID is registered.
func (r *Registry) SetThemeID(id string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	theme, exists := r.themes[id]
	if !exists {
		return false
	}
	r.current = theme
	return true
}

// GetTheme retrieves a theme by its ID.
// Returns the theme and true if found, nil and false otherwise.
func (r *Registry) GetTheme(id string) (Theme, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	theme, exists := r.themes[id]
	return theme, exists
}

// GetCurrentTheme returns the currently active theme.
// May return nil if no themes are registered.
func (r *Registry) GetCurrentTheme() Theme {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.current
}

// Themes returns all registered themes in alphabetical order by ID.
func (r *Registry) Themes() []Theme {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]Theme, 0, len(r.sorted))
	for _, id := range r.sorted {
		result = append(result, r.themes[id])
	}
	return result
}

// ThemeIDs returns all registered theme IDs in alphabetical order.
func (r *Registry) ThemeIDs() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]string, len(r.sorted))
	copy(result, r.sorted)
	return result
}

// Count returns the number of registered themes.
func (r *Registry) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return len(r.themes)
}

// NextTheme switches to the next theme in alphabetical order.
// Wraps around to the first theme after the last.
func (r *Registry) NextTheme() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(r.sorted) == 0 || r.current == nil {
		return
	}

	currentIdx := r.findCurrentIndex()
	nextIdx := (currentIdx + 1) % len(r.sorted)
	r.current = r.themes[r.sorted[nextIdx]]
}

// PreviousTheme switches to the previous theme in alphabetical order.
// Wraps around to the last theme before the first.
func (r *Registry) PreviousTheme() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(r.sorted) == 0 || r.current == nil {
		return
	}

	currentIdx := r.findCurrentIndex()
	prevIdx := currentIdx - 1
	if prevIdx < 0 {
		prevIdx = len(r.sorted) - 1
	}
	r.current = r.themes[r.sorted[prevIdx]]
}

// findCurrentIndex returns the index of the current theme in the sorted list.
// Must be called with lock held.
func (r *Registry) findCurrentIndex() int {
	if r.current == nil {
		return 0
	}
	if idx := slices.Index(r.sorted, r.current.ID()); idx >= 0 {
		return idx
	}
	return 0
}

// updateSorted rebuilds the sorted ID list.
// Must be called with lock held.
func (r *Registry) updateSorted() {
	r.sorted = slices.Sorted(maps.Keys(r.themes))
}

// Convenience methods that delegate to the current theme.
// These methods return zero values if no theme is set.

// ID returns the current theme's ID.
func (r *Registry) ID() string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if r.current == nil {
		return ""
	}
	return r.current.ID()
}

// DisplayName returns the current theme's display name.
func (r *Registry) DisplayName() string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if r.current == nil {
		return ""
	}
	return r.current.DisplayName()
}

// IsDark returns whether the current theme is dark.
func (r *Registry) IsDark() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if r.current == nil {
		return false
	}
	return r.current.IsDark()
}

// Background returns the current theme's background color.
func (r *Registry) Background() Color {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if r.current == nil {
		return Color{}
	}
	return r.current.Background()
}

// TextPrimary returns the current theme's primary text color.
func (r *Registry) TextPrimary() Color {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if r.current == nil {
		return Color{}
	}
	return r.current.TextPrimary()
}

// Accent returns the current theme's accent color.
func (r *Registry) Accent() Color {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if r.current == nil {
		return Color{}
	}
	return r.current.Accent()
}
