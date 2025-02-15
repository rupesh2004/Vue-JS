import { createRouter, createWebHistory } from 'vue-router'
import Login from '@/components/Login.vue'
import Register from '@/components/Register.vue'
import Dashboard from '@/components/Dashboard.vue'
import Admin from '@/components/Admin.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Login',
      component: Login,
    },
    {
      path: '/register',
      name: 'Register',
      component : Register,
    },
    {
      path: '/dashboard',
      name: 'Dashboard',
      component : Dashboard,
      beforeEnter:(to,from,next)=>{
        const token  = sessionStorage.getItem("token")
        if(!token){
          next({name:'Login'})
        }else{
          next()
        }
      }
    },
    {
      name:"Admin",
      path:"/admin",
      component: Admin,
    }
  ],
})

export default router
