# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0] - 2025-12-07

### Added

#### Core Theme System
- `Theme` interface with 45+ methods for comprehensive color access
- `BaseTheme` implementation with full documentation
- `ThemeBuilder` for fluent theme construction
- Auto-derivation of secondary colors from primary colors
- Theme metadata support (author, license, source, description)

#### Color System
- `Color` type with hex, RGB, HSL, and OKLCH support
- CSS color output in multiple formats
- Color manipulation utilities (lighten, darken, alpha)
- Thread-safe color operations

#### Theme Registry
- `Registry` for managing multiple themes
- Thread-safe theme switching with `sync.RWMutex`
- Theme navigation (next/previous)
- Implements `Theme` interface for seamless integration

#### Output Formats
- **CSS**: Variables with configurable prefix, selector, and color space
- **SCSS**: Variables and mixins
- **JSON**: Key-value color mapping
- **DTCG**: Design Tokens Community Group v1 compliant tokens

#### Syntax Highlighting
- Prism.js CSS generation
- Highlight.js CSS generation
- Chroma (Go) CSS generation

#### Pre-built Themes
- 451 themes from iTerm2-Color-Schemes
- Automatic dark/light mode detection
- Full ANSI 16-color support
- Code syntax highlighting colors

#### Validation & Analysis
- `ValidateTheme()` for checking theme completeness
- `ValidateContrast()` for WCAG AA/AAA compliance
- `ValidateStrict()` for CI/CD integration
- `AnalyzeTheme()` for color statistics
- `FilterAccessible()` for finding accessible themes
- `SortByAccessibility()` for ranking themes

#### WCAG Contrast
- `pkg/contrast` package for accessibility calculations
- WCAG 2.1 compliant luminance calculations
- Contrast ratio calculations (1:1 to 21:1 scale)
- AA and AAA level checking

#### Examples
- Vanilla HTML/CSS with theme switcher
- Echo framework with theme API
- HTMX with server-side rendering
- Templ integration example
- React CSS/JSON generation

### Technical Details

- Go 1.25+ required
- Zero external runtime dependencies (stdlib only)
- golangci-lint clean with strict settings
- Parallel test execution throughout
- 70%+ test coverage

### Theme Categories

Each theme includes colors for:
- **Backgrounds**: primary, secondary, surface, surface-secondary
- **Text**: primary, secondary, muted, inverted
- **Accent**: primary, secondary, brand
- **Borders**: default, subtle, strong
- **Semantic**: success, warning, error, info (each with background/border/text)
- **ANSI**: 16 terminal colors (8 standard + 8 bright)
- **Code**: 12 syntax highlighting colors

[Unreleased]: https://github.com/tj-smith47/gothememe/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/tj-smith47/gothememe/releases/tag/v1.0.0
