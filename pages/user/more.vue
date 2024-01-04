<template>
	<view>
		
		<view class="u-flex u-flex-col u-col-top u-padding-10">
			<text class="u-margin-top-10 u-font-28 u-color-balck3">基本信息</text>
			<text class="u-margin-top-10 u-font-24 u-color-balck9">运营者</text>
			<text class="u-margin-top-10 u-font-28 u-color-balck3">{{company}}</text>
			<text class="u-margin-top-10 u-font-24 u-color-balck9">平台名称</text>
			<text class="u-margin-top-10 u-font-28 u-color-balck3">{{name}}</text>
			<text class="u-margin-top-10 u-font-24 u-color-balck9">服务及数据由以下网址提供</text>
			<text class="u-margin-top-10 u-font-28 u-color-balck3">{{url}}</text>
		</view>
		
		<view class="u-flex u-flex-col u-col-top u-padding-10">
			<text class="u-margin-top-10 u-font-24 u-color-balck9">服务声明</text>
			<text class="u-margin-top-10 u-font-26 u-color-balck6" style="line-height: 25px;">本系统服务由运营者向用户提供，运营者对本系统服务信息内容、数据资料及其运营行为的真实性、合法性及有效性承担全部责任。</text>
		</view>
		
		<view class="u-flex u-flex-col u-col-center u-padding-10">
			<text class="u-margin-top-10 u-font-24 u-color-balck9">© 海之心通用预约系统提供技术支持</text>
		</view>
		
	</view>
</template>

<script>
	export default {
		data() {
			return {
				name:'',
				version:'',
				company:'',
				url:''
			}
		},
		onLoad() {
			this.url=this.$puburl;
			this.getData();
		},
		methods: {
			getData() {
				var _this=this;
				this.$api.getsite().then(res => {
					_this.name=res.data.data.name;
					_this.company=res.data.data.company;
					_this.version=res.data.data.version;
				})
			},
			tiao(url){
				uni.navigateTo({
					url:url
				})
			}
		}
	}
</script>

<style>
	.bootx{
		width: 750rpx;
		height: 50rpx;
		display: flex;
		justify-content: center;
		align-items: center;
		margin-top: 100rpx;
	}
</style>
