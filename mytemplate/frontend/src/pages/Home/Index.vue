<script setup lang="ts">
import { ref } from 'vue'
import { applyTheme } from '../../utils/ThemeManager'
  const bc = new BroadcastChannel('theme')
  bc.onmessage = (e) => {
    applyTheme(e.data)
  }
</script>

<template> 
  <div class="flex  pt-4 h-screen w-screen h-full w-full bg-gray-100/90 drag-region dark:bg-gray-900"> 
      <div class="pt-4 pl-4 pr-4 dark:text-white">
        <div>
          <h2>Second</h2>
          <p>Click on the Wails logo to learn more</p></div>
      </div>
  </div>
</template>
