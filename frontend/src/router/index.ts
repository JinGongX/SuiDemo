import { createRouter, createWebHashHistory} from 'vue-router'
import MainPage from '../components/Main/Index.vue'
import SecondPage from '../components/About/Index.vue'

const routes = [
  { path: '/', component: MainPage },
  { path: '/second', component: SecondPage }
]

export default createRouter({
  history: createWebHashHistory(), // Use createWebHashHistory for hash-based routing
  routes
})