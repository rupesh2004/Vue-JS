import { createRouter, createWebHistory } from 'vue-router'

import Login from '@/components/Login.vue'
import UserDashboard from '@/components/UserDashboard.vue'
import AdminDashboard from '@/components/AdminDashboard.vue'
import UserDetails from '@/components/UserDetails.vue'
import ManageUsers from '@/components/ManageUsers.vue'
import UpdateUser from '@/components/UpdateUser.vue'
import CreateExam from '@/components/CreateExam.vue'
import ExamRegistration from '@/components/ExamRegistration.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Login',
      component: Login,
      meta:{
        authRequired: false,
        showHeader: true,
        showFooter: true,
      }
    },
    {
      path: '/userDashboard',
      name: 'UserDashboard',
      component: UserDashboard,
      meta: {
        authRequired: true,
        role:['user'],
        showHeader: true,
        showFooter: true,
      }
    },
    {
      path: '/adminDashboard',
      name: 'AdminDashboard',
      component: AdminDashboard,
      meta: {
        authRequired: true,
        role: ['admin'],
        showHeader: true,
        showFooter: true,
      },
    },
    {
      path: '/userDetails',
      name: 'UserDetails',
      component: UserDetails,
      props: true  
    },
    {
      path: '/userManagement',
      name: 'ManageUsers',
      component: ManageUsers,
      props: true  
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
      path: '/examRegistration',
      name: 'ExamRegistration',
      component: ExamRegistration,
    },
  ],
})

export default router;

