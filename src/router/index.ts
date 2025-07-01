import { createRouter, createWebHistory } from 'vue-router';
import Edit from '@/views/edit.vue';
import Preview from '@/views/preview.vue';

const routes = [
  { path: '/', redirect: '/edit' },
  { path: '/edit', name: 'edit', component: Edit },
  { path: '/preview', name: 'preview', component: Preview }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
