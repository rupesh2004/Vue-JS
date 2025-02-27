import {createApp} from 'vue'
import {createPinia} from 'pinia'
import App from './App.vue'
import piniaPersist from 'pinia-plugin-persistedstate';

const app = createApp(App)
const pinia = createPinia()

pinia.use(piniaPersist)
app.use(pinia)
app.mount('#app')
export default pinia