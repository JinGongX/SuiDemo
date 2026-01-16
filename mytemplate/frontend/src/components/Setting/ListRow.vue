<script setup lang="ts">
defineProps<{
   label: string
   subLabel?: string
 }>()
</script>

<template> 
   <div class="flex items-center justify-between px-4 py-3 ">
      <div style="text-align: left;">
        <div class="text-sm text-gray-900 font-medium dark:text-white">{{ label }}</div>
        <div v-if="subLabel" class="text-xs text-gray-500 mt-1">{{ subLabel }}</div>
      </div>
      <label class="inline-flex items-center cursor-pointer">
        <slot />
      </label>
    </div> 
</template>
