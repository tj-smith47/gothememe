package gothememe

import (
	"cmp"
	"slices"

	"github.com/tj-smith47/gothememe/pkg/contrast"
)

// ThemeStats provides analysis of a theme's color usage and accessibility.
type ThemeStats struct {
	// ColorCount is the total number of defined colors.
	ColorCount int

	// UniqueColors is the number of unique color values.
	UniqueColors int

	// ContrastScore is the average contrast ratio of text/background pairs (0-21).
	ContrastScore float64

	// AccessiblePairs is the number of color pairs meeting AA requirements.
	AccessiblePairs int

	// TotalPairs is the total number of color pairs checked.
	TotalPairs int

	// AccessibilityPercent is the percentage of pairs meeting AA (0-100).
	AccessibilityPercent float64

	// IsDark indicates if the theme is classified as dark.
	IsDark bool

	// AverageTextLuminance is the average luminance of text colors.
	AverageTextLuminance float64

	// BackgroundLuminance is the luminance of the primary background.
	BackgroundLuminance float64
}

// AnalyzeTheme returns statistics about a theme's colors and accessibility.
func AnalyzeTheme(t Theme) ThemeStats {
	stats := ThemeStats{
		IsDark: t.IsDark(),
	}

	// Collect all colors
	colors := collectColors(t)
	stats.ColorCount = len(colors)
	stats.UniqueColors = countUnique(colors)

	// Calculate background luminance
	if !t.Background().IsEmpty() {
		stats.BackgroundLuminance = contrast.LuminanceHex(t.Background().Hex())
	}

	// Calculate text luminance average
	textColors := []Color{
		t.TextPrimary(),
		t.TextSecondary(),
		t.TextMuted(),
	}
	stats.AverageTextLuminance = averageLuminance(textColors)

	// Analyze contrast pairs
	pairs := getContrastPairs(t)
	stats.TotalPairs = len(pairs)

	var totalRatio float64
	for _, pair := range pairs {
		if pair.fg.IsEmpty() || pair.bg.IsEmpty() {
			continue
		}
		ratio := contrast.RatioHex(pair.fg.Hex(), pair.bg.Hex())
		totalRatio += ratio
		if ratio >= contrast.MinAA {
			stats.AccessiblePairs++
		}
	}

	if stats.TotalPairs > 0 {
		stats.ContrastScore = totalRatio / float64(stats.TotalPairs)
		stats.AccessibilityPercent = float64(stats.AccessiblePairs) / float64(stats.TotalPairs) * 100
	}

	return stats
}

// colorPair represents a foreground/background color combination.
type colorPair struct {
	fg, bg Color
}

// collectColors gathers all defined colors from a theme.
func collectColors(t Theme) []Color {
	return []Color{
		t.Background(),
		t.BackgroundSecondary(),
		t.Surface(),
		t.SurfaceSecondary(),
		t.TextPrimary(),
		t.TextSecondary(),
		t.TextMuted(),
		t.TextInverted(),
		t.Accent(),
		t.AccentSecondary(),
		t.Brand(),
		t.Border(),
		t.BorderSubtle(),
		t.BorderStrong(),
		t.Success().Background,
		t.Success().Border,
		t.Success().Text,
		t.Warning().Background,
		t.Warning().Border,
		t.Warning().Text,
		t.Error().Background,
		t.Error().Border,
		t.Error().Text,
		t.Info().Background,
		t.Info().Border,
		t.Info().Text,
		t.Black(),
		t.Red(),
		t.Green(),
		t.Yellow(),
		t.Blue(),
		t.Purple(),
		t.Cyan(),
		t.White(),
		t.BrightBlack(),
		t.BrightRed(),
		t.BrightGreen(),
		t.BrightYellow(),
		t.BrightBlue(),
		t.BrightPurple(),
		t.BrightCyan(),
		t.BrightWhite(),
		t.CodeBackground(),
		t.CodeText(),
		t.CodeComment(),
		t.CodeKeyword(),
		t.CodeString(),
		t.CodeNumber(),
		t.CodeFunction(),
		t.CodeOperator(),
		t.CodePunctuation(),
		t.CodeVariable(),
		t.CodeConstant(),
		t.CodeType(),
	}
}

// countUnique counts the number of unique, non-empty color values.
func countUnique(colors []Color) int {
	seen := make(map[string]struct{})
	for _, c := range colors {
		if !c.IsEmpty() {
			seen[c.Hex()] = struct{}{}
		}
	}
	return len(seen)
}

// averageLuminance calculates the average luminance of non-empty colors.
func averageLuminance(colors []Color) float64 {
	var total float64
	var count int
	for _, c := range colors {
		if !c.IsEmpty() {
			total += contrast.LuminanceHex(c.Hex())
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return total / float64(count)
}

// getContrastPairs returns all foreground/background pairs to check.
// Uses the shared pair specifications from internal/pairs.
func getContrastPairs(t Theme) []colorPair {
	sharedPairs := getColorPairsFromTheme(t)
	result := make([]colorPair, 0, len(sharedPairs))
	for _, p := range sharedPairs {
		result = append(result, colorPair{
			fg: Hex(p.FgHex),
			bg: Hex(p.BgHex),
		})
	}
	return result
}

// CompareThemes returns a comparison of two themes.
func CompareThemes(a, b Theme) ThemeComparison {
	statsA := AnalyzeTheme(a)
	statsB := AnalyzeTheme(b)

	return ThemeComparison{
		ThemeA:         a.ID(),
		ThemeB:         b.ID(),
		StatsA:         statsA,
		StatsB:         statsB,
		ContrastDiff:   statsB.ContrastScore - statsA.ContrastScore,
		AccessDiff:     statsB.AccessibilityPercent - statsA.AccessibilityPercent,
		UniqueDiff:     statsB.UniqueColors - statsA.UniqueColors,
		SameDarkMode:   statsA.IsDark == statsB.IsDark,
		MoreAccessible: b.ID(),
	}
}

// ThemeComparison holds the result of comparing two themes.
type ThemeComparison struct {
	ThemeA         string
	ThemeB         string
	StatsA         ThemeStats
	StatsB         ThemeStats
	ContrastDiff   float64 // Positive means B has higher contrast
	AccessDiff     float64 // Positive means B is more accessible
	UniqueDiff     int     // Positive means B has more unique colors
	SameDarkMode   bool    // True if both themes have same dark/light mode
	MoreAccessible string  // ID of the more accessible theme
}

// init sets MoreAccessible based on accessibility percentages.
func (c *ThemeComparison) init() {
	if c.StatsA.AccessibilityPercent >= c.StatsB.AccessibilityPercent {
		c.MoreAccessible = c.ThemeA
	}
}

// AnalyzeAll returns stats for all provided themes.
func AnalyzeAll(themes []Theme) []ThemeStats {
	stats := make([]ThemeStats, len(themes))
	for i, t := range themes {
		stats[i] = AnalyzeTheme(t)
	}
	return stats
}

// FilterAccessible returns themes meeting the specified accessibility level.
func FilterAccessible(themes []Theme, minPercent float64) []Theme {
	var accessible []Theme
	for _, t := range themes {
		stats := AnalyzeTheme(t)
		if stats.AccessibilityPercent >= minPercent {
			accessible = append(accessible, t)
		}
	}
	return accessible
}

// themeWithStats pairs a theme with its cached stats for efficient sorting.
type themeWithStats struct {
	theme Theme
	stats ThemeStats
}

// SortByAccessibility returns themes sorted by accessibility percentage (highest first).
// Uses slices.SortFunc with cached stats for O(n log n) performance.
func SortByAccessibility(themes []Theme) []Theme {
	if len(themes) == 0 {
		return themes
	}

	// Cache stats to avoid recalculation during sorting
	wrapped := make([]themeWithStats, len(themes))
	for i, t := range themes {
		wrapped[i] = themeWithStats{theme: t, stats: AnalyzeTheme(t)}
	}

	// Sort by accessibility percentage (descending)
	slices.SortFunc(wrapped, func(a, b themeWithStats) int {
		return cmp.Compare(b.stats.AccessibilityPercent, a.stats.AccessibilityPercent)
	})

	// Extract sorted themes
	result := make([]Theme, len(wrapped))
	for i, w := range wrapped {
		result[i] = w.theme
	}
	return result
}
