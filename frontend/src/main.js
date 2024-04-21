// main.js
import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import Main from './Main.vue'
import App from './App.vue'
import './app.css'
import TrendingPage from './TrendingPage.vue'
import QuizPlayPage from './QuizPlayPage.vue'
import SignupPage from './SignupPage.vue'
import LoginPage from './LoginPage.vue'


const routes = [
  { path: '/', component: Main },
  { path: "/trending", component: TrendingPage },
  { path: "/quiz/:id", component: QuizPlayPage, props: true},
  { path: "/signup", component: SignupPage },
  { path: "/login", component: LoginPage}
]

const router = createRouter({
  history: createWebHistory(),
  routes 
})

const app = createApp(App)
app.use(router)
app.mount('#app')