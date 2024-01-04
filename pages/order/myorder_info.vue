<template>
	<view v-if="list && list.id">
		
		<view class="u-padding-20">
			<view class="u-flex-col u-bg-white" style="padding: 0 10px;">
				<view class="u-flex u-col-center u-row-between" style="line-height: 40px;border-bottom: 1px solid #F6F6F6;font-size: 14px;color: #333333;">
					<view class="u-flex">
						<i class="icon zhuti_color u-font-40">&#xe75c;</i>
						<span class="u-margin-left-20">订单信息</span>
					</view>
					<text class="u-font-24 u-color-balck3">ID：{{list.id}}</text>
				</view>
				<view class="">
					<view class="u-flex border-bootom-1" style="padding: 20rpx 0;">
						<image :src="list.img" style="width: 100rpx;height: 100rpx;border-radius: 20rpx;" mode=""></image>
						<view class="u-flex-col u-margin-left-20">
							<view class="u-flex">
								<text class="u-font-26 u-color-balck3 u-margin-left-8 u-font-dan-sheng" style="width: 540rpx;height: 40rpx;line-height: 40rpx;">{{list.title}}</text>
							</view>
							<text class="u-font-26 u-margin-top-20" style="color: #2979ff;">{{list.tit}}</text>
						</view>
					</view>
					<view class="u-flex-col u-margin-top-10">
						<view class="u-flex txtli u-row-between" v-if="list.money>0">
							<text class="u-font-24 u-color-balck3">商品总额</text>
							<text class="u-font-24 u-color-balck3">￥{{list.money}}</text>
						</view>
						<view class="u-flex txtli u-row-between">
							<text class="u-font-24 u-color-balck3">订单状态：</text>
							<view class="u-flex">
								<text class="u-font-24 u-color-balck3">{{list.star}}</text>
							</view>
						</view>
						<view class="u-flex txtli u-row-between" v-if="list.paymo>0">
							<text class="u-font-24 u-color-balck3">实付款：</text>
							<view class="u-flex">
								<text class="u-font-28 zhuti_color_hong u-font-bold">￥{{list.paymo}}</text>
							</view>
						</view>
						<view class="u-flex txtli u-row-between">
							<text class="u-font-24 u-color-balck3">提交时间：</text>
							<view class="u-flex">
								<text class="u-font-24 u-color-balck3">{{list.addtime}}</text>
							</view>
						</view>
					</view>
				</view>
			</view>
		</view>
		
		<view class="" style="padding: 0rpx 20rpx;">
			<view class="u-flex-col u-bg-white" style="padding: 0 10px;">
				<view class="u-flex" style="line-height: 40px;border-bottom: 1px solid #F6F6F6;font-size: 14px;color: #333333;">
					<text class="icon zhuti_color u-font-40">&#xe75c;</text>
					<text class="u-margin-left-20">预约信息</text>
				</view>
				<view class="u-flex u-flex-wrap" style="line-height: 60rpx;color: #333333;font-size: 24rpx;">
					<text style="width: 200rpx;">预约日期</text>
					<text>{{list.riqi}}</text>
				</view>
				<view class="u-flex u-flex-wrap" style="line-height: 60rpx;color: #333333;font-size: 24rpx;">
					<text style="width: 200rpx;">预约时间</text>
					<text>{{list.time}}</text>
				</view>
				<view class="u-flex u-flex-wrap" style="line-height: 60rpx;color: #333333;font-size: 24rpx;" v-for="(item,index) in list.res" :key="index">
					<text style="width: 200rpx;">{{item.name}}</text>
					<text v-if="item.type!=5">{{item.val}}</text>
					<view class="u-flex" v-if="item.type==5">
						<image :src="it" mode="" v-for="(it,index2) in item.val" :key="index2" @click="proimg(it)" style="width: 100rpx;height: 100rpx;"></image>
					</view>
				</view>
			</view>
			
		</view>
		
		<view class="" style="padding: 0rpx 20rpx;margin-top: 20rpx;" v-if="1==2">
			<view class="u-flex-col u-bg-white" style="padding: 0 10px;">
				<view class="u-flex" style="line-height: 40px;border-bottom: 1px solid #F6F6F6;font-size: 14px;color: #333333;">
					<text class="icon zhuti_color u-font-40">&#xe75c;</text>
					<text class="u-margin-left-20">支付信息</text>
				</view>
				<view style="line-height: 60rpx;color: #333333;font-size: 24rpx;" v-if="list.status>1 && list.status<5">支付单号 ： {{list.trade_no}}</view>
				<view style="line-height: 60rpx;color: #333333;font-size: 24rpx;" v-if="list.status>1 && list.status<5">付款时间 ： {{list.paytime}}</view>
			</view>
		</view>
		
		<!--底部-->
		<!--未付款-->
		<view style="height: 100rpx;" v-if="list.status==1"></view>
		<view class="myorderboot u-bg-white u-flex u-row-between" v-if="list.status==1">
			<view @click="qxyu()" class="u-flex u-row-center u-col-center zhuti_border zhuti_color u-margin-left-20" style="width: 300rpx;height: 70rpx;border-radius: 35rpx;">取消订单</view>
			<view  @click="getpay(list.id)" class="u-flex u-row-center u-col-center zhuti_bg u-color-white" style="width: 300rpx;height: 70rpx;border-radius: 35rpx;margin-right: 20rpx;">立即支付</view>
		</view>
		<!--待核销-->
		<view style="height: 100rpx;" v-if="list.status==2"></view>
		<view class="myorderboot u-bg-white u-flex u-row-between" v-if="list.status==2">
			<view @click="istui(list.id)" class="u-flex u-row-center u-col-center zhuti_border zhuti_color u-margin-left-20" style="width: 300rpx;height: 70rpx;border-radius: 35rpx;">
				<text v-if="list.paymo>0">我要退款</text>
				<text v-if="list.paymo==0">我要取消</text>
			</view>
			<view @click="open()" class="u-flex u-row-center u-col-center zhuti_bg u-color-white" style="width: 300rpx;height: 70rpx;border-radius: 35rpx;margin-right: 20rpx;">核销二维码</view>
		</view>
		<!--取消订单-->
		<view style="height: 100rpx;" v-if="list.status==4"></view>
		<view class="myorderboot u-bg-white u-flex u-row-center" v-if="list.status==4">
			<view @click="tiao('../yuyue/info?id='+list.yuyue_id)" class="u-flex u-row-center u-col-center zhuti_bg u-color-white" style="width:600rpx;height: 70rpx;border-radius: 35rpx;margin-right: 20rpx;">
				再次预约
			</view>
		</view>
		<!--订单已完成-->
		<view style="height: 100rpx;" v-if="list.status==3"></view>
		<view class="myorderboot u-bg-white u-flex u-row-center" v-if="list.status==3">
			<view @click="tiao('../yuyue/info?id='+list.yuyue_id)" class="u-flex u-row-center u-col-center zhuti_bg u-color-white" style="width:600rpx;height: 70rpx;border-radius: 35rpx;margin-right: 20rpx;">
				再次预约
			</view>
		</view>
		
		<uni-popup ref="popup" type="center">
		    <view class="popbox">
				<view class="topvix">
					<text class="u-font-28 u-color-balck3">核销凭证</text>
				</view>
				<view class="u-flex u-row-center u-col-center u-padding-20">
					<view class="erlist">
						<text class="u-font-32 u-color-balck6">核销码：NO{{id}}</text>
						<image :src="list.code" mode="" style="width: 360rpx;height: 360rpx;"></image>
						<!-- <canvas v-if="id" id="qrcode" canvas-id="qrcode" style="width: 180px;height: 180px;" /> -->
					</view>
				</view>
				<view class="topvix">
					<text class="u-font-24 u-color-balck3">店铺地址：{{list.address}}</text>
				</view>
				<view v-if="list.paymo>0">
					<view class="topvix">
						<text class="u-font-24 u-color-balck3">商品信息：{{list.title}}</text>
					</view>
					<view class="topvix" style="border:none;">
						<text class="u-font-24 u-color-balck3">实付金额：{{list.paymo}}</text>
					</view>
				</view>
				<view class="topvix" style="border:none;" v-else>
					<text class="u-font-24 u-color-balck3">商品信息：{{list.title}}</text>
				</view>
				
		    </view>
			<view class="u-flex u-row-center u-col-center u-margin-top-30" @click="close()">
				<text class="icon" style="color:#ffffff;font-size: 64rpx;">&#xe66f;</text>
			</view>
		</uni-popup>
		
	</view>
</template>

<script>
	import uQRCode from '@/components/uqrcode/common/uqrcode.js';
	export default {
		data() {
			return {
				id:0,
				navst:0,
				list:[],
				title: 'Hello',
				token:'',
				status: 'loadmore',
				type:0,
				count: 10,
				res_count:0,
				page: 1,
				iswx:1,//网页
			}
		},
		onLoad(e) {
			if(e.id){
				this.id=e.id;
			}
			this.iswx=this.$iswx;
			this.GetToken();
			
		},
		onPullDownRefresh() {
			this.getData();
			setTimeout(function () {
			    uni.stopPullDownRefresh();
			}, 1000);
		},
		onShow() {
			this.list=[];
			this.GetToken();
		},
		onBackPress() {
			uni.redirectTo({
				url:'myorder'
			});
			return true;
		},
		//#ifdef H5
		mounted() {
			if(this.iswx==2){
				if (window.location.href.indexOf("?#") < 0) {
				  window.location.href = window.location.href.replace("#", "?#");
				}
			}
		},
		//#endif
		methods: {
			istui(id){
				var _this=this;
				if(this.list.paymo>0){
					var tit='你正在进行退款操作，确定退款吗？';
					uni.showModal({
						content: tit,
						confirmText: "确定",
						cancelText: "取消",
						cancelColor:"#71ADF5",
						confirmColor:"#71ADF5",
							success: function (res) {
								if (res.confirm) {
									_this.tuikuan(id);
								}
							}
					})
				}else{
					var tit='你正在进行取消操作，确定取消吗？';
					uni.showModal({
						content: tit,
						confirmText: "确定",
						cancelText: "取消",
						cancelColor:"#71ADF5",
						confirmColor:"#71ADF5",
							success: function (res) {
								if (res.confirm) {
									_this.qxyu(id);
								}
							}
					})
				}
			},
			tuikuan(id){
				var _this=this;
				var st=3;
				//#ifdef H5
				if(this.iswx==2){
					var st=1;
				}else{
					var st=2;
				}
				//#endif
				uni.request({
					url: _this.$puburl+'resource/Payco/h5refund',
				    data: {
				        id: id,
						type:st,
						token:_this.token
				    },
				    method: 'GET',
				    header: {
				        'content-type': 'application/json'
				    },
				    success: (res) => {
						if(res.data.code==200){
							uni.showToast({
							    title: res.data.message,
								icon:'none',
							    duration: 2000
							});
							_this.getData();
						}else{
							uni.showToast({
							    title: res.data.message,
								icon:'none',
							    duration: 2000
							});
						}
					}
				})
			},
			open() {
				this.$refs.popup.open();
				//#ifdef H5
				var textxx=this.$puburl+'#/pages/user/hexiao?id='+this.id;
				//#endif
				//#ifdef MP-WEIXIN
				var textxx=this.id+'&1';
				//#endif
				// uQRCode.make({
				//     canvasId: 'qrcode',
				//     componentInstance: this,
				//     size: 180,
				//     margin: 10,
				//     text: textxx,
				//     backgroundColor: '#ffffff',
				//     foregroundColor: '#333333',
				//     fileType: 'png',
				//     errorCorrectLevel: uQRCode.errorCorrectLevel.H
				// })
				// .then(res => {
				//     console.log(res)
				// })
			},
			close() {
				this.$refs.popup.close()
			},
			getData(){
				var _this=this;
				//#ifdef H5
				var type=1;
				//#endif
				//#ifdef MP-WEIXIN
				var type=2;
				//#endif
				this.$api.myorderinfo({'token':this.token,id:this.id,type:type}).then((res)=>{
					_this.list=res.data.data;
					_this.money=res.data.data.money;
					console.log(_this.money);
					uni.setNavigationBarTitle({
					    title: res.data.data.star
					});
				})
			},
			onBridgeReady(id){
				var _this=this;
				uni.request({
					url: _this.$puburl+'resource/Payco/wxh5',
				    data: {
				        id: id,
						st:1,
						token:_this.token
				    },
				    method: 'GET',
				    header: {
				        'content-type': 'application/json'
				    },
				    success: (res) => {
						var data=res.data;
						WeixinJSBridge.invoke(
						   'getBrandWCPayRequest', data,
						   function(res){
							console.log('xsssss='+JSON.stringify(res));
						   if(res.err_msg == "get_brand_wcpay_request:ok" ){
						   // 使用以上方式判断前端返回,微信团队郑重提示：
							   uni.showToast({
								title: '支付成功',
								icon:"none",
								duration: 2000
							   });
							   setTimeout(() => {
								uni.navigateTo({
									url:'../pay/payok?id='+id+'&money='+_this.money
								})
							   },1000)
						         //res.err_msg将在用户支付成功后返回ok，但并不保证它绝对可靠。
						   }else{
							   uni.showToast({
							   	title: '支付失败',
							   	icon:"none",
							   	duration: 2000
							   });
						   }
						});
				    }
				});
			},
			getpay(id){
				//#ifdef H5
				if(this.iswx==2){
					this.onBridgeReady(id);
				}else{
					//网页
					uni.navigateTo({
						url:'../pay/h5pay?id='+id+'&st=1'
					})
				}
				return false;
				//#endif
				var _this=this;
				//#ifdef MP-WEIXIN
				uni.request({
					url: _this.$puburl+'resource/Payco/wxxcx',
				    data: {
				        id: id,
						st:1,
						token:_this.token
				    },
				    method: 'GET',
				    header: {
				        'content-type': 'application/json'
				    },
				    success: (res) => {
						var nonceStr=res.data.data.nonceStr;
						var packages=res.data.data.package;
						var paySign=res.data.data.paySign;
						var timeStamp=res.data.data.timeStamp;
						uni.requestPayment({
						    provider: 'wxpay',
						    timeStamp:timeStamp,
						    nonceStr: nonceStr,
						    package: packages,
						    signType: 'MD5',
						    paySign: paySign,
						    success: function (res) {
						        console.log('success:' + JSON.stringify(res));
								uni.showToast({
									title: '支付成功',
									icon:"none",
									duration: 2000
								});
								setTimeout(() => {
									uni.navigateTo({
										url:'../pay/payok?id='+id+'&money='+_this.money
									})
								},1000)
						    },
						    fail: function (err) {
								uni.showToast({
									title: '支付失败',
									icon:"none",
									duration: 2000
								});
						        console.log('fail:' + JSON.stringify(err));
						    }
						});
						//console.log(JSON.stringify(res));
						
				    }
				});
				//#endif
			},
			qxyu(){
				var _this=this;
				this.$api.quxiao_yuyue({id:this.id,token:this.token}).then(res => {
					if(res.data.code==200){
						uni.showToast({
						    title: res.data.message,
							icon:'none',
						    duration: 2000
						});
						_this.getData();
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
			tiao(url){
				uni.navigateTo({
					url:url
				})
			}
		}
	}
</script>

<style>
	.erlist{
		width: 540rpx;
		height: 400rpx;
		background-color: #FFFFFF;
		margin-top: 10rpx;
		margin-bottom: 10rpx;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
	}
	.topvix{
		padding: 20rpx;
		border-bottom: 2rpx dashed #CCCCCC;
	}
	.popbox{
		width: 600rpx;
		background-color: #F7F7F7;
	}
	
	
	.txtli{
		line-height: 40rpx;
		margin-top: 10rpx;
	}
	.myorderboot{
		height: 100rpx;
		width: 100%;
		position: fixed;
		bottom: 0px;
		left: 0px;
		border-top: 1rpx solid #F6F6F6;
	}
	.mobg{
		background-color: #999999;
	}
</style>
