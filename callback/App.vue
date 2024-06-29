<script setup>
import axios from 'axios'
import UrlParse from 'url-search-params-polyfill'


const url = new UrlParse(window.location.href, true)
const code = url.query.code
axios({
      url: 'https://' + import.meta.env.VITE_API_HOSTNAME + '/' + import.meta.env.VITE_API_VERSION + '/user/login',
      method: 'GET',
      params: {
        code: code
      },
    }
).then(function (response){
  const json = JSON.parse(response.data)
  localStorage.setItem('access_token', json.access_token)
  localStorage.setItem('expires_in', String(json.expires_in))
})
window.location.href = '/index.html'
</script>

<template>

</template>

<style scoped>

</style>