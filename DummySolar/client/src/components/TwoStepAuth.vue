<template>
<div>
  <h1>-----TWO STEP VERIFICATION ----</h1>
    <div class="columns">
    <div class="column"> </div>
    <div class="column"> 
    <div>
       
        <!--<p>{{ showqrimage }} {{isShowQr}}</p>-->

        <div v-if="!showqrimage">
          <p>scan the following image and enter OTP IN BOX</p>
          <img :src="qrcode_path" width="200" height="200">
        </div>        
      
        <input class="input" type="text" v-model="google_pin" placeholder="enter otp here"/>
        <button class="button is-info" type="submit" @click="verify_google_pin()">VERIFY</button>
              <!--</p>-->
      </div>
<button @click="downloadImage()">DOWNLOAD</button>
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
      otp_mobile: JSON.parse(JSON.stringify(this.$store.state.config.purpose.settings.twostepauth.type1.sms)),
      otp_email: JSON.parse(JSON.stringify(this.$store.state.config.purpose.settings.twostepauth.type1.email)),
      google_auth: JSON.parse(JSON.stringify(this.$store.state.config.purpose.settings.twostepauth.type1.google_authenticator)),
      qrcode_path: '',
      result: '',
      google_pin: '',
      returnURL: JSON.parse(JSON.stringify(this.$store.state.config.purpose.returnurl)),
      showqrimage: false,
      test: 'deepak'

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
        axios.post(`o/login/googleauthappcheck`, { 'username': this.username, 'otp': this.google_pin })
            .then(Response => {
              // console.log('response=' + Response.data)
              this.result = Response.data
              if (Response.status === 208) {
                loginForm.$toasted.error('max session limit reached')
              }
              if (Response.status === 200) {
                var finalpath = loginForm.returnURL
                loginForm.$store.commit('SET_TOKEN', Response.headers.authorization)
                loginForm.$store.commit('SET_IS_LOGGEDIN')
                loginForm.$toasted.success('otp is correct')
                loginForm.$toasted.success('Logged In Successfully')
                window.location.href = finalpath
              }
            })
            .catch(e => {
              if (e.response.status === 417) {
                loginForm.$toasted.error('otp is incorrect')
              }
            //  if (e.response.status === 500) {
            //    loginForm.$toasted.error('Captcha verification failed')
            //  }
            })
      } else {
        loginForm.$toasted.error('username is Blank')
      }
    },
    downloadImage: function () {
      console.log('hi')
      var url = 'http://localhost:8080/server/sushmita.png'
      window.open(url, '_blank')
    }
  },
  mounted: function () {
    // var con = this
    this.username = this.$route.params.username
    this.showqrimage = (this.$route.params.showqrimage === 'true')
    // console.log('1 show qr code:', this.showqrimage)
    // console.log('2:', this.$route.params.username)
    if (this.username !== '') {
      var loginForm = this
      axios.post(`o/login/googleauthapp`, { 'username': this.username })
          .then(Response => {
            // console.log('response=' + Response.data)
            this.qrcode_path = 'http://localhost:8080/server/' + loginForm.username + '.png'
          })
    }
  }
}
</script>

<style>

</style>
