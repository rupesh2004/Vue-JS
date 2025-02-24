import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/js/bootstrap.bundle.min.js";
import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import '@fortawesome/fontawesome-free/css/all.css';

const app = createApp(App);


router.beforeEach((to, from, next) => {
  if (to.name === undefined) {
    next(from.path);
    return;
  }


  const token = sessionStorage.getItem("jwt_token");
  if (
    to.meta.authRequired &&
    (token == null || token == undefined || token == "")
  ) {
    next("/");
    return;
  }

  if (to.meta.role && to.meta.role.length > 0) {
    const currentUserRole = sessionStorage.getItem("user_role");
    if(!to.meta.role.includes(currentUserRole)){
      next("/");
      return;
    }
  }
  next();
});
app.use(router);


app.mount("#app");
