var axios = require('axios')

var config = {
  clientId: null,
  mkclLoginServerURL: null,
  mkclLoginClientURL: null
}
var loginWindow
var clientConfig = null
// const MKCL'_LOGIN_URL = 'http://10.4.0.46:8080/#/login'
const LOGOUT_API = '/r/logout'
const VALIDATE_TOKEN_API = '/o/isValidToken'
const GET_CLIENT_CONFIG_API = '/o/getclientconfig'
// const sid = config.clientId

// window.addEventListener('hashchange', function (e) {
//   alert('hash change')
//   initializeLoginPage()
//   // checkAuthentication()
// })

//  function checkAuthentication() {
//   if (!isLoginPage() && !isSessionPresent()) {
//     // means user is not authorize and trying to access unAuthorized pages
//     // alert('not login page session not present')
//     redirectToLogin()
//   }
//   if (isLoginPage() && isSessionPresent()) {
//     // alert(e.oldURL)
//     // alert('login page session present')
//     // redirect(e.oldURL)
//     redirect(returnUrl)
//   }
//  }

// window.addEventListener('load', function (e) {
//   if (isReturnUrlWithToken()) {
//     var token = getParameterByName('token')
//     setSession(token)
//   } else {
//     checkAuthentication()
//   }
// })

// function initializeLoginPage () {
//   var localloginBtn = document.getElementById('mkclSignIn')
//   // if (isLoginPage() && !isSesionPresent()) {
//   if (isLoginPage()) {
//     loginBtn = localloginBtn
//     returnUrl = localloginBtn.getAttribute('data-myReturnUrl')
//     // attaching listener on sign in clicked
//     // document.cookie = 'returnUrl=' + returnUrl
//     loginBtn.addEventListener('click', loginClicked)
//   }
// }

// function loginClicked () {
//   // alert('login clicked')
//   redirectToLogin()
// }

// function redirectToLogin () {
//   // var currentUrl = window.location.href
//   var loginUrlMKCLLoginUrl = MKCLLoginUrl + '?clientId=' + clientId
//   window.location.href = loginUrlMKCLLoginUrl
// }

function login () {
  // var currentUrl = window.location.href
  var MKCLLoginUrl = config.mkclLoginClientURL + '?clientId=' + config.clientId
  window.location.href = MKCLLoginUrl
}

function loginPoc () {
  // var currentUrl = window.location.href
  var MKCLLoginUrl = config.mkclLoginClientURL + '?clientId=' + config.clientId
  loginWindow = window.open(MKCLLoginUrl, '_self')
  // window.location.href = MKCLLoginUrl
}
window.addEventListener('message', receiveMessage, false)

function receiveMessage (event) {
  // console.log(event.source.location.href)
  if (event.origin === 'http://10.4.0.54:8080') {
    alert('postmessage' + event.data)
    console.log(event)
  }
}
// function redirect (redirectTo) {
//   window.location.href = redirectTo
// }

function getClientConfig () {
  var executed = false
  return new Promise((resolve, reject) => {
    // alert('getclientconfig')
    var getClientConfigUrl = config.mkclLoginServerURL + GET_CLIENT_CONFIG_API
    if (!executed) {
      axios.post(getClientConfigUrl, {
        'clientId': clientConfig.clientId
      }).then(function () {
        resolve(true)
      }).catch(function () {
        reject(new Error('failed'))
      })
    }
  })
}
// function getClientConfig () {
//   // alert('getClientCOnfigCalled')
//   var executed = false
//   if (!executed) {
//     executed = true
//     var xmlhttp = new XMLHttpRequest()
//     var url = getClientConfigUrl
//     xmlhttp.onreadystatechange = function () {
//       if (this.readyState === 4 && this.status === 200) {
//         var responseData = JSON.parse(this.responseText)
//         // alert(responseData)
//         // console.log(responseData)
//         clientConfig = responseData
//         // sid = clientConfig.group.groupcode
//         //  alert('completed sid = '+ sid)
//       }
//     }
//     xmlhttp.open('POST', url, false)
//     xmlhttp.setRequestHeader('Content-type', 'application/json')
//     xmlhttp.send(JSON.stringify({'clientId': config.clientId}))
//   }
// }

// function isReturnUrlWithToken () {
//   var currentUrl = window.location.href
//   if (clientConfig !== undefined && clientConfig.purpose !== undefined) {
//     returnUrl = clientConfig.purpose.returnurl
//   }
//   // alert(returnUrl)
//   if (currentUrl.startsWith(returnUrl)) {
//     var token = getParameterByName('token', currentUrl)
//     if (isTokenNotNull(token)) {
//       return true
//     } else {
//       return false
//     }
//   } else {
//     return false
//   }
// }

// function isReturnUrl () {
//   var currentUrl = window.location.href
//   if (clientConfig !== undefined && clientConfig.purpose !== undefined) {
//     returnUrl = clientConfig.purpose.returnurl
//   }
//   // alert(returnUrl)
//   if (currentUrl.startsWith(returnUrl)) {
//     return true
//   } else {
//     return false
//   }
// }

// function getParameterByName (name, url) {
//   if (!url) url = window.location.href
//   name = name.replace(/[[]]/g, '\\$&')
//   var regex = new RegExp('[?&]' + name + '(=([^&#]*)|&|#|$)')
//   var results = regex.exec(url)
//   if (!results) return null
//   if (!results[2]) return ''
//   return decodeURIComponent(results[2].replace(/\+/g, ' '))
// }

function getCookieVal (offset) {
  var endstr = document.cookie.indexOf(';', offset)
  if (endstr === -1) {
    endstr = document.cookie.length
  }
  return unescape(document.cookie.substring(offset, endstr))
}

function GetCookie (name) {
  var arg = name + '='
  var alen = arg.length
  var clen = document.cookie.length
  var i = 0
  while (i < clen) {
    var j = i + alen
    if (document.cookie.substring(i, j) === arg) return getCookieVal(j)
    i = document.cookie.indexOf(' ', i) + 1
    if (i === 0) break
  }
  return null
}
function removeToken () {
  deleteCookie(config.clientId)
}

function isSessionPresent () {
  var token = GetCookie(config.clientId)
  return new Promise((resolve, reject) => {
    var validateTokenUrl = config.mkclLoginServerURL + VALIDATE_TOKEN_API
    axios.post(validateTokenUrl, {'clientId': config.clientId, 'token': token}).then(function (res) {
      resolve(res)
    }).catch(function (err) {
      removeToken()
      reject(err)
    })
  })
}

// function isLoginPage () {
//   var localloginBtn = document.getElementById('mkclSignIn')
//   if (localloginBtn != null) { return true } else { return false }
// }

// function setSession (token) {
//   if (isTokenNotNull(token)) {
//     // sessionStorage.setItem(sid, token)
//     document.cookie = sid + '=' + token
//   }
// }

function isLoggedIn () {
  // alert('function isLoggedIn()')
  var status = false
  isSessionPresent().then(function () {
    status = true
  }).catch(function () {
    status = false
  })
  return status
}

function deleteCookie (name) {
  document.cookie = name + '=;expires=Thu, 01 Jan 1970 00:00:01 GMT;'
}

// function isTokenNotNull (token) {
//   if (token === undefined || token === null || token === '' || token.length === 0) {
//     return false
//   }
//   return true
// }

function removeGroupMembersCookies () {
  console.log(clientConfig)
  var groupmembers = clientConfig.group.groupmembers
  for (var i = 0; i < groupmembers.length; i++) {
    var sid = groupmembers[i]
    // alert(groupmembers[i])
    deleteCookie(sid)
  }
}
function logout () {
  return new Promise((resolve, reject) => {
    var logoutUrl = config.mkclLoginServerURL + LOGOUT_API
    axios.get(logoutUrl).then(function () {
      // removeGroupMembersCookies()
      deleteCookie(config.clientId)
      resolve()
    }).catch(function (e) {
      reject(e)
    })
    // getClientConfig().then(function () {
    //   var logoutUrl = 'http://' + config.mkclLoginServerIp + LOGOUT_API
    //   axios.get(logoutUrl).then(function () {
    //     removeGroupMembersCookies()
    //     resolve()
    //   }).catch(function () {
    //     throw new Error('failed to logout')
    //   })
    // }).catch(function (e) {
    //   reject(e)
    // })
  })
}

function getToken () {
  return GetCookie(config.clientId)
}

// export {logout, isLoggedIn, getToken, login, isSessionPresent}

// var MKCLAPI = (function () {
//   return {
//     config: config,
//     getToken: getToken,
//     login: login,
//     logout: logout,
//     isLoggedIn: isLoggedIn,
//     isSessionPresent: isSessionPresent
//   }
// })()

function init (options) {
  Object.keys(options).forEach(function (key) {
    config[key] = options[key]
  })
}

export default {
  config: config,
  init: init,
  getToken: getToken,
  login: login,
  logout: logout,
  isLoggedIn: isLoggedIn,
  isSessionPresent: isSessionPresent,
  loginPoc: loginPoc
}
