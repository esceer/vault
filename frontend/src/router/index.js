import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import CreateView from '../views/CreateView.vue'
import ListView from '../views/ListView.vue'
import SearchView from '../views/SearchView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/list',
      name: 'list',
      component: ListView
    },
    {
      path: '/create',
      name: 'create',
      component: CreateView
    },
    {
      path: '/search',
      name: 'search',
      component: SearchView
    },
    {
      path: '/security',
      name: 'security',
      // lazy-loaded when the route is visited.
      component: () => import('../views/SecurityView.vue')
    }
  ]
})

export default router
