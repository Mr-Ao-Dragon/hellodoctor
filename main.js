import Vue from 'vue'
import App from './App'

import api from '@/common/vmeitime-http/'

//#ifdef MP-WEIXIN
import fenxiang from './static/js/fenxiang.js'
Vue.mixin(fenxiang)
//#endif

Vue.config.productionTip = false

//Vue.prototype.$puburl = 'http://www.ty.com/'
Vue.prototype.$puburl = 'https://ty.ohyu.cn/'
Vue.prototype.$wxurl = 'http://wxpublic.ohyu.cn/'
Vue.prototype.$wxappid = 'wxd749b5d48bc7e52c';//微信公众号appid跳转使用
Vue.prototype.$api = api

//定义一下h5是网页1  还是  微信2
//#ifdef H5
let status = navigator.userAgent.toLowerCase();
if (status.match(/MicroMessenger/i) == "micromessenger") {
	//console.log('微信浏览器');
	Vue.prototype.$iswx = 2;//微信浏览器
}else{
	//console.log('普通浏览器');
	Vue.prototype.$iswx = 1;//网页
}
//#endif

App.mpType = 'app'

const app = new Vue({
    ...App
})
app.$mount()
