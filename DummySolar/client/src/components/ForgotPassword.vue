<template>
  <div id="ForgotPassword">
    <div class="columns">
      <div class="column"></div>
        <div id="sendOTP" class="column">
          <h1>Reset Password..</h1>
          <div class="field">
          <label class="label">{{ $t('ForgotPassword.username') }}</label>
            <div class="control">
              <input class="input" type="text" ref="txt_username" :placeholder="$t('ForgotPassword.placeholder.username')" @focusout="validate_username(username)" v-model="username">
            </div>
          </div>
          <div class="field">
            <div class="control">
                <p v-if="showEmail">
                  <input type="checkbox" id="isemail" v-model="checked">
                  <label for="isemail">{{ $t('ForgotPassword.isEmail') }}</label>
                </p>
            </div>
          </div>
          <div class="field is-grouped">
            <p class="control">
                <button id="send_otp" class="button is-primary" @click="sendotp(username)">{{ $t('ForgotPassword.button.sendotp') }}</button>
            </p>
            <p class="control">
                <button id="cancel" class="button is-info" @click="cancel()">{{ $t('ForgotPassword.button.cancel') }}</button>
            </p>
          </div>  
            
        </div>
      <div class="column"></div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'ForgotPassword',
  data () {
    return {
      profiledata: {},
      result: '',
      username: '',
      checked: '',
      showEmail: JSON.parse(JSON.stringify(this.$store.state.config.purpose.settings.otprequired.type2.email))
    }
  },
  methods: {
    validate_username: function (username) {
      if (this.username !== '') {
        axios.post(`o/login/validateusername`, { 'username': username })
           .then(Response => {
             var forgotpassword = this
             this.result = Response.data
             if (this.result === true) {
              //  var validatedUsername = username
               forgotpassword.$toasted.success(forgotpassword.$t('ForgotPassword.alerts.usernamevalid'))
             } else {
               forgotpassword.$toasted.error(forgotpassword.$t('ForgotPassword.alerts.usernameinvalid'))
             }
           })
      } else {
        var forgotpassword = this
        forgotpassword.$toasted.error(forgotpassword.$t('ForgotPassword.alerts.usernameblank'))
      }
    },
    sendotp: function (username1) {
      if (this.showEmail) {
        axios.post(`o/forgotpassword/sendotp`, {'username': username1, 'isemail': this.checked})
        .then(Response => {
          var forgotpassword = this
          this.profiledata = Response.data
          if (this.profiledata !== '') {
            forgotpassword.$toasted.success(forgotpassword.$t('ForgotPassword.alerts.otpsuccessme'))
            var mypath = '/forgotpasswordotpverification/' + username1
            this.$router.push({ path: mypath })
          }
        })
      } else {
        axios.post(`o/forgotpassword/sendotp`, {'username': username1})
        .then(Response => {
          var forgotpassword = this
          this.profiledata = Response.data
          if (this.profiledata !== '') {
            forgotpassword.$toasted.success(forgotpassword.$t('ForgotPassword.alerts.otpsuccessm'))
            var mypath = '/forgotpasswordotpverification/' + username1
            this.$router.push({ path: mypath })
          }
        })
      }
    },
    cancel: function () {
      var forgotpassword = this
      forgotpassword.$toasted.info(forgotpassword.$t('ForgotPassword.alerts.backtologin'))
      var home = '/'
      this.$router.push({ path: home })
    }
  },
  mounted: function () {
    this.$refs.txt_username.focus()
  }
}
</script>

<style>
#sendOTP {
    border-radius: 5px;
    background-color: #f2f2f2;
    padding: 20px;
}
.h1 {
  font-family: Helvetica Bold;
  color: #333;
}
</style>
