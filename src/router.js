import Vue from 'vue'
import Router from 'vue-router'
import Index from './pages/Index'
import ClassList from './pages/ClassList'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'index',
      component: Index
    },
    {
      path: '/class-list',
      name: 'classList',
      component: ClassList
    }
  ]
})