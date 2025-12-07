// Example: templ template engine with gothememe
//
// This example demonstrates integrating gothememe with templ templates.
// Note: This example uses html/template as a stand-in for templ syntax demonstration.
// In a real templ project, you would use .templ files.
//
// Run: go run main.go
// Then open http://localhost:8080 in your browser.
package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/tj-smith47/gothememe"
	"github.com/tj-smith47/gothememe/themes"
)

var (
	allThemes []gothememe.Theme
	tmpl      *template.Template
)

func init() {
	// Get all themes
	allThemes = themes.All()

	// Parse template (in real templ, this would be auto-generated)
	tmpl = template.Must(template.New("page").Parse(pageHTML))
}

func main() {
	// Serve theme CSS
	http.HandleFunc("/themes.css", handleCSS)

	// Main page
	http.HandleFunc("/", handleIndex)

	log.Println("Starting server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCSS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")

	// Only generate CSS for a subset for performance
	var subset []gothememe.Theme
	for _, t := range allThemes {
		if len(subset) >= 20 {
			break
		}
		subset = append(subset, t)
	}

	css := gothememe.GenerateAllThemesCSS(subset, gothememe.CSSOptions{
		UseDataAttribute: true,
	})
	w.Write([]byte(css))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Themes     []gothememe.Theme
		ThemeCount int
	}{
		Themes:     allThemes[:min(20, len(allThemes))],
		ThemeCount: len(allThemes),
	}
	tmpl.Execute(w, data)
}

const pageHTML = `<!DOCTYPE html>
<html lang="en" data-theme="dracula">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>GoThemeMe - Templ Example</title>
    <link rel="stylesheet" href="/themes.css">
    <style>
        body {
            font-family: system-ui, sans-serif;
            background: var(--theme-background);
            color: var(--theme-text-primary);
            padding: 2rem;
            transition: all 0.3s;
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
        select {
            background: var(--theme-background);
            color: var(--theme-text-primary);
            border: 1px solid var(--theme-border);
            padding: 0.5rem;
            border-radius: 4px;
            font-size: 1rem;
        }
        .code {
            background: var(--theme-code-background, var(--theme-surface));
            color: var(--theme-code-text, var(--theme-text-primary));
            padding: 1rem;
            border-radius: 4px;
            font-family: monospace;
            overflow-x: auto;
        }
        .badge {
            display: inline-block;
            padding: 0.25rem 0.5rem;
            border-radius: 4px;
            font-size: 0.8rem;
            margin-right: 0.5rem;
        }
        .badge-success {
            background: var(--theme-success-background);
            color: var(--theme-success-text);
        }
        .badge-info {
            background: var(--theme-info-background);
            color: var(--theme-info-text);
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>GoThemeMe + Templ</h1>

        <p style="color: var(--theme-text-secondary);">
            <span class="badge badge-success">{{.ThemeCount}} themes</span>
            <span class="badge badge-info">Server-side rendered</span>
        </p>

        <div class="card">
            <label for="theme">Select Theme:</label>
            <select id="theme" onchange="setTheme(this.value)">
                {{range .Themes}}
                <option value="{{.ID}}">{{.DisplayName}}{{if .IsDark}} (Dark){{else}} (Light){{end}}</option>
                {{end}}
            </select>
        </div>

        <div class="card">
            <h3>Example Code</h3>
            <pre class="code">func main() {
    theme := themes.ByID("dracula")
    css := gothememe.GenerateCSS(theme, gothememe.CSSOptions{
        IncludeRoot: true,
    })
    fmt.Println(css)
}</pre>
        </div>

        <div class="card">
            <h3>Semantic Colors</h3>
            <p class="badge badge-success">Success state</p>
            <p style="background: var(--theme-warning-background); color: var(--theme-warning-text); padding: 0.5rem; border-radius: 4px; margin-top: 0.5rem;">Warning state</p>
            <p style="background: var(--theme-error-background); color: var(--theme-error-text); padding: 0.5rem; border-radius: 4px; margin-top: 0.5rem;">Error state</p>
        </div>
    </div>

    <script>
        function setTheme(id) {
            document.documentElement.setAttribute('data-theme', id);
            localStorage.setItem('theme', id);
        }
        // Restore saved theme
        const saved = localStorage.getItem('theme');
        if (saved) {
            setTheme(saved);
            document.getElementById('theme').value = saved;
        }
    </script>
</body>
</html>`
