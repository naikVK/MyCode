// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import store from '../store/index.js'
import axios from 'axios'
import VeeValidate from 'vee-validate'
import vuexI18n from 'vuex-i18n'
import Toasted from 'vue-toasted'

import translationsEn from '../static/localization/en.json'
import translationsMr from '../static/localization/mr.json'
// import {getToken} from './utils/utils'
// import * as mkclApi from '../static/mkclLogin.js'
import MKCLAPI from '../static/mkclLogin.js'

// window.mkclApi = mkclApi
window.MKCLAPI = MKCLAPI
window.MKCLAPI.init({
  clientId: 'SOLAR_2',
  mkclLoginServerURL: 'http://10.4.0.34:3030',
  mkclLoginClientURL: 'http://10.4.0.34:8080/#/login'
})

window.axios = axios

Vue.use(VeeValidate, {
  events: 'blur'
})
Vue.use(vuexI18n.plugin, store)
Vue.use(Toasted, {
  duration: 1000,
  theme: 'bubble'
})

Vue.i18n.add('en', translationsEn)
Vue.i18n.add('mr', translationsMr)

// set the start locale to use
Vue.i18n.set('en')

Vue.config.productionTip = false
// axios.defaults.withCredentials = true
// axios.defaults.baseURL = 'http://localhost:8080/server/'
// window.serverPath = axios.defaults.baseURL
/* eslint-disable no-new */

new Vue({
  el: '#app',
  store,
  router,
  template: '<App/>',
  components: { App }
})

axios.interceptors.request.use(function (config) {
  // Do something before request is sent
  var clientId = store.getters.getClientId
  if (clientId != null) {
    config.headers.ClientId = clientId
  }
  // alert(clientId)
  var tokenString = window.MKCLAPI.getToken()
  // alert(tokenString)
  // alert(tokenString)
  if (tokenString != null) {
    config.headers.Authorization = tokenString
  }
  return config
}, function (error) {
  // Do something with request error
  console.log('intercept error')
  return Promise.reject(error)
})

router.beforeEach(function (to, from, next) {
  if (to.matched.some(record => record.meta.requireAuth)) {
    window.MKCLAPI.isSessionPresent().then(function () {
      next()
    }).catch(function () {
      router.push({ path: '/login' })
    })
  } else {
    next()
  }
})
