<template>
  <div class="w-full h-full flex items-center justify-center ">
    <!-- 中央轻卡片 -->
    <div
      class="
        w-[420px]
        rounded-2xl
        bg-neutral-50 dark:bg-gray-800
        px-8 py-10
        text-center
        shadow-sm
        border border-neutral-100 dark:border-gray-700
      "
    >
      <!-- 轻符号 / 视觉锚点 -->
      <transition name="fade-up">
        <div
          v-if="show"
          class="flex justify-center mb-4"
        >
          <div
            class="
              w-10 h-10
              rounded-full
              bg-orange-100
              text-orange-500
              flex items-center justify-center
              text-lg 
            "
          >
            •
          </div>
        </div>
      </transition>
      <!-- 标题 -->
      <transition name="fade-up">
        <h1
          v-if="show"
          class="text-2xl font-semibold text-neutral-800 mb-1 dark:text-white"
          style="transition-delay: 80ms"
        >
          SuiDemo
        </h1>
      </transition>
      <!-- 副标题 -->
      <transition name="fade-up">
        <p
          v-if="show"
          class="text-sm text-neutral-500 mb-6 dark:text-neutral-300"
          style="transition-delay: 120ms"
        >
          Wails3 · 桌面工具
        </p>
      </transition>
      <!-- 主文案 -->
      <transition name="fade-up">
        <div
          v-if="show"
          class="text-sm text-neutral-700 leading-relaxed space-y-3 mb-6 dark:text-white"
          style="transition-delay: 180ms"
        >
          <p></p>
          <p></p>
          <div class="h-2"></div>
          <p>只是在合适的时候，</p>
          <p class="font-medium">
            让复杂的事物变得简单一些。
          </p>
        </div>
      </transition>
      <!-- 分割线 -->
      <transition name="fade-up">
        <div
          v-if="show"
          class="flex justify-center mb-6"
          style="transition-delay: 240ms"
        >
          <span class="w-12 h-px bg-neutral-200"></span>
        </div>
      </transition>
      <!-- 底部信息 -->
      <transition name="fade-up">
        <div
          v-if="show"
          class="text-xs text-neutral-400 space-y-1 dark:text-white"
          style="transition-delay: 300ms"
        >
          <p>Designed & Built by Jin Gong</p>
          <p>© 2026 Light Prompt</p>
        </div>
      </transition>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
const show = ref(false)
onMounted(() => {
  // 延迟出现，避免“页面一闪全出”
  setTimeout(() => {
    show.value = true
  }, 80)
})
</script>

<style scoped>
.fade-up-enter-from {
  opacity: 0;
  transform: translateY(6px);
}

.fade-up-enter-to {
  opacity: 1;
  transform: translateY(0);
}

.fade-up-enter-active {
  transition: all 0.4s ease;
}
</style>