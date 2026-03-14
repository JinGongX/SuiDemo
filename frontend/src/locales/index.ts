import en from './en.json'
import zh from './zh.json'
import zhHK from './zh-HK.json'
export const availableLocales = ['en', 'zh','zh-HK'] as const
export type Locale = typeof availableLocales[number]

export async function loadLocaleMessages(locale: Locale) {
  const messagesMap = {
    en,
    zh,
    'zh-HK':zhHK
  }
  return messagesMap[locale]
}

 