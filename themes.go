package gothememe

// DefaultThemes returns an empty slice for the main package.
// To get all built-in themes, import the themes subpackage:
//
//	import "github.com/tj-smith47/gothememe/themes"
//
//	allThemes := themes.All()
//
// This design avoids import cycles while keeping the themes
// package cleanly separated.
func DefaultThemes() []Theme {
	return nil
}

// Note: Direct theme access is available via the themes subpackage:
//
//	import "github.com/tj-smith47/gothememe/themes"
//
//	theme := themes.ThemeDracula
//	theme := themes.ByID("nord")
//	allThemes := themes.All()
//
// Popular themes include:
//   - themes.ThemeDracula
//   - themes.ThemeNord
//   - themes.ThemeGruvboxDark
//   - themes.ThemeSolarizedDark
//   - themes.ThemeCatppuccinMocha
//   - themes.ThemeTokyoNight
//   - And 445+ more...
