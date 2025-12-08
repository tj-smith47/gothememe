# Default Themes

GoThemeMe includes **452** pre-built themes sourced from [iTerm2-Color-Schemes](https://github.com/mbadolato/iTerm2-Color-Schemes).

## Usage

```go
import "github.com/tj-smith47/gothememe/themes"

// Access by variable
theme := themes.ThemeDracula

// Access by ID
theme := themes.ByID("dracula")

// List all themes
for _, t := range themes.All() {
    fmt.Println(t.ID(), "-", t.DisplayName())
}
```

## Theme List

| Preview | ID | Display Name | Type | Variable |
|---------|----|--------------| ------|----------|
| ![0x96f](themes/0x96f.svg) | `0x96f` | 0x96f | 🌙 Dark | `themes.ThemeN0x96f` |
| ![12-bit Rainbow](themes/12_bit_rainbow.svg) | `12_bit_rainbow` | 12-bit Rainbow | 🌙 Dark | `themes.ThemeN12BitRainbow` |
| ![3024 Day](themes/3024_day.svg) | `3024_day` | 3024 Day | ☀️ Light | `themes.ThemeN3024Day` |
| ![3024 Night](themes/3024_night.svg) | `3024_night` | 3024 Night | 🌙 Dark | `themes.ThemeN3024Night` |
| ![Aardvark Blue](themes/aardvark_blue.svg) | `aardvark_blue` | Aardvark Blue | 🌙 Dark | `themes.ThemeAardvarkBlue` |
| ![Abernathy](themes/abernathy.svg) | `abernathy` | Abernathy | 🌙 Dark | `themes.ThemeAbernathy` |
| ![Adventure](themes/adventure.svg) | `adventure` | Adventure | 🌙 Dark | `themes.ThemeAdventure` |
| ![Adventure Time](themes/adventure_time.svg) | `adventure_time` | Adventure Time | 🌙 Dark | `themes.ThemeAdventureTime` |
| ![Adwaita](themes/adwaita.svg) | `adwaita` | Adwaita | ☀️ Light | `themes.ThemeAdwaita` |
| ![Adwaita Dark](themes/adwaita_dark.svg) | `adwaita_dark` | Adwaita Dark | 🌙 Dark | `themes.ThemeAdwaitaDark` |
| ![Afterglow](themes/afterglow.svg) | `afterglow` | Afterglow | 🌙 Dark | `themes.ThemeAfterglow` |
| ![Alabaster](themes/alabaster.svg) | `alabaster` | Alabaster | ☀️ Light | `themes.ThemeAlabaster` |
| ![Alien Blood](themes/alien_blood.svg) | `alien_blood` | Alien Blood | 🌙 Dark | `themes.ThemeAlienBlood` |
| ![Andromeda](themes/andromeda.svg) | `andromeda` | Andromeda | 🌙 Dark | `themes.ThemeAndromeda` |
| ![Apple Classic](themes/apple_classic.svg) | `apple_classic` | Apple Classic | 🌙 Dark | `themes.ThemeAppleClassic` |
| ![Apple System Colors](themes/apple_system_colors.svg) | `apple_system_colors` | Apple System Colors | 🌙 Dark | `themes.ThemeAppleSystemColors` |
| ![Apple System Colors Light](themes/apple_system_colors_light.svg) | `apple_system_colors_light` | Apple System Colors Light | ☀️ Light | `themes.ThemeAppleSystemColorsLight` |
| ![Arcoiris](themes/arcoiris.svg) | `arcoiris` | Arcoiris | 🌙 Dark | `themes.ThemeArcoiris` |
| ![Ardoise](themes/ardoise.svg) | `ardoise` | Ardoise | 🌙 Dark | `themes.ThemeArdoise` |
| ![Argonaut](themes/argonaut.svg) | `argonaut` | Argonaut | 🌙 Dark | `themes.ThemeArgonaut` |
| ![Arthur](themes/arthur.svg) | `arthur` | Arthur | 🌙 Dark | `themes.ThemeArthur` |
| ![Atelier Sulphurpool](themes/atelier_sulphurpool.svg) | `atelier_sulphurpool` | Atelier Sulphurpool | 🌙 Dark | `themes.ThemeAtelierSulphurpool` |
| ![Atom](themes/atom.svg) | `atom` | Atom | 🌙 Dark | `themes.ThemeAtom` |
| ![Atom One Dark](themes/atom_one_dark.svg) | `atom_one_dark` | Atom One Dark | 🌙 Dark | `themes.ThemeAtomOneDark` |
| ![Atom One Light](themes/atom_one_light.svg) | `atom_one_light` | Atom One Light | ☀️ Light | `themes.ThemeAtomOneLight` |
| ![Aura](themes/aura.svg) | `aura` | Aura | 🌙 Dark | `themes.ThemeAura` |
| ![Aurora](themes/aurora.svg) | `aurora` | Aurora | 🌙 Dark | `themes.ThemeAurora` |
| ![Ayu](themes/ayu.svg) | `ayu` | Ayu | 🌙 Dark | `themes.ThemeAyu` |
| ![Ayu Light](themes/ayu_light.svg) | `ayu_light` | Ayu Light | ☀️ Light | `themes.ThemeAyuLight` |
| ![Ayu Mirage](themes/ayu_mirage.svg) | `ayu_mirage` | Ayu Mirage | 🌙 Dark | `themes.ThemeAyuMirage` |
| ![Banana Blueberry](themes/banana_blueberry.svg) | `banana_blueberry` | Banana Blueberry | 🌙 Dark | `themes.ThemeBananaBlueberry` |
| ![Batman](themes/batman.svg) | `batman` | Batman | 🌙 Dark | `themes.ThemeBatman` |
| ![Belafonte Day](themes/belafonte_day.svg) | `belafonte_day` | Belafonte Day | 🌙 Dark | `themes.ThemeBelafonteDay` |
| ![Belafonte Night](themes/belafonte_night.svg) | `belafonte_night` | Belafonte Night | 🌙 Dark | `themes.ThemeBelafonteNight` |
| ![Birds Of Paradise](themes/birds_of_paradise.svg) | `birds_of_paradise` | Birds Of Paradise | 🌙 Dark | `themes.ThemeBirdsOfParadise` |
| ![Black Metal](themes/black_metal.svg) | `black_metal` | Black Metal | 🌙 Dark | `themes.ThemeBlackMetal` |
| ![Black Metal (Bathory)](themes/black_metal_bathory.svg) | `black_metal_bathory` | Black Metal (Bathory) | 🌙 Dark | `themes.ThemeBlackMetalBathory` |
| ![Black Metal (Burzum)](themes/black_metal_burzum.svg) | `black_metal_burzum` | Black Metal (Burzum) | 🌙 Dark | `themes.ThemeBlackMetalBurzum` |
| ![Black Metal (Dark Funeral)](themes/black_metal_dark_funeral.svg) | `black_metal_dark_funeral` | Black Metal (Dark Funeral) | 🌙 Dark | `themes.ThemeBlackMetalDarkFuneral` |
| ![Black Metal (Gorgoroth)](themes/black_metal_gorgoroth.svg) | `black_metal_gorgoroth` | Black Metal (Gorgoroth) | 🌙 Dark | `themes.ThemeBlackMetalGorgoroth` |
| ![Black Metal (Immortal)](themes/black_metal_immortal.svg) | `black_metal_immortal` | Black Metal (Immortal) | 🌙 Dark | `themes.ThemeBlackMetalImmortal` |
| ![Black Metal (Khold)](themes/black_metal_khold.svg) | `black_metal_khold` | Black Metal (Khold) | 🌙 Dark | `themes.ThemeBlackMetalKhold` |
| ![Black Metal (Marduk)](themes/black_metal_marduk.svg) | `black_metal_marduk` | Black Metal (Marduk) | 🌙 Dark | `themes.ThemeBlackMetalMarduk` |
| ![Black Metal (Mayhem)](themes/black_metal_mayhem.svg) | `black_metal_mayhem` | Black Metal (Mayhem) | 🌙 Dark | `themes.ThemeBlackMetalMayhem` |
| ![Black Metal (Nile)](themes/black_metal_nile.svg) | `black_metal_nile` | Black Metal (Nile) | 🌙 Dark | `themes.ThemeBlackMetalNile` |
| ![Black Metal (Venom)](themes/black_metal_venom.svg) | `black_metal_venom` | Black Metal (Venom) | 🌙 Dark | `themes.ThemeBlackMetalVenom` |
| ![Blazer](themes/blazer.svg) | `blazer` | Blazer | 🌙 Dark | `themes.ThemeBlazer` |
| ![Blue Berry Pie](themes/blue_berry_pie.svg) | `blue_berry_pie` | Blue Berry Pie | 🌙 Dark | `themes.ThemeBlueBerryPie` |
| ![Blue Dolphin](themes/blue_dolphin.svg) | `blue_dolphin` | Blue Dolphin | 🌙 Dark | `themes.ThemeBlueDolphin` |
| ![Blue Matrix](themes/blue_matrix.svg) | `blue_matrix` | Blue Matrix | 🌙 Dark | `themes.ThemeBlueMatrix` |
| ![Bluloco Dark](themes/bluloco_dark.svg) | `bluloco_dark` | Bluloco Dark | 🌙 Dark | `themes.ThemeBlulocoDark` |
| ![Bluloco Light](themes/bluloco_light.svg) | `bluloco_light` | Bluloco Light | ☀️ Light | `themes.ThemeBlulocoLight` |
| ![Borland](themes/borland.svg) | `borland` | Borland | 🌙 Dark | `themes.ThemeBorland` |
| ![Box](themes/box.svg) | `box` | Box | 🌙 Dark | `themes.ThemeBox` |
| ![branch](themes/branch.svg) | `branch` | branch | 🌙 Dark | `themes.ThemeBranch` |
| ![Breadog](themes/breadog.svg) | `breadog` | Breadog | ☀️ Light | `themes.ThemeBreadog` |
| ![Breeze](themes/breeze.svg) | `breeze` | Breeze | 🌙 Dark | `themes.ThemeBreeze` |
| ![Bright Lights](themes/bright_lights.svg) | `bright_lights` | Bright Lights | 🌙 Dark | `themes.ThemeBrightLights` |
| ![Broadcast](themes/broadcast.svg) | `broadcast` | Broadcast | 🌙 Dark | `themes.ThemeBroadcast` |
| ![Brogrammer](themes/brogrammer.svg) | `brogrammer` | Brogrammer | 🌙 Dark | `themes.ThemeBrogrammer` |
| ![Builtin Dark](themes/builtin_dark.svg) | `builtin_dark` | Builtin Dark | 🌙 Dark | `themes.ThemeBuiltinDark` |
| ![Builtin Light](themes/builtin_light.svg) | `builtin_light` | Builtin Light | ☀️ Light | `themes.ThemeBuiltinLight` |
| ![Builtin Pastel Dark](themes/builtin_pastel_dark.svg) | `builtin_pastel_dark` | Builtin Pastel Dark | 🌙 Dark | `themes.ThemeBuiltinPastelDark` |
| ![Builtin Solarized Dark](themes/builtin_solarized_dark.svg) | `builtin_solarized_dark` | Builtin Solarized Dark | 🌙 Dark | `themes.ThemeBuiltinSolarizedDark` |
| ![Builtin Solarized Light](themes/builtin_solarized_light.svg) | `builtin_solarized_light` | Builtin Solarized Light | ☀️ Light | `themes.ThemeBuiltinSolarizedLight` |
| ![Builtin Tango Dark](themes/builtin_tango_dark.svg) | `builtin_tango_dark` | Builtin Tango Dark | 🌙 Dark | `themes.ThemeBuiltinTangoDark` |
| ![Builtin Tango Light](themes/builtin_tango_light.svg) | `builtin_tango_light` | Builtin Tango Light | ☀️ Light | `themes.ThemeBuiltinTangoLight` |
| ![C64](themes/c64.svg) | `c64` | C64 | 🌙 Dark | `themes.ThemeC64` |
| ![Calamity](themes/calamity.svg) | `calamity` | Calamity | 🌙 Dark | `themes.ThemeCalamity` |
| ![Carbonfox](themes/carbonfox.svg) | `carbonfox` | Carbonfox | 🌙 Dark | `themes.ThemeCarbonfox` |
| ![Catppuccin Frappe](themes/catppuccin_frappe.svg) | `catppuccin_frappe` | Catppuccin Frappe | 🌙 Dark | `themes.ThemeCatppuccinFrappe` |
| ![Catppuccin Latte](themes/catppuccin_latte.svg) | `catppuccin_latte` | Catppuccin Latte | ☀️ Light | `themes.ThemeCatppuccinLatte` |
| ![Catppuccin Macchiato](themes/catppuccin_macchiato.svg) | `catppuccin_macchiato` | Catppuccin Macchiato | 🌙 Dark | `themes.ThemeCatppuccinMacchiato` |
| ![Catppuccin Mocha](themes/catppuccin_mocha.svg) | `catppuccin_mocha` | Catppuccin Mocha | 🌙 Dark | `themes.ThemeCatppuccinMocha` |
| ![CGA](themes/cga.svg) | `cga` | CGA | 🌙 Dark | `themes.ThemeCga` |
| ![Chalk](themes/chalk.svg) | `chalk` | Chalk | 🌙 Dark | `themes.ThemeChalk` |
| ![Chalkboard](themes/chalkboard.svg) | `chalkboard` | Chalkboard | 🌙 Dark | `themes.ThemeChalkboard` |
| ![Challenger Deep](themes/challenger_deep.svg) | `challenger_deep` | Challenger Deep | 🌙 Dark | `themes.ThemeChallengerDeep` |
| ![Chester](themes/chester.svg) | `chester` | Chester | 🌙 Dark | `themes.ThemeChester` |
| ![Ciapre](themes/ciapre.svg) | `ciapre` | Ciapre | 🌙 Dark | `themes.ThemeCiapre` |
| ![Citruszest](themes/citruszest.svg) | `citruszest` | Citruszest | 🌙 Dark | `themes.ThemeCitruszest` |
| ![CLRS](themes/clrs.svg) | `clrs` | CLRS | ☀️ Light | `themes.ThemeClrs` |
| ![Cobalt2](themes/cobalt2.svg) | `cobalt2` | Cobalt2 | 🌙 Dark | `themes.ThemeCobalt2` |
| ![Cobalt Neon](themes/cobalt_neon.svg) | `cobalt_neon` | Cobalt Neon | 🌙 Dark | `themes.ThemeCobaltNeon` |
| ![Cobalt Next](themes/cobalt_next.svg) | `cobalt_next` | Cobalt Next | 🌙 Dark | `themes.ThemeCobaltNext` |
| ![Cobalt Next Dark](themes/cobalt_next_dark.svg) | `cobalt_next_dark` | Cobalt Next Dark | 🌙 Dark | `themes.ThemeCobaltNextDark` |
| ![Cobalt Next Minimal](themes/cobalt_next_minimal.svg) | `cobalt_next_minimal` | Cobalt Next Minimal | 🌙 Dark | `themes.ThemeCobaltNextMinimal` |
| ![Coffee Theme](themes/coffee_theme.svg) | `coffee_theme` | Coffee Theme | ☀️ Light | `themes.ThemeCoffeeTheme` |
| ![Crayon Pony Fish](themes/crayon_pony_fish.svg) | `crayon_pony_fish` | Crayon Pony Fish | 🌙 Dark | `themes.ThemeCrayonPonyFish` |
| ![Cursor Dark](themes/cursor_dark.svg) | `cursor_dark` | Cursor Dark | 🌙 Dark | `themes.ThemeCursorDark` |
| ![Cutie Pro](themes/cutie_pro.svg) | `cutie_pro` | Cutie Pro | 🌙 Dark | `themes.ThemeCutiePro` |
| ![Cyberdyne](themes/cyberdyne.svg) | `cyberdyne` | Cyberdyne | 🌙 Dark | `themes.ThemeCyberdyne` |
| ![Cyberpunk](themes/cyberpunk.svg) | `cyberpunk` | Cyberpunk | 🌙 Dark | `themes.ThemeCyberpunk` |
| ![Cyberpunk Scarlet Protocol](themes/cyberpunk_scarlet_protocol.svg) | `cyberpunk_scarlet_protocol` | Cyberpunk Scarlet Protocol | 🌙 Dark | `themes.ThemeCyberpunkScarletProtocol` |
| ![Dark+](themes/dark.svg) | `dark` | Dark+ | 🌙 Dark | `themes.ThemeDark` |
| ![Dark Modern](themes/dark_modern.svg) | `dark_modern` | Dark Modern | 🌙 Dark | `themes.ThemeDarkModern` |
| ![Dark Pastel](themes/dark_pastel.svg) | `dark_pastel` | Dark Pastel | 🌙 Dark | `themes.ThemeDarkPastel` |
| ![Darkermatrix](themes/darkermatrix.svg) | `darkermatrix` | Darkermatrix | 🌙 Dark | `themes.ThemeDarkermatrix` |
| ![Darkmatrix](themes/darkmatrix.svg) | `darkmatrix` | Darkmatrix | 🌙 Dark | `themes.ThemeDarkmatrix` |
| ![Darkside](themes/darkside.svg) | `darkside` | Darkside | 🌙 Dark | `themes.ThemeDarkside` |
| ![Dawnfox](themes/dawnfox.svg) | `dawnfox` | Dawnfox | ☀️ Light | `themes.ThemeDawnfox` |
| ![Dayfox](themes/dayfox.svg) | `dayfox` | Dayfox | ☀️ Light | `themes.ThemeDayfox` |
| ![Deep](themes/deep.svg) | `deep` | Deep | 🌙 Dark | `themes.ThemeDeep` |
| ![Desert](themes/desert.svg) | `desert` | Desert | 🌙 Dark | `themes.ThemeDesert` |
| ![Detuned](themes/detuned.svg) | `detuned` | Detuned | 🌙 Dark | `themes.ThemeDetuned` |
| ![Dimidium](themes/dimidium.svg) | `dimidium` | Dimidium | 🌙 Dark | `themes.ThemeDimidium` |
| ![Dimmed Monokai](themes/dimmed_monokai.svg) | `dimmed_monokai` | Dimmed Monokai | 🌙 Dark | `themes.ThemeDimmedMonokai` |
| ![Django](themes/django.svg) | `django` | Django | 🌙 Dark | `themes.ThemeDjango` |
| ![Django Reborn Again](themes/django_reborn_again.svg) | `django_reborn_again` | Django Reborn Again | 🌙 Dark | `themes.ThemeDjangoRebornAgain` |
| ![Django Smooth](themes/django_smooth.svg) | `django_smooth` | Django Smooth | 🌙 Dark | `themes.ThemeDjangoSmooth` |
| ![Doom One](themes/doom_one.svg) | `doom_one` | Doom One | 🌙 Dark | `themes.ThemeDoomOne` |
| ![Doom Peacock](themes/doom_peacock.svg) | `doom_peacock` | Doom Peacock | 🌙 Dark | `themes.ThemeDoomPeacock` |
| ![Dot Gov](themes/dot_gov.svg) | `dot_gov` | Dot Gov | 🌙 Dark | `themes.ThemeDotGov` |
| ![Dracula+](themes/dracula.svg) | `dracula` | Dracula+ | 🌙 Dark | `themes.ThemeDracula` |
| ![Dracula](themes/dracula_2.svg) | `dracula_2` | Dracula | 🌙 Dark | `themes.ThemeDracula2` |
| ![Dracula Pro](themes/dracula_pro.svg) | `dracula_pro` | Dracula Pro | 🌙 Dark | `themes.ThemeDraculaPro` |
| ![Duckbones](themes/duckbones.svg) | `duckbones` | Duckbones | 🌙 Dark | `themes.ThemeDuckbones` |
| ![Duotone Dark](themes/duotone_dark.svg) | `duotone_dark` | Duotone Dark | 🌙 Dark | `themes.ThemeDuotoneDark` |
| ![Duskfox](themes/duskfox.svg) | `duskfox` | Duskfox | 🌙 Dark | `themes.ThemeDuskfox` |
| ![Earthsong](themes/earthsong.svg) | `earthsong` | Earthsong | 🌙 Dark | `themes.ThemeEarthsong` |
| ![Electron Highlighter](themes/electron_highlighter.svg) | `electron_highlighter` | Electron Highlighter | 🌙 Dark | `themes.ThemeElectronHighlighter` |
| ![Elegant](themes/elegant.svg) | `elegant` | Elegant | 🌙 Dark | `themes.ThemeElegant` |
| ![Elemental](themes/elemental.svg) | `elemental` | Elemental | 🌙 Dark | `themes.ThemeElemental` |
| ![Elementary](themes/elementary.svg) | `elementary` | Elementary | 🌙 Dark | `themes.ThemeElementary` |
| ![Embark](themes/embark.svg) | `embark` | Embark | 🌙 Dark | `themes.ThemeEmbark` |
| ![Embers Dark](themes/embers_dark.svg) | `embers_dark` | Embers Dark | 🌙 Dark | `themes.ThemeEmbersDark` |
| ![ENCOM](themes/encom.svg) | `encom` | ENCOM | 🌙 Dark | `themes.ThemeEncom` |
| ![Espresso](themes/espresso.svg) | `espresso` | Espresso | 🌙 Dark | `themes.ThemeEspresso` |
| ![Espresso Libre](themes/espresso_libre.svg) | `espresso_libre` | Espresso Libre | 🌙 Dark | `themes.ThemeEspressoLibre` |
| ![Everblush](themes/everblush.svg) | `everblush` | Everblush | 🌙 Dark | `themes.ThemeEverblush` |
| ![Everforest Dark Hard](themes/everforest_dark_hard.svg) | `everforest_dark_hard` | Everforest Dark Hard | 🌙 Dark | `themes.ThemeEverforestDarkHard` |
| ![Everforest Light Med](themes/everforest_light_med.svg) | `everforest_light_med` | Everforest Light Med | ☀️ Light | `themes.ThemeEverforestLightMed` |
| ![Fahrenheit](themes/fahrenheit.svg) | `fahrenheit` | Fahrenheit | 🌙 Dark | `themes.ThemeFahrenheit` |
| ![Fairyfloss](themes/fairyfloss.svg) | `fairyfloss` | Fairyfloss | 🌙 Dark | `themes.ThemeFairyfloss` |
| ![Farmhouse Dark](themes/farmhouse_dark.svg) | `farmhouse_dark` | Farmhouse Dark | 🌙 Dark | `themes.ThemeFarmhouseDark` |
| ![Farmhouse Light](themes/farmhouse_light.svg) | `farmhouse_light` | Farmhouse Light | ☀️ Light | `themes.ThemeFarmhouseLight` |
| ![Fideloper](themes/fideloper.svg) | `fideloper` | Fideloper | 🌙 Dark | `themes.ThemeFideloper` |
| ![Firefly Traditional](themes/firefly_traditional.svg) | `firefly_traditional` | Firefly Traditional | 🌙 Dark | `themes.ThemeFireflyTraditional` |
| ![Firefox Dev](themes/firefox_dev.svg) | `firefox_dev` | Firefox Dev | 🌙 Dark | `themes.ThemeFirefoxDev` |
| ![Firewatch](themes/firewatch.svg) | `firewatch` | Firewatch | 🌙 Dark | `themes.ThemeFirewatch` |
| ![Fish Tank](themes/fish_tank.svg) | `fish_tank` | Fish Tank | 🌙 Dark | `themes.ThemeFishTank` |
| ![Flat](themes/flat.svg) | `flat` | Flat | 🌙 Dark | `themes.ThemeFlat` |
| ![Flatland](themes/flatland.svg) | `flatland` | Flatland | 🌙 Dark | `themes.ThemeFlatland` |
| ![Flexoki Dark](themes/flexoki_dark.svg) | `flexoki_dark` | Flexoki Dark | 🌙 Dark | `themes.ThemeFlexokiDark` |
| ![Flexoki Light](themes/flexoki_light.svg) | `flexoki_light` | Flexoki Light | ☀️ Light | `themes.ThemeFlexokiLight` |
| ![Floraverse](themes/floraverse.svg) | `floraverse` | Floraverse | 🌙 Dark | `themes.ThemeFloraverse` |
| ![Forest Blue](themes/forest_blue.svg) | `forest_blue` | Forest Blue | 🌙 Dark | `themes.ThemeForestBlue` |
| ![Framer](themes/framer.svg) | `framer` | Framer | 🌙 Dark | `themes.ThemeFramer` |
| ![Front End Delight](themes/front_end_delight.svg) | `front_end_delight` | Front End Delight | 🌙 Dark | `themes.ThemeFrontEndDelight` |
| ![Fun Forrest](themes/fun_forrest.svg) | `fun_forrest` | Fun Forrest | 🌙 Dark | `themes.ThemeFunForrest` |
| ![Galaxy](themes/galaxy.svg) | `galaxy` | Galaxy | 🌙 Dark | `themes.ThemeGalaxy` |
| ![Galizur](themes/galizur.svg) | `galizur` | Galizur | 🌙 Dark | `themes.ThemeGalizur` |
| ![Ghostty Default Style Dark](themes/ghostty_default_style_dark.svg) | `ghostty_default_style_dark` | Ghostty Default Style Dark | 🌙 Dark | `themes.ThemeGhosttyDefaultStyleDark` |
| ![GitHub](themes/github.svg) | `github` | GitHub | ☀️ Light | `themes.ThemeGithub` |
| ![GitHub Dark](themes/github_dark.svg) | `github_dark` | GitHub Dark | 🌙 Dark | `themes.ThemeGithubDark` |
| ![GitHub Dark Colorblind](themes/github_dark_colorblind.svg) | `github_dark_colorblind` | GitHub Dark Colorblind | 🌙 Dark | `themes.ThemeGithubDarkColorblind` |
| ![GitHub Dark Default](themes/github_dark_default.svg) | `github_dark_default` | GitHub Dark Default | 🌙 Dark | `themes.ThemeGithubDarkDefault` |
| ![GitHub Dark Dimmed](themes/github_dark_dimmed.svg) | `github_dark_dimmed` | GitHub Dark Dimmed | 🌙 Dark | `themes.ThemeGithubDarkDimmed` |
| ![GitHub Dark High Contrast](themes/github_dark_high_contrast.svg) | `github_dark_high_contrast` | GitHub Dark High Contrast | 🌙 Dark | `themes.ThemeGithubDarkHighContrast` |
| ![GitHub Light Colorblind](themes/github_light_colorblind.svg) | `github_light_colorblind` | GitHub Light Colorblind | ☀️ Light | `themes.ThemeGithubLightColorblind` |
| ![GitHub Light Default](themes/github_light_default.svg) | `github_light_default` | GitHub Light Default | ☀️ Light | `themes.ThemeGithubLightDefault` |
| ![GitHub Light High Contrast](themes/github_light_high_contrast.svg) | `github_light_high_contrast` | GitHub Light High Contrast | ☀️ Light | `themes.ThemeGithubLightHighContrast` |
| ![GitLab Dark](themes/gitlab_dark.svg) | `gitlab_dark` | GitLab Dark | 🌙 Dark | `themes.ThemeGitlabDark` |
| ![GitLab Dark Grey](themes/gitlab_dark_grey.svg) | `gitlab_dark_grey` | GitLab Dark Grey | 🌙 Dark | `themes.ThemeGitlabDarkGrey` |
| ![GitLab Light](themes/gitlab_light.svg) | `gitlab_light` | GitLab Light | ☀️ Light | `themes.ThemeGitlabLight` |
| ![Glacier](themes/glacier.svg) | `glacier` | Glacier | 🌙 Dark | `themes.ThemeGlacier` |
| ![Grape](themes/grape.svg) | `grape` | Grape | 🌙 Dark | `themes.ThemeGrape` |
| ![Grass](themes/grass.svg) | `grass` | Grass | 🌙 Dark | `themes.ThemeGrass` |
| ![Grey Green](themes/grey_green.svg) | `grey_green` | Grey Green | 🌙 Dark | `themes.ThemeGreyGreen` |
| ![Gruber Darker](themes/gruber_darker.svg) | `gruber_darker` | Gruber Darker | 🌙 Dark | `themes.ThemeGruberDarker` |
| ![Gruvbox Dark](themes/gruvbox_dark.svg) | `gruvbox_dark` | Gruvbox Dark | 🌙 Dark | `themes.ThemeGruvboxDark` |
| ![Gruvbox Dark Hard](themes/gruvbox_dark_hard.svg) | `gruvbox_dark_hard` | Gruvbox Dark Hard | 🌙 Dark | `themes.ThemeGruvboxDarkHard` |
| ![Gruvbox Light](themes/gruvbox_light.svg) | `gruvbox_light` | Gruvbox Light | ☀️ Light | `themes.ThemeGruvboxLight` |
| ![Gruvbox Light Hard](themes/gruvbox_light_hard.svg) | `gruvbox_light_hard` | Gruvbox Light Hard | ☀️ Light | `themes.ThemeGruvboxLightHard` |
| ![Gruvbox Material](themes/gruvbox_material.svg) | `gruvbox_material` | Gruvbox Material | 🌙 Dark | `themes.ThemeGruvboxMaterial` |
| ![Gruvbox Material Dark](themes/gruvbox_material_dark.svg) | `gruvbox_material_dark` | Gruvbox Material Dark | 🌙 Dark | `themes.ThemeGruvboxMaterialDark` |
| ![Gruvbox Material Light](themes/gruvbox_material_light.svg) | `gruvbox_material_light` | Gruvbox Material Light | ☀️ Light | `themes.ThemeGruvboxMaterialLight` |
| ![Guezwhoz](themes/guezwhoz.svg) | `guezwhoz` | Guezwhoz | 🌙 Dark | `themes.ThemeGuezwhoz` |
| ![Hacktober](themes/hacktober.svg) | `hacktober` | Hacktober | 🌙 Dark | `themes.ThemeHacktober` |
| ![Hardcore](themes/hardcore.svg) | `hardcore` | Hardcore | 🌙 Dark | `themes.ThemeHardcore` |
| ![Harper](themes/harper.svg) | `harper` | Harper | 🌙 Dark | `themes.ThemeHarper` |
| ![Havn Daggry](themes/havn_daggry.svg) | `havn_daggry` | Havn Daggry | ☀️ Light | `themes.ThemeHavnDaggry` |
| ![Havn Skumring](themes/havn_skumring.svg) | `havn_skumring` | Havn Skumring | 🌙 Dark | `themes.ThemeHavnSkumring` |
| ![HaX0R Blue](themes/hax0r_blue.svg) | `hax0r_blue` | HaX0R Blue | 🌙 Dark | `themes.ThemeHax0rBlue` |
| ![HaX0R Gr33N](themes/hax0r_gr33n.svg) | `hax0r_gr33n` | HaX0R Gr33N | 🌙 Dark | `themes.ThemeHax0rGr33n` |
| ![HaX0R R3D](themes/hax0r_r3d.svg) | `hax0r_r3d` | HaX0R R3D | 🌙 Dark | `themes.ThemeHax0rR3d` |
| ![Heeler](themes/heeler.svg) | `heeler` | Heeler | 🌙 Dark | `themes.ThemeHeeler` |
| ![Highway](themes/highway.svg) | `highway` | Highway | 🌙 Dark | `themes.ThemeHighway` |
| ![Hipster Green](themes/hipster_green.svg) | `hipster_green` | Hipster Green | 🌙 Dark | `themes.ThemeHipsterGreen` |
| ![Hivacruz](themes/hivacruz.svg) | `hivacruz` | Hivacruz | 🌙 Dark | `themes.ThemeHivacruz` |
| ![Homebrew](themes/homebrew.svg) | `homebrew` | Homebrew | 🌙 Dark | `themes.ThemeHomebrew` |
| ![Hopscotch](themes/hopscotch.svg) | `hopscotch` | Hopscotch | 🌙 Dark | `themes.ThemeHopscotch` |
| ![Hopscotch.256](themes/hopscotch_256.svg) | `hopscotch_256` | Hopscotch.256 | 🌙 Dark | `themes.ThemeHopscotch256` |
| ![Horizon](themes/horizon.svg) | `horizon` | Horizon | 🌙 Dark | `themes.ThemeHorizon` |
| ![Horizon Bright](themes/horizon_bright.svg) | `horizon_bright` | Horizon Bright | ☀️ Light | `themes.ThemeHorizonBright` |
| ![Hot Dog Stand](themes/hot_dog_stand.svg) | `hot_dog_stand` | Hot Dog Stand | 🌙 Dark | `themes.ThemeHotDogStand` |
| ![Hot Dog Stand (Mustard)](themes/hot_dog_stand_mustard.svg) | `hot_dog_stand_mustard` | Hot Dog Stand (Mustard) | ☀️ Light | `themes.ThemeHotDogStandMustard` |
| ![Hurtado](themes/hurtado.svg) | `hurtado` | Hurtado | 🌙 Dark | `themes.ThemeHurtado` |
| ![Hybrid](themes/hybrid.svg) | `hybrid` | Hybrid | 🌙 Dark | `themes.ThemeHybrid` |
| ![IBM 5153 CGA](themes/ibm_5153_cga.svg) | `ibm_5153_cga` | IBM 5153 CGA | 🌙 Dark | `themes.ThemeIbm5153Cga` |
| ![IBM 5153 CGA (Black)](themes/ibm_5153_cga_black.svg) | `ibm_5153_cga_black` | IBM 5153 CGA (Black) | 🌙 Dark | `themes.ThemeIbm5153CgaBlack` |
| ![IC Green PPL](themes/ic_green_ppl.svg) | `ic_green_ppl` | IC Green PPL | 🌙 Dark | `themes.ThemeIcGreenPpl` |
| ![IC Orange PPL](themes/ic_orange_ppl.svg) | `ic_orange_ppl` | IC Orange PPL | 🌙 Dark | `themes.ThemeIcOrangePpl` |
| ![Iceberg Dark](themes/iceberg_dark.svg) | `iceberg_dark` | Iceberg Dark | 🌙 Dark | `themes.ThemeIcebergDark` |
| ![Iceberg Light](themes/iceberg_light.svg) | `iceberg_light` | Iceberg Light | ☀️ Light | `themes.ThemeIcebergLight` |
| ![Idea](themes/idea.svg) | `idea` | Idea | 🌙 Dark | `themes.ThemeIdea` |
| ![Idle Toes](themes/idle_toes.svg) | `idle_toes` | Idle Toes | 🌙 Dark | `themes.ThemeIdleToes` |
| ![IR Black](themes/ir_black.svg) | `ir_black` | IR Black | 🌙 Dark | `themes.ThemeIrBlack` |
| ![IRIX Console](themes/irix_console.svg) | `irix_console` | IRIX Console | 🌙 Dark | `themes.ThemeIrixConsole` |
| ![IRIX Terminal](themes/irix_terminal.svg) | `irix_terminal` | IRIX Terminal | 🌙 Dark | `themes.ThemeIrixTerminal` |
| ![iTerm2 Dark Background](themes/iterm2_dark_background.svg) | `iterm2_dark_background` | iTerm2 Dark Background | 🌙 Dark | `themes.ThemeIterm2DarkBackground` |
| ![iTerm2 Default](themes/iterm2_default.svg) | `iterm2_default` | iTerm2 Default | 🌙 Dark | `themes.ThemeIterm2Default` |
| ![iTerm2 Light Background](themes/iterm2_light_background.svg) | `iterm2_light_background` | iTerm2 Light Background | ☀️ Light | `themes.ThemeIterm2LightBackground` |
| ![iTerm2 Pastel Dark Background](themes/iterm2_pastel_dark_background.svg) | `iterm2_pastel_dark_background` | iTerm2 Pastel Dark Background | 🌙 Dark | `themes.ThemeIterm2PastelDarkBackground` |
| ![iTerm2 Smoooooth](themes/iterm2_smoooooth.svg) | `iterm2_smoooooth` | iTerm2 Smoooooth | 🌙 Dark | `themes.ThemeIterm2Smoooooth` |
| ![iTerm2 Solarized Dark](themes/iterm2_solarized_dark.svg) | `iterm2_solarized_dark` | iTerm2 Solarized Dark | 🌙 Dark | `themes.ThemeIterm2SolarizedDark` |
| ![iTerm2 Solarized Light](themes/iterm2_solarized_light.svg) | `iterm2_solarized_light` | iTerm2 Solarized Light | ☀️ Light | `themes.ThemeIterm2SolarizedLight` |
| ![iTerm2 Tango Dark](themes/iterm2_tango_dark.svg) | `iterm2_tango_dark` | iTerm2 Tango Dark | 🌙 Dark | `themes.ThemeIterm2TangoDark` |
| ![iTerm2 Tango Light](themes/iterm2_tango_light.svg) | `iterm2_tango_light` | iTerm2 Tango Light | ☀️ Light | `themes.ThemeIterm2TangoLight` |
| ![Jackie Brown](themes/jackie_brown.svg) | `jackie_brown` | Jackie Brown | 🌙 Dark | `themes.ThemeJackieBrown` |
| ![Japanesque](themes/japanesque.svg) | `japanesque` | Japanesque | 🌙 Dark | `themes.ThemeJapanesque` |
| ![Jellybeans](themes/jellybeans.svg) | `jellybeans` | Jellybeans | 🌙 Dark | `themes.ThemeJellybeans` |
| ![JetBrains Darcula](themes/jetbrains_darcula.svg) | `jetbrains_darcula` | JetBrains Darcula | 🌙 Dark | `themes.ThemeJetbrainsDarcula` |
| ![Jubi](themes/jubi.svg) | `jubi` | Jubi | 🌙 Dark | `themes.ThemeJubi` |
| ![Kanagawa Dragon](themes/kanagawa_dragon.svg) | `kanagawa_dragon` | Kanagawa Dragon | 🌙 Dark | `themes.ThemeKanagawaDragon` |
| ![Kanagawa Wave](themes/kanagawa_wave.svg) | `kanagawa_wave` | Kanagawa Wave | 🌙 Dark | `themes.ThemeKanagawaWave` |
| ![Kanagawabones](themes/kanagawabones.svg) | `kanagawabones` | Kanagawabones | 🌙 Dark | `themes.ThemeKanagawabones` |
| ![Kibble](themes/kibble.svg) | `kibble` | Kibble | 🌙 Dark | `themes.ThemeKibble` |
| ![Kitty Default](themes/kitty_default.svg) | `kitty_default` | Kitty Default | 🌙 Dark | `themes.ThemeKittyDefault` |
| ![Kitty Low Contrast](themes/kitty_low_contrast.svg) | `kitty_low_contrast` | Kitty Low Contrast | 🌙 Dark | `themes.ThemeKittyLowContrast` |
| ![Kolorit](themes/kolorit.svg) | `kolorit` | Kolorit | 🌙 Dark | `themes.ThemeKolorit` |
| ![Konsolas](themes/konsolas.svg) | `konsolas` | Konsolas | 🌙 Dark | `themes.ThemeKonsolas` |
| ![Kurokula](themes/kurokula.svg) | `kurokula` | Kurokula | 🌙 Dark | `themes.ThemeKurokula` |
| ![Lab Fox](themes/lab_fox.svg) | `lab_fox` | Lab Fox | 🌙 Dark | `themes.ThemeLabFox` |
| ![Laser](themes/laser.svg) | `laser` | Laser | 🌙 Dark | `themes.ThemeLaser` |
| ![Later This Evening](themes/later_this_evening.svg) | `later_this_evening` | Later This Evening | 🌙 Dark | `themes.ThemeLaterThisEvening` |
| ![Lavandula](themes/lavandula.svg) | `lavandula` | Lavandula | 🌙 Dark | `themes.ThemeLavandula` |
| ![Light Owl](themes/light_owl.svg) | `light_owl` | Light Owl | ☀️ Light | `themes.ThemeLightOwl` |
| ![Liquid Carbon](themes/liquid_carbon.svg) | `liquid_carbon` | Liquid Carbon | 🌙 Dark | `themes.ThemeLiquidCarbon` |
| ![Liquid Carbon Transparent](themes/liquid_carbon_transparent.svg) | `liquid_carbon_transparent` | Liquid Carbon Transparent | 🌙 Dark | `themes.ThemeLiquidCarbonTransparent` |
| ![Lovelace](themes/lovelace.svg) | `lovelace` | Lovelace | 🌙 Dark | `themes.ThemeLovelace` |
| ![Man Page](themes/man_page.svg) | `man_page` | Man Page | ☀️ Light | `themes.ThemeManPage` |
| ![Mariana](themes/mariana.svg) | `mariana` | Mariana | 🌙 Dark | `themes.ThemeMariana` |
| ![Material](themes/material.svg) | `material` | Material | ☀️ Light | `themes.ThemeMaterial` |
| ![Material Dark](themes/material_dark.svg) | `material_dark` | Material Dark | 🌙 Dark | `themes.ThemeMaterialDark` |
| ![Material Darker](themes/material_darker.svg) | `material_darker` | Material Darker | 🌙 Dark | `themes.ThemeMaterialDarker` |
| ![Material Design Colors](themes/material_design_colors.svg) | `material_design_colors` | Material Design Colors | 🌙 Dark | `themes.ThemeMaterialDesignColors` |
| ![Material Ocean](themes/material_ocean.svg) | `material_ocean` | Material Ocean | 🌙 Dark | `themes.ThemeMaterialOcean` |
| ![Mathias](themes/mathias.svg) | `mathias` | Mathias | 🌙 Dark | `themes.ThemeMathias` |
| ![Matrix](themes/matrix.svg) | `matrix` | Matrix | 🌙 Dark | `themes.ThemeMatrix` |
| ![Matte Black](themes/matte_black.svg) | `matte_black` | Matte Black | 🌙 Dark | `themes.ThemeMatteBlack` |
| ![Medallion](themes/medallion.svg) | `medallion` | Medallion | 🌙 Dark | `themes.ThemeMedallion` |
| ![Melange Dark](themes/melange_dark.svg) | `melange_dark` | Melange Dark | 🌙 Dark | `themes.ThemeMelangeDark` |
| ![Melange Light](themes/melange_light.svg) | `melange_light` | Melange Light | ☀️ Light | `themes.ThemeMelangeLight` |
| ![Mellifluous](themes/mellifluous.svg) | `mellifluous` | Mellifluous | 🌙 Dark | `themes.ThemeMellifluous` |
| ![Mellow](themes/mellow.svg) | `mellow` | Mellow | 🌙 Dark | `themes.ThemeMellow` |
| ![Miasma](themes/miasma.svg) | `miasma` | Miasma | 🌙 Dark | `themes.ThemeMiasma` |
| ![Midnight In Mojave](themes/midnight_in_mojave.svg) | `midnight_in_mojave` | Midnight In Mojave | 🌙 Dark | `themes.ThemeMidnightInMojave` |
| ![Mirage](themes/mirage.svg) | `mirage` | Mirage | 🌙 Dark | `themes.ThemeMirage` |
| ![Misterioso](themes/misterioso.svg) | `misterioso` | Misterioso | 🌙 Dark | `themes.ThemeMisterioso` |
| ![Molokai](themes/molokai.svg) | `molokai` | Molokai | 🌙 Dark | `themes.ThemeMolokai` |
| ![Mona Lisa](themes/mona_lisa.svg) | `mona_lisa` | Mona Lisa | 🌙 Dark | `themes.ThemeMonaLisa` |
| ![Monokai Classic](themes/monokai_classic.svg) | `monokai_classic` | Monokai Classic | 🌙 Dark | `themes.ThemeMonokaiClassic` |
| ![Monokai Pro](themes/monokai_pro.svg) | `monokai_pro` | Monokai Pro | 🌙 Dark | `themes.ThemeMonokaiPro` |
| ![Monokai Pro Light](themes/monokai_pro_light.svg) | `monokai_pro_light` | Monokai Pro Light | ☀️ Light | `themes.ThemeMonokaiProLight` |
| ![Monokai Pro Light Sun](themes/monokai_pro_light_sun.svg) | `monokai_pro_light_sun` | Monokai Pro Light Sun | ☀️ Light | `themes.ThemeMonokaiProLightSun` |
| ![Monokai Pro Machine](themes/monokai_pro_machine.svg) | `monokai_pro_machine` | Monokai Pro Machine | 🌙 Dark | `themes.ThemeMonokaiProMachine` |
| ![Monokai Pro Octagon](themes/monokai_pro_octagon.svg) | `monokai_pro_octagon` | Monokai Pro Octagon | 🌙 Dark | `themes.ThemeMonokaiProOctagon` |
| ![Monokai Pro Ristretto](themes/monokai_pro_ristretto.svg) | `monokai_pro_ristretto` | Monokai Pro Ristretto | 🌙 Dark | `themes.ThemeMonokaiProRistretto` |
| ![Monokai Pro Spectrum](themes/monokai_pro_spectrum.svg) | `monokai_pro_spectrum` | Monokai Pro Spectrum | 🌙 Dark | `themes.ThemeMonokaiProSpectrum` |
| ![Monokai Remastered](themes/monokai_remastered.svg) | `monokai_remastered` | Monokai Remastered | 🌙 Dark | `themes.ThemeMonokaiRemastered` |
| ![Monokai Soda](themes/monokai_soda.svg) | `monokai_soda` | Monokai Soda | 🌙 Dark | `themes.ThemeMonokaiSoda` |
| ![Monokai Vivid](themes/monokai_vivid.svg) | `monokai_vivid` | Monokai Vivid | 🌙 Dark | `themes.ThemeMonokaiVivid` |
| ![Moonfly](themes/moonfly.svg) | `moonfly` | Moonfly | 🌙 Dark | `themes.ThemeMoonfly` |
| ![N0Tch2K](themes/n0tch2k.svg) | `n0tch2k` | N0Tch2K | 🌙 Dark | `themes.ThemeN0tch2k` |
| ![Neobones Dark](themes/neobones_dark.svg) | `neobones_dark` | Neobones Dark | 🌙 Dark | `themes.ThemeNeobonesDark` |
| ![Neobones Light](themes/neobones_light.svg) | `neobones_light` | Neobones Light | ☀️ Light | `themes.ThemeNeobonesLight` |
| ![Neon](themes/neon.svg) | `neon` | Neon | 🌙 Dark | `themes.ThemeNeon` |
| ![Neopolitan](themes/neopolitan.svg) | `neopolitan` | Neopolitan | 🌙 Dark | `themes.ThemeNeopolitan` |
| ![Neutron](themes/neutron.svg) | `neutron` | Neutron | 🌙 Dark | `themes.ThemeNeutron` |
| ![Night Lion V1](themes/night_lion_v1.svg) | `night_lion_v1` | Night Lion V1 | 🌙 Dark | `themes.ThemeNightLionV1` |
| ![Night Lion V2](themes/night_lion_v2.svg) | `night_lion_v2` | Night Lion V2 | 🌙 Dark | `themes.ThemeNightLionV2` |
| ![Night Owl](themes/night_owl.svg) | `night_owl` | Night Owl | 🌙 Dark | `themes.ThemeNightOwl` |
| ![Night Owlish Light](themes/night_owlish_light.svg) | `night_owlish_light` | Night Owlish Light | ☀️ Light | `themes.ThemeNightOwlishLight` |
| ![Nightfox](themes/nightfox.svg) | `nightfox` | Nightfox | 🌙 Dark | `themes.ThemeNightfox` |
| ![Niji](themes/niji.svg) | `niji` | Niji | 🌙 Dark | `themes.ThemeNiji` |
| ![No Clown Fiesta](themes/no_clown_fiesta.svg) | `no_clown_fiesta` | No Clown Fiesta | 🌙 Dark | `themes.ThemeNoClownFiesta` |
| ![No Clown Fiesta Light](themes/no_clown_fiesta_light.svg) | `no_clown_fiesta_light` | No Clown Fiesta Light | 🌙 Dark | `themes.ThemeNoClownFiestaLight` |
| ![Nocturnal Winter](themes/nocturnal_winter.svg) | `nocturnal_winter` | Nocturnal Winter | 🌙 Dark | `themes.ThemeNocturnalWinter` |
| ![Nord](themes/nord.svg) | `nord` | Nord | 🌙 Dark | `themes.ThemeNord` |
| ![Nord Light](themes/nord_light.svg) | `nord_light` | Nord Light | ☀️ Light | `themes.ThemeNordLight` |
| ![Nord Wave](themes/nord_wave.svg) | `nord_wave` | Nord Wave | 🌙 Dark | `themes.ThemeNordWave` |
| ![Nordfox](themes/nordfox.svg) | `nordfox` | Nordfox | 🌙 Dark | `themes.ThemeNordfox` |
| ![Novel](themes/novel.svg) | `novel` | Novel | 🌙 Dark | `themes.ThemeNovel` |
| ![novmbr](themes/novmbr.svg) | `novmbr` | novmbr | 🌙 Dark | `themes.ThemeNovmbr` |
| ![Nvim Dark](themes/nvim_dark.svg) | `nvim_dark` | Nvim Dark | 🌙 Dark | `themes.ThemeNvimDark` |
| ![Nvim Light](themes/nvim_light.svg) | `nvim_light` | Nvim Light | ☀️ Light | `themes.ThemeNvimLight` |
| ![Obsidian](themes/obsidian.svg) | `obsidian` | Obsidian | 🌙 Dark | `themes.ThemeObsidian` |
| ![Ocean](themes/ocean.svg) | `ocean` | Ocean | 🌙 Dark | `themes.ThemeOcean` |
| ![Oceanic Material](themes/oceanic_material.svg) | `oceanic_material` | Oceanic Material | 🌙 Dark | `themes.ThemeOceanicMaterial` |
| ![Oceanic Next](themes/oceanic_next.svg) | `oceanic_next` | Oceanic Next | 🌙 Dark | `themes.ThemeOceanicNext` |
| ![Ollie](themes/ollie.svg) | `ollie` | Ollie | 🌙 Dark | `themes.ThemeOllie` |
| ![One Dark Two](themes/one_dark_two.svg) | `one_dark_two` | One Dark Two | 🌙 Dark | `themes.ThemeOneDarkTwo` |
| ![One Double Dark](themes/one_double_dark.svg) | `one_double_dark` | One Double Dark | 🌙 Dark | `themes.ThemeOneDoubleDark` |
| ![One Double Light](themes/one_double_light.svg) | `one_double_light` | One Double Light | ☀️ Light | `themes.ThemeOneDoubleLight` |
| ![One Half Dark](themes/one_half_dark.svg) | `one_half_dark` | One Half Dark | 🌙 Dark | `themes.ThemeOneHalfDark` |
| ![One Half Light](themes/one_half_light.svg) | `one_half_light` | One Half Light | ☀️ Light | `themes.ThemeOneHalfLight` |
| ![Operator Mono Dark](themes/operator_mono_dark.svg) | `operator_mono_dark` | Operator Mono Dark | 🌙 Dark | `themes.ThemeOperatorMonoDark` |
| ![Overnight Slumber](themes/overnight_slumber.svg) | `overnight_slumber` | Overnight Slumber | 🌙 Dark | `themes.ThemeOvernightSlumber` |
| ![owl](themes/owl.svg) | `owl` | owl | 🌙 Dark | `themes.ThemeOwl` |
| ![Oxocarbon](themes/oxocarbon.svg) | `oxocarbon` | Oxocarbon | 🌙 Dark | `themes.ThemeOxocarbon` |
| ![Pale Night Hc](themes/pale_night_hc.svg) | `pale_night_hc` | Pale Night Hc | 🌙 Dark | `themes.ThemePaleNightHc` |
| ![Pandora](themes/pandora.svg) | `pandora` | Pandora | 🌙 Dark | `themes.ThemePandora` |
| ![Paraiso Dark](themes/paraiso_dark.svg) | `paraiso_dark` | Paraiso Dark | 🌙 Dark | `themes.ThemeParaisoDark` |
| ![Paul Millr](themes/paul_millr.svg) | `paul_millr` | Paul Millr | 🌙 Dark | `themes.ThemePaulMillr` |
| ![Pencil Dark](themes/pencil_dark.svg) | `pencil_dark` | Pencil Dark | 🌙 Dark | `themes.ThemePencilDark` |
| ![Pencil Light](themes/pencil_light.svg) | `pencil_light` | Pencil Light | ☀️ Light | `themes.ThemePencilLight` |
| ![Peppermint](themes/peppermint.svg) | `peppermint` | Peppermint | 🌙 Dark | `themes.ThemePeppermint` |
| ![Phala Green Dark](themes/phala_green_dark.svg) | `phala_green_dark` | Phala Green Dark | 🌙 Dark | `themes.ThemePhalaGreenDark` |
| ![Piatto Light](themes/piatto_light.svg) | `piatto_light` | Piatto Light | ☀️ Light | `themes.ThemePiattoLight` |
| ![Pnevma](themes/pnevma.svg) | `pnevma` | Pnevma | 🌙 Dark | `themes.ThemePnevma` |
| ![Poimandres](themes/poimandres.svg) | `poimandres` | Poimandres | 🌙 Dark | `themes.ThemePoimandres` |
| ![Poimandres Darker](themes/poimandres_darker.svg) | `poimandres_darker` | Poimandres Darker | 🌙 Dark | `themes.ThemePoimandresDarker` |
| ![Poimandres Storm](themes/poimandres_storm.svg) | `poimandres_storm` | Poimandres Storm | 🌙 Dark | `themes.ThemePoimandresStorm` |
| ![Poimandres White](themes/poimandres_white.svg) | `poimandres_white` | Poimandres White | ☀️ Light | `themes.ThemePoimandresWhite` |
| ![Popping And Locking](themes/popping_and_locking.svg) | `popping_and_locking` | Popping And Locking | 🌙 Dark | `themes.ThemePoppingAndLocking` |
| ![Powershell](themes/powershell.svg) | `powershell` | Powershell | 🌙 Dark | `themes.ThemePowershell` |
| ![Primary](themes/primary.svg) | `primary` | Primary | ☀️ Light | `themes.ThemePrimary` |
| ![Pro](themes/pro.svg) | `pro` | Pro | 🌙 Dark | `themes.ThemePro` |
| ![Pro Light](themes/pro_light.svg) | `pro_light` | Pro Light | ☀️ Light | `themes.ThemeProLight` |
| ![Purple Rain](themes/purple_rain.svg) | `purple_rain` | Purple Rain | 🌙 Dark | `themes.ThemePurpleRain` |
| ![Purplepeter](themes/purplepeter.svg) | `purplepeter` | Purplepeter | 🌙 Dark | `themes.ThemePurplepeter` |
| ![Rapture](themes/rapture.svg) | `rapture` | Rapture | 🌙 Dark | `themes.ThemeRapture` |
| ![Raycast Dark](themes/raycast_dark.svg) | `raycast_dark` | Raycast Dark | 🌙 Dark | `themes.ThemeRaycastDark` |
| ![Raycast Light](themes/raycast_light.svg) | `raycast_light` | Raycast Light | ☀️ Light | `themes.ThemeRaycastLight` |
| ![Rebecca](themes/rebecca.svg) | `rebecca` | Rebecca | 🌙 Dark | `themes.ThemeRebecca` |
| ![Red Alert](themes/red_alert.svg) | `red_alert` | Red Alert | 🌙 Dark | `themes.ThemeRedAlert` |
| ![Red Planet](themes/red_planet.svg) | `red_planet` | Red Planet | 🌙 Dark | `themes.ThemeRedPlanet` |
| ![Red Sands](themes/red_sands.svg) | `red_sands` | Red Sands | 🌙 Dark | `themes.ThemeRedSands` |
| ![Relaxed](themes/relaxed.svg) | `relaxed` | Relaxed | 🌙 Dark | `themes.ThemeRelaxed` |
| ![Retro](themes/retro.svg) | `retro` | Retro | 🌙 Dark | `themes.ThemeRetro` |
| ![Retro Legends](themes/retro_legends.svg) | `retro_legends` | Retro Legends | 🌙 Dark | `themes.ThemeRetroLegends` |
| ![Rippedcasts](themes/rippedcasts.svg) | `rippedcasts` | Rippedcasts | 🌙 Dark | `themes.ThemeRippedcasts` |
| ![Rose Pine](themes/rose_pine.svg) | `rose_pine` | Rose Pine | 🌙 Dark | `themes.ThemeRosePine` |
| ![Rose Pine Dawn](themes/rose_pine_dawn.svg) | `rose_pine_dawn` | Rose Pine Dawn | ☀️ Light | `themes.ThemeRosePineDawn` |
| ![Rose Pine Moon](themes/rose_pine_moon.svg) | `rose_pine_moon` | Rose Pine Moon | 🌙 Dark | `themes.ThemeRosePineMoon` |
| ![Rouge 2](themes/rouge_2.svg) | `rouge_2` | Rouge 2 | 🌙 Dark | `themes.ThemeRouge2` |
| ![Royal](themes/royal.svg) | `royal` | Royal | 🌙 Dark | `themes.ThemeRoyal` |
| ![Ryuuko](themes/ryuuko.svg) | `ryuuko` | Ryuuko | 🌙 Dark | `themes.ThemeRyuuko` |
| ![Sakura](themes/sakura.svg) | `sakura` | Sakura | 🌙 Dark | `themes.ThemeSakura` |
| ![Scarlet Protocol](themes/scarlet_protocol.svg) | `scarlet_protocol` | Scarlet Protocol | 🌙 Dark | `themes.ThemeScarletProtocol` |
| ![Sea Shells](themes/sea_shells.svg) | `sea_shells` | Sea Shells | 🌙 Dark | `themes.ThemeSeaShells` |
| ![Seafoam Pastel](themes/seafoam_pastel.svg) | `seafoam_pastel` | Seafoam Pastel | 🌙 Dark | `themes.ThemeSeafoamPastel` |
| ![Selenized Black](themes/selenized_black.svg) | `selenized_black` | Selenized Black | 🌙 Dark | `themes.ThemeSelenizedBlack` |
| ![Selenized Dark](themes/selenized_dark.svg) | `selenized_dark` | Selenized Dark | 🌙 Dark | `themes.ThemeSelenizedDark` |
| ![Selenized Light](themes/selenized_light.svg) | `selenized_light` | Selenized Light | ☀️ Light | `themes.ThemeSelenizedLight` |
| ![Seoulbones Dark](themes/seoulbones_dark.svg) | `seoulbones_dark` | Seoulbones Dark | 🌙 Dark | `themes.ThemeSeoulbonesDark` |
| ![Seoulbones Light](themes/seoulbones_light.svg) | `seoulbones_light` | Seoulbones Light | ☀️ Light | `themes.ThemeSeoulbonesLight` |
| ![Seti](themes/seti.svg) | `seti` | Seti | 🌙 Dark | `themes.ThemeSeti` |
| ![Shades Of Purple](themes/shades_of_purple.svg) | `shades_of_purple` | Shades Of Purple | 🌙 Dark | `themes.ThemeShadesOfPurple` |
| ![Shaman](themes/shaman.svg) | `shaman` | Shaman | 🌙 Dark | `themes.ThemeShaman` |
| ![Slate](themes/slate.svg) | `slate` | Slate | 🌙 Dark | `themes.ThemeSlate` |
| ![Sleepy Hollow](themes/sleepy_hollow.svg) | `sleepy_hollow` | Sleepy Hollow | 🌙 Dark | `themes.ThemeSleepyHollow` |
| ![Smyck](themes/smyck.svg) | `smyck` | Smyck | 🌙 Dark | `themes.ThemeSmyck` |
| ![Snazzy](themes/snazzy.svg) | `snazzy` | Snazzy | 🌙 Dark | `themes.ThemeSnazzy` |
| ![Snazzy Soft](themes/snazzy_soft.svg) | `snazzy_soft` | Snazzy Soft | 🌙 Dark | `themes.ThemeSnazzySoft` |
| ![Soft Server](themes/soft_server.svg) | `soft_server` | Soft Server | 🌙 Dark | `themes.ThemeSoftServer` |
| ![Solarized Darcula](themes/solarized_darcula.svg) | `solarized_darcula` | Solarized Darcula | 🌙 Dark | `themes.ThemeSolarizedDarcula` |
| ![Solarized Dark Higher Contrast](themes/solarized_dark_higher_contrast.svg) | `solarized_dark_higher_contrast` | Solarized Dark Higher Contrast | 🌙 Dark | `themes.ThemeSolarizedDarkHigherContrast` |
| ![Solarized Dark Patched](themes/solarized_dark_patched.svg) | `solarized_dark_patched` | Solarized Dark Patched | 🌙 Dark | `themes.ThemeSolarizedDarkPatched` |
| ![Solarized Osaka Night](themes/solarized_osaka_night.svg) | `solarized_osaka_night` | Solarized Osaka Night | 🌙 Dark | `themes.ThemeSolarizedOsakaNight` |
| ![Sonokai](themes/sonokai.svg) | `sonokai` | Sonokai | 🌙 Dark | `themes.ThemeSonokai` |
| ![Spacedust](themes/spacedust.svg) | `spacedust` | Spacedust | 🌙 Dark | `themes.ThemeSpacedust` |
| ![Spacegray](themes/spacegray.svg) | `spacegray` | Spacegray | 🌙 Dark | `themes.ThemeSpacegray` |
| ![Spacegray Bright](themes/spacegray_bright.svg) | `spacegray_bright` | Spacegray Bright | 🌙 Dark | `themes.ThemeSpacegrayBright` |
| ![Spacegray Eighties](themes/spacegray_eighties.svg) | `spacegray_eighties` | Spacegray Eighties | 🌙 Dark | `themes.ThemeSpacegrayEighties` |
| ![Spacegray Eighties Dull](themes/spacegray_eighties_dull.svg) | `spacegray_eighties_dull` | Spacegray Eighties Dull | 🌙 Dark | `themes.ThemeSpacegrayEightiesDull` |
| ![Spiderman](themes/spiderman.svg) | `spiderman` | Spiderman | 🌙 Dark | `themes.ThemeSpiderman` |
| ![Spring](themes/spring.svg) | `spring` | Spring | ☀️ Light | `themes.ThemeSpring` |
| ![Square](themes/square.svg) | `square` | Square | 🌙 Dark | `themes.ThemeSquare` |
| ![Squirrelsong Dark](themes/squirrelsong_dark.svg) | `squirrelsong_dark` | Squirrelsong Dark | 🌙 Dark | `themes.ThemeSquirrelsongDark` |
| ![Srcery](themes/srcery.svg) | `srcery` | Srcery | 🌙 Dark | `themes.ThemeSrcery` |
| ![Starlight](themes/starlight.svg) | `starlight` | Starlight | 🌙 Dark | `themes.ThemeStarlight` |
| ![Sublette](themes/sublette.svg) | `sublette` | Sublette | 🌙 Dark | `themes.ThemeSublette` |
| ![Subliminal](themes/subliminal.svg) | `subliminal` | Subliminal | 🌙 Dark | `themes.ThemeSubliminal` |
| ![Sugarplum](themes/sugarplum.svg) | `sugarplum` | Sugarplum | 🌙 Dark | `themes.ThemeSugarplum` |
| ![Sundried](themes/sundried.svg) | `sundried` | Sundried | 🌙 Dark | `themes.ThemeSundried` |
| ![Symfonic](themes/symfonic.svg) | `symfonic` | Symfonic | 🌙 Dark | `themes.ThemeSymfonic` |
| ![Synthwave](themes/synthwave.svg) | `synthwave` | Synthwave | 🌙 Dark | `themes.ThemeSynthwave` |
| ![Synthwave Alpha](themes/synthwave_alpha.svg) | `synthwave_alpha` | Synthwave Alpha | 🌙 Dark | `themes.ThemeSynthwaveAlpha` |
| ![Synthwave Everything](themes/synthwave_everything.svg) | `synthwave_everything` | Synthwave Everything | 🌙 Dark | `themes.ThemeSynthwaveEverything` |
| ![Tango Adapted](themes/tango_adapted.svg) | `tango_adapted` | Tango Adapted | ☀️ Light | `themes.ThemeTangoAdapted` |
| ![Tango Half Adapted](themes/tango_half_adapted.svg) | `tango_half_adapted` | Tango Half Adapted | ☀️ Light | `themes.ThemeTangoHalfAdapted` |
| ![Tearout](themes/tearout.svg) | `tearout` | Tearout | 🌙 Dark | `themes.ThemeTearout` |
| ![Teerb](themes/teerb.svg) | `teerb` | Teerb | 🌙 Dark | `themes.ThemeTeerb` |
| ![Terafox](themes/terafox.svg) | `terafox` | Terafox | 🌙 Dark | `themes.ThemeTerafox` |
| ![Terminal Basic](themes/terminal_basic.svg) | `terminal_basic` | Terminal Basic | ☀️ Light | `themes.ThemeTerminalBasic` |
| ![Terminal Basic Dark](themes/terminal_basic_dark.svg) | `terminal_basic_dark` | Terminal Basic Dark | 🌙 Dark | `themes.ThemeTerminalBasicDark` |
| ![Thayer Bright](themes/thayer_bright.svg) | `thayer_bright` | Thayer Bright | 🌙 Dark | `themes.ThemeThayerBright` |
| ![The Hulk](themes/the_hulk.svg) | `the_hulk` | The Hulk | 🌙 Dark | `themes.ThemeTheHulk` |
| ![Tinacious Design Dark](themes/tinacious_design_dark.svg) | `tinacious_design_dark` | Tinacious Design Dark | 🌙 Dark | `themes.ThemeTinaciousDesignDark` |
| ![Tinacious Design Light](themes/tinacious_design_light.svg) | `tinacious_design_light` | Tinacious Design Light | ☀️ Light | `themes.ThemeTinaciousDesignLight` |
| ![TokyoNight](themes/tokyonight.svg) | `tokyonight` | TokyoNight | 🌙 Dark | `themes.ThemeTokyonight` |
| ![TokyoNight Day](themes/tokyonight_day.svg) | `tokyonight_day` | TokyoNight Day | ☀️ Light | `themes.ThemeTokyonightDay` |
| ![TokyoNight Moon](themes/tokyonight_moon.svg) | `tokyonight_moon` | TokyoNight Moon | 🌙 Dark | `themes.ThemeTokyonightMoon` |
| ![TokyoNight Night](themes/tokyonight_night.svg) | `tokyonight_night` | TokyoNight Night | 🌙 Dark | `themes.ThemeTokyonightNight` |
| ![TokyoNight Storm](themes/tokyonight_storm.svg) | `tokyonight_storm` | TokyoNight Storm | 🌙 Dark | `themes.ThemeTokyonightStorm` |
| ![Tomorrow](themes/tomorrow.svg) | `tomorrow` | Tomorrow | ☀️ Light | `themes.ThemeTomorrow` |
| ![Tomorrow Night](themes/tomorrow_night.svg) | `tomorrow_night` | Tomorrow Night | 🌙 Dark | `themes.ThemeTomorrowNight` |
| ![Tomorrow Night Blue](themes/tomorrow_night_blue.svg) | `tomorrow_night_blue` | Tomorrow Night Blue | 🌙 Dark | `themes.ThemeTomorrowNightBlue` |
| ![Tomorrow Night Bright](themes/tomorrow_night_bright.svg) | `tomorrow_night_bright` | Tomorrow Night Bright | 🌙 Dark | `themes.ThemeTomorrowNightBright` |
| ![Tomorrow Night Burns](themes/tomorrow_night_burns.svg) | `tomorrow_night_burns` | Tomorrow Night Burns | 🌙 Dark | `themes.ThemeTomorrowNightBurns` |
| ![Tomorrow Night Eighties](themes/tomorrow_night_eighties.svg) | `tomorrow_night_eighties` | Tomorrow Night Eighties | 🌙 Dark | `themes.ThemeTomorrowNightEighties` |
| ![Toy Chest](themes/toy_chest.svg) | `toy_chest` | Toy Chest | 🌙 Dark | `themes.ThemeToyChest` |
| ![traffic](themes/traffic.svg) | `traffic` | traffic | 🌙 Dark | `themes.ThemeTraffic` |
| ![Treehouse](themes/treehouse.svg) | `treehouse` | Treehouse | 🌙 Dark | `themes.ThemeTreehouse` |
| ![Twilight](themes/twilight.svg) | `twilight` | Twilight | 🌙 Dark | `themes.ThemeTwilight` |
| ![Ubuntu](themes/ubuntu.svg) | `ubuntu` | Ubuntu | 🌙 Dark | `themes.ThemeUbuntu` |
| ![Ultra Dark](themes/ultra_dark.svg) | `ultra_dark` | Ultra Dark | 🌙 Dark | `themes.ThemeUltraDark` |
| ![Ultra Violent](themes/ultra_violent.svg) | `ultra_violent` | Ultra Violent | 🌙 Dark | `themes.ThemeUltraViolent` |
| ![Under The Sea](themes/under_the_sea.svg) | `under_the_sea` | Under The Sea | 🌙 Dark | `themes.ThemeUnderTheSea` |
| ![Unikitty](themes/unikitty.svg) | `unikitty` | Unikitty | 🌙 Dark | `themes.ThemeUnikitty` |
| ![urban](themes/urban.svg) | `urban` | urban | 🌙 Dark | `themes.ThemeUrban` |
| ![Urple](themes/urple.svg) | `urple` | Urple | 🌙 Dark | `themes.ThemeUrple` |
| ![Vague](themes/vague.svg) | `vague` | Vague | 🌙 Dark | `themes.ThemeVague` |
| ![Vaughn](themes/vaughn.svg) | `vaughn` | Vaughn | 🌙 Dark | `themes.ThemeVaughn` |
| ![Vercel](themes/vercel.svg) | `vercel` | Vercel | 🌙 Dark | `themes.ThemeVercel` |
| ![Vesper](themes/vesper.svg) | `vesper` | Vesper | 🌙 Dark | `themes.ThemeVesper` |
| ![Vibrant Ink](themes/vibrant_ink.svg) | `vibrant_ink` | Vibrant Ink | 🌙 Dark | `themes.ThemeVibrantInk` |
| ![Vimbones](themes/vimbones.svg) | `vimbones` | Vimbones | ☀️ Light | `themes.ThemeVimbones` |
| ![Violet Dark](themes/violet_dark.svg) | `violet_dark` | Violet Dark | 🌙 Dark | `themes.ThemeVioletDark` |
| ![Violet Light](themes/violet_light.svg) | `violet_light` | Violet Light | ☀️ Light | `themes.ThemeVioletLight` |
| ![Violite](themes/violite.svg) | `violite` | Violite | 🌙 Dark | `themes.ThemeViolite` |
| ![Warm Neon](themes/warm_neon.svg) | `warm_neon` | Warm Neon | 🌙 Dark | `themes.ThemeWarmNeon` |
| ![Wez](themes/wez.svg) | `wez` | Wez | 🌙 Dark | `themes.ThemeWez` |
| ![Whimsy](themes/whimsy.svg) | `whimsy` | Whimsy | 🌙 Dark | `themes.ThemeWhimsy` |
| ![Wild Cherry](themes/wild_cherry.svg) | `wild_cherry` | Wild Cherry | 🌙 Dark | `themes.ThemeWildCherry` |
| ![Wilmersdorf](themes/wilmersdorf.svg) | `wilmersdorf` | Wilmersdorf | 🌙 Dark | `themes.ThemeWilmersdorf` |
| ![Wombat](themes/wombat.svg) | `wombat` | Wombat | 🌙 Dark | `themes.ThemeWombat` |
| ![Wryan](themes/wryan.svg) | `wryan` | Wryan | 🌙 Dark | `themes.ThemeWryan` |
| ![Xcode Dark](themes/xcode_dark.svg) | `xcode_dark` | Xcode Dark | 🌙 Dark | `themes.ThemeXcodeDark` |
| ![Xcode Dark hc](themes/xcode_dark_hc.svg) | `xcode_dark_hc` | Xcode Dark hc | 🌙 Dark | `themes.ThemeXcodeDarkHc` |
| ![Xcode Light](themes/xcode_light.svg) | `xcode_light` | Xcode Light | ☀️ Light | `themes.ThemeXcodeLight` |
| ![Xcode Light hc](themes/xcode_light_hc.svg) | `xcode_light_hc` | Xcode Light hc | ☀️ Light | `themes.ThemeXcodeLightHc` |
| ![Xcode WWDC](themes/xcode_wwdc.svg) | `xcode_wwdc` | Xcode WWDC | 🌙 Dark | `themes.ThemeXcodeWwdc` |
| ![Zenbones](themes/zenbones.svg) | `zenbones` | Zenbones | ☀️ Light | `themes.ThemeZenbones` |
| ![Zenbones Dark](themes/zenbones_dark.svg) | `zenbones_dark` | Zenbones Dark | 🌙 Dark | `themes.ThemeZenbonesDark` |
| ![Zenbones Light](themes/zenbones_light.svg) | `zenbones_light` | Zenbones Light | ☀️ Light | `themes.ThemeZenbonesLight` |
| ![Zenburn](themes/zenburn.svg) | `zenburn` | Zenburn | 🌙 Dark | `themes.ThemeZenburn` |
| ![Zenburned](themes/zenburned.svg) | `zenburned` | Zenburned | 🌙 Dark | `themes.ThemeZenburned` |
| ![Zenwritten Dark](themes/zenwritten_dark.svg) | `zenwritten_dark` | Zenwritten Dark | 🌙 Dark | `themes.ThemeZenwrittenDark` |
| ![Zenwritten Light](themes/zenwritten_light.svg) | `zenwritten_light` | Zenwritten Light | ☀️ Light | `themes.ThemeZenwrittenLight` |

## Popular Themes

### Dark Themes
- **Dracula** - A dark theme with purple accents
- **Nord** - An arctic, north-bluish color palette
- **Gruvbox Dark** - Retro groove color scheme
- **Tokyo Night** - A dark theme inspired by Tokyo
- **Catppuccin Mocha** - Soothing pastel theme
- **One Dark** - Atom's iconic dark theme
- **Monokai** - Classic syntax highlighting theme

### Light Themes
- **Solarized Light** - Precision colors for light backgrounds
- **Gruvbox Light** - Retro groove for light mode
- **Catppuccin Latte** - Catppuccin's light variant
- **One Light** - Atom's light companion theme

## Adding Custom Themes

See the [Contributing Guide](docs/CONTRIBUTING.md) for instructions on adding new themes.
