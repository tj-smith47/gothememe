package gothememe

import (
	"fmt"
	"image/color"
	"math"
	"regexp"
	"strings"

	"github.com/lucasb-eyer/go-colorful"

	"github.com/tj-smith47/gothememe/internal/colorutil"
)

// Color represents a theme color with various output and manipulation methods.
// Colors are stored internally as hex values and can be converted to multiple
// color spaces including RGB, HSL, and OKLCH.
type Color struct {
	value string // hex value without # prefix
}

// SemanticColor represents a set of related colors for semantic states
// like success, warning, error, and info. Each semantic color includes
// a background, border, and text color for consistent UI implementation.
type SemanticColor struct {
	Background Color
	Border     Color
	Text       Color
}

var hexPattern = regexp.MustCompile(`^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6}|[0-9a-fA-F]{8})$`)

// Hex creates a Color from a hex string.
// Accepts formats: "#RGB", "#RRGGBB", "#RRGGBBAA", "RGB", "RRGGBB", "RRGGBBAA".
// Returns an empty color if the input is invalid.
func Hex(hex string) Color {
	hex = strings.TrimPrefix(hex, "#")
	if !hexPattern.MatchString(hex) {
		return Color{}
	}

	// Expand shorthand hex (RGB -> RRGGBB)
	if len(hex) == 3 {
		hex = string([]byte{hex[0], hex[0], hex[1], hex[1], hex[2], hex[2]})
	}

	return Color{value: strings.ToLower(hex)}
}

// RGB creates a Color from red, green, and blue components (0-255).
func RGB(r, g, b uint8) Color {
	return Color{value: fmt.Sprintf("%02x%02x%02x", r, g, b)}
}

// RGBA creates a Color from red, green, blue, and alpha components (0-255).
func RGBA(r, g, b, a uint8) Color {
	return Color{value: fmt.Sprintf("%02x%02x%02x%02x", r, g, b, a)}
}

// HSL creates a Color from hue (0-360), saturation (0-1), and lightness (0-1).
func HSL(h, s, l float64) Color {
	c := colorful.Hsl(h, s, l)
	return Color{value: strings.TrimPrefix(c.Hex(), "#")}
}

// OKLCH creates a Color from OKLCH color space values.
// L is lightness (0-1), C is chroma (typically 0-0.4), H is hue (0-360).
func OKLCH(l, c, h float64) Color {
	// Convert OKLCH to sRGB via Lab
	col := colorful.LuvLCh(l*100, c*100, h)
	return Color{value: strings.TrimPrefix(col.Clamped().Hex(), "#")}
}

// IsEmpty returns true if the color has no value.
func (c Color) IsEmpty() bool {
	return c.value == ""
}

// Hex returns the color as a hex string with # prefix.
func (c Color) Hex() string {
	if c.value == "" {
		return ""
	}
	return "#" + c.value
}

// HexNoPrefix returns the color as a hex string without # prefix.
func (c Color) HexNoPrefix() string {
	return c.value
}

// RGB returns the red, green, and blue components (0-255).
func (c Color) RGB() (r, g, b uint8) {
	return colorutil.HexToRGB(c.value)
}

// RGBAComponents returns the red, green, blue, and alpha components (0-255).
// If no alpha is specified, returns 255 (fully opaque).
func (c Color) RGBAComponents() (r, g, b, a uint8) {
	r, g, b = c.RGB()
	a = 255
	if len(c.value) >= 8 {
		a = colorutil.ParseHexByte(c.value[6:8])
	}
	return r, g, b, a
}

// HSLValues returns the hue (0-360), saturation (0-1), and lightness (0-1).
func (c Color) HSLValues() (h, s, l float64) {
	col := c.colorful()
	return col.Hsl()
}

// OKLCHValues returns the OKLCH color space values.
// L is lightness (0-1), C is chroma (typically 0-0.4), H is hue (0-360).
func (c Color) OKLCHValues() (l, ch, h float64) {
	col := c.colorful()
	ll, cc, hh := col.LuvLCh()
	return ll / 100, cc / 100, hh
}

// CSS returns the color formatted for CSS.
// Returns the hex value by default.
func (c Color) CSS() string {
	return c.Hex()
}

// CSSVar returns the color as a CSS variable reference.
// Example: CSSVar("background") returns "var(--theme-background)"
func (c Color) CSSVar(name string) string {
	return fmt.Sprintf("var(--theme-%s)", name)
}

// CSSRGB returns the color as CSS rgb() function.
func (c Color) CSSRGB() string {
	r, g, b := c.RGB()
	return fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
}

// CSSRGBA returns the color as CSS rgba() function.
func (c Color) CSSRGBA() string {
	r, g, b, a := c.RGBAComponents()
	alpha := float64(a) / 255.0
	return fmt.Sprintf("rgba(%d, %d, %d, %.3f)", r, g, b, alpha)
}

// CSSHSL returns the color as CSS hsl() function.
func (c Color) CSSHSL() string {
	h, s, l := c.HSLValues()
	return fmt.Sprintf("hsl(%.1f, %.1f%%, %.1f%%)", h, s*100, l*100)
}

// WithAlpha returns a new color with the specified alpha value (0-1).
func (c Color) WithAlpha(alpha float64) Color {
	if alpha < 0 {
		alpha = 0
	}
	if alpha > 1 {
		alpha = 1
	}
	r, g, b := c.RGB()
	a := uint8(alpha * 255)
	return RGBA(r, g, b, a)
}

// Lighten returns a new color lightened by the specified amount (0-1).
func (c Color) Lighten(amount float64) Color {
	h, s, l := c.HSLValues()
	l = math.Min(1.0, l+amount)
	return HSL(h, s, l)
}

// Darken returns a new color darkened by the specified amount (0-1).
func (c Color) Darken(amount float64) Color {
	h, s, l := c.HSLValues()
	l = math.Max(0.0, l-amount)
	return HSL(h, s, l)
}

// Saturate returns a new color with increased saturation by the specified amount (0-1).
func (c Color) Saturate(amount float64) Color {
	h, s, l := c.HSLValues()
	s = math.Min(1.0, s+amount)
	return HSL(h, s, l)
}

// Desaturate returns a new color with decreased saturation by the specified amount (0-1).
func (c Color) Desaturate(amount float64) Color {
	h, s, l := c.HSLValues()
	s = math.Max(0.0, s-amount)
	return HSL(h, s, l)
}

// Mix blends two colors together. The ratio parameter controls the mix:
// 0.0 = entirely this color, 1.0 = entirely the other color, 0.5 = equal mix.
func (c Color) Mix(other Color, ratio float64) Color {
	c1 := c.colorful()
	c2 := other.colorful()
	mixed := c1.BlendLab(c2, ratio)
	return Color{value: strings.TrimPrefix(mixed.Hex(), "#")}
}

// Complement returns the complementary color (opposite on the color wheel).
func (c Color) Complement() Color {
	h, s, l := c.HSLValues()
	h = math.Mod(h+180, 360)
	return HSL(h, s, l)
}

// Invert returns the inverted color.
func (c Color) Invert() Color {
	r, g, b := c.RGB()
	return RGB(255-r, 255-g, 255-b)
}

// RelativeLuminance calculates the relative luminance of the color
// according to WCAG 2.1 specification. Returns a value between 0 and 1.
func (c Color) RelativeLuminance() float64 {
	r, g, b := c.RGB()
	return colorutil.RelativeLuminance(r, g, b)
}

// IsDark returns true if the color is considered dark (luminance < 0.5).
func (c Color) IsDark() bool {
	return c.RelativeLuminance() < 0.5
}

// IsLight returns true if the color is considered light (luminance >= 0.5).
func (c Color) IsLight() bool {
	return !c.IsDark()
}

// String implements the Stringer interface.
func (c Color) String() string {
	return c.Hex()
}

// GoString implements the GoStringer interface for fmt %#v.
func (c Color) GoString() string {
	return fmt.Sprintf("gothememe.Hex(%q)", c.Hex())
}

// colorful converts to go-colorful Color type for advanced operations.
func (c Color) colorful() colorful.Color {
	if c.value == "" {
		return colorful.Color{}
	}
	col, err := colorful.Hex("#" + c.value)
	if err != nil {
		return colorful.Color{}
	}
	return col
}

// StdColor returns the color as a standard library color.Color.
func (c Color) StdColor() color.Color {
	r, g, b, a := c.RGBAComponents()
	return color.RGBA{R: r, G: g, B: b, A: a}
}

// FromStdColor creates a Color from a standard library color.Color.
func FromStdColor(c color.Color) Color {
	r, g, b, a := c.RGBA()
	// color.RGBA() returns 16-bit values, convert to 8-bit
	return RGBA(uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8))
}
