package gothememe

import (
	"testing"
)

func TestHex(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		input     string
		wantHex   string
		wantEmpty bool
	}{
		{"valid 6-char with hash", "#ff5555", "#ff5555", false},
		{"valid 6-char without hash", "ff5555", "#ff5555", false},
		{"valid 3-char with hash", "#f55", "#ff5555", false},
		{"valid 3-char without hash", "f55", "#ff5555", false},
		{"valid 8-char with alpha", "#ff555580", "#ff555580", false},
		{"uppercase", "#FF5555", "#ff5555", false},
		{"mixed case", "#Ff5555", "#ff5555", false},
		{"invalid empty", "", "", true},
		{"invalid short", "#ff", "", true},
		{"invalid chars", "#gggggg", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := Hex(tt.input)
			if tt.wantEmpty {
				if !c.IsEmpty() {
					t.Errorf("Hex(%q) = %q, want empty", tt.input, c.Hex())
				}
				return
			}
			if c.Hex() != tt.wantHex {
				t.Errorf("Hex(%q).Hex() = %q, want %q", tt.input, c.Hex(), tt.wantHex)
			}
		})
	}
}

func TestRGB(t *testing.T) {
	t.Parallel()
	c := RGB(255, 85, 85)
	if c.Hex() != "#ff5555" {
		t.Errorf("RGB(255, 85, 85).Hex() = %q, want #ff5555", c.Hex())
	}

	r, g, b := c.RGB()
	if r != 255 || g != 85 || b != 85 {
		t.Errorf("RGB() = (%d, %d, %d), want (255, 85, 85)", r, g, b)
	}
}

func TestColorManipulation(t *testing.T) {
	t.Parallel()
	c := Hex("#808080") // Gray

	// Test lighten
	lighter := c.Lighten(0.1)
	_, _, l1 := c.HSLValues()
	_, _, l2 := lighter.HSLValues()
	if l2 <= l1 {
		t.Errorf("Lighten() should increase lightness: %f -> %f", l1, l2)
	}

	// Test darken
	darker := c.Darken(0.1)
	_, _, l3 := darker.HSLValues()
	if l3 >= l1 {
		t.Errorf("Darken() should decrease lightness: %f -> %f", l1, l3)
	}

	// Test alpha
	alpha := c.WithAlpha(0.5)
	_, _, _, a := alpha.RGBAComponents()
	if a != 127 && a != 128 { // Allow for rounding
		t.Errorf("WithAlpha(0.5) alpha = %d, want ~127", a)
	}
}

func TestColorRelativeLuminance(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		hex  string
		want float64
		tol  float64
	}{
		{"white", "#ffffff", 1.0, 0.001},
		{"black", "#000000", 0.0, 0.001},
		{"red", "#ff0000", 0.2126, 0.001},
		{"green", "#00ff00", 0.7152, 0.001},
		{"blue", "#0000ff", 0.0722, 0.001},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := Hex(tt.hex)
			got := c.RelativeLuminance()
			if got < tt.want-tt.tol || got > tt.want+tt.tol {
				t.Errorf("RelativeLuminance() = %f, want %f (±%f)", got, tt.want, tt.tol)
			}
		})
	}
}

func TestColorIsDark(t *testing.T) {
	t.Parallel()
	tests := []struct {
		hex      string
		wantDark bool
	}{
		{"#000000", true},
		{"#282a36", true}, // Dracula background
		{"#ffffff", false},
		{"#f8f8f2", false}, // Dracula foreground
	}

	for _, tt := range tests {
		t.Run(tt.hex, func(t *testing.T) {
			t.Parallel()
			c := Hex(tt.hex)
			if c.IsDark() != tt.wantDark {
				t.Errorf("Hex(%q).IsDark() = %v, want %v", tt.hex, c.IsDark(), tt.wantDark)
			}
		})
	}
}

func TestSemanticColor(t *testing.T) {
	t.Parallel()
	sc := SemanticColor{
		Background: Hex("#50fa7b"),
		Border:     Hex("#50fa7b"),
		Text:       Hex("#50fa7b"),
	}

	if sc.Background.IsEmpty() {
		t.Error("SemanticColor.Background should not be empty")
	}
	if sc.Border.IsEmpty() {
		t.Error("SemanticColor.Border should not be empty")
	}
	if sc.Text.IsEmpty() {
		t.Error("SemanticColor.Text should not be empty")
	}
}

func TestRGBA(t *testing.T) {
	t.Parallel()
	c := RGBA(255, 85, 85, 128)
	if c.Hex() != "#ff555580" {
		t.Errorf("RGBA(255, 85, 85, 128).Hex() = %q, want #ff555580", c.Hex())
	}

	r, g, b, a := c.RGBAComponents()
	if r != 255 || g != 85 || b != 85 || a != 128 {
		t.Errorf("RGBAComponents() = (%d, %d, %d, %d), want (255, 85, 85, 128)", r, g, b, a)
	}
}

func TestHSL(t *testing.T) {
	t.Parallel()
	// Red at full saturation
	c := HSL(0, 1.0, 0.5)
	r, g, b := c.RGB()
	if r != 255 || g != 0 || b != 0 {
		t.Errorf("HSL(0, 1.0, 0.5).RGB() = (%d, %d, %d), want (255, 0, 0)", r, g, b)
	}
}

func TestOKLCH(t *testing.T) {
	t.Parallel()
	c := OKLCH(0.5, 0.2, 180)
	if c.IsEmpty() {
		t.Error("OKLCH() should not return empty color")
	}
}

func TestHexNoPrefix(t *testing.T) {
	t.Parallel()
	c := Hex("#ff5555")
	if c.HexNoPrefix() != "ff5555" {
		t.Errorf("HexNoPrefix() = %q, want %q", c.HexNoPrefix(), "ff5555")
	}
}

func TestHSLValues(t *testing.T) {
	t.Parallel()
	// Test red color
	c := Hex("#ff0000")
	h, s, l := c.HSLValues()
	if h < -1 || h > 1 { // Hue should be near 0 for red
		t.Errorf("HSLValues() hue = %f, want near 0", h)
	}
	if s < 0.99 || s > 1.01 {
		t.Errorf("HSLValues() saturation = %f, want 1.0", s)
	}
	if l < 0.49 || l > 0.51 {
		t.Errorf("HSLValues() lightness = %f, want 0.5", l)
	}
}

func TestOKLCHValues(t *testing.T) {
	t.Parallel()
	c := Hex("#ff0000")
	l, ch, h := c.OKLCHValues()
	if l < 0 || l > 1 {
		t.Errorf("OKLCHValues() lightness = %f, want 0-1", l)
	}
	if ch < 0 {
		t.Errorf("OKLCHValues() chroma = %f, want >= 0", ch)
	}
	_ = h // Hue can be any value
}

func TestCSS(t *testing.T) {
	t.Parallel()
	c := Hex("#ff5555")
	if c.CSS() != "#ff5555" {
		t.Errorf("CSS() = %q, want %q", c.CSS(), "#ff5555")
	}
}

func TestCSSVar(t *testing.T) {
	t.Parallel()
	c := Hex("#ff5555")
	if c.CSSVar("background") != "var(--theme-background)" {
		t.Errorf("CSSVar() = %q, want %q", c.CSSVar("background"), "var(--theme-background)")
	}
}

func TestCSSRGB(t *testing.T) {
	t.Parallel()
	c := Hex("#ff5555")
	if c.CSSRGB() != "rgb(255, 85, 85)" {
		t.Errorf("CSSRGB() = %q, want %q", c.CSSRGB(), "rgb(255, 85, 85)")
	}
}

func TestCSSRGBA(t *testing.T) {
	t.Parallel()
	c := Hex("#ff555580")
	got := c.CSSRGBA()
	// Alpha 128/255 ≈ 0.502
	if got != "rgba(255, 85, 85, 0.502)" {
		t.Errorf("CSSRGBA() = %q, want %q", got, "rgba(255, 85, 85, 0.502)")
	}
}

func TestCSSHSL(t *testing.T) {
	t.Parallel()
	c := Hex("#ff0000")
	got := c.CSSHSL()
	// Red: hue=0, sat=100%, light=50%
	if got != "hsl(0.0, 100.0%, 50.0%)" {
		t.Errorf("CSSHSL() = %q, want %q", got, "hsl(0.0, 100.0%, 50.0%)")
	}
}

func TestSaturate(t *testing.T) {
	t.Parallel()
	c := Hex("#808080") // Gray (0% saturation)
	saturated := c.Saturate(0.5)
	_, s1, _ := c.HSLValues()
	_, s2, _ := saturated.HSLValues()
	if s2 <= s1 {
		t.Errorf("Saturate() should increase saturation: %f -> %f", s1, s2)
	}
}

func TestDesaturate(t *testing.T) {
	t.Parallel()
	c := Hex("#ff0000") // Full red (100% saturation)
	desaturated := c.Desaturate(0.5)
	_, s1, _ := c.HSLValues()
	_, s2, _ := desaturated.HSLValues()
	if s2 >= s1 {
		t.Errorf("Desaturate() should decrease saturation: %f -> %f", s1, s2)
	}
}

func TestMix(t *testing.T) {
	t.Parallel()
	black := Hex("#000000")
	white := Hex("#ffffff")

	// Mix 50/50 should give gray
	mixed := black.Mix(white, 0.5)
	r, g, b := mixed.RGB()
	if r < 100 || r > 150 || g < 100 || g > 150 || b < 100 || b > 150 {
		t.Errorf("Mix(black, white, 0.5).RGB() = (%d, %d, %d), want near (128, 128, 128)", r, g, b)
	}

	// Mix 0.0 should give first color
	same := black.Mix(white, 0.0)
	if same.Hex() != "#000000" {
		t.Errorf("Mix(_, 0.0) = %q, want %q", same.Hex(), "#000000")
	}
}

func TestComplement(t *testing.T) {
	t.Parallel()
	red := Hex("#ff0000")
	comp := red.Complement()
	h, _, _ := comp.HSLValues()
	// Complement of red (hue=0) should be cyan (hue=180)
	if h < 170 || h > 190 {
		t.Errorf("Complement() hue = %f, want near 180", h)
	}
}

func TestInvert(t *testing.T) {
	t.Parallel()
	black := Hex("#000000")
	white := black.Invert()
	if white.Hex() != "#ffffff" {
		t.Errorf("Invert(#000000) = %q, want %q", white.Hex(), "#ffffff")
	}

	c := Hex("#ff0000")
	inv := c.Invert()
	if inv.Hex() != "#00ffff" {
		t.Errorf("Invert(#ff0000) = %q, want %q", inv.Hex(), "#00ffff")
	}
}

func TestIsLight(t *testing.T) {
	t.Parallel()
	white := Hex("#ffffff")
	black := Hex("#000000")

	if !white.IsLight() {
		t.Error("White should be light")
	}
	if black.IsLight() {
		t.Error("Black should not be light")
	}
}

func TestColorString(t *testing.T) {
	t.Parallel()
	c := Hex("#ff5555")
	if c.String() != "#ff5555" {
		t.Errorf("String() = %q, want %q", c.String(), "#ff5555")
	}
}

func TestColorGoString(t *testing.T) {
	t.Parallel()
	c := Hex("#ff5555")
	if c.GoString() != `gothememe.Hex("#ff5555")` {
		t.Errorf("GoString() = %q, want %q", c.GoString(), `gothememe.Hex("#ff5555")`)
	}
}

func TestStdColor(t *testing.T) {
	t.Parallel()
	c := Hex("#ff5555")
	stdC := c.StdColor()
	r, g, b, a := stdC.RGBA()
	// Standard library returns 16-bit values
	if r>>8 != 255 || g>>8 != 85 || b>>8 != 85 || a>>8 != 255 {
		t.Errorf("StdColor().RGBA() = (%d, %d, %d, %d), want (255<<8, 85<<8, 85<<8, 255<<8)", r, g, b, a)
	}
}

func TestFromStdColor(t *testing.T) {
	t.Parallel()
	c := Hex("#ff5555")
	stdC := c.StdColor()
	back := FromStdColor(stdC)
	// FromStdColor returns RGBA, so includes alpha channel
	if back.Hex() != "#ff5555ff" {
		t.Errorf("FromStdColor(c.StdColor()).Hex() = %q, want %q", back.Hex(), "#ff5555ff")
	}
}

func TestColorfulConversion(t *testing.T) {
	t.Parallel()
	c := Hex("#ff5555")
	// Verify colorful conversion doesn't panic
	h, s, l := c.HSLValues()
	if h < 0 || s < 0 || l < 0 {
		t.Error("HSLValues should return non-negative values")
	}

	// Test empty color
	empty := Color{}
	eh, es, el := empty.HSLValues()
	if eh != 0 && es != 0 && el != 0 {
		t.Log("Empty color HSL values:", eh, es, el)
	}
}

func TestWithAlphaEdgeCases(t *testing.T) {
	t.Parallel()
	c := Hex("#ff5555")

	// Test alpha < 0 (should clamp to 0)
	a1 := c.WithAlpha(-0.5)
	_, _, _, alpha1 := a1.RGBAComponents()
	if alpha1 != 0 {
		t.Errorf("WithAlpha(-0.5) alpha = %d, want 0", alpha1)
	}

	// Test alpha > 1 (should clamp to 1)
	a2 := c.WithAlpha(1.5)
	_, _, _, alpha2 := a2.RGBAComponents()
	if alpha2 != 255 {
		t.Errorf("WithAlpha(1.5) alpha = %d, want 255", alpha2)
	}
}

func TestEmptyColorMethods(t *testing.T) {
	t.Parallel()
	c := Color{}

	if !c.IsEmpty() {
		t.Error("Empty color should be empty")
	}
	if c.Hex() != "" {
		t.Errorf("Empty color Hex() = %q, want empty", c.Hex())
	}

	r, g, b := c.RGB()
	if r != 0 || g != 0 || b != 0 {
		t.Errorf("Empty color RGB() = (%d, %d, %d), want (0, 0, 0)", r, g, b)
	}
}

// Fuzz tests

func FuzzHexParsing(f *testing.F) {
	// Add seed corpus
	f.Add("#ff5555")
	f.Add("#FF5555")
	f.Add("ff5555")
	f.Add("#f55")
	f.Add("f55")
	f.Add("#ff555580")
	f.Add("")
	f.Add("#")
	f.Add("gggggg")
	f.Add("#gggggg")
	f.Add("rgb(255,0,0)")
	f.Add("123456")
	f.Add("#123")
	f.Add("#12345678")

	f.Fuzz(func(t *testing.T, input string) {
		// Hex() should never panic
		c := Hex(input)

		// If we got a valid color, verify properties
		if !c.IsEmpty() {
			// Hex output should start with #
			hex := c.Hex()
			if hex != "" && hex[0] != '#' {
				t.Errorf("valid color Hex() = %q, should start with #", hex)
			}

			// RGB method should return valid values (uint8 is always 0-255)
			r, g, b := c.RGB()
			_ = r
			_ = g
			_ = b

			// Color manipulation methods should not panic
			_ = c.Lighten(0.1)
			_ = c.Darken(0.1)
			_ = c.WithAlpha(0.5)
			_ = c.Complement()
			_ = c.Invert()
			_ = c.CSSRGB()
			_ = c.CSSHSL()
		}
	})
}

func FuzzRGBConstruction(f *testing.F) {
	// Add seed corpus
	f.Add(uint8(255), uint8(0), uint8(0))
	f.Add(uint8(0), uint8(255), uint8(0))
	f.Add(uint8(0), uint8(0), uint8(255))
	f.Add(uint8(0), uint8(0), uint8(0))
	f.Add(uint8(255), uint8(255), uint8(255))
	f.Add(uint8(128), uint8(128), uint8(128))

	f.Fuzz(func(t *testing.T, r, g, b uint8) {
		// RGB() should never panic
		c := RGB(r, g, b)

		// Should always produce a valid (non-empty) color
		if c.IsEmpty() {
			t.Errorf("RGB(%d, %d, %d) produced empty color", r, g, b)
		}

		// Round-trip should preserve values
		gotR, gotG, gotB := c.RGB()
		if gotR != r || gotG != g || gotB != b {
			t.Errorf("RGB(%d, %d, %d).RGB() = (%d, %d, %d)", r, g, b, gotR, gotG, gotB)
		}
	})
}

func FuzzColorManipulation(f *testing.F) {
	// Add seed corpus with amount values
	f.Add("#ff0000", 0.0)
	f.Add("#ff0000", 0.5)
	f.Add("#ff0000", 1.0)
	f.Add("#000000", 0.25)
	f.Add("#ffffff", 0.75)
	f.Add("#808080", 0.1)

	f.Fuzz(func(t *testing.T, hex string, amount float64) {
		c := Hex(hex)
		if c.IsEmpty() {
			return // Skip invalid hex inputs
		}

		// Manipulation methods should not panic
		_ = c.Lighten(amount)
		_ = c.Darken(amount)
		_ = c.Saturate(amount)
		_ = c.Desaturate(amount)
		_ = c.WithAlpha(amount)

		// Mix should not panic
		other := Hex("#808080")
		_ = c.Mix(other, amount)
	})
}
