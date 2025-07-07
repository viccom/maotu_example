import { createRouter, createWebHistory } from 'vue-router';
import Edit from '@/views/edit.vue';
import Preview from '@/views/preview.vue';
import Test from '@/views/test.vue';


const routes = [
  { path: '/', redirect: '/edit' },
  { path: '/test', name: 'test', component: Test },
  { path: '/edit', name: 'edit', component: Edit },
  { path: '/preview', name: 'preview', component: Preview }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
