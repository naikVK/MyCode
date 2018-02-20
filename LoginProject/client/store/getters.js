import { getCookie } from '../src/utils/utils.js'

export default {
  getToken: function (state) {
    // console.log('Inside getter ::getToken')
    return state.token
  },
  getConfig: function (state) {
    // console.log('Inside getter ::getConfig')
    return state.config
  },
  getClientId: function (state) {
    // console.log('Inside getter ::getClientId')
    return state.clientId
  },
  getRestrict_token: function (state) {
    // console.log('Inside getter ::getClientId')
    return state.Restrict_token
  },
  getIsLoggedIn: function (state) {
    var sid = state.clientid
    var token = getCookie(sid)
    if (token === undefined || token === null || token === '' || token.length === 0) {
      return false
    }
    return true
  },
  getSessionToken : function(state) {
    if (state.config != null) {
      var sid = state.config.clientid
      return getCookie(sid)
    }
    return null
  },
  getUsername: function (state) {
    return state.username
  },
  getShowQRcode: function (state) {
    return state.showqrcodeimage
  }
}
