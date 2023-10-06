import { createRouter, createWebHistory } from 'vue-router';
import Error404 from '/static/template/error404.vue';

const routes = [
  {
    path: '/:catchAll(.*)', // Übereinstimmung mit unbekannten Routen
    name: 'Error404',
    component: Error404, // Verweisen Sie hier auf Ihre Vue.js-Komponente
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
