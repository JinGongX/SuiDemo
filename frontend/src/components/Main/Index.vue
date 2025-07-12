<template>
  <div class="h-screen w-screen   flex flex-col bg-white text-black  dark:text-white ">
    <div class="flex flex-1    overflow-hidden">
      <!-- 左侧请求栏 -->
      <div class="w-44 bg-gray-100 font-bold text-base p-3 space-y-2  dark:bg-gray-800" style="padding-top:40px">
        <div v-for="(item, index) in requests" :key="index" @click="selected = item.id"
            :class="[selected === item.id ? 'bg-orange-300 cursor-pointer' : '', 'cursor-pointer p-2 rounded hover:bg-gray-400']">
          {{ item.label }} 
        </div>
         <!-- <button @click="openSecond" style="width:80%">打开子窗口</button> -->
      </div>

      <!-- 主内容 -->
      <div class="flex-1 p-4 overflow-y-auto max-h-screen scroll-container bg-gray-50 dark:bg-gray-950">
        <component :is="getComponent"  />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref ,computed,watch ,nextTick} from 'vue'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()

const requests = computed(() => [
  {id:'general',label:t('menus.general')},
  {id:'shortcut',label:t('menus.shortcut')},
  {id:'about',label:t('menus.about')}
])
const selected = ref('general')
import General from '../General/Index.vue'
import Shortcut from '../Shortcut/Index.vue'
import About from '../About/Index.vue'

const components = {
  general: General,
  shortcut: Shortcut,
  about: About
}

const getComponent = computed(() => components[selected.value])
// 菜单变化时滚动到顶部
watch(selected, () => {
  nextTick(() => {
    const el = document.querySelector('.scroll-container')
    if (el) el.scrollTop = 0
  })
})

// import { OpenSecondWindow } from  '../../../bindings/changeme/appservice'
// function openSecond() {
//   OpenSecondWindow()
// }
</script>

<style scoped>
html, body {
  margin: 0;
  overflow-x: hidden; /* 禁止横向滚动 */
  overflow-y: hidden;
}
.drag-region {
  -webkit-app-region: drag;
}
</style>