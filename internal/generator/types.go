// Package generator provides theme generation from external sources.
package generator

// WindowsTerminalTheme represents the JSON structure from iTerm2-Color-Schemes.
type WindowsTerminalTheme struct {
	Name                string `json:"name"`
	Background          string `json:"background"`
	Foreground          string `json:"foreground"`
	CursorColor         string `json:"cursorColor"`
	SelectionBackground string `json:"selectionBackground"`
	Black               string `json:"black"`
	Red                 string `json:"red"`
	Green               string `json:"green"`
	Yellow              string `json:"yellow"`
	Blue                string `json:"blue"`
	Purple              string `json:"purple"`
	Cyan                string `json:"cyan"`
	White               string `json:"white"`
	BrightBlack         string `json:"brightBlack"`
	BrightRed           string `json:"brightRed"`
	BrightGreen         string `json:"brightGreen"`
	BrightYellow        string `json:"brightYellow"`
	BrightBlue          string `json:"brightBlue"`
	BrightPurple        string `json:"brightPurple"`
	BrightCyan          string `json:"brightCyan"`
	BrightWhite         string `json:"brightWhite"`
}

// ThemeMetadata contains additional metadata for generated themes.
type ThemeMetadata struct {
	ID          string
	DisplayName string
	Description string
	Author      string
	License     string
	Source      string
	IsDark      bool
}
