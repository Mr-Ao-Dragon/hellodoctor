<template>
	<view>
		
		<div class="u-content">
			<div style="padding: 10px;border-bottom: 1px solid #F7F7F7;">
				<h2 style="font-size: 14px;margin: 0 0 10px 0;font-weight: 600;">{{info.title}}</h2>
				<p style="margin: 0;font-size:10px;">{{info.addtime}}</p>
			</div>
			<div v-if="content" style="color: #333333;padding: 10px;font-size: 14px;">
				<u-parse :content="content"></u-parse>
			</div>
		</div>
		
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
				title: 'Hello',
				info:[],
				id:0,
				st:1,
				content:''
			}
		},
		onLoad(e) {
			if(e.id){
				this.id=e.id;
			}
			this.getData();
		},
		methods: {
			//zong_imglist
			getData(){
				this.$api.noti_info({id:this.id,st:this.st}).then((res)=>{
					this.info=res.data.data;
					this.content=res.data.data.text;
					if(res.data.data && res.data.data.title){
						uni.setNavigationBarTitle({
						    title: res.data.data.title
						});
					}
				})
			}
		}
	}
</script>

<style>
	
</style>
