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

// Metadata methods

// WithDescription sets the theme description.
func (b *ThemeBuilder) WithDescription(description string) *ThemeBuilder {
	b.theme.description = description
	return b
}

// WithAuthor sets the theme author.
func (b *ThemeBuilder) WithAuthor(author string) *ThemeBuilder {
	b.theme.author = author
	return b
}

// WithLicense sets the theme license.
func (b *ThemeBuilder) WithLicense(license string) *ThemeBuilder {
	b.theme.license = license
	return b
}

// WithSource sets the theme source URL.
func (b *ThemeBuilder) WithSource(source string) *ThemeBuilder {
	b.theme.source = source
	return b
}

// WithIsDark sets whether this is a dark theme.
func (b *ThemeBuilder) WithIsDark(isDark bool) *ThemeBuilder {
	b.theme.isDark = isDark
	return b
}

// Background colors

// WithBackground sets the primary background color.
func (b *ThemeBuilder) WithBackground(c Color) *ThemeBuilder {
	b.theme.background = c
	return b
}

// WithBackgroundSecondary sets the secondary background color.
func (b *ThemeBuilder) WithBackgroundSecondary(c Color) *ThemeBuilder {
	b.theme.backgroundSecondary = c
	return b
}

// WithSurface sets the surface color.
func (b *ThemeBuilder) WithSurface(c Color) *ThemeBuilder {
	b.theme.surface = c
	return b
}

// WithSurfaceSecondary sets the secondary surface color.
func (b *ThemeBuilder) WithSurfaceSecondary(c Color) *ThemeBuilder {
	b.theme.surfaceSecondary = c
	return b
}

// Text colors

// WithTextPrimary sets the primary text color.
func (b *ThemeBuilder) WithTextPrimary(c Color) *ThemeBuilder {
	b.theme.textPrimary = c
	return b
}

// WithTextSecondary sets the secondary text color.
func (b *ThemeBuilder) WithTextSecondary(c Color) *ThemeBuilder {
	b.theme.textSecondary = c
	return b
}

// WithTextMuted sets the muted text color.
func (b *ThemeBuilder) WithTextMuted(c Color) *ThemeBuilder {
	b.theme.textMuted = c
	return b
}

// WithTextInverted sets the inverted text color.
func (b *ThemeBuilder) WithTextInverted(c Color) *ThemeBuilder {
	b.theme.textInverted = c
	return b
}

// Accent/Brand colors

// WithAccent sets the primary accent color.
func (b *ThemeBuilder) WithAccent(c Color) *ThemeBuilder {
	b.theme.accent = c
	return b
}

// WithAccentSecondary sets the secondary accent color.
func (b *ThemeBuilder) WithAccentSecondary(c Color) *ThemeBuilder {
	b.theme.accentSecondary = c
	return b
}

// WithBrand sets the brand color.
func (b *ThemeBuilder) WithBrand(c Color) *ThemeBuilder {
	b.theme.brand = c
	return b
}

// Border colors

// WithBorder sets the default border color.
func (b *ThemeBuilder) WithBorder(c Color) *ThemeBuilder {
	b.theme.border = c
	return b
}

// WithBorderSubtle sets the subtle border color.
func (b *ThemeBuilder) WithBorderSubtle(c Color) *ThemeBuilder {
	b.theme.borderSubtle = c
	return b
}

// WithBorderStrong sets the strong border color.
func (b *ThemeBuilder) WithBorderStrong(c Color) *ThemeBuilder {
	b.theme.borderStrong = c
	return b
}

// Semantic colors

// WithSuccess sets the success semantic colors.
func (b *ThemeBuilder) WithSuccess(sc SemanticColor) *ThemeBuilder {
	b.theme.success = sc
	return b
}

// WithWarning sets the warning semantic colors.
func (b *ThemeBuilder) WithWarning(sc SemanticColor) *ThemeBuilder {
	b.theme.warning = sc
	return b
}

// WithError sets the error semantic colors.
func (b *ThemeBuilder) WithError(sc SemanticColor) *ThemeBuilder {
	b.theme.errorColor = sc
	return b
}

// WithInfo sets the info semantic colors.
func (b *ThemeBuilder) WithInfo(sc SemanticColor) *ThemeBuilder {
	b.theme.info = sc
	return b
}

// ANSI colors

// WithBlack sets the ANSI black color.
func (b *ThemeBuilder) WithBlack(c Color) *ThemeBuilder {
	b.theme.black = c
	return b
}

// WithRed sets the ANSI red color.
func (b *ThemeBuilder) WithRed(c Color) *ThemeBuilder {
	b.theme.red = c
	return b
}

// WithGreen sets the ANSI green color.
func (b *ThemeBuilder) WithGreen(c Color) *ThemeBuilder {
	b.theme.green = c
	return b
}

// WithYellow sets the ANSI yellow color.
func (b *ThemeBuilder) WithYellow(c Color) *ThemeBuilder {
	b.theme.yellow = c
	return b
}

// WithBlue sets the ANSI blue color.
func (b *ThemeBuilder) WithBlue(c Color) *ThemeBuilder {
	b.theme.blue = c
	return b
}

// WithPurple sets the ANSI purple color.
func (b *ThemeBuilder) WithPurple(c Color) *ThemeBuilder {
	b.theme.purple = c
	return b
}

// WithCyan sets the ANSI cyan color.
func (b *ThemeBuilder) WithCyan(c Color) *ThemeBuilder {
	b.theme.cyan = c
	return b
}

// WithWhite sets the ANSI white color.
func (b *ThemeBuilder) WithWhite(c Color) *ThemeBuilder {
	b.theme.white = c
	return b
}

// WithBrightBlack sets the bright black color.
func (b *ThemeBuilder) WithBrightBlack(c Color) *ThemeBuilder {
	b.theme.brightBlack = c
	return b
}

// WithBrightRed sets the bright red color.
func (b *ThemeBuilder) WithBrightRed(c Color) *ThemeBuilder {
	b.theme.brightRed = c
	return b
}

// WithBrightGreen sets the bright green color.
func (b *ThemeBuilder) WithBrightGreen(c Color) *ThemeBuilder {
	b.theme.brightGreen = c
	return b
}

// WithBrightYellow sets the bright yellow color.
func (b *ThemeBuilder) WithBrightYellow(c Color) *ThemeBuilder {
	b.theme.brightYellow = c
	return b
}

// WithBrightBlue sets the bright blue color.
func (b *ThemeBuilder) WithBrightBlue(c Color) *ThemeBuilder {
	b.theme.brightBlue = c
	return b
}

// WithBrightPurple sets the bright purple color.
func (b *ThemeBuilder) WithBrightPurple(c Color) *ThemeBuilder {
	b.theme.brightPurple = c
	return b
}

// WithBrightCyan sets the bright cyan color.
func (b *ThemeBuilder) WithBrightCyan(c Color) *ThemeBuilder {
	b.theme.brightCyan = c
	return b
}

// WithBrightWhite sets the bright white color.
func (b *ThemeBuilder) WithBrightWhite(c Color) *ThemeBuilder {
	b.theme.brightWhite = c
	return b
}

// Code/Syntax highlighting colors

// WithCodeBackground sets the code background color.
func (b *ThemeBuilder) WithCodeBackground(c Color) *ThemeBuilder {
	b.theme.codeBackground = c
	return b
}

// WithCodeText sets the default code text color.
func (b *ThemeBuilder) WithCodeText(c Color) *ThemeBuilder {
	b.theme.codeText = c
	return b
}

// WithCodeComment sets the code comment color.
func (b *ThemeBuilder) WithCodeComment(c Color) *ThemeBuilder {
	b.theme.codeComment = c
	return b
}

// WithCodeKeyword sets the code keyword color.
func (b *ThemeBuilder) WithCodeKeyword(c Color) *ThemeBuilder {
	b.theme.codeKeyword = c
	return b
}

// WithCodeString sets the code string color.
func (b *ThemeBuilder) WithCodeString(c Color) *ThemeBuilder {
	b.theme.codeString = c
	return b
}

// WithCodeNumber sets the code number color.
func (b *ThemeBuilder) WithCodeNumber(c Color) *ThemeBuilder {
	b.theme.codeNumber = c
	return b
}

// WithCodeFunction sets the code function color.
func (b *ThemeBuilder) WithCodeFunction(c Color) *ThemeBuilder {
	b.theme.codeFunction = c
	return b
}

// WithCodeOperator sets the code operator color.
func (b *ThemeBuilder) WithCodeOperator(c Color) *ThemeBuilder {
	b.theme.codeOperator = c
	return b
}

// WithCodePunctuation sets the code punctuation color.
func (b *ThemeBuilder) WithCodePunctuation(c Color) *ThemeBuilder {
	b.theme.codePunctuation = c
	return b
}

// WithCodeVariable sets the code variable color.
func (b *ThemeBuilder) WithCodeVariable(c Color) *ThemeBuilder {
	b.theme.codeVariable = c
	return b
}

// WithCodeConstant sets the code constant color.
func (b *ThemeBuilder) WithCodeConstant(c Color) *ThemeBuilder {
	b.theme.codeConstant = c
	return b
}

// WithCodeType sets the code type color.
func (b *ThemeBuilder) WithCodeType(c Color) *ThemeBuilder {
	b.theme.codeType = c
	return b
}

// Build finalizes the theme and returns it.
// Missing colors will be auto-derived from set colors where possible.
func (b *ThemeBuilder) Build() Theme {
	// Auto-derive missing colors
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
