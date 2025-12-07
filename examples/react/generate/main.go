// Generate CSS and JSON for React application.
//
// Run from the examples/react directory:
//
//	go run ./generate
//
// Or with custom output:
//
//	go run ./generate ./custom-dir
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tj-smith47/gothememe"
	"github.com/tj-smith47/gothememe/themes"
)

// ThemeData represents the JSON structure for theme metadata.
type ThemeData struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	IsDark      bool   `json:"isDark"`
}

func main() {
	// Get output directory (default: public/ relative to current directory)
	outDir := "public"
	if len(os.Args) > 1 {
		outDir = os.Args[1]
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Select popular themes for the demo (keeps bundle size reasonable)
	ids := []string{
		"dracula", "nord", "gruvbox_dark", "atom_one_dark",
		"builtin_solarized_light", "builtin_solarized_dark",
		"catppuccin_mocha", "tokyonight", "github_dark", "monokai_pro",
	}
	var selectedThemes []gothememe.Theme
	var themeData []ThemeData

	for _, id := range ids {
		if t := themes.ByID(id); t != nil {
			selectedThemes = append(selectedThemes, t)
			themeData = append(themeData, ThemeData{
				ID:          t.ID(),
				DisplayName: t.DisplayName(),
				IsDark:      t.IsDark(),
			})
		}
	}

	// Generate CSS with data-theme attribute selector
	css := gothememe.GenerateAllThemesCSS(selectedThemes, gothememe.CSSOptions{
		UseDataAttribute: true,
		Prefix:           "theme",
	})

	cssPath := filepath.Join(outDir, "themes.css")
	if err := os.WriteFile(cssPath, []byte(css), 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing CSS: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Generated %s\n", cssPath)

	// Generate JSON for theme metadata
	jsonData, err := json.MarshalIndent(themeData, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	jsonPath := filepath.Join(outDir, "themes.json")
	if err := os.WriteFile(jsonPath, jsonData, 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing JSON: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Generated %s\n", jsonPath)

	fmt.Printf("\nGenerated %d themes\n", len(selectedThemes))
}
