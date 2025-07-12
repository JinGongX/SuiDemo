<script setup lang="ts">
import Main from './components/Main/Index.vue'
import { onMounted} from 'vue'
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
    <Main  />
</template>

<style scoped>
html, body {
  margin: 0;
  overflow-x: hidden; /* 禁止横向滚动 */
}
/* 可选：全局背景适配 */
body {
  @apply bg-white dark:bg-gray-900 text-black dark:text-white;
}
</style>
