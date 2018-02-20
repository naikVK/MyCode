<template>
  <div>
      <h1>EMAIL VERIFYING
          </h1>
      </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'EmailVerification',
  data () {
    return {
      email: '',
      result: ''
    }
  },
  mounted: function () {
    axios.post(`http://localhost:8080/server/o/login/validateuser`, {})
          .then(Response => {
            console.log('response=' + Response.data)
            this.email = Response.data
          })
  }
}
</script>

<style>

</style>
