<template>
  <div >
    <a-select  class="dark:text-white"
              ref="select"
              v-model:value="locale"
              style="width: 100px"
            >
              <a-select-option value="zh" >简体中文</a-select-option>
              <a-select-option value="en">English</a-select-option>
              <a-select-option value="zh-HK">繁體中文</a-select-option>
            </a-select>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { i18n, setupI18n } from '../../utils/i18n'
import { Locale } from '../../locales'
import { SetLanguage } from '../../../bindings/changeme/internal/services/appservice'
const locale = computed({
  get: () => i18n.global.locale.value,
  set: (val) => {
    switchLang(val as Locale)
  }
})
const switchLang =async  (lang: Locale) => {
  await setupI18n(lang) // 切换语言并更新 i18n 实例
  await SetLanguage(lang)// 更新配置项中的语言设置 and 更新菜单语言
};
</script>