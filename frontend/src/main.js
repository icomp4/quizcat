// main.js
import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import Main from './Main.vue'
import QuizComponent from './Quiz.vue'
import App from './App.vue'
import './app.css'

const routes = [
  { path: '/', component: Main },
  { path: '/quiz/:id', component: QuizComponent, props: true }
]

const router = createRouter({
  history: createWebHistory(),
  routes 
})

const app = createApp(App)
app.use(router)
app.mount('#app')