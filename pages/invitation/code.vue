<template>
	<view>
		
		<view class="tutu" v-if="img">
			<image :src="img" mode="widthFix"></image>
		</view>
		
		<view class="tubtn"></view>
		
		<view class="bootvv u-flex u-row-center u-col-center">
			<view class="bootright zhuti_bg" @tap="save()">
				<text>保存邀请二维码</text>
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
				img:''
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
			save() {
				console.log('url:',this.img);
				// #ifdef H5
				uni.showToast({
					title: '请直接保存',
					icon: 'none',
					duration: 2200
				});
				// #endif
				
				var url=this.img;
				// #ifdef MP-WEIXIN
				uni.downloadFile({
					url: url,
					success: function(image) {
						console.log('图片信息：', JSON.stringify(image));
						uni.saveImageToPhotosAlbum({
							filePath: image.tempFilePath,
							success: function() {
								console.log('save success');
								uni.showToast({
									title: '图片保存成功',
									icon: 'none',
									duration: 2200
								});
							}
						});
					}
				});
				// #endif
			},
			getData(){
				var _this=this;
				//#ifdef H5
				var type=1;
				//#endif
				//#ifdef MP-WEIXIN
				var type=2;
				//#endif
				this.$api.addcode({token:this.token,type:type}).then((res)=>{
					_this.img=res.data.data.img;
				})
			}
		}
	}
</script>

<style>
	page{
		background: #FFFFFF;
	}
	.bootvv{
		width: 750rpx;
		height: 60px;
		position: fixed;
		bottom: 0px;
		left: 0rpx;
		background: #FFFFFF;
	}
	.bootleft{
		width: 340rpx;
		height: 40px;
		float: left;
		text-align: center;
		line-height: 40px;
		font-size: 14px;
		color: #F44F01;
		border: 1px solid #F44F01;
		border-radius: 20px;
		margin-top: 10px;
		margin-left: 25rpx;
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
		margin-right: 25rpx;
	}
	.tutu{
		width: 750rpx;
		background: #FFFFFF;
		margin: auto;
		overflow: hidden;
	}
	.tutu image{
		width: 100%;
	}
	.tubtn{
		width: 100px;
		height: 30px;
		font-size: 14px;
		color: #FFFFFF;
		text-align: center;
		line-height: 30px;
		border-radius: 10px;
		margin:20px auto;
	}
</style>