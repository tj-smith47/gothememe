# GoThemeMe Implementation Plan

> **⚠️ IMPORTANT: DO NOT assign priority levels, defer items, or mark anything as "low priority" or "optional". DO NOT determine that any task is "not needed" or skip implementation. ALL incomplete items are equally essential and must be completed exactly as specified. The user determines importance and necessity, not the implementer.**

---

## Original Requirements

> We are going to create an open-source go theme package similar to lrstanley/bubbletint, but for web instead of in the terminal. Your task specifically is to create a checklist formatted plan in this directory that future sessions will execute on - that said, be thorough and do not skimp due to context limitations. Do not declare any features as "optional" or assign relative importance - if something improves code quality or UX, it is essential. It should have the same registry of themes, complete integration & contributing documentation, an MIT license, a goreleaser config and auto-tag (there is a reusable workflow for this, don't reinvent the wheel) and release workflows, a thorough golangci-lint config tuned to enforce all go best practices and modern Go (1.25) features / idioms / techniques. It should be able to be used by all modern web frameworks, and have support for code block formatting as well. The godoc comments / context should be immaculate - this is going to be a first of its' kind, official go web theming package. It should also support custom theme generation / registration.

**Key requirements from user:**
- Same registry of themes as bubbletint (**128+ themes** - research found bubbletint has 128+ themes)
- Complete integration & contributing documentation
- MIT license
- goreleaser config with auto-tag using `anothrNick/github-tag-action`
- Thorough golangci-lint config for Go 1.25
- Support for all modern web frameworks
- Support for code block formatting
- Immaculate godoc comments
- Custom theme generation/registration
- DTCG design tokens support (user confirmed)

**Repository**: `github.com/tj-smith47/gothememe`
**Go Version**: 1.25.5
**License**: MIT

---

## Implementation Status

### Legend
- [x] Complete
- [~] Partially complete (needs more work)
- [ ] Not started

---

## Session 5 Progress

### Accomplished
1. **Generator test coverage**: 25.3% → 88.6% with httptest mocks
2. **DRY refactoring in validation.go**: Replaced 27-case switch with `colorGetters` and `semanticColorGetters` maps
3. **DRY refactoring in builder.go**: Consolidated 16 ANSI if-blocks with `ansiColorEntry` struct
4. **Modern Go 1.25 idioms**:
   - `slices.SortFunc` with `cmp.Compare` in `SortByAccessibility()`
   - `slices.Index` in `findCurrentIndex()`
   - `map[string]struct{}` for set pattern in `countUnique()`
5. **SVG preview generation**: Created `internal/generator/svg.go` and `svg_test.go`
6. **AutoFixContrast function**: Implemented in `validation.go`
7. **Godoc examples**: Added 8 new examples (ValidateTheme, ValidateContrast, ValidateStrict, AnalyzeTheme, CompareThemes, AutoFixContrast, FilterAccessible, SortByAccessibility)
8. **Benchmarks**: Added CSS generation and registry operation benchmarks
9. **Fuzz tests**: Added FuzzHexParsing, FuzzRGBConstruction, FuzzColorManipulation

---

## Session 4 Progress

### Accomplished
1. **Created internal/colorutil package** - Shared luminance and hex parsing utilities
2. **Created internal/pairs package** - Shared contrast pairs for validation/analysis
3. **Added comprehensive tests**:
   - `default_registry_test.go` - Nil registry handling, delegated methods
   - `registry_test.go` - GetTheme, SetThemeDirect, UnregisterCurrent, EmptyNavigation
   - `output_test.go` - DefaultCSSOptions, DefaultSyntaxOptions, SCSS options
   - `tokens_test.go` - GenerateAllDesignTokens, DesignTokensWithDescriptions
   - `pkg/contrast/wcag_test.go` - MeetsAA_RGB, MeetsAAA_RGB, MeetsUIComponent, Check_RGB
   - `themes/themes_test.go` - AllThemeMethodsCoverage (100% coverage)
4. **Completed React example**:
   - `ThemeContext.tsx` - React context provider with localStorage
   - `ThemeSwitcher.tsx` - Dropdown component with navigation
   - `App.tsx`, `App.css`, `main.tsx` - Demo application
   - `package.json`, `tsconfig.json`, `vite.config.ts` - Build configuration
   - `generate/main.go` - Updated theme generator
5. **Created example READMEs**:
   - `examples/vanilla/README.md`
   - `examples/htmx/README.md`
   - `examples/templ/README.md`
   - `examples/echo/README.md`
   - `examples/react/README.md`
6. **Fixed linting** - All tests pass, 0 lint issues

### Current Coverage (Session 5 Update)
- Main package: **94.5%** (was 93.2%)
- internal/colorutil: **95.2%**
- internal/pairs: **100.0%**
- pkg/contrast: **97.4%**
- themes: **100.0%**
- internal/generator: **88.6%** (was 25.3%)
- **Total: 99.5%**

---

## Phase 1: Project Foundation

### 1.1 Repository Structure
- [x] Create directory structure
- [x] `.github/workflows/` - CI, release, auto-tag
- [x] `cmd/themegen/` - Theme generator CLI - FULLY WORKING
- [x] `internal/generator/` - Theme generation logic - IMPLEMENTED
- [x] `internal/colorutil/` - Shared color utilities
- [x] `internal/pairs/` - Shared contrast pairs
- [x] `themes/` - Generated theme files (451 themes)
- [x] `pkg/contrast/` - WCAG accessibility
- [x] `examples/` - Framework integration examples
- [x] `docs/` - CONTRIBUTING.md, INTEGRATION.md

### 1.2 Core Files
- [x] `go.mod` with Go 1.25.5
- [x] `LICENSE` (MIT)
- [x] `Makefile`
- [x] `.gitignore`

---

## Phase 2: Core Theme Interface & Types

### 2.1 Theme Interface (`theme.go`)
- [x] Define `Theme` interface with all 45+ methods
- [x] `BaseTheme` struct implementation with godoc comments

### 2.2 Supporting Types (`color.go`)
- [x] `Color` type with hex/RGB/HSL methods
- [x] OKLCH conversion (via go-colorful)
- [x] Color manipulation (Lighten, Darken, WithAlpha, Mix, etc.)
- [x] `SemanticColor` struct
- [x] Relative luminance calculation (uses internal/colorutil)

### 2.3 Registry (`registry.go`)
- [x] Thread-safe `Registry` struct
- [x] All CRUD operations (Register, Unregister, SetTheme, etc.)
- [x] Navigation (NextTheme, PreviousTheme)
- [x] Alphabetical sorting

### 2.4 Default Registry (`default_registry.go`)
- [x] Package-level facade functions
- [x] `NewDefaultRegistry()`

### 2.5 Theme Builder (`builder.go`)
- [x] Fluent `ThemeBuilder` API
- [x] Auto-derivation of missing colors
- [x] `GenerateThemeFromPalette()`
- [x] `DeriveTheme()` for creating variants

---

## Phase 3: Output Formats

### 3.1 CSS Output (`output.go`)
- [x] `GenerateCSS()` with options
- [x] `GenerateAllThemesCSS()` for multi-theme output
- [x] `data-theme` attribute selectors
- [x] Multiple color spaces (Hex, RGB, HSL, OKLCH)
- [x] Minification option
- [x] Metadata comments

### 3.2 SCSS Output
- [x] `GenerateSCSS()` function

### 3.3 JSON Output
- [x] `GenerateJSON()` function

### 3.4 Design Tokens (`tokens.go`)
- [x] DTCG v1 compliant output
- [x] `GenerateDesignTokens()` function
- [x] Hierarchical token structure

### 3.5 Syntax Highlighting
- [x] `GenerateSyntaxCSS()` function
- [x] Prism.js compatible output
- [x] Highlight.js compatible output
- [x] Chroma (Go) compatible output

---

## Phase 4: Theme Generation

### 4.1 Theme Generator (`cmd/themegen/`)
- [x] CLI with flags for output directory and source
- [x] iTerm2 Color Schemes fetcher - WORKING
- [x] Go file generation from fetched themes - WORKING
- [x] DEFAULT_THEMES.md generation - WORKING
- [x] SVG preview generation (`internal/generator/svg.go`)
- [ ] Integrate SVG into generator pipeline
- [ ] Update DEFAULT_THEMES.md with inline SVG previews

### 4.2 Built-in Themes
- [x] **451 themes generated** from iTerm2-Color-Schemes
- [x] Includes Dracula, Nord, Gruvbox, Tokyo Night, Catppuccin, Solarized, etc.

---

## Phase 5: Accessibility & Validation

### 5.1 WCAG Contrast Checking (`pkg/contrast/`)
- [x] `Luminance()` calculation
- [x] `Ratio()` / `RatioHex()` contrast ratio
- [x] `MeetsAA()` / `MeetsAAA()` compliance checks
- [x] `Check()` returns compliance level
- [x] `Level` type with String()
- [x] `Issue` struct (renamed from `ContrastIssue`)

### 5.2 Theme Validation (`validation.go`)
- [x] `ValidateTheme()` function
- [x] `ValidateContrast()` for all color pairs
- [x] `ValidateStrict()` returns error if validation fails
- [x] `AutoFixContrast()` - auto-adjust colors for compliance (Session 5)

### 5.3 Theme Analysis (`analysis.go`)
- [x] `AnalyzeTheme()` function
- [x] `ThemeStats` struct with ColorCount, UniqueColors, ContrastScore, AccessibilityPercent

---

## Phase 6: Framework Integration Examples

### 6.1 Vanilla HTML/CSS Example (`examples/vanilla/`)
- [x] HTTP server with theme CSS endpoint
- [x] Theme switching with localStorage
- [x] System preference detection (prefers-color-scheme)
- [x] README.md documentation

### 6.2 HTMX Example (`examples/htmx/`)
- [x] Go server with HTMX integration
- [x] Server-side theme rendering
- [x] Theme selector with live preview
- [x] README.md documentation

### 6.3 React Example (`examples/react/`)
- [x] `generate/main.go` - CSS generator
- [x] `src/ThemeContext.tsx` - React context provider
- [x] `src/ThemeSwitcher.tsx` - Dropdown component with useTheme
- [x] `src/App.tsx` - Demo application
- [x] `src/App.css` - Styles
- [x] `src/main.tsx` - Entry point
- [x] `package.json` - Dependencies
- [x] `tsconfig.json` - TypeScript config
- [x] `vite.config.ts` - Vite configuration
- [x] `index.html` - HTML entry
- [x] README.md documentation

### 6.4 templ Example (`examples/templ/`)
- [x] Pure templ example with Go server
- [x] README.md documentation

### 6.5 Echo Example (`examples/echo/`)
- [x] Echo framework integration example
- [x] README.md documentation

---

## Phase 7: Documentation

### 7.1 Godoc Comments
- [x] `doc.go` package documentation
- [x] All exported types documented
- [x] BaseTheme methods documented

### 7.2 README.md
- [x] Project description
- [x] Installation instructions
- [x] Quick start guide
- [x] Three usage patterns
- [x] Output format examples
- [x] Custom theme creation
- [x] Contributing section
- [ ] README badges (Go Reference, Go Report Card, CI Status, Coverage)

### 7.3 CONTRIBUTING.md
- [x] Development setup
- [x] Code style guidelines
- [x] Adding themes guide
- [x] PR process

### 7.4 INTEGRATION.md
- [x] Vanilla HTML/CSS guide
- [x] HTMX + templ guide
- [x] React guide
- [x] Vue guide
- [x] Svelte guide
- [x] Tailwind CSS guide

### 7.5 DEFAULT_THEMES.md
- [x] Theme listing with 451 themes
- [ ] SVG previews

### 7.6 CHANGELOG.md
- [x] v1.0.0 release notes

---

## Phase 8: Testing

### 8.1 Unit Tests
- [x] `color_test.go` - Color parsing, manipulation (with t.Parallel())
- [x] `registry_test.go` - Registry operations, concurrency (with t.Parallel())
- [x] `default_registry_test.go` - Default registry nil handling, delegated methods
- [x] `pkg/contrast/wcag_test.go` - WCAG calculations (with t.Parallel())
- [x] `output_test.go` - CSS/SCSS/JSON output tests
- [x] `tokens_test.go` - Design tokens validation
- [x] `builder_test.go` - ThemeBuilder tests
- [x] `validation_test.go` - Theme validation tests
- [x] `themes/themes_test.go` - Generated themes tests (100% coverage)
- [x] `internal/generator/generator_test.go` - Generator logic tests
- [x] `internal/colorutil/colorutil_test.go` - Color utility tests

### 8.2 Integration Tests
- [ ] CSS syntax validation
- [ ] DTCG schema validation

### 8.3 Example Tests
- [x] `example_test.go` for godoc
- [x] `ExampleValidateTheme()` - validation usage (Session 5)
- [x] `ExampleValidateContrast()` - contrast validation (Session 5)
- [x] `ExampleValidateStrict()` - strict validation (Session 5)
- [x] `ExampleAnalyzeTheme()` - theme analysis (Session 5)
- [x] `ExampleCompareThemes()` - theme comparison (Session 5)
- [x] `ExampleAutoFixContrast()` - auto-fix contrast (Session 5)
- [x] `ExampleFilterAccessible()` - filter accessible themes (Session 5)
- [x] `ExampleSortByAccessibility()` - sort by accessibility (Session 5)

### 8.4 Benchmarks
- [x] `BenchmarkRatioHex` in wcag_test.go
- [x] `BenchmarkLuminanceHex` in wcag_test.go
- [x] `BenchmarkRelativeLuminance` in colorutil_test.go
- [x] `BenchmarkHexToRGB` in colorutil_test.go
- [x] `BenchmarkContrastRatioHex` in colorutil_test.go
- [x] `BenchmarkGenerateCSS` in output_test.go (Session 5)
- [x] `BenchmarkGenerateCSS_Minified` in output_test.go (Session 5)
- [x] `BenchmarkGenerateAllThemesCSS` in output_test.go (Session 5)
- [x] `BenchmarkGenerateSyntaxCSS` in output_test.go (Session 5)
- [x] `BenchmarkGenerateJSON` in output_test.go (Session 5)
- [x] `BenchmarkRegistrySetTheme` in registry_test.go (Session 5)
- [x] `BenchmarkRegistryGetTheme` in registry_test.go (Session 5)
- [x] `BenchmarkRegistryNavigation` in registry_test.go (Session 5)

### 8.5 Fuzz Tests
- [x] `FuzzHexParsing` in color_test.go (Session 5)
- [x] `FuzzRGBConstruction` in color_test.go (Session 5)
- [x] `FuzzColorManipulation` in color_test.go (Session 5)
- [ ] `FuzzThemeIDLookup` in registry_test.go

### 8.6 Coverage
- [x] Main package: 94.5% (target: 80%+) - Session 5
- [x] internal/colorutil: 95.2%
- [x] internal/pairs: 100.0%
- [x] pkg/contrast: 97.4%
- [x] themes: 100.0%
- [x] internal/generator: 88.6% (was 25.3%) - Session 5
- [x] **Total: 99.5%** - Session 5

---

## Phase 9: CI/CD & Tooling

### 9.1 golangci-lint (`.golangci.yml`)
- [x] Complete v2 format configuration
- [x] All linters passing (zero issues)
- [x] paralleltest linter enabled
- [x] errcheck globally enabled

### 9.2 goreleaser (`.goreleaser.yaml`)
- [x] Multi-platform builds
- [x] Changelog generation
- [x] GitHub release configuration

### 9.3 GitHub Actions
- [x] CI workflow (`ci.yml`)
- [x] Auto-tag workflow (`auto-tag.yml`) - commit message parsing (#major, #minor, #patch, #none)
- [x] Release workflow (`release.yml`) - skip release support ([skip release])

### 9.4 Makefile
- [x] All standard targets

---

## Phase 10: Polish & Release Prep

### 10.1 Code Quality
- [x] All golangci-lint issues fixed properly
- [x] Main package 80%+ test coverage (93.2%)
- [ ] internal/generator 80%+ test coverage (25.3%)

### 10.2 Theme Quality
- [x] All 451 themes generated
- [ ] Run WCAG validation on all themes
- [ ] Generate SVG previews
- [x] Verify metadata

### 10.3 DRY Refactoring
- [x] Extract shared luminance utilities to internal/colorutil
- [x] Extract shared contrast pairs to internal/pairs
- [x] Extract semantic color derivation helper (`deriveSemanticColor()`) in builder.go
- [x] Replace switch with map in validation.go (`colorGetters`, `semanticColorGetters`) - Session 5
- [x] Replace switch with map in builder.go (`overrideHandlers`) - already done
- [x] Consolidate ANSI color setting (`ansiColorEntry`, `getANSIColorEntries()`) - Session 5
- [ ] Consolidate ValidateStrict/ValidateStrictAAA into shared internal
- [ ] Code generation for builder setters (go:generate)

### 10.4 Modern Go Idioms
- [x] Use `slices` package (`slices.SortFunc`, `slices.Index`) - Session 5
- [x] Use `cmp` package for comparisons (`cmp.Compare`) - Session 5
- [x] Use `map[string]struct{}` for sets - Session 5
- [ ] Use `maps` package where applicable
- [ ] Review error handling for best practices

### 10.5 Pre-Release Checklist
- [x] CHANGELOG.md created
- [ ] README badges working
- [ ] pkg.go.dev renders correctly

---

## Remaining Tasks Summary (Updated Session 5)

### ✅ COMPLETED in Session 5
- [x] Generator test coverage (25.3% → 88.6%)
- [x] SVG preview generation (`internal/generator/svg.go`)
- [x] AutoFixContrast function
- [x] DRY refactoring (colorGetters, semanticColorGetters, ansiColorEntry)
- [x] Modern Go idioms (slices.SortFunc, slices.Index, cmp.Compare, struct{})
- [x] 8 godoc examples
- [x] CSS/registry benchmarks
- [x] Color fuzz tests (3 tests)

### Still Incomplete

#### SVG Preview Integration
- [ ] Integrate SVG generation into generator pipeline (`generator.go`)
- [ ] Update markdown template for inline SVG previews
- [ ] Regenerate DEFAULT_THEMES.md with SVG previews

#### Additional Tests
- [ ] `FuzzThemeIDLookup` in registry_test.go
- [ ] Integration tests: CSS syntax validation, DTCG schema validation

#### DRY Refactoring
- [ ] Consolidate `ValidateStrict()` and `ValidateStrictAAA()` into shared internal
- [ ] go:generate for builder setters

#### README Badges
- [ ] Add Go Reference badge
- [ ] Add Go Report Card badge
- [ ] Add CI Status badge
- [ ] Add Coverage badge

#### Theme Validation
- [ ] Run WCAG validation on all 451 themes
- [ ] Document any themes with contrast issues

#### Modern Go
- [ ] Use `maps` package where applicable
- [ ] Review error handling for best practices

#### pkg.go.dev
- [ ] Verify documentation renders correctly after first release

---

## File Structure (Current)

```
gothememe/
├── .github/
│   └── workflows/
│       ├── ci.yml
│       ├── release.yml
│       └── auto-tag.yml
├── cmd/
│   └── themegen/
│       └── main.go
├── docs/
│   ├── CONTRIBUTING.md
│   └── INTEGRATION.md
├── examples/
│   ├── echo/
│   │   ├── go.mod
│   │   ├── main.go
│   │   └── README.md
│   ├── htmx/
│   │   ├── go.mod
│   │   ├── main.go
│   │   └── README.md
│   ├── react/
│   │   ├── generate/
│   │   │   ├── go.mod
│   │   │   └── main.go
│   │   ├── public/
│   │   │   ├── themes.css
│   │   │   └── themes.json
│   │   ├── src/
│   │   │   ├── App.css
│   │   │   ├── App.tsx
│   │   │   ├── main.tsx
│   │   │   ├── ThemeContext.tsx
│   │   │   └── ThemeSwitcher.tsx
│   │   ├── index.html
│   │   ├── package.json
│   │   ├── tsconfig.json
│   │   ├── tsconfig.node.json
│   │   ├── vite.config.ts
│   │   └── README.md
│   ├── templ/
│   │   ├── go.mod
│   │   ├── main.go
│   │   └── README.md
│   └── vanilla/
│       ├── go.mod
│       ├── main.go
│       └── README.md
├── internal/
│   ├── colorutil/
│   │   ├── colorutil.go
│   │   └── colorutil_test.go
│   ├── generator/
│   │   ├── types.go
│   │   ├── fetcher.go
│   │   ├── generator.go
│   │   ├── generator_test.go
│   │   ├── svg.go              # NEW - Session 5
│   │   └── svg_test.go         # NEW - Session 5
│   └── pairs/
│       ├── pairs.go
│       └── pairs_test.go
├── pkg/
│   └── contrast/
│       ├── wcag.go
│       └── wcag_test.go
├── themes/
│   ├── doc.go
│   ├── themes.go
│   ├── themes_test.go
│   └── theme_*.go (451 files)
├── analysis.go
├── analysis_test.go
├── builder.go
├── builder_test.go
├── color.go
├── color_test.go
├── default_registry.go
├── default_registry_test.go
├── doc.go
├── example_test.go
├── go.mod
├── go.sum
├── .golangci.yml
├── .goreleaser.yaml
├── .gitignore
├── LICENSE
├── Makefile
├── output.go
├── output_test.go
├── README.md
├── registry.go
├── registry_test.go
├── theme.go
├── themes.go
├── tokens.go
├── tokens_test.go
├── validation.go
├── validation_test.go
├── CHANGELOG.md
├── DEFAULT_THEMES.md
└── PLAN.md
```

---

## Dependencies

```go
require github.com/lucasb-eyer/go-colorful v1.2.0

tool github.com/golangci/golangci-lint/v2/cmd/golangci-lint
```

---

## Verification Commands

```bash
# Build
cd /db/appdata/gothememe && go build ./...

# Test with race detection
cd /db/appdata/gothememe && go test -race ./...

# Test with coverage
cd /db/appdata/gothememe && go test -cover ./...

# Regenerate themes (fetches from GitHub)
cd /db/appdata/gothememe && go run ./cmd/themegen -output themes

# Run linter
cd /db/appdata/gothememe && go tool golangci-lint run

# Build examples
cd /db/appdata/gothememe/examples/vanilla && go build .
cd /db/appdata/gothememe/examples/htmx && go build .
cd /db/appdata/gothememe/examples/templ && go build .
cd /db/appdata/gothememe/examples/echo && go build .
cd /db/appdata/gothememe/examples/react/generate && go build .
```

---

## Success Criteria (from original requirements)

- [x] All 128+ themes generate valid CSS/SCSS/JSON/DTCG output (451 themes!)
- [ ] WCAG AA contrast validation passes for all themes
- [x] Full godoc coverage with examples (20 examples total) - Session 5
- [x] CI/CD pipeline working with auto-tagging
- [x] Framework examples working (vanilla, htmx, templ, echo, react complete)
- [x] golangci-lint passes with zero warnings
- [x] 80%+ test coverage (main: 94.5%, generator: 88.6%, total: 99.5%) - Session 5
- [ ] pkg.go.dev documentation renders correctly
