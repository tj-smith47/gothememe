import React from 'react';
import { useTheme } from './ThemeContext';
import { ThemeSwitcher, ThemeIndicator } from './ThemeSwitcher';
import './App.css';

function App() {
  const { currentTheme, isDark } = useTheme();

  return (
    <div className="app">
      <header className="header">
        <h1>GoThemeMe React Example</h1>
        <ThemeSwitcher />
      </header>

      <main className="main">
        <section className="card">
          <h2>Current Theme</h2>
          <ThemeIndicator />
          {currentTheme && (
            <div className="theme-info">
              <p><strong>ID:</strong> {currentTheme.id}</p>
              <p><strong>Mode:</strong> {isDark ? 'Dark' : 'Light'}</p>
            </div>
          )}
        </section>

        <section className="card">
          <h2>Color Palette</h2>
          <div className="color-grid">
            <ColorSwatch name="Background" varName="--theme-background" />
            <ColorSwatch name="Surface" varName="--theme-surface" />
            <ColorSwatch name="Text Primary" varName="--theme-text-primary" />
            <ColorSwatch name="Text Secondary" varName="--theme-text-secondary" />
            <ColorSwatch name="Accent" varName="--theme-accent" />
            <ColorSwatch name="Border" varName="--theme-border" />
          </div>
        </section>

        <section className="card">
          <h2>Semantic Colors</h2>
          <div className="color-grid">
            <ColorSwatch name="Success" varName="--theme-success-text" />
            <ColorSwatch name="Warning" varName="--theme-warning-text" />
            <ColorSwatch name="Error" varName="--theme-error-text" />
            <ColorSwatch name="Info" varName="--theme-info-text" />
          </div>
        </section>

        <section className="card">
          <h2>ANSI Colors</h2>
          <div className="color-grid ansi-grid">
            <ColorSwatch name="Black" varName="--theme-black" />
            <ColorSwatch name="Red" varName="--theme-red" />
            <ColorSwatch name="Green" varName="--theme-green" />
            <ColorSwatch name="Yellow" varName="--theme-yellow" />
            <ColorSwatch name="Blue" varName="--theme-blue" />
            <ColorSwatch name="Purple" varName="--theme-purple" />
            <ColorSwatch name="Cyan" varName="--theme-cyan" />
            <ColorSwatch name="White" varName="--theme-white" />
          </div>
        </section>

        <section className="card">
          <h2>Code Example</h2>
          <pre className="code-block">
            <code>{`function greet(name: string): string {
  // Returns a greeting message
  const message = \`Hello, \${name}!\`;
  console.log(message);
  return message;
}`}</code>
          </pre>
        </section>
      </main>

      <footer className="footer">
        <p>
          Powered by <a href="https://github.com/tj-smith47/gothememe">GoThemeMe</a>
        </p>
      </footer>
    </div>
  );
}

function ColorSwatch({ name, varName }: { name: string; varName: string }) {
  return (
    <div className="color-swatch">
      <div
        className="swatch-color"
        style={{ backgroundColor: `var(${varName})` }}
      />
      <div className="swatch-info">
        <span className="swatch-name">{name}</span>
        <code className="swatch-var">{varName}</code>
      </div>
    </div>
  );
}

export default App;
