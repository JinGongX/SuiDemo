<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { setTheme, getCurrentTheme, ThemeMode } from '../../utils/ThemeManager'


const themevalue = ref<ThemeMode>('light')

const themeChange = (val: ThemeMode) => {
  setTheme(val)

  const id = 'antd-theme'
  let link = document.getElementById(id) as HTMLLinkElement
  if (!link) {
    link = document.createElement('link')
    link.id = id
    link.rel = 'stylesheet'
    document.head.appendChild(link)
  }
  link.href = val === 'dark'
    ? 'https://cdn.jsdelivr.net/npm/ant-design-vue@latest/dist/antd.dark.css'
    : 'https://cdn.jsdelivr.net/npm/ant-design-vue@latest/dist/antd.css'

    // 通知子窗口也可以加上 BroadcastChannel
  // const bc = new BroadcastChannel('theme')
  // bc.postMessage(val)
}

onMounted(() => {
  themevalue.value = getCurrentTheme()
})
</script>

<template>
  <div class="space-y-2">
    <a-select
      v-model:value="themevalue"
      style="width: 100px"
      @change="themeChange"
    >
      <a-select-option value="light">{{$t('components.themesetting.select.opt_light')}}</a-select-option>
      <a-select-option value="dark">{{$t('components.themesetting.select.opt_dark')}}</a-select-option>
    </a-select>
  </div>
</template>