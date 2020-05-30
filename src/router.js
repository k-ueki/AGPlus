import Vue from 'vue'
import Router from 'vue-router'
import Index from './components/Index'
import ClassList from './pages/ClassList'
import FacultyDepartment from './pages/FacultyDepartment'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    // {
    //   path: '/',
    //   name: 'index',
    //   component: Index
    // },
    {
      path: '/welcome',
      name: 'index',
      component: Index
    },
    {
      path: '/classes',
      name: 'classList',
      component: ClassList
    },
    {
      path: '/faculties',
      name: 'faculties',
      component: FacultyDepartment
    },
  ]
})