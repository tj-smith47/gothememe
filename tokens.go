package gothememe

import (
	"encoding/json"
	"fmt"
)

// DesignToken represents a DTCG v1 compliant design token.
type DesignToken struct {
	Value       string `json:"$value"`
	Type        string `json:"$type,omitempty"`
	Description string `json:"$description,omitempty"`
}

// DesignTokenGroup represents a group of design tokens.
type DesignTokenGroup struct {
	Type   string                 `json:"$type,omitempty"`
	Tokens map[string]interface{} `json:"-"`
}

// TokenOptions configures design token output.
type TokenOptions struct {
	// IncludeDescriptions adds $description to each token.
	IncludeDescriptions bool

	// Indent is the JSON indentation string (default: "  ").
	Indent string
}

// DefaultTokenOptions returns sensible default token options.
func DefaultTokenOptions() TokenOptions {
	return TokenOptions{
		IncludeDescriptions: true,
		Indent:              "  ",
	}
}

// GenerateDesignTokens generates DTCG v1 compliant design tokens JSON.
func GenerateDesignTokens(t Theme, opts TokenOptions) (string, error) {
	if opts.Indent == "" {
		opts.Indent = "  "
	}

	tokens := buildTokenStructure(t, opts)

	var data []byte
	var err error

	if opts.Indent != "" {
		data, err = json.MarshalIndent(tokens, "", opts.Indent)
	} else {
		data, err = json.Marshal(tokens)
	}

	if err != nil {
		return "", fmt.Errorf("failed to marshal design tokens: %w", err)
	}

	return string(data), nil
}

// buildTokenStructure creates the hierarchical token structure.
func buildTokenStructure(t Theme, opts TokenOptions) map[string]interface{} {
	makeToken := func(c Color, desc string) map[string]interface{} {
		token := map[string]interface{}{
			"$value": c.Hex(),
			"$type":  "color",
		}
		if opts.IncludeDescriptions && desc != "" {
			token["$description"] = desc
		}
		return token
	}

	makeSemanticGroup := func(sc SemanticColor, baseName string) map[string]interface{} {
		return map[string]interface{}{
			"background": makeToken(sc.Background, fmt.Sprintf("%s background color", baseName)),
			"border":     makeToken(sc.Border, fmt.Sprintf("%s border color", baseName)),
			"text":       makeToken(sc.Text, fmt.Sprintf("%s text color", baseName)),
		}
	}

	tokens := map[string]interface{}{
		"$description": fmt.Sprintf("Design tokens for %s theme", t.DisplayName()),
		"color": map[string]interface{}{
			"$type": "color",
			"background": map[string]interface{}{
				"primary":   makeToken(t.Background(), "Primary background color"),
				"secondary": makeToken(t.BackgroundSecondary(), "Secondary background color"),
			},
			"surface": map[string]interface{}{
				"primary":   makeToken(t.Surface(), "Primary surface color for cards/modals"),
				"secondary": makeToken(t.SurfaceSecondary(), "Secondary surface color"),
			},
			"text": map[string]interface{}{
				"primary":   makeToken(t.TextPrimary(), "Primary text color"),
				"secondary": makeToken(t.TextSecondary(), "Secondary text color"),
				"muted":     makeToken(t.TextMuted(), "Muted text color for placeholders"),
				"inverted":  makeToken(t.TextInverted(), "Inverted text for colored backgrounds"),
			},
			"accent": map[string]interface{}{
				"primary":   makeToken(t.Accent(), "Primary accent color"),
				"secondary": makeToken(t.AccentSecondary(), "Secondary accent color"),
			},
			"brand": makeToken(t.Brand(), "Brand/logo color"),
			"border": map[string]interface{}{
				"default": makeToken(t.Border(), "Default border color"),
				"subtle":  makeToken(t.BorderSubtle(), "Subtle border color"),
				"strong":  makeToken(t.BorderStrong(), "Strong/emphasized border color"),
			},
			"semantic": map[string]interface{}{
				"success": makeSemanticGroup(t.Success(), "Success"),
				"warning": makeSemanticGroup(t.Warning(), "Warning"),
				"error":   makeSemanticGroup(t.Error(), "Error"),
				"info":    makeSemanticGroup(t.Info(), "Info"),
			},
			"ansi": map[string]interface{}{
				"black":         makeToken(t.Black(), "ANSI black"),
				"red":           makeToken(t.Red(), "ANSI red"),
				"green":         makeToken(t.Green(), "ANSI green"),
				"yellow":        makeToken(t.Yellow(), "ANSI yellow"),
				"blue":          makeToken(t.Blue(), "ANSI blue"),
				"purple":        makeToken(t.Purple(), "ANSI purple/magenta"),
				"cyan":          makeToken(t.Cyan(), "ANSI cyan"),
				"white":         makeToken(t.White(), "ANSI white"),
				"bright-black":  makeToken(t.BrightBlack(), "Bright ANSI black"),
				"bright-red":    makeToken(t.BrightRed(), "Bright ANSI red"),
				"bright-green":  makeToken(t.BrightGreen(), "Bright ANSI green"),
				"bright-yellow": makeToken(t.BrightYellow(), "Bright ANSI yellow"),
				"bright-blue":   makeToken(t.BrightBlue(), "Bright ANSI blue"),
				"bright-purple": makeToken(t.BrightPurple(), "Bright ANSI purple"),
				"bright-cyan":   makeToken(t.BrightCyan(), "Bright ANSI cyan"),
				"bright-white":  makeToken(t.BrightWhite(), "Bright ANSI white"),
			},
			"code": map[string]interface{}{
				"background":  makeToken(t.CodeBackground(), "Code block background"),
				"text":        makeToken(t.CodeText(), "Default code text"),
				"comment":     makeToken(t.CodeComment(), "Code comment color"),
				"keyword":     makeToken(t.CodeKeyword(), "Code keyword color"),
				"string":      makeToken(t.CodeString(), "Code string literal color"),
				"number":      makeToken(t.CodeNumber(), "Code number literal color"),
				"function":    makeToken(t.CodeFunction(), "Code function name color"),
				"operator":    makeToken(t.CodeOperator(), "Code operator color"),
				"punctuation": makeToken(t.CodePunctuation(), "Code punctuation color"),
				"variable":    makeToken(t.CodeVariable(), "Code variable color"),
				"constant":    makeToken(t.CodeConstant(), "Code constant color"),
				"type":        makeToken(t.CodeType(), "Code type name color"),
			},
		},
		"meta": map[string]interface{}{
			"id": map[string]interface{}{
				"$value": t.ID(),
				"$type":  "string",
			},
			"name": map[string]interface{}{
				"$value": t.DisplayName(),
				"$type":  "string",
			},
			"description": map[string]interface{}{
				"$value": t.Description(),
				"$type":  "string",
			},
			"author": map[string]interface{}{
				"$value": t.Author(),
				"$type":  "string",
			},
			"license": map[string]interface{}{
				"$value": t.License(),
				"$type":  "string",
			},
			"source": map[string]interface{}{
				"$value": t.Source(),
				"$type":  "string",
			},
			"isDark": map[string]interface{}{
				"$value": t.IsDark(),
				"$type":  "boolean",
			},
		},
	}

	return tokens
}

// GenerateAllDesignTokens generates design tokens for multiple themes.
func GenerateAllDesignTokens(themes []Theme, opts TokenOptions) (string, error) {
	if opts.Indent == "" {
		opts.Indent = "  "
	}

	allTokens := make(map[string]interface{})
	allTokens["$description"] = "Design tokens collection"

	for _, t := range themes {
		allTokens[t.ID()] = buildTokenStructure(t, opts)
	}

	var data []byte
	var err error

	if opts.Indent != "" {
		data, err = json.MarshalIndent(allTokens, "", opts.Indent)
	} else {
		data, err = json.Marshal(allTokens)
	}

	if err != nil {
		return "", fmt.Errorf("failed to marshal design tokens: %w", err)
	}

	return string(data), nil
}
