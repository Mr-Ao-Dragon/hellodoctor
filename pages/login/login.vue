<template>
	<view class="wrap">
		<!-- #ifdef H5 -->
		<view class="top"></view>
		<view class="content" v-if="iswx==1">
			<view class="title">欢迎登录</view>
			<view class="u-flex u-row-between u-col-center">
				<input class="u-border-bottom" style="width:430rpx;font-size: 28rpx;" type="number" v-model="tel" placeholder="请输入手机号" />
				<view class="" @tap="submit">
					<text :class="status?'u-color-balck3':'u-color-balck6'">{{sendtit}}</text>
				</view>
			</view>
			<view class="tips">未注册的手机号验证后自动创建账号</view>
			<input class="u-border-bottom" type="number" v-model="code" placeholder="请输入验证码" style="margin-bottom: 30rpx;font-size: 28rpx;"/>
			<button :style="[inputStyle]" class="getCaptcha" @click="regbtn">立即登录</button>
			<view class="alternative" v-if="1==2">
				<view class="password">密码登录</view>
				<view class="issue">遇到问题</view>
			</view>
		</view>
		<view class="content" v-if="iswx==2">
			<view class="title">欢迎登录</view>
		</view>
		<view class="buttom u-margin-top-20">
			<view class="loginType u-flex u-row-center u-col-center">
				
				<view class="wechat item" @click="loginwx" v-if="iswx==2">
					<i class="icon u-color-balck6" style="color: #53C240;font-size: 68rpx;">&#xe617;</i>
					<text class="u-margin-top-10" style="color: #53C240;">微信登录</text>
				</view>
				<!-- <view class="QQ item">
					<view class="icon"><u-icon size="70" name="qq-fill" color="rgb(17,183,233)"></u-icon></view>
					QQ
				</view> -->
			</view>
			<view class="hint">
				登录代表同意
				<text class="link" @click="tiao('../texter/text_xieyi?id=1')">用户协议、</text>
				<text class="link" @click="tiao('../texter/text_xieyi?id=2')">隐私政策，</text>
				并授权使用您的账号信息（如昵称、头像）以便您统一管理
			</view>
		</view>
		<!--#endif-->
		<!-- #ifdef MP-WEIXIN -->
		<view v-if="isCanUse">
			<view>
				<view class='header'>
					<image src='../../static/img/login/wx_login.png'></image>
				</view>
				<view class='content' v-if="yz==1">
					<view>申请获取以下权限</view>
					<text>获得你的手机号信息</text>
				</view>
				<view class='content' v-if="yz==2">
					<view>申请获取以下权限</view>
					<text>获得你的公开信息(昵称，头像等)</text>
				</view>
		
				<button class='bottom' v-if="yz==2" @click="wxGetUserInfo">
					个人信息授权登录
				</button>
				<button class='bottom' open-type="getPhoneNumber" v-if="yz==1" @getphonenumber="onGetPhoneNumber">获取手机号</button>  
				
			</view>
		</view>
		<!-- #endif -->
	</view>
</template>

<script>
//var jweixin = require('@/components/jweixin-module/index.js');
export default {
	data() {
		return {
			yz:2, //1获取手机号  2登录
			tel: '',
			sendtit:'发送验证码',
			code:'',
			time1:false,
			second:60,
			status:true,
			isCanUse: uni.getStorageSync('isCanUse')||true,//默认为true
			wxcode:'',
			token:'',
			encryptedData:'',
			iswx:1,//1网页  2微信
		}
	},
	computed: {
		inputStyle() {
			let style = {};
			if(this.tel && this.code) {
				style.color = "#fff";
				style.backgroundColor = '#19be6b';
			}
			return style; 
		}
	},
	onLoad(option) {
		
		//#ifdef H5
		let href = window.location.href;
		if (href.includes("cn/?code")) {  //url包括 cn/?code 证明为从微信跳转回来的
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
		this.iswx=this.$iswx;
		this.GetToken();
		//#endif
		//#ifdef MP-WEIXIN
		this.login();
		//#endif
		console.log(option);
	},
	onUnload() {
		clearInterval(this.time1);
	},
	methods: {
		GetToken(){
			var _this=this;
			uni.getStorage({
			    key: 'userData',
				success:function (res) {
					_this.token=res.data.vuex_token;
				},
				fail() {
					if(_this.iswx==2){
						_this.getUserInfo();
					}
				}
			});
		},
		loginwx(){
			//#ifdef H5
			this.getUserInfo();
			//#endif
			//#ifdef MP-WEIXIN
			this.wxGetUserInfo();
			//#endif
		},
		//第一授权获取用户信息===》按钮触发
		wxGetUserInfo() {
			let _this = this;
				console.log('xxxx');
				uni.getUserProfile({
					desc: '获取用户信息用于登录',
					lang:'zh_CN',
					success: function(infoRes) {
						_this.encryptedData = infoRes.encryptedData; 
						_this.iv = infoRes.iv;
					console.log(JSON.stringify(infoRes));
		            uni.setStorageSync('isCanUse', false);//记录是否第一次授权  false:表示不是第一次授权
		            _this.updateUserInfo();
		           },
		           fail(res) {
						console.log(JSON.stringify(res));
					}
		       });
			},　　　　　　//登录
			login() {
				let _this = this;
				uni.showLoading({
					title: '登录中...'
				});
				// 1.wx获取登录用户code
		       uni.login({
		           provider: 'weixin',
		           success: function(loginRes) {
						console.log(loginRes);
		               let code = loginRes.code;
		    //            if (!_this.isCanUse) {
		    //                //非第一次授权获取用户信息
		    //                uni.getUserProfile({
		    //                    desc: '获取用户信息用于登录',
		    //                    lang:'zh_CN',
		    //                    success: function(infoRes) { 　　　　　　　　　　　　　　　　　　　　　　//获取用户信息后向调用信息更新方法
		    //                       console.log(JSON.stringify(infoRes));
						// 				_this.encryptedData = infoRes.encryptedData;
						// 				_this.iv = infoRes.iv;
						// 				_this.updateUserInfo();//调用更新信息方法
		    //                    }
						// 	});
						// }	
						var uid=0;
						//获取师傅id
						const value = uni.getStorageSync('invitation');
						if (value) {
							if(value.uid){
								uid=value.uid;
							}
						}
		               //2.将用户登录code传递到后台置换用户SessionKey、OpenId等信息
		               uni.request({
							url: _this.$puburl+'resource/wxregister',
		                   data: {
		                       code: code,
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
									//_this.wxGetUserInfo(); //需要手动确认
								}else if (res.data.code == "200") {
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
									setTimeout(() => {
										// 此为uView的方法，详见路由相关文档
										uni.reLaunch({//信息更新成功后跳转到小程序首页
											url: '/pages/index/index'
										});
									},300)
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
		           },
		       });
		   },
		//向后台更新信息
	   updateUserInfo() {
		   let _this = this;
		   uni.request({
			   url:_this.$puburl+'resource/getinfo',//服务器端地址
			   data: {
				   encryptedData: _this.encryptedData,
				   iv: _this.iv,
					token:_this.token
			   },
			   method: 'POST',
			   header: {
				   'content-type': 'application/json'
			   },
			   success: (res) => {
				   if (res.data.code == "200") {
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
						setTimeout(() => {
							// 此为uView的方法，详见路由相关文档
							uni.reLaunch({//信息更新成功后跳转到小程序首页
								url: '/pages/index/index'
							});
						},300)
				   }else{
						uni.showToast({
							title: res.data.message,
							icon:"none",
							duration: 2000
						});
					}
			   }
			  
		   });
	   },
		wxh5login(){
			jweixin.ready(function(){  
			    // TODO  
				
			});
		},
		// 点击授权方法
		getUserInfo() {
			if(this.$wxurl){
				var xurl=this.$wxurl+'?type=2';
			}else{
				var xurl=this.$puburl;
			}
			window.location.href = 'https://open.weixin.qq.com/connect/oauth2/authorize?appid='+this.$wxappid+'&redirect_uri='+xurl+'&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect'
			console.log(this.code);
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
					url: _this.$puburl+'resource/wxregister',
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
		tiao(url){
			uni.navigateTo({
				url:url
			})
		},
		appLoginWx(){
			uni.login({
			  provider: 'weixin',
			  success: function (loginRes) {
			    console.log(loginRes);
			  }
			});
			
		},
		codest(){
			var _this=this;
			this.sendtit='倒计时60s';
			this.time1 = setInterval(() => {
				let times = --_this.second;
				_this.sendtit='倒计时'+times+'s';
				console.log(times);
				if(times<1){
					clearInterval(_this.time1);
					_this.status=true;
					_this.sendtit='发送验证码';
				}
			}, 1000)
		},
		regbtn(){
			var _this=this;
			if(!this.tel){
				uni.showToast({
				    title: '请输入手机号',
					icon:'none',
				    duration: 2000
				});
				return false;
			}
			if(!this.code){
				uni.showToast({
				    title: '请输入验证码',
					icon:'none',
				    duration: 2000
				});
				return false;
			}
			var uid=0;
			//获取师傅id
			const value = uni.getStorageSync('invitation');
			if (value) {
				if(value.uid){
					uid=value.uid;
				}
			}
			this.$api.phonelogin({mobile:this.tel,code:this.code,uid:uid}).then(res => {
				if(res.data.code==200){
					uni.showToast({
					    title: res.data.message,
						icon:'none',
					    duration: 2000
					});
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
					setTimeout(() => {
						// 此为uView的方法，详见路由相关文档
						uni.switchTab({
							url: '/pages/index/index',
						});
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
		submit() {
			var _this=this;
			if(!this.status){
				uni.showToast({
				    title: '请不要频繁发送',
					icon:'none',
				    duration: 2000
				});
				return false;
			}
			if(!this.tel){
				uni.showToast({
				    title: '请输入手机号',
					icon:'none',
				    duration: 2000
				});
				return false;
			}
			// 调用getInfo接口
			this.$api.smssend({mobile: this.tel,'type':'login'}).then(res => {
				//console.log(JSON.stringify(res));
				if(res.data.code==200){
					uni.showToast({
					    title: res.data.message,
						icon:'none',
					    duration: 2000
					});
					_this.codest();
					_this.status=false;
				}else{
					uni.showToast({
					    title: res.data.message,
						icon:'none',
					    duration: 2000
					});
				}
			})
		}
	}
};
</script>

<style lang="scss" scoped>
	/**授权样式**/
	/*#ifdef MP-WEIXIN*/
	.header {
		margin: 90rpx 0 90rpx 50rpx;
		border-bottom: 1px solid #ccc;
		text-align: center;
		width: 650rpx;
		height: 300rpx;
		line-height: 450rpx;
	}
	.header image {
		width: 200rpx;
		height: 200rpx;
	}
	.content {
		margin-left: 50rpx;
		margin-bottom: 90rpx;
	}
	.content text {
		display: block;
		color: #9d9d9d;
		margin-top: 40rpx;
	}
	.bottom {
		border-radius: 80rpx;
		margin: 70rpx 50rpx;
		font-size: 35rpx;
	}
	/*#endif*/
.wrap {
	font-size: 28rpx;
	.content {
		width: 600rpx;
		margin: 80rpx auto 0;

		.title {
			text-align: left;
			font-size: 60rpx;
			font-weight: 500;
			margin-bottom: 100rpx;
		}
		input {
			text-align: left;
			margin-bottom: 10rpx;
			padding-bottom: 6rpx;
		}
		.tips {
			color: #666666;
			margin-bottom: 30rpx;
			margin-top: 18rpx;
			font-size: 24rpx;
		}
		.getCaptcha {
			background-color: #19be6b;
			color: #FFFFFF;
			border: none;
			font-size: 30rpx;
			padding: 12rpx 0;
			
			&::after {
				border: none;
			}
		}
		.alternative {
			color: #FF7474;
			display: flex;
			justify-content: space-between;
			margin-top: 30rpx;
		}
	}
	.buttom {
		.loginType {
			display: flex;
			padding: 80rpx 80rpx 80rpx 80rpx;
			justify-content:center;
			
			.item {
				display: flex;
				flex-direction: column;
				align-items: center;
				color: #FF7474;
				font-size: 28rpx;
			}
		}
		
		.hint {
			padding: 20rpx 40rpx;
			font-size: 24rpx;
			color: #444444;
			
			.link {
				color: #007AFF;
			}
		}
	}
}
</style>
