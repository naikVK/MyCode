import axios from 'axios'

function getCookieVal (offset) {
  var endstr = document.cookie.indexOf(';', offset)
  if (endstr === -1) {
    endstr = document.cookie.length
  }
  return unescape(document.cookie.substring(offset, endstr))
}

export function getCookie (name) {
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

function setCookie (name, value) {
  document.cookie = name + '=' + value
}

function deleteCookie (name) {
  document.cookie = name + '=;expires=Thu, 01 Jan 1970 00:00:01 GMT;'
}

export function isSessionPresent (sid) {
  var token = getCookie(sid)
  if (token === undefined || token === null || token === '' || token.length === 0) {
    return false
  }
  return true
}

export function getSession (sid) {
  var token = getCookie(sid)
  return token
}

export function setSession (sid, token) {
  setCookie(sid, token)
}

export function deleteSession (sid) {
  deleteCookie(sid)
}

export function redirect (redirectTo) {
  window.location.href = redirectTo
}

export function isTokenNotNull (token) {
  if (token === undefined || token === null || token === '' || token.length === 0) {
    return false
  }
  return true
}
export function isNull (val) {
  if (val === null || val === undefined || val.length === 0) {
    return true
  } else {
    return false
  }
}
export function isTokenValid (clientId, token) {
  return new Promise((resolve, reject) => {
    if (isNull(clientId) || isNull(token)) {
      reject(new Error('bad request'))
    } else {
      axios.post('/o/isValidToken', {'clientId': clientId, 'token': token}).then(function (response) {
        if (response.status === 200) {
          resolve(true)
        }
      }).catch(function (err) {
        reject(err)
      })
    }
  })
}
