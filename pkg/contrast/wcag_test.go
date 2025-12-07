package contrast

import (
	"math"
	"testing"

	"github.com/tj-smith47/gothememe/internal/colorutil"
)

func TestLuminance(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Luminance(tt.r, tt.g, tt.b)
			if math.Abs(got-tt.want) > tt.tol {
				t.Errorf("Luminance(%d, %d, %d) = %f, want %f", tt.r, tt.g, tt.b, got, tt.want)
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
		{"ffffff", 1.0, 0.001},
		{"#000000", 0.0, 0.001},
		{"#f8f8f2", 0.935, 0.01}, // Dracula foreground
		{"#282a36", 0.024, 0.01}, // Dracula background
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

func TestRatio(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		r1, g1, b1 uint8
		r2, g2, b2 uint8
		want       float64
		tol        float64
	}{
		{"black/white", 0, 0, 0, 255, 255, 255, 21.0, 0.1},
		{"same color", 128, 128, 128, 128, 128, 128, 1.0, 0.001},
		{"dracula bg/fg", 40, 42, 54, 248, 248, 242, 13.27, 0.1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Ratio(tt.r1, tt.g1, tt.b1, tt.r2, tt.g2, tt.b2)
			if math.Abs(got-tt.want) > tt.tol {
				t.Errorf("Ratio() = %f, want %f", got, tt.want)
			}
		})
	}
}

func TestRatioHex(t *testing.T) {
	t.Parallel()
	tests := []struct {
		hex1, hex2 string
		want       float64
		tol        float64
	}{
		{"#000000", "#ffffff", 21.0, 0.1},
		{"#282a36", "#f8f8f2", 13.27, 0.5}, // Dracula
		{"#2e3440", "#eceff4", 10.84, 0.5}, // Nord
	}

	for _, tt := range tests {
		t.Run(tt.hex1+"/"+tt.hex2, func(t *testing.T) {
			t.Parallel()
			got := RatioHex(tt.hex1, tt.hex2)
			if math.Abs(got-tt.want) > tt.tol {
				t.Errorf("RatioHex(%q, %q) = %f, want %f", tt.hex1, tt.hex2, got, tt.want)
			}
		})
	}
}

func TestMeetsAA(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		hex1, hex2 string
		largeText  bool
		want       bool
	}{
		{"dracula normal", "#282a36", "#f8f8f2", false, true},
		{"dracula large", "#282a36", "#f8f8f2", true, true},
		{"low contrast normal", "#666666", "#888888", false, false},
		{"low contrast large", "#666666", "#888888", true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := MeetsAAHex(tt.hex1, tt.hex2, tt.largeText)
			if got != tt.want {
				t.Errorf("MeetsAAHex(%q, %q, %v) = %v, want %v", tt.hex1, tt.hex2, tt.largeText, got, tt.want)
			}
		})
	}
}

func TestMeetsAA_RGB(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		r1, g1, b1 uint8
		r2, g2, b2 uint8
		largeText  bool
		want       bool
	}{
		{"black/white normal", 0, 0, 0, 255, 255, 255, false, true},
		{"black/white large", 0, 0, 0, 255, 255, 255, true, true},
		{"low contrast normal", 102, 102, 102, 136, 136, 136, false, false},
		{"medium contrast large", 51, 51, 51, 153, 153, 153, true, true}, // ratio ~4.2
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := MeetsAA(tt.r1, tt.g1, tt.b1, tt.r2, tt.g2, tt.b2, tt.largeText)
			if got != tt.want {
				t.Errorf("MeetsAA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMeetsAAA(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		hex1, hex2 string
		largeText  bool
		want       bool
	}{
		{"black/white normal", "#000000", "#ffffff", false, true},
		{"dracula normal", "#282a36", "#f8f8f2", false, true},
		{"medium contrast normal", "#333333", "#cccccc", false, true}, // ratio ~8.5
		{"medium contrast large", "#333333", "#cccccc", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := MeetsAAAHex(tt.hex1, tt.hex2, tt.largeText)
			if got != tt.want {
				t.Errorf("MeetsAAAHex(%q, %q, %v) = %v, want %v", tt.hex1, tt.hex2, tt.largeText, got, tt.want)
			}
		})
	}
}

func TestMeetsAAA_RGB(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		r1, g1, b1 uint8
		r2, g2, b2 uint8
		largeText  bool
		want       bool
	}{
		{"black/white normal", 0, 0, 0, 255, 255, 255, false, true},
		{"black/white large", 0, 0, 0, 255, 255, 255, true, true},
		{"medium contrast normal fail", 85, 85, 85, 170, 170, 170, false, false}, // ratio ~3.5
		{"medium contrast large pass", 51, 51, 51, 170, 170, 170, true, true},    // ratio ~5.2
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := MeetsAAA(tt.r1, tt.g1, tt.b1, tt.r2, tt.g2, tt.b2, tt.largeText)
			if got != tt.want {
				t.Errorf("MeetsAAA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMeetsUIComponent(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		r1, g1, b1 uint8
		r2, g2, b2 uint8
		want       bool
	}{
		{"black/white", 0, 0, 0, 255, 255, 255, true},
		{"high contrast", 50, 50, 50, 200, 200, 200, true},
		{"low contrast", 128, 128, 128, 160, 160, 160, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := MeetsUIComponent(tt.r1, tt.g1, tt.b1, tt.r2, tt.g2, tt.b2)
			if got != tt.want {
				t.Errorf("MeetsUIComponent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMeetsUIComponentHex(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		hex1, hex2 string
		want       bool
	}{
		{"black/white", "#000000", "#ffffff", true},
		{"high contrast", "#333333", "#cccccc", true},
		{"low contrast", "#808080", "#a0a0a0", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := MeetsUIComponentHex(tt.hex1, tt.hex2)
			if got != tt.want {
				t.Errorf("MeetsUIComponentHex(%q, %q) = %v, want %v", tt.hex1, tt.hex2, got, tt.want)
			}
		})
	}
}

func TestCheck_RGB(t *testing.T) {
	t.Parallel()
	// Note: LevelAA (4.5) and LevelAAALarge (4.5) have same threshold.
	// Since AAALarge is checked first, ratios 4.5-6.99 return AAALarge.
	tests := []struct {
		name       string
		r1, g1, b1 uint8
		r2, g2, b2 uint8
		want       Level
	}{
		{"black/white AAA", 0, 0, 0, 255, 255, 255, LevelAAA},        // 21:1
		{"dracula AAA", 40, 42, 54, 248, 248, 242, LevelAAA},         // ~13:1
		{"AAA large only", 60, 60, 60, 200, 200, 200, LevelAAALarge}, // ~5.8 (4.5-6.99)
		{"AA large only", 90, 90, 90, 180, 180, 180, LevelAALarge},   // ~3.2 (3.0-4.49)
		{"fail", 140, 140, 140, 160, 160, 160, LevelFail},            // ~1.3 (<3.0)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Check(tt.r1, tt.g1, tt.b1, tt.r2, tt.g2, tt.b2)
			if got != tt.want {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIssueStruct(t *testing.T) {
	t.Parallel()
	issue := Issue{
		ForegroundName: "TextPrimary",
		BackgroundName: "Background",
		Foreground:     "#333333",
		Background:     "#444444",
		Ratio:          1.5,
		Level:          LevelFail,
		Required:       LevelAA,
	}

	if issue.ForegroundName != "TextPrimary" {
		t.Errorf("Issue.ForegroundName = %q, want %q", issue.ForegroundName, "TextPrimary")
	}
	if issue.Ratio != 1.5 {
		t.Errorf("Issue.Ratio = %f, want %f", issue.Ratio, 1.5)
	}
	if issue.Level != LevelFail {
		t.Errorf("Issue.Level = %v, want %v", issue.Level, LevelFail)
	}
}

func TestCheck(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		hex1, hex2 string
		want       Level
	}{
		{"black/white", "#000000", "#ffffff", LevelAAA},
		{"dracula", "#282a36", "#f8f8f2", LevelAAA},
		{"medium", "#333333", "#bbbbbb", LevelAAALarge}, // ratio ~6.1
		{"low", "#555555", "#888888", LevelFail},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := CheckHex(tt.hex1, tt.hex2)
			if got != tt.want {
				t.Errorf("CheckHex(%q, %q) = %v, want %v", tt.hex1, tt.hex2, got, tt.want)
			}
		})
	}
}

func TestLevelString(t *testing.T) {
	t.Parallel()
	tests := []struct {
		level Level
		want  string
	}{
		{LevelAAA, "AAA"},
		{LevelAAALarge, "AAA (large text only)"},
		{LevelAA, "AA"},
		{LevelAALarge, "AA (large text only)"},
		{LevelFail, "Fail"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			t.Parallel()
			got := tt.level.String()
			if got != tt.want {
				t.Errorf("Level.String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestHexToRGB(t *testing.T) {
	t.Parallel()
	tests := []struct {
		hex     string
		r, g, b uint8
	}{
		{"#ff5555", 255, 85, 85},
		{"ff5555", 255, 85, 85},
		{"#f55", 255, 85, 85},
		{"f55", 255, 85, 85},
		{"#000000", 0, 0, 0},
		{"#ffffff", 255, 255, 255},
	}

	for _, tt := range tests {
		t.Run(tt.hex, func(t *testing.T) {
			t.Parallel()
			r, g, b := colorutil.HexToRGB(tt.hex)
			if r != tt.r || g != tt.g || b != tt.b {
				t.Errorf("HexToRGB(%q) = (%d, %d, %d), want (%d, %d, %d)", tt.hex, r, g, b, tt.r, tt.g, tt.b)
			}
		})
	}
}

func BenchmarkRatioHex(b *testing.B) {
	for b.Loop() {
		RatioHex("#282a36", "#f8f8f2")
	}
}

func BenchmarkLuminanceHex(b *testing.B) {
	for b.Loop() {
		LuminanceHex("#282a36")
	}
}
