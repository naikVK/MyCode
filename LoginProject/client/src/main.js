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

window.axios = axios
// window.axios.defaults.baseURL = window.location.origin

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
axios.defaults.baseURL = 'http://10.4.0.34:8080/server/'
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
  // console.log(config)
  // var username = store.getters.getUsername
  // if (username != null) {
  //   var f = username.charAt(0)
  //   console.log(f)
  //   if (f.length === 1 && f.match(/[a-z]/i)) {
  //     config.baseURL = 'http://10.4.2.64:3030/'
  //   }
  // }
  // Do something before request is sent
  var tokenString = store.getters.getSessionToken
  if (tokenString != null) {
    config.headers.Authorization = tokenString
  }
  var clientId = store.getters.getClientId
  if (clientId != null) {
    config.headers.ClientId = clientId
  }
  var RestrictToken = store.getters.getRestrict_token
  if (clientId != null) {
    config.headers.Authorization = RestrictToken
  }
  return config
}, function (error) {
  // Do something with request error
  console.log('intercept error')
  return Promise.reject(error)
})
