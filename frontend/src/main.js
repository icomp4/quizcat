// main.js
import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'
import './app.css'

const routes = [
  { path: '/', component: App },
]

const router = createRouter({
  history: createWebHistory(),
  routes 
})

const app = createApp(App)
app.use(router)
app.mount('#app')