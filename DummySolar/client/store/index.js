import Vue from 'vue'
import Vuex from 'vuex'
import state from './state'
import getters from './getters'
import mutations from './mutations'
import createPersistedState from 'vuex-persistedstate'

Vue.use(Vuex)
var myplugins = [];

(function () {
  if(window.navigator.cookieEnabled) {
    // alert('cookies enabled')
    myplugins.push(createPersistedState({ storage: window.sessionStorage }))
  }
})()

export default new Vuex.Store({
  state,
  mutations,
  getters,
  // plugins: [createPersistedState({ storage: window.sessionStorage })]
  plugins: myplugins
})
