package generator

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// testWriteJSON is a helper that encodes JSON for test handlers.
// Error handling is relaxed since test failures will surface issues.
func testWriteJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v) //nolint:errcheck // test helper
}

// testWriteBytes writes bytes for test handlers.
func testWriteBytes(w http.ResponseWriter, b []byte) {
	_, _ = w.Write(b) //nolint:errcheck // test helper
}

// testReadFile reads a file in tests (safe path from t.TempDir).
func testReadFile(t *testing.T, path string) []byte {
	t.Helper()
	content, err := os.ReadFile(path) //nolint:gosec // G304: path is from t.TempDir
	if err != nil {
		t.Fatalf("reading file %s: %v", path, err)
	}
	return content
}

func TestToThemeID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want string
	}{
		{"Tokyo Night", "tokyo_night"},
		{"Dracula", "dracula"},
		{"Gruvbox Dark", "gruvbox_dark"},
		{"One Dark", "one_dark"},
		{"Catppuccin Mocha", "catppuccin_mocha"},
		{"3024 Day", "3024_day"},
		{"  Spaced  Out  ", "spaced_out"},
		{"Special!@#Chars", "special_chars"},
		{"UPPERCASE", "uppercase"},
		{"lowercase", "lowercase"},
		{"MixedCase Theme", "mixedcase_theme"},
		{"Multiple   Spaces", "multiple_spaces"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := toThemeID(tt.name)
			if got != tt.want {
				t.Errorf("toThemeID(%q) = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}

func TestToVarName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want string
	}{
		{"Tokyo Night", "TokyoNight"},
		{"Dracula", "Dracula"},
		{"Gruvbox Dark", "GruvboxDark"},
		{"One Dark", "OneDark"},
		{"3024 Day", "N3024Day"},
		{"  Spaced  Out  ", "SpacedOut"},
		{"Special!@#Chars", "SpecialChars"},
		{"lowercase", "Lowercase"},
		{"ALL CAPS", "AllCaps"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := toVarName(tt.name)
			if got != tt.want {
				t.Errorf("toVarName(%q) = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}

func TestIsDarkTheme(t *testing.T) {
	t.Parallel()

	tests := []struct {
		background string
		wantDark   bool
	}{
		{"#000000", true},
		{"#282a36", true},  // Dracula
		{"#1e1e1e", true},  // VS Code Dark
		{"#2e3440", true},  // Nord
		{"#ffffff", false}, // White
		{"#f8f8f2", false}, // Light
		{"#fdf6e3", false}, // Solarized Light
		{"#808080", true},  // Gray - luminance ~0.2, still considered dark by threshold
		{"invalid", true},  // Invalid defaults to dark
		{"#ff", true},      // Too short defaults to dark
		{"", true},         // Empty defaults to dark
	}

	for _, tt := range tests {
		t.Run(tt.background, func(t *testing.T) {
			t.Parallel()
			got := isDarkTheme(tt.background)
			if got != tt.wantDark {
				t.Errorf("isDarkTheme(%q) = %v, want %v", tt.background, got, tt.wantDark)
			}
		})
	}
}

func TestNewGenerator(t *testing.T) {
	t.Parallel()

	gen := NewGenerator("/tmp/test-output")

	if gen == nil {
		t.Fatal("NewGenerator returned nil")
	}

	if gen.outputDir != "/tmp/test-output" {
		t.Errorf("outputDir = %q, want /tmp/test-output", gen.outputDir)
	}
}

func TestGeneratorSetThemes(t *testing.T) {
	t.Parallel()

	gen := NewGenerator("/tmp/test")

	themes := []*WindowsTerminalTheme{
		{Name: "Theme1", Background: "#000000"},
		{Name: "Theme2", Background: "#ffffff"},
	}

	gen.SetThemes(themes)

	if len(gen.themes) != 2 {
		t.Errorf("themes length = %d, want 2", len(gen.themes))
	}
}

// ============================================================================
// Fetcher Tests
// ============================================================================

func TestNewFetcher(t *testing.T) {
	t.Parallel()

	f := NewFetcher()

	if f == nil {
		t.Fatal("NewFetcher returned nil")
	}
	if f.client == nil {
		t.Error("client is nil")
	}
	if f.client.Timeout != 30*1e9 { // 30 seconds in nanoseconds
		t.Errorf("client timeout = %v, want 30s", f.client.Timeout)
	}
}

func TestFetcher_WithBaseURL(t *testing.T) {
	t.Parallel()

	f := NewFetcher().WithBaseURL("http://test.example.com")

	if f.baseURL != "http://test.example.com" {
		t.Errorf("baseURL = %q, want http://test.example.com", f.baseURL)
	}
}

func TestFetcher_WithClient(t *testing.T) {
	t.Parallel()

	customClient := &http.Client{Timeout: 60 * 1e9}
	f := NewFetcher().WithClient(customClient)

	if f.client != customClient {
		t.Error("client was not set correctly")
	}
}

func TestFetcher_apiURL(t *testing.T) {
	t.Parallel()

	t.Run("default", func(t *testing.T) {
		t.Parallel()
		f := NewFetcher()
		if f.apiURL() != APIURL {
			t.Errorf("apiURL() = %q, want %q", f.apiURL(), APIURL)
		}
	})

	t.Run("custom", func(t *testing.T) {
		t.Parallel()
		f := NewFetcher().WithBaseURL("http://custom.example.com")
		if f.apiURL() != "http://custom.example.com" {
			t.Errorf("apiURL() = %q, want http://custom.example.com", f.apiURL())
		}
	})
}

func TestListThemes_Success(t *testing.T) {
	t.Parallel()

	contents := []GitHubContent{
		{Name: "Dracula.json", Path: "windowsterminal/Dracula.json", DownloadURL: "http://example.com/Dracula.json", Type: "file"},
		{Name: "Nord.json", Path: "windowsterminal/Nord.json", DownloadURL: "http://example.com/Nord.json", Type: "file"},
		{Name: "README.md", Path: "windowsterminal/README.md", DownloadURL: "http://example.com/README.md", Type: "file"},
		{Name: "subdir", Path: "windowsterminal/subdir", DownloadURL: "", Type: "dir"},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		testWriteJSON(w, contents)
	}))
	defer server.Close()

	f := NewFetcher().WithBaseURL(server.URL)
	themes, err := f.ListThemes()

	if err != nil {
		t.Fatalf("ListThemes() error = %v", err)
	}

	// Should only include JSON files
	if len(themes) != 2 {
		t.Errorf("len(themes) = %d, want 2", len(themes))
	}

	for _, theme := range themes {
		if !strings.HasSuffix(theme.Name, ".json") {
			t.Errorf("non-JSON file included: %s", theme.Name)
		}
	}
}

func TestListThemes_EmptyResponse(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		testWriteBytes(w, []byte("[]"))
	}))
	defer server.Close()

	f := NewFetcher().WithBaseURL(server.URL)
	themes, err := f.ListThemes()

	if err != nil {
		t.Fatalf("ListThemes() error = %v", err)
	}

	if len(themes) != 0 {
		t.Errorf("len(themes) = %d, want 0", len(themes))
	}
}

func TestListThemes_NoJsonFiles(t *testing.T) {
	t.Parallel()

	contents := []GitHubContent{
		{Name: "README.md", Path: "windowsterminal/README.md", Type: "file"},
		{Name: "LICENSE", Path: "windowsterminal/LICENSE", Type: "file"},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		testWriteJSON(w, contents)
	}))
	defer server.Close()

	f := NewFetcher().WithBaseURL(server.URL)
	themes, err := f.ListThemes()

	if err != nil {
		t.Fatalf("ListThemes() error = %v", err)
	}

	if len(themes) != 0 {
		t.Errorf("len(themes) = %d, want 0", len(themes))
	}
}

func TestListThemes_HttpError(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		statusCode int
	}{
		{"NotFound", http.StatusNotFound},
		{"InternalServerError", http.StatusInternalServerError},
		{"Forbidden", http.StatusForbidden},
		{"RateLimited", http.StatusTooManyRequests},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(tc.statusCode)
				testWriteBytes(w, []byte("error response body"))
			}))
			defer server.Close()

			f := NewFetcher().WithBaseURL(server.URL)
			_, err := f.ListThemes()

			if err == nil {
				t.Fatal("ListThemes() expected error, got nil")
			}

			if !strings.Contains(err.Error(), "GitHub API") {
				t.Errorf("error should mention GitHub API: %v", err)
			}
		})
	}
}

func TestListThemes_InvalidJSON(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		testWriteBytes(w, []byte("not valid json"))
	}))
	defer server.Close()

	f := NewFetcher().WithBaseURL(server.URL)
	_, err := f.ListThemes()

	if err == nil {
		t.Fatal("ListThemes() expected error for invalid JSON")
	}

	if !strings.Contains(err.Error(), "decoding") {
		t.Errorf("error should mention decoding: %v", err)
	}
}

func TestFetchTheme_Success(t *testing.T) {
	t.Parallel()

	theme := WindowsTerminalTheme{
		Name:        "Dracula",
		Background:  "#282a36",
		Foreground:  "#f8f8f2",
		CursorColor: "#f8f8f2",
		Black:       "#21222c",
		Red:         "#ff5555",
		Green:       "#50fa7b",
		Yellow:      "#f1fa8c",
		Blue:        "#bd93f9",
		Purple:      "#ff79c6",
		Cyan:        "#8be9fd",
		White:       "#f8f8f2",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		testWriteJSON(w, theme)
	}))
	defer server.Close()

	f := NewFetcher()
	result, err := f.FetchTheme(server.URL)

	if err != nil {
		t.Fatalf("FetchTheme() error = %v", err)
	}

	if result.Name != "Dracula" {
		t.Errorf("Name = %q, want Dracula", result.Name)
	}
	if result.Background != "#282a36" {
		t.Errorf("Background = %q, want #282a36", result.Background)
	}
}

func TestFetchTheme_HttpError(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	f := NewFetcher()
	_, err := f.FetchTheme(server.URL)

	if err == nil {
		t.Fatal("FetchTheme() expected error for HTTP 404")
	}

	if !strings.Contains(err.Error(), "HTTP 404") {
		t.Errorf("error should mention HTTP 404: %v", err)
	}
}

func TestFetchTheme_InvalidJSON(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		testWriteBytes(w, []byte("{invalid json"))
	}))
	defer server.Close()

	f := NewFetcher()
	_, err := f.FetchTheme(server.URL)

	if err == nil {
		t.Fatal("FetchTheme() expected error for invalid JSON")
	}

	if !strings.Contains(err.Error(), "decoding") {
		t.Errorf("error should mention decoding: %v", err)
	}
}

func TestFetchAllThemes_AllSucceed(t *testing.T) {
	t.Parallel()

	mux := http.NewServeMux()

	// List endpoint
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		contents := []GitHubContent{
			{Name: "Dracula.json", DownloadURL: "/theme/dracula", Type: "file"},
			{Name: "Nord.json", DownloadURL: "/theme/nord", Type: "file"},
		}
		testWriteJSON(w, contents)
	})

	// Theme endpoints - will be called with absolute URLs
	mux.HandleFunc("/theme/dracula", func(w http.ResponseWriter, _ *http.Request) {
		theme := WindowsTerminalTheme{Name: "Dracula", Background: "#282a36"}
		testWriteJSON(w, theme)
	})

	mux.HandleFunc("/theme/nord", func(w http.ResponseWriter, _ *http.Request) {
		theme := WindowsTerminalTheme{Name: "Nord", Background: "#2e3440"}
		testWriteJSON(w, theme)
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	// Update download URLs to use actual server URL
	listHandler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		contents := []GitHubContent{
			{Name: "Dracula.json", DownloadURL: server.URL + "/theme/dracula", Type: "file"},
			{Name: "Nord.json", DownloadURL: server.URL + "/theme/nord", Type: "file"},
		}
		testWriteJSON(w, contents)
	})

	listServer := httptest.NewServer(listHandler)
	defer listServer.Close()

	f := NewFetcher().WithBaseURL(listServer.URL).WithClient(server.Client())
	themes, err := f.FetchAllThemes()

	if err != nil {
		t.Fatalf("FetchAllThemes() error = %v", err)
	}

	if len(themes) != 2 {
		t.Errorf("len(themes) = %d, want 2", len(themes))
	}
}

func TestFetchAllThemes_PartialFailure(t *testing.T) {
	t.Parallel()

	mux := http.NewServeMux()

	mux.HandleFunc("/theme/dracula", func(w http.ResponseWriter, _ *http.Request) {
		theme := WindowsTerminalTheme{Name: "Dracula", Background: "#282a36"}
		testWriteJSON(w, theme)
	})

	mux.HandleFunc("/theme/broken", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	listHandler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		contents := []GitHubContent{
			{Name: "Dracula.json", DownloadURL: server.URL + "/theme/dracula", Type: "file"},
			{Name: "Broken.json", DownloadURL: server.URL + "/theme/broken", Type: "file"},
		}
		testWriteJSON(w, contents)
	})

	listServer := httptest.NewServer(listHandler)
	defer listServer.Close()

	f := NewFetcher().WithBaseURL(listServer.URL)
	themes, err := f.FetchAllThemes()

	// Should not return error, but skip failed themes
	if err != nil {
		t.Fatalf("FetchAllThemes() error = %v", err)
	}

	// Should have 1 successful theme
	if len(themes) != 1 {
		t.Errorf("len(themes) = %d, want 1", len(themes))
	}

	if themes[0].Name != "Dracula" {
		t.Errorf("first theme name = %q, want Dracula", themes[0].Name)
	}
}

func TestFetchAllThemes_ListError(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	f := NewFetcher().WithBaseURL(server.URL)
	_, err := f.FetchAllThemes()

	if err == nil {
		t.Fatal("FetchAllThemes() expected error when ListThemes fails")
	}
}

// ============================================================================
// Generator Tests
// ============================================================================

func TestGenerate_ValidThemes(t *testing.T) {
	t.Parallel()

	outputDir := t.TempDir()
	gen := NewGenerator(outputDir)

	themes := []*WindowsTerminalTheme{
		createTestTheme("Dracula", "#282a36"),
		createTestTheme("Nord", "#2e3440"),
		createTestTheme("Gruvbox", "#282828"),
	}
	gen.SetThemes(themes)

	err := gen.Generate()
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	// Check that files were created
	files := []string{"dracula.go", "nord.go", "gruvbox.go", "themes.go", "doc.go"}
	for _, f := range files {
		path := filepath.Join(outputDir, f)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("expected file %s to exist", f)
		}
	}

	// Verify themes.go content
	themesContent := testReadFile(t, filepath.Join(outputDir, "themes.go"))

	if !strings.Contains(string(themesContent), "func All()") {
		t.Error("themes.go should contain All() function")
	}
	if !strings.Contains(string(themesContent), "func ByID(") {
		t.Error("themes.go should contain ByID() function")
	}
}

func TestGenerate_DuplicateIds(t *testing.T) {
	t.Parallel()

	outputDir := t.TempDir()
	gen := NewGenerator(outputDir)

	// Two themes that normalize to the same ID
	themes := []*WindowsTerminalTheme{
		createTestTheme("Test Theme", "#000000"),
		createTestTheme("Test-Theme", "#111111"), // Same ID: test_theme
	}
	gen.SetThemes(themes)

	err := gen.Generate()
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	// Check both files exist with different names
	if _, err := os.Stat(filepath.Join(outputDir, "test_theme.go")); os.IsNotExist(err) {
		t.Error("expected test_theme.go to exist")
	}
	if _, err := os.Stat(filepath.Join(outputDir, "test_theme_2.go")); os.IsNotExist(err) {
		t.Error("expected test_theme_2.go to exist")
	}
}

func TestGenerate_DuplicateVarNames(t *testing.T) {
	t.Parallel()

	outputDir := t.TempDir()
	gen := NewGenerator(outputDir)

	// Two themes that normalize to similar var names
	themes := []*WindowsTerminalTheme{
		createTestTheme("Test", "#000000"),
		createTestTheme("TEST", "#111111"), // Same var name: Test
	}
	gen.SetThemes(themes)

	err := gen.Generate()
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	// Verify themes.go contains both variants
	content := testReadFile(t, filepath.Join(outputDir, "themes.go"))

	if !strings.Contains(string(content), "themeTestInstance") {
		t.Error("themes.go should contain themeTestInstance")
	}
	if !strings.Contains(string(content), "themeTest2Instance") {
		t.Error("themes.go should contain themeTest2Instance")
	}
}

func TestGenerate_EmptyThemesList(t *testing.T) {
	t.Parallel()

	outputDir := t.TempDir()
	gen := NewGenerator(outputDir)
	gen.SetThemes([]*WindowsTerminalTheme{})

	err := gen.Generate()
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	// themes.go and doc.go should still be created
	if _, err := os.Stat(filepath.Join(outputDir, "themes.go")); os.IsNotExist(err) {
		t.Error("expected themes.go to exist even with empty themes")
	}
	if _, err := os.Stat(filepath.Join(outputDir, "doc.go")); os.IsNotExist(err) {
		t.Error("expected doc.go to exist even with empty themes")
	}
}

func TestGenerate_DirectoryCreation(t *testing.T) {
	t.Parallel()

	baseDir := t.TempDir()
	outputDir := filepath.Join(baseDir, "nested", "output", "dir")

	gen := NewGenerator(outputDir)
	gen.SetThemes([]*WindowsTerminalTheme{
		createTestTheme("Test", "#000000"),
	})

	err := gen.Generate()
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		t.Error("expected output directory to be created")
	}
}

func TestGenerateMarkdown_Success(t *testing.T) {
	t.Parallel()

	outputDir := t.TempDir()
	gen := NewGenerator(outputDir)

	themes := []*WindowsTerminalTheme{
		createTestTheme("Dracula", "#282a36"),
		createTestTheme("Nord", "#2e3440"),
	}
	gen.SetThemes(themes)

	outputPath := filepath.Join(outputDir, "DEFAULT_THEMES.md")
	err := gen.GenerateMarkdown(outputPath)
	if err != nil {
		t.Fatalf("GenerateMarkdown() error = %v", err)
	}

	content := testReadFile(t, outputPath)

	if !strings.Contains(string(content), "# Default Themes") {
		t.Error("markdown should contain title")
	}
	if !strings.Contains(string(content), "Dracula") {
		t.Error("markdown should contain Dracula")
	}
	if !strings.Contains(string(content), "Nord") {
		t.Error("markdown should contain Nord")
	}
	if !strings.Contains(string(content), "**2**") {
		t.Error("markdown should contain theme count")
	}
}

func TestGenerateMarkdown_WithDuplicates(t *testing.T) {
	t.Parallel()

	outputDir := t.TempDir()
	gen := NewGenerator(outputDir)

	themes := []*WindowsTerminalTheme{
		createTestTheme("Test", "#000000"),
		createTestTheme("Test", "#111111"), // Duplicate name
	}
	gen.SetThemes(themes)

	outputPath := filepath.Join(outputDir, "DEFAULT_THEMES.md")
	err := gen.GenerateMarkdown(outputPath)
	if err != nil {
		t.Fatalf("GenerateMarkdown() error = %v", err)
	}

	content := testReadFile(t, outputPath)

	// Both should be listed with different IDs
	if !strings.Contains(string(content), "`test`") {
		t.Error("markdown should contain test ID")
	}
	if !strings.Contains(string(content), "`test_2`") {
		t.Error("markdown should contain test_2 ID for duplicate")
	}
}

func TestGenerateThemeFile_Success(t *testing.T) {
	t.Parallel()

	outputDir := t.TempDir()
	gen := NewGenerator(outputDir)

	theme := createTestTheme("Dracula", "#282a36")
	seenIDs := make(map[string]int)
	seenVarNames := make(map[string]int)

	info, err := gen.generateThemeFile(theme, seenIDs, seenVarNames)
	if err != nil {
		t.Fatalf("generateThemeFile() error = %v", err)
	}

	if info.ID != "dracula" {
		t.Errorf("ID = %q, want dracula", info.ID)
	}
	if info.VarName != "Dracula" {
		t.Errorf("VarName = %q, want Dracula", info.VarName)
	}
	if !info.IsDark {
		t.Error("Dracula should be detected as dark theme")
	}

	// Verify file exists and is valid Go
	path := filepath.Join(outputDir, "dracula.go")
	content := testReadFile(t, path)

	if !strings.Contains(string(content), "package themes") {
		t.Error("theme file should have package themes")
	}
	if !strings.Contains(string(content), `func (t *themeDracula) ID() string`) {
		t.Error("theme file should have ID method")
	}
}

func TestGenerateThemeFile_Duplicates(t *testing.T) {
	t.Parallel()

	outputDir := t.TempDir()
	gen := NewGenerator(outputDir)

	seenIDs := map[string]int{"test": 1}
	seenVarNames := map[string]int{"Test": 1}

	theme := createTestTheme("Test", "#000000")
	info, err := gen.generateThemeFile(theme, seenIDs, seenVarNames)
	if err != nil {
		t.Fatalf("generateThemeFile() error = %v", err)
	}

	if info.ID != "test_2" {
		t.Errorf("ID = %q, want test_2", info.ID)
	}
	if info.VarName != "Test2" {
		t.Errorf("VarName = %q, want Test2", info.VarName)
	}
}

func TestGenerateThemesFile_ValidThemes(t *testing.T) {
	t.Parallel()

	outputDir := t.TempDir()
	gen := NewGenerator(outputDir)

	themes := []themeInfo{
		{ID: "alpha", VarName: "Alpha", InstanceName: "themeAlphaInstance", DisplayName: "Alpha"},
		{ID: "beta", VarName: "Beta", InstanceName: "themeBetaInstance", DisplayName: "Beta"},
	}

	err := gen.generateThemesFile(themes)
	if err != nil {
		t.Fatalf("generateThemesFile() error = %v", err)
	}

	content := testReadFile(t, filepath.Join(outputDir, "themes.go"))

	// Check function definitions
	if !strings.Contains(string(content), "func All() []gothememe.Theme") {
		t.Error("themes.go should have All() function")
	}
	if !strings.Contains(string(content), "func ByID(id string) gothememe.Theme") {
		t.Error("themes.go should have ByID() function")
	}
	if !strings.Contains(string(content), "func IDs() []string") {
		t.Error("themes.go should have IDs() function")
	}

	// Check theme entries
	if !strings.Contains(string(content), "themeAlphaInstance") {
		t.Error("themes.go should reference themeAlphaInstance")
	}
	if !strings.Contains(string(content), "themeBetaInstance") {
		t.Error("themes.go should reference themeBetaInstance")
	}
}

func TestGenerateDocFile_ValidCount(t *testing.T) {
	t.Parallel()

	outputDir := t.TempDir()
	gen := NewGenerator(outputDir)

	err := gen.generateDocFile(451)
	if err != nil {
		t.Fatalf("generateDocFile() error = %v", err)
	}

	content := testReadFile(t, filepath.Join(outputDir, "doc.go"))

	if !strings.Contains(string(content), "Package themes provides 451+ pre-built") {
		t.Error("doc.go should contain theme count")
	}
	if !strings.Contains(string(content), "package themes") {
		t.Error("doc.go should have package declaration")
	}
}

func TestGenerateDocFile_ZeroCount(t *testing.T) {
	t.Parallel()

	outputDir := t.TempDir()
	gen := NewGenerator(outputDir)

	err := gen.generateDocFile(0)
	if err != nil {
		t.Fatalf("generateDocFile() error = %v", err)
	}

	content := testReadFile(t, filepath.Join(outputDir, "doc.go"))

	if !strings.Contains(string(content), "0+ pre-built") {
		t.Error("doc.go should handle zero count")
	}
}

// ============================================================================
// Helper Functions
// ============================================================================

func createTestTheme(name, background string) *WindowsTerminalTheme {
	return &WindowsTerminalTheme{
		Name:                name,
		Background:          background,
		Foreground:          "#f8f8f2",
		CursorColor:         "#f8f8f2",
		SelectionBackground: "#44475a",
		Black:               "#21222c",
		Red:                 "#ff5555",
		Green:               "#50fa7b",
		Yellow:              "#f1fa8c",
		Blue:                "#bd93f9",
		Purple:              "#ff79c6",
		Cyan:                "#8be9fd",
		White:               "#f8f8f2",
		BrightBlack:         "#6272a4",
		BrightRed:           "#ff6e6e",
		BrightGreen:         "#69ff94",
		BrightYellow:        "#ffffa5",
		BrightBlue:          "#d6acff",
		BrightPurple:        "#ff92df",
		BrightCyan:          "#a4ffff",
		BrightWhite:         "#ffffff",
	}
}
