import {createRouter, createWebHistory} from 'vue-router';
import Sender from './components/sender.vue';
import Receiver from './components/Receiver.vue';

const routes = [
  { path: '/', component: Sender },
  { path: '/second-page', name: 'secondPage', component: Receiver }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
