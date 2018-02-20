<template>
  <div id="ForgotPasswordOTPVerification">
    <div class="columns">
      <div class="column"></div>
      <div id="verifyOTP"class="column">
        <h1>Verify OTP..</h1>
        <div class="field">
          <div class="control">
            <label class="label">{{ $t('ForgotPassword.enterotp') }}</label>
            <input class="input" ref="txt_otp" type="text" tabindex="1" :placeholder="$t('ForgotPassword.placeholder.otp')" v-model="otp">
          </div>
        </div>
        <div class="field is-grouped">
            <p class="control">
              <button id="verify_otp" class="button is-primary" @click="verifyotp()">{{ $t('ForgotPassword.button.verifyotp') }}</button>
            </p>
            <p class="control">
              <button id="resendotp" class="button is-info" @click="resendotp(username1)">{{ $t('ForgotPassword.button.resendotp') }}</button>
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
  name: 'ForgotPasswordOTPVerification',
  data () {
    return {
      profiledata: {},
      result: '',
      otp: '',
      username1: JSON.parse(JSON.stringify(this.$store.state.username))
    }
  },
  methods: {
    verifyotp: function () {
      var forgotpassword = this
      axios.post(`o/forgotpassword/verifyotp`, {'username': this.username1, 'otp': this.otp})
      .then(Response => {
        var otpverify = this
        this.result = Response.data
        if (Response.status === 200) {
          otpverify.$store.commit('SET_RESTRICT_TOKEN', Response.headers.authorization)
          otpverify.$toasted.success(forgotpassword.$t('ForgotPassword.alerts.otpverify'))
          var mypath = '/changepassword'
          this.$router.push({ path: mypath })
        }
      })
      .catch(e => {
        if (e.response.status === 500) {
          var forgotpassword = this
          forgotpassword.$toasted.error(forgotpassword.$t('ForgotPassword.alerts.otpfail'))
        }
      })
    },
    resendotp: function (username1) {
      axios.post(`o/forgotpassword/resendotp`, {'username': username1})
      .then(Response => {
        var forgotpassword = this
        this.profiledata = Response.data
        forgotpassword.$toasted.info(forgotpassword.$t('ForgotPassword.alerts.otpresend'))
      })
    }
  },
  mounted: function () {
    this.$refs.txt_otp.focus()
  }
}
</script>

<style>
 #verifyOTP {
    border-radius: 5px;
    background-color: #f2f2f2;
    padding: 20px;
} 
.h1 {
  font-family: Helvetica Bold;
  color: #333;
}
</style>
