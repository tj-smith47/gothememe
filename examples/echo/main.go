// Example: Echo framework with theme API
//
// This example demonstrates how to use gothememe with the Echo web framework
// to provide a theme API endpoint and serve theme-aware pages.
//
// Run: go run main.go
// Then open http://localhost:8080 in your browser.
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/tj-smith47/gothememe"
	"github.com/tj-smith47/gothememe/themes"
)

// ThemeInfo represents theme metadata for the API.
type ThemeInfo struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	IsDark      bool   `json:"isDark"`
	Author      string `json:"author,omitempty"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Get all available themes
	allThemes := themes.All()

	// API: List all themes
	e.GET("/api/themes", func(c echo.Context) error {
		var list []ThemeInfo
		for _, t := range allThemes {
			list = append(list, ThemeInfo{
				ID:          t.ID(),
				DisplayName: t.DisplayName(),
				IsDark:      t.IsDark(),
				Author:      t.Author(),
			})
		}
		return c.JSON(http.StatusOK, list)
	})

	// API: Get theme CSS
	e.GET("/api/themes/:id/css", func(c echo.Context) error {
		t := themes.ByID(c.Param("id"))
		if t == nil {
			return c.String(http.StatusNotFound, "Theme not found")
		}

		css := gothememe.GenerateCSS(t, gothememe.CSSOptions{
			IncludeRoot:     true,
			IncludeMetadata: true,
		})

		c.Response().Header().Set("Content-Type", "text/css")
		return c.String(http.StatusOK, css)
	})

	// API: Get theme as JSON
	e.GET("/api/themes/:id/json", func(c echo.Context) error {
		t := themes.ByID(c.Param("id"))
		if t == nil {
			return c.String(http.StatusNotFound, "Theme not found")
		}

		json := gothememe.GenerateJSON(t, gothememe.CSSOptions{})
		c.Response().Header().Set("Content-Type", "application/json")
		return c.String(http.StatusOK, json)
	})

	// API: Get theme design tokens (DTCG format)
	e.GET("/api/themes/:id/tokens", func(c echo.Context) error {
		t := themes.ByID(c.Param("id"))
		if t == nil {
			return c.String(http.StatusNotFound, "Theme not found")
		}

		tokens, err := gothememe.GenerateDesignTokens(t, gothememe.DefaultTokenOptions())
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		c.Response().Header().Set("Content-Type", "application/json")
		return c.String(http.StatusOK, tokens)
	})

	// API: Get theme analysis
	e.GET("/api/themes/:id/stats", func(c echo.Context) error {
		t := themes.ByID(c.Param("id"))
		if t == nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Theme not found"})
		}

		stats := gothememe.AnalyzeTheme(t)
		return c.JSON(http.StatusOK, stats)
	})

	// API: Validate theme
	e.GET("/api/themes/:id/validate", func(c echo.Context) error {
		t := themes.ByID(c.Param("id"))
		if t == nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Theme not found"})
		}

		validationErrors := gothememe.ValidateTheme(t)
		contrastIssues := gothememe.ValidateContrast(t, gothememe.ContrastLevelAA)

		return c.JSON(http.StatusOK, map[string]interface{}{
			"valid":          len(validationErrors) == 0 && len(contrastIssues) == 0,
			"errors":         validationErrors,
			"contrastIssues": contrastIssues,
		})
	})

	// Serve all themes CSS
	e.GET("/themes.css", func(c echo.Context) error {
		css := gothememe.GenerateAllThemesCSS(allThemes, gothememe.CSSOptions{
			UseDataAttribute: true,
		})
		c.Response().Header().Set("Content-Type", "text/css")
		return c.String(http.StatusOK, css)
	})

	// Serve index page
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, indexHTML)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

const indexHTML = `<!DOCTYPE html>
<html lang="en" data-theme="dracula">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>GoThemeMe - Echo Example</title>
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
        }
        .api-links a {
            color: var(--theme-accent);
            margin-right: 1rem;
        }
    </style>
</head>
<body>
    <h1>GoThemeMe + Echo</h1>

    <div class="card">
        <label>Theme: </label>
        <select id="theme" onchange="setTheme(this.value)"></select>
    </div>

    <div class="card">
        <h3>API Endpoints</h3>
        <div class="api-links" id="links">
            <a href="/api/themes">All Themes</a>
        </div>
    </div>

    <div class="card">
        <h3>Theme Stats</h3>
        <pre id="stats"></pre>
    </div>

    <script>
        fetch('/api/themes')
            .then(r => r.json())
            .then(themes => {
                const select = document.getElementById('theme');
                themes.forEach(t => {
                    const opt = document.createElement('option');
                    opt.value = t.id;
                    opt.textContent = t.displayName;
                    select.appendChild(opt);
                });
                setTheme(localStorage.getItem('theme') || 'dracula');
            });

        function setTheme(id) {
            document.documentElement.setAttribute('data-theme', id);
            localStorage.setItem('theme', id);
            document.getElementById('theme').value = id;

            // Update API links
            document.getElementById('links').innerHTML =
                '<a href="/api/themes/' + id + '/css">CSS</a>' +
                '<a href="/api/themes/' + id + '/json">JSON</a>' +
                '<a href="/api/themes/' + id + '/tokens">Tokens</a>' +
                '<a href="/api/themes/' + id + '/stats">Stats</a>' +
                '<a href="/api/themes/' + id + '/validate">Validate</a>';

            // Load stats
            fetch('/api/themes/' + id + '/stats')
                .then(r => r.json())
                .then(s => {
                    document.getElementById('stats').textContent = JSON.stringify(s, null, 2);
                });
        }
    </script>
</body>
</html>`
