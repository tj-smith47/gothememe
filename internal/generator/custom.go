package generator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// CustomTheme represents metadata extracted from a custom theme file.
type CustomTheme struct {
	ID          string
	VarName     string
	DisplayName string
	IsDark      bool
	FileName    string
}

// ScanCustomThemes scans the themes directory for custom_*.go theme files
// and extracts their metadata.
func ScanCustomThemes(themesDir string) ([]CustomTheme, error) {
	// Check if themes directory exists
	if _, err := os.Stat(themesDir); os.IsNotExist(err) {
		return nil, nil // No themes directory
	}

	entries, err := os.ReadDir(themesDir)
	if err != nil {
		return nil, fmt.Errorf("reading themes directory: %w", err)
	}

	var customs []CustomTheme
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		// Only process files that start with "custom_" and end with ".go"
		if !strings.HasPrefix(entry.Name(), "custom_") || !strings.HasSuffix(entry.Name(), ".go") {
			continue
		}

		// Skip test files
		if strings.HasSuffix(entry.Name(), "_test.go") {
			continue
		}

		filePath := filepath.Join(themesDir, entry.Name())
		custom, err := parseCustomTheme(filePath)
		if err != nil {
			fmt.Printf("Warning: failed to parse custom theme %s: %v\n", entry.Name(), err)
			continue
		}

		customs = append(customs, custom)
	}

	return customs, nil
}

// parseCustomTheme extracts metadata from a custom theme Go file.
func parseCustomTheme(filePath string) (CustomTheme, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return CustomTheme{}, fmt.Errorf("parsing file: %w", err)
	}

	var custom CustomTheme
	custom.FileName = filepath.Base(filePath)

	// Extract ID and other metadata from function implementations
	ast.Inspect(node, func(n ast.Node) bool {
		// Look for method declarations
		if fn, ok := n.(*ast.FuncDecl); ok && fn.Recv != nil {
			methodName := fn.Name.Name

			// Extract string literals from return statements
			if fn.Body != nil && len(fn.Body.List) > 0 {
				if ret, ok := fn.Body.List[0].(*ast.ReturnStmt); ok && len(ret.Results) > 0 {
					if lit, ok := ret.Results[0].(*ast.BasicLit); ok && lit.Kind == token.STRING {
						value := strings.Trim(lit.Value, "\"")

						switch methodName {
						case "ID":
							custom.ID = value
						case "DisplayName":
							custom.DisplayName = value
						}
					} else if ident, ok := ret.Results[0].(*ast.Ident); ok {
						// Handle boolean returns like "true" or "false"
						if methodName == "IsDark" {
							custom.IsDark = ident.Name == "true"
						}
					}
				}
			}
		}

		// Look for variable declarations (themeXInstance)
		if decl, ok := n.(*ast.GenDecl); ok && decl.Tok == token.VAR {
			for _, spec := range decl.Specs {
				if vspec, ok := spec.(*ast.ValueSpec); ok {
					for _, name := range vspec.Names {
						if strings.HasSuffix(name.Name, "Instance") {
							// Extract VarName from instance variable name
							// e.g., "themeDraculaProInstance" -> "DraculaPro"
							instanceName := name.Name
							if strings.HasPrefix(instanceName, "theme") && strings.HasSuffix(instanceName, "Instance") {
								custom.VarName = strings.TrimSuffix(strings.TrimPrefix(instanceName, "theme"), "Instance")
							}
						}
					}
				}
			}
		}

		return true
	})

	// Validate we got the required fields
	if custom.ID == "" {
		return CustomTheme{}, fmt.Errorf("could not extract ID from theme")
	}
	if custom.VarName == "" {
		// Try to derive VarName from ID
		custom.VarName = toVarName(custom.ID)
	}
	if custom.DisplayName == "" {
		custom.DisplayName = custom.ID
	}

	return custom, nil
}
