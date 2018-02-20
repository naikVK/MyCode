import { setSession, deleteSession } from '../src/utils/utils.js'

export default {
  SET_TOKEN (state, token) {
    var sid = state.config.clientid
    state.isLoggedIn = true
    setSession(sid, token)
  },
  REMOVE_TOKEN (state) {
    // console.log('removing token')
    state.token = null
    var sid = state.config.clientid
    deleteSession(sid)
  },
  SET_CONFIG (state, config) {
    // console.log('setting config')
    state.config = config
  },
  REMOVE_CONFIG (state) {
    state.config = null
    // console.log("CONFIG CLEAR")
  },
  SET_USERNAME (state, username) {
    // console.log('setting username')
    state.username = username
      // console.log(username)
  },
  SET_SHOWQRCODE (state, showqrcodeimage) {
    // console.log('setting showqrcodeimage')
    state.showqrcodeimage = showqrcodeimage
      // console.log(showqrcodeimage)
  },
  SET_RESTRICT_TOKEN (state, Restrict_token) {
    // console.log('setting showqrcodeimage')
    state.Restrict_token = Restrict_token
      // console.log(Restrict_token)
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
  }, 
  SET_CLIENT_ID (state, cId) {
    state.clientId = cId
  }
}
