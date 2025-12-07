// Example: HTMX theme switcher with server-side rendering
//
// This example demonstrates how to use gothememe with HTMX
// for dynamic theme switching without full page reloads.
//
// Run: go run main.go
// Then open http://localhost:8080 in your browser.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/tj-smith47/gothememe"
	"github.com/tj-smith47/gothememe/themes"
)

var (
	selectedThemes []gothememe.Theme
	tmpl           *template.Template
)

func init() {
	// Select themes
	ids := []string{"dracula", "nord", "gruvbox_dark", "one_dark", "solarized_light", "catppuccin_mocha"}
	for _, id := range ids {
		if t := themes.ByID(id); t != nil {
			selectedThemes = append(selectedThemes, t)
		}
	}

	// Parse templates
	tmpl = template.Must(template.New("").Parse(pageTemplate + themeOptionsTemplate + previewTemplate))
}

func main() {
	// Serve theme CSS
	http.HandleFunc("/themes.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		css := gothememe.GenerateAllThemesCSS(selectedThemes, gothememe.CSSOptions{
			UseDataAttribute: true,
		})
		fmt.Fprint(w, css)
	})

	// Main page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"Themes":  selectedThemes,
			"Current": "dracula",
		}
		tmpl.ExecuteTemplate(w, "page", data)
	})

	// HTMX: Get theme options
	http.HandleFunc("/theme-options", func(w http.ResponseWriter, r *http.Request) {
		current := r.URL.Query().Get("current")
		data := map[string]interface{}{
			"Themes":  selectedThemes,
			"Current": current,
		}
		tmpl.ExecuteTemplate(w, "theme-options", data)
	})

	// HTMX: Get preview panel
	http.HandleFunc("/preview", func(w http.ResponseWriter, r *http.Request) {
		themeID := r.URL.Query().Get("theme")
		t := themes.ByID(themeID)
		if t == nil {
			http.Error(w, "Theme not found", 404)
			return
		}

		stats := gothememe.AnalyzeTheme(t)
		data := map[string]interface{}{
			"Theme": t,
			"Stats": stats,
		}
		tmpl.ExecuteTemplate(w, "preview", data)
	})

	log.Println("Starting server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

const pageTemplate = `{{define "page"}}<!DOCTYPE html>
<html lang="en" data-theme="{{.Current}}">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>GoThemeMe - HTMX Example</title>
    <script src="https://unpkg.com/htmx.org@2.0.0"></script>
    <link rel="stylesheet" href="/themes.css">
    <style>
        body {
            font-family: system-ui, sans-serif;
            background: var(--theme-background);
            color: var(--theme-text-primary);
            padding: 2rem;
            transition: background-color 0.3s, color 0.3s;
        }
        h1 { color: var(--theme-accent); }
        .container { max-width: 800px; margin: 0 auto; }
        .card {
            background: var(--theme-surface);
            border: 1px solid var(--theme-border);
            border-radius: 8px;
            padding: 1rem;
            margin: 1rem 0;
        }
        .theme-btn {
            background: var(--theme-surface);
            color: var(--theme-text-primary);
            border: 2px solid var(--theme-border);
            padding: 0.5rem 1rem;
            margin: 0.25rem;
            border-radius: 4px;
            cursor: pointer;
        }
        .theme-btn:hover, .theme-btn.active {
            border-color: var(--theme-accent);
            background: var(--theme-accent);
            color: var(--theme-text-inverted);
        }
        .stats { font-family: monospace; font-size: 0.9rem; }
    </style>
</head>
<body>
    <div class="container">
        <h1>GoThemeMe + HTMX</h1>

        <div class="card" id="theme-selector">
            {{template "theme-options" .}}
        </div>

        <div class="card" id="preview"
             hx-get="/preview?theme={{.Current}}"
             hx-trigger="load">
            Loading...
        </div>
    </div>

    <script>
        function setTheme(id) {
            document.documentElement.setAttribute('data-theme', id);
            localStorage.setItem('theme', id);
        }
    </script>
</body>
</html>{{end}}`

const themeOptionsTemplate = `{{define "theme-options"}}
<p>Select a theme:</p>
{{range .Themes}}
<button class="theme-btn {{if eq .ID $.Current}}active{{end}}"
        onclick="setTheme('{{.ID}}')"
        hx-get="/preview?theme={{.ID}}"
        hx-target="#preview"
        hx-swap="innerHTML">
    {{.DisplayName}}
</button>
{{end}}
{{end}}`

const previewTemplate = `{{define "preview"}}
<h3>{{.Theme.DisplayName}}</h3>
<p style="color: var(--theme-text-secondary)">
    {{if .Theme.IsDark}}Dark{{else}}Light{{end}} theme
    {{if .Theme.Author}}by {{.Theme.Author}}{{end}}
</p>
<div class="stats">
    <p>Colors: {{.Stats.ColorCount}} ({{.Stats.UniqueColors}} unique)</p>
    <p>Contrast Score: {{printf "%.2f" .Stats.ContrastScore}}:1</p>
    <p>Accessibility: {{printf "%.0f" .Stats.AccessibilityPercent}}%</p>
</div>
{{end}}`
