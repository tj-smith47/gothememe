// Command themegen generates Go theme files from external theme sources.
//
// Usage:
//
//	themegen [flags]
//
// Flags:
//
//	-output string
//	      Output directory for generated files (default "themes")
//	-source string
//	      Theme source to fetch from (default "iterm2")
//
// Sources:
//
//	iterm2    - iTerm2-Color-Schemes repository (Windows Terminal JSON format)
//
// Example:
//
//	themegen -source iterm2 -output themes
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tj-smith47/gothememe/internal/generator"
)

// Version information set by goreleaser.
var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

func main() {
	var (
		outputDir = flag.String("output", "themes", "Output directory for generated files")
		source    = flag.String("source", "iterm2", "Theme source (iterm2)")
		version   = flag.Bool("version", false, "Print version information")
	)

	flag.Parse()

	if *version {
		fmt.Printf("themegen %s (%s) built %s\n", Version, Commit, Date)
		os.Exit(0)
	}

	fmt.Printf("GoThemeMe Theme Generator\n")
	fmt.Printf("Version: %s\n", Version)
	fmt.Printf("Source: %s\n", *source)
	fmt.Printf("Output: %s\n", *outputDir)
	fmt.Println()

	switch *source {
	case "iterm2":
		if err := generateFromITerm2(*outputDir); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown source: %s\n", *source)
		os.Exit(1)
	}
}

func generateFromITerm2(outputDir string) error {
	fmt.Println("Fetching themes from iTerm2-Color-Schemes...")

	fetcher := generator.NewFetcher()

	// List available themes
	list, err := fetcher.ListThemes()
	if err != nil {
		return fmt.Errorf("listing themes: %w", err)
	}
	fmt.Printf("Found %d themes\n", len(list))

	// Fetch all themes
	themes, err := fetcher.FetchAllThemes()
	if err != nil {
		return fmt.Errorf("fetching themes: %w", err)
	}
	fmt.Printf("Successfully fetched %d themes\n", len(themes))

	// Ensure output directory is absolute or relative to current dir
	absOutput, err := filepath.Abs(outputDir)
	if err != nil {
		return fmt.Errorf("resolving output path: %w", err)
	}

	// Generate Go files
	fmt.Printf("Generating Go files in %s...\n", absOutput)
	gen := generator.NewGenerator(absOutput)
	gen.SetThemes(themes)

	if err := gen.Generate(); err != nil {
		return fmt.Errorf("generating themes: %w", err)
	}

	// Generate DEFAULT_THEMES.md in the parent directory
	mdPath := filepath.Join(filepath.Dir(absOutput), "DEFAULT_THEMES.md")
	fmt.Printf("Generating %s...\n", mdPath)
	if err := gen.GenerateMarkdown(mdPath); err != nil {
		return fmt.Errorf("generating markdown: %w", err)
	}

	fmt.Printf("\nSuccessfully generated %d theme files!\n", len(themes))
	fmt.Println()
	fmt.Println("To use the themes, import:")
	fmt.Println()
	fmt.Println("    import \"github.com/tj-smith47/gothememe/themes\"")
	fmt.Println()
	fmt.Println("Then access themes directly:")
	fmt.Println()
	fmt.Println("    theme := themes.ThemeDracula")
	fmt.Println("    theme := themes.ByID(\"nord\")")
	fmt.Println()

	return nil
}
