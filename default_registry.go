package gothememe

// defaultRegistry is the package-level registry used by convenience functions.
var defaultRegistry *Registry

// NewDefaultRegistry initializes the default registry with all built-in themes.
// The default theme is set to Dracula. Call this function once at application startup.
//
// Example:
//
//	func main() {
//	    gothememe.NewDefaultRegistry()
//	    css := gothememe.CSS(gothememe.DefaultCSSOptions())
//	    // Use the CSS...
//	}
func NewDefaultRegistry() {
	themes := DefaultThemes()
	if len(themes) == 0 {
		defaultRegistry = NewRegistry(nil)
		return
	}

	// Find Dracula as default, or use first theme
	var defaultTheme Theme
	for _, t := range themes {
		if t.ID() == "dracula" {
			defaultTheme = t
			break
		}
	}
	if defaultTheme == nil {
		defaultTheme = themes[0]
	}

	defaultRegistry = NewRegistry(defaultTheme, themes...)
}

// GetDefaultRegistry returns the default registry.
// Returns nil if NewDefaultRegistry has not been called.
func GetDefaultRegistry() *Registry {
	return defaultRegistry
}

// SetTheme sets the current theme in the default registry.
// Returns false if the theme is not registered.
func SetTheme(theme Theme) bool {
	if defaultRegistry == nil {
		return false
	}
	return defaultRegistry.SetTheme(theme)
}

// SetThemeID sets the current theme by ID in the default registry.
// Returns false if no theme with the given ID is registered.
func SetThemeID(id string) bool {
	if defaultRegistry == nil {
		return false
	}
	return defaultRegistry.SetThemeID(id)
}

// GetTheme retrieves a theme by ID from the default registry.
func GetTheme(id string) (Theme, bool) {
	if defaultRegistry == nil {
		return nil, false
	}
	return defaultRegistry.GetTheme(id)
}

// GetCurrentTheme returns the currently active theme from the default registry.
func GetCurrentTheme() Theme {
	if defaultRegistry == nil {
		return nil
	}
	return defaultRegistry.GetCurrentTheme()
}

// Themes returns all themes from the default registry.
func Themes() []Theme {
	if defaultRegistry == nil {
		return nil
	}
	return defaultRegistry.Themes()
}

// ThemeIDs returns all theme IDs from the default registry.
func ThemeIDs() []string {
	if defaultRegistry == nil {
		return nil
	}
	return defaultRegistry.ThemeIDs()
}

// NextTheme switches to the next theme in the default registry.
func NextTheme() {
	if defaultRegistry != nil {
		defaultRegistry.NextTheme()
	}
}

// PreviousTheme switches to the previous theme in the default registry.
func PreviousTheme() {
	if defaultRegistry != nil {
		defaultRegistry.PreviousTheme()
	}
}

// Register adds themes to the default registry.
func Register(themes ...Theme) {
	if defaultRegistry != nil {
		defaultRegistry.Register(themes...)
	}
}

// Unregister removes themes from the default registry.
func Unregister(themes ...Theme) {
	if defaultRegistry != nil {
		defaultRegistry.Unregister(themes...)
	}
}

// ID returns the current theme's ID from the default registry.
func ID() string {
	if defaultRegistry == nil {
		return ""
	}
	return defaultRegistry.ID()
}

// DisplayName returns the current theme's display name from the default registry.
func DisplayName() string {
	if defaultRegistry == nil {
		return ""
	}
	return defaultRegistry.DisplayName()
}

// IsDark returns whether the current theme is dark.
func IsDark() bool {
	if defaultRegistry == nil {
		return false
	}
	return defaultRegistry.IsDark()
}

// Background returns the current theme's background color.
func Background() Color {
	if defaultRegistry == nil {
		return Color{}
	}
	return defaultRegistry.Background()
}

// TextPrimary returns the current theme's primary text color.
func TextPrimary() Color {
	if defaultRegistry == nil {
		return Color{}
	}
	return defaultRegistry.TextPrimary()
}

// Accent returns the current theme's accent color.
func Accent() Color {
	if defaultRegistry == nil {
		return Color{}
	}
	return defaultRegistry.Accent()
}

// CSS generates CSS for the current theme using default options.
// Returns an empty string if no theme is set.
func CSS(opts CSSOptions) string {
	theme := GetCurrentTheme()
	if theme == nil {
		return ""
	}
	return GenerateCSS(theme, opts)
}

// AllThemesCSS generates CSS for all registered themes.
// Each theme's CSS is wrapped in a [data-theme="ID"] selector.
func AllThemesCSS(opts CSSOptions) string {
	if defaultRegistry == nil {
		return ""
	}
	return GenerateAllThemesCSS(defaultRegistry.Themes(), opts)
}
