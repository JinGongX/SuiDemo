<template>
  <div class="flex h-screen  font-sans text-slate-900">
      <!-- 左侧请求栏 -->
      <aside class="w-44 border-r border-slate-200  bg-white flex flex-col dark:bg-gray-800 dark:text-white dark:border-slate-600" :class="ismacos?'pt-10 ':''">
          <nav class="flex-1  px-4 space-y-1 dark:bg-gray-800"  >
            <div  v-for="(item, index) in requests" :key="index" @click="handleMenu(item)"  :class="[' flex items-center gap-3 px-3 py-2  text-indigo-700 cursor-pointer rounded-xl font-medium dark:text-white',selected === item.id ?'bg-indigo-50 dark:bg-slate-500':' hover:bg-slate-50 text-slate-500 dark:hover:bg-slate-700'  ]">
             <component :is="item.icon" />   
             {{ item.label }}
            </div>
          </nav>
          <div class="pl-4 pr-4 pt-2 pb-2 border-t border-slate-100 flex items-center gap-3 dark:border-slate-600 ">
            <img src="https://api.dicebear.com/7.x/avataaars/svg?seed=JinGong8" class="w-10 h-10 rounded-full bg-slate-200" />
            <div class="text-sm">
              <p class="font-bold mb-0">Jin Gong</p>
              <p class="text-slate-500 text-xs mb-0">Pro Plan</p>
            </div>
          </div>
      </aside>
      <!-- 主内容 -->
      <main class="flex-1 p-4 overflow-y-auto max-h-screen scroll-container bg-gray-50 dark:bg-gray-950 scrollbar-thin">
        <component :is="getComponent"  />
      </main>
  </div>
</template>

<script setup lang="ts">
import { ref ,computed,watch ,nextTick,onMounted} from 'vue'
import { useI18n } from 'vue-i18n' 
import { SettingOutlined,BulbOutlined ,ScheduleOutlined,DashboardOutlined } from '@ant-design/icons-vue';
import { OpenSecondWindow } from  '../../../bindings/changeme/internal/services/appservice'
import {  getOS,OS_READY } from '../../utils/osinfo'
const ismacos =ref(false)
const { t } = useI18n()
//菜单结构
type MenuItem =
  | {
      id: string
      label: string
      icon: any
      type: 'component'
    }
  | {
      id: string
      label: string
      icon: any
      type: 'window'
      action: () => void
    }
 
const requests = computed<MenuItem[]>(() => [
  {id:'dashboard',label:t('menus.dashboard'),icon:DashboardOutlined,type: 'component'},
  {id: 'light-tip',label: t('menus.second'),icon: ScheduleOutlined,type: 'window',action: OpenSecondWindow},
  {id:'setting',label:t('menus.setting'),icon:SettingOutlined,type: 'component'},
  {id:'about',label:t('menus.about'),icon:BulbOutlined,type: 'component'},
])
const selected = ref('dashboard')
import Dashboard from '../Dashboard/Index.vue'
import Setting from '../General/Index.vue'
import About from '../About/Index.vue'


const components = {
  dashboard: Dashboard,
  setting: Setting,
  about: About
}

const getComponent = computed(() => components[selected.value])
function handleMenu(item: MenuItem) {
  if (item.type === 'component') {
    selected.value = item.id
  }
  if (item.type === 'window') {
    item.action()
  }
}

// 菜单变化时滚动到顶部
watch(selected, () => {
  nextTick(() => {
    const el = document.querySelector('.scroll-container')
    if (el) el.scrollTop = 0
  })
})

onMounted(async() => {
   await OS_READY 
   const osname = getOS() //判断是否为macos 
   if (osname==='darwin'){
      ismacos.value=true;
   }
})
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