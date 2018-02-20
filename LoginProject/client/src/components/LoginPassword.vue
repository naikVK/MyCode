<template>
    <div class="login">
      <div class="columns">
        <div class="column"></div>
        <div class="column">
          <div class="field">
            <div v-if="showpassword" class="control">
              <label class="label">{{ $t('login.Password') }}</label>
              <input class="input" @keyup.enter="trigger1"  ref="txt_password" type="password" :placeholder="$t('login.placeholder.Password')" v-model="password">
            </div>
            <div v-if="showOTPphone" class=control >
            <p> {{ $t('login.send_otpPhone_msg') }}</p><p>{{ phoneNumber}}
                    <br/>
                    <br>
                    <input class="input"  @keyup.enter="trigger2" autofocus type="text" :placeholder="$t('login.placeholder.OTP')" v-model="OTP">   
                    </input>
                    <br>
                    <br />
                    <button id="submit_OTP" class="button is-primary" type="submit" @click="verify_OTP()">{{ $t('login.button.verify_otp') }}</button>
                    <button id="resendotp" class="button is-primary" @click="Send_OTP_Phone(username1)">{{ $t('login.button.resend_otp') }}</button>
                    </form>

                    </br>

                </p>
                </div>
                <div v-if="showOTPemail" class="control">
                  <p> {{ $t('login.send_otpPhone_msg_email') }}</p><p>{{ emailid }}
                    <br/>
                    <br>
                    <input class="input" autofocus @keyup.enter="trigger2"  type="text" :placeholder="$t('login.placeholder.OTP')" v-model="OTP">   
                    </input>
                    <br>
                    <br />
                    <button id="submit_OTP" class="button is-primary" type="submit" @click="verify_OTP()">{{ $t('login.button.verify_otp') }}</button>
                    <button id="resendotp" class="button is-primary" @click="Send_OTP_Email(username1)">{{ $t('login.button.resend_otp') }}</button>
                    </form>

                    </br>

                </p>
                </div>
          </div>
          <div class="field">
            <div class="control">
                   <p v-if="OTPoption">
        <!--<div class="dropdown">-->
          <!--<span>{{ $t('login.OptiontoPassword') }}</span>-->
          <!--<div class="dropdown-content">-->
            <button class="button is-warning" v-if="OTPoption_Phone" @click="showModal()">{{ $t('login.button.send_otpPhone') }}</button>
            <br>
            <br>
            <button class="button is-warning" v-if="OTPoption_Email" @click="showModal2()">{{ $t('login.button.send_otpEmail') }}</button>
            
          <!--</div>-->
        <!--</div>-->
        </p>
            </div>
          </div>
          <div class="field">
            <div class="control">
               <p v-if="captcha">
        <vue-recaptcha ref="captcha" @render="captchaRendered" @verify="verifyCapRes" :sitekey="rcapt_sig_key"></vue-recaptcha>
        </p>
            </div>
          </div>
          <div class="field">
            <div class="control">
               <button v-if="showloginbutton" class="button is-primary" id="" @click="validate_login(username, password)">{{ $t('login.button.login') }}</button>
            </div>
          </div>
        </div>
        <div class="column">
        </div>
      </div>
    </div>
</template>
<script>

import VModal from 'vue-js-modal'
import VueRecaptcha from 'vue-recaptcha'
import axios from 'axios'
import Vue from 'vue'
Vue.use(VModal)

export default {
  name: 'hello',
  components: { VueRecaptcha },
  data () {
    return {
      // user: {}
      // username1: '',
      username: this.$store.state.username,
      password: '',
      result: '',
      captcha: this.$store.state.config.purpose.settings.captcha,
      captchaResponse: '',
      rcapt_sig_key: '6Lf-hzoUAAAAAJEo2lPMFw2zu-cm46bQc4dJ4fP6',
      rcapt_id: 0,
      res: '',
      twostepauth: this.$store.state.config.purpose.settings.twostepauth.set,
      google_authenticator: this.$store.state.config.purpose.settings.twostepauth.type_two_step_auth.google_authenticator,
      returnURL: this.$store.state.config.purpose.returnurl,
      OTP: '',
      phoneNumber: '',
      OTPverified: '',
      OTPoption: JSON.parse(JSON.stringify(this.$store.state.config.purpose.settings.otprequired.set)),
      emailid: '',
      OTPoption_Phone: JSON.parse(JSON.stringify(this.$store.state.config.purpose.settings.otprequired.type_otp.mobile)),
      OTPoption_Email: JSON.parse(JSON.stringify(this.$store.state.config.purpose.settings.otprequired.type_otp.email)),
      showOTPphone: false,
      showpassword: true,
      showOTPemail: false,
      showloginbutton: true
    }
  },
  methods: {
    showModal (modal1) {
      this.Send_OTP_Phone()
      // this.$modal.show('modal1')
      this.showOTPphone = true
      this.showpassword = false
      this.OTPoption_Phone = false
      this.captcha = false
      this.showloginbutton = false
      this.showOTPemail = false
    },
    Send_OTP_Phone () {
      var loginForm = this
      axios.post('o/login/sendOTPonPhone', { 'username': this.username })
           .then(Response => {
             this.phoneNumber = Response.data
             var phone = this.phoneNumber.substring(10, this.phoneNumber.length - 3)
             this.phoneNumber = '*******' + phone
             if (Response.status === 200) {
               loginForm.$toasted.success('OTP Sent on Mobile')
             }
           })
           .catch(e => {
             if (e.response.status === 417) {
               loginForm.$toasted.error('OTP not sent', {
                 duration: 5000
               })
             }
           })
    },
    verify_OTP () {
      var loginForm = this
      axios.post('o/login/verifyOTP', { 'username': this.username, 'otp': this.OTP })
           .then(Response => {
             this.result = Response.data
             if (Response.status === 200 & this.google_authenticator) {
               loginForm.$toasted.success('OTP is correct')
               loginForm.$toasted.info('Two step verification on')
               loginForm.$store.commit('SET_RESTRICT_TOKEN', Response.headers.authorization)
               loginForm.$store.commit('SET_SHOWQRCODE', this.result)
               var mypath = '/twostepverification'
               this.$router.push({ path: mypath })
             }
             if (Response.status === 200 & !this.twostepauth) {
               loginForm.$toasted.success('OTP is correct')
               loginForm.$store.commit('SET_TOKEN', Response.headers.authorization)
               loginForm.$store.commit('SET_IS_LOGGEDIN')
               var finalpath = loginForm.returnURLe
               if (finalpath !== '' && finalpath !== undefined && finalpath !== null) {
                 window.location.href = finalpath
               }
             }
           })
           .catch(e => {
             if (e.response.status === 417) {
               loginForm.$toasted.error('OTP is incorrect', {
                 duration: 5000
               })
               this.OTP = ''
             }
             if (e.response.status === 403) {
               loginForm.$toasted.error('ClientId not matched at server', {
                 duration: 5000
               })
             }
           })
    },
    showModal2 (modal2) {
      // this.$modal.show('modal2')
      this.Send_OTP_Email()
      this.showOTPemail = true
      this.showpassword = false
      this.OTPoption_Email = false
      this.captcha = false
      this.showloginbutton = false
      this.showOTPphone = false
    },
    Send_OTP_Email () {
      var loginForm = this
      axios.post('o/login/sendOTPonEmail', { 'username': this.username })
        .then(Response => {
          this.emailid = Response.data
          var fields = this.emailid.split(/@/)
          var name = fields[1]
          this.emailid = '***********' + name
          if (Response.status === 200) {
            loginForm.$toasted.success('OTP Sent on Email')
          }
        })
           .catch(e => {
             if (e.response.status === 417) {
               loginForm.$toasted.error('OTP not sent', {
                 duration: 5000
               })
             }
           })
    },
    captchaRendered: function (id) {
      // console.log(id)
    },
    verifyCapRes: function (res) {
      this.captchaResponse = res
    },
    validate_login: function () {
      var loginForm = this
      if (this.captcha && this.captchaResponse === '') {
        loginForm.$toasted.info('Please verify captcha')
        return
      }
      if (this.username !== '' && this.password !== '') {
        axios.post('o/login/validateuser', { 'username': this.username, 'password': this.password, 'captchaResponse': this.captchaResponse })
           .then(Response => {
             this.result = Response.data
             if (Response.status === 208) {
               loginForm.$toasted.error('max session limit reached')
             }
             if (Response.status === 200 & this.google_authenticator) {
               loginForm.$toasted.success('password is correct')
               loginForm.$toasted.info('Two step verification on')
               loginForm.$store.commit('SET_RESTRICT_TOKEN', Response.headers.authorization)
               loginForm.$store.commit('SET_SHOWQRCODE', this.result)
               var mypath = '/twostepverification'
               this.$router.push({ path: mypath })
             }
             if (Response.status === 200 & !this.twostepauth) {
               loginForm.$toasted.success('password is correct')
               var finalpath = loginForm.returnURL
               loginForm.$store.commit('SET_TOKEN', Response.headers.authorization)
               loginForm.$store.commit('SET_IS_LOGGEDIN')
               if (finalpath !== '' && finalpath !== undefined && finalpath !== null) {
                 window.location.href = finalpath
               } else {
                 this.$router.push({ path: '/home' })
               }
             }
           })
           .catch(e => {
             if (e.response.status === 417) {
               loginForm.$toasted.error('Password is incorrect', {
                 duration: 5000
               })
               this.password = ''
             }
             if (e.response.status === 500) {
               loginForm.$toasted.error('Captcha verification failed', {
                 duration: 5000
               })
             }
             if (e.response.status === 403) {
               loginForm.$toasted.error('ClientId not matched at server', {
                 duration: 5000
               })
             }
           })
      } else {
        var loginForm1 = this
        loginForm1.$toasted.success('Password is Blank')
      }
    },
    trigger1 () {
      this.validate_login(this.username, this.password)
    },
    trigger2 () {
      this.verify_OTP()
    }
  },
  mounted: function () {
    this.$refs.txt_password.focus()
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
.dropdown {
    position: relative;
    display: inline-block;
}

.dropdown-content {
    display: none;
    position: absolute;
    background-color: #f9f9f9;
    min-width: 200px;
    box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
    padding: 12px 16px;
    z-index: 1;
}

.dropdown:hover .dropdown-content {
     outline: 1px solid #ccc;
    display: block;
}
</style>