<template>
  <div class="h-screen w-screen   flex flex-col  text-black  dark:text-white ">
    <div class="flex flex-1    overflow-hidden"> 
      <!-- 左侧请求栏 -->
      <div class="w-44 bg-gray-100/20 font-bold text-base p-3 space-y-2  dark:bg-gray-800" style="padding-top:40px"> 
          <div v-for="(item, index) in requests" :key="index" @click="handleMenu(item)" 
             :class="[' cursor-pointer p-2 rounded text-left',selected === item.id ? 'bg-orange-300/90' : 'hover:bg-gray-300/70']">
            <component :is="item.icon" :style="['margin-right: 10px;vertical-align: middle;',item.id==='shortcut'?'font-size: 19px':'font-size: 18px']" />
             <span>{{ item.label }}</span>
        </div>
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
import { SettingOutlined,BulbOutlined ,FormOutlined,SlackOutlined ,SendOutlined } from '@ant-design/icons-vue';
import { OpenSecondWindow } from  '../../../bindings/changeme/appservice'

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
  {id:'general',label:t('menus.general'),icon:SettingOutlined,type: 'component'},
  {id:'shortcut',label:t('menus.shortcut'),icon:FormOutlined,type: 'component'},
  {id: 'light-tip',label: t('menus.second'),icon: SlackOutlined,type: 'window',action: OpenSecondWindow},
  {id:'about',label:t('menus.about'),icon:BulbOutlined,type: 'component'},
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