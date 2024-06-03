<view
	v-if="typeof accessToken !== 'undefined' && typeof expiresIn !== 'undefined' && expiresIn < Math.floor((new Date()).valueOf() / 1000) && !ready">
	<!-- 当accessToken和expiresIn都定义，且expiresIn过期且ready为false时渲染的内容 -->
	<wxlogin :appid="wx295f40fae672dd46" :scope="snsapi_userinfo"
		:redirect_uri=`https://${window.location.hostname}/jump`></wxlogin>
</view>
<script>
	import wxlogin from 'vue-wxlogin';
	Vue.component('wxlogin', wxlogin);
	uni.setStorageSync('backUrl', window.location.href)
	export default {
		async onload(options) {},
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