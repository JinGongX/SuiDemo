import { createApp } from 'vue'
import Antd from 'ant-design-vue';//add
import App from './App.vue'
import './style.css' 
import 'ant-design-vue/dist/reset.css';
import {i18n,setupI18n } from './utils/i18n'
import { initTheme } from './utils/ThemeManager'//add
import router from './router/index'

initTheme()
if (window.location.hash === '' || window.location.hash === '#/') {
  router.replace('/second')
}
async function bootstrap(){
    await setupI18n('zh')//默认语言或从后端读取
    createApp(App).use(Antd).use(router).use(i18n).mount('#app')
}
bootstrap()
