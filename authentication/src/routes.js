import {createWebHistory, createRouter} from "vue-router"
import LoginForm from './components/loginForm.vue'
import SignUpForm from "./components/signupForm.vue"
import NotFound from "./components/pageNotFound.vue"
const routes = [
    {
        name : "LoginForm",
        path :"/",
        component : LoginForm,
        
    },
    {
        name : "SignUpForm",
        path : "/signup",
        component : SignUpForm
    },
    {
        name:"NotFound",
        path:"/:pathMatch(.*)*",
        component:NotFound
    }
]
const router = createRouter({
    history:createWebHistory(),
    routes
});
export default router