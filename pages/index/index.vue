<template>
	<view class="content">
		
		<view class="u-flex u-col-center u-row-between u-padding-20">
			<view class="imglist">
				<swiper class="swiper" :indicator-dots="true" :autoplay="true">
					<swiper-item v-for="(item,index) in lunbo" :key="index" @click="turl(item.style,item.url)">
						<image :src="item.img" mode="" class="swiimg" style="border-radius: 20rpx;"></image>
					</swiper-item>
				</swiper>
			</view>
		</view>
		
		<!--公告-->
		<view class="notilist u-border-ra20" v-if="noti && noti.length>0">
			<view class="u-flex u-row-between u-row-center u-padding-20 u-col-center">
				<text class="icon zhuti_color u-font-28">&#xe609;</text>
				<view class="noticent">
					<swiper  :vertical="true" :autoplay="true" style="height: 40rpx;line-height: 40rpx;">
						<swiper-item class="u-font-dan-sheng" @click="tiao('../texter/noti_info?id='+item.id)" v-for="(item,index) in noti" :key="index">
							<text class="u-font-28 u-color-balck3 u-font-dan-sheng" style="width: 620rpx;">{{item.title}}</text>
						</swiper-item>
					</swiper>
				</view>
				<text class="icon u-color-balck9 u-font-28">&#xe605;</text>
			</view>
		</view>
		
		<!--小按钮-->
		<view class="u-flex u-flex-wrap u-row-between" style="width:710rpx;background-color: #F6F6F6;">
			<view class="u-flex-col u-col-center  u-margin-top-15" style="width:140rpx;background-color: #FFFFFF;padding: 10rpx;" v-for="(item,index) in kuai" :key="index" @click="turl(item.style,item.url)">
				<image :src="item.img" style="width: 90rpx;height:90rpx"></image>
				<text class="u-font-26 u-color-balck3 u-margin-top-20">{{item.title}}</text>
			</view>
		</view>
		
		<!--推荐信息-->
		<view class="u-flex-col gxcent u-margin-bottom-20 u-border-ra20 u-bg-white u-margin-top-20" v-if="list && list.length>0">
			<view class="u-padding-20 u-flex border-bootom-1">
				<!-- <view class="xiaoli zhuti_bg"></view> -->
				<text class="icon zhuti_color u-font-28">&#xe6be;</text>
				<span class="u-font-28 u-color-balck3 u-margin-left-10">推荐信息</span>
			</view>
			<view v-for="(item,index) in list" :key="index" class="u-padding-20 u-flex border-bootom-1" @click="tiao('../yuyue/info?id='+item.id)">
				<image :src="item.img" style="width: 200rpx;height: 140rpx;border-radius: 20rpx;" mode=""></image>
				<view class="u-flex-col u-margin-left-20 clix u-row-between">
					<text class="u-font-28 u-color-balck3">{{item.title}}</text>
					<text class="u-font-24 u-color-balck6 u-margin-top-10">{{item.address}}</text>
					<view class="u-flex u-row-between">
						<text class="u-font-24 zhuti_color" style="color: #19be6b;">电话:{{item.mobile}}</text>
						<view class="lxbtn">
							<text class="u-font-24 zhuti_color" style="color: #19be6b;">立即预约</text>
						</view>
					</view>
				</view>
			</view>
		</view>
		
		<!--门店地址-->
		<view class="u-flex u-margin-bottom-20" v-if="stores && stores.title && stores.address && stores.phone">
			<view class="u-flex-col u-padding-20 u-bg-white u-border-ra20" style="width: 670rpx;">
				<view class="u-flex">
					<image :src="stores.imgUrl" alt="" style="width: 80rpx;height: 80rpx;border-radius: 40rpx;">
					<view class="u-flex u-margin-left-20">
						<text class="u-font-28 u-color-balck3 u-font-bold">{{stores.title}}</text>
					</view>
				</view>
				<view class="u-flex-col">
					<view class="u-flex u-margin-top-20" v-if="stores.address">
						<text class="icon u-color-balck6 u-font-28">&#xe640;</text>
						<text class="u-font-26 u-color-balck6 u-margin-left-20">{{stores.address}}</text>
					</view>
					<view class="u-flex u-margin-top-20" v-if="stores.time">
						<text class="icon u-color-balck6 u-font-24">&#xe602;</text>
						<text class="u-font-26 u-color-balck6 u-margin-left-20">营业时间：{{stores.time}}</text>
					</view>
				</view>
				<view class="u-flex u-margin-top-20" style="border-top: 2rpx solid #F6F6F6;" v-if="stores.phone && stores.lat">
					<view class="u-flex-1 u-row-center u-col-center u-margin-top-10" style="border-right: 2rpx solid #F6F6F6;height: 60rpx;" @click="openditu()">
						<text class="icon zhuti_color_lan u-font-36">&#xe611;</text>
					</view>
					<view class="u-flex-1 u-row-center u-col-center u-margin-top-10" style="height: 60rpx;" @click="botel(stores.phone)">
						<text class="icon zhuti_color u-font-36">&#xe601;</text>
					</view>
				</view>
			</view>
		</view>
		
		
		<view @click="openwxkf()" v-if="kfurl" class="faguang kefuvie">
			<text class="icon zhuti_color" style="font-size: 50rpx;">&#xe753;</text>
		</view>
		
	</view>
</template>

<script>
	//#ifdef H5
	var jweixin = require('@/components/jweixin-module/index.js');
	//#endif
	export default {
		data() {
			return {
				name:'',
				mobile:'',
				title: 'Hello',
				lunbo:[],
				kuai:[],
				token:'',
				list:[],
				noti:[],
				guang:[],
				classr:[],
				status: 'loadmore',
				page:1,
				count: 10,
				res_count:0,
				wxcode:'',
				appid:'',
				kfurl:'',
				iswx:1,
				uuid:'',
				index:0,
				sharedata:[],
				stores:[],
				identity:[],
				fenxiang:{
					path:'/pages/index/index'
				}
			}
		},
		onLoad(option) {
			//推广者  绑定上级id
			if (option.scene) {
				var scene = decodeURIComponent(option.scene);
				var scene = scene.split("&");
				if(scene){
					var data=scene[0];
					var va=data.split("=");
					if(va){
						var uid=va[1];//推荐人id
						//缓存师傅id
						uni.setStorage({
							key: 'invitation',
							data: {
								'uid' :uid
							}
						});
					}
				}
			}
			if(option.uid){
				var uid=option.uid;
				console.log(uid);
				//缓存师傅id
				uni.setStorage({
					key: 'invitation',
					data: {
						'uid' :uid
					}
				});
			}
			
			//#ifdef H5
			//微信公众号登录
			let href = window.location.href;
			if (href.includes("cn/?code")) {  //url包括 com/?code 证明为从微信跳转回来的
				var url = href.substring(0, href.length - 2); //vue自动在末尾加了 #/ 符号，截取去掉
				var jingPosit = url.indexOf("cn/") + 3; //获取域名结束的位置
				var urlLeft = url.substring(0, jingPosit);//url左侧部分
				var urlRight = url.substring(jingPosit, url.length); //url右侧部分
				var url=urlLeft + "#/" + urlRight;
				window.location = url;//拼接跳转
				console.log(url);return false;
			}
			if (href.includes("com/?code")) {  //url包括 com/?code 证明为从微信跳转回来的
				var url = href.substring(0, href.length - 2); //vue自动在末尾加了 #/ 符号，截取去掉
				var jingPosit = url.indexOf("com/") + 4; //获取域名结束的位置
				var urlLeft = url.substring(0, jingPosit);//url左侧部分
				var urlRight = url.substring(jingPosit, url.length); //url右侧部分
				var url=urlLeft + "#/" + urlRight;
				window.location = url;//拼接跳转
				console.log(url);return false;
			}
			if (href.includes("net/?code")) {  //url包括 net/?code 证明为从微信跳转回来的
				var url = href.substring(0, href.length - 2); //vue自动在末尾加了 #/ 符号，截取去掉
				var jingPosit = url.indexOf("net/") + 4; //获取域名结束的位置
				var urlLeft = url.substring(0, jingPosit);//url左侧部分
				var urlRight = url.substring(jingPosit, url.length); //url右侧部分
				var url=urlLeft + "#/" + urlRight;
				window.location = url;//拼接跳转
				console.log(url);return false;
			}
			if(option.code){
				this.wxcode = option.code;
				this.getopenid();
				return false;
			}
			//end
			//#endif
			
			this.iswx=this.$iswx;
			var that=this;
			//#ifdef H5
			if(this.iswx==2){
				var apiUrl = location.href.split("#")[0];
				uni.request({
					url: that.$puburl+'resource/Wxlogin/getSignPackage',
					data: {
						url: encodeURIComponent(apiUrl),//当前页面的域名
						api: ['scanQRCode','checkJsApi'],//调用的方法去接口列表里找
					},
					success: function(res) {
						var wxData = res.data.data.data;
						that.wx_sanCode(wxData)
					}
				})
			}
			//#endif
			
			this.getData();
			this.iswx=this.$iswx;//1网页 2微信
			this.GetToken();
			
		},
		onPullDownRefresh() {
			this.list=[];
			this.page=1;
			this.status='loadmore';
			this.getData();
			setTimeout(function () {
			    uni.stopPullDownRefresh();
			}, 1000);
		},
		methods: {
			//扫码
			wx_sanCode: function(wxData) {
				var _this=this;
				console.log(wxData.appId);
				jweixin.config({
					debug: false, // 开启调试模式,调用的所有api的返回值会在客户端alert出来，若要查看传入的参数，可以在pc端打开，参数信息会通过log打出，仅在pc端时才会打印。
					appId: wxData.appId, // 必填，公众号的唯一标识
					timestamp: wxData.timestamp, // 必填，生成签名的时间戳
					nonceStr: wxData.nonceStr, // 必填，生成签名的随机串
					signature: wxData.signature, // 必填，签名
					jsApiList: [
						"checkJsApi",
						"updateAppMessageShareData",
						"updateTimelineShareData",
						"openLocation"
					], // 必填，需要使用的JS接口列表
				})
				var apiUrl = window.location.href;
				console.log(apiUrl);
				jweixin.ready(()=> {
					jweixin.updateAppMessageShareData({
						title:_this.sharedata.title, // 分享标题
						desc: _this.sharedata.desc, // 分享描述
						link: apiUrl, // 分享链接，该链接域名或路径必须与当前页面对应的公众号JS安全域名一致
						imgUrl:_this.sharedata.imgUrl, // 分享图标									
						success:function(data){
						  // 设置成功
							console.log('updateAppMessageShareData success:', data);
						},
						fail: function(error) {
							console.log('updateAppMessageShareData error:', error);
						}
					});
					jweixin.updateTimelineShareData({
						title: _this.sharedata.title, // 分享标题
						desc: _this.sharedata.desc, // 分享描述
						link: apiUrl, // 分享链接，该链接域名或路径必须与当前页面对应的公众号JS安全域名一致
						imgUrl: _this.sharedata.imgUrl, // 分享图标
					})
				});
				jweixin.error(function(res){
				  // config信息验证失败会执行error函数，如签名过期导致验证失败，具体错误信息可以打开config的debug模式查看，也可以在返回的res参数中查看，对于SPA可以在这里更新签名。
				console.log(res,'接口验证失败')
				});
			},
			bindPickerChange: function(e) {
				this.index = e.detail.value
			},
			open() {
				this.$refs.popup.open();
			},
			close() {
				this.$refs.popup.close();
			},
			botel(tel){
				uni.makePhoneCall({
				    phoneNumber: tel
				});
			},
			openditu(){
				//#ifdef H5 || MP-WEIXIN
				if(this.iswx==2){
					var latitude=parseFloat(this.stores.lat);
					var longitude=parseFloat(this.stores.lng);
					var company=this.stores.title;
					var address=this.stores.address;
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
				var latitude=parseFloat(this.stores.lat);
				var longitude=parseFloat(this.stores.lng);
				var company=this.stores.title;
				var address=this.stores.address;
				uni.openLocation({
					latitude: latitude,
					longitude: longitude,
					address: address,
					name: company
				});
				//#endif
			},
			openwxkf(){
				var kfurl=this.kfurl;
				var appid=this.appid;
				console.log('打开客服');
				//#ifdef MP-WEIXIN
				wx.openCustomerServiceChat({
				  extInfo: {url: kfurl},
				  corpId: appid,
				  success(res) {
					console.log(res);
				  }
				})
				//#endif
				//#ifdef H5
				this.turl(1,kfurl);
				//#endif
			},
			//#ifndef H5
			GetToken(){
				var _this=this;
				uni.getStorage({
				    key: 'userData',
					success:function (res) {
						_this.token=res.data.vuex_token;
						_this.uid=value.vuex_user.id;
						_this.fenxiang.path='/pages/index/index?uid='+_this.uid;
					}
				});
			},
			//#endif
			//#ifdef H5
			GetToken(){
				var _this=this;
				uni.getStorage({
				    key: 'userData',
					success:function (res) {
						_this.token=res.data.vuex_token;
						_this.uid=value.vuex_user.id;
						_this.fenxiang.path='/pages/index/index?uid='+_this.uid;
					},
					fail() {
						if(_this.iswx==2){
							_this.getUserInfo();
						}
					}
				});
			},
			// 点击授权方法
			getUserInfo() {
				if(this.$wxurl){
					var xurl=this.$wxurl+'?type=2';
				}else{
					var xurl=this.$puburl;
				}
				var url=encodeURI('https://open.weixin.qq.com/connect/oauth2/authorize?appid='+this.$wxappid+'&redirect_uri='+xurl+'&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect');
				window.location.href = url;
				if(this.code){
					//将onLoad中的获取token方法放进来即可
				} else {
					//没有code就去获取code
				}
			},
			getopenid(){
				var _this=this;
				if(this.wxcode){
					uni.showLoading({
						title: '登录中...'
					});
					var uid=0;
					//获取师傅id
					const value = uni.getStorageSync('invitation');
					if (value) {
						if(value.uid){
							uid=value.uid;
						}
					}
					uni.request({
						url: _this.$puburl+'resource/wxh5register',
					    data: {
					        code: this.wxcode,
							uid:uid
					    },
					    method: 'GET',
					    header: {
					        'content-type': 'application/json'
					    },
					    success: (res) => {
							console.log(JSON.stringify(res));
							if(res.data.code==1){
								_this.token=res.data.data.token;
								var token=res.data.data.token;
								var userlist=res.data.data.userlist;
								//console.log('uu==='+userlist);
								//登录成功  跳转到首页
								uni.setStorage({
									key: 'userData',
									data: {
										'vuex_token' :token,
										'vuex_user' : userlist
									},
									success: function () {
										console.log('success');
									}
								});
								window.location = _this.$puburl;//跳转
								_this.getData();
								//_this.wxGetUserInfo(); //需要手动确认
							}else{
								uni.showToast({
									title: res.data.message,
									icon:"none",
									duration: 2000
								});
							}
					        //openId、或SessionKdy存储//隐藏loading
					        uni.hideLoading();
					    },fail(e) {
							console.log(e);
					    	uni.hideLoading();
					    }
					});
				} else {
					uni.showToast({
						icon:'none',
						title:'请先登录'
					})
					//没有code在走一次流程去获取code
				}
			},
			//#endif
			getData(){
				console.log('xxxx');
				this.$api.indeximg().then((res)=>{
					this.lunbo=res.data.data.lunbo;
					this.kuai=res.data.data.kuai;
					this.list=res.data.data.list;
					this.noti=res.data.data.noti;
					this.guang=res.data.data.guang;
					this.appid=res.data.data.appid;
					this.kfurl=res.data.data.kfurl;
					this.sharedata=res.data.data.sharedata;
					this.stores=res.data.data.stores;
					if(res.data.data.sharedata && res.data.data.sharedata.title){
						uni.setNavigationBarTitle({
						    title: res.data.data.sharedata.title
						});
					}
				})
			},
			tiao(url){
				uni.navigateTo({
					url:url
				})
			},
			turl(st,url){
				if(st==1){
					uni.navigateTo({
						url:'weburl?url='+url
					})
				}else if(st==2){
					//#ifdef APP-PLUS
					plus.runtime.openURL(url);
					//#endif
					//#ifdef H5
					window.open(url);
					//#endif
				}else if(st==4){
					uni.navigateTo({
						url:url
					})
				}else if(st==5){
					uni.switchTab({
						url:url
					})
				}
			}
		}
	}
</script>

<style>
	page{
		background-color: #F6F6F6;
	}
	.kefuvie{
		position:fixed;
		right:40rpx;
		bottom:200rpx;
		width:78rpx;
		height:78rpx;background:#ffffff;
		border-radius:40rpx;z-index:9;
		display: flex;
		justify-content: center;
		align-items: center;
	}
	.faguang{
	    box-shadow: 0px 0px 16rpx #cccccc;
	    border-radius: 40rpx;
	}
	.putx{
		width: 500rpx;
		height: 70rpx;
		background-color: #F0F0F0;
		margin-bottom: 15rpx;
		margin-top: 8rpx;
		text-indent: 20rpx;
		font-size: 28rpx;
	}
	.popbox{
		padding: 20rpx;
		background-color: #FFFFFF;
		border-radius: 10rpx;
	}
	.clix{
		width: 490rpx;
		height: 130rpx;
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
	.gxcent{
		width: 710rpx;
	}
	.xiaoli{
		width: 6rpx;
		height: 36rpx;
	}
	.noticent{
		width: 620rpx;
		height: 40rpx;
	}
	.notilist{
		width: 710rpx;
		background-color: #FFFFFF;
	}
	.content {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}
	.an_li{
		width: 177.5rpx;
	}
	.imglist{
		width: 710rpx;
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
	}
	.swiimg{
		width: 710rpx;
		height: 300rpx;
	}
	
	
</style>
