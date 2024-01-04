<template>
	<view class="u-content">
		<view style="padding: 10px;border-bottom: 1px solid #F7F7F7;" v-if="title">
			<h2 style="font-size: 14px;margin: 0 0 10px 0;font-weight: 600;">{{title}}</h2>
			<p style="margin: 0;font-size:10px;">{{addtime}}</p>
		</view>
		<view v-if="content" style="color: #333333;padding: 10px;font-size: 14px;">
			<u-parse v-if="content" :content="content"></u-parse>
		</view>
	</view>
</template>

<script>
	import uParse from '@/components/gaoyia-parse/parse.vue'
	export default {
		components: {
		    uParse
		  },
		data() {
			return {
				content: '',
				title:'',
				addtime:'',
				id:0
			}
		},
		onLoad(e) {
			if(e.id){
				this.id=e.id;
				this.getData();
			}
		},
		methods: {
			getData(){
				this.$api.noti_info({id:this.id,st:2}).then(res => {
					this.title=res.data.data.title;
					this.addtime=res.data.data.addtime;
					this.content=res.data.data.texter;
					uni.setNavigationBarTitle({
					    title: res.data.data.title
					});
				})
			}
		}
	}
</script>

<style lang="scss" scoped>
	
</style>