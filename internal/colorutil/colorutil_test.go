package colorutil

import (
	"math"
	"testing"
)

func TestRelativeLuminance(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		r, g, b uint8
		want    float64
		tol     float64
	}{
		{"white", 255, 255, 255, 1.0, 0.001},
		{"black", 0, 0, 0, 0.0, 0.001},
		{"red", 255, 0, 0, 0.2126, 0.001},
		{"green", 0, 255, 0, 0.7152, 0.001},
		{"blue", 0, 0, 255, 0.0722, 0.001},
		{"gray", 128, 128, 128, 0.2158, 0.001},
		{"dracula background", 40, 42, 54, 0.0237, 0.001},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := RelativeLuminance(tt.r, tt.g, tt.b)
			if math.Abs(got-tt.want) > tt.tol {
				t.Errorf("RelativeLuminance(%d, %d, %d) = %f, want %f", tt.r, tt.g, tt.b, got, tt.want)
			}
		})
	}
}

func TestLinearize(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		v    float64
		want float64
		tol  float64
	}{
		{"below threshold", 0.02, 0.00155, 0.001},
		{"at threshold", 0.03928, 0.00304, 0.001},
		{"above threshold", 0.5, 0.214, 0.001},
		{"at 1.0", 1.0, 1.0, 0.001},
		{"at 0.0", 0.0, 0.0, 0.001},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := linearize(tt.v)
			if math.Abs(got-tt.want) > tt.tol {
				t.Errorf("linearize(%f) = %f, want %f", tt.v, got, tt.want)
			}
		})
	}
}

func TestHexToRGB(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		hex     string
		r, g, b uint8
	}{
		{"6-char with hash", "#ff5555", 255, 85, 85},
		{"6-char without hash", "ff5555", 255, 85, 85},
		{"3-char with hash", "#f55", 255, 85, 85},
		{"3-char without hash", "f55", 255, 85, 85},
		{"black", "#000000", 0, 0, 0},
		{"white", "#ffffff", 255, 255, 255},
		{"uppercase", "#ABCDEF", 171, 205, 239},
		{"mixed case", "#AbCdEf", 171, 205, 239},
		{"empty string", "", 0, 0, 0},
		{"too short", "#ff", 0, 0, 0},
		{"invalid chars", "#gggggg", 0, 0, 0},
		{"8-char with alpha", "#ff5555aa", 255, 85, 85},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r, g, b := HexToRGB(tt.hex)
			if r != tt.r || g != tt.g || b != tt.b {
				t.Errorf("HexToRGB(%q) = (%d, %d, %d), want (%d, %d, %d)", tt.hex, r, g, b, tt.r, tt.g, tt.b)
			}
		})
	}
}

func TestLuminanceHex(t *testing.T) {
	t.Parallel()
	tests := []struct {
		hex  string
		want float64
		tol  float64
	}{
		{"#ffffff", 1.0, 0.001},
		{"#000000", 0.0, 0.001},
		{"#282a36", 0.0237, 0.001},
		{"#f8f8f2", 0.935, 0.01},
		{"282a36", 0.0237, 0.001},
	}

	for _, tt := range tests {
		t.Run(tt.hex, func(t *testing.T) {
			t.Parallel()
			got := LuminanceHex(tt.hex)
			if math.Abs(got-tt.want) > tt.tol {
				t.Errorf("LuminanceHex(%q) = %f, want %f", tt.hex, got, tt.want)
			}
		})
	}
}

func TestContrastRatio(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                   string
		r1, g1, b1, r2, g2, b2 uint8
		want                   float64
		tol                    float64
	}{
		{"black on white", 0, 0, 0, 255, 255, 255, 21.0, 0.1},
		{"white on black", 255, 255, 255, 0, 0, 0, 21.0, 0.1},
		{"same color", 128, 128, 128, 128, 128, 128, 1.0, 0.001},
		{"dracula fg/bg", 248, 248, 242, 40, 42, 54, 13.4, 0.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := ContrastRatio(tt.r1, tt.g1, tt.b1, tt.r2, tt.g2, tt.b2)
			if math.Abs(got-tt.want) > tt.tol {
				t.Errorf("ContrastRatio() = %f, want %f", got, tt.want)
			}
		})
	}
}

func TestContrastRatioHex(t *testing.T) {
	t.Parallel()
	tests := []struct {
		hex1, hex2 string
		want       float64
		tol        float64
	}{
		{"#000000", "#ffffff", 21.0, 0.1},
		{"#282a36", "#f8f8f2", 13.4, 0.5},
		{"#ffffff", "#ffffff", 1.0, 0.001},
	}

	for _, tt := range tests {
		t.Run(tt.hex1+"_"+tt.hex2, func(t *testing.T) {
			t.Parallel()
			got := ContrastRatioHex(tt.hex1, tt.hex2)
			if math.Abs(got-tt.want) > tt.tol {
				t.Errorf("ContrastRatioHex(%q, %q) = %f, want %f", tt.hex1, tt.hex2, got, tt.want)
			}
		})
	}
}

func TestRatioFromLuminance(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		l1, l2 float64
		want   float64
		tol    float64
	}{
		{"same luminance", 0.5, 0.5, 1.0, 0.001},
		{"black and white", 0.0, 1.0, 21.0, 0.1},
		{"reversed order", 1.0, 0.0, 21.0, 0.1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := RatioFromLuminance(tt.l1, tt.l2)
			if math.Abs(got-tt.want) > tt.tol {
				t.Errorf("RatioFromLuminance(%f, %f) = %f, want %f", tt.l1, tt.l2, got, tt.want)
			}
		})
	}
}

func TestParseHexByte(t *testing.T) {
	t.Parallel()
	tests := []struct {
		hex  string
		want uint8
	}{
		{"ff", 255},
		{"00", 0},
		{"80", 128},
		{"FF", 255},
		{"Ab", 171},
		{"f", 0},     // too short
		{"", 0},      // empty
		{"gg", 0},    // invalid
		{"zz", 0},    // invalid
		{"abc", 171}, // extra chars ignored
	}

	for _, tt := range tests {
		t.Run(tt.hex, func(t *testing.T) {
			t.Parallel()
			got := ParseHexByte(tt.hex)
			if got != tt.want {
				t.Errorf("ParseHexByte(%q) = %d, want %d", tt.hex, got, tt.want)
			}
		})
	}
}

func BenchmarkRelativeLuminance(b *testing.B) {
	for b.Loop() {
		RelativeLuminance(40, 42, 54)
	}
}

func BenchmarkHexToRGB(b *testing.B) {
	for b.Loop() {
		HexToRGB("#282a36")
	}
}

func BenchmarkContrastRatioHex(b *testing.B) {
	for b.Loop() {
		ContrastRatioHex("#282a36", "#f8f8f2")
	}
}
