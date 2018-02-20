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
  getIsLoggedIn: function (state) {
    return state.isLoggedIn
  }
}
