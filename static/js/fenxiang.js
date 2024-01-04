export default {
	data() {
		return {
			fenxiang:{
				title:'海之心通用预约系统',
				path:'/pages/index/index',
				imageUrl:'http://www.ohyu.cn/logo.png',
				desc:'海之心通用预约系统！',
				content:'海之心专做预约系统，多种预约系统成品。'
			}
		}
	},
	onShareAppMessage(res) {
		return {
			title: this.fenxiang.title,
			path: this.fenxiang.path,
			imageUrl: this.fenxiang.imageUrl,
			desc: this.fenxiang.desc,
			content: this.fenxiang.content
		}
	},
	onShareTimeline(res) {
		return {
			title: this.fenxiang.title,
			path: this.fenxiang.path,
			imageUrl: this.fenxiang.imageUrl,
			desc: this.fenxiang.desc,
			content: this.fenxiang.content
		}
	}
}