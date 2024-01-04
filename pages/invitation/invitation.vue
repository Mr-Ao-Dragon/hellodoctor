<template>
	<view>
		
		<view class="u-flex-col" v-if="isshen==1">
			<view class="toptop u-flex-col">
				<text class="p1">我的邀请</text>
				<text class="p2 u-margin-top-10">{{tuno}}人</text>
				<view class="d4" ></view>
			</view>
			
			<view class="u-flex u-padding-20 u-row-center u-col-center u-margin-top-20">
				<text style="font-size:26rpx;color:#19be6b;">累计贡献</text>
				<text class="u-margin-left-10" style="font-size:26rpx;color:#666666;">{{zmo}}</text>
				<text class="u-margin-left-10" style="font-size:26rpx;color:#19be6b;">元</text>
			</view>
			
			<view class="u-flex u-padding-20">
				<view class="u-flex-col u-flex-1">
					<view class="u-flex u-row-between" @click="tiao('list?type=1')">
						<text class="u-font-26 u-color-balck6">我邀请数</text>
						<view class="u-flex">
							<text class="u-font-26 u-color-balck6">{{tuno}}</text>
							<text class="icon u-font-26 u-margin-left-10" style="color:#19be6b;">&#xe623;</text>
						</view>
					</view>
					<view class="u-padding-30 u-flex-col zhuti_bg2 u-border-ra10 u-margin-top-20 u-row-center u-col-center">
						<text class="u-font-24 u-color-white">我邀请累计贡献</text>
						<text class="u-font-24 u-color-white u-margin-top-20">{{tumo}}</text>
					</view>
				</view>
				<view class="u-flex-col u-flex-1 u-margin-left-20" style="margin-left:30rpx;">
					<view class="u-flex u-row-between" @click="tiao('list?type=2')">
						<text class="u-font-26 u-color-balck6">下级邀请数</text>
						<view class="u-flex">
							<text class="u-font-26 u-color-balck6">{{sunno}}</text>
							<text class="icon u-font-26 u-margin-left-10" style="color:#19be6b;">&#xe623;</text>
						</view>
					</view>
					<view class="u-padding-30 u-flex-col zhuti_bg2 u-border-ra10 u-margin-top-20 u-row-center u-col-center">
						<text class="u-font-24 u-color-white">下级邀请累计贡献</text>
						<text class="u-font-24 u-color-white u-margin-top-20">{{sunmo}}</text>
					</view>
				</view>
			</view>
			
			<view class="u-flex u-padding-20 u-row-center u-col-center">
				<text class="u-font-28 u-color-balck6">----</text>
				<text class="u-font-28 u-color-balck6 u-margin-left-20">邀请奖励规则</text>
				<text class="u-font-28 u-color-balck6 u-margin-left-20">----</text>
			</view>
			
			<view class="yaocent">
				<rich-text :nodes="msg" style="line-height:24px;color:#666666;"></rich-text>
			</view>
			
			<view style="height: 60px;width: 750rpx;"></view>
			
			<view class="bootvv u-flex u-row-center u-col-center">
				<view class="bootright zhuti_bg" @click="tiao('code')">
					<text>生成邀请卡</text>
				</view>
			</view>
		</view>
		
		
		<view class="u-flex" v-if="isshen==0">
			<view class="u-margin-top-20">
				<view class="yaocent">
					<rich-text :nodes="msg" style="line-height:24px;color:#666666;"></rich-text>
				</view>
			</view>
			<view class="bootvv u-flex u-row-center u-col-center">
				<view class="bootright zhuti_bg" @click="tiao('shenqing')">
					<text>申请成为推广员</text>
				</view>
			</view>
		</view>
		
	</view>
</template>

<script>
	export default {
		data() {
			return {
				data:[],
				token:'',
				reg:'',
				regmo:0,
				tuno:0,
				sunno:0,
				sunmo:0,
				tumo:0,
				zmo:0,
				msg:'',
				isshen:0
			}
		},
		onLoad: function(){
			this.GetToken();
		},
		methods: {
			GetToken(){
				var _this=this;
				uni.getStorage({
				    key: 'userData',
					success:function (res) {
						 _this.token=res.data.vuex_token;
						 _this.getData();
					},
					fail() {
						//清理本地
						uni.removeStorage({
							key: 'userData'
						});
						uni.navigateTo({
							url:'/pages/login/login'
						})
						return false;
					}
				});
			},
			getData(){
				var _this=this;
				this.$api.getinvitation({token:this.token}).then((res)=>{
					_this.isshen=res.data.data.isshen;
					_this.msg=res.data.data.msg;
					_this.tuno=res.data.data.tuno;
					_this.sunno=res.data.data.sunno;
					_this.tumo=res.data.data.tumo;
					_this.sunmo=res.data.data.sunmo;
					_this.zmo=res.data.data.zmo;
				})
			},
			tiao(url){
				if(!this.token){
					uni.navigateTo({
						url:'/pages/login/login'
					})
					return false;
				}
				uni.navigateTo({
					url:url
				})
			}
		}
		
	}
</script>

<style>
	page{
		background: #FFFFFF;
	}
	.u-padding-30{
		padding:30rpx;
	}
	.bootvv{
		width: 90%;
		height: 60px;
		position: fixed;
		bottom: 0px;
		left: 5%;
		background: #FFFFFF;
	}
	.bootright{
		width: 80%;
		height: 40px;
		float: right;
		text-align: center;
		line-height: 40px;
		font-size: 14px;
		border-radius: 20px;
		color: #FFFFFF;
		margin-top: 10px;
	}
	.yaocent{
		width: 710rpx;
		min-height: 30px;
		background: #FFFFFF;
		line-height: 20px;
		font-size: 13px;
		margin-bottom: 70px;
		margin-left: 20rpx;
	}
	.toptop{
		width: 100%;
		height: 110px;
		background: linear-gradient(to bottom, #19BE6B, #6FB995);
		text-align: center;
		color: #FFFFFF;
		position: relative;
	}
	.p1{
		line-height: 50px;
		font-size: 14px;
	}
	.p2{
		line-height: 20px;
		font-size:42px;
	}
	.d4{
	 position: absolute;
	 bottom:0px;
	 left: 50%;
	 margin-left: -5px;
	 float: left;
	 width: 0; 
	 height: 0;
	 border-width: 10px;
	 border-style: solid;
	 border-color: transparent #FFFFFF transparent transparent;
	 transform: rotate(90deg); /*顺时针旋转90°*/
	}
</style>
