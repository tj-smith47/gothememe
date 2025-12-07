import React, { createContext, useContext, useState, useEffect, ReactNode } from 'react';

interface Theme {
  id: string;
  displayName: string;
  isDark: boolean;
}

interface ThemeContextType {
  currentTheme: Theme | null;
  themes: Theme[];
  setTheme: (themeId: string) => void;
  nextTheme: () => void;
  previousTheme: () => void;
  isDark: boolean;
}

const ThemeContext = createContext<ThemeContextType | undefined>(undefined);

interface ThemeProviderProps {
  children: ReactNode;
  defaultTheme?: string;
  storageKey?: string;
}

export function ThemeProvider({
  children,
  defaultTheme = 'dracula',
  storageKey = 'gothememe-theme',
}: ThemeProviderProps) {
  const [themes, setThemes] = useState<Theme[]>([]);
  const [currentThemeId, setCurrentThemeId] = useState<string>(() => {
    if (typeof window !== 'undefined') {
      return localStorage.getItem(storageKey) || defaultTheme;
    }
    return defaultTheme;
  });

  // Load themes from JSON
  useEffect(() => {
    fetch('/themes.json')
      .then((res) => res.json())
      .then((data: Theme[]) => setThemes(data))
      .catch((err) => console.error('Failed to load themes:', err));
  }, []);

  // Apply theme to document
  useEffect(() => {
    document.documentElement.setAttribute('data-theme', currentThemeId);
    localStorage.setItem(storageKey, currentThemeId);
  }, [currentThemeId, storageKey]);

  const currentTheme = themes.find((t) => t.id === currentThemeId) || null;

  const setTheme = (themeId: string) => {
    if (themes.some((t) => t.id === themeId)) {
      setCurrentThemeId(themeId);
    }
  };

  const nextTheme = () => {
    const idx = themes.findIndex((t) => t.id === currentThemeId);
    const nextIdx = (idx + 1) % themes.length;
    setCurrentThemeId(themes[nextIdx]?.id || currentThemeId);
  };

  const previousTheme = () => {
    const idx = themes.findIndex((t) => t.id === currentThemeId);
    const prevIdx = (idx - 1 + themes.length) % themes.length;
    setCurrentThemeId(themes[prevIdx]?.id || currentThemeId);
  };

  const value: ThemeContextType = {
    currentTheme,
    themes,
    setTheme,
    nextTheme,
    previousTheme,
    isDark: currentTheme?.isDark ?? true,
  };

  return <ThemeContext.Provider value={value}>{children}</ThemeContext.Provider>;
}

export function useTheme(): ThemeContextType {
  const context = useContext(ThemeContext);
  if (context === undefined) {
    throw new Error('useTheme must be used within a ThemeProvider');
  }
  return context;
}
