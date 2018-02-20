<template>
  <div id="registerationForm" class="columns">
    <div class="column"></div>
    <div class="panel column is-two-thirds">
    <p class="panel-heading">Registration Form</p>
    <div class="panel-block">
    <div class="form">  
    <div class="control">
    <div class="field">
      <label class="label">{{ $t('registration.fullName') }}</label>
      <div class="control is-expanded">
        <input class="input" ref="txt_fullname" type="text" v-validate="'required'" name="fullname" v-model="signUpForm.personalDetails.fullName" :placeholder="$t('registration.placeholder.fullname')">
      </div>
      <p v-show="errors.has('fullname')" class="help is-danger">{{ errors.first('fullname') }}</p>
    </div>

      <div class="field">
      <label class="label">{{ $t('registration.dob') }}</label>
      <div class="control has-icons-left">
        <input class="input" type="date" id="dobCalendar" v-validate="'required'" name="dob" v-model="signUpForm.personalDetails.dob">
        <i class="icon is-left fa fa-calendar"></i>
      </div>
      <p v-show="errors.has('dob')" class="help is-danger">{{ errors.first('dob') }}</p>
    </div>

    <div class="field">
      <label class="label">{{ $t('registration.gender') }}</label>
      <div class="control">
        <label class="radio">
          <input type="radio" v-validate="'required'" v-model="signUpForm.personalDetails.gender" value="M" name="gender">
          {{ $t('registration.gender_male') }}
        </label>
        <label class="radio">
          <input type="radio"  v-validate="'required'" v-model="signUpForm.personalDetails.gender" value="F" name="gender">
          {{ $t('registration.gender_female') }}
        </label>
      </div>
      <p v-show="errors.has('gender')" class="help is-danger">{{ errors.first('gender') }}</p>
    </div>

    <div class="field">
      <label class="label">{{ $t('registration.email') }}</label>
      <div class="control has-icons-left has-icons-right">
        <input v-validate="'required|email'" :class="{'input': true}" name="email" type="email"   v-model="signUpForm.contactDetails.email.address" :placeholder="$t('registration.placeholder.email')">
        <span class="icon is-small is-left">
          <i class="fa fa-envelope"></i>
        </span>
        <span class="icon is-small is-right">
          <i class="fa fa-warning"></i>
        </span>
      </div>
      <p v-show="errors.has('email')" class="help is-danger">{{ errors.first('email') }}</p>
    </div>

    <div class="field">
      <label class="label">{{ $t('registration.mobile') }}</label>
      <div class="control has-icons-left has-icons-right">
        <input class="input" type="text" id="txt_mobNo" @focusout="mobileNoChanged" v-validate="'required|digits:10'" name="mobile" v-model="signUpForm.contactDetails.mobile.number" :placeholder="$t('registration.placeholder.mobile')">
        <span class="icon is-small is-left">
          <i class="fa fa-mobile"></i>
        </span>
      </div>
      <p v-show="errors.has('mobile')" class="help is-danger">{{ errors.first('mobile') }}</p>
    </div>

    <div class="field">
      <label class="label">{{ $t('registration.username') }}</label>
      <div class="control is-expanded has-icons-left has-icons-right">
        <input :class="{'input' : true, 'is-success': !errors.has('username')&& isUsernameAvailable}" type="text" @focusout="checkUsernameAvailability" v-validate="'required'" name="username" v-model="signUpForm.userName" :placeholder="$t('registration.placeholder.username')">
        <span class="icon is-small is-left">
          <i class="fa fa-user"></i>
        </span>
        <span v-if="isUsernameAvailable"  class="icon is-small is-right">
          <i style="color:green;"class="fa fa-check"></i>
        </span>
        <span v-if="!isUsernameAvailable" class="icon is-small is-right">
          <i class="fa fa-times"></i>
        </span>
      </div>
      <p v-show="errors.has('username')" class="help is-danger">{{errors.first('username')}}</p>
      <div v-if="usernameSuggessions != null && !isUsernameAvailable">
        <div>Suggested usernames :
          <a style="margin-left:10px;"v-for="username in usernameSuggessions" @click.prevent="selectSuggestedUsername(username)" :key="username">{{username}}</a>
        </div>
      </div>
    </div>

    <div class="field">
      <label class="label">{{ $t('registration.password') }}</label>
      <div class="columns">
      <div class="control has-icons-left column is-four-fifths">
        <input :class="{'input': true}" id="txt_password" type="password" name="password" v-validate="'required'" @keypress="passwordChanged" v-model="signUpForm.password" :placeholder="$t('registration.placeholder.password')"/>
        <span class="icon is-small is-left" style="margin:3%;">
          <i class="fa fa-unlock-alt"></i>
        </span>
      </div>
      <div class="column control" id="showPass"> 
          <span  class="button" @click="showHidePassword" v-if="!isShowPass">
            <i v-if="!isShowPass" class="fa fa-eye"></i>
            Show
          </span>
           <span  class="button" @click="showHidePassword" v-if="isShowPass">
            <i class="fa fa-eye-slash"></i>
            Hide
          </span>
      </div>
      </div>
        <p v-show="!errors.has('password')" id="strength"></p> 
      <p v-show="errors.has('password')" class="help is-danger"> {{errors.first('password') }}</p>
    </div>

  <div class="field">
      <p class="control">
              By clicking Register You agree to th e <a href="#">Terms of Service</a> and <a href="#"> Privacy Policy</a>.
      </p>
  </div>

    <div class="field is-grouped">
      <p class="control">
         <button type="button" class="button is-primary" @click.prevent="register">{{ $t('registration.buttonType.register') }}</button>
      </p>
      <p class="control">
         <button type="button"  class="button is-light" @click.prevent="cancel">{{ $t('registration.buttonType.cancel') }}</button>
      </p>
    </div>

    </div>
    
    </div>
    </div>
    </div>
    <div class="column"></div>
</div>
</template>

<script>
export default {
  name: 'registerationForm',
  data () {
    return {
      signUpForm: {
        contactDetails: {
          email: {
            address: null
          },
          mobile: {
            number: null,
            verifiedOn: null
          }
        },
        personalDetails: {
          dob: null,
          fullName: null,
          gender: null
        },
        userName: null,
        password: null
      },
      isUsernameAvailable: false,
      usernameSuggessions: null,
      isShowPass: false,
      isEnableRegistration: false
    }
  },
  methods: {
    mobileNoChanged () {
      var txtMobNo = document.getElementById('txt_mobNo')
      var mobRegExp = new RegExp('^[7-9][0-9]{9}$')
      if (!mobRegExp.test(txtMobNo.value)) {
        this.errors.add('mobile', 'Invalid Mobile Number')
      } else {
        this.errors.remove('mobile')
      }
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
    },
    selectSuggestedUsername (selectedUsername) {
      this.isUsernameAvailable = true
      this.signUpForm.userName = selectedUsername
    },
    cancel () {
      this.$router.push({ path: '/' })
    },
    checkUsernameAvailability () {
      var registrationForm = this
      this.$validator
        .validateAll({
          fullname: registrationForm.signUpForm.personalDetails.fullName,
          dob: registrationForm.signUpForm.personalDetails.dob,
          username: registrationForm.signUpForm.userName
        })
        .then(result => {
          if (result) {
            window.axios
              .post(
                'o/isUsernameAvailable',
                registrationForm.signUpForm
              )
              .then(function (response) {
                if (response.status === 200) {
                  registrationForm.isUsernameAvailable = true
                  registrationForm.errors.remove('username')
                  registrationForm.$toasted.info(registrationForm.$t('registration.alerts.USERNAME_AVAILABLE'))
                }
                if (response.status === 208) {
                  registrationForm.isUsernameAvailable = false
                  registrationForm.usernameSuggessions =
                    response.data.usernames
                  // adding error
                  registrationForm.errors.add(
                    'username',
                    'Username Already Taken'
                  )
                  registrationForm.$toasted.info(registrationForm.$t('registration.alerts.USERNAME_NOT_AVAILABLE'))
                }
              })
              .catch(function () {
                registrationForm.isUsernameAvailable = false
                registrationForm.$toasted.error(registrationForm.$t('registration.alerts.FAIL_TO_FIND_USERNAME_AVAILABILITY'))
              })
          } else {
            registrationForm.$toasted.error(registrationForm.$t('registration.alerts.DATA_NOT_VALID'))
          }
        })
    },
    register () {
      var registrationForm = this
      if (this.errors.count() > 0) {
        registrationForm.$toasted.error(registrationForm.$t('registration.alerts.DATA_NOT_VALID'))
        return
      }
      this.$validator.validateAll().then(result => {
        if (result) {
          window.axios.post('o/register', registrationForm.signUpForm)
            .then(function (response) {
              if (response.status === 200) {
                registrationForm.$toasted.success(registrationForm.$t('registration.alerts.REGISTRATION_SUCCESS'))
                registrationForm.$router.push({ path: '/' })
              }
            })
            .catch(function () {
              registrationForm.$toasted.error(registrationForm.$t('registration.alerts.REGISTRATION_FAIL'))
            })
        } else {
          registrationForm.$toasted.error(registrationForm.$t('registration.alerts.DATA_NOT_VALID'))
        }
      })
    }
  },
  mounted: function () {
    this.$refs.txt_fullname.focus()
    var registrationForm = this
    window.axios.post(`/o/getclientconfig`, { 'clientid': registrationForm.$store.getters.getClientId })
          .then(Response => {
            this.CONFIG = Response.data
            registrationForm.$store.commit('SET_CONFIG', Response.data)
          })
  }
}
</script>

<style>

#showPass {
  cursor: pointer;
}
input[type="date"] {
  position: relative;
}
input[type="date"]::-webkit-inner-spin-button {
  -webkit-appearance: none;
  display: none;
}
input[type="date"]::-webkit-calendar-picker-indicator {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  width: auto;
  height: auto;
  color: transparent;
  background: transparent;
}
</style>
