<script setup>
import weChatLogin from 'vue-wxlogin';

// 使用defineComponent和useStyle进行组件导入和使用（如果vue-wxlogin需要这样的用法）
// 但通常第三方组件直接import后在模板中使用即可，无需额外注册
// 注意：下面这行不是必须的，具体取决于vue-wxlogin的使用文档
// const WxLogin = defineComponent(weChatLogin);

const theme = window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
</script>

<template>
  <we-chat-login
      :appid="import.meta.env.VITE_APPID"
      :scope="'snsapi_userinfo'"
      :redirect-uri="encodeURI('https://' + window.location.hostname + '/callback')"
      :theme="theme"
  >
  </we-chat-login>
</template>

<style scoped>
/* 你的样式代码 */
</style>
