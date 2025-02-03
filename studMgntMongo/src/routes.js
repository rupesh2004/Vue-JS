import App from "@/App.vue";
import AddStudent from "@/components/AddStudent.vue";
import Dashboard from "@/components/Dashboard.vue";
import FetchAllStudent from "@/components/FetchAllStudent.vue";
import { compile, name } from "ejs";
import {createWebHistory, createRouter} from "vue-router";
const routes = [
    {
        name : "Dashboard",
        path: "/",
        component: Dashboard,
    },
    {
        name :"AddStudent",
        path: "/addStudent",
        component : AddStudent
    },
    {
        name:"FetchAllStudent",
        path: "/fetchAllStudent",
        component:FetchAllStudent
    }
]
const router = createRouter({
    history: createWebHistory(),
    routes: routes
})
export default router