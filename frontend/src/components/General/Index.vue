<script setup lang="ts">
import { ref,h } from 'vue'
import ListItem from '../Setting/ListRow.vue' 
import LanguageSwitcher from '../Setting/LanguageSwitcher.vue'
import ThemeSetting from '../Setting/ThemeSetting.vue'

const modelValue = ref(false)
</script>

<template > 
<div >
  <section >

      <!-- 应用设置 -->
  <h2 class="text-lg font-bold text-gray-800 dark:text-white" >{{ $t('components.general.title.application') }}</h2>
      <div class="bg-white rounded-lg shadow divide-y divide-gray-200 dark:bg-gray-800 dark:divide-gray-700">
          <ListItem :label="$t('components.general.label.startup')" subLabel="">
            <input type="checkbox" class="sr-only peer" v-model="modelValue" />
            <div
              class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-blue-500 rounded-full peer dark:bg-gray-300 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-500 relative"
            ></div>
          </ListItem>
      </div>
    <!-- 外观设置 -->
        <h2 class="text-lg font-bold text-gray-800 dark:text-white"  >{{ $t('components.general.title.exterior') }}</h2>
      <div class="bg-white rounded-lg shadow divide-y divide-gray-200 dark:bg-gray-800 dark:divide-gray-700">
         <ListItem :label="$t('components.general.label.language')" subLabel="">
            <LanguageSwitcher />
          </ListItem>
          <ListItem :label="$t('components.general.label.theme')" subLabel="">
            <ThemeSetting />
          </ListItem>
      </div>
<!-- 应用更新 -->
     <h2 class="text-lg font-bold text-gray-800 dark:text-white" >{{ $t('components.general.title.update') }}</h2>
      <div class="mt-4 bg-white rounded-lg shadow divide-y divide-gray-200 dark:bg-gray-800 dark:divide-gray-700">
         <ListItem :label="$t('components.general.label.automatic_up')" subLabel="">
            <input type="checkbox" class="sr-only peer" v-model="modelValue" />
            <div
              class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-blue-500 rounded-full peer dark:bg-gray-300 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-500 relative"
            ></div>
          </ListItem>
          <ListItem :label="$t('components.general.label.next_up')" subLabel="">
            <input type="checkbox" class="sr-only peer" v-model="modelValue" />
            <div
              class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-blue-500 rounded-full peer dark:bg-gray-300 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-500 relative"
            ></div>
          </ListItem>
      </div> 


  </section>
  </div>
</template>
<style scoped>
 h2{
    text-align: left;
    margin:15px 15px 4px 15px;
    font-size:15px;
 }
 .rg_desc{
    text-align: left;
    margin:4px 0px 4px 8px;
    font-size:13px;
    color:#999;
 }
</style>