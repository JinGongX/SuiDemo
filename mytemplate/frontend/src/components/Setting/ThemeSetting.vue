<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { setTheme, getCurrentTheme, ThemeMode } from '../../utils/ThemeManager'

const themevalue = ref<ThemeMode>('systemdefault')

const themeChange = (val: ThemeMode) => {
  setTheme(val) 
  // 通知子窗口也可以加上 BroadcastChannel
   const bc = new BroadcastChannel('theme')
   bc.postMessage(val)
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