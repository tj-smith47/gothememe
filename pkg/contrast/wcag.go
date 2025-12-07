// Package contrast provides WCAG 2.1 compliant contrast ratio calculations
// and accessibility validation utilities.
package contrast

import (
	"github.com/tj-smith47/gothememe/internal/colorutil"
)

// MinAA is the minimum contrast ratio for WCAG AA compliance (normal text).
const MinAA = 4.5

// MinAALarge is the minimum contrast ratio for WCAG AA compliance (large text).
const MinAALarge = 3.0

// MinAAA is the minimum contrast ratio for WCAG AAA compliance (normal text).
const MinAAA = 7.0

// MinAAALarge is the minimum contrast ratio for WCAG AAA compliance (large text).
const MinAAALarge = 4.5

// MinUIComponent is the minimum contrast ratio for UI components and graphical objects.
const MinUIComponent = 3.0

// Luminance calculates the relative luminance of a color according to WCAG 2.1.
// RGB values should be in the range 0-255.
// Returns a value between 0 (black) and 1 (white).
func Luminance(r, g, b uint8) float64 {
	return colorutil.RelativeLuminance(r, g, b)
}

// LuminanceHex calculates the relative luminance from a hex color string.
// Accepts formats with or without # prefix (e.g., "#RRGGBB" or "RRGGBB").
func LuminanceHex(hex string) float64 {
	return colorutil.LuminanceHex(hex)
}

// Ratio calculates the contrast ratio between two colors.
// RGB values should be in the range 0-255.
// Returns a value between 1 (same color) and 21 (black/white).
func Ratio(r1, g1, b1, r2, g2, b2 uint8) float64 {
	return colorutil.ContrastRatio(r1, g1, b1, r2, g2, b2)
}

// RatioHex calculates the contrast ratio between two hex colors.
func RatioHex(hex1, hex2 string) float64 {
	return colorutil.ContrastRatioHex(hex1, hex2)
}

// MeetsAA checks if two colors meet WCAG AA contrast requirements.
// Set largeText to true for 18pt+ or 14pt+ bold text.
func MeetsAA(r1, g1, b1, r2, g2, b2 uint8, largeText bool) bool {
	ratio := Ratio(r1, g1, b1, r2, g2, b2)
	if largeText {
		return ratio >= MinAALarge
	}
	return ratio >= MinAA
}

// MeetsAAHex checks if two hex colors meet WCAG AA contrast requirements.
func MeetsAAHex(hex1, hex2 string, largeText bool) bool {
	ratio := RatioHex(hex1, hex2)
	if largeText {
		return ratio >= MinAALarge
	}
	return ratio >= MinAA
}

// MeetsAAA checks if two colors meet WCAG AAA contrast requirements.
// Set largeText to true for 18pt+ or 14pt+ bold text.
func MeetsAAA(r1, g1, b1, r2, g2, b2 uint8, largeText bool) bool {
	ratio := Ratio(r1, g1, b1, r2, g2, b2)
	if largeText {
		return ratio >= MinAAALarge
	}
	return ratio >= MinAAA
}

// MeetsAAAHex checks if two hex colors meet WCAG AAA contrast requirements.
func MeetsAAAHex(hex1, hex2 string, largeText bool) bool {
	ratio := RatioHex(hex1, hex2)
	if largeText {
		return ratio >= MinAAALarge
	}
	return ratio >= MinAAA
}

// MeetsUIComponent checks if a color meets WCAG 2.1 UI component contrast requirements.
// This applies to UI components and graphical objects (3:1 minimum).
func MeetsUIComponent(r1, g1, b1, r2, g2, b2 uint8) bool {
	return Ratio(r1, g1, b1, r2, g2, b2) >= MinUIComponent
}

// MeetsUIComponentHex checks if two hex colors meet UI component contrast requirements.
func MeetsUIComponentHex(hex1, hex2 string) bool {
	return RatioHex(hex1, hex2) >= MinUIComponent
}

// Level represents a WCAG compliance level.
type Level int

const (
	// LevelFail indicates the colors do not meet any WCAG contrast requirements.
	LevelFail Level = iota
	// LevelAALarge indicates colors meet AA requirements for large text only.
	LevelAALarge
	// LevelAA indicates colors meet AA requirements for all text.
	LevelAA
	// LevelAAALarge indicates colors meet AAA requirements for large text only.
	LevelAAALarge
	// LevelAAA indicates colors meet AAA requirements for all text.
	LevelAAA
)

// String returns a human-readable string for the compliance level.
func (l Level) String() string {
	switch l {
	case LevelAAA:
		return "AAA"
	case LevelAAALarge:
		return "AAA (large text only)"
	case LevelAA:
		return "AA"
	case LevelAALarge:
		return "AA (large text only)"
	default:
		return "Fail"
	}
}

// Check returns the highest WCAG compliance level achieved by the color pair.
func Check(r1, g1, b1, r2, g2, b2 uint8) Level {
	ratio := Ratio(r1, g1, b1, r2, g2, b2)
	return levelFromRatio(ratio)
}

// CheckHex returns the highest WCAG compliance level achieved by two hex colors.
func CheckHex(hex1, hex2 string) Level {
	ratio := RatioHex(hex1, hex2)
	return levelFromRatio(ratio)
}

// levelFromRatio determines the compliance level from a contrast ratio.
func levelFromRatio(ratio float64) Level {
	switch {
	case ratio >= MinAAA:
		return LevelAAA
	case ratio >= MinAAALarge:
		return LevelAAALarge
	case ratio >= MinAA:
		return LevelAA
	case ratio >= MinAALarge:
		return LevelAALarge
	default:
		return LevelFail
	}
}

// Issue represents a contrast accessibility issue.
type Issue struct {
	ForegroundName string
	BackgroundName string
	Foreground     string // hex color
	Background     string // hex color
	Ratio          float64
	Level          Level
	Required       Level
}
