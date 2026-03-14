<script setup lang="ts">
import { applyTheme } from '../../utils/ThemeManager'
  const bc = new BroadcastChannel('theme')
  bc.onmessage = (e) => {
    applyTheme(e.data)
  }

interface PromptCard {
  id: number
  type: string //PromptType
  title: string
}

// 根据类型获取样式
function getVisual(type: string) {
  const map: Record<string, { bg: string; image: string }> = {
    life: { bg: 'bg-amber-50', image: '/wails.png' },
    work: { bg: 'bg-blue-50', image: '/wails.png' },
    device: { bg: 'bg-teal-50', image: '/wails.png' },
    security: { bg: 'bg-emerald-50', image: '/wails.png' },
    system: { bg: 'bg-neutral-100', image: '/wails.png' },
    default: { bg: 'bg-neutral-100', image: '/wails.png' },
  }
  return map[type] || map.default
} 

// ---------- 示例数据 ----------
const cards: PromptCard[] = [
  {
    id: 1,
    type: 'life',
    title: 'The kitchen is messy. Please provide a five-step tidying plan',
  },
  {
    id: 2,
    type: 'device',
    title: 'My storage space is insufficient. Please guide me on how to back up photos',
  },
  {
    id: 3,
    type: 'security',
    title: 'I want to enhance account security. Please help me set up a password manager',
  },
]

//
const onCardClick = (card) => {
  console.log('点击卡片', card)
} 
</script>

<template>
  <div class="flex pt-8 h-screen w-screen h-full w-full  font-sans select-none bg-gray-100 drag-region dark:bg-gray-900">
    <TransitionGroup name="card" tag="div">
      <h2  class="text-neutral-900 dark:text-white">Second</h2>
      <div
        v-for="card in cards"
        :key="card.id"
        @click="onCardClick(card)"
        class="group flex gap-3 mt-1 ml-2 mr-2  rounded-2xl bg-white/80  dark:bg-gray-700 dark:text-white backdrop-blur-md px-1 py-1 shadow-[0_8px_24px_rgba(0,0,0,0.12)] transition"
      >
        <!-- Visual Block -->
        <div
          class="flex h-12 w-12 shrink-0 items-center justify-center rounded-xl"
          :class="getVisual(card.type).bg"
        >
          <img
            :src="getVisual(card.type).image"
            class="h-10 w-10 object-contain"
            draggable="false"
          />
        </div>
        <!-- Content -->
        <div
          class="flex-1 text-sm py-1 text-neutral-900 line-clamp-2 break-words dark:text-white"
          :title="card.title"
        >
          {{ card.title }}
        </div>
        <!-- Close -->
        <div class="flex h-10 w-8">
          <button
            class="opacity-0 group-hover:opacity-100 text-neutral-400 transition hover:text-neutral-600"
          >
            ✓
          </button>
        </div>
      </div>
    </TransitionGroup>
  </div>
</template>

<style scoped>
.card-enter-active,
.card-leave-active {
  transition: all 0.3s ease;
}
.card-leave-to {
  opacity: 0;
  height: 0;
  margin: 0;
  padding: 0;
}
</style>