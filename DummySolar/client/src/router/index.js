import Vue from 'vue'
import Router from 'vue-router'
import LoginUserName from '@/components/LoginUserName'
import Home from '@/components/home'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      redirect: '/login'
      // component: LogetTinUserName
    },
    {
      path: '/login',
      name: 'LoginUserName',
      component: LoginUserName,
      meta: {requireAuth: false}
    },
    {
      path: '/home',
      name: 'Home',
      component: Home,
      meta: {requireAuth: true}
    }
  ]
})
