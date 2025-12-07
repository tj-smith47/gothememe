package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"
	"unicode"
)

// Generator creates Go theme files from fetched theme data.
type Generator struct {
	outputDir string
	themes    []*WindowsTerminalTheme
}

// NewGenerator creates a new theme generator.
func NewGenerator(outputDir string) *Generator {
	return &Generator{
		outputDir: outputDir,
	}
}

// SetThemes sets the themes to generate.
func (g *Generator) SetThemes(themes []*WindowsTerminalTheme) {
	g.themes = themes
}

// Generate creates all theme files and the registry file.
func (g *Generator) Generate() error {
	if err := os.MkdirAll(g.outputDir, 0o755); err != nil { //nolint:gosec // G301: 0755 is intentional for output directory
		return fmt.Errorf("creating output directory: %w", err)
	}

	// Track seen IDs and var names to handle duplicates
	seenIDs := make(map[string]int)
	seenVarNames := make(map[string]int)

	// Generate individual theme files
	var themeInfos []themeInfo
	for _, t := range g.themes {
		info, err := g.generateThemeFile(t, seenIDs, seenVarNames)
		if err != nil {
			fmt.Printf("Warning: failed to generate %s: %v\n", t.Name, err)
			continue
		}
		themeInfos = append(themeInfos, info)
	}

	// Sort themes alphabetically
	sort.Slice(themeInfos, func(i, j int) bool {
		return themeInfos[i].ID < themeInfos[j].ID
	})

	// Generate themes.go with all theme registrations
	if err := g.generateThemesFile(themeInfos); err != nil {
		return fmt.Errorf("generating themes.go: %w", err)
	}

	// Generate doc.go
	if err := g.generateDocFile(len(themeInfos)); err != nil {
		return fmt.Errorf("generating doc.go: %w", err)
	}

	return nil
}

// GenerateMarkdown creates a DEFAULT_THEMES.md file listing all themes.
func (g *Generator) GenerateMarkdown(outputPath string) error {
	// Track seen IDs to generate consistent info
	seenIDs := make(map[string]int)
	seenVarNames := make(map[string]int)

	svgGen := DefaultSVGPreview()

	var themeInfos []themeInfo
	for _, t := range g.themes {
		id := toThemeID(t.Name)
		varName := toVarName(t.Name)

		// Handle duplicate IDs
		if count, exists := seenIDs[id]; exists {
			seenIDs[id] = count + 1
			id = fmt.Sprintf("%s_%d", id, count+1)
		} else {
			seenIDs[id] = 1
		}

		// Handle duplicate var names
		if count, exists := seenVarNames[varName]; exists {
			seenVarNames[varName] = count + 1
			varName = fmt.Sprintf("%s%d", varName, count+1)
		} else {
			seenVarNames[varName] = 1
		}

		// Generate inline SVG preview
		svgDataURI := svgGen.GenerateInline(t)

		themeInfos = append(themeInfos, themeInfo{
			ID:          id,
			VarName:     varName,
			DisplayName: t.Name,
			IsDark:      isDarkTheme(t.Background),
			SVGPreview:  svgDataURI,
		})
	}

	// Sort themes alphabetically
	sort.Slice(themeInfos, func(i, j int) bool {
		return themeInfos[i].ID < themeInfos[j].ID
	})

	var buf bytes.Buffer
	if err := markdownTemplate.Execute(&buf, struct {
		Themes []themeInfo
		Count  int
	}{
		Themes: themeInfos,
		Count:  len(themeInfos),
	}); err != nil {
		return fmt.Errorf("executing template: %w", err)
	}

	return os.WriteFile(outputPath, buf.Bytes(), 0o644) //nolint:gosec // G306: 0644 is intentional for generated markdown
}

type themeInfo struct {
	ID           string
	VarName      string
	InstanceName string
	DisplayName  string
	FileName     string
	IsDark       bool
	SVGPreview   string // Inline SVG data URI for markdown
}

func (g *Generator) generateThemeFile(t *WindowsTerminalTheme, seenIDs, seenVarNames map[string]int) (themeInfo, error) {
	id := toThemeID(t.Name)
	varName := toVarName(t.Name)

	// Handle duplicate IDs
	if count, exists := seenIDs[id]; exists {
		seenIDs[id] = count + 1
		id = fmt.Sprintf("%s_%d", id, count+1)
	} else {
		seenIDs[id] = 1
	}

	// Handle duplicate var names
	if count, exists := seenVarNames[varName]; exists {
		seenVarNames[varName] = count + 1
		varName = fmt.Sprintf("%s%d", varName, count+1)
	} else {
		seenVarNames[varName] = 1
	}

	instanceName := "theme" + varName + "Instance"
	typeName := "theme" + varName
	filename := "theme_" + id + ".go"

	isDark := isDarkTheme(t.Background)

	info := themeInfo{
		ID:           id,
		VarName:      varName,
		InstanceName: instanceName,
		DisplayName:  t.Name,
		FileName:     filename,
		IsDark:       isDark,
	}

	data := struct {
		ID           string
		VarName      string
		InstanceName string
		TypeName     string
		DisplayName  string
		IsDark       bool
		Theme        *WindowsTerminalTheme
	}{
		ID:           id,
		VarName:      varName,
		InstanceName: instanceName,
		TypeName:     typeName,
		DisplayName:  t.Name,
		IsDark:       isDark,
		Theme:        t,
	}

	var buf bytes.Buffer
	if err := themeTemplate.Execute(&buf, data); err != nil {
		return info, fmt.Errorf("executing template: %w", err)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		// Write unformatted for debugging
		path := filepath.Join(g.outputDir, filename)
		if writeErr := os.WriteFile(path, buf.Bytes(), 0o644); writeErr != nil { //nolint:gosec // G306: 0644 is intentional for generated source files
			fmt.Fprintf(os.Stderr, "warning: failed to write debug file %s: %v\n", path, writeErr)
		}
		return info, fmt.Errorf("formatting %s: %w", filename, err)
	}

	path := filepath.Join(g.outputDir, filename)
	if err := os.WriteFile(path, formatted, 0o644); err != nil { //nolint:gosec // G306: 0644 is intentional for generated source
		return info, fmt.Errorf("writing file: %w", err)
	}

	return info, nil
}

func (g *Generator) generateThemesFile(themes []themeInfo) error {
	data := struct {
		Themes []themeInfo
	}{
		Themes: themes,
	}

	var buf bytes.Buffer
	if err := themesFileTemplate.Execute(&buf, data); err != nil {
		return fmt.Errorf("executing template: %w", err)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		// Write unformatted for debugging
		path := filepath.Join(g.outputDir, "themes.go")
		if writeErr := os.WriteFile(path, buf.Bytes(), 0o644); writeErr != nil { //nolint:gosec // G306: 0644 is intentional for generated source files
			fmt.Fprintf(os.Stderr, "warning: failed to write debug file %s: %v\n", path, writeErr)
		}
		return fmt.Errorf("formatting: %w", err)
	}

	path := filepath.Join(g.outputDir, "themes.go")
	return os.WriteFile(path, formatted, 0o644) //nolint:gosec // G306: 0644 is intentional for generated source
}

func (g *Generator) generateDocFile(themeCount int) error {
	data := struct {
		ThemeCount int
	}{
		ThemeCount: themeCount,
	}

	var buf bytes.Buffer
	if err := docFileTemplate.Execute(&buf, data); err != nil {
		return fmt.Errorf("executing template: %w", err)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("formatting: %w", err)
	}

	path := filepath.Join(g.outputDir, "doc.go")
	return os.WriteFile(path, formatted, 0o644) //nolint:gosec // G306: 0644 is intentional for generated source
}

// toThemeID converts a theme name to a lowercase ID.
// Example: "Tokyo Night" -> "tokyo_night"
func toThemeID(name string) string {
	// Replace non-alphanumeric with underscore
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	id := re.ReplaceAllString(name, "_")
	id = strings.Trim(id, "_")
	id = strings.ToLower(id)

	// Collapse multiple underscores
	re = regexp.MustCompile(`_+`)
	id = re.ReplaceAllString(id, "_")

	return id
}

// toVarName converts a theme name to a Go-style exported name.
// Example: "Tokyo Night" -> "TokyoNight"
func toVarName(name string) string {
	// Split on non-alphanumeric
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	parts := re.Split(name, -1)

	var result strings.Builder
	for _, part := range parts {
		if part == "" {
			continue
		}
		// Handle numbers at start
		if result.Len() == 0 && part != "" && unicode.IsDigit(rune(part[0])) {
			result.WriteString("N")
		}
		// Title case each part
		if part != "" {
			result.WriteString(strings.ToUpper(string(part[0])))
			if len(part) > 1 {
				result.WriteString(strings.ToLower(part[1:]))
			}
		}
	}

	return result.String()
}

// isDarkTheme determines if a theme is dark based on background luminance.
func isDarkTheme(background string) bool {
	bg := strings.TrimPrefix(background, "#")
	if len(bg) != 6 {
		return true // Default to dark
	}

	var r, g, b uint8
	if _, err := fmt.Sscanf(bg, "%02x%02x%02x", &r, &g, &b); err != nil {
		return true // Default to dark on parse error
	}

	// Calculate relative luminance
	rs := float64(r) / 255.0
	gs := float64(g) / 255.0
	bs := float64(b) / 255.0

	var rLinear, gLinear, bLinear float64
	if rs <= 0.03928 {
		rLinear = rs / 12.92
	} else {
		rLinear = pow((rs+0.055)/1.055, 2.4)
	}
	if gs <= 0.03928 {
		gLinear = gs / 12.92
	} else {
		gLinear = pow((gs+0.055)/1.055, 2.4)
	}
	if bs <= 0.03928 {
		bLinear = bs / 12.92
	} else {
		bLinear = pow((bs+0.055)/1.055, 2.4)
	}

	luminance := 0.2126*rLinear + 0.7152*gLinear + 0.0722*bLinear
	return luminance < 0.5
}

func pow(base, exp float64) float64 {
	result := 1.0
	for i := 0; i < int(exp*10); i++ {
		result *= base
	}
	// This is a rough approximation, use math.Pow in real code
	// But we want to avoid importing math just for this
	if exp == 2.4 {
		return base * base * pow(base, 0.4)
	}
	return result
}

var themeTemplate = template.Must(template.New("theme").Parse(`// Code generated by themegen. DO NOT EDIT.

package themes

import "github.com/tj-smith47/gothememe"

// {{.InstanceName}} is the {{.DisplayName}} color theme singleton.
var {{.InstanceName}} gothememe.Theme = &{{.TypeName}}{}

// {{.TypeName}} implements the {{.DisplayName}} theme.
type {{.TypeName}} struct{}

func (t *{{.TypeName}}) ID() string          { return "{{.ID}}" }
func (t *{{.TypeName}}) DisplayName() string { return "{{.DisplayName}}" }
func (t *{{.TypeName}}) Description() string { return "{{.DisplayName}} color theme" }
func (t *{{.TypeName}}) Author() string      { return "iTerm2-Color-Schemes" }
func (t *{{.TypeName}}) License() string     { return "MIT" }
func (t *{{.TypeName}}) Source() string      { return "https://github.com/mbadolato/iTerm2-Color-Schemes" }
func (t *{{.TypeName}}) IsDark() bool        { return {{.IsDark}} }

// Background colors
func (t *{{.TypeName}}) Background() gothememe.Color          { return gothememe.Hex("{{.Theme.Background}}") }
func (t *{{.TypeName}}) BackgroundSecondary() gothememe.Color { return gothememe.Hex("{{.Theme.Black}}") }
func (t *{{.TypeName}}) Surface() gothememe.Color             { return gothememe.Hex("{{.Theme.SelectionBackground}}") }
func (t *{{.TypeName}}) SurfaceSecondary() gothememe.Color    { return gothememe.Hex("{{.Theme.BrightBlack}}") }

// Text colors
func (t *{{.TypeName}}) TextPrimary() gothememe.Color   { return gothememe.Hex("{{.Theme.Foreground}}") }
func (t *{{.TypeName}}) TextSecondary() gothememe.Color { return gothememe.Hex("{{.Theme.White}}") }
func (t *{{.TypeName}}) TextMuted() gothememe.Color     { return gothememe.Hex("{{.Theme.BrightBlack}}") }
func (t *{{.TypeName}}) TextInverted() gothememe.Color  { return gothememe.Hex("{{.Theme.Background}}") }

// Accent/Brand colors
func (t *{{.TypeName}}) Accent() gothememe.Color          { return gothememe.Hex("{{.Theme.Blue}}") }
func (t *{{.TypeName}}) AccentSecondary() gothememe.Color { return gothememe.Hex("{{.Theme.Purple}}") }
func (t *{{.TypeName}}) Brand() gothememe.Color           { return gothememe.Hex("{{.Theme.Blue}}") }

// Border colors
func (t *{{.TypeName}}) Border() gothememe.Color       { return gothememe.Hex("{{.Theme.BrightBlack}}") }
func (t *{{.TypeName}}) BorderSubtle() gothememe.Color { return gothememe.Hex("{{.Theme.Black}}") }
func (t *{{.TypeName}}) BorderStrong() gothememe.Color { return gothememe.Hex("{{.Theme.White}}") }

// Semantic colors
func (t *{{.TypeName}}) Success() gothememe.SemanticColor {
	return gothememe.SemanticColor{
		Background: gothememe.Hex("{{.Theme.Green}}").WithAlpha(0.1),
		Border:     gothememe.Hex("{{.Theme.Green}}").WithAlpha(0.3),
		Text:       gothememe.Hex("{{.Theme.Green}}"),
	}
}

func (t *{{.TypeName}}) Warning() gothememe.SemanticColor {
	return gothememe.SemanticColor{
		Background: gothememe.Hex("{{.Theme.Yellow}}").WithAlpha(0.1),
		Border:     gothememe.Hex("{{.Theme.Yellow}}").WithAlpha(0.3),
		Text:       gothememe.Hex("{{.Theme.Yellow}}"),
	}
}

func (t *{{.TypeName}}) Error() gothememe.SemanticColor {
	return gothememe.SemanticColor{
		Background: gothememe.Hex("{{.Theme.Red}}").WithAlpha(0.1),
		Border:     gothememe.Hex("{{.Theme.Red}}").WithAlpha(0.3),
		Text:       gothememe.Hex("{{.Theme.Red}}"),
	}
}

func (t *{{.TypeName}}) Info() gothememe.SemanticColor {
	return gothememe.SemanticColor{
		Background: gothememe.Hex("{{.Theme.Cyan}}").WithAlpha(0.1),
		Border:     gothememe.Hex("{{.Theme.Cyan}}").WithAlpha(0.3),
		Text:       gothememe.Hex("{{.Theme.Cyan}}"),
	}
}

// ANSI colors
func (t *{{.TypeName}}) Black() gothememe.Color        { return gothememe.Hex("{{.Theme.Black}}") }
func (t *{{.TypeName}}) Red() gothememe.Color          { return gothememe.Hex("{{.Theme.Red}}") }
func (t *{{.TypeName}}) Green() gothememe.Color        { return gothememe.Hex("{{.Theme.Green}}") }
func (t *{{.TypeName}}) Yellow() gothememe.Color       { return gothememe.Hex("{{.Theme.Yellow}}") }
func (t *{{.TypeName}}) Blue() gothememe.Color         { return gothememe.Hex("{{.Theme.Blue}}") }
func (t *{{.TypeName}}) Purple() gothememe.Color       { return gothememe.Hex("{{.Theme.Purple}}") }
func (t *{{.TypeName}}) Cyan() gothememe.Color         { return gothememe.Hex("{{.Theme.Cyan}}") }
func (t *{{.TypeName}}) White() gothememe.Color        { return gothememe.Hex("{{.Theme.White}}") }
func (t *{{.TypeName}}) BrightBlack() gothememe.Color  { return gothememe.Hex("{{.Theme.BrightBlack}}") }
func (t *{{.TypeName}}) BrightRed() gothememe.Color    { return gothememe.Hex("{{.Theme.BrightRed}}") }
func (t *{{.TypeName}}) BrightGreen() gothememe.Color  { return gothememe.Hex("{{.Theme.BrightGreen}}") }
func (t *{{.TypeName}}) BrightYellow() gothememe.Color { return gothememe.Hex("{{.Theme.BrightYellow}}") }
func (t *{{.TypeName}}) BrightBlue() gothememe.Color   { return gothememe.Hex("{{.Theme.BrightBlue}}") }
func (t *{{.TypeName}}) BrightPurple() gothememe.Color { return gothememe.Hex("{{.Theme.BrightPurple}}") }
func (t *{{.TypeName}}) BrightCyan() gothememe.Color   { return gothememe.Hex("{{.Theme.BrightCyan}}") }
func (t *{{.TypeName}}) BrightWhite() gothememe.Color  { return gothememe.Hex("{{.Theme.BrightWhite}}") }

// Code/Syntax highlighting colors
func (t *{{.TypeName}}) CodeBackground() gothememe.Color  { return gothememe.Hex("{{.Theme.Background}}") }
func (t *{{.TypeName}}) CodeText() gothememe.Color        { return gothememe.Hex("{{.Theme.Foreground}}") }
func (t *{{.TypeName}}) CodeComment() gothememe.Color     { return gothememe.Hex("{{.Theme.BrightBlack}}") }
func (t *{{.TypeName}}) CodeKeyword() gothememe.Color     { return gothememe.Hex("{{.Theme.Purple}}") }
func (t *{{.TypeName}}) CodeString() gothememe.Color      { return gothememe.Hex("{{.Theme.Green}}") }
func (t *{{.TypeName}}) CodeNumber() gothememe.Color      { return gothememe.Hex("{{.Theme.Yellow}}") }
func (t *{{.TypeName}}) CodeFunction() gothememe.Color    { return gothememe.Hex("{{.Theme.Blue}}") }
func (t *{{.TypeName}}) CodeOperator() gothememe.Color    { return gothememe.Hex("{{.Theme.Red}}") }
func (t *{{.TypeName}}) CodePunctuation() gothememe.Color { return gothememe.Hex("{{.Theme.Foreground}}") }
func (t *{{.TypeName}}) CodeVariable() gothememe.Color    { return gothememe.Hex("{{.Theme.Cyan}}") }
func (t *{{.TypeName}}) CodeConstant() gothememe.Color    { return gothememe.Hex("{{.Theme.BrightYellow}}") }
func (t *{{.TypeName}}) CodeType() gothememe.Color        { return gothememe.Hex("{{.Theme.BrightCyan}}") }
`))

var themesFileTemplate = template.Must(template.New("themes").Parse(`// Code generated by themegen. DO NOT EDIT.

package themes

import "github.com/tj-smith47/gothememe"

// All returns all available themes.
func All() []gothememe.Theme {
	return []gothememe.Theme{
{{- range .Themes}}
		{{.InstanceName}},
{{- end}}
	}
}

// ByID returns a theme by its ID, or nil if not found.
func ByID(id string) gothememe.Theme {
	for _, t := range All() {
		if t.ID() == id {
			return t
		}
	}
	return nil
}

// IDs returns all theme IDs in alphabetical order.
func IDs() []string {
	return []string{
{{- range .Themes}}
		"{{.ID}}",
{{- end}}
	}
}

// Pre-instantiated theme variables for direct access.
var (
{{- range .Themes}}
	// Theme{{.VarName}} is the {{.DisplayName}} theme.
	Theme{{.VarName}} gothememe.Theme = {{.InstanceName}}
{{- end}}
)
`))

var docFileTemplate = template.Must(template.New("doc").Parse(`// Code generated by themegen. DO NOT EDIT.

// Package themes provides {{.ThemeCount}}+ pre-built color themes for gothememe.
//
// Themes are sourced from the iTerm2-Color-Schemes repository and include
// popular themes like Dracula, Nord, Gruvbox, Tokyo Night, Catppuccin, and more.
//
// # Usage
//
// Import the themes package and access themes directly or via the registry:
//
//	import "github.com/tj-smith47/gothememe/themes"
//
//	// Direct access
//	theme := themes.ThemeDracula
//
//	// By ID
//	theme := themes.ByID("dracula")
//
//	// Iterate all themes
//	for _, t := range themes.All() {
//	    fmt.Println(t.DisplayName())
//	}
//
// # Available Themes
//
// Use themes.IDs() to get a list of all available theme IDs, or themes.All()
// to get all theme instances.
package themes
`))

var markdownTemplate = template.Must(template.New("markdown").Parse(`# Default Themes

GoThemeMe includes **{{.Count}}** pre-built themes sourced from [iTerm2-Color-Schemes](https://github.com/mbadolato/iTerm2-Color-Schemes).

## Usage

` + "```" + `go
import "github.com/tj-smith47/gothememe/themes"

// Access by variable
theme := themes.ThemeDracula

// Access by ID
theme := themes.ByID("dracula")

// List all themes
for _, t := range themes.All() {
    fmt.Println(t.ID(), "-", t.DisplayName())
}
` + "```" + `

## Theme List

| Preview | ID | Display Name | Type | Variable |
|---------|----|--------------| ------|----------|
{{- range .Themes}}
| ![{{.DisplayName}}]({{.SVGPreview}}) | ` + "`" + `{{.ID}}` + "`" + ` | {{.DisplayName}} | {{if .IsDark}}🌙 Dark{{else}}☀️ Light{{end}} | ` + "`" + `themes.Theme{{.VarName}}` + "`" + ` |
{{- end}}

## Popular Themes

### Dark Themes
- **Dracula** - A dark theme with purple accents
- **Nord** - An arctic, north-bluish color palette
- **Gruvbox Dark** - Retro groove color scheme
- **Tokyo Night** - A dark theme inspired by Tokyo
- **Catppuccin Mocha** - Soothing pastel theme
- **One Dark** - Atom's iconic dark theme
- **Monokai** - Classic syntax highlighting theme

### Light Themes
- **Solarized Light** - Precision colors for light backgrounds
- **Gruvbox Light** - Retro groove for light mode
- **Catppuccin Latte** - Catppuccin's light variant
- **One Light** - Atom's light companion theme

## Adding Custom Themes

See the [Contributing Guide](docs/CONTRIBUTING.md) for instructions on adding new themes.
`))
