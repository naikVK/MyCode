<template>
    <div class="login">
        <loading v-show="isLoading"></loading>
        <div class="columns" v-if="!isLoading">
        <div class="column"></div>
        <div class="column">
          <div class="field">
            <div class="control">
              <label class="label">{{ $t('login.username') }}</label>
              <div class="control has-icon has-icon-left">
                <input class="input" @keyup.enter="trigger" type="text" ref="txt_username" :placeholder="$t('login.placeholder.username')" v-model="username">      
                  <span class="icon is-small is-left">
                    <i class="fa fa-user"></i>
                  </span>
              </div>
            </div>
          </div>
          <p class="control">
            <router-link to="/forgotpassword">{{ $t('forgotPassword.url') }}</router-link>
          </p>
           <div class="field is-grouped">
              <p class="control">
                <button class="button is-primary" ref='btn_next' id="next_button" @click="validate_username(username)">{{ $t('login.button.next') }}</button>
              </p>
          </div>  
          <div class="field">
            <div class="control">
              <i>{{ $t('login.Reg_msg') }}</i>
              <router-link to="/register">{{ $t('login.Reg_SignUp') }}</router-link>
            </div>
          </div>
        </div>
        <div class="column"></div>
      </div>
  </div>
</div>
</template>
<script>
import axios from 'axios'
import {getCookie, isTokenNotNull, isTokenValid, redirect, deleteSession, isNull} from '@/utils/utils.js'
import loading from './Loading'
export default {
  name: 'hello',
  components: {
    loading
  },
  data () {
    return {
      username: '',
      result: '',
      CONFIG: {},
      isLoading: true
    }
  },
  methods: {
    validate_username: function (username) {
      if (this.username !== '') {
        var loginForm = this
        loginForm.$store.commit('SET_USERNAME', this.username)
        axios.post(`o/login/validateusername`, { 'username': username })
           .then(Response => {
             this.result = Response.data
            //  console.log(this.result)
             if (Response.status === 200) {
               loginForm.$toasted.success('Username is correct')
              //  console.log(username)
               loginForm.$store.commit('SET_USERNAME', this.result)
               var mypath = '/loginpassword'
               this.$router.push({ path: mypath })
             }
           })
           .catch(Response => {
             loginForm.$toasted.error('Username is incorrect', {
               duration: 5000
             })
           })
      } else {
        var loginForm1 = this
        loginForm1.$toasted.success('UserName is Blank')
      }
    },
    trigger () {
      this.validate_username(this.username)
    },
    getClientConfig () {
      var configForm = this
      var clientId = this.$route.query.clientId
      if (isNull(clientId)) {
        var clieIDFromStore = this.$store.getters.getClientId
        if (!isNull(clieIDFromStore)) {
          clientId = clieIDFromStore
        } else {
          configForm.$toasted.error('ClientId not provided')
          return
        }
      }
      this.$store.commit('SET_CLIENT_ID', clientId)
      window.axios.post(`/o/getclientconfig`, { 'clientid': clientId })
            .then(Response => {
              this.CONFIG = Response.data
              if (Response.status === 200) {
                configForm.$store.commit('SET_CONFIG', Response.data)
                configForm.$toasted.success('Client info found')
                configForm.checkSessionAvailability()
              }
            })
            .catch(Response => {
              configForm.$toasted.error('Client info not found', {
                duration: 5000
              })
            })
    },
    checkForGroupSession () {
      // alert('checkForGroupSession')
      var loginForm = this
      // alert(clientId)
      var clientId = this.$route.query.clientId
      var clientConfig = loginForm.$store.getters.getConfig
      var returnUrl = clientConfig.purpose.returnurl
      var groupmembers = clientConfig.group.groupmembers
      groupmembers.forEach(function (member, index, groupmembers) {
        var tokenCookie = getCookie(member)
        // alert(member + ' ' + tokenCookie)
        // alert(member + '->' + tokenCookie)
        if (isTokenNotNull(tokenCookie)) {
          // alert('cookie found==>' + tokenCookie)
          // alert('before tokenCookie' + tokenCookie)
          isTokenValid(clientId, tokenCookie).then(function () {
            // alert('after tokenCookie' + foundToken)
            loginForm.$store.commit('SET_TOKEN', tokenCookie)
            loginForm.$store.commit('SET_IS_LOGGEDIN')
            loginForm.isLoading = true
            redirect(returnUrl)
          }).catch(function () {
            loginForm.isLoading = false
            // delete invalid or expired token cookie
            deleteSession(member)
          }).next(function () {

          })
        } else {
          loginForm.isLoading = false
        }
      })
    },
    checkSessionAvailability () {
      // alert('checkSessionAvailability')
      var loginForm = this
      var clientId = this.$route.query.clientId
      // alert(clientId)
      var clientConfig = loginForm.$store.getters.getConfig
      var returnUrl = clientConfig.purpose.returnurl
      if (!isNull(clientId)) {
        // alert('client id not null')
        var token = getCookie(clientId)
        // alert('token =>' + token)
        isTokenValid(clientId, token).then(function (res) {
          redirect(returnUrl)
        }).catch(function () {
          loginForm.checkForGroupSession()
        })
      }
    }
  },
  beforeMount: function () {
    this.getClientConfig()
    if (!window.navigator.cookieEnabled) {
      this.$toasted.error('Cookies are disabled, Kindly enable it', {
        duration: null,
        position: 'top-center',
        fullWidth: 'true',
        theme: 'bubble'
      })
      this.$refs.btn_next.disabled = true
    }
  },
  //   getClientConfig () {
  //     var configForm = this
  //     var clientId = this.$route.query.clientId
  //     if (isNull(clientId)) {
  //       configForm.$toasted.error('ClientId not provided')
  //       return
  //     }
  //     this.$store.commit('SET_CLIENT_ID', clientId)
  //     axios.post(`/o/getclientconfig`, { 'clientid': clientId })
  //           .then(Response => {
  //             this.CONFIG = Response.data
  //             if (Response.status === 200) {
  //               configForm.$store.commit('SET_CONFIG', Response.data)
  //               configForm.$toasted.success('Client info found')
  //               configForm.checkSessionAvailability()
  //             }
  //           })
  //           .catch(Response => {
  //             configForm.$toasted.error('Client info not found', {
  //               duration: 5000
  //             })
  //           })
  //   },
  //   checkForGroupSession () {
  //     var loginForm = this
  //     // alert(clientId)
  //     var clientId = this.$route.query.clientId
  //     var clientConfig = loginForm.$store.getters.getConfig
  //     var returnUrl = clientConfig.purpose.returnurl
  //     var groupmembers = clientConfig.group.groupmembers
  //     // console.log(groupmembers)
  //     groupmembers.forEach(function (member, index, groupmembers) {
  //       var tokenCookie = getCookie(member)
  //       // alert(groupmembers[i] + '==>' + tokenCookie)
  //       if (isTokenNotNull(tokenCookie)) {
  //         // alert('cookie found==>' + tokenCookie)
  //         // alert('before tokenCookie' + tokenCookie)
  //         isTokenValid(clientId, tokenCookie).then(function () {
  //           // alert('after tokenCookie' + foundToken)
  //           loginForm.$store.commit('SET_TOKEN', tokenCookie)
  //           loginForm.$store.commit('SET_IS_LOGGEDIN')
  //           redirect(returnUrl)
  //         }).catch(function () {
  //           // delete invalid or expired token cookie
  //           deleteSession(member)
  //         })
  //       }
  //     })
  //   },
  //   checkSessionAvailability () {
  //     var loginForm = this
  //     var clientId = this.$route.query.clientId
  //     // alert(clientId)
  //     var clientConfig = loginForm.$store.getters.getConfig
  //     var returnUrl = clientConfig.purpose.returnurl
  //     if (!isNull(clientId)) {
  //       // alert('client id not null')
  //       var token = getCookie(clientId)
  //       // alert('token =>' + token)
  //       isTokenValid(clientId, token).then(function (res) {
  //         redirect(returnUrl)
  //       }).catch(function () {
  //         loginForm.checkForGroupSession()
  //       })
  //     }
  //   }
  // },
  // beforeMount: function () {
  //   this.getClientConfig()
  //   if (!window.navigator.cookieEnabled) {
  //     this.$toasted.error('Cookies are disabled, Kindly enable it', {
  //       duration: null,
  //       position: 'top-center',
  //       fullWidth: 'true',
  //       theme: 'bubble'
  //     })
  //     this.$refs.btn_next.disabled = true
  //   }
  // },
  mounted () {
    // this.$refs.txt_username.focus()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#login {
    border-style: solid;
    top: 100px;
    padding:100px ;
    border-color: black
}
#login_button 
{
    background-color: lawngreen
  
}

</style>