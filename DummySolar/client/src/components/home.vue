<template>
  <div id="home">
    <h2>Welcome</h2>
    <div class="columns">
      <div class="column">
         <p class="button is-primary" @click="getData">getData</p> 
        </div  
      <div class="column">
      </div> 
      <div class="column">
        <p class="button is-info" @click="logout">Logout</p>
      </div>
    </div> 
 </div>
</template>

<script>
export default {
  name: 'home',
  methods: {
    logout () {
      var form = this
      window.MKCLAPI.logout().then(function () {
        form.$toasted.success('logout successfull')
        form.$router.push({path: '/login'})
      }).catch(function () {
        form.$toasted.error('failed to logout')
      })
    },
    getData () {
      var form = this
      window.axios.get('/server/r/getData')
          .then(Response => {
            if (Response.status === 200) {
              form.$toasted.success('success')
            }
          }).catch(function (err) {
            console.log(err)
            form.$toasted.error('UnAuthorized Access')
          })
    }
  },
  mounted () {
  }
}
</script>

<style>

</style>
