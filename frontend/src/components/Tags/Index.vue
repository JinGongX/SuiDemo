<template>
  <div class="flex h-full ">
    <!-- Left List -->
    <div class=" pl-1 pb-2 pr-1 flex flex-col" :class="currentTag === null ? 'w-3/5 ' : 'w-2/5'">
      <div class="mr-2 ml-1">
        <h2 class="text-lg font-bold text-gray-800 dark:text-white float-left">
          {{ $t('components.tags.title.manage_tags') }}
        </h2>
        <button @click="createNew"
          class="w-[100px] float-right mb-2 bg-indigo-600 hover:bg-indigo-700 text-white text-sm px-2 py-2 rounded-xl flex items-center gap-2 font-medium transition-colors shadow-sm">
          <span class="text-xl pb-1">+</span> {{ $t('components.tags.buttons.newtag') }}
        </button>
      </div>
      <!-- Search -->
      <div class="relative w-full pr-2">
        <FilterOutlined class="absolute pl-1 top-[15px] -translate-y-1/2 text-gray-400 pointer-events-none" />
        <input v-model="keyword"  placeholder="Search Tag..."
          class="w-full h-[30px] outline-none mb-2 pl-6 pr-3 py-1 border rounded-lg text-sm dark:bg-gray-900 dark:text-white dark:border-slate-500" />
      </div>
      <!-- List -->
      <div class="flex-1 overflow-auto space-y-1  scrollbar-thin">
        <div v-for="item in filteredTags" :key="item.id" class="item" :class="{ selected: currentTag?.id === item.id }"
          @click="select(item)">
          <div class="flex h-8 w-8 shrink-0 items-center justify-center rounded mr-1 ml-[-4px] mt-1"
            :class="item.color">
            <img :src="item.icon" class="h-8 w-8 object-contain" draggable="false" />
          </div>
          <div class="flex-1 ">
            <div class=" text-sm text-left line-clamp-1 break-words dark:text-white">
              {{ item.name }}
            </div>
            <div class="text-sm  text-left  text-gray-500 line-clamp-2 dark:text-gray-300">
              {{ item.keywords.join(', ') }}
            </div>
          </div>
        </div>
        <a-divider dashed plain>END</a-divider>
      </div>
    </div>

    <!-- Right Editor -->
    <div class="flex w-3/5 pt-2 pb-6   bg-white  rounded-lg dark:bg-gray-800">
      <div v-if="currentTag" class="w-full">
        <div class="w-full mb-4 px-5  border-b gap-2  items-center  dark:border-gray-700">
          <div class="float-right flex gap-2 pt-1 w-[20%] justify-end">
            <DeleteOutlined v-if="currentTag.is_default != 1 && currentTag.id != 0" @click="deleteTag"
              class="px-1 cursor-pointer float-left  text-red-600" />
            <CheckCircleOutlined class=" px-1  cursor-pointer float-right text-green-600"
              @click="editTag(currentTag)" />
          </div>
          <div class="pb-3  itmes-left">
            <input v-model="currentTag.name"
              class="text-sm  w-[80%] font-semibold  outline-none dark:bg-gray-800 dark:text-white "
              placeholder="Tag Title" />
          </div>
        </div>
        <div class="px-4">
          <div class="mb-4">
            <div class="flex justify-between items-center mb-3">
              <div class="text-xs text-gray-500 dark:text-gray-400 tracking-wider">{{ $t('components.tags.label.tagicon') }}</div>
              <!-- 当前预览 -->
              <div class="flex items-center gap-2 text-xs text-gray-500 dark:text-gray-400">
                <span>{{ $t('components.tags.label.current_preview') }}:</span>
                <div class="w-9 h-9 rounded-lg flex items-center justify-center text-white" :class="selectedColor">
                  <img v-if="customIcon" :src="customIcon" class="w-6 h-6 object-contain" />
                  <img v-else-if="selectedIconComponent" :src="selectedIconComponent" class="w-6 h-6 object-contain" />
                </div>
              </div>
            </div>
            <!-- icon 列表 -->
            <div class="flex items-center gap-3 flex-wrap ">
              <div v-for="item in iconOptions" :key="item.name" @click="selectIcon(item.name)"
                class="shrink-0 w-10 h-10 rounded-lg flex items-center justify-center cursor-pointer transition" :class="[
                  selectedIcon === item.name && !customIcon
                    ? 'bg-indigo-600 text-white shadow'
                    : 'bg-gray-100 text-gray-500 hover:bg-gray-200'
                ]">
                <img :src="item.component" class="w-8 h-8 object-contain" />
              </div>
              <div class="w-px h-6 bg-gray-200 mx-1"></div>
              <label for="fileInput" @click="handleUpload"
                class="w-10 h-10 rounded-lg border-2 border-dashed border-gray-300 dark:border-gray-600 flex items-center justify-center cursor-pointer hover:bg-gray-50 dark:hover:bg-gray-700">
                <UploadOutlined class="text-gray-500 dark:text-gray-200" />
              </label>
            </div>
            <!-- 提示 -->
            <div class="text-xs text-gray-400 mt-2">
              SVG, PNG or JPG (max 100x100px recommended)
            </div>
          </div>
        </div>
        <div class="px-4">
          <div class="mb-2">
            <span class="text-xs text-left  text-gray-500 dark:text-gray-400">
              {{ $t('components.tags.label.tagcolor') }}
            </span>
          </div>
          <div class="flex flex-wrap gap-3">
            <div v-for="color in colors" :key="color" @click="selectedColor = color"
              class="w-5 h-5 rounded-full cursor-pointer border-2 shrink-0 relative"
              :class="[color, selectedColor === color ? 'ring-2 ring-white ring-offset-2 ring-offset-black/20 scale-100 ' : 'border-transparent']">
              <div v-if="color === 'bg-transparent'"
                class="absolute inset-0 bg-[linear-gradient(45deg,#ccc_25%,transparent_25%),linear-gradient(-45deg,#ccc_25%,transparent_25%),linear-gradient(45deg,transparent_75%,#ccc_75%),linear-gradient(-45deg,transparent_75%,#ccc_75%)] bg-[length:6px_6px] bg-[position:0_0,0_3px,3px_-3px,-3px_0px]">
              </div>
            </div>
          </div>
        </div>
        <!-- TAG 子标签 -->
        <div class="px-4">
          <div class="mb-2">
            <span class="text-xs text-left  text-gray-500 dark:text-gray-400">{{ $t('components.tags.label.subtags') }}</span>
          </div> 
          <div class="flex gap-2 mb-3">
            <input v-model="newSubTag" @keyup.enter="addSubTag"
              class="flex-1 h-[30px] bg-gray-100 rounded-lg px-3 py-2 outline-none text-sm  dark:bg-gray-900 dark:text-white"
              placeholder="Add sub tag..." />
            <button @click="addSubTag" class="px-4 py bg-black text-white rounded-lg text-sm ">
              {{ $t('components.tags.buttons.add_subtag') }}
            </button>
          </div>
          <!-- 标签列表 -->
          <div class="flex flex-wrap gap-2">
            <div v-for="(tag, index) in currentTag.keywords" :key="index"
              class="flex items-center gap-2 bg-gray-100 px-2 py-1.5 rounded-lg text-sm  dark:bg-gray-900 dark:text-white">
              <span class="px-1">{{ tag }}</span>
              <button @click="removeSubTag(index)"
                class="w-4 h-4 flex items-center justify-center text-gray-400 hover:text-red-500 ">
                ✕
              </button>
            </div>
          </div>

          <!-- 空状态 -->
          <div v-if="currentTag.keywords.length === 0" class="text-xs text-gray-400 mt-2">
            No sub tags yet
          </div>
        </div>
        <div class="text-right px-4">

        </div>
      </div>
      <div v-else class="text-gray-400 p-6 text-center  mx-auto ">
        <img src="/system/mychoice.svg" class="mx-auto mb-4" style="width: 120px; height: 120px;" />
        Select a left tags to view details
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, createVNode } from 'vue'
import { i18n } from '../../utils/i18n'
import { CheckCircleOutlined, DeleteOutlined, UploadOutlined, ExclamationCircleOutlined, FilterOutlined} from '@ant-design/icons-vue'
 import { message, Modal } from 'ant-design-vue';

const keyword = ref('')
//const tags = ref<{id:number, name:string, color:string}[]>([])
const currentTag = ref<Tag | null>(null)
interface Tag {
  id: number
  tag_key: string
  name: string
  color: string
  icon: string
  keywords: string[]
  is_default: number
  sort: number
  created_at?: number
  updated_at?: number
}

const Tags = ref<Tag[]>([
  {
    id: 1,
    tag_key: 'work',
    name: 'Work',
    color: 'bg-blue-100',
    icon: '/visual/work.png',
    keywords: ['office', 'project', 'meeting'],
    is_default: 0,
    sort: 1
  },
  {
    id: 2,
    tag_key: 'life',
    name: 'Life',
    color: 'bg-green-100',
    icon: '/visual/eat.png',
    keywords: ['home', 'family', 'health'],
    is_default: 0,
    sort: 2
  },
  {
    id: 3,
    tag_key: 'rest',
    name: 'Rest',
    color: 'bg-yellow-100',
    icon: '/visual/coffee.png',
    keywords: ['break', 'relax', 'leisure'],
    is_default: 0,
    sort: 3
  },
  {
    id: 4,
    tag_key: 'read',
    name: 'Read',
    color: 'bg-purple-100',
    icon: '/visual/read.png',
    keywords: ['book', 'article', 'study'],
    is_default: 0,
    sort: 4
  }
]) 
const filteredTags = computed(() => {
  return Tags.value.filter(t => t.name.toLowerCase().includes(keyword.value.toLowerCase()))
})
function createNew() {
  currentTag.value = null
  selectedColor.value = 'bg-slate-100'
  selectedIcon.value = 'Default'
  customIcon.value = null
  currentTag.value = {
    id: 0,
    tag_key: '',
    name: '',
    color: 'bg-slate-100',
    icon: '/visual/default.png',
    keywords: [],
    is_default: 0,
    sort: 0
  }
}

function editTag(tag) {
  if (currentTag.value == null) return
  tag.color = selectedColor.value
  tag.icon = customIcon.value || selectedIconComponent.value || '/visual/default.png'
  if (tag.id == 0) {
    tag.Lang = i18n.global.locale.value
    tag.tag_key = tag.name.toLowerCase().replace(/\s+/g, '-') 
    tag.id = Date.now() + Math.random()
    Tags.value.unshift(tag)
    currentTag.value = null 
  } else {
      const index = Tags.value.findIndex(t => t.id === tag.id)
      if (index !== -1) {
        Tags.value[index] = tag
      }
      currentTag.value = null
  }
}

function deleteTag() {
  if (currentTag.value == null) return
  Modal.confirm({
    title: `Do you want to delete the tag "${currentTag.value?.name}"?`,
    icon: createVNode(ExclamationCircleOutlined),
    content: 'When clicked the OK button, this dialog will be closed after 1 second',
    okText: 'Yes',
    okType: 'danger',
    cancelText: 'No',
    async onOk() {
      try {
        Tags.value = Tags.value.filter(t => t.id !== currentTag.value?.id)
        message.success('Tag deleted successfully.')
        currentTag.value = null
      } catch (error) {
        message.error('Failed to delete the tag.')
      }
    },
    onCancel() { },
  });
}
// 选择
function select(p: Tag) {
  selectedColor.value = p.color
  customIcon.value = p.icon
  currentTag.value = JSON.parse(JSON.stringify(p))
}
function fetchTags() {
  // 调用 API 获取标签列表 
}
onMounted(async () => {
  await fetchTags()
})

const selectedColor = ref('bg-slate-100') // 默认颜色
const colors = [
  'bg-transparent',
  'bg-slate-100',
  'bg-gray-100',
  'bg-red-100',
  'bg-yellow-100',
  'bg-green-100',
  'bg-blue-100',
  'bg-indigo-100',
  'bg-purple-100',
  'bg-slate-400',
  'bg-red-400',
  'bg-yellow-400',
  'bg-green-400',
  'bg-blue-400',
  'bg-indigo-400',
  'bg-purple-400',
  'bg-red-700',
  'bg-yellow-700',
  'bg-green-700',
  'bg-zinc-800'
]
// icon 列表 
const iconOptions = [
  { name: 'Default', component: '/visual/default.png' },
  { name: 'Work', component: '/visual/work.png' },
  { name: 'Life', component: '/visual/eat.png' },
  { name: 'Rest', component: '/visual/coffee.png' },
  { name: 'read', component: '/visual/read.png' },
]

const selectedIcon = ref('Default')
const customIcon = ref<string | null>(null)
const selectedIconComponent = computed(() => {
  return iconOptions.find(i => i.name === selectedIcon.value)?.component
})

function selectIcon(name: string) {
  selectedIcon.value = name
  customIcon.value = null
}
// 上传处理
async function handleUpload(e: Event) {
  console.log('上传事件:')
} 

// 子标签
const newSubTag = ref('')
// 添加
function addSubTag() {
  const val = newSubTag.value.trim()
  if (!val) return
  if (currentTag.value?.keywords.includes(val)) return
  currentTag.value?.keywords.push(val) 
  newSubTag.value = ''
}
// 删除
function removeSubTag(index: number) {
  currentTag.value?.keywords.splice(index, 1)
}
</script>
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

.item {
  @apply flex p-2 hover:bg-gray-100 cursor-pointer rounded dark:hover:bg-slate-700 dark:border-gray-700;
}

.selected {
  @apply border-l-4 border-indigo-500 bg-white dark:bg-slate-500 dark:border-slate-600;
}

.tag {
  @apply px-2 py-1 text-xs bg-gray-100 rounded mr-2;
}

.tag.active {
  @apply bg-purple-100 text-purple-600;
}
</style>