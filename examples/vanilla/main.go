// Example: Vanilla HTML/CSS theme switcher
//
// This example demonstrates how to use gothememe to generate CSS
// for a static HTML page with theme switching.
//
// Run: go run main.go
// Then open http://localhost:8080 in your browser.
package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/tj-smith47/gothememe"
	"github.com/tj-smith47/gothememe/themes"
)

//go:embed static
var staticFS embed.FS

func main() {
	// Select some themes to use
	selectedThemes := []gothememe.Theme{
		themes.ByID("dracula"),
		themes.ByID("nord"),
		themes.ByID("gruvbox_dark"),
		themes.ByID("one_dark"),
		themes.ByID("solarized_light"),
	}

	// Filter out nil themes
	var validThemes []gothememe.Theme
	for _, t := range selectedThemes {
		if t != nil {
			validThemes = append(validThemes, t)
		}
	}

	// Generate CSS for all themes with data-theme attribute
	css := gothememe.GenerateAllThemesCSS(validThemes, gothememe.CSSOptions{
		UseDataAttribute: true,
		IncludeMetadata:  true,
	})

	// Set up HTTP server
	http.HandleFunc("/themes.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		fmt.Fprint(w, css)
	})

	http.HandleFunc("/themes.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[`)
		for i, t := range validThemes {
			if i > 0 {
				fmt.Fprint(w, ",")
			}
			fmt.Fprintf(w, `{"id":%q,"name":%q,"dark":%t}`, t.ID(), t.DisplayName(), t.IsDark())
		}
		fmt.Fprint(w, `]`)
	})

	// Serve static files
	http.Handle("/", http.FileServer(http.FS(staticFS)))

	addr := ":8080"
	log.Printf("Starting server at http://localhost%s", addr)
	log.Printf("Available themes: %d", len(validThemes))
	log.Fatal(http.ListenAndServe(addr, nil))
}
