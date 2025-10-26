import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  { path: '/', redirect: '/products' },
  { path: '/products', component: () => import('@/views/ProductView.vue') },
  { path: '/pricetags', component: () => import('@/views/PriceTagView.vue') },
  {
    path: '/pricetags-templates',
    component: () => import('@/views/TemplateView.vue'),
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
