<script>
import request from './api/request'
const wxLoginData = {
  appid: 'wx277005e156d46f0a',//填写appid
  redirect_uri: 'https://'+window.location.hostname+'/index.html',//填写这个网页的地址
  response_type: 'code',
  scope: 'snsapi_userinfo',
}
const loginLink = `https://open.weixin.qq.com/connect/oauth2/authorize?appid=${wxLoginData.appid}&redirect_uri=${wxLoginData.redirect_uri}&response_type=${wxLoginData.response_type}&scope=${wxLoginData.scope}#wechat_redirect`

export default {
  async onload(options) {
    
  },
  onShow() {
    const token = uni.getStorageSync('token')
    const expries_in = uni.getStorageSync('expries_in')
    if (!token || !expries_in || expries_in < new Date().getTime()) {
      console.log("未登录")
      window.location.href = loginLink
    }
  }
}
</script>

<style>
@import "static/css/common.css";
@import "static/css/u-ui.css";
@import url("/components/gaoyia-parse/parse.css");
::-webkit-scrollbar {
  display: none;
  width: 0 !important;
  height: 0 !important;
  -webkit-appearance: none;
  background: transparent;
}
@font-face {
  font-family: 'iconfont';
  src: url('static/icon/iconfont.woff2?t=1705233337736') format('woff2'),
       url('static/icon/iconfont.woff?t=1705233337736') format('woff'),
       url('static/icon/iconfont.ttf?t=1705233337736') format('truetype');
}
.iconfont {
  font-family: "iconfont" !important;
  font-size: 16px;
  font-style: normal;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

</style>
