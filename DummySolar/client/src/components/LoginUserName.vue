<template>
    <div class="login">
      <div class="columns">
        <div class="column"></div>
        <div class="column">
          <div class="field">
            <div class="control">
              <div id="mkclSignIn" @click="loginWithMkcl" class="button is-info" data-myReturnUrl="http://10.4.1.186:8082/#/home">Login with Mkcl</div>
              <!-- <a href="#" @click.prevent="openLogin">Login with mkcl</a>     -->
              <!-- <label class="label">{{ $t('login.username') }}</label>
              <div class="control has-icon has-icon-left">
                <input class="input" type="text" :placeholder="$t('login.placeholder.username')" v-model="username">      
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
                <button class="button is-primary" id="next_button" @click="validate_username(username)">{{ $t('login.button.next') }}</button>
              </p>
          </div>  
          <div class="field">
            <div class="control">
              <i>{{ $t('login.Reg_msg') }}</i>
              <router-link to="/register">{{ $t('login.Reg_SignUp') }}</router-link>
            </div> -->
          </div> 
        </div>
        <div class="column">
          <!-- <p class="button" @click="goHome"> home</p>-->
        </div>
      </div>
  </div>
</div>
</template>
<script>
import axios from 'axios'
export default {
  name: 'hello',
  data () {
    return {
      username1: '',
      username: '',
      result: '',
      CONFIG: {}
    }
  },
  methods: {
    loginWithMkcl () {
      // window.MKCLAPI.login()
      window.MKCLAPI.loginPoc()
    },
    goHome () {
      this.$router.push({path: '/home'})
    },
    openLogin () {
      var path = 'http://localhost:8080/#/login?clientId=' + this.$store.state.clientId
      window.location.replace(path)
    },
    validate_username: function (username1) {
      if (this.username !== '') {
        var loginForm = this
        axios.post('/server/o/login/validateusername', { 'username': username1 })
           .then(Response => {
             this.result = Response.data
            //  loginForm.$store.commit('SET_TOKEN', Response.headers.authorization)
             // alert("Login Successfull")
             if (Response.status === 200) {
               loginForm.$toasted.success('Username is correct')
               var mypath = '/loginpassword/' + username1
               this.$router.push({ path: mypath })
             }
           })
           .catch(Response => {
             loginForm.$toasted.error('Username is incorrect')
           })
      } else {
        var loginForm1 = this
        loginForm1.$toasted.success('UserName is Blank')
      }
    }
  },
  mounted () {
    console.log(window)
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