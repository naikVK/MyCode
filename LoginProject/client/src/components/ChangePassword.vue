<template>
<div id="ChangePassword">
  <div class="columns">
    <div class="column"></div>
      <div id="changePass" class="control column is-half">
        <div class="field">
          <label class="label is-left">{{ $t('ForgotPassword.newpass') }}</label>
          <div class="columns">
            <div class="control column is-three-quarters">
              <input class="input" ref="txt_newpass" type="password" id="txt_password" tabindex="1" :placeholder="$t('ForgotPassword.placeholder.newpass')" @keypress="passwordChanged" v-model="user.password">
            </div>
            <div class="control column">
              <p id="showPass">        
                <span  class="button" @click="showHidePassword" v-if="!isShowPass">
                  <i v-if="!isShowPass" class="fa fa-eye"></i>
                  Show
                </span>
                <span  class="button" @click="showHidePassword" v-if="isShowPass">
                  <i class="fa fa-eye-slash"></i>
                  Hide
                </span>     
              </p>
            </div>
          </div>
           <p v-show="!errors.has('password')" id="strength"></p> 
      </div>

      <div class="field">
        <label class="label">{{ $t('ForgotPassword.confirmpass') }}</label>
        <div class="control column">
          <input class="input" type="password" :placeholder="$t('ForgotPassword.placeholder.confirmpass')" v-model="user.confirmpassword">
        </div>
      </div>
      <div class="field is-grouped">
        <p class="control">
        <button id="change_pass" class="button is-primary"@click="changePassword()">{{ $t('ForgotPassword.button.submit') }}</button>
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
  name: 'ChangePassword',
  data () {
    return {
      user: {},
      password: '',
      confirmpassword: '',
      isShowPass: ''
    }
  },
  methods: {
    changePassword: function () {
      var changepassword = this
      if ((this.user.password) !== (this.user.confirmpassword)) {
        changepassword.$toasted.error(changepassword.$t('ForgotPassword.alerts.nomatch'))
      }
      axios.post(`o/forgotpassword/changepassword`, {'username': this.username1, 'password': this.user.password, 'confirmpassword': this.user.confirmpassword})
        .then(Response => {
          this.user = Response.data
          if (Response.status === 200) {
            changepassword.$store.commit('SET_USERNAME', this.user.username)
            changepassword.$toasted.success(changepassword.$t('ForgotPassword.alerts.changesuccess'))
            changepassword.$toasted.info(changepassword.$t('ForgotPassword.alerts.login'))
            var mypath = '/'
            this.$router.push({ path: mypath })
          }
        })
        .catch(e => {
          if (e.response.status === 417) {
            var changepassword = this
            changepassword.$toasted.error('Validation Failed')
          }
        })
    },
    passwordChanged () {
      var strength = document.getElementById('strength')
      var strongRegex = new RegExp(
        '^(?=.{8,})(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9])(?=.*\\W).*$',
        'g'
      )
      var mediumRegex = new RegExp(
        '^(?=.{7,})(((?=.*[A-Z])(?=.*[a-z]))|((?=.*[A-Z])(?=.*[0-9]))|((?=.*[a-z])(?=.*[0-9]))).*$',
        'g'
      )
      var enoughRegex = new RegExp('(?=.{6,}).*', 'g')
      var pwd = document.getElementById('txt_password')
      if (strongRegex.test(pwd.value)) {
        strength.innerHTML = "<span style='color:green'>Strong!</span>"
      } else if (mediumRegex.test(pwd.value)) {
        strength.innerHTML = "<span style='color:orange'>Medium!</span>"
      } else if (enoughRegex.test(pwd.value)) {
        strength.innerHTML = "<span style='color:red'>Weak!</span>"
      } else {
        strength.innerHTML = '<span style="color:red">Very Weak!</span>'
      }
    },
    showHidePassword () {
      if (this.isShowPass) {
        this.isShowPass = false
        document
          .getElementById('txt_password')
          .setAttribute('type', 'password')
      } else {
        this.isShowPass = true
        document.getElementById('txt_password').setAttribute('type', 'text')
      }
    }
  },
  mounted: function () {
    this.$refs.txt_newpass.focus()
  }
}
</script>

<style>
#changePass {
    border-radius: 5px;
    background-color: #f2f2f2;
    padding: 20px;
}
.h1 {
  font-family: Helvetica Bold;
  color: #333;
}
</style>
