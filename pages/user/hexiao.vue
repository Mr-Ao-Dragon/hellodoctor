<template>
	<view class="u-padding-20">
		
		<view class="" v-if="status>0">
			<view class="centxs">
				<view class="topvix">
					<text class="u-font-32 u-color-balck6">预约码：NO{{list.id}}</text>
				</view>
				<view class="topvix" style="border: none;">
					<text class="u-font-24 u-color-balck3">预约活动：{{list.title}}</text>
				</view>
			</view>
			
			<view class="centxs u-margin-top-30">
				<view class="topvix u-flex" v-for="(item,index) in list.res" :key="index">
					<view class="" style="width: 200rpx;">
						<text class="u-font-24 u-color-balck3">{{item.name}}</text>
					</view>
					<text class="u-font-24 u-color-balck3" v-if="item.type!=5">{{item.val}}</text>
					<view class="u-flex" v-if="item.type==5">
						<image :src="it" mode="" v-for="(it,index2) in item.val" :key="index2" @click="proimg(it)" style="width: 100rpx;height: 100rpx;"></image>
					</view>
					
					<!-- <text class="u-font-24 u-color-balck3">姓名</text>
					<text class="u-font-24 u-color-balck3">{{list.name}}</text> -->
				</view>
				<view class="topvix u-flex">
					<view class="" style="width: 200rpx;">
						<text class="u-font-24 u-color-balck3">预约日期</text>
					</view>
					<text class="u-font-24 u-color-balck3">{{list.y_data}}</text>
				</view>
				<view class="topvix u-flex">
					<view class="" style="width: 200rpx;">
						<text class="u-font-24 u-color-balck3">预约时间</text>
					</view>
					<text class="u-font-24 u-color-balck3">{{list.y_time}}</text>
				</view>
				<view class="topvix u-flex" style="border: none;">
					<view class="" style="width: 200rpx;">
						<text class="u-font-24 u-color-balck3">订单创建时间</text>
					</view>
					<text class="u-font-24 u-color-balck3">{{list.addtime}}</text>
				</view>
			</view>
			
			<view style="text-align: center;" v-if="status==200">
				<view class="tbtn zhuti_bg" @click="add_hexiao()">
					<text class="u-font-28 u-color-white">立即核销</text>
				</view>
			</view>
			<view style="text-align: center;margin-top: 20rpx;" v-if="status!=200">
				<text class="u-font-26 u-color-balck6">{{msg}}</text>
			</view>
			
		</view>
		
		<view style="text-align: center;" v-if="status==0">
			<text class="u-font-26 u-color-balck6">{{msg}}</text>
		</view>
		
	</view>
</template>

<script>
	export default {
		data() {
			return {
				id:8,
				list:[],
				res:[],
				token:'',
				status:0,
				msg:''
			}
		},
		onLoad(e){
			if(e.id){
				this.id=e.id;
			}
			this.GetToken();
		},
		onShow() {
			if(this.id){
				this.GetToken();
			}
		},
		methods: {
			proimg(url){
				var ll=[];
				ll[0]=url;
				// 预览图片
				uni.previewImage({
					current:1,
					urls: ll
				});
			},
			getData(){
				var _this=this;
				this.$api.xiaodata({'token':this.token,id:this.id}).then((res)=>{
					_this.list=res.data.data;
					_this.msg=res.data.message;
					_this.status=res.data.code;
				})
			},
			add_hexiao(){
				var _this=this;
				this.$api.hexiao({id:this.id,token:this.token}).then((res)=>{
					uni.showToast({
						title: res.data.message,
						icon:"none",
						duration: 2000
					});
					_this.getData();
				})
			},
			GetToken(){
				var _this=this;
				uni.getStorage({
				    key: 'userData',
				    success: function (res) {
						_this.token=res.data.vuex_token;
						_this.getData();
				    },
					fail() {
						uni.navigateTo({
							url:'/pages/login/login'
						})
					}
				});
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
	page{
		background-color: #F6F6F6;
	}
	.tbtn{
		margin-top: 40rpx;
		border-radius: 10rpx;
		height: 90rpx;
		display: flex;
		justify-content: center;
		align-items: center;
	}
	.topvix{
		padding: 20rpx;
		border-bottom: 2rpx dashed #CCCCCC;
	}
	.centxs{
		background-color: #FFFFFF;
		padding: 20rpx;
		border: 2rpx dashed #CCCCCC;
	}
</style>
