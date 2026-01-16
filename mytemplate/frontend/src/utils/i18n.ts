// src/i18n.ts
import { createI18n } from 'vue-i18n'
import { Locale, loadLocaleMessages } from '../locales'

const defaultLocale: Locale = 'zh'
export const i18n = createI18n({
  legacy: false, // 使用 Composition API
  locale: defaultLocale, // 默认语言
  fallbackLocale: 'en',
  messages: {},
})

//export default i18n
// 初始化语言（只加载一次）
export async function setupI18n(locale: Locale) {
  const messages = await loadLocaleMessages(locale)
  i18n.global.setLocaleMessage(locale, messages)
  i18n.global.locale.value = locale
}