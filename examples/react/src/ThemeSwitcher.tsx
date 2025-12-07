import React from 'react';
import { useTheme } from './ThemeContext';

interface ThemeSwitcherProps {
  showNavigation?: boolean;
  className?: string;
}

export function ThemeSwitcher({ showNavigation = true, className = '' }: ThemeSwitcherProps) {
  const { currentTheme, themes, setTheme, nextTheme, previousTheme } = useTheme();

  if (themes.length === 0) {
    return <div className={className}>Loading themes...</div>;
  }

  return (
    <div className={`theme-switcher ${className}`}>
      {showNavigation && (
        <button
          type="button"
          onClick={previousTheme}
          aria-label="Previous theme"
          className="theme-nav-btn"
        >
          ←
        </button>
      )}

      <select
        value={currentTheme?.id || ''}
        onChange={(e) => setTheme(e.target.value)}
        aria-label="Select theme"
        className="theme-select"
      >
        {themes.map((theme) => (
          <option key={theme.id} value={theme.id}>
            {theme.displayName} {theme.isDark ? '🌙' : '☀️'}
          </option>
        ))}
      </select>

      {showNavigation && (
        <button
          type="button"
          onClick={nextTheme}
          aria-label="Next theme"
          className="theme-nav-btn"
        >
          →
        </button>
      )}
    </div>
  );
}

export function ThemeIndicator({ className = '' }: { className?: string }) {
  const { currentTheme, isDark } = useTheme();

  if (!currentTheme) {
    return null;
  }

  return (
    <div className={`theme-indicator ${className}`}>
      <span className="theme-name">{currentTheme.displayName}</span>
      <span className="theme-mode">{isDark ? '🌙 Dark' : '☀️ Light'}</span>
    </div>
  );
}
