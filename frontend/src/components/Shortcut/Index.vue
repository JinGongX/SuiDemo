<script setup lang="ts">
import { ref,watch ,onMounted,nextTick} from 'vue'
import ListItem from '../Setting/ListRow.vue' 
import ShortcutInput from '../Setting/ShortcutInput.vue';
import WinShortcutInput from '../Setting/WinShortcutInput.vue';
import { parseShortcutToHotkeyWin,parseShortcutToHotkey,formatHotkeyStringmac ,formatHotkeyStringWin} from '../../utils/hotkeyUtils'; // 🔁 引入工具函数

import { UpHotkey,GetHotkeys} from '../../../bindings/changeme/internal/services/hotkeyService';
import {  IsmacOS } from '../../utils/osinfo'
import { message } from 'ant-design-vue';
const ismacos=ref(false)
// const modelValue = ref(false)
const OpenShortcut = ref('');
const OpenSetting = ref('');

const isInitialized = ref(false); 

let lastSaved = '';
//const action = 'OpenSearch';
// 监听变化自动保存
watch(OpenShortcut, async (newShortcut) => {
   if (!isInitialized.value) return;

  if (!newShortcut || newShortcut === lastSaved) return;

  try {
   if(ismacos.value){
      await SendHanld(1,newShortcut); 
    }else{
      await SendHanldWin(1,newShortcut); 
    }
  } catch (e: any) {
    message.success('快捷键保存失败: ' + e.message); 
  }
},{ flush: 'post' }
);


watch(OpenSetting, async (newShortcut) => {
   if (!isInitialized.value) return;

  if (!newShortcut || newShortcut === lastSaved) return;
  try {
   if(ismacos.value){
      await SendHanld(2,newShortcut); 
    }else{
      await SendHanldWin(2,newShortcut); 
    }
  } catch (e: any) {
    message.success('快捷键保存失败: ' + e.message);
  }
},{ flush: 'post' }
 );


const SendHanld = async (item,newShortcut) => {
   try {
     const parsed = parseShortcutToHotkey(newShortcut);
     if (!parsed) {
       console.error('❌ 快捷键格式错误:', newShortcut);
       return;
     }
    message.success('快捷键已保存'); 
    await UpHotkey(item,parsed.key, parsed.modifier);
  } catch (e: any) {
    message.success('快捷键保存失败: ' + e.message);
  }
};

const SendHanldWin = async (item,newShortcut) => {
   try {
     const parsed = parseShortcutToHotkeyWin(newShortcut);
     if (!parsed) {
       console.error('❌ 快捷键格式错误:', newShortcut);
       return;
     }
    message.success('快捷键已保存'); 
    await UpHotkey(item,parsed.key, parsed.modifier);
  } catch (e: any) {
    message.success('快捷键保存失败: ' + e.message);
  }
};

interface HotkeyItem {
  id: number;
  keycode: number;
  modifiers: number;
}
const hotkeyentry = ref<HotkeyItem[]>([]);

const Gethotkey = async () => {
  hotkeyentry.value = await GetHotkeys(); 
  if (hotkeyentry.value && hotkeyentry.value.length > 0) { 
    if(ismacos.value){
       OpenShortcut.value = formatHotkeyStringmac(hotkeyentry.value[0].keycode, hotkeyentry.value[0].modifiers);
       OpenSetting.value = formatHotkeyStringmac(hotkeyentry.value[1].keycode, hotkeyentry.value[1].modifiers);
    }
    else{
      OpenShortcut.value = formatHotkeyStringWin(hotkeyentry.value[0].keycode, hotkeyentry.value[0].modifiers);
      OpenSetting.value = formatHotkeyStringWin(hotkeyentry.value[1].keycode, hotkeyentry.value[1].modifiers);
    }
    await nextTick();
    isInitialized.value = true;
  }
};


 onMounted(async() => {
    ismacos.value=IsmacOS()
    Gethotkey(); 
     // 等下一轮 DOM 渲染完成后才启用监听，避免初始赋值触发 watch
 });
 

</script>

<template> 
  <h2 class="text-lg font-bold text-gray-800 dark:text-white" >{{ $t('components.general.title.shortcut') }}</h2>
  <div> 
       <div class="bg-white rounded-lg shadow divide-y divide-gray-200 dark:bg-gray-800 dark:divide-gray-700">
         <ListItem :label="$t('components.general.label.open_second')" :subLabel="$t('components.general.subLabel.ht_shortcut')">
             <ShortcutInput v-if="ismacos" v-model:modelValue="OpenShortcut"  />
            <WinShortcutInput v-else v-model:modelValue="OpenShortcut" />
          </ListItem>
          <ListItem :label="$t('components.general.label.open_setting')" subLabel="">
              <ShortcutInput v-if="ismacos" v-model:modelValue="OpenSetting" />
             <WinShortcutInput v-else v-model:modelValue="OpenSetting" />
          </ListItem> 
      </div>
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