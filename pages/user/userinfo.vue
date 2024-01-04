<template>
	<view>
		<view class="jianju-10"></view>
		<view class="stopdiv">
			<text class="teleft">头像</text>
			<image v-if="list.headimg" :src="list.headimg" mode=""class="phoimg teright"></image>
		</view>
		<view class="stopdiv" >
			<text class="teleft">昵称</text>
			<view v-if="list.name" style="line-height: 45px;height: 45px;font-size: 14px;float: right;margin-right: 15rpx;">
				<text>{{list.name}}</text>
				<text class="icon u-color-balck9 u-font-26 u-margin-left-5">&#xe61b;</text>
			</view>
		</view>
		<view class="stopdiv">
			<text class="teleft">账号ID</text>
			<text class="teright" style="color: #CBCBCB;" v-if="list && list.id">
				{{list.id}}
			</text>
		</view>
		
		<view class="jianju-10"></view>
		
		
		<view class="stopdiv" @click="tiao('../texter/text_xieyi?id=1')">
			<text class="teleft">平台协议</text>
			<text class="teright"><text class="icon">&#xe623;</text></text>
		</view>
		<view class="stopdiv" @click="tiao('../texter/text_xieyi?id=2')">
			<text class="teleft">隐私政策</text>
			<text class="teright"><text class="icon">&#xe623;</text></text>
		</view>
		
		<!--#ifdef H5-->
		<view class="jianju-10"></view>
		<view class="botdivxx" @tap="outlo()">
			安全退出
		</view>
		<!--#endif-->
		
		
	</view>
</template>

<script>
export default {
	data() {
		return {
			show:false,
			token: '',
			vuex_user:[],
			list:[],
			name:'',
			headimg:'',
			id:''
		}
	},
	onLoad() {
		this.GetToken();
	},
	methods: {
		tiao(url){
			uni.navigateTo({
				url:url
			})
		},
		open() {
			this.$refs.popup.open();
		},
		close() {
			this.$refs.popup.close();
		},
		GetToken(){
			var _this=this;
			uni.getStorage({
			    key: 'userData',
				success:function (res) {
					console.log(res);
					_this.vuex_user=res.data.vuex_user;
					_this.id=res.data.vuex_user.id;
					_this.name=res.data.vuex_user.name;
					_this.headimg=res.data.vuex_user.headimg;
					_this.token=res.data.vuex_token;
					_this.getData();
				},fail() {
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
		outlo(){
			var th=this;
			uni.showModal({
				content: "你正在进行账号退出操作，确定退出当前用户吗？",
				confirmText: "确定",
				cancelText: "取消",
				cancelColor:"#71ADF5",
				confirmColor:"#71ADF5",
					success: function (res) {
						if (res.confirm) {
							uni.removeStorage({
								key: 'userData'
							});
							//解析数据
							uni.showToast({
							    title: '退出成功',
								icon:'none',
							    duration: 2000
							});
							uni.reLaunch({
								url: '/pages/login/login',
							});
							
						}
					}
			})
		},
		getData() {
			this.$api.user_info({token: this.token}).then(res => {
				this.list=res.data.data;
			})
		}
	}
};
</script>

<style lang="scss" scoped>
	.qdbtn{
		width: 120rpx;
		height: 56rpx;
		border-radius: 10rpx;
		display: flex;
		justify-content: center;
		align-items: center;
		margin-right: 20rpx;
	}
	.rinput{
		width: 300rpx;
		height: 60rpx;
		margin-right: 20rpx;
		margin-left: 20rpx;
		font-size: 28rpx;
		color: #666666;
		border-bottom: 1rpx solid #f0f0f0;
	}
	.phoimg{
		width: 35px;
		height: 35px;
		margin-top: 5px;
		border-radius: 50%;
	}
	.stopdiv {
		width: 750rpx;
		height: 90rpx;
		border-bottom: 1px solid #F7F7F7;
		background: #FFFFFF;
		line-height: 90rpx;
		font-size:28rpx;
		display: flex;
		justify-content: space-between;
		align-items: center;
	}
	.teleft{
		color: #333333;
		margin-left: 2%;
	}
	.teright{
		color: #666666;
		line-height: 90rpx;
		margin-right: 2%;
		text-align: right;
	}
	.botdivxx{
		width: 750rpx;
		margin: auto;
		line-height: 100rpx;
		text-align: center;
		color: #333;
		background: #FFFFFF;
		font-size:28rpx;
	}
	page{
		background-color: #F0F0F0;
	}
</style>
