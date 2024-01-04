<script>
	export default {
		onLaunch: function() {
			console.warn('© ohyu.cn 海之心通用预约系统 1.0');
			console.warn('%c严禁任何单位或个人利用本系统从事任何违法违规的活动。', 'color: red');
		},
		onShow: function() {
			console.warn(
			  '%c这不是一个自由软件！在未得到官方有效书面许可的前提下禁止对程序代码进行修改和使用\n%c任何企业和个人不允许对程序代码以任何形式任何目的再发布\n%c开发不易，请通过正规渠道获得本系统',
			  'color: blue', 'color: red', 'color: green' // CSS Style
			)
		},
		onHide: function() {
			console.log('App Hide')
		}
	}
</script>

<style>
	/*每个页面公共css */
	@import "static/css/common.css";
	@import "static/css/u-ui.css";
	/**富文本样式**/
	@import url("/components/gaoyia-parse/parse.css");
	/**隐藏横向滚动条**/
	::-webkit-scrollbar {
	    display: none;  
	    width: 0 !important;  
	    height: 0 !important;  
	    -webkit-appearance: none;  
	    background: transparent;  
	}
</style>
