<template>
<div>
  <h1>{{ $t('login.twostep.header') }}</h1>
    <div class="columns">
    <div class="column"> </div>
    <div class="column"> 
    <div>
        <div v-if="!showqrimage">
          <p>{{ $t('login.twostep.msg') }}</p>
          <img :src="qrcode_path" width="200" height="200">
        </div>        
      
        <input class="input" type="text" autofocus @keyup.enter="trigger1" v-model="google_pin" :placeholder="$t('login.placeholder.OTP')"/>
        <br>
        <button class="button is-info" type="submit" @click="verify_google_pin()">{{ $t('login.button.verify_otp') }}</button>
      </div>
      <br>
<button class="button is-info" type="submit" @click="GetSecretKey()">{{ $t('login.button.getSecretKey') }}</button>
    </div>
    <div class="column"> </div>
    </div>

  
 </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'hello',
  data () {
    return {
      otp_mobile: JSON.parse(JSON.stringify(this.$store.state.config.purpose.settings.twostepauth.type_two_step_auth.sms)),
      otp_email: JSON.parse(JSON.stringify(this.$store.state.config.purpose.settings.twostepauth.type_two_step_auth.email)),
      google_auth: JSON.parse(JSON.stringify(this.$store.state.config.purpose.settings.twostepauth.type_two_step_auth.google_authenticator)),
      qrcode_path: '',
      result: '',
      google_pin: '',
      returnURL: JSON.parse(JSON.stringify(this.$store.state.config.purpose.returnurl)),
      showqrimage: this.$store.state.showqrcodeimage,
      test: 'deepak',
      username: JSON.parse(JSON.stringify(this.$store.state.username))
    }
  },
  computed: {
    isShowQr: function () {
      return this.showqrimage
    }
  },
  methods: {
    verify_google_pin: function () {
      // console.log(this.username)
      if (this.username !== '') {
        var loginForm = this
        axios.post(`o/login/googleauthappcheck`, { 'otp': this.google_pin })
            .then(Response => {
              // console.log('response=' + Response.data)
              this.result = Response.data
              if (Response.status === 208) {
                loginForm.$toasted.error('max session limit reached')
              }
              if (Response.status === 200) {
                var finalpath = loginForm.returnURL
                // alert(finalpath)
                loginForm.$store.commit('SET_TOKEN', Response.headers.authorization)
                loginForm.$store.commit('SET_IS_LOGGEDIN')
                loginForm.$toasted.success('otp is correct')
                loginForm.$toasted.success('Logged In Successfully')
                window.location.href = finalpath
              }
            })
            .catch(e => {
              if (e.response.status === 417) {
                loginForm.$toasted.error('otp is incorrect', {
                  duration: 5000
                })
                this.google_pin = ''
              }
            })
      } else {
        loginForm.$toasted.error('username is Blank')
      }
    },
    GetSecretKey: function () {
      var loginForm = this
      axios.post(`o/login/googleauthgetkey`, { 'username': this.username })
          .then(Response => {
            // console.log(Response.data)
            if (Response.status === 200) {
              loginForm.$toasted.success('SECRETKEY IS SEND TO MOBILE ')
            }
          })
    },
    trigger1 () {
      this.verify_google_pin()
    }
  },
  mounted: function () {
    // console.log(this.showqrimage)
    if (this.username !== '' & !this.showqrimage) {
      // var loginForm = this
      axios.post(`o/login/googleauthapp`, { 'username': this.username })
          .then(Response => {
            this.qrcode_path = axios.defaults.baseURL + Response.data
          })
    }
  }
}
</script>

<style>

</style>
