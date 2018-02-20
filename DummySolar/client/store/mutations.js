export default {
  SET_TOKEN (state, token) {
    console.log('setting token')
    state.token = token
    state.isLoggedIn = true
  },
  REMOVE_TOKEN (state) {
    // console.log('removing token')
    state.token = null
  },
  SET_CONFIG (state, config) {
    // console.log('setting config')
    state.config = config
  },
  REMOVE_CONFIG (state) {
    state.config = null
    // console.log("CONFIG CLEAR")
  },
  FLUSH_DATA (state) {
    // console.log('flushing all data')
    state.token = null
    state.config = null
    state.isLoggedIn = false
    // console.log("data flush")
  },
  SET_IS_LOGGEDIN (state) {
    // console.log('Setting is logged in')
    state.isLoggedIn = true
  }
}
