package generator

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	// BaseURL is the raw GitHub content URL for iTerm2-Color-Schemes.
	BaseURL = "https://raw.githubusercontent.com/mbadolato/iTerm2-Color-Schemes/master"

	// APIURL is the GitHub API URL for listing themes.
	APIURL = "https://api.github.com/repos/mbadolato/iTerm2-Color-Schemes/contents/windowsterminal"
)

// GitHubContent represents a file entry from the GitHub API.
type GitHubContent struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	DownloadURL string `json:"download_url"`
	Type        string `json:"type"`
}

// Fetcher handles fetching themes from external sources.
type Fetcher struct {
	client  *http.Client
	baseURL string // Override for testing, empty uses default APIURL
}

// NewFetcher creates a new theme fetcher.
func NewFetcher() *Fetcher {
	return &Fetcher{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// WithBaseURL returns a new Fetcher with a custom base URL for testing.
func (f *Fetcher) WithBaseURL(url string) *Fetcher {
	return &Fetcher{
		client:  f.client,
		baseURL: url,
	}
}

// WithClient returns a new Fetcher with a custom HTTP client.
func (f *Fetcher) WithClient(client *http.Client) *Fetcher {
	return &Fetcher{
		client:  client,
		baseURL: f.baseURL,
	}
}

// apiURL returns the API URL to use (custom or default).
func (f *Fetcher) apiURL() string {
	if f.baseURL != "" {
		return f.baseURL
	}
	return APIURL
}

// ListThemes returns a list of available theme files from the source.
func (f *Fetcher) ListThemes() ([]GitHubContent, error) {
	req, err := http.NewRequest(http.MethodGet, f.apiURL(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("User-Agent", "gothememe-themegen/1.0")

	resp, err := f.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetching theme list: %w", err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			fmt.Printf("warning: failed to close response body: %v\n", cerr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return nil, fmt.Errorf("GitHub API returned %d (failed to read body: %w)", resp.StatusCode, readErr)
		}
		return nil, fmt.Errorf("GitHub API returned %d: %s", resp.StatusCode, string(body))
	}

	var contents []GitHubContent
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	// Filter to only JSON files
	var themes []GitHubContent
	for _, c := range contents {
		if c.Type == "file" && strings.HasSuffix(c.Name, ".json") {
			themes = append(themes, c)
		}
	}

	return themes, nil
}

// FetchTheme downloads and parses a single theme file.
func (f *Fetcher) FetchTheme(url string) (*WindowsTerminalTheme, error) {
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("User-Agent", "gothememe-themegen/1.0")

	resp, err := f.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetching theme: %w", err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			fmt.Printf("warning: failed to close response body: %v\n", cerr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d fetching %s", resp.StatusCode, url)
	}

	var theme WindowsTerminalTheme
	if err := json.NewDecoder(resp.Body).Decode(&theme); err != nil {
		return nil, fmt.Errorf("decoding theme: %w", err)
	}

	return &theme, nil
}

// FetchAllThemes downloads all themes from the source.
func (f *Fetcher) FetchAllThemes() ([]*WindowsTerminalTheme, error) {
	list, err := f.ListThemes()
	if err != nil {
		return nil, err
	}

	var themes []*WindowsTerminalTheme
	for _, entry := range list {
		theme, err := f.FetchTheme(entry.DownloadURL)
		if err != nil {
			// Log error but continue with other themes
			fmt.Printf("Warning: failed to fetch %s: %v\n", entry.Name, err)
			continue
		}
		themes = append(themes, theme)
	}

	return themes, nil
}
