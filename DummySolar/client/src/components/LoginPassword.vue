<template>
    <div class="login">
      <div class="columns">
        <div class="column"></div>
        <div class="column">
          <div class="field">
            <div class="control">
              <label class="label">{{ $t('login.Password') }}</label>
              <input class="input" type="password" :placeholder="$t('login.placeholder.Password')" v-model="password">
            </div>
          </div>
          <div class="field">
            <div class="control">
                   <p v-if="OTPoption">
        <div class="dropdown">
          <span>{{ $t('login.OptiontoPassword') }}</span>
          <div class="dropdown-content">
            <p v-if="OTPoption_Phone" @click="showModal()">{{ $t('login.button.send_otpPhone') }}</p>
            <p v-if="OTPoption_Email" @click="showModal2()">{{ $t('login.button.send_otpEmail') }}</p>
            <modal name="modal1">
                <br/>
                <p> {{ $t('login.send_otpPhone_msg') }}</p><p>{{ phoneNumber}}
                    <i class="fa fa-times-circle" aria-hidden="true" @click="hideModal()"></i>
                    <br/>
                    <input type="text" :placeholder="$t('login.placeholder.OTP')" v-model="OTP">   
                    </input>
                    <br />
                    <button id="submit_OTP" type="submit" @click="verify_OTP()">{{ $t('login.button.verify_otp') }}</button>
                    <button id="resendotp" class="button is-primary" @click="Send_OTP_Phone(username1)">{{ $t('login.button.resend_otp') }}</button>
                    </form>

                    </br>

                </p>
            </modal>
             <modal name="modal2">
                <br/>
                 <p> {{ $t('login.send_otpPhone_msg_email') }}</p><p>{{ emailid }}
                    <i class="fa fa-times-circle" aria-hidden="true" @click="hideModal2()"></i>
                    <br/>
                    <input type="text" :placeholder="$t('login.placeholder.OTP')" v-model="OTP">   
                    </input>
                    <br />
                    <button id="submit_OTP" type="submit" @click="verify_OTP()">{{ $t('login.button.verify_otp') }}</button>
                    <button id="resendotp" class="button is-primary" @click="Send_OTP_Email(username1)">{{ $t('login.button.resend_otp') }}</button>
                    </form>

                    </br>

                </p>
            </modal>
          </div>
        </div>
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
               <button class="button is-primary" id="" @click="validate_login(username, password)">{{ $t('login.button.login') }}</button>
            </div>
          </div>
        </div>
        <div class="column"></div>
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
      username1: '',
      username: '',
      password: '',
      result: '',
      captcha: JSON.parse(JSON.stringify(this.$store.state.config.purpose.settings.captcha)),
      captchaResponse: '',
      rcapt_sig_key: '6Lf9KjkUAAAAALuoIeGeKhlfn6p1nSEk83Z0ruwC',
      rcapt_id: 0,
      res: '',
      twostepauth: JSON.parse(JSON.stringify(this.$store.state.config.purpose.settings.twostepauth.set)),
      google_authenticator: JSON.parse(JSON.stringify(this.$store.state.config.purpose.settings.twostepauth.type1.google_authenticator)),
      returnURL: JSON.parse(JSON.stringify(this.$store.state.config.purpose.returnurl)),
      OTP: '',
      phoneNumber: '',
      OTPverified: '',
      OTPoption: JSON.parse(JSON.stringify(this.$store.state.config.purpose.settings.otprequired.set)),
      emailid: '',
      OTPoption_Phone: JSON.parse(JSON.stringify(this.$store.state.config.purpose.settings.otprequired.type2.mobile)),
      OTPoption_Email: JSON.parse(JSON.stringify(this.$store.state.config.purpose.settings.otprequired.type2.email))
    }
  },
  methods: {
    showModal (modal1) {
      this.Send_OTP_Phone()
      this.$modal.show('modal1')
    },
    Send_OTP_Phone () {
      var loginForm = this
      axios.post('o/login/sendOTPonPhone', { 'username': this.username1 })
           .then(Response => {
             this.phoneNumber = Response.data
             if (Response.status === 200) {
               loginForm.$toasted.success('OTP Sent on Mobile')
             }
           })
           .catch(e => {
             if (e.response.status === 417) {
               loginForm.$toasted.error('OTP not sent')
             }
           })
    },
    verify_OTP () {
      var loginForm = this
      axios.post('o/login/verifyOTP', { 'username': this.username1, 'otp': this.OTP })
           .then(Response => {
             this.result = Response.data
             if (Response.status === 200 & this.google_authenticator) {
               loginForm.$toasted.success('OTP is correct')
               loginForm.$toasted.info('Two step verification on')
               var mypath = '/twostepverification/' + this.username1 + '/' + this.result
               this.$router.push({ path: mypath })
             }
             if (Response.status === 200 & !this.twostepauth) {
               loginForm.$toasted.success('OTP is correct')
               loginForm.$store.commit('SET_TOKEN', Response.headers.authorization)
               loginForm.$store.commit('SET_IS_LOGGEDIN')
               var finalpath = loginForm.returnURL
               window.location.href = finalpath
             }
           })
           .catch(e => {
             if (e.response.status === 417) {
               loginForm.$toasted.error('OTP is incorrect')
             }
             if (e.response.status === 403) {
               loginForm.$toasted.error('ClientId not matched at server')
             }
           })
    },
    hideModal (modal1) {
      this.$modal.hide('modal1')
    },
    showModal2 (modal2) {
      this.$modal.show('modal2')
      this.Send_OTP_Email()
    },
    Send_OTP_Email () {
      var loginForm = this
      axios.post('o/login/sendOTPonEmail', { 'username': this.username1 })
        .then(Response => {
          this.emailid = Response.data
          if (Response.status === 200) {
            loginForm.$toasted.success('OTP Sent on Email')
          }
        })
           .catch(e => {
             if (e.response.status === 417) {
               loginForm.$toasted.error('OTP not sent')
             }
           })
    },
    hideModal2 (showModal2) {
      this.$modal.hide('modal2')
    },
    captchaRendered: function (id) {
      console.log(id)
    },
    verifyCapRes: function (res) {
      this.captchaResponse = res
    },
    validate_login: function () {
      alert()
      var loginForm = this
      if (this.captcha && this.captchaResponse === '') {
        loginForm.$toasted.info('Please verify captcha')
        return
      }
      if (this.username1 !== '' && this.password !== '') {
        axios.post('/server/o/login/validateuser', { 'username': this.username1, 'password': this.password, 'captchaResponse': this.captchaResponse })
           .then(Response => {
             alert()
             this.result = Response.data
             // alert("Login Successfull")
             if (Response.status === 200) {
               alert(Response.data)
             }
             if (Response.status === 208) {
               loginForm.$toasted.error('max session limit reached')
             }
             if (Response.status === 200 & this.google_authenticator) {
               loginForm.$toasted.success('password is correct')
               loginForm.$toasted.info('Two step verification on')
               var mypath = '/twostepverification/' + this.username1 + '/' + this.result
               this.$router.push({ path: mypath })
             }
             if (Response.status === 200 & !this.twostepauth) {
               loginForm.$toasted.success('password is correct')
              //  var finalpath = loginForm.returnURL
               loginForm.$store.commit('SET_TOKEN', Response.headers.authorization)
               loginForm.$store.commit('SET_IS_LOGGEDIN')
              //  this.$router.push({ path: finalpath })
              //  window.location.href = finalpath
             }
           })
           .catch(e => {
             if (e.response.status === 417) {
               loginForm.$toasted.error('Password is incorrect')
             }
             if (e.response.status === 500) {
               loginForm.$toasted.error('Captcha verification failed')
             }
             if (e.response.status === 403) {
               loginForm.$toasted.error('ClientId not matched at server')
             }
           })
      } else {
        var loginForm1 = this
        loginForm1.$toasted.success('Password is Blank')
      }
    }
  },
  mounted: function () {
    // var con = this
    // alert(document.cookie)
    this.username1 = this.$route.params.username
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
    display: block;
}
</style>