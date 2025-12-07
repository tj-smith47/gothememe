package gothememe

// ThemeBuilder provides a fluent interface for constructing custom themes.
// All setter methods return the builder to allow method chaining.
//
// Example:
//
//	theme := gothememe.NewThemeBuilder("my-theme", "My Custom Theme").
//	    WithBackground(gothememe.Hex("#1a1a2e")).
//	    WithTextPrimary(gothememe.Hex("#e4e4e4")).
//	    WithAccent(gothememe.Hex("#e94560")).
//	    Build()
type ThemeBuilder struct {
	theme *BaseTheme
}

// NewThemeBuilder creates a new ThemeBuilder with the given ID and display name.
// The ID should be lowercase, alphanumeric with underscores.
func NewThemeBuilder(id, displayName string) *ThemeBuilder {
	return &ThemeBuilder{
		theme: &BaseTheme{
			id:          id,
			displayName: displayName,
		},
	}
}

//go:generate go run ./internal/cmd/genbuilder -output builder_gen.go

// Build finalizes the theme and returns it as a Theme interface.
// It applies default values for any unset colors.
func (b *ThemeBuilder) Build() Theme {
	b.deriveMissingColors()
	return b.theme
}

// deriveMissingColors fills in unset colors based on the colors that are set.
func (b *ThemeBuilder) deriveMissingColors() {
	t := b.theme

	// Detect dark mode if not explicitly set
	if !t.background.IsEmpty() && !t.isDark {
		t.isDark = t.background.IsDark()
	}

	// Derive background variants
	if t.backgroundSecondary.IsEmpty() && !t.background.IsEmpty() {
		if t.isDark {
			t.backgroundSecondary = t.background.Lighten(0.03)
		} else {
			t.backgroundSecondary = t.background.Darken(0.03)
		}
	}
	if t.surface.IsEmpty() && !t.background.IsEmpty() {
		if t.isDark {
			t.surface = t.background.Lighten(0.05)
		} else {
			t.surface = t.background.Darken(0.02)
		}
	}
	if t.surfaceSecondary.IsEmpty() && !t.surface.IsEmpty() {
		if t.isDark {
			t.surfaceSecondary = t.surface.Lighten(0.03)
		} else {
			t.surfaceSecondary = t.surface.Darken(0.02)
		}
	}

	// Derive text variants
	if t.textSecondary.IsEmpty() && !t.textPrimary.IsEmpty() {
		t.textSecondary = t.textPrimary.WithAlpha(0.7)
	}
	if t.textMuted.IsEmpty() && !t.textPrimary.IsEmpty() {
		t.textMuted = t.textPrimary.WithAlpha(0.5)
	}
	if t.textInverted.IsEmpty() && !t.background.IsEmpty() {
		t.textInverted = t.background
	}

	// Derive accent variants
	if t.accentSecondary.IsEmpty() && !t.accent.IsEmpty() {
		h, s, l := t.accent.HSLValues()
		t.accentSecondary = HSL(h+30, s, l) // Analogous color
	}
	if t.brand.IsEmpty() && !t.accent.IsEmpty() {
		t.brand = t.accent
	}

	// Derive border variants
	if t.border.IsEmpty() && !t.textPrimary.IsEmpty() {
		t.border = t.textPrimary.WithAlpha(0.2)
	}
	if t.borderSubtle.IsEmpty() && !t.border.IsEmpty() {
		t.borderSubtle = t.border.WithAlpha(0.1)
	}
	if t.borderStrong.IsEmpty() && !t.border.IsEmpty() {
		t.borderStrong = t.border.WithAlpha(0.4)
	}

	// Derive semantic colors from ANSI colors or defaults
	t.success = deriveSemanticColor(t.success, t.green, Hex("#22c55e"))
	t.warning = deriveSemanticColor(t.warning, t.yellow, Hex("#eab308"))
	t.errorColor = deriveSemanticColor(t.errorColor, t.red, Hex("#ef4444"))
	t.info = deriveSemanticColor(t.info, t.blue, Hex("#3b82f6"))

	// Derive code colors from theme colors
	if t.codeBackground.IsEmpty() && !t.background.IsEmpty() {
		if t.isDark {
			t.codeBackground = t.background.Darken(0.02)
		} else {
			t.codeBackground = t.background.Darken(0.05)
		}
	}
	if t.codeText.IsEmpty() && !t.textPrimary.IsEmpty() {
		t.codeText = t.textPrimary
	}
	if t.codeComment.IsEmpty() && !t.textMuted.IsEmpty() {
		t.codeComment = t.textMuted
	}
	if t.codeKeyword.IsEmpty() && !t.purple.IsEmpty() {
		t.codeKeyword = t.purple
	}
	if t.codeString.IsEmpty() && !t.green.IsEmpty() {
		t.codeString = t.green
	}
	if t.codeNumber.IsEmpty() && !t.yellow.IsEmpty() {
		t.codeNumber = t.yellow
	}
	if t.codeFunction.IsEmpty() && !t.blue.IsEmpty() {
		t.codeFunction = t.blue
	}
	if t.codeOperator.IsEmpty() && !t.cyan.IsEmpty() {
		t.codeOperator = t.cyan
	}
	if t.codePunctuation.IsEmpty() && !t.textSecondary.IsEmpty() {
		t.codePunctuation = t.textSecondary
	}
	if t.codeVariable.IsEmpty() && !t.textPrimary.IsEmpty() {
		t.codeVariable = t.textPrimary
	}
	if t.codeConstant.IsEmpty() && !t.yellow.IsEmpty() {
		t.codeConstant = t.yellow
	}
	if t.codeType.IsEmpty() && !t.cyan.IsEmpty() {
		t.codeType = t.cyan
	}
}

// Palette represents a minimal set of colors that can be used to generate
// a complete theme through automatic color derivation.
type Palette struct {
	Background Color
	Foreground Color
	Accent     Color
	// ANSI colors (all optional - will be derived if not set)
	Black        Color
	Red          Color
	Green        Color
	Yellow       Color
	Blue         Color
	Purple       Color
	Cyan         Color
	White        Color
	BrightBlack  Color
	BrightRed    Color
	BrightGreen  Color
	BrightYellow Color
	BrightBlue   Color
	BrightPurple Color
	BrightCyan   Color
	BrightWhite  Color
}

// ansiColorEntry holds a palette color and its corresponding builder setter.
type ansiColorEntry struct {
	color  Color
	setter func(*ThemeBuilder, Color) *ThemeBuilder
}

// getANSIColorEntries extracts ANSI color entries from a Palette for bulk application.
func getANSIColorEntries(p Palette) []ansiColorEntry {
	return []ansiColorEntry{
		{p.Black, (*ThemeBuilder).WithBlack},
		{p.Red, (*ThemeBuilder).WithRed},
		{p.Green, (*ThemeBuilder).WithGreen},
		{p.Yellow, (*ThemeBuilder).WithYellow},
		{p.Blue, (*ThemeBuilder).WithBlue},
		{p.Purple, (*ThemeBuilder).WithPurple},
		{p.Cyan, (*ThemeBuilder).WithCyan},
		{p.White, (*ThemeBuilder).WithWhite},
		{p.BrightBlack, (*ThemeBuilder).WithBrightBlack},
		{p.BrightRed, (*ThemeBuilder).WithBrightRed},
		{p.BrightGreen, (*ThemeBuilder).WithBrightGreen},
		{p.BrightYellow, (*ThemeBuilder).WithBrightYellow},
		{p.BrightBlue, (*ThemeBuilder).WithBrightBlue},
		{p.BrightPurple, (*ThemeBuilder).WithBrightPurple},
		{p.BrightCyan, (*ThemeBuilder).WithBrightCyan},
		{p.BrightWhite, (*ThemeBuilder).WithBrightWhite},
	}
}

// GenerateThemeFromPalette creates a complete theme from a minimal color palette.
// Missing ANSI colors are derived from the base colors.
func GenerateThemeFromPalette(id, displayName string, palette Palette) Theme {
	builder := NewThemeBuilder(id, displayName).
		WithBackground(palette.Background).
		WithTextPrimary(palette.Foreground).
		WithAccent(palette.Accent)

	// Set ANSI colors if provided
	for _, entry := range getANSIColorEntries(palette) {
		if !entry.color.IsEmpty() {
			entry.setter(builder, entry.color)
		}
	}

	return builder.Build()
}

// DeriveTheme creates a new theme based on an existing theme with color overrides.
// Any colors specified in the overrides map will replace the base theme's colors.
func DeriveTheme(base Theme, id, displayName string, overrides map[string]Color) Theme {
	builder := NewThemeBuilder(id, displayName).
		WithDescription(base.Description()).
		WithAuthor(base.Author()).
		WithLicense(base.License()).
		WithSource(base.Source()).
		WithIsDark(base.IsDark()).
		WithBackground(base.Background()).
		WithBackgroundSecondary(base.BackgroundSecondary()).
		WithSurface(base.Surface()).
		WithSurfaceSecondary(base.SurfaceSecondary()).
		WithTextPrimary(base.TextPrimary()).
		WithTextSecondary(base.TextSecondary()).
		WithTextMuted(base.TextMuted()).
		WithTextInverted(base.TextInverted()).
		WithAccent(base.Accent()).
		WithAccentSecondary(base.AccentSecondary()).
		WithBrand(base.Brand()).
		WithBorder(base.Border()).
		WithBorderSubtle(base.BorderSubtle()).
		WithBorderStrong(base.BorderStrong()).
		WithSuccess(base.Success()).
		WithWarning(base.Warning()).
		WithError(base.Error()).
		WithInfo(base.Info()).
		WithBlack(base.Black()).
		WithRed(base.Red()).
		WithGreen(base.Green()).
		WithYellow(base.Yellow()).
		WithBlue(base.Blue()).
		WithPurple(base.Purple()).
		WithCyan(base.Cyan()).
		WithWhite(base.White()).
		WithBrightBlack(base.BrightBlack()).
		WithBrightRed(base.BrightRed()).
		WithBrightGreen(base.BrightGreen()).
		WithBrightYellow(base.BrightYellow()).
		WithBrightBlue(base.BrightBlue()).
		WithBrightPurple(base.BrightPurple()).
		WithBrightCyan(base.BrightCyan()).
		WithBrightWhite(base.BrightWhite()).
		WithCodeBackground(base.CodeBackground()).
		WithCodeText(base.CodeText()).
		WithCodeComment(base.CodeComment()).
		WithCodeKeyword(base.CodeKeyword()).
		WithCodeString(base.CodeString()).
		WithCodeNumber(base.CodeNumber()).
		WithCodeFunction(base.CodeFunction()).
		WithCodeOperator(base.CodeOperator()).
		WithCodePunctuation(base.CodePunctuation()).
		WithCodeVariable(base.CodeVariable()).
		WithCodeConstant(base.CodeConstant()).
		WithCodeType(base.CodeType())

	// Apply overrides
	for name, color := range overrides {
		applyOverride(builder, name, color)
	}

	return builder.Build()
}

// overrideHandlers maps color names to their corresponding builder setter methods.
var overrideHandlers = map[string]func(*ThemeBuilder, Color) *ThemeBuilder{
	"background":           (*ThemeBuilder).WithBackground,
	"background_secondary": (*ThemeBuilder).WithBackgroundSecondary,
	"surface":              (*ThemeBuilder).WithSurface,
	"surface_secondary":    (*ThemeBuilder).WithSurfaceSecondary,
	"text_primary":         (*ThemeBuilder).WithTextPrimary,
	"text_secondary":       (*ThemeBuilder).WithTextSecondary,
	"text_muted":           (*ThemeBuilder).WithTextMuted,
	"text_inverted":        (*ThemeBuilder).WithTextInverted,
	"accent":               (*ThemeBuilder).WithAccent,
	"accent_secondary":     (*ThemeBuilder).WithAccentSecondary,
	"brand":                (*ThemeBuilder).WithBrand,
	"border":               (*ThemeBuilder).WithBorder,
	"border_subtle":        (*ThemeBuilder).WithBorderSubtle,
	"border_strong":        (*ThemeBuilder).WithBorderStrong,
	"black":                (*ThemeBuilder).WithBlack,
	"red":                  (*ThemeBuilder).WithRed,
	"green":                (*ThemeBuilder).WithGreen,
	"yellow":               (*ThemeBuilder).WithYellow,
	"blue":                 (*ThemeBuilder).WithBlue,
	"purple":               (*ThemeBuilder).WithPurple,
	"cyan":                 (*ThemeBuilder).WithCyan,
	"white":                (*ThemeBuilder).WithWhite,
	"bright_black":         (*ThemeBuilder).WithBrightBlack,
	"bright_red":           (*ThemeBuilder).WithBrightRed,
	"bright_green":         (*ThemeBuilder).WithBrightGreen,
	"bright_yellow":        (*ThemeBuilder).WithBrightYellow,
	"bright_blue":          (*ThemeBuilder).WithBrightBlue,
	"bright_purple":        (*ThemeBuilder).WithBrightPurple,
	"bright_cyan":          (*ThemeBuilder).WithBrightCyan,
	"bright_white":         (*ThemeBuilder).WithBrightWhite,
	"code_background":      (*ThemeBuilder).WithCodeBackground,
	"code_text":            (*ThemeBuilder).WithCodeText,
	"code_comment":         (*ThemeBuilder).WithCodeComment,
	"code_keyword":         (*ThemeBuilder).WithCodeKeyword,
	"code_string":          (*ThemeBuilder).WithCodeString,
	"code_number":          (*ThemeBuilder).WithCodeNumber,
	"code_function":        (*ThemeBuilder).WithCodeFunction,
	"code_operator":        (*ThemeBuilder).WithCodeOperator,
	"code_punctuation":     (*ThemeBuilder).WithCodePunctuation,
	"code_variable":        (*ThemeBuilder).WithCodeVariable,
	"code_constant":        (*ThemeBuilder).WithCodeConstant,
	"code_type":            (*ThemeBuilder).WithCodeType,
}

// applyOverride applies a single color override to the builder.
func applyOverride(b *ThemeBuilder, name string, c Color) {
	if handler, ok := overrideHandlers[name]; ok {
		handler(b, c)
	}
}

// deriveSemanticColor creates a SemanticColor from a base color if the existing
// color is empty. Uses the baseColor if available, otherwise falls back to fallback.
func deriveSemanticColor(existing SemanticColor, baseColor, fallback Color) SemanticColor {
	if !existing.Text.IsEmpty() {
		return existing
	}

	c := baseColor
	if c.IsEmpty() {
		c = fallback
	}

	return SemanticColor{
		Background: c.WithAlpha(0.1),
		Border:     c.WithAlpha(0.3),
		Text:       c,
	}
}
