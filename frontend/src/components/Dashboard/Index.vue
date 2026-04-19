<template>
  <div class="  text-slate-800 font-sans  dark:text-white">
    <div class="max-w-4xl mx-auto pl-6 pr-6">
      <header class="flex items-center justify-between mb-6">
        <div class="relative w-2/3">
          <span class="absolute inset-y-0 left-0 pl-3 flex items-center text-slate-400">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </span>
          <input 
            type="text" 
            placeholder="Search reminders..." 
            class="w-full dark:bg-gray-800 pl-10 pr-4 py-2 bg-slate-100 border-none rounded-xl focus:ring-2 focus:ring-indigo-500 transition-all outline-none text-sm"
          />
        </div>
        <div class="flex items-center gap-3">
          <button @click="showMsgModal"  class="relative  text-slate-600 w-[40px]">
            <div class="absolute mr-2 top-2 right-2 w-2 h-2 bg-red-500 rounded-full border-2 border-white dark:border-gray-200"></div>
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 dark:text-gray-200" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
            </svg>
          </button>
          <button @click="OpenSecondWindow" class="w-[110px] bg-indigo-600 hover:bg-indigo-700 text-white text-sm px-3 py-2 rounded-xl flex items-center gap-2 font-medium transition-colors shadow-sm">
            <span class="text-xl pb-1">+</span> {{ $t('components.dashboard.buttons.newtask') }}
          </button>
        </div>
      </header>

      <section class="mb-4  dark:text-white ">  
        <span class="text-xl font-bold text-slate-900 text-left dark:text-white">{{ greetingMessage }}, {{ profile.name }}</span>
        <p class="text-slate-500 mt-1 text-left dark:text-slate-300" v-if="today_pending>0">You have <span class="font-semibold text-slate-600 dark:text-slate-100" v-text="today_pending"></span> tasks to complete today.</p>
      </section>

      <section class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-10">
        <div v-for="stat in stats" :key="stat.label" class="dark:bg-gray-800 dark:text-white  bg-white pt-2 pr-4 pl-4  rounded-3xl shadow-sm 0">
          <div class="flex justify-between  mb-2">
            <div :class="`p-1 rounded-xl  ${stat.iconBg} `">
              <component :is="stat.icon" class="w-6 h-6 flex items-center justify-center " :class="stat.iconColor" />
            </div>
            <span :class="tagStyles(stat.tagType)">{{ stat.tag }}</span>
          </div>
          <p class="text-slate-500 text-sm font-medium text-left dark:text-white">{{ stat.label }}</p>
          <h2 class="text-2xl font-bold mt-1 text-slate-800 dark:text-white">{{ stat.value }}</h2>
        </div>
      </section>

      <section class="mb-4 flex items-center justify-between">
        <span class="text-base font-bold text-slate-800 dark:text-white">{{ $t('components.dashboard.title.upcoming_task') }}</span>
        <div class="flex bg-slate-200/60 p-1 rounded-lg text-sm dark:bg-gray-800">
          <button class="px-4 py-1 rounded-md bg-white shadow-sm font-medium dark:text-white dark:bg-gray-900">{{ $t('components.dashboard.buttons.upcoming_all') }}</button>
          <button class="px-4 py-1 w-[66px] rounded-md text-slate-500 hover:text-slate-700 dark:text-slate-300">{{ $t('components.dashboard.buttons.upcoming_priority') }}</button>
        </div>
      </section>

      <div class="space-y-1 ">
        <div v-if="tasks.length === 0" class="text-center py-10">
          <img src="/system/empty.svg" class=" mx-auto mt-1 w-32 "  />
          <span class="text-sm">{{ $t('components.dashboard.label.no_tasks') }}</span>
        </div>
        <div v-else v-for="task in tasks" :key="task.id" 
             class="dark:bg-gray-800 dark:text-white dark:border-slate-500 group flex items-center bg-white p-2 rounded-3xl border border-slate-100 hover:shadow-md transition-shadow cursor-pointer">
          <div class="w-6 h-6 rounded-full border-2 border-slate-200 mr-4 flex items-center justify-center group-hover:border-indigo-400 transition-colors">
            <div v-if="task.isCompleted" class="w-4 h-4 bg-indigo-500 rounded-full"></div>
          </div> 
          <div class="flex-1 pl-3">
            <span class="text-slate-800 text-left text-sm dark:text-white line-clamp-1">{{ task.title }}</span>
            <div class="flex items-center gap-4 ">
              <span class="flex items-center text-xs text-slate-400 font-medium">
                <svg v-if="task.time.includes(':')" xmlns="http://www.w3.org/2000/svg" class="h-3.5 w-3.5 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-3.5 w-3.5 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
                
                {{task.time}}
              </span>
              
              <span v-if="task.priority" class="flex items-center text-xs text-orange-500 font-bold uppercase tracking-wider">
                <span class="mr-1">🚩</span> {{ task.priority }}
              </span>

              <span  :class="categoryStyles(task.category)">
                {{ task.category }}
              </span>
            </div>
          </div>
        </div>

      </div>
    </div>
  </div>
  <a-modal v-model:open="openmessage" title="" width="400px" :closable="false" :footer="null">
   <div class="bg-white h-[360px]  pl-1 pr-1 pb-4 pt-1 ">  
      <Notifications :handle-cancel="openmessage" @notification-change="handleOk" /> 
  </div>
  </a-modal>
</template>

<script setup lang="ts">
import { onMounted, ref ,computed} from 'vue'
import { ProfileOutlined,FileDoneOutlined ,ExceptionOutlined  } from '@ant-design/icons-vue'
import  Notifications from '../Setting/Notifications.vue'
import { OpenSecondWindow } from  '../../../bindings/changeme/internal/services/appservice'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
const today_pending = ref(0)
const profile = ref({
  name: 'Jin Gong',
  avatar: 'https://i.pravatar.cc/150?img=3'
})
const openmessage = ref<boolean>(false);
const stats = ref([
  { label: t('components.dashboard.label.total_tasks'), value: 520, tag: '+12%', tagType: 'percentage', icon: ProfileOutlined, iconBg: 'bg-indigo-100', iconColor: 'text-indigo-600' },
  { label: t('components.dashboard.label.pending_tasks'), value: 13, tag: 'Urgent', tagType: 'status', icon: ExceptionOutlined, iconBg: 'bg-orange-100', iconColor: 'text-orange-600' },
  { label: t('components.dashboard.label.completed_tasks'), value: 14, tag: 'Daily', tagType: 'default', icon: FileDoneOutlined, iconBg: 'bg-emerald-100', iconColor: 'text-emerald-600' },
])
interface Task {
  id: number
  title: string
  time: string
  category: string
  priority?: string
  isCompleted: boolean
}
const tasks = ref<Task[]>([])
tasks.value = [
    // { id: 1, title: 'Quarterly Budget Review', time: '10:30 AM', category: 'WORK', priority: 'High Priority', isCompleted: true },
    // { id: 2, title: 'Buy Weekly Groceries', time: '05:00 PM', category: 'PERSONAL', isCompleted: false },
    // { id: 3, title: 'Client Presentation Design', time: 'Tomorrow', category: 'WORK', isCompleted: false },
    // { id: 4, title: 'Book Gym Slot', time: '08:00 AM', category: 'PERSONAL', isCompleted: false }, 
  ]
const showMsgModal = () => {
  openmessage.value = true;
};
const handleOk = () => {
  openmessage.value = false;
};
const tagStyles = (type: string) => {
  const base = "text-[10px] px-2 my-1 rounded-lg font-bold "
  if (type === 'percentage') return base + "bg-emerald-50 text-emerald-500"
  if (type === 'status') return base + "bg-orange-50 text-orange-500"
  return base + "bg-slate-100 text-slate-400"
}

const categoryStyles = (cat: string) => {
  const base = "text-[8px] px-2 rounded-md font-bold "
  return cat === 'WORK' 
    ? base + "bg-blue-100 text-blue-500" 
    : base + "bg-emerald-100 text-emerald-500"
}
onMounted(async () => {
})
 
const now = ref(new Date())

// 计算问候语
const greetingMessage = computed(() => {
  const hour = now.value.getHours()
  if (hour >= 5 && hour < 12) return 'Good Morning'
  if (hour >= 12 && hour < 18) return 'Good Afternoon'
  if (hour >= 18 && hour < 22) return 'Good Evening'
  return 'Good Night'
})

</script>

<style scoped>
/* 可以在此添加特定字体，如 Inter 或 Plus Jakarta Sans */
</style>