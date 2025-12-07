# Framework Integration Guide

GoThemeMe generates CSS that works with any web framework. This guide shows integration patterns for popular frameworks.

## Table of Contents

- [Vanilla HTML/CSS](#vanilla-htmlcss)
- [HTMX + templ](#htmx--templ)
- [React](#react)
- [Vue](#vue)
- [Svelte](#svelte)
- [Tailwind CSS](#tailwind-css)

## Vanilla HTML/CSS

### Basic Setup

1. Generate CSS at build time or serve dynamically:

```go
package main

import (
    "net/http"
    "github.com/tj-smith47/gothememe"
)

func main() {
    gothememe.NewDefaultRegistry()

    // Serve theme CSS
    http.HandleFunc("/theme.css", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/css")
        css := gothememe.AllThemesCSS(gothememe.DefaultCSSOptions())
        w.Write([]byte(css))
    })

    http.ListenAndServe(":8080", nil)
}
```

2. Include in HTML:

```html
<!DOCTYPE html>
<html data-theme="dracula">
<head>
    <link rel="stylesheet" href="/theme.css">
    <style>
        body {
            background-color: var(--theme-background);
            color: var(--theme-text-primary);
        }
        a {
            color: var(--theme-accent);
        }
        .card {
            background: var(--theme-surface);
            border: 1px solid var(--theme-border);
        }
    </style>
</head>
<body>
    <h1>Hello, World!</h1>
    <button onclick="toggleTheme()">Toggle Theme</button>

    <script>
        function toggleTheme() {
            const html = document.documentElement;
            const current = html.getAttribute('data-theme');
            html.setAttribute('data-theme', current === 'dracula' ? 'nord' : 'dracula');
        }
    </script>
</body>
</html>
```

### System Preference Detection

```javascript
// Detect system preference
const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
document.documentElement.setAttribute('data-theme', prefersDark ? 'dracula' : 'nord-light');

// Listen for changes
window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', e => {
    document.documentElement.setAttribute('data-theme', e.matches ? 'dracula' : 'nord-light');
});
```

## HTMX + templ

### Server-Side Theme Switching

```go
// handlers.go
package main

import (
    "net/http"
    "github.com/tj-smith47/gothememe"
)

func init() {
    gothememe.NewDefaultRegistry()
}

func handleThemeSwitch(w http.ResponseWriter, r *http.Request) {
    theme := r.FormValue("theme")
    if gothememe.SetThemeID(theme) {
        // Set cookie for persistence
        http.SetCookie(w, &http.Cookie{
            Name:   "theme",
            Value:  theme,
            Path:   "/",
            MaxAge: 31536000, // 1 year
        })
    }

    // Return updated body with new theme
    w.Header().Set("HX-Trigger", "theme-changed")
    renderPage(w, theme)
}

func handleCSS(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/css")
    css := gothememe.AllThemesCSS(gothememe.DefaultCSSOptions())
    w.Write([]byte(css))
}
```

```templ
// components/layout.templ
package components

templ Layout(theme string, content templ.Component) {
    <!DOCTYPE html>
    <html data-theme={ theme }>
    <head>
        <link rel="stylesheet" href="/theme.css">
        <script src="https://unpkg.com/htmx.org@2"></script>
    </head>
    <body style="background: var(--theme-background); color: var(--theme-text-primary);">
        @ThemeSwitcher(theme)
        @content
    </body>
    </html>
}

templ ThemeSwitcher(current string) {
    <select
        hx-post="/api/theme"
        hx-target="body"
        hx-swap="outerHTML"
        name="theme"
    >
        <option value="dracula" selected?={ current == "dracula" }>Dracula</option>
        <option value="nord" selected?={ current == "nord" }>Nord</option>
    </select>
}
```

## React

### Theme Context Provider

```tsx
// ThemeContext.tsx
import React, { createContext, useContext, useEffect, useState } from 'react';

type Theme = 'dracula' | 'nord' | 'gruvbox-dark';

interface ThemeContextValue {
    theme: Theme;
    setTheme: (theme: Theme) => void;
}

const ThemeContext = createContext<ThemeContextValue | null>(null);

export function ThemeProvider({ children }: { children: React.ReactNode }) {
    const [theme, setTheme] = useState<Theme>(() => {
        return (localStorage.getItem('theme') as Theme) || 'dracula';
    });

    useEffect(() => {
        document.documentElement.setAttribute('data-theme', theme);
        localStorage.setItem('theme', theme);
    }, [theme]);

    return (
        <ThemeContext.Provider value={{ theme, setTheme }}>
            {children}
        </ThemeContext.Provider>
    );
}

export function useTheme() {
    const context = useContext(ThemeContext);
    if (!context) {
        throw new Error('useTheme must be used within ThemeProvider');
    }
    return context;
}
```

```tsx
// App.tsx
import { ThemeProvider, useTheme } from './ThemeContext';

function ThemeSwitcher() {
    const { theme, setTheme } = useTheme();

    return (
        <select value={theme} onChange={e => setTheme(e.target.value as any)}>
            <option value="dracula">Dracula</option>
            <option value="nord">Nord</option>
        </select>
    );
}

function App() {
    return (
        <ThemeProvider>
            <div style={{
                background: 'var(--theme-background)',
                color: 'var(--theme-text-primary)',
                minHeight: '100vh'
            }}>
                <ThemeSwitcher />
                <h1>Hello, React!</h1>
            </div>
        </ThemeProvider>
    );
}
```

### TypeScript Types

Generate TypeScript types from Go (add to your build):

```go
// Generate TypeScript definitions
func generateTypeScript() string {
    return `
export type Theme = 'dracula' | 'nord' | 'gruvbox-dark' | 'tokyo-night';

export interface ThemeColors {
    '--theme-background': string;
    '--theme-text-primary': string;
    '--theme-accent': string;
    // ... add all variables
}
`
}
```

## Vue

### Composable

```vue
<!-- useTheme.ts -->
<script setup lang="ts">
import { ref, watch, onMounted } from 'vue';

type Theme = 'dracula' | 'nord';

const theme = ref<Theme>('dracula');

onMounted(() => {
    const saved = localStorage.getItem('theme') as Theme;
    if (saved) theme.value = saved;
    document.documentElement.setAttribute('data-theme', theme.value);
});

watch(theme, (newTheme) => {
    document.documentElement.setAttribute('data-theme', newTheme);
    localStorage.setItem('theme', newTheme);
});

function toggleTheme() {
    theme.value = theme.value === 'dracula' ? 'nord' : 'dracula';
}
</script>

<!-- App.vue -->
<template>
    <div class="app">
        <button @click="toggleTheme">Toggle Theme</button>
        <h1>Hello, Vue!</h1>
    </div>
</template>

<style>
.app {
    background: var(--theme-background);
    color: var(--theme-text-primary);
    min-height: 100vh;
}

button {
    background: var(--theme-accent);
    color: var(--theme-text-inverted);
    border: none;
    padding: 0.5rem 1rem;
    cursor: pointer;
}
</style>
```

## Svelte

### Store-Based Theme

```svelte
<!-- stores/theme.ts -->
<script context="module" lang="ts">
import { writable } from 'svelte/store';

type Theme = 'dracula' | 'nord';

function createThemeStore() {
    const { subscribe, set } = writable<Theme>('dracula');

    return {
        subscribe,
        set: (theme: Theme) => {
            document.documentElement.setAttribute('data-theme', theme);
            localStorage.setItem('theme', theme);
            set(theme);
        },
        toggle: () => {
            const current = document.documentElement.getAttribute('data-theme');
            const next = current === 'dracula' ? 'nord' : 'dracula';
            document.documentElement.setAttribute('data-theme', next);
            localStorage.setItem('theme', next);
            set(next);
        }
    };
}

export const theme = createThemeStore();
</script>

<!-- App.svelte -->
<script lang="ts">
import { theme } from './stores/theme';
import { onMount } from 'svelte';

onMount(() => {
    const saved = localStorage.getItem('theme') as 'dracula' | 'nord';
    if (saved) theme.set(saved);
});
</script>

<div class="app">
    <button on:click={() => theme.toggle()}>Toggle Theme</button>
    <h1>Hello, Svelte!</h1>
</div>

<style>
.app {
    background: var(--theme-background);
    color: var(--theme-text-primary);
    min-height: 100vh;
}
</style>
```

## Tailwind CSS

### Extending Tailwind with Theme Variables

```javascript
// tailwind.config.js
module.exports = {
    theme: {
        extend: {
            colors: {
                theme: {
                    bg: 'var(--theme-background)',
                    'bg-secondary': 'var(--theme-background-secondary)',
                    surface: 'var(--theme-surface)',
                    text: 'var(--theme-text-primary)',
                    'text-secondary': 'var(--theme-text-secondary)',
                    'text-muted': 'var(--theme-text-muted)',
                    accent: 'var(--theme-accent)',
                    border: 'var(--theme-border)',
                    success: 'var(--theme-success-text)',
                    warning: 'var(--theme-warning-text)',
                    error: 'var(--theme-error-text)',
                }
            }
        }
    }
}
```

Usage:

```html
<div class="bg-theme-bg text-theme-text border-theme-border">
    <h1 class="text-theme-accent">Hello, Tailwind!</h1>
    <p class="text-theme-text-secondary">Themed with GoThemeMe</p>
</div>
```

## Embedding CSS at Build Time

For static sites, generate CSS at build time:

```go
//go:generate go run ./cmd/generate-css

package static

import _ "embed"

//go:embed theme.css
var ThemeCSS string
```

```go
// cmd/generate-css/main.go
package main

import (
    "os"
    "github.com/tj-smith47/gothememe"
)

func main() {
    gothememe.NewDefaultRegistry()
    css := gothememe.AllThemesCSS(gothememe.DefaultCSSOptions())
    os.WriteFile("static/theme.css", []byte(css), 0644)
}
```

## Server-Side Rendering

For SSR applications, include the theme in the initial HTML:

```go
func renderPage(w http.ResponseWriter, theme string) {
    data := struct {
        Theme string
        CSS   string
    }{
        Theme: theme,
        CSS:   gothememe.CSS(gothememe.CSSOptions{
            Prefix:           "theme",
            UseDataAttribute: true,
        }),
    }

    tmpl.Execute(w, data)
}
```

Template:
```html
<html data-theme="{{.Theme}}">
<head>
    <style>{{.CSS}}</style>
</head>
```
