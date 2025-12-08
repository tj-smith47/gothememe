# Custom Themes

Custom themes allow you to add your own manually-crafted themes that are preserved during theme regeneration.

## How It Works

When you run `go run ./cmd/themegen`, the generator:

1. Fetches themes from iTerm2-Color-Schemes
2. Scans for `custom_*.go` files in the `themes/` directory
3. Merges both sets into the generated registry files
4. Generates SVG previews for all themes

Files matching `custom_*.go` will **not** be deleted or overwritten during regeneration.

## Adding a Custom Theme

1. Create a new `.go` file with the `custom_` prefix (e.g., `custom_my_theme.go`)
2. Implement the theme following the pattern in `custom_dracula_pro.go`:

```go
package themes

import "github.com/tj-smith47/gothememe"

var themeMyThemeInstance gothememe.Theme = &themeMyTheme{}

type themeMyTheme struct{}

func (t *themeMyTheme) ID() string          { return "my_theme" }
func (t *themeMyTheme) DisplayName() string { return "My Theme" }
// ... implement all Theme interface methods
```

3. Run `go run ./cmd/themegen` to regenerate the registry
4. Your theme will be available as `themes.ThemeMyTheme` and `themes.ByID("my_theme")`

## Theme Requirements

- File must be in package `themes`
- Must have a variable named `theme<Name>Instance` of type `gothememe.Theme`
- Must implement the full `gothememe.Theme` interface
- ID should match the filename (without `.go` extension)

## Examples

See `custom_dracula_pro.go` for a complete example of a custom theme.

## Important Notes

- Custom theme files **must** start with `custom_` prefix
- All custom themes must be in the `themes/` directory (not subdirectories)
- The generator will automatically include them in the registry
- SVG previews are preserved if they exist, or auto-generated as placeholders
