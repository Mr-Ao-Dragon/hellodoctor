<template>
	<view>
		
		<view class="" v-if="istui==0">
			<view class="u-flex-col u-bg-white u-margin-top-20" style="padding: 0 20rpx;border-radius: 20rpx;">
				<view class="u-flex border-bottom-1" style="padding: 30rpx 0;">
					<view class="u-flex u-row-left u-col-center" style="width: 150rpx;">
						<text class="u-font-28 u-color-balck">联系人</text>
						<text class="u-font-24 zhuti_color_hong u-margin-left-5">*</text>
					</view>
					<view class="u-flex u-row-left u-col-center" style="width: 400rpx;">
						<input type="text" v-model="name" placeholder="请输入联系人" class="u-font-28" style="width: 440rpx;"/>
					</view>
					<view class="u-flex u-row-right u-col-center" style="width: 140rpx;">
						
					</view>
				</view>
				<view class="u-flex border-bottom-1" style="padding: 30rpx 0;">
					<view class="u-flex u-row-left u-col-center" style="width: 150rpx;">
						<text class="u-font-28 u-color-balck">联系电话</text>
						<text class="u-font-24 zhuti_color_hong u-margin-left-5">*</text>
					</view>
					<view class="u-flex u-row-left u-col-center" style="width: 400rpx;">
						<input v-model="phone" type="text" placeholder="请输入联系人电话" class="u-font-28" style="width: 360rpx;"/>
					</view>
					<view class="u-flex u-row-right u-col-center" style="width:160rpx;">
						
					</view>
				</view>
			</view>
			
			<view class="u-flex-col u-bg-white u-margin-top-20" style="padding: 0 20rpx;border-radius: 20rpx;">
				<view class="u-flex u-row-between border-bottom-1" style="padding: 30rpx 0;">
					<view class="u-flex">
						<text class="u-font-28 u-color-balck u-margin-left-5">申请说明：</text>
						<text class="u-font-24 zhuti_color_hong u-margin-left-5">*</text>
					</view>
				</view>
				<view class="u-flex border-bottom-1" style="padding: 30rpx 0;">
					<view class="u-flex u-row-left u-col-center" style="width: 640rpx;">
						<textarea v-model="message" placeholder="请输入申请说明（500字内）" class="u-font-28" style="width: 640rpx;height: 130rpx;" />
					</view>
				</view>
			</view>
			
			<view @click="addinvata()" class="bootbtn zhuti_bg u-flex u-row-center u-col-center">
				<text class="u-font-30 u-color-white">提交</text>
			</view>
		</view>
		
		<view class="u-padding-20 u-flex u-row-center u-col-center" v-if="istui==1">
			<text v-if="list.status==0" class="u-font-28 zhuti_color_hong">待审核..</text>
			<text v-if="list.status==2" class="u-font-28 zhuti_color_hong">审核失败：{{list.error}}</text>
		</view>
		
		
	</view>
</template>

<script>
	export default {
		data() {
			return {
				list:[],
				id:0,
				cid:0,
				token:'',
				name:'',
				phone:'',
				message:'',
				istui:1
			}
		},
		onLoad(e) {
			if(e.id){
				this.id=e.id;
			}
			this.GetToken();
		},
		methods: {
			addinvata(){
				var _this=this;
				this.$api.addinvata({id:this.id,token:this.token,name:this.name,phone:this.phone,msg:this.message}).then(res => {
					if(res.data.code==200){
						uni.showToast({
						    title: res.data.message,
							icon:'none',
						    duration: 2000
						});
						setTimeout(() => {
							// 此为uView的方法，详见路由相关文档
							uni.navigateBack();
						},1200)
						
					}else{
						uni.showToast({
						    title: res.data.message,
							icon:'none',
						    duration: 2000
						});
					}
				})
			},
			GetToken(){
				var _this=this;
				uni.getStorage({
				    key: 'userData',
				    success: function (res) {
						_this.token=res.data.vuex_token;
						_this.getData();
				    }
				});
			},
			getData(){
				var _this=this;
				this.$api.getinvata({token:this.token}).then((res)=>{
					_this.istui=res.data.data.istui;
					_this.list=res.data.data.list;
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
	.bootbtn{
		width: 690rpx;
		height: 90rpx;
		margin-top: 50rpx;
		margin-left: 30rpx;
		border-radius: 20rpx;
	}
</style>
