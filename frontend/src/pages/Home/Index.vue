<script setup lang="ts">
import { ref } from 'vue'
import { applyTheme } from '../../utils/ThemeManager'
import { RecognizeImageBase64 } from '../../../bindings/changeme/internal/services/ocrservice'
const bc = new BroadcastChannel('theme')
bc.onmessage = (e) => {
  applyTheme(e.data)
}

//-- OCR功能 ---
declare global {
  interface Window {
    go: any
  }
}

const imageUrl = ref<string>('')
const imageBase64 = ref<string>('')
const result = ref<string>('')
const loading = ref<boolean>(false)

// 文件处理
function handleFile(file: File) {
  const reader = new FileReader()
  reader.onload = () => {
    const base64 = reader.result as string
    imageUrl.value = base64
    imageBase64.value = base64.split(',')[1]
    result.value = '' // 清空旧结果
  }
  reader.readAsDataURL(file)
}

// 点击选择
function selectImage() {
  const input = document.createElement('input')
  input.type = 'file'
  input.accept = 'image/*'
  input.onchange = (e: any) => {
    const file = e.target.files[0]
    if (file) handleFile(file)
  }
  input.click()
}

// 拖拽
function onDrop(e: DragEvent) {
  const file = e.dataTransfer?.files[0]
  if (file) handleFile(file)
}

// OCR 调用
async function handleOCR() {
  if (!imageBase64.value) return
  loading.value = true
  result.value = ''
  try {
    const res = await RecognizeImageBase64(imageBase64.value)
    result.value = res || '未识别到内容'

  } catch (err) {
    console.error(err)
    result.value = '识别失败'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div
    class="flex  items-start justify-center pt-8 h-screen w-screen h-full w-full  font-sans select-none bg-gray-100 drag-region dark:bg-gray-900">
    <TransitionGroup name="card" tag="div"> 
      <div
        class="flex-1  w-[300px] rounded-2xl bg-white/80  dark:bg-gray-700 dark:text-white backdrop-blur-md px-4 py-2 shadow-[0_8px_24px_rgba(0,0,0,0.12)] transition">
        <!-- 图片区域 -->
        <div
          class="h-20 w-full  border-2 border-dashed border-gray-300 rounded-xl flex items-center justify-center cursor-pointer overflow-hidden hover:border-blue-400 transition"
          @click="selectImage" @drop.prevent="onDrop" @dragover.prevent>
          <img v-if="imageUrl" :src="imageUrl" class="w-full h-full object-cover" />
          <span v-else class="text-gray-400 text-sm">
            拖拽图片
          </span>
        </div>
        <button @click="handleOCR" :disabled="!imageBase64 || loading" class="w-full my-2 py-1 rounded-lg text-white transition
        bg-blue-500 hover:bg-blue-600
        disabled:bg-gray-300 disabled:cursor-not-allowed">
          <span v-if="loading">识别中...</span>
          <span v-else>OCR识别</span>
        </button>
        <div v-if="result" class="bg-gray-50 rounded-lg p-3 text-sm text-gray-700 max-h-40 overflow-auto dark:bg-gray-800 dark:text-gray-300">
          <div class="font-medium mb-1">识别结果：</div>
          <pre class="whitespace-pre-wrap">{{ result }}</pre>
        </div>
      </div>
    </TransitionGroup>
  </div>

</template>

<style scoped>
</style>