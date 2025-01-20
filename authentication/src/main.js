import { createApp } from 'vue';
import { createPinia } from 'pinia';
import App from './App.vue';
import router from './routes.js';
import piniaPersist from 'pinia-plugin-persistedstate';

const app = createApp(App);

// Create Pinia instance and use the persist plugin
const pinia = createPinia();
pinia.use(piniaPersist);  // Apply the persist plugin

app.use(pinia);  // Use pinia
app.use(router);  // Use router
app.mount('#app');  // Mount app

export default pinia;
