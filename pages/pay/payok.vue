<template>
	<view>
		
		<view class="topvie zhuti_bg u-flex u-row-center u-col-center">
			<i class="icon u-font-56 u-color-white" style="font-size: 68rpx;">&#xe72e;</i>
			<view class="u-flex-col u-margin-left-20">
				<text class="u-font-36 u-color-white" v-if="money>0">支付成功</text>
				<text class="u-font-36 u-color-white" v-if="money==0">预约成功</text>
				<text class="u-font-24 u-color-white u-margin-top-10">感谢您的预约</text>
			</view>
		</view>
		
		<view class="centvie u-flex-col">
			<text class="u-font-32 u-color-balck3" style="line-height: 90rpx;" v-if="money>0">￥{{money}}</text>
			<view class="u-flex centli">
				<text class="u-font-26 u-color-balck6">订单编号：</text>
				<text class="u-font-26 u-color-balck3">{{id}}</text>
			</view>
			<view class="u-flex centli">
				<text class="u-font-26 u-color-balck6">下单时间：</text>
				<text class="u-font-26 u-color-balck3">{{time}}</text>
			</view>
			<view class="u-flex centli">
				<text class="u-font-26 u-color-balck6">订单状态：</text>
				<text class="u-font-26 u-color-balck3" v-if="money>0">支付成功</text>
				<text class="u-font-26 u-color-balck3" v-if="money==0">预约成功</text>
			</view>
		</view>
		
		<view class="u-flex u-padding-20" style="width: 640rpx;margin-left: 55rpx;line-height: 50rpx;">
			<text class="u-font-26 zhuti_color_hong">请在预定时间内凭核销码前往，若有事无法前往，请提前取消订单。</text>
		</view>
		
		<view class="paybtn zhuti_bg" @click="tiaoord()">
			<text class="u-font-28 u-color-white">查看订单</text>
		</view>
		
		<view class="paybtn" style="margin-top: 30rpx;border:1rpx solid #cccccc;" @click="tiaoindex()">
			<text class="u-font-28 u-color-balck3">返回首页</text>
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
				img:'',
				money:0,
				id:0,
				time:'',
				st:1
			}
		},
		onLoad(e) {
			if(e.id){
				this.id=e.id;
			}
			if(e.money){
				this.money=e.money;
			}
			if(e.st){
				this.st=e.st;
			}
			this.GetToken();
			this.getNow();
		},
		onBackPress() {
			uni.redirectTo({
			    url: '../order/myorder'
			});
			return true;
		},
		methods: {
			getNow() {
			   var now = new Date(),
			   y = now.getFullYear(),
			   month = now.getMonth(),
			   d = now.getDate(),
			   h = now.getHours(),
			   m = now.getMinutes(),
			   s = now.getSeconds();
			   this.time=y + "-" + (month + 1) + "-" + d + " " + h + ":" + m + ":" + s;
			},
			GetToken(){
				var _this=this;
				uni.getStorage({
				    key: 'userData',
				    success: function (res) {
						_this.token=res.data.vuex_token;
				    }
				});
			},
			tiao(url){
				uni.navigateTo({
					url:url
				})
			},
			tiaoord(){
				if(this.st==1){
					uni.navigateTo({//信息更新成功后跳转到小程序首页
						url: '/pages/order/myorder_info?id='+this.id
					});
					return false;
				}else if(this.st==4){
					uni.navigateTo({//信息更新成功后跳转到小程序首页
						url: '/pages/vip/viplist?id='+this.id
					});
					return false;
				}
				
			},
			tiaoindex(){
				uni.switchTab({
					url:'../index/index'
				})
			}
		}
	}
</script>

<style>
	.paybtn{
		width: 400rpx;
		height: 80rpx;
		border-radius: 45rpx;
		display: flex;
		justify-content: center;
		align-items: center;
		margin-top: 160rpx;
		margin-left: 175rpx;
	}
	.centli{
		width: 600rpx;
		margin-left: 20rpx;
		line-height: 70rpx;
	}
	.centvie{
		width: 640rpx;
		margin-top: -60rpx;
		margin-left: 55rpx;
		background-color: #FFFFFF;
		display: flex;
		justify-content: center;
		align-items: center;
		border-radius: 20rpx;
	}
	.topvie{
		height: 300rpx;
	}
	.u-font-56{
		font-size: 56rpx;
	}
</style>
