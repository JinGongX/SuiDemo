// src/utils/ThemeManager.ts
import { theme } from 'ant-design-vue'
import type { ThemeConfig } from 'ant-design-vue/es/config-provider/context'
import { ref } from 'vue'

export const lightTheme = {
  algorithm: theme.defaultAlgorithm,
}

export const darkTheme = {
  algorithm: theme.darkAlgorithm,
}

const THEME_KEY = 'theme'
export type ThemeMode = 'light' | 'dark' | 'systemdefault'
export function applyTheme(mode: ThemeMode) { 
  let resolvedTheme = mode
  if (mode === 'systemdefault') {
    resolvedTheme = window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light'
  }

  document.documentElement.classList.remove('dark', 'light')
  document.documentElement.classList.add(resolvedTheme)

  themeConfig.value = getAntdCurrentTheme(resolvedTheme)
}

export function initTheme() {
  const savedTheme = (localStorage.getItem(THEME_KEY) || 'light') as ThemeMode
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
  return (localStorage.getItem(THEME_KEY) || 'light') as ThemeMode
}
// 全局响应式主题状态
//export const themeMode = ref<ThemeMode>(getCurrentTheme())
export const currentMode = ref<'light' | 'dark'>(localStorage.getItem('theme') as 'light' | 'dark' || 'light')
export const themeConfig = ref<ThemeConfig>(getAntdCurrentTheme(currentMode.value))
export function getAntdCurrentTheme(mode: ThemeMode): ThemeConfig {
   return mode === 'dark' ? darkTheme : lightTheme
}
