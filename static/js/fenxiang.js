export default {
	data() {
		return {
			fenxiang:{
				title:'通源堂预约',
				path:'/pages/index/index',
				imageUrl:'http://www.ohyu.cn/logo.png',
				desc:'通源堂预约系统！',
				content:'通源堂预约系统。'
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