<template>
	<view>
		
		<view class="u-flex u-col-center u-row-between" v-if="lunbo && lunbo.length>0">
			<view class="imglist">
				<swiper class="swiper" :indicator-dots="true" :autoplay="true" style="height: 400rpx;">
					<swiper-item v-for="(item,index) in lunbo" :key="index">
						<image :src="item.img" mode="" class="swiimg"></image>
					</swiper-item>
				</swiper>
			</view>
		</view>
		
		<view class="u-flex-col u-padding-20">
			<text class="u-font-28 u-color-balck6 u-margin-top-10">{{list.title}}</text>
			<view class="u-flex u-margin-top-10 u-row-between">
				<text class="u-font-28 zhuti_color_huang">{{list.money>0?'￥'+list.money:'免费'}} </text>
				<view class="u-flex">
					<text class="u-font-28 u-color-balck9">已报名</text>
					<text class="u-font-28 zhuti_color_huang">{{list.bao}}</text>
					<text class="u-font-28 u-color-balck9">人</text>
				</view>
			</view>
			
			
			<view class="jianju-10 u-margin-top-20"></view>
			<view class="u-flex border_bottom_1" style="padding: 20rpx 0;">
				<text class="u-font-26 u-color-balck3">预约地址</text>
			</view>
			<view class="u-flex border_bottom_1 u-row-between" style="padding: 20rpx 0;">
				<view class="u-flex">
					<i class="icon u-font-28 u-color-balck9">&#xe640;</i>
					<text class="u-font-26 u-color-balck6 u-margin-left-5">{{list.address}}</text>
				</view>
				<view class="u-flex u-row-center u-col-center" style="border-left: 1rpx solid #F6F6F6;" @click="openditu()">
					<text class="u-font-26 zhuti_color">查看地址</text>
				</view>
			</view>
			<view class="u-flex border_bottom_1 u-row-between" style="padding: 20rpx 0;">
				<view class="u-flex">
					<i class="icon u-font-28 u-color-balck9">&#xe601;</i>
					<text class="u-font-26 u-color-balck6 u-margin-left-5">{{list.name}}</text>
				</view>
				<view class="u-flex u-row-center u-col-center" style="border-left: 1rpx solid #F6F6F6;" @click="botel(list.mobile)">
					<text class="u-font-26 zhuti_color">电话咨询</text>
				</view>
			</view>
			<view class="jianju-10"></view>
		</view>
		
		<view class="u-padding-20 u-font-26 u-color-balck6">
			<view class="u-flex zhuti-border-bootom-1 u-row-center u-col-center" style="padding: 20rpx 0;">
				<text class="u-font-26 u-color-balck3">图文介绍</text>
			</view>
			<view style="padding: 20rpx 0;line-height: 40rpx;">
				<u-parse v-if="texter" :content="texter"></u-parse>
			</view>
		</view>
		
		<view style="height: 100rpx;"></view>
		<view class="u-flex tabbat">
			<view @click="backurl()" class="tabbatli u-flex-col u-col-center u-row-center" style="width: 225rpx;">
				<i class="icon u-font-44 zhuti_color">&#xe613;</i>
				<text class="u-font-24 u-margin-top-1 zhuti_color">返回首页</text>
			</view>
			<view v-if="start==0" @click="open()" class="tabbatli u-flex-col u-col-center u-row-center zhuti_bg" style="width: 525rpx;">
				<text class="u-color-white u-font-28 u-margin-top-1">立即预约</text>
			</view>
			<view v-if="start==1" class="tabbatli u-flex-col u-col-center u-row-center" style="width: 525rpx;background-color: #CCCCCC;">
				<text class="u-color-white u-font-28 u-margin-top-1">预约已结束</text>
			</view>
			<view v-if="start==2" class="tabbatli u-flex-col u-col-center u-row-center" style="width: 525rpx;background-color: #CCCCCC;">
				<text class="u-color-white u-font-28 u-margin-top-1">预约已满</text>
			</view>
		</view>
		
		<!--选择弹框  底部选择-->
		<uni-popup ref="popup" type="share">
		    <view class="popbox">
				<view class="u-flex u-padding-20 border-bootom-1 u-row-between u-col-center">
					<view class="u-flex">
						<text class="icon u-font-32 zhuti_color">&#xe75c;</text>
						<text class="u-font-24 u-color-balck9 u-margin-left-5">请选择预约类型</text>
					</view>
					<view class="u-flex" @click="close()">
						<text class="icon u-font-32 u-color-balck9">&#xe66f;</text>
					</view>
				</view>
				<view class="u-flex u-flex-wrap">
					<view class="u-flex u-row-center u-col-center u-padding-10"  v-for="(item,index) in classarr" :key="index">
						<view v-if="item.no>0" @click="xuanst=index" class="zhongbtn u-flex-col u-row-center u-col-center u-margin-bottom-10" :class="xuanst==index?'xuanno u-color-white':'u-color-balck9'">
							<text class="u-font-24 ">{{item.title}}</text>
							<text class="u-font-24  u-margin-top-5">剩余：{{item.no}}</text>
						</view>
						<view v-if="item.no<=0" class="zhongbtn2 u-flex-col u-row-center u-col-center u-margin-bottom-10">
							<text class="u-font-24 ">{{item.title}}</text>
							<text class="u-font-24  u-margin-top-5">预约已满</text>
						</view>
					</view>
				</view>
				<view class="bootcent">
					<view class="bootbtn zhutibtn u-flex u-row-center u-col-center" @click="tiao('order?id='+list.id)">
						<text class="u-font-28 u-color-white">确定</text>
					</view>
				</view>
			</view>
		</uni-popup>
		
	</view>
</template>

<script>
	//#ifdef H5
	var jweixin = require('@/components/jweixin-module/index.js');
	//#endif
	import uParse from '@/components/gaoyia-parse/parse.vue'
	export default {
		components: {
		    uParse
		  },
		data() {
			return {
				endtime:'',
				classarr:[],
				id:0,
				xuanst:-1,
				texter:'',
				token:'',
				list:[],
				lunbo:[],
				start:-1,
				lat:'',
				lng:'',
				opst:0,//
			}
		},
		onLoad(e){
			if(e.id){
				this.id=e.id;
			}
			this.GetToken();
			this.iswx=this.$iswx;
			//#ifdef H5
			var that=this;
			if(this.iswx==2){
				var apiUrl = location.href.split("#")[0];
				uni.request({
					url: that.$puburl+'resource/Wxlogin/getSignPackage',
					data: {
						url: encodeURIComponent(apiUrl),//当前页面的域名
						api: ['openLocation','checkJsApi'],//调用的方法去接口列表里找
					},
					success: function(res) {
						var wxData = res.data.data.data;
						that.wx_sanCode(wxData)
					}
				})
			}
			//#endif
		},
		methods: {
			open() {
				console.log(this.opst);
				//打开时间  还是  直接打开订单
				if(this.opst==1){
					this.tiao('time?id='+this.list.id);
				}else{
					this.tiao('order?id='+this.list.id);
				}
				return false;
				this.$refs.popup.open()
			},
			close() {
				this.$refs.popup.close()
			},
			botel(tel){
				uni.makePhoneCall({
				    phoneNumber: tel
				});
			},
			// openditu(){
			// 	var _this=this;
			// 	var latitude=parseFloat(this.lat);
			// 	var longitude=parseFloat(this.lng);
			// 	uni.openLocation({
			// 		latitude: latitude,
			// 		longitude: longitude,
			// 		name:'地图地址',
			// 		address: _this.list.address,
			// 		success: function () {
			// 			console.log('success');
			// 		},
			// 		fail(e) {
			// 			console.log(JSON.stringify(e));
			// 		}
			// 	});
			// },
			getData(){
				var _this=this;
				this.$api.yuyue_info({id:this.id,token:this.token}).then((res)=>{
					_this.lunbo=res.data.data.lunbo;
					_this.classarr=res.data.data.classarr;
					_this.list=res.data.data.list;
					_this.texter=res.data.data.list.texter;
					_this.start=res.data.data.list.start;
					_this.lat=res.data.data.list.lat;
					_this.lng=res.data.data.list.lng;
					_this.opst=res.data.data.opst;
				})
			},
			GetToken(){
				var _this=this;
				uni.getStorage({
				    key: 'userData',
					success:function (res) {
						_this.token=res.data.vuex_token;
						_this.getData();
					},
					fail() {
						_this.getData();
					}
				});
			},
			openditu(){
				//#ifdef H5 || MP-WEIXIN
				if(this.iswx==2){
					var latitude=parseFloat(this.latitude);
					var longitude=parseFloat(this.longitude);
					var company=this.list.title;
					var address=this.list.address;
					//微信公众号则需要
					jweixin.openLocation({
						latitude: latitude,//目的地latitude
						longitude: longitude,//目的地longitude
						name: company,
						address: address,
						scale: 15//地图缩放大小，可根据情况具体调整
					});
					return false;
				}
				var latitude=parseFloat(this.latitude);
				var longitude=parseFloat(this.longitude);
				var company=this.list.title;
				var address=this.list.address;
				uni.openLocation({
					latitude: latitude,
					longitude: longitude,
					address: address,
					name: company
				});
				//#endif
			},
			//扫码
			wx_sanCode: function(wxData) {
				jweixin.config({
					debug: false, // 开启调试模式,调用的所有api的返回值会在客户端alert出来，若要查看传入的参数，可以在pc端打开，参数信息会通过log打出，仅在pc端时才会打印。
					appId: wxData.appId, // 必填，公众号的唯一标识
					timestamp: wxData.timestamp, // 必填，生成签名的时间戳
					nonceStr: wxData.nonceStr, // 必填，生成签名的随机串
					signature: wxData.signature, // 必填，签名
					jsApiList: [
						"checkJsApi",
						"onMenushareTimeline",
						"onMenushareAppMessage",
						"openLocation"
					], // 必填，需要使用的JS接口列表
				})
				jweixin.ready(function() {
					// config信息验证后会执行ready方法，所有接口调用都必须在config接口获得结果之后，config是一个客户端的异步操作，所以如果需要在页面加载时就调用相关接口，则须把相关接口放在ready函数中调用来确保正确执行。对于用户触发时才调用的接口，则可以直接调用，不需要放在ready函数中。
					jweixin.checkJsApi({
						jsApiList: ['openLocation'], // 需要检测的JS接口列表，所有JS接口列表见附录2,
						success: function(res) {
							console.log('返回信息')
							console.log(res)
							// 以键值对的形式返回，可用的api值true，不可用为false
							// 如：{"checkResult":{"chooseImage":true},"errMsg":"checkJsApi:ok"}
						}
					});
				});
				jweixin.error(function(res){
				  // config信息验证失败会执行error函数，如签名过期导致验证失败，具体错误信息可以打开config的debug模式查看，也可以在返回的res参数中查看，对于SPA可以在这里更新签名。
				console.log(res,'接口验证失败')
				});
			},
			backurl(){
				uni.navigateBack();
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
	.xuanno{
		background-image: linear-gradient(to right, #77FAB8, #23D0A2);
		border:1rpx solid #77FAB8!important;
	}
	.zhongbtn2{
		color: #CCCCCC;
		padding: 8rpx 10rpx;
		border-radius: 5rpx;
		border:1rpx solid #CCCCCC;
	}
	.zhongbtn{
		padding: 8rpx 10rpx;
		border-radius: 5rpx;
		border:1rpx solid #CCCCCC;
	}
	.bootcent{
		padding: 20rpx;
		box-shadow: 0px -4px 5px -3px rgba(20, 20, 20, 0.06);
	}
	.yinyin{
		box-shadow: 0px 10px 20px 0px rgba(20, 20, 20, 0.06);
	}
	.bootbtn{
		border-radius: 10px;
		height: 90rpx;
	}
	.popbox{
		background-color: #FFFFFF;
	}
	
	.jianju-10{
		width: 100%;
		height: 20rpx;
		background-color: #f6f6f6;
	}
	.bobtn{
		background-image: linear-gradient(to right, #00B7ED, #2979FF);
	}
	/**倒计时**/
	.baibtn{
		width: 40rpx;
		height: 40rpx;
		border-radius: 4rpx;
		color: #FC4B54;
		margin-left: 6rpx;
		margin-right: 6rpx;
	}
	.daobtn{
		width: 710rpx;
		height: 80rpx;
		border-top-left-radius: 10rpx;
		border-bottom-right-radius: 10rpx;
	}
	/**倒计时**/
	
	.border_bottom_1{
		border-bottom: 1rpx solid #F6F6F6;
	}
	.imglist{
		width: 750rpx;
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
		height: 400rpx;
	}
	.swiimg{
		width: 750rpx;
		height: 400rpx;
	}
	.tabbat{
		width: 750rpx;
		height: 100rpx;
		position: fixed;
		bottom: 0px;
		left: 0px;
		border-top: 1rpx solid #F6F6F6;
		background-color: #FFFFFF;
	}
	.tabbatli{
		height: 100rpx;
	}
</style>
