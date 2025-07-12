// src/utils/ThemeManager.ts
const THEME_KEY = 'theme'

export type ThemeMode = 'light' | 'dark' | 'systemdefault'

export function applyTheme(mode: ThemeMode) {
  let resolvedTheme = mode
  if (mode === 'systemdefault') {
    resolvedTheme = window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light'
  }

  document.documentElement.classList.remove('dark', 'light')
  document.documentElement.classList.add(resolvedTheme)
}

export function initTheme() {
  const savedTheme = (localStorage.getItem(THEME_KEY) || 'systemdefault') as ThemeMode
  applyTheme(savedTheme)

  // 监听系统变化，仅在 systemdefault 时响应
  window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', () => {
    const current = getCurrentTheme()
    if (current === 'systemdefault') {
      applyTheme('systemdefault')
    }
  })
}

export function setTheme(mode: ThemeMode) {
  localStorage.setItem(THEME_KEY, mode)
  applyTheme(mode)
}

export function getCurrentTheme(): ThemeMode {
  return (localStorage.getItem(THEME_KEY) || 'systemdefault') as ThemeMode
}