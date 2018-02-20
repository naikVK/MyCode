<template>
  <div id="app" v-show="redirection">
    <nav class="navbar is-fixed-top is-light  " role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
          <a class="navbar-item">
            <img src="./assets/logo.png">
          </a>
      </div>
      <div class="navbar-menu">
      <div class="navbar-start">
        <!-- navbar items -->
      </div>

      <div class="navbar-end">
          <div class="navbar-item">
            <div class="field">
            <div class="control">
              <div class="select">
                <select v-model="selectedLanguage" @change="changeLanguage">
                  <option disabled>{{ this.$t('selectLanguage')}}</option>
                  <option>English</option>
                  <option>Marathi</option>
                </select>
              </div>
            </div>
            </div>
          </div>
          <div v-if="isSignedIn" class="navbar-item" @click="mylogout">
            <p class="control">
              <a class="button is-static">
                  Logout
              </a>
            </p>
          </div>
  </div>
</div>
    </nav>
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'app',
  data () {
    return {
      flag: false,
      selectedLanguage: this.$t('selectLanguage')
    }
  },
  methods: {
    changeLanguage () {
      if (this.selectedLanguage === 'English') {
        this.$i18n.set('en')
      } else if (this.selectedLanguage === 'Marathi') {
        this.$i18n.set('mr')
      }
    },
    mylogout () {
      var form = this
      window.MKCLAPI.logout().then(function () {
        form.$toasted.success('logout successfull')
        form.$router.push({path: '/login'})
      }).catch(function () {
        form.$toasted.error('failed to logout')
      })
    }
  },
  computed: {
    isSignedIn: function () {
      // alert(window.MKCLAPI.getToken())
      if (window.MKCLAPI.isLoggedIn()) {
        return true
      } else {
        return false
      }
    },
    redirection: function () {
      var form = this
      if (this.$route.path !== 'This/Is/Never/Gonna/Match') {
        if (!(this.$route.path === '/' || this.$route.path === '/login')) {
          window.MKCLAPI.isSessionPresent().then(function () {
          }).catch(function () {
            form.$router.push({ path: '/login' })
            form.$toasted.info('Session Expired. Please login again.')
          })
        }
      }
      return true
    }
  },
  mounted () {
    this.flag = window.MKCLAPI.isLoggedIn()
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
