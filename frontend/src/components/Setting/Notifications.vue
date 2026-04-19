<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { message } from 'ant-design-vue';
import { CloseOutlined, DeleteOutlined, SyncOutlined } from '@ant-design/icons-vue'
const props = defineProps<{
  handleCancel: boolean,
}>()
const emit = defineEmits<{
  (e: 'notification-change', value: true | false): void
}>()
function onCloseChange() {
  emit('notification-change', false)
}
interface Activity {
  id: number
  title: string
  content?: string
  type: string
  created_at: number
}

const activities_data = ref<Activity[]>([
  {
    id: 1,
    title: 'System Update',
    content: 'Your system was updated successfully.',
    type: 'system',
    created_at: Math.floor(Date.now() / 1000) - 300
  },
  {
    id: 2,
    title: 'Auto-complete Task',
    content: 'Task "Write report" was auto-completed.',
    type: 'auto_complete',
    created_at: Math.floor(Date.now() / 1000) - 7200
  },
  {
    id: 3,
    title: 'Alert: High CPU Usage',
    content: 'Your CPU usage has been above 90% for the last hour.',
    type: 'alert',
    created_at: Math.floor(Date.now() / 1000) - 90000
  },
  {
    id: 4,
    title: 'New Message',
    content: 'You have received a new message from John.',
    type: 'message',
    created_at: Math.floor(Date.now() / 1000) - 120
  },
  {
    id: 5,
    title: 'Backup Completed',
    content: 'Your files were backed up successfully.',
    type: 'backup',
    created_at: Math.floor(Date.now() / 1000) - 3600
  },
  {
    id: 6,
    title: 'Friend Request',
    content: 'Alice has sent you a friend request.',
    type: 'social',
    created_at: Math.floor(Date.now() / 1000) - 18000
  }
])
const activities = ref<Activity[]>([])
// 时间格式
function formatTime(ts: number) {
  const now = Date.now() / 1000
  const diff = now - ts
  if (diff < 60) return 'just now'
  if (diff < 3600) return Math.floor(diff / 60) + ' mins ago'
  if (diff < 86400) return Math.floor(diff / 3600) + ' hours ago'
  return 'yesterday'
}

// 类型样式
function getTypeStyle(type: string) {
  switch (type) {
    case 'system':
      return {
        icon: 'i-carbon-information',
        bg: 'bg-indigo-100',
        iconColor: 'text-indigo-600'
      }
    case 'auto_complete':
      return {
        icon: 'i-carbon-checkmark-filled',
        bg: 'bg-green-100',
        iconColor: 'text-green-600'
      }
    case 'complete':
      return {
        icon: 'i-carbon-checkmark-filled',
        bg: 'bg-green-200',
        iconColor: 'text-green-800'
      }
    case 'alert':
      return {
        icon: 'i-carbon-warning-filled',
        bg: 'bg-orange-100',
        iconColor: 'text-orange-600'
      }
    case 'create':
      return {
        icon: 'i-carbon-add-filled',
        bg: 'bg-blue-100',
        iconColor: 'text-blue-600'
      }
    default:
      return {
        icon: 'i-carbon-dot-mark',
        bg: 'bg-gray-100',
        iconColor: 'text-gray-500'
      }
  }
}  
onMounted(async () => {
  await onloadActivities()
})
const onloadActivities = async () => {
  try {
    activities.value = activities_data.value
  } catch (error) {
    console.error('Failed to load activities:', error)
  }
}
const clearActivities = async () => {
  try {
    activities.value = []
  } catch (error) {
    console.error('Failed to clear activities:', error)
  }
}
const clearActivity = async (id: number) => {
  try {
    activities.value = activities.value.filter(item => item.id !== id)
    message.success('Deleted successfully.')
  } catch (error) {
    console.error('Failed to clear activity:', error)
  }
}
</script>

<template>
  <div class="flex items-center justify-between mb-4 mt-[-4px]">
    <div class="text-sm font-semibold flex items-center gap-2 text-slate-800 dark:text-white">
      <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 dark:text-gray-200" fill="none" viewBox="0 0 24 24"
        stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
          d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
      </svg>
      Notifications
    </div>
    <div class=" cursor-pointer text-items-center text-sm flex gap-4">
      <DeleteOutlined @click="clearActivities" class="text-red-400 hover:text-red-600" />
      <SyncOutlined @click="onloadActivities" class="text-yellow-400 hover:text-yellow-600" />
      <CloseOutlined @click="onCloseChange" class="text-gray-400 hover:text-gray-600" />
    </div>
  </div>

  <div class="bg-white h-[320px] ">
    <!-- 空状态 -->
    <div v-if="activities.length === 0" class="text-center text-gray-400 text-sm py-6">
      <img src="/system/reminders.svg" alt="No activity" class="w-28 h-28 mx-auto mb-2" />
      No activity yet
    </div>
    <div v-else class="space-y-1 overflow-y-auto h-full scrollbar-thin pr-2 ">
      <div v-for="item in activities" :key="item.id" class="flex gap-2 group"> 
        <div class="flex-1 p-1 rounded-xl transition-all duration-200
                 hover:bg-gray-50 ">
          <div class="flex items-start justify-between gap-3">
            <div class="flex items-start gap-3"> 
              <div class="w-9 h-9 rounded-xl flex items-center justify-center" :class="getTypeStyle(item.type).bg">
                <i :class="[
                  getTypeStyle(item.type).icon,
                  getTypeStyle(item.type).iconColor,
                  'text-lg'
                ]" />
              </div>
            </div> 
            <div class="flex-1  sm:w-auto">
              <div class="flex items-start justify-between gap-2 ">
                <!-- 标题 -->
                <div class="text-xs  font-medium text-gray-800   break-words line-clamp-1">
                  {{ item.title }}
                </div>
                <!-- 时间 -->
                <div class="text-xs text-gray-400 whitespace-nowrap shrink-0  ">
                  {{ formatTime(item.created_at) }} 
                    <CloseOutlined @click="clearActivity(item.id)"
                      class="text-red-400 hover:text-red-600 pl-1  group-hover:opacity-100 transition-opacity cursor-pointer" />
                </div>
              </div>
              <div v-if="item.content" class="text-xs text-gray-500 mt-1">
                {{ item.content }}
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="w-[60%]  items-center justify-center  mx-auto  ">
        <a-divider dashed plain>END</a-divider>
      </div>
    </div>
  </div>
</template>
<style scoped>
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

.ant-popover {
  font-size: 10px;
}
</style>