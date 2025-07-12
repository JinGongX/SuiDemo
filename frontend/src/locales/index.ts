import en from './en.json'
import zh from './zh.json'
export const availableLocales = ['en', 'zh'] as const
export type Locale = typeof availableLocales[number]

export async function loadLocaleMessages(locale: Locale) {
  const messagesMap = {
    en,
    zh
  }
  return messagesMap[locale]
}