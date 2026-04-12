<template>
  <div class="flex h-screen  font-sans text-slate-900">
    <!-- 左侧请求栏 -->
    <aside
      class="w-44 border-r border-slate-200  bg-white flex flex-col dark:bg-gray-800 dark:text-white dark:border-slate-600"
      :class="ismacos ? 'pt-10 ' : 'pt-8'">
      <nav class="flex-1  px-4 space-y-1 dark:bg-gray-800">
        <div v-for="(item, index) in requests" :key="index" @click="handleMenu(item)"
          :class="[' flex items-center gap-3 px-3 py-2  text-indigo-700 cursor-pointer rounded-xl font-medium dark:text-white', selected === item.id ? 'bg-indigo-50 dark:bg-slate-500' : ' hover:bg-slate-50 text-slate-500 dark:hover:bg-slate-700']">
          <component :is="item.icon" />
          {{ item.label }}
        </div>
      </nav>
      <div class="pl-4 pr-4 pt-2 pb-2 border-t border-slate-100 flex items-center gap-3 dark:border-slate-600 ">
        <img src="https://api.dicebear.com/7.x/avataaars/svg?seed=JinGong8"
          class="w-10 h-10 rounded-full bg-slate-200" />
        <div class="text-sm">
          <p class="font-bold mb-0">Jin Gong</p>
          <p class="text-slate-500 text-xs mb-0">Pro Plan</p>
        </div>
      </div>
    </aside>
    <!-- 右侧整体 -->
    <div class="flex-1 flex flex-col bg-gray-50 dark:bg-gray-950">
      <div v-if="iswindows" class="titlebar drag-region pt-4 px-4 mb-2  text-white flex items-center justify-between">
        <div class="drag-area">
        </div>
        <div class="window-controls">
          <button @click="Window.Minimise()" >
            <LineOutlined class="text-gray-400 hover:text-gray-600 dark:text-gray-300 dark:hover:text-gray-100" />
          </button>
          <button  @click="toggleMaximize()">
            <BorderOutlined v-if="!isMaximized"  class="text-gray-400 hover:text-gray-600 dark:text-gray-300 dark:hover:text-gray-100" />
            <SwitcherOutlined v-else   class="text-gray-400 hover:text-gray-600 dark:text-gray-300 dark:hover:text-gray-100" />
          </button> 
          <button @click="Window.Hide()" >
            <CloseOutlined class="text-gray-400 hover:text-gray-600 dark:text-gray-300 dark:hover:text-gray-100" />
          </button>
        </div>
      </div>
      <!-- 主内容 -->
      <main class="flex-1  overflow-y-auto max-h-screen scroll-container  scrollbar-thin" :class="iswindows ? 'px-4 pt-px pb-4' : 'p-4'">
        <component :is="getComponent" />
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, nextTick, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import { SettingOutlined, BulbOutlined, ScheduleOutlined, DashboardOutlined } from '@ant-design/icons-vue';
import { OpenSecondWindow } from '../../../bindings/changeme/internal/services/appservice'
import { getOS, OS_READY } from '../../utils/osinfo'
import { CloseOutlined, LineOutlined, BorderOutlined, SwitcherOutlined } from '@ant-design/icons-vue'

const ismacos = ref(false)
const iswindows= ref(false)
const { t } = useI18n()
const isMaximized = ref(false)
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
  { id: 'dashboard', label: t('menus.dashboard'), icon: DashboardOutlined, type: 'component' },
  { id: 'light-tip', label: t('menus.second'), icon: ScheduleOutlined, type: 'window', action: OpenSecondWindow },
  { id: 'setting', label: t('menus.setting'), icon: SettingOutlined, type: 'component' },
  { id: 'about', label: t('menus.about'), icon: BulbOutlined, type: 'component' },
])
const selected = ref('dashboard')
import Dashboard from '../Dashboard/Index.vue'
import Setting from '../General/Index.vue'
import About from '../About/Index.vue'
import { Window } from '@wailsio/runtime';


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

onMounted(async () => {
  await OS_READY
  const osname = getOS() //判断是否为macos 
  if (osname === 'darwin') {
    ismacos.value = true;
  }else if (osname === 'windows') {
    iswindows.value = true;
    isMaximized.value = await Window.IsMaximised()
  }else{ 
  }
})
async function toggleMaximize() {
  isMaximized.value = !isMaximized.value
  await Window.ToggleMaximise()
}
</script>

<style scoped>
html,
body {
  margin: 0;
  overflow-x: hidden;
  /* 禁止横向滚动 */
  overflow-y: hidden;
}

.drag-region {
  -webkit-app-region: drag;
  --wails-draggable: drag;
}

.titlebar {
  height: 40px;
  display: flex;
  justify-content: center;
  align-items: center;
  color: #fff;
  user-select: none;
}

.drag-area {
  flex: 1;
  padding-left: 12px;
  /* 👇 关键：允许拖动窗口 */
  -webkit-app-region: drag;
  --wails-draggable: drag;
}

.window-controls {
  display: flex;
   /* 👇 按钮不能拖动 */
  -webkit-app-region: no-drag;
  --wails-draggable: no-drag;
}

.window-controls button {
  margin-right: 8px;
  width: 20px;
  height: 20px;
  font-size: 13px;
  border: none;
  background: transparent;
  color: white;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
 
}

/* 全局样式（可放在 main.css 或 tailwind.css） */
.scrollbar-thin::-webkit-scrollbar {
  width: 6px;
  height: 6px;
}

.scrollbar-thin::-webkit-scrollbar-thumb {
  background-color: #c1c1c1;
  border-radius: 4px;
}

.dark .scrollbar-thin::-webkit-scrollbar-thumb {
  background-color: #555;
}

.scrollbar-thin::-webkit-scrollbar-track {
  background-color: transparent;
}
</style>