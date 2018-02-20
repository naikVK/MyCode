<template>
  <div id="app">
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
          <div v-if="isLoggedIn" class="navbar-item" @click="logout">
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
    logout () {
      var form = this
      window.axios.get('/r/logout').then(function (response) {
        if (response.status === 200) {
          form.$store.commit('REMOVE_TOKEN')
          form.$store.commit('FLUSH_DATA')
          form.$toasted.success(form.$t('logout.alerts.LOGOUT_SUCCESS'))
          form.$router.push({path: '/'})
        }
      }).catch(function () {
        form.$toasted.error(form.$t('logout.alerts.LOGOUT_FAIL'))
      })
    }
  },
  computed: {
    isLoggedIn: function () {
      return this.$store.getters.getIsLoggedIn
    }
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
