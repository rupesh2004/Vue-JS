import { createRouter, createWebHistory } from 'vue-router'

import Register from '@/components/Register.vue'
import Login from '@/components/Login.vue'
import UserDashboard from '@/components/UserDashboard.vue'
import AdminDashboard from '@/components/AdminDashboard.vue'
import UserDetails from '@/components/UserDetails.vue'
import ManageUsers from '@/components/ManageUsers.vue'
import ViewAllUsers from '@/components/ViewAllUsers.vue'
import UpdateUser from '@/components/UpdateUser.vue'
import CreateExam from '@/components/CreateExam.vue'
import ViewExamsVue from '@/components/ViewExams.vue'
import ExamRegistration from '@/components/ExamRegistration.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Login',
      component: Login,
    },
    {
      path: '/addUser',
      name: 'Register',
      component: Register,
    },
    {
      path: '/userDashboard',
      name: 'UserDashboard',
      component: UserDashboard,
    },
    {
      path: '/adminDashboard',
      name: 'AdminDashboard',
      component: AdminDashboard,
    },
    {
      path: '/userDetails',
      name: 'UserDetails',
      component: UserDetails,
      props:true  
    },
    {
      path: '/userManagement',
      name: 'ManageUsers',
      component: ManageUsers,
      props:true  
    },
    {
      path: '/viewAllUsers',
      name: 'ViewAllUsers',
      component: ViewAllUsers,
      props:true  
    },
    {
      path: '/updateUser',
      name: 'UpdateUser',
      component: UpdateUser,
    },
    {
      path: '/createExam',
      name: 'CreateExam',
      component: CreateExam,
    },
    {
      path: '/viewExam',
      name: 'ViewExam',
      component: ViewExamsVue,
    },
    {
      path: '/examRegistration',
      name: 'ExamRegistration',
      component: ExamRegistration,
    },
    
  ],
})

export default router
