<template>
	<view class="u-flex-col u-row-center u-col-center">
		
		<view class="u-flex-col gxcent">
			<view v-for="(item,index) in list" :key="index" class="u-padding-20 u-margin-top-20 u-border-ra20 u-bg-white u-flex border-bootom-1" @click="tiao('../yuyue/info?id='+item.id)">
				<image :src="item.img" style="width: 200rpx;height: 140rpx;border-radius: 20rpx;" mode=""></image>
				<view class="u-flex-col u-margin-left-20 clix u-row-between" style="width: 300rpx;">
					<text class="u-font-28 u-color-balck3">{{item.title}}</text>
					<text class="u-font-24 u-color-balck6 u-margin-top-10">{{item.address}}</text>
					<view class="u-flex u-row-between">
						<text class="u-font-24 zhuti_color" style="color: #19be6b;">电话:{{item.mobile}}</text>
					</view>
				</view>
				<view class="lxbtn">
					<text class="u-font-24 zhuti_color" style="color: #19be6b;">立即预约</text>
				</view>
			</view>
		</view>
		
	</view>
</template>

<script>
	export default {
		data() {
			return {
				title: 'Hello',
				list:[],
				id:0
			}
		},
		onLoad(e) {
			if(e.id){
				this.id=e.id;
			}
			this.getData();
		},
		methods: {
			tiao(url){
				uni.navigateTo({
					url:url
				})
			},
			getData(){
				this.$api.yuelist({id:this.id}).then((res)=>{
					this.list=res.data.data.list;
					if(res.data.data.title){
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
	page{
		background-color: #F6F6F6;
	}
	.gxcent{
		width: 710rpx;
	}
	.lxbtn{
		width: 140rpx;
		height: 48rpx;
		border: 1rpx solid #19be6b;
		display: flex;
		justify-content: center;
		align-items: center;
		border-radius: 10rpx;
	}
</style>