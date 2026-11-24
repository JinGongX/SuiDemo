<script setup lang="ts">
import { onMounted} from 'vue'
import { themeConfig } from './utils/ThemeManager' 
const applyTheme = () => {
  const userTheme = localStorage.getItem('theme')
  const systemPrefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches
  const isDark = userTheme === 'dark' || (!userTheme && systemPrefersDark)
  document.documentElement.classList.toggle('dark', isDark)
}

onMounted(() => {
  applyTheme()
  // 可监听系统主题变化自动更新（可选）
  window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', () => {
    applyTheme()
  })
})
</script>

<template>
  <ConfigProvider :theme="themeConfig" >
      <router-view />
  </ConfigProvider>
</template>

<style>
html,
body {
  margin: 0;
  overflow-x: hidden;
}

body {
  @apply bg-white dark:bg-gray-900 text-black dark:text-white;
}
</style>
