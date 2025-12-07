package gothememe

import (
	"errors"
	"fmt"
	"strings"

	"github.com/tj-smith47/gothememe/internal/pairs"
	"github.com/tj-smith47/gothememe/pkg/contrast"
)

// ValidationSeverity indicates how critical a validation issue is.
type ValidationSeverity string

const (
	// SeverityError indicates a critical issue that should be fixed.
	SeverityError ValidationSeverity = "error"
	// SeverityWarning indicates a non-critical issue that may affect usability.
	SeverityWarning ValidationSeverity = "warning"
)

// ValidationError represents a theme validation failure.
type ValidationError struct {
	Field    string             // The field or color pair that has the issue
	Message  string             // Description of the issue
	Severity ValidationSeverity // How critical the issue is
}

// Error implements the error interface.
func (e ValidationError) Error() string {
	return fmt.Sprintf("[%s] %s: %s", e.Severity, e.Field, e.Message)
}

// ContrastLevel specifies the WCAG compliance level to check against.
type ContrastLevel int

const (
	// ContrastLevelAA requires 4.5:1 for normal text, 3:1 for large text.
	ContrastLevelAA ContrastLevel = iota
	// ContrastLevelAAA requires 7:1 for normal text, 4.5:1 for large text.
	ContrastLevelAAA
)

// ContrastIssue represents a contrast ratio failure.
type ContrastIssue struct {
	Background     string  // Background color hex
	Foreground     string  // Foreground color hex
	Ratio          float64 // Actual contrast ratio
	RequiredRatio  float64 // Minimum required ratio
	Level          string  // WCAG level (AA, AAA)
	BackgroundName string  // Name of the background color
	ForegroundName string  // Name of the foreground color
}

// Error returns a formatted error message.
func (i ContrastIssue) Error() string {
	return fmt.Sprintf("%s on %s: %.2f:1 (requires %.1f:1 for %s)",
		i.ForegroundName, i.BackgroundName, i.Ratio, i.RequiredRatio, i.Level)
}

// ValidateTheme checks a theme for common issues.
// Returns a slice of validation errors found.
func ValidateTheme(t Theme) []ValidationError {
	var errs []ValidationError

	// Check required fields
	if t.ID() == "" {
		errs = append(errs, ValidationError{
			Field:    "ID",
			Message:  "theme ID is required",
			Severity: SeverityError,
		})
	}

	if t.DisplayName() == "" {
		errs = append(errs, ValidationError{
			Field:    "DisplayName",
			Message:  "display name is required",
			Severity: SeverityError,
		})
	}

	// Check for empty required colors
	if t.Background().IsEmpty() {
		errs = append(errs, ValidationError{
			Field:    "Background",
			Message:  "background color is required",
			Severity: SeverityError,
		})
	}

	if t.TextPrimary().IsEmpty() {
		errs = append(errs, ValidationError{
			Field:    "TextPrimary",
			Message:  "primary text color is required",
			Severity: SeverityError,
		})
	}

	// Warn about transparent ANSI colors
	ansiColors := []struct {
		name  string
		color Color
	}{
		{"Black", t.Black()},
		{"Red", t.Red()},
		{"Green", t.Green()},
		{"Blue", t.Blue()},
		{"White", t.White()},
	}

	for _, c := range ansiColors {
		if c.color.IsEmpty() || c.color.Hex() == "transparent" {
			errs = append(errs, ValidationError{
				Field:    c.name,
				Message:  "ANSI color is empty or transparent",
				Severity: SeverityWarning,
			})
		}
	}

	// Check for mode/color mismatch
	if !t.Background().IsEmpty() && !t.TextPrimary().IsEmpty() {
		bgLum := contrast.LuminanceHex(t.Background().Hex())
		textLum := contrast.LuminanceHex(t.TextPrimary().Hex())

		isDarkBg := bgLum < 0.5
		isDarkText := textLum < 0.5

		if t.IsDark() && !isDarkBg {
			errs = append(errs, ValidationError{
				Field:    "IsDark",
				Message:  "theme is marked as dark but has a light background",
				Severity: SeverityWarning,
			})
		}

		if !t.IsDark() && isDarkBg {
			errs = append(errs, ValidationError{
				Field:    "IsDark",
				Message:  "theme is marked as light but has a dark background",
				Severity: SeverityWarning,
			})
		}

		if isDarkBg == isDarkText {
			errs = append(errs, ValidationError{
				Field:    "TextPrimary",
				Message:  "text and background have similar luminance, may be hard to read",
				Severity: SeverityWarning,
			})
		}
	}

	return errs
}

// ValidateContrast checks all standard color pairs for WCAG compliance.
// Returns a slice of contrast issues found.
func ValidateContrast(t Theme, level ContrastLevel) []ContrastIssue {
	var issues []ContrastIssue

	var minRatio float64
	var levelName string

	switch level {
	case ContrastLevelAAA:
		minRatio = contrast.MinAAA
		levelName = "AAA"
	default:
		minRatio = contrast.MinAA
		levelName = "AA"
	}

	// Get color pairs using shared pair specifications
	colorPairs := getColorPairsFromTheme(t)

	for _, pair := range colorPairs {
		// Skip if either color is empty
		if pair.FgHex == "" || pair.BgHex == "" {
			continue
		}

		ratio := contrast.RatioHex(pair.FgHex, pair.BgHex)

		if ratio < minRatio {
			issues = append(issues, ContrastIssue{
				Background:     pair.BgHex,
				Foreground:     pair.FgHex,
				Ratio:          ratio,
				RequiredRatio:  minRatio,
				Level:          levelName,
				BackgroundName: pair.BgName,
				ForegroundName: pair.FgName,
			})
		}
	}

	return issues
}

// getColorPairsFromTheme extracts color pairs from a theme based on standard pair specs.
func getColorPairsFromTheme(t Theme) []pairs.ColorPair {
	specs := pairs.StandardPairSpecs()
	result := make([]pairs.ColorPair, 0, len(specs))

	for _, spec := range specs {
		fg := getThemeColor(t, spec.FgName)
		bg := getThemeColor(t, spec.BgName)
		result = append(result, pairs.ColorPair{
			FgName: spec.FgName,
			BgName: spec.BgName,
			FgHex:  fg.Hex(),
			BgHex:  bg.Hex(),
		})
	}

	return result
}

// colorGetter is a function that extracts a color from a Theme.
type colorGetter func(Theme) Color

// colorGetters maps color names to their extraction functions.
var colorGetters = map[string]colorGetter{
	"Background":          Theme.Background,
	"BackgroundSecondary": Theme.BackgroundSecondary,
	"Surface":             Theme.Surface,
	"SurfaceSecondary":    Theme.SurfaceSecondary,
	"TextPrimary":         Theme.TextPrimary,
	"TextSecondary":       Theme.TextSecondary,
	"TextMuted":           Theme.TextMuted,
	"TextInverted":        Theme.TextInverted,
	"Accent":              Theme.Accent,
	"AccentSecondary":     Theme.AccentSecondary,
	"Brand":               Theme.Brand,
	"Border":              Theme.Border,
	"BorderSubtle":        Theme.BorderSubtle,
	"BorderStrong":        Theme.BorderStrong,
	"CodeText":            Theme.CodeText,
	"CodeBackground":      Theme.CodeBackground,
	"CodeComment":         Theme.CodeComment,
	"CodeKeyword":         Theme.CodeKeyword,
	"CodeString":          Theme.CodeString,
}

// semanticColorGetters maps semantic color names to their extraction functions.
var semanticColorGetters = map[string]func(Theme) Color{
	"Success.Text":       func(t Theme) Color { return t.Success().Text },
	"Success.Background": func(t Theme) Color { return t.Success().Background },
	"Warning.Text":       func(t Theme) Color { return t.Warning().Text },
	"Warning.Background": func(t Theme) Color { return t.Warning().Background },
	"Error.Text":         func(t Theme) Color { return t.Error().Text },
	"Error.Background":   func(t Theme) Color { return t.Error().Background },
	"Info.Text":          func(t Theme) Color { return t.Info().Text },
	"Info.Background":    func(t Theme) Color { return t.Info().Background },
}

// getThemeColor retrieves a color from a theme by name.
func getThemeColor(t Theme, name string) Color {
	if getter, ok := colorGetters[name]; ok {
		return getter(t)
	}
	if getter, ok := semanticColorGetters[name]; ok {
		return getter(t)
	}
	return Color{}
}

// ValidateStrict performs all validations and returns an error if any fail.
// This is useful for CI/CD pipelines or strict mode checking.
func ValidateStrict(t Theme) error {
	return validateStrictInternal(t, ContrastLevelAA)
}

// ValidateStrictAAA performs strict AAA-level validation.
// Returns an error if any validation fails.
func ValidateStrictAAA(t Theme) error {
	return validateStrictInternal(t, ContrastLevelAAA)
}

// validateStrictInternal performs validation at the specified contrast level.
func validateStrictInternal(t Theme, level ContrastLevel) error {
	validationErrs := ValidateTheme(t)
	contrastIssues := ValidateContrast(t, level)

	var errMsgs []string

	// Collect validation errors (only actual errors, not warnings)
	for _, err := range validationErrs {
		if err.Severity == SeverityError {
			errMsgs = append(errMsgs, err.Error())
		}
	}

	// Collect contrast issues
	for _, issue := range contrastIssues {
		errMsgs = append(errMsgs, issue.Error())
	}

	if len(errMsgs) == 0 {
		return nil
	}

	return errors.New(strings.Join(errMsgs, "; "))
}

// AutoFixContrast creates a new theme with colors adjusted to meet WCAG contrast requirements.
// It returns a theme with the same ID (suffixed with "-fixed") where problematic color pairs
// have been adjusted to meet the specified contrast level.
//
// The function prioritizes adjusting foreground colors while preserving the overall
// theme aesthetic. For dark themes, foreground colors are lightened; for light themes,
// they are darkened.
func AutoFixContrast(t Theme, level ContrastLevel) Theme {
	issues := ValidateContrast(t, level)
	if len(issues) == 0 {
		// No issues, return a copy with the same colors
		return copyTheme(t, t.ID(), t.DisplayName())
	}

	// Build new theme with fixed colors
	builder := NewThemeBuilder(t.ID()+"-fixed", t.DisplayName()+" (Fixed)").
		WithDescription(t.Description() + " - WCAG contrast adjusted").
		WithAuthor(t.Author()).
		WithLicense(t.License()).
		WithSource(t.Source()).
		WithIsDark(t.IsDark())

	// Copy all original colors first
	copyAllColors(builder, t)

	// Track which colors we've already fixed
	fixed := make(map[string]Color)

	// Calculate required ratio based on level
	requiredRatio := 4.5 // AA default
	if level == ContrastLevelAAA {
		requiredRatio = 7.0
	}

	// Fix each issue
	for _, issue := range issues {
		// Skip if we've already fixed this foreground color
		if _, ok := fixed[issue.ForegroundName]; ok {
			continue
		}

		fg := Hex(issue.Foreground)
		bg := Hex(issue.Background)

		// Adjust the foreground color to meet contrast requirements
		adjustedFg := adjustColorForContrast(fg, bg, requiredRatio, t.IsDark())
		fixed[issue.ForegroundName] = adjustedFg

		// Apply the fixed color
		applyColorFix(builder, issue.ForegroundName, adjustedFg)
	}

	return builder.Build()
}

// copyTheme creates a copy of a theme with a new ID and name.
func copyTheme(t Theme, id, name string) Theme {
	builder := NewThemeBuilder(id, name).
		WithDescription(t.Description()).
		WithAuthor(t.Author()).
		WithLicense(t.License()).
		WithSource(t.Source()).
		WithIsDark(t.IsDark())

	copyAllColors(builder, t)
	return builder.Build()
}

// copyAllColors copies all colors from a theme to a builder.
func copyAllColors(builder *ThemeBuilder, t Theme) {
	builder.
		WithBackground(t.Background()).
		WithBackgroundSecondary(t.BackgroundSecondary()).
		WithSurface(t.Surface()).
		WithSurfaceSecondary(t.SurfaceSecondary()).
		WithTextPrimary(t.TextPrimary()).
		WithTextSecondary(t.TextSecondary()).
		WithTextMuted(t.TextMuted()).
		WithTextInverted(t.TextInverted()).
		WithAccent(t.Accent()).
		WithAccentSecondary(t.AccentSecondary()).
		WithBrand(t.Brand()).
		WithBorder(t.Border()).
		WithBorderSubtle(t.BorderSubtle()).
		WithBorderStrong(t.BorderStrong()).
		WithSuccess(t.Success()).
		WithWarning(t.Warning()).
		WithError(t.Error()).
		WithInfo(t.Info()).
		WithBlack(t.Black()).
		WithRed(t.Red()).
		WithGreen(t.Green()).
		WithYellow(t.Yellow()).
		WithBlue(t.Blue()).
		WithPurple(t.Purple()).
		WithCyan(t.Cyan()).
		WithWhite(t.White()).
		WithBrightBlack(t.BrightBlack()).
		WithBrightRed(t.BrightRed()).
		WithBrightGreen(t.BrightGreen()).
		WithBrightYellow(t.BrightYellow()).
		WithBrightBlue(t.BrightBlue()).
		WithBrightPurple(t.BrightPurple()).
		WithBrightCyan(t.BrightCyan()).
		WithBrightWhite(t.BrightWhite()).
		WithCodeBackground(t.CodeBackground()).
		WithCodeText(t.CodeText()).
		WithCodeComment(t.CodeComment()).
		WithCodeKeyword(t.CodeKeyword()).
		WithCodeString(t.CodeString()).
		WithCodeNumber(t.CodeNumber()).
		WithCodeFunction(t.CodeFunction()).
		WithCodeOperator(t.CodeOperator()).
		WithCodePunctuation(t.CodePunctuation()).
		WithCodeVariable(t.CodeVariable()).
		WithCodeConstant(t.CodeConstant()).
		WithCodeType(t.CodeType())
}

// adjustColorForContrast adjusts a foreground color to meet the required contrast ratio.
func adjustColorForContrast(fg, bg Color, requiredRatio float64, isDark bool) Color {
	fgHex := fg.Hex()
	bgHex := bg.Hex()

	// Check if already meets requirement
	currentRatio := contrast.RatioHex(fgHex, bgHex)
	if currentRatio >= requiredRatio {
		return fg
	}

	// Adjust in small steps until we meet the requirement
	adjusted := fg
	maxIterations := 50 // Prevent infinite loops

	for i := 0; i < maxIterations; i++ {
		currentRatio = contrast.RatioHex(adjusted.Hex(), bgHex)
		if currentRatio >= requiredRatio {
			break
		}

		// For dark themes, lighten the foreground; for light themes, darken it
		if isDark {
			adjusted = adjusted.Lighten(0.05)
		} else {
			adjusted = adjusted.Darken(0.05)
		}
	}

	return adjusted
}

// applyColorFix applies a fixed color to the appropriate builder method.
func applyColorFix(builder *ThemeBuilder, colorName string, color Color) {
	// Map color names to builder methods
	colorFixers := map[string]func(*ThemeBuilder, Color) *ThemeBuilder{
		"TextPrimary":     (*ThemeBuilder).WithTextPrimary,
		"TextSecondary":   (*ThemeBuilder).WithTextSecondary,
		"TextMuted":       (*ThemeBuilder).WithTextMuted,
		"TextInverted":    (*ThemeBuilder).WithTextInverted,
		"Accent":          (*ThemeBuilder).WithAccent,
		"AccentSecondary": (*ThemeBuilder).WithAccentSecondary,
		"Brand":           (*ThemeBuilder).WithBrand,
		"Border":          (*ThemeBuilder).WithBorder,
		"BorderSubtle":    (*ThemeBuilder).WithBorderSubtle,
		"BorderStrong":    (*ThemeBuilder).WithBorderStrong,
		"CodeText":        (*ThemeBuilder).WithCodeText,
		"CodeComment":     (*ThemeBuilder).WithCodeComment,
		"CodeKeyword":     (*ThemeBuilder).WithCodeKeyword,
		"CodeString":      (*ThemeBuilder).WithCodeString,
	}

	if fixer, ok := colorFixers[colorName]; ok {
		fixer(builder, color)
	}
}
