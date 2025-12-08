// Package themes provides the Dracula Pro theme implementation.
// Dracula Pro is a premium dark theme with a darker background variant.
package themes

import "github.com/tj-smith47/gothememe"

// themeDraculaProInstance is the Dracula Pro color theme singleton.
var themeDraculaProInstance gothememe.Theme = &themeDraculaPro{}

// themeDraculaPro implements the Dracula Pro theme.
type themeDraculaPro struct{}

func (t *themeDraculaPro) ID() string          { return "dracula_pro" }
func (t *themeDraculaPro) DisplayName() string { return "Dracula Pro" }
func (t *themeDraculaPro) Description() string { return "Dracula Pro premium dark theme" }
func (t *themeDraculaPro) Author() string      { return "Dracula Theme" }
func (t *themeDraculaPro) License() string     { return "MIT" }
func (t *themeDraculaPro) Source() string      { return "https://draculatheme.com/pro" }
func (t *themeDraculaPro) IsDark() bool        { return true }

// Background colors
func (t *themeDraculaPro) Background() gothememe.Color          { return gothememe.Hex("#1e1f29") }
func (t *themeDraculaPro) BackgroundSecondary() gothememe.Color { return gothememe.Hex("#343746") }
func (t *themeDraculaPro) Surface() gothememe.Color             { return gothememe.Hex("#282a36") }
func (t *themeDraculaPro) SurfaceSecondary() gothememe.Color    { return gothememe.Hex("#44475a") }

// Text colors
func (t *themeDraculaPro) TextPrimary() gothememe.Color   { return gothememe.Hex("#f8f8f2") }
func (t *themeDraculaPro) TextSecondary() gothememe.Color { return gothememe.Hex("#e0e0e0") }
func (t *themeDraculaPro) TextMuted() gothememe.Color     { return gothememe.Hex("#6272a4") }
func (t *themeDraculaPro) TextInverted() gothememe.Color  { return gothememe.Hex("#1e1f29") }

// Accent/Brand colors
func (t *themeDraculaPro) Accent() gothememe.Color          { return gothememe.Hex("#bd93f9") } // Purple
func (t *themeDraculaPro) AccentSecondary() gothememe.Color { return gothememe.Hex("#ff79c6") } // Pink
func (t *themeDraculaPro) Brand() gothememe.Color           { return gothememe.Hex("#bd93f9") }

// Border colors
func (t *themeDraculaPro) Border() gothememe.Color {
	return gothememe.Hex("#bd93f9").WithAlpha(0.55)
} // BorderAccent
func (t *themeDraculaPro) BorderSubtle() gothememe.Color {
	return gothememe.Hex("#bd93f9").WithAlpha(0.35)
}
func (t *themeDraculaPro) BorderStrong() gothememe.Color { return gothememe.Hex("#f8f8f2") }

// Semantic colors
func (t *themeDraculaPro) Success() gothememe.SemanticColor {
	return gothememe.SemanticColor{
		Background: gothememe.Hex("#50fa7b").WithAlpha(0.1),
		Border:     gothememe.Hex("#50fa7b").WithAlpha(0.3),
		Text:       gothememe.Hex("#50fa7b"),
	}
}

func (t *themeDraculaPro) Warning() gothememe.SemanticColor {
	return gothememe.SemanticColor{
		Background: gothememe.Hex("#ffb86c").WithAlpha(0.1),
		Border:     gothememe.Hex("#ffb86c").WithAlpha(0.3),
		Text:       gothememe.Hex("#ffb86c"),
	}
}

func (t *themeDraculaPro) Error() gothememe.SemanticColor {
	return gothememe.SemanticColor{
		Background: gothememe.Hex("#ff5555").WithAlpha(0.1),
		Border:     gothememe.Hex("#ff5555").WithAlpha(0.3),
		Text:       gothememe.Hex("#ff5555"),
	}
}

func (t *themeDraculaPro) Info() gothememe.SemanticColor {
	return gothememe.SemanticColor{
		Background: gothememe.Hex("#8be9fd").WithAlpha(0.1),
		Border:     gothememe.Hex("#8be9fd").WithAlpha(0.3),
		Text:       gothememe.Hex("#8be9fd"),
	}
}

// ANSI colors
func (t *themeDraculaPro) Black() gothememe.Color        { return gothememe.Hex("#1e1f29") }
func (t *themeDraculaPro) Red() gothememe.Color          { return gothememe.Hex("#ff5555") }
func (t *themeDraculaPro) Green() gothememe.Color        { return gothememe.Hex("#50fa7b") }
func (t *themeDraculaPro) Yellow() gothememe.Color       { return gothememe.Hex("#f1fa8c") }
func (t *themeDraculaPro) Blue() gothememe.Color         { return gothememe.Hex("#bd93f9") }
func (t *themeDraculaPro) Purple() gothememe.Color       { return gothememe.Hex("#ff79c6") }
func (t *themeDraculaPro) Cyan() gothememe.Color         { return gothememe.Hex("#8be9fd") }
func (t *themeDraculaPro) White() gothememe.Color        { return gothememe.Hex("#f8f8f2") }
func (t *themeDraculaPro) BrightBlack() gothememe.Color  { return gothememe.Hex("#6272a4") }
func (t *themeDraculaPro) BrightRed() gothememe.Color    { return gothememe.Hex("#ff6e6e") }
func (t *themeDraculaPro) BrightGreen() gothememe.Color  { return gothememe.Hex("#69ff94") }
func (t *themeDraculaPro) BrightYellow() gothememe.Color { return gothememe.Hex("#ffffa5") }
func (t *themeDraculaPro) BrightBlue() gothememe.Color   { return gothememe.Hex("#d6acff") }
func (t *themeDraculaPro) BrightPurple() gothememe.Color { return gothememe.Hex("#ff92df") }
func (t *themeDraculaPro) BrightCyan() gothememe.Color   { return gothememe.Hex("#a4ffff") }
func (t *themeDraculaPro) BrightWhite() gothememe.Color  { return gothememe.Hex("#ffffff") }

// Code/Syntax highlighting colors
func (t *themeDraculaPro) CodeBackground() gothememe.Color  { return gothememe.Hex("#1e1f29") }
func (t *themeDraculaPro) CodeText() gothememe.Color        { return gothememe.Hex("#f8f8f2") }
func (t *themeDraculaPro) CodeComment() gothememe.Color     { return gothememe.Hex("#6272a4") }
func (t *themeDraculaPro) CodeKeyword() gothememe.Color     { return gothememe.Hex("#ff79c6") }
func (t *themeDraculaPro) CodeString() gothememe.Color      { return gothememe.Hex("#50fa7b") }
func (t *themeDraculaPro) CodeNumber() gothememe.Color      { return gothememe.Hex("#f1fa8c") }
func (t *themeDraculaPro) CodeFunction() gothememe.Color    { return gothememe.Hex("#bd93f9") }
func (t *themeDraculaPro) CodeOperator() gothememe.Color    { return gothememe.Hex("#ff5555") }
func (t *themeDraculaPro) CodePunctuation() gothememe.Color { return gothememe.Hex("#f8f8f2") }
func (t *themeDraculaPro) CodeVariable() gothememe.Color    { return gothememe.Hex("#8be9fd") }
func (t *themeDraculaPro) CodeConstant() gothememe.Color    { return gothememe.Hex("#ffb86c") }
func (t *themeDraculaPro) CodeType() gothememe.Color        { return gothememe.Hex("#a4ffff") }
