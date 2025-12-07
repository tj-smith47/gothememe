package gothememe

import (
	"fmt"
	"strings"
)

// OutputFormat specifies the output format for theme CSS generation.
type OutputFormat int

const (
	// FormatCSS generates standard CSS with custom properties.
	FormatCSS OutputFormat = iota

	// FormatSCSS generates SCSS variables.
	FormatSCSS

	// FormatJSON generates a JSON representation.
	FormatJSON

	// FormatDesignTokens generates DTCG v1 compliant design tokens.
	FormatDesignTokens
)

// ColorSpace specifies the color space for output.
type ColorSpace int

const (
	// ColorSpaceHex outputs colors as hex values (#RRGGBB).
	ColorSpaceHex ColorSpace = iota

	// ColorSpaceRGB outputs colors as rgb() functions.
	ColorSpaceRGB

	// ColorSpaceHSL outputs colors as hsl() functions.
	ColorSpaceHSL

	// ColorSpaceOKLCH outputs colors as oklch() functions.
	ColorSpaceOKLCH
)

// CSSOptions configures CSS output generation.
type CSSOptions struct {
	// Prefix for CSS variable names (default: "theme").
	// Variables will be named --{prefix}-{property}.
	Prefix string

	// IncludeRoot wraps variables in :root {} selector.
	IncludeRoot bool

	// UseDataAttribute wraps variables in [data-theme="ID"] selector.
	UseDataAttribute bool

	// UseLightDark uses CSS light-dark() function for automatic mode switching.
	// Requires a light and dark theme pair.
	UseLightDark bool

	// IncludeMediaQuery adds @media (prefers-color-scheme: dark) query.
	IncludeMediaQuery bool

	// ColorSpace for output colors.
	ColorSpace ColorSpace

	// Minify removes whitespace and newlines.
	Minify bool

	// IncludeMetadata adds theme metadata as CSS comments.
	IncludeMetadata bool
}

// DefaultCSSOptions returns sensible default CSS options.
func DefaultCSSOptions() CSSOptions {
	return CSSOptions{
		Prefix:           "theme",
		IncludeRoot:      true,
		UseDataAttribute: false,
		ColorSpace:       ColorSpaceHex,
		IncludeMetadata:  true,
	}
}

// GenerateCSS generates CSS custom properties from a theme.
func GenerateCSS(t Theme, opts CSSOptions) string {
	if opts.Prefix == "" {
		opts.Prefix = "theme"
	}

	var sb strings.Builder

	// Add metadata comment
	if opts.IncludeMetadata && !opts.Minify {
		sb.WriteString(fmt.Sprintf("/* Theme: %s (%s) */\n", t.DisplayName(), t.ID()))
		if t.Author() != "" {
			sb.WriteString(fmt.Sprintf("/* Author: %s */\n", t.Author()))
		}
		if t.License() != "" {
			sb.WriteString(fmt.Sprintf("/* License: %s */\n", t.License()))
		}
		sb.WriteString("\n")
	}

	// Determine selector
	var selector string
	if opts.UseDataAttribute {
		selector = fmt.Sprintf("[data-theme=%q]", t.ID())
	} else if opts.IncludeRoot {
		selector = ":root"
	}

	// Build CSS
	if selector != "" {
		sb.WriteString(selector)
		if opts.Minify {
			sb.WriteString("{")
		} else {
			sb.WriteString(" {\n")
		}
	}

	// Generate variables
	vars := generateVariables(t, opts)
	for _, v := range vars {
		if opts.Minify {
			sb.WriteString(fmt.Sprintf("--%s-%s:%s;", opts.Prefix, v.name, v.value))
		} else {
			sb.WriteString(fmt.Sprintf("    --%s-%s: %s;\n", opts.Prefix, v.name, v.value))
		}
	}

	if selector != "" {
		if opts.Minify {
			sb.WriteString("}")
		} else {
			sb.WriteString("}\n")
		}
	}

	return sb.String()
}

// GenerateAllThemesCSS generates CSS for multiple themes using data-theme selectors.
func GenerateAllThemesCSS(themes []Theme, opts CSSOptions) string {
	var sb strings.Builder

	// Force data attribute mode for multi-theme output
	opts.UseDataAttribute = true
	opts.IncludeRoot = false

	for i, t := range themes {
		if i > 0 && !opts.Minify {
			sb.WriteString("\n")
		}
		sb.WriteString(GenerateCSS(t, opts))
	}

	return sb.String()
}

// SyntaxOptions configures syntax highlighting CSS generation.
type SyntaxOptions struct {
	// Format specifies the syntax highlighting library to target.
	Format SyntaxFormat

	// Prefix for CSS variable references.
	Prefix string

	// UseVariables references theme CSS variables instead of inline colors.
	UseVariables bool

	// Minify removes whitespace.
	Minify bool
}

// SyntaxFormat specifies the target syntax highlighting library.
type SyntaxFormat int

const (
	// SyntaxPrism generates Prism.js compatible CSS.
	SyntaxPrism SyntaxFormat = iota

	// SyntaxHighlightJS generates Highlight.js compatible CSS.
	SyntaxHighlightJS

	// SyntaxChroma generates Chroma (Go) compatible CSS.
	SyntaxChroma
)

// DefaultSyntaxOptions returns sensible default syntax options.
func DefaultSyntaxOptions() SyntaxOptions {
	return SyntaxOptions{
		Format:       SyntaxPrism,
		Prefix:       "theme",
		UseVariables: true,
	}
}

// GenerateSyntaxCSS generates syntax highlighting CSS for a theme.
func GenerateSyntaxCSS(t Theme, opts SyntaxOptions) string {
	var sb strings.Builder

	getColor := func(c Color, varName string) string {
		if opts.UseVariables {
			return fmt.Sprintf("var(--%s-%s)", opts.Prefix, varName)
		}
		return c.Hex()
	}

	nl := "\n"
	indent := "    "
	if opts.Minify {
		nl = ""
		indent = ""
	}

	switch opts.Format {
	case SyntaxPrism:
		// Prism.js token classes
		rules := []struct {
			selector string
			color    Color
			varName  string
		}{
			{".token.comment, .token.prolog, .token.doctype, .token.cdata", t.CodeComment(), "code-comment"},
			{".token.punctuation", t.CodePunctuation(), "code-punctuation"},
			{".token.property, .token.tag, .token.boolean, .token.number, .token.constant, .token.symbol", t.CodeNumber(), "code-number"},
			{".token.selector, .token.attr-name, .token.string, .token.char, .token.builtin", t.CodeString(), "code-string"},
			{".token.operator, .token.entity, .token.url", t.CodeOperator(), "code-operator"},
			{".token.atrule, .token.attr-value, .token.keyword", t.CodeKeyword(), "code-keyword"},
			{".token.function, .token.class-name", t.CodeFunction(), "code-function"},
			{".token.regex, .token.important, .token.variable", t.CodeVariable(), "code-variable"},
		}

		for _, rule := range rules {
			sb.WriteString(fmt.Sprintf("%s {%s%scolor: %s;%s}%s",
				rule.selector, nl, indent, getColor(rule.color, rule.varName), nl, nl))
		}

	case SyntaxHighlightJS:
		// Highlight.js classes
		rules := []struct {
			selector string
			color    Color
			varName  string
		}{
			{".hljs-comment, .hljs-quote", t.CodeComment(), "code-comment"},
			{".hljs-keyword, .hljs-selector-tag", t.CodeKeyword(), "code-keyword"},
			{".hljs-string, .hljs-doctag", t.CodeString(), "code-string"},
			{".hljs-number, .hljs-literal", t.CodeNumber(), "code-number"},
			{".hljs-title, .hljs-section, .hljs-selector-id", t.CodeFunction(), "code-function"},
			{".hljs-variable, .hljs-template-variable", t.CodeVariable(), "code-variable"},
			{".hljs-type, .hljs-class .hljs-title", t.CodeType(), "code-type"},
			{".hljs-symbol, .hljs-bullet", t.CodeConstant(), "code-constant"},
			{".hljs-attribute", t.CodeOperator(), "code-operator"},
		}

		for _, rule := range rules {
			sb.WriteString(fmt.Sprintf("%s {%s%scolor: %s;%s}%s",
				rule.selector, nl, indent, getColor(rule.color, rule.varName), nl, nl))
		}

	case SyntaxChroma:
		// Chroma (Go) classes
		rules := []struct {
			selector string
			color    Color
			varName  string
		}{
			{".chroma .c, .chroma .cm, .chroma .c1, .chroma .cs", t.CodeComment(), "code-comment"},
			{".chroma .k, .chroma .kc, .chroma .kd, .chroma .kn, .chroma .kp, .chroma .kr", t.CodeKeyword(), "code-keyword"},
			{".chroma .s, .chroma .sa, .chroma .sb, .chroma .sc, .chroma .dl, .chroma .sd, .chroma .s2, .chroma .se, .chroma .sh, .chroma .si, .chroma .sx, .chroma .sr, .chroma .s1, .chroma .ss", t.CodeString(), "code-string"},
			{".chroma .m, .chroma .mb, .chroma .mf, .chroma .mh, .chroma .mi, .chroma .il, .chroma .mo", t.CodeNumber(), "code-number"},
			{".chroma .nf, .chroma .fm", t.CodeFunction(), "code-function"},
			{".chroma .nv, .chroma .vc, .chroma .vg, .chroma .vi, .chroma .vm", t.CodeVariable(), "code-variable"},
			{".chroma .nc, .chroma .no, .chroma .nd, .chroma .ni, .chroma .ne, .chroma .nf, .chroma .nl, .chroma .nn, .chroma .nt", t.CodeType(), "code-type"},
			{".chroma .o, .chroma .ow", t.CodeOperator(), "code-operator"},
			{".chroma .p", t.CodePunctuation(), "code-punctuation"},
		}

		for _, rule := range rules {
			sb.WriteString(fmt.Sprintf("%s {%s%scolor: %s;%s}%s",
				rule.selector, nl, indent, getColor(rule.color, rule.varName), nl, nl))
		}
	}

	return sb.String()
}

// cssVariable represents a CSS custom property.
type cssVariable struct {
	name  string
	value string
}

// generateVariables creates all CSS variables for a theme.
func generateVariables(t Theme, opts CSSOptions) []cssVariable {
	formatColor := func(c Color) string {
		if c.IsEmpty() {
			return "transparent"
		}
		switch opts.ColorSpace {
		case ColorSpaceRGB:
			return c.CSSRGB()
		case ColorSpaceHSL:
			return c.CSSHSL()
		case ColorSpaceOKLCH:
			l, ch, h := c.OKLCHValues()
			return fmt.Sprintf("oklch(%.3f %.3f %.1f)", l, ch, h)
		default:
			return c.Hex()
		}
	}

	vars := []cssVariable{
		// Background colors
		{"background", formatColor(t.Background())},
		{"background-secondary", formatColor(t.BackgroundSecondary())},
		{"surface", formatColor(t.Surface())},
		{"surface-secondary", formatColor(t.SurfaceSecondary())},

		// Text colors
		{"text-primary", formatColor(t.TextPrimary())},
		{"text-secondary", formatColor(t.TextSecondary())},
		{"text-muted", formatColor(t.TextMuted())},
		{"text-inverted", formatColor(t.TextInverted())},

		// Accent/Brand colors
		{"accent", formatColor(t.Accent())},
		{"accent-secondary", formatColor(t.AccentSecondary())},
		{"brand", formatColor(t.Brand())},

		// Border colors
		{"border", formatColor(t.Border())},
		{"border-subtle", formatColor(t.BorderSubtle())},
		{"border-strong", formatColor(t.BorderStrong())},

		// Semantic colors
		{"success-background", formatColor(t.Success().Background)},
		{"success-border", formatColor(t.Success().Border)},
		{"success-text", formatColor(t.Success().Text)},
		{"warning-background", formatColor(t.Warning().Background)},
		{"warning-border", formatColor(t.Warning().Border)},
		{"warning-text", formatColor(t.Warning().Text)},
		{"error-background", formatColor(t.Error().Background)},
		{"error-border", formatColor(t.Error().Border)},
		{"error-text", formatColor(t.Error().Text)},
		{"info-background", formatColor(t.Info().Background)},
		{"info-border", formatColor(t.Info().Border)},
		{"info-text", formatColor(t.Info().Text)},

		// ANSI colors
		{"black", formatColor(t.Black())},
		{"red", formatColor(t.Red())},
		{"green", formatColor(t.Green())},
		{"yellow", formatColor(t.Yellow())},
		{"blue", formatColor(t.Blue())},
		{"purple", formatColor(t.Purple())},
		{"cyan", formatColor(t.Cyan())},
		{"white", formatColor(t.White())},
		{"bright-black", formatColor(t.BrightBlack())},
		{"bright-red", formatColor(t.BrightRed())},
		{"bright-green", formatColor(t.BrightGreen())},
		{"bright-yellow", formatColor(t.BrightYellow())},
		{"bright-blue", formatColor(t.BrightBlue())},
		{"bright-purple", formatColor(t.BrightPurple())},
		{"bright-cyan", formatColor(t.BrightCyan())},
		{"bright-white", formatColor(t.BrightWhite())},

		// Code colors
		{"code-background", formatColor(t.CodeBackground())},
		{"code-text", formatColor(t.CodeText())},
		{"code-comment", formatColor(t.CodeComment())},
		{"code-keyword", formatColor(t.CodeKeyword())},
		{"code-string", formatColor(t.CodeString())},
		{"code-number", formatColor(t.CodeNumber())},
		{"code-function", formatColor(t.CodeFunction())},
		{"code-operator", formatColor(t.CodeOperator())},
		{"code-punctuation", formatColor(t.CodePunctuation())},
		{"code-variable", formatColor(t.CodeVariable())},
		{"code-constant", formatColor(t.CodeConstant())},
		{"code-type", formatColor(t.CodeType())},
	}

	return vars
}

// GenerateSCSS generates SCSS variables from a theme.
func GenerateSCSS(t Theme, opts CSSOptions) string {
	if opts.Prefix == "" {
		opts.Prefix = "theme"
	}

	var sb strings.Builder

	// Add metadata comment
	if opts.IncludeMetadata && !opts.Minify {
		sb.WriteString(fmt.Sprintf("// Theme: %s (%s)\n", t.DisplayName(), t.ID()))
		if t.Author() != "" {
			sb.WriteString(fmt.Sprintf("// Author: %s\n", t.Author()))
		}
		sb.WriteString("\n")
	}

	vars := generateVariables(t, opts)
	for _, v := range vars {
		if opts.Minify {
			sb.WriteString(fmt.Sprintf("$%s-%s:%s;", opts.Prefix, v.name, v.value))
		} else {
			sb.WriteString(fmt.Sprintf("$%s-%s: %s;\n", opts.Prefix, v.name, v.value))
		}
	}

	return sb.String()
}

// GenerateJSON generates a JSON representation of the theme colors.
func GenerateJSON(t Theme, opts CSSOptions) string {
	var sb strings.Builder

	nl := "\n"
	indent := "    "
	if opts.Minify {
		nl = ""
		indent = ""
	}

	sb.WriteString("{" + nl)

	vars := generateVariables(t, opts)
	for i, v := range vars {
		comma := ","
		if i == len(vars)-1 {
			comma = ""
		}
		sb.WriteString(fmt.Sprintf("%s%q: %q%s%s", indent, v.name, v.value, comma, nl))
	}

	sb.WriteString("}")

	return sb.String()
}
