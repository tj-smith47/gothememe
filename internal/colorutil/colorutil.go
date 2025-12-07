// Package colorutil provides shared color manipulation utilities.
// This package consolidates color-related functions used across
// the gothememe package to avoid code duplication.
package colorutil

import (
	"math"
	"strconv"
	"strings"
)

// RelativeLuminance calculates the relative luminance of a color according
// to WCAG 2.1 specification. RGB values should be in the range 0-255.
// Returns a value between 0 (black) and 1 (white).
//
// The formula follows the W3C WCAG 2.1 guidelines:
// https://www.w3.org/TR/WCAG21/#dfn-relative-luminance
func RelativeLuminance(r, g, b uint8) float64 {
	// Convert to sRGB (0-1 range)
	rs := float64(r) / 255.0
	gs := float64(g) / 255.0
	bs := float64(b) / 255.0

	// Apply gamma correction (linearize sRGB)
	rLinear := linearize(rs)
	gLinear := linearize(gs)
	bLinear := linearize(bs)

	// Calculate relative luminance using ITU-R BT.709 coefficients
	return 0.2126*rLinear + 0.7152*gLinear + 0.0722*bLinear
}

// linearize applies gamma correction to a single sRGB channel value.
// This converts from sRGB to linear RGB space.
func linearize(v float64) float64 {
	if v <= 0.03928 {
		return v / 12.92
	}
	return math.Pow((v+0.055)/1.055, 2.4)
}

// HexToRGB converts a hex color string to RGB values.
// Accepts formats with or without # prefix: "#RRGGBB", "RRGGBB", "#RGB", "RGB".
// Returns (0, 0, 0) for invalid input.
func HexToRGB(hex string) (r, g, b uint8) {
	// Remove # prefix if present
	hex = strings.TrimPrefix(hex, "#")

	// Expand shorthand (RGB -> RRGGBB)
	if len(hex) == 3 {
		hex = string([]byte{hex[0], hex[0], hex[1], hex[1], hex[2], hex[2]})
	}

	if len(hex) < 6 {
		return 0, 0, 0
	}

	// Parse hex values
	rr, err := strconv.ParseUint(hex[0:2], 16, 8)
	if err != nil {
		return 0, 0, 0
	}
	gg, err := strconv.ParseUint(hex[2:4], 16, 8)
	if err != nil {
		return 0, 0, 0
	}
	bb, err := strconv.ParseUint(hex[4:6], 16, 8)
	if err != nil {
		return 0, 0, 0
	}

	return uint8(rr), uint8(gg), uint8(bb)
}

// LuminanceHex calculates the relative luminance from a hex color string.
// Accepts formats with or without # prefix.
func LuminanceHex(hex string) float64 {
	r, g, b := HexToRGB(hex)
	return RelativeLuminance(r, g, b)
}

// ContrastRatio calculates the contrast ratio between two colors.
// RGB values should be in the range 0-255.
// Returns a value between 1 (same color) and 21 (black/white).
func ContrastRatio(r1, g1, b1, r2, g2, b2 uint8) float64 {
	l1 := RelativeLuminance(r1, g1, b1)
	l2 := RelativeLuminance(r2, g2, b2)
	return RatioFromLuminance(l1, l2)
}

// ContrastRatioHex calculates the contrast ratio between two hex colors.
func ContrastRatioHex(hex1, hex2 string) float64 {
	l1 := LuminanceHex(hex1)
	l2 := LuminanceHex(hex2)
	return RatioFromLuminance(l1, l2)
}

// RatioFromLuminance calculates contrast ratio from two luminance values.
func RatioFromLuminance(l1, l2 float64) float64 {
	// Ensure l1 is the lighter color
	if l1 < l2 {
		l1, l2 = l2, l1
	}
	return (l1 + 0.05) / (l2 + 0.05)
}

// ParseHexByte parses a two-character hex string to a uint8.
// Returns 0 for invalid input.
func ParseHexByte(hex string) uint8 {
	if len(hex) < 2 {
		return 0
	}
	v, err := strconv.ParseUint(hex[0:2], 16, 8)
	if err != nil {
		return 0
	}
	return uint8(v)
}
