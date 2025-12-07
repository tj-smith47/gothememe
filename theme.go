package gothememe

// Theme defines the interface that all themes must implement.
// It provides methods for metadata, mode detection, and accessing all theme colors.
//
// Themes are organized into several color categories:
//   - Background/Surface: Page and component backgrounds
//   - Text: Primary, secondary, muted, and inverted text colors
//   - Accent/Brand: Interactive and brand colors
//   - Border: Various border intensities
//   - Semantic: Success, warning, error, and info states
//   - ANSI: Terminal-compatible colors for code blocks
//   - Code: Syntax highlighting colors
type Theme interface {
	// Metadata methods

	// ID returns the unique identifier for the theme (e.g., "dracula", "nord").
	// IDs should be lowercase, alphanumeric with underscores.
	ID() string

	// DisplayName returns the human-readable name for the theme (e.g., "Dracula", "Nord").
	DisplayName() string

	// Description returns a brief description of the theme.
	Description() string

	// Author returns the original author or maintainer of the theme.
	Author() string

	// License returns the license under which the theme is distributed.
	License() string

	// Source returns a URL or reference to the original theme source.
	Source() string

	// Mode detection

	// IsDark returns true if this is a dark theme (dark background, light text).
	IsDark() bool

	// Background colors

	// Background returns the primary page background color.
	Background() Color

	// BackgroundSecondary returns a secondary/alternate background color.
	BackgroundSecondary() Color

	// Surface returns the color for elevated surfaces like cards and modals.
	Surface() Color

	// SurfaceSecondary returns a secondary surface color for nested components.
	SurfaceSecondary() Color

	// Text colors

	// TextPrimary returns the primary text color (highest contrast).
	TextPrimary() Color

	// TextSecondary returns a secondary text color (medium contrast).
	TextSecondary() Color

	// TextMuted returns a muted text color for disabled/placeholder content.
	TextMuted() Color

	// TextInverted returns text color for use on colored backgrounds.
	TextInverted() Color

	// Accent and Brand colors

	// Accent returns the primary accent color for interactive elements.
	Accent() Color

	// AccentSecondary returns a secondary accent color.
	AccentSecondary() Color

	// Brand returns the brand/logo color (often same as accent).
	Brand() Color

	// Border colors

	// Border returns the default border color.
	Border() Color

	// BorderSubtle returns a subtle/light border color.
	BorderSubtle() Color

	// BorderStrong returns an emphasized border color.
	BorderStrong() Color

	// Semantic colors

	// Success returns the semantic colors for success states.
	Success() SemanticColor

	// Warning returns the semantic colors for warning states.
	Warning() SemanticColor

	// Error returns the semantic colors for error states.
	Error() SemanticColor

	// Info returns the semantic colors for informational states.
	Info() SemanticColor

	// ANSI colors (standard 16-color palette)

	// Black returns the ANSI black color.
	Black() Color

	// Red returns the ANSI red color.
	Red() Color

	// Green returns the ANSI green color.
	Green() Color

	// Yellow returns the ANSI yellow color.
	Yellow() Color

	// Blue returns the ANSI blue color.
	Blue() Color

	// Purple returns the ANSI purple/magenta color.
	Purple() Color

	// Cyan returns the ANSI cyan color.
	Cyan() Color

	// White returns the ANSI white color.
	White() Color

	// BrightBlack returns the bright/bold ANSI black color.
	BrightBlack() Color

	// BrightRed returns the bright/bold ANSI red color.
	BrightRed() Color

	// BrightGreen returns the bright/bold ANSI green color.
	BrightGreen() Color

	// BrightYellow returns the bright/bold ANSI yellow color.
	BrightYellow() Color

	// BrightBlue returns the bright/bold ANSI blue color.
	BrightBlue() Color

	// BrightPurple returns the bright/bold ANSI purple/magenta color.
	BrightPurple() Color

	// BrightCyan returns the bright/bold ANSI cyan color.
	BrightCyan() Color

	// BrightWhite returns the bright/bold ANSI white color.
	BrightWhite() Color

	// Code/Syntax highlighting colors

	// CodeBackground returns the background color for code blocks.
	CodeBackground() Color

	// CodeText returns the default text color for code.
	CodeText() Color

	// CodeComment returns the color for comments in code.
	CodeComment() Color

	// CodeKeyword returns the color for keywords (if, else, return, etc.).
	CodeKeyword() Color

	// CodeString returns the color for string literals.
	CodeString() Color

	// CodeNumber returns the color for numeric literals.
	CodeNumber() Color

	// CodeFunction returns the color for function names.
	CodeFunction() Color

	// CodeOperator returns the color for operators (+, -, =, etc.).
	CodeOperator() Color

	// CodePunctuation returns the color for punctuation ({, }, (, ), etc.).
	CodePunctuation() Color

	// CodeVariable returns the color for variables.
	CodeVariable() Color

	// CodeConstant returns the color for constants.
	CodeConstant() Color

	// CodeType returns the color for type names.
	CodeType() Color
}

// BaseTheme provides a default implementation of the Theme interface.
// It can be embedded in custom themes to provide sensible defaults
// while allowing selective overrides.
type BaseTheme struct {
	id                  string
	displayName         string
	description         string
	author              string
	license             string
	source              string
	isDark              bool
	background          Color
	backgroundSecondary Color
	surface             Color
	surfaceSecondary    Color
	textPrimary         Color
	textSecondary       Color
	textMuted           Color
	textInverted        Color
	accent              Color
	accentSecondary     Color
	brand               Color
	border              Color
	borderSubtle        Color
	borderStrong        Color
	success             SemanticColor
	warning             SemanticColor
	errorColor          SemanticColor
	info                SemanticColor
	black               Color
	red                 Color
	green               Color
	yellow              Color
	blue                Color
	purple              Color
	cyan                Color
	white               Color
	brightBlack         Color
	brightRed           Color
	brightGreen         Color
	brightYellow        Color
	brightBlue          Color
	brightPurple        Color
	brightCyan          Color
	brightWhite         Color
	codeBackground      Color
	codeText            Color
	codeComment         Color
	codeKeyword         Color
	codeString          Color
	codeNumber          Color
	codeFunction        Color
	codeOperator        Color
	codePunctuation     Color
	codeVariable        Color
	codeConstant        Color
	codeType            Color
}

// ID implements [Theme.ID] and returns the unique lowercase identifier.
func (t *BaseTheme) ID() string { return t.id }

// DisplayName implements [Theme.DisplayName] and returns the human-readable name.
func (t *BaseTheme) DisplayName() string { return t.displayName }

// Description implements [Theme.Description] and returns a brief theme description.
func (t *BaseTheme) Description() string { return t.description }

// Author implements [Theme.Author] and returns the theme author.
func (t *BaseTheme) Author() string { return t.author }

// License implements [Theme.License] and returns the theme license.
func (t *BaseTheme) License() string { return t.license }

// Source implements [Theme.Source] and returns a reference to the original theme.
func (t *BaseTheme) Source() string { return t.source }

// IsDark implements [Theme.IsDark] and returns true for dark themes.
func (t *BaseTheme) IsDark() bool { return t.isDark }

// Background implements [Theme.Background] and returns the primary page background color.
func (t *BaseTheme) Background() Color { return t.background }

// BackgroundSecondary implements [Theme.BackgroundSecondary] and returns the alternate background.
func (t *BaseTheme) BackgroundSecondary() Color { return t.backgroundSecondary }

// Surface implements [Theme.Surface] and returns the elevated surface color.
func (t *BaseTheme) Surface() Color { return t.surface }

// SurfaceSecondary implements [Theme.SurfaceSecondary] and returns the nested surface color.
func (t *BaseTheme) SurfaceSecondary() Color { return t.surfaceSecondary }

// TextPrimary implements [Theme.TextPrimary] and returns the primary text color.
func (t *BaseTheme) TextPrimary() Color { return t.textPrimary }

// TextSecondary implements [Theme.TextSecondary] and returns the secondary text color.
func (t *BaseTheme) TextSecondary() Color { return t.textSecondary }

// TextMuted implements [Theme.TextMuted] and returns the muted text color.
func (t *BaseTheme) TextMuted() Color { return t.textMuted }

// TextInverted implements [Theme.TextInverted] and returns text for colored backgrounds.
func (t *BaseTheme) TextInverted() Color { return t.textInverted }

// Accent implements [Theme.Accent] and returns the primary accent color.
func (t *BaseTheme) Accent() Color { return t.accent }

// AccentSecondary implements [Theme.AccentSecondary] and returns the secondary accent.
func (t *BaseTheme) AccentSecondary() Color { return t.accentSecondary }

// Brand implements [Theme.Brand] and returns the brand/logo color.
func (t *BaseTheme) Brand() Color { return t.brand }

// Border implements [Theme.Border] and returns the default border color.
func (t *BaseTheme) Border() Color { return t.border }

// BorderSubtle implements [Theme.BorderSubtle] and returns the subtle border color.
func (t *BaseTheme) BorderSubtle() Color { return t.borderSubtle }

// BorderStrong implements [Theme.BorderStrong] and returns the emphasized border color.
func (t *BaseTheme) BorderStrong() Color { return t.borderStrong }

// Success implements [Theme.Success] and returns semantic colors for success states.
func (t *BaseTheme) Success() SemanticColor { return t.success }

// Warning implements [Theme.Warning] and returns semantic colors for warning states.
func (t *BaseTheme) Warning() SemanticColor { return t.warning }

// Error implements [Theme.Error] and returns semantic colors for error states.
func (t *BaseTheme) Error() SemanticColor { return t.errorColor }

// Info implements [Theme.Info] and returns semantic colors for informational states.
func (t *BaseTheme) Info() SemanticColor { return t.info }

// Black implements [Theme.Black] and returns the ANSI black color.
func (t *BaseTheme) Black() Color { return t.black }

// Red implements [Theme.Red] and returns the ANSI red color.
func (t *BaseTheme) Red() Color { return t.red }

// Green implements [Theme.Green] and returns the ANSI green color.
func (t *BaseTheme) Green() Color { return t.green }

// Yellow implements [Theme.Yellow] and returns the ANSI yellow color.
func (t *BaseTheme) Yellow() Color { return t.yellow }

// Blue implements [Theme.Blue] and returns the ANSI blue color.
func (t *BaseTheme) Blue() Color { return t.blue }

// Purple implements [Theme.Purple] and returns the ANSI purple/magenta color.
func (t *BaseTheme) Purple() Color { return t.purple }

// Cyan implements [Theme.Cyan] and returns the ANSI cyan color.
func (t *BaseTheme) Cyan() Color { return t.cyan }

// White implements [Theme.White] and returns the ANSI white color.
func (t *BaseTheme) White() Color { return t.white }

// BrightBlack implements [Theme.BrightBlack] and returns the bright ANSI black.
func (t *BaseTheme) BrightBlack() Color { return t.brightBlack }

// BrightRed implements [Theme.BrightRed] and returns the bright ANSI red.
func (t *BaseTheme) BrightRed() Color { return t.brightRed }

// BrightGreen implements [Theme.BrightGreen] and returns the bright ANSI green.
func (t *BaseTheme) BrightGreen() Color { return t.brightGreen }

// BrightYellow implements [Theme.BrightYellow] and returns the bright ANSI yellow.
func (t *BaseTheme) BrightYellow() Color { return t.brightYellow }

// BrightBlue implements [Theme.BrightBlue] and returns the bright ANSI blue.
func (t *BaseTheme) BrightBlue() Color { return t.brightBlue }

// BrightPurple implements [Theme.BrightPurple] and returns the bright ANSI purple.
func (t *BaseTheme) BrightPurple() Color { return t.brightPurple }

// BrightCyan implements [Theme.BrightCyan] and returns the bright ANSI cyan.
func (t *BaseTheme) BrightCyan() Color { return t.brightCyan }

// BrightWhite implements [Theme.BrightWhite] and returns the bright ANSI white.
func (t *BaseTheme) BrightWhite() Color { return t.brightWhite }

// CodeBackground implements [Theme.CodeBackground] and returns the code block background.
func (t *BaseTheme) CodeBackground() Color { return t.codeBackground }

// CodeText implements [Theme.CodeText] and returns the default code text color.
func (t *BaseTheme) CodeText() Color { return t.codeText }

// CodeComment implements [Theme.CodeComment] and returns the comment color.
func (t *BaseTheme) CodeComment() Color { return t.codeComment }

// CodeKeyword implements [Theme.CodeKeyword] and returns the keyword color.
func (t *BaseTheme) CodeKeyword() Color { return t.codeKeyword }

// CodeString implements [Theme.CodeString] and returns the string literal color.
func (t *BaseTheme) CodeString() Color { return t.codeString }

// CodeNumber implements [Theme.CodeNumber] and returns the numeric literal color.
func (t *BaseTheme) CodeNumber() Color { return t.codeNumber }

// CodeFunction implements [Theme.CodeFunction] and returns the function name color.
func (t *BaseTheme) CodeFunction() Color { return t.codeFunction }

// CodeOperator implements [Theme.CodeOperator] and returns the operator color.
func (t *BaseTheme) CodeOperator() Color { return t.codeOperator }

// CodePunctuation implements [Theme.CodePunctuation] and returns the punctuation color.
func (t *BaseTheme) CodePunctuation() Color { return t.codePunctuation }

// CodeVariable implements [Theme.CodeVariable] and returns the variable color.
func (t *BaseTheme) CodeVariable() Color { return t.codeVariable }

// CodeConstant implements [Theme.CodeConstant] and returns the constant color.
func (t *BaseTheme) CodeConstant() Color { return t.codeConstant }

// CodeType implements [Theme.CodeType] and returns the type name color.
func (t *BaseTheme) CodeType() Color { return t.codeType }

// Ensure BaseTheme implements Theme interface.
var _ Theme = (*BaseTheme)(nil)
