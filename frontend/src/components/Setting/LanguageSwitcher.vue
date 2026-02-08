<template>
  <div >
    <a-select  class="dark:text-white"
              ref="select"
              v-model:value="locale"
              style="width: 100px"
              @change="switchLang"
            >
              <a-select-option value="zh" >简体中文</a-select-option>
              <a-select-option value="en">English</a-select-option>
              <a-select-option value="zhhk">繁體中文</a-select-option>
            </a-select>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { i18n, setupI18n } from '../../utils/i18n'
import { Locale } from '../../locales'

const locale = ref(i18n.global.locale.value)

const switchLang =async  () => {
  await setupI18n(locale.value as Locale)
};
</script>