package pairs

import (
	"testing"
)

func TestStandardPairSpecs(t *testing.T) {
	t.Parallel()

	specs := StandardPairSpecs()

	// Should have exactly 14 pairs
	if len(specs) != 14 {
		t.Errorf("StandardPairSpecs() returned %d pairs, want 14", len(specs))
	}

	// Verify expected pairs exist
	expectedPairs := []StandardPairSpec{
		{FgName: "TextPrimary", BgName: "Background"},
		{FgName: "TextSecondary", BgName: "Background"},
		{FgName: "TextMuted", BgName: "Background"},
		{FgName: "TextPrimary", BgName: "BackgroundSecondary"},
		{FgName: "TextPrimary", BgName: "Surface"},
		{FgName: "Accent", BgName: "Background"},
		{FgName: "Success.Text", BgName: "Success.Background"},
		{FgName: "Warning.Text", BgName: "Warning.Background"},
		{FgName: "Error.Text", BgName: "Error.Background"},
		{FgName: "Info.Text", BgName: "Info.Background"},
		{FgName: "CodeText", BgName: "CodeBackground"},
		{FgName: "CodeComment", BgName: "CodeBackground"},
		{FgName: "CodeKeyword", BgName: "CodeBackground"},
		{FgName: "CodeString", BgName: "CodeBackground"},
	}

	for i, expected := range expectedPairs {
		if i >= len(specs) {
			t.Errorf("Missing pair at index %d: %+v", i, expected)
			continue
		}
		if specs[i] != expected {
			t.Errorf("Pair %d = %+v, want %+v", i, specs[i], expected)
		}
	}
}

func TestStandardPairSpecsNonEmpty(t *testing.T) {
	t.Parallel()

	specs := StandardPairSpecs()

	for i, spec := range specs {
		if spec.FgName == "" {
			t.Errorf("Pair %d has empty FgName", i)
		}
		if spec.BgName == "" {
			t.Errorf("Pair %d has empty BgName", i)
		}
	}
}

func TestColorPairStruct(t *testing.T) {
	t.Parallel()

	pair := ColorPair{
		FgName: "TestFg",
		BgName: "TestBg",
		FgHex:  "#ffffff",
		BgHex:  "#000000",
	}

	if pair.FgName != "TestFg" {
		t.Errorf("FgName = %q, want %q", pair.FgName, "TestFg")
	}
	if pair.BgName != "TestBg" {
		t.Errorf("BgName = %q, want %q", pair.BgName, "TestBg")
	}
	if pair.FgHex != "#ffffff" {
		t.Errorf("FgHex = %q, want %q", pair.FgHex, "#ffffff")
	}
	if pair.BgHex != "#000000" {
		t.Errorf("BgHex = %q, want %q", pair.BgHex, "#000000")
	}
}
