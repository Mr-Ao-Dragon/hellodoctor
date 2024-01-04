<template>
	<view>
		
		<view class="usertop zhuti_bg" @click="tiao('userinfo')" style="overflow: hidden;border-bottom-left-radius:20px;border-bottom-right-radius:20px;">
			<view class="u-flex u-col-center u-row-center" style="width:100%;margin-top:25px;">
				<image :src="headimg?headimg:'../../static/img/user/mo.png'" style="width: 50px;height:50px;border-radius:50%;border:3px solid #FFFFFF;">
				<view class="u-flex-col u-margin-left-10" style="width:65%;">
					<text class="u-font-28 u-font-bold u-color-white">{{name?name:'请登录'}}</text>
					<text class="u-font-24 u-color-white u-margin-top-10">UID：{{id?id:'0'}}</text>
				</view>
				<!-- <view style="width:10%;text-align:right;">
					<i class="icon u-font-48 u-color-white">&#xe606;</i>
				</view> -->
			</view>
		</view>
		
		<view class="u-padding-20" style="padding-bottom:0px;position: absolute;top: 200rpx;width: 710rpx;">
			<view class="u-flex-col u-bg-white" style="border-radius:10px;">
				<view class="u-flex u-row-between" style="border-bottom:1px solid #f6f6f6;padding:15px 20px;"  onclick="turl('{php echo $this->createMobileUrl('myorder',array('st'=>0))}')">
					<view class="u-flex">
						<text class="icon u-font-42 zhuti_color">&#xe63c;</text>
						<text class="u-font-24 u-color-balck3 u-margin-left-5">全部订单</text>
					</view>
					<view class="u-flex"  @click="tiao('../order/myorder')">
						<text class="u-font-24 u-color-balck9">查看</text>
						<text class="icon u-color-balck9 u-font-26 u-margin-left-5">&#xe623;</text>
					</view>
				</view>
				<view class="u-flex u-padding-20">
					<view class="u-flex-col u-flex-1 u-row-center u-col-center u-rela" @click="tiao('../order/myorder?type=1')">
						<text class="icon u-font-42 zhuti_color">&#xe60e;</text>
						<text class="u-font-26 u-color-balck9 u-margin-top-10">待付款</text>
						<view class="u-abso quan u-flex u-row-center u-col-center" v-if="dayno>0">
							<text class="u-font-22 u-color-white">{{dayno}}</text>
						</view>
					</view>
					<view class="u-flex-col u-flex-1 u-row-center u-col-center u-rela" @click="tiao('../order/myorder?type=2')">
						<text class="icon u-font-42 zhuti_color">&#xe62d;</text>
						<text class="u-font-26 u-color-balck9 u-margin-top-10">待核销</text>
						<view class="u-abso quan u-flex u-row-center u-col-center" v-if="heno>0">
							<text class="u-font-22 u-color-white">{{heno}}</text>
						</view>
					</view>
					<view class="u-flex-col u-flex-1 u-row-center u-col-center" @click="tiao('../order/myorder?type=3')">
						<text class="icon u-font-42 zhuti_color">&#xe6a6;</text>
						<text class="u-font-26 u-color-balck9 u-margin-top-10">已完成</text>
					</view>
					<view class="u-flex-col u-flex-1 u-row-center u-col-center" @click="tiao('../order/myorder?type=4')">
						<text class="icon u-font-42 zhuti_color">&#xe625;</text>
						<text class="u-font-26 u-color-balck9 u-margin-top-10">已取消</text>
					</view>
				</view>
			</view>
		</view>
		
		<view class="userlist u-flex-col u-margin-top-20" style="margin-top: 190rpx;">
			<view class="userli u-flex u-row-between" @click="tiao('../invitation/invitation')" v-if="istui==1">
				<view class="iconlist u-flex u-row-center u-col-center">
					<text class="icon u-font-42 zhuti_color">&#xe604;</text>
				</view>
				<view class="user_cent">
					<text class="u-font-28 u-color-balck3">推荐有礼</text>
				</view>
				<text class="icon u-color-balck9">&#xe623;</text>
			</view>
			<view class="userli u-flex u-row-between u-bg-white" @click="tiao('../pay/paylist')">
				<view class="iconlist u-flex u-row-center u-col-center">
					<i class="icon u-font-44 zhuti_color">&#xe63c;</i>
				</view>
				<view class="user_cent">
					<span class="u-font-28 u-color-balck3">支付记录</span>
				</view>
				<text class="icon u-color-balck9">&#xe623;</text>
			</view>
			<view class="userli u-flex u-row-between" @click="tiao('about')">
				<view class="iconlist u-flex u-row-center u-col-center">
					<text class="icon u-font-38 zhuti_color">&#xe60b;</text>
				</view>
				<view class="user_cent">
					<text class="u-font-28 u-color-balck3">关于我们</text>
				</view>
				<text class="icon u-color-balck9">&#xe623;</text>
			</view>
			<view v-if="xiao==1" class="userli u-flex u-row-between" @click="tiao('../hexiao/list')">
				<view class="iconlist u-flex u-row-center u-col-center">
					<text class="icon u-font-38 zhuti_color">&#xe6fd;</text>
				</view>
				<view class="user_cent">
					<text class="u-font-28 u-color-balck3">核销记录</text>
				</view>
				<text class="icon u-color-balck9">&#xe623;</text>
			</view>
			
		</view>
		<view style="height: 30rpx;"></view>
		
	</view>
</template>

<script>
	export default {
		data() {
			return {
				xiao:0,
				title:'Hello',
				headimg:'',
				name:'请登录',
				id:0,
				token:'',
				heno:0,
				dayno:0,
				istui:0
			}
		},
		onShow() {
			this.GetToken();
		},
		methods: {
			GetToken(){
				var _this=this;
				uni.getStorage({
				    key: 'userData',
					success:function (res) {
						_this.id=res.data.vuex_user.id;
						_this.name=res.data.vuex_user.name;
						_this.headimg=res.data.vuex_user.headimg;
						_this.token=res.data.vuex_token;
						_this.getData();
					}
				});
			},
			saoma(){
				if(!this.token){
					uni.navigateTo({
						url:'/pages/login/login'
					})
					return false;
				}
				if(this.xiao==0){
					uni.showToast({
						title: '没有权限',
						icon:"none",
						duration: 2000
					});
					return false;
				}
				// 只允许通过相机扫码
				uni.scanCode({
				    onlyFromCamera: true,
				    success: function (res) {
						console.log('条码内容：' + res.result);
						var str=res.result;
						//拆分id和type出来 跳转
						let arr = str.split('&');
						var id=arr[0];
						if(arr[1]==1){
							uni.navigateTo({
								url:'hexiao?id='+id
							})
						}
				    }
				});
			},
			getData() {
				var _this=this;
				this.$api.user_info({token: this.token}).then(res => {
					_this.dayno=res.data.data.dayno;
					_this.heno=res.data.data.heno;
					_this.istui=res.data.data.istui;
					//_this.xiao=res.data.data.xiao;
				})
			},
			tiao(url){
				if(!this.token){
					//清理本地
					uni.removeStorage({
						key: 'userData'
					});
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
		background-color: #F6F6F6;
	}
	.quan{
		width: 36rpx;
		height: 36rpx;
		border-radius: 20rpx;
		top: -16rpx;
		right: 40rpx;
		background-color: #FF5722;
	}
	.xfvie{
		width: 720rpx;
		height: 160rpx;
		background: #FFFFFF;
		position: absolute;
		left:15rpx;
		top:200rpx;
		border-radius: 20rpx;
	}
	.zxian{
		width: 4rpx;
		height: 80rpx;
		background: #f6f6f6;
	}
	.viptop{
		height: 100rpx;
		background-color: #FFFFFF;
	}
	.topcenttop{
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		padding: 20rpx 20rpx;
		border-bottom: 1rpx solid #F6F6F6;
	}
	.topcent{
		width: 710rpx;
		margin-top: 20rpx;
		margin-bottom: 20rpx;
		margin-left: 20rpx;
		background-color: #FFFFFF;
		border-radius: 10rpx;
	}
	
	.usertop{
		width: 750rpx;
		height: 280rpx;
		border-bottom-left-radius: 20rpx;
		border-bottom-right-radius: 20rpx;
		background-image:url('/static/img/user/zbg.png');
		background-repeat:repeat-x;
	}
	.userli{
		height: 60px;
		border-bottom: 1rpx solid #F6F6F6;
		padding: 0 10px;
	}
	.userlist{
		width: 710rpx;
		margin-left: 20rpx;
		background-color: #FFFFFF;
		border-radius: 10rpx;
	}
	.iconlist{
		width: 30px;
	}
	.user_cent{
		width: 86%;
	}
	
</style>
