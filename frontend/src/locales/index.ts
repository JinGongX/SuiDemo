import en from './en.json'
import zh from './zh.json'
import zhhk from './zh-HK.json'
export const availableLocales = ['en', 'zh','zhhk'] as const
export type Locale = typeof availableLocales[number]

export async function loadLocaleMessages(locale: Locale) {
  const messagesMap = {
    en,
    zh,
    zhhk
  }
  return messagesMap[locale]
}

 