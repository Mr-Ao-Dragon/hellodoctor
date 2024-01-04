<template>
  <h3>学习Vue</h3>
  <p>{{ msg }}</p>
  <p>{{ hello }}</p>
</template>

<script>
export default {
  data() {
    return {
      msg: 'Hello Vue!',
      hello: 'Hello World!'
    }
  }
}
</script>

<style scoped>
h3 {
  color: red;
}
</style>
