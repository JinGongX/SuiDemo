<script setup lang="ts">
import { ref,watch ,onMounted,nextTick} from 'vue'
import ListItem from '../Setting/ListRow.vue' 
import ShortcutInput from '../Setting/ShortcutInput.vue';
import { parseShortcutToHotkey,formatHotkeyStringmac } from '../../utils/hotkeyUtils' ; // 🔁 引入工具函数

import { UpHotkey,GetHotkeys} from '../../../bindings/changeme/services/suistore'

import { message } from 'ant-design-vue';

// const modelValue = ref(false)
const OpenShortcut = ref('');
const OpenSetting = ref('');

const isInitialized = ref(false); 

let lastSaved = '';
//const action = 'OpenSearch';
// 监听变化自动保存
watch(OpenShortcut, async (newShortcut) => {
  //message.info('快捷键已更改: ' + isInitialized.value);
   if (!isInitialized.value) return;

  if (!newShortcut || newShortcut === lastSaved) return;

  try {
   await SendHanld(1,newShortcut); 
  } catch (e: any) {
    message.success('快捷键保存失败: ' + e.message);
    //console.error('快捷键保存失败:', e.message);
  }
},{ flush: 'post' }
);


watch(OpenSetting, async (newShortcut) => {
   if (!isInitialized.value) return;

  if (!newShortcut || newShortcut === lastSaved) return;
  try {
   await SendHanld(2,newShortcut); 
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

interface HotkeyItem {
  id: number;
  keycode: number;
  modifiers: number;
}
const hotkeyentry = ref<HotkeyItem[]>([]);

const Gethotkey = async () => {
  hotkeyentry.value = await GetHotkeys(); 
  if (hotkeyentry.value && hotkeyentry.value.length > 0) {
   OpenShortcut.value = formatHotkeyStringmac(hotkeyentry.value[0].keycode, hotkeyentry.value[0].modifiers);
   OpenSetting.value = formatHotkeyStringmac(hotkeyentry.value[1].keycode, hotkeyentry.value[1].modifiers);
    await nextTick();
    isInitialized.value = true;
  }
};


 onMounted(async() => {
    Gethotkey(); 
     // 等下一轮 DOM 渲染完成后才启用监听，避免初始赋值触发 watch
 });
 

</script>

<template> 
  <h2 class="text-lg font-bold text-gray-800 dark:text-white" >快捷键</h2>
  <div> 
       <div class="bg-white rounded-lg shadow divide-y divide-gray-200 dark:bg-gray-800 dark:divide-gray-700">
         <ListItem label="打开主窗口" subLabel="按下组合键，必须包含至少一个修饰键和一个主键">
            <ShortcutInput v-model:modelValue="OpenShortcut" />
          </ListItem>
          <ListItem label="打开偏好设置窗口" subLabel="">
             <ShortcutInput v-model:modelValue="OpenSetting" />
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