# React Example

TypeScript React application with theme context and components.

## Features

- React Context for global theme state
- `useTheme()` hook for easy access
- Theme switcher component with navigation
- Color palette preview
- localStorage persistence
- Vite for fast development

## Setup

### 1. Generate Theme Files

```bash
cd generate
go run . ../public
```

This creates:
- `public/themes.css` - CSS with `[data-theme="id"]` selectors
- `public/themes.json` - Theme metadata (id, displayName, isDark)

### 2. Install Dependencies

```bash
npm install
```

### 3. Run Development Server

```bash
npm run dev
```

Then open http://localhost:3000

## Components

### ThemeProvider

Wraps your app to provide theme context:

```tsx
import { ThemeProvider } from './ThemeContext';

<ThemeProvider defaultTheme="dracula">
  <App />
</ThemeProvider>
```

### useTheme Hook

Access theme state anywhere:

```tsx
import { useTheme } from './ThemeContext';

function MyComponent() {
  const { currentTheme, setTheme, nextTheme, isDark } = useTheme();

  return (
    <button onClick={() => setTheme('nord')}>
      Switch to Nord
    </button>
  );
}
```

### ThemeSwitcher

Dropdown with optional navigation buttons:

```tsx
import { ThemeSwitcher } from './ThemeSwitcher';

<ThemeSwitcher showNavigation={true} />
```

## Project Structure

```
react/
├── generate/           # Go theme generator
│   ├── main.go
│   └── go.mod
├── public/             # Generated assets
│   ├── themes.css
│   └── themes.json
├── src/
│   ├── App.tsx         # Main app component
│   ├── App.css         # App styles
│   ├── ThemeContext.tsx # React context + provider
│   ├── ThemeSwitcher.tsx # UI components
│   └── main.tsx        # Entry point
├── index.html
├── package.json
├── tsconfig.json
└── vite.config.ts
```

## Customization

### Adding More Themes

Edit `generate/main.go`:

```go
ids := []string{
    "dracula", "nord", "gruvbox_dark",
    // Add more theme IDs...
}
```

Then regenerate: `cd generate && go run . ../public`

### Using All Themes

To include all 451+ themes:

```go
// In generate/main.go
allThemes := themes.All()
css := gothememe.GenerateAllThemesCSS(allThemes, opts)
```

Note: This creates a larger CSS file (~2MB).
