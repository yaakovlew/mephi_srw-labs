import { RouteRecordRaw } from 'vue-router';

const MainLayout = () => import('layouts/MainLayout.vue');
const ErrorPage = () => import('pages/ErrorNotFound.vue');

import AuthPageVue from 'src/pages/auth/AuthPage.vue';
import Lab3 from 'src/pages/lab3/lab-3.vue';
import lab3BVue from 'src/pages/lab3b/lab-3-b.vue';
import lab3C from 'src/pages/lab3c/lab-3-c.vue';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: MainLayout,
    children: [
      {
        path: '/lab3a',
        name: 'lab3-a',
        component: Lab3,
      },
      {
        path: '/lab3b',
        name: 'lab3-b',
        component: lab3BVue,
      },
      {
        path: '/lab3c',
        name: 'lab3-c',
        component: lab3C,
      }
    ],
  },
  {
    path: '/:catchAll(.*)*',
    component: ErrorPage,
  },
];

export default routes;
