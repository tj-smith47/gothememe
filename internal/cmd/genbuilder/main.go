// genbuilder generates ThemeBuilder setter methods from a field definition.
package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

type field struct {
	Name    string // Go field name (e.g., "background")
	Method  string // Method name (e.g., "WithBackground")
	Type    string // Field type (e.g., "Color", "string", "bool")
	Doc     string // Method documentation
	Section string // Section comment (e.g., "// Background colors")
}

var fields = []field{
	// Metadata
	{Name: "description", Method: "WithDescription", Type: "string", Doc: "sets the theme description.", Section: "// Metadata methods"},
	{Name: "author", Method: "WithAuthor", Type: "string", Doc: "sets the theme author."},
	{Name: "license", Method: "WithLicense", Type: "string", Doc: "sets the theme license."},
	{Name: "source", Method: "WithSource", Type: "string", Doc: "sets the theme source URL."},
	{Name: "isDark", Method: "WithIsDark", Type: "bool", Doc: "sets whether this is a dark theme."},

	// Background colors
	{Name: "background", Method: "WithBackground", Type: "Color", Doc: "sets the primary background color.", Section: "// Background colors"},
	{Name: "backgroundSecondary", Method: "WithBackgroundSecondary", Type: "Color", Doc: "sets the secondary background color."},
	{Name: "surface", Method: "WithSurface", Type: "Color", Doc: "sets the surface color."},
	{Name: "surfaceSecondary", Method: "WithSurfaceSecondary", Type: "Color", Doc: "sets the secondary surface color."},

	// Text colors
	{Name: "textPrimary", Method: "WithTextPrimary", Type: "Color", Doc: "sets the primary text color.", Section: "// Text colors"},
	{Name: "textSecondary", Method: "WithTextSecondary", Type: "Color", Doc: "sets the secondary text color."},
	{Name: "textMuted", Method: "WithTextMuted", Type: "Color", Doc: "sets the muted text color."},
	{Name: "textInverted", Method: "WithTextInverted", Type: "Color", Doc: "sets the inverted text color."},

	// Accent/Brand colors
	{Name: "accent", Method: "WithAccent", Type: "Color", Doc: "sets the primary accent color.", Section: "// Accent/Brand colors"},
	{Name: "accentSecondary", Method: "WithAccentSecondary", Type: "Color", Doc: "sets the secondary accent color."},
	{Name: "brand", Method: "WithBrand", Type: "Color", Doc: "sets the brand color."},

	// Border colors
	{Name: "border", Method: "WithBorder", Type: "Color", Doc: "sets the primary border color.", Section: "// Border colors"},
	{Name: "borderSubtle", Method: "WithBorderSubtle", Type: "Color", Doc: "sets the subtle border color."},
	{Name: "borderStrong", Method: "WithBorderStrong", Type: "Color", Doc: "sets the strong border color."},

	// Semantic colors (these are SemanticColor not Color)
	{Name: "success", Method: "WithSuccess", Type: "SemanticColor", Doc: "sets the success color.", Section: "// Semantic colors"},
	{Name: "warning", Method: "WithWarning", Type: "SemanticColor", Doc: "sets the warning color."},
	{Name: "errorColor", Method: "WithError", Type: "SemanticColor", Doc: "sets the error color."},
	{Name: "info", Method: "WithInfo", Type: "SemanticColor", Doc: "sets the info color."},

	// ANSI colors
	{Name: "black", Method: "WithBlack", Type: "Color", Doc: "sets the black ANSI color.", Section: "// ANSI colors"},
	{Name: "red", Method: "WithRed", Type: "Color", Doc: "sets the red ANSI color."},
	{Name: "green", Method: "WithGreen", Type: "Color", Doc: "sets the green ANSI color."},
	{Name: "yellow", Method: "WithYellow", Type: "Color", Doc: "sets the yellow ANSI color."},
	{Name: "blue", Method: "WithBlue", Type: "Color", Doc: "sets the blue ANSI color."},
	{Name: "purple", Method: "WithPurple", Type: "Color", Doc: "sets the purple ANSI color."},
	{Name: "cyan", Method: "WithCyan", Type: "Color", Doc: "sets the cyan ANSI color."},
	{Name: "white", Method: "WithWhite", Type: "Color", Doc: "sets the white ANSI color."},
	{Name: "brightBlack", Method: "WithBrightBlack", Type: "Color", Doc: "sets the bright black ANSI color."},
	{Name: "brightRed", Method: "WithBrightRed", Type: "Color", Doc: "sets the bright red ANSI color."},
	{Name: "brightGreen", Method: "WithBrightGreen", Type: "Color", Doc: "sets the bright green ANSI color."},
	{Name: "brightYellow", Method: "WithBrightYellow", Type: "Color", Doc: "sets the bright yellow ANSI color."},
	{Name: "brightBlue", Method: "WithBrightBlue", Type: "Color", Doc: "sets the bright blue ANSI color."},
	{Name: "brightPurple", Method: "WithBrightPurple", Type: "Color", Doc: "sets the bright purple ANSI color."},
	{Name: "brightCyan", Method: "WithBrightCyan", Type: "Color", Doc: "sets the bright cyan ANSI color."},
	{Name: "brightWhite", Method: "WithBrightWhite", Type: "Color", Doc: "sets the bright white ANSI color."},

	// Code/Syntax highlighting colors
	{Name: "codeBackground", Method: "WithCodeBackground", Type: "Color", Doc: "sets the code background color.", Section: "// Code/Syntax highlighting colors"},
	{Name: "codeText", Method: "WithCodeText", Type: "Color", Doc: "sets the code text color."},
	{Name: "codeComment", Method: "WithCodeComment", Type: "Color", Doc: "sets the code comment color."},
	{Name: "codeKeyword", Method: "WithCodeKeyword", Type: "Color", Doc: "sets the code keyword color."},
	{Name: "codeString", Method: "WithCodeString", Type: "Color", Doc: "sets the code string color."},
	{Name: "codeNumber", Method: "WithCodeNumber", Type: "Color", Doc: "sets the code number color."},
	{Name: "codeFunction", Method: "WithCodeFunction", Type: "Color", Doc: "sets the code function color."},
	{Name: "codeOperator", Method: "WithCodeOperator", Type: "Color", Doc: "sets the code operator color."},
	{Name: "codePunctuation", Method: "WithCodePunctuation", Type: "Color", Doc: "sets the code punctuation color."},
	{Name: "codeVariable", Method: "WithCodeVariable", Type: "Color", Doc: "sets the code variable color."},
	{Name: "codeConstant", Method: "WithCodeConstant", Type: "Color", Doc: "sets the code constant color."},
	{Name: "codeType", Method: "WithCodeType", Type: "Color", Doc: "sets the code type color."},
}

const methodTemplate = `
{{- if .Section}}

{{.Section}}
{{- end}}

// {{.Method}} {{.Doc}}
func (b *ThemeBuilder) {{.Method}}({{if eq .Type "bool"}}{{camelCase .Name}}{{else}}c{{end}} {{.Type}}) *ThemeBuilder {
	b.theme.{{.Name}} = {{if eq .Type "bool"}}{{camelCase .Name}}{{else}}c{{end}}
	return b
}
`

func main() {
	outputFile := flag.String("output", "", "output file (defaults to stdout)")
	flag.Parse()

	tmpl := template.Must(template.New("method").Funcs(template.FuncMap{
		"camelCase": func(s string) string {
			return s
		},
	}).Parse(methodTemplate))

	// Determine output destination
	var out *os.File
	var needsClose bool
	if *outputFile != "" {
		f, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		out = f
		needsClose = true
	} else {
		out = os.Stdout
	}

	// Write header
	if _, err := fmt.Fprintln(out, "// Code generated by go generate; DO NOT EDIT."); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing header: %v\n", err)
		os.Exit(1)
	}
	if _, err := fmt.Fprintln(out, "// This file is generated from internal/cmd/genbuilder/main.go"); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing header: %v\n", err)
		os.Exit(1)
	}
	if _, err := fmt.Fprintln(out); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing header: %v\n", err)
		os.Exit(1)
	}
	if _, err := fmt.Fprintln(out, "package gothememe"); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing package: %v\n", err)
		os.Exit(1)
	}
	if _, err := fmt.Fprintln(out); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing newline: %v\n", err)
		os.Exit(1)
	}

	// Generate methods
	for _, f := range fields {
		if err := tmpl.Execute(out, f); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing template: %v\n", err)
			if needsClose {
				if closeErr := out.Close(); closeErr != nil {
					fmt.Fprintf(os.Stderr, "Error closing file: %v\n", closeErr)
				}
			}
			os.Exit(1)
		}
	}

	// Close output file if needed
	if needsClose {
		if err := out.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "Error closing output file: %v\n", err)
			os.Exit(1)
		}
	}
}
