<template>
	<view style="padding: 0 30rpx;">
		
		<view class="u-flex-col u-bg-white u-margin-top-20" style="border-radius: 20rpx;">
			<!-- <view  class="u-padding-20 u-flex u-col-center" style="border-bottom: 1rpx solid #F4F5F7;">
				<view class="" style="width: 150rpx;">
					<text class="u-font-28 u-color-balck3">姓名</text>
				</view>
				<view class="u-margin-left-20" style="width: 300rpx;">
					<input type="text" class="u-font-28 u-color-balck3" placeholder="请输入您的姓名"/>
				</view>
				<view class="" style="width: 150rpx;">
					
				</view>
			</view> -->
			<view v-for="(item,index) in form" :key="index" >
				<!--文本框-->
				<view class="u-flex-col border-bootom" style="padding: 0 20rpx;">
					<view class="u-flex border-bottom-1" style="padding: 30rpx 0;" v-if="item.type==1">
						<view class="u-flex u-row-left u-col-center" style="width: 160rpx;">
							<text class="u-font-26 u-color-balck">{{item.name}}</text>
							<text class="u-font-24 zhuti_color_hong u-margin-left-5" v-if="item.mandatory==1">*</text>
						</view>
						<view class="u-flex u-row-left u-col-center u-margin-left-20" style="width: 410rpx;">
							<input type="text" v-model="item.value" :placeholder="item.title" class="u-font-26" style="width: 440rpx;"/>
						</view>
						<view class="u-flex u-row-right u-col-center" style="width: 100rpx;">
							<text class="icon zhuti_color_hui u-font-32" v-if="item.value" @click="item.value=''">&#xe66f;</text>
						</view>
					</view>
				</view>
				<!--文本域-->
				<view class="u-flex-col border-bootom" style="padding: 0 20rpx;" v-if="item.type==2">
					<view class="u-flex u-row-between" style="padding: 30rpx 0;">
						<view class="u-flex">
							<text class="u-font-26 u-color-balck">{{item.name}}：</text>
							<text class="u-font-24 zhuti_color_hong u-margin-left-5" v-if="item.mandatory==1">*</text>
						</view>
					</view>
					<view class="u-flex border-bottom-1" style="padding: 30rpx 0;">
						<view class="u-flex u-row-left u-col-center" style="width: 640rpx;">
							<textarea v-model="item.value" :placeholder="item.title" class="u-font-28" style="width: 640rpx;height: 130rpx;" />
						</view>
					</view>
				</view>
				<!--单选-->
				<view class="u-flex-col border-bootom" style="padding: 0 20rpx;" v-if="item.type==3">
					<view class="u-flex border-bottom-1" style="padding: 30rpx 0;">
						<view class="u-flex" style="width: 160rpx;">
							<text class="u-font-28 u-color-balck u-margin-left-5">{{item.name}}</text>
							<text class="u-font-24 zhuti_color_hong u-margin-left-5" v-if="item.mandatory==1">*</text>
						</view>
						<view class="u-flex u-margin-right-20 u-margin-top-10" v-for="(ite,index2) in item.arr" :key="index2">
							<text class="u-font-28 u-margin-right-5">{{ite}}</text>
							<text class="icon zhuti_color_lan u-font-36" v-if="item.value!=ite" @click="danxuanval(ite,index,index2)">&#xe8bb;</text>
							<text class="icon zhuti_color_lan u-font-36" v-if="item.value==ite">&#xe6c5;</text>
						</view>
					</view>
				</view>
				<!--多选-->
				<view class="u-flex-col border-bootom" style="padding: 0 20rpx;" v-if="item.type==4">
					<view class="u-flex border-bottom-1" style="padding: 30rpx 0;">
						<view class="u-flex" style="width: 160rpx;">
							<text class="u-font-28 u-color-balck u-margin-left-5">{{item.name}}</text>
							<text class="u-font-24 zhuti_color_hong u-margin-left-5" v-if="item.mandatory==1">*</text>
						</view>
						<view class="u-flex u-flex-wrap u-margin-left-20">
							<checkbox-group class="u-flex u-flex-wrap" @change="checkboxChange" id="25" style="width: 500rpx;">
								<label class="u-flex u-flex-wrap u-margin-top-10" v-for="(ite,index2) in item.arr" :key="index2">
									<view>
										<checkbox :value="index+'|'+ite"/>
									</view>
									<view>{{ite}}</view>
								</label>
							</checkbox-group>
						</view>
					</view>
				</view>
				
				<!--单选下拉框-->
				<view class="u-flex-col border-bootom" style="padding: 0 20rpx;" v-if="item.type==6">
					<view class="u-flex border-bottom-1  u-row-between" style="padding: 30rpx 0;">
						<view class="u-flex" style="width: 170rpx;">
							<text class="u-font-26 u-color-balck u-margin-left-5">{{item.name}}</text>
							<text class="u-font-24 zhuti_color_hong u-margin-left-5" v-if="item.mandatory==1">*</text>
						</view>
						<picker @change="bindPickerChange" :id="index" :value="indexx" :range="item.arr">
						    <view class="u-flex u-row-between" style="width: 500rpx;">
						    	<view class="u-font-28 u-color-balck3">{{item.arr[indexx]}}</view>
								<text class="icon zhuti_color_hui u-font-28 u-margin-left-20">&#xe623;</text>
						    </view>
						</picker>
					</view>
				</view>
				
				<!--图片上传-->
				<view class="u-flex-col border-bootom" style="padding: 0 20rpx;" v-if="item.type==5">
					<view class="u-flex u-row-between" style="padding: 30rpx 0;">
						<view class="u-flex">
							<text class="u-font-26 u-color-balck">{{item.name}}：</text>
							<text class="u-font-24 zhuti_color_hong u-margin-left-5" v-if="item.mandatory==1">*</text>
						</view>
					</view>
					<view class="u-flex border-bottom-1" style="padding: 20rpx 0;">
						<view v-for="(itex,index3) in item.imageList" :key="index3" style="margin-right: 10rpx;">
							<image :src="itex" mode="" style="width: 100rpx;height: 100rpx;border-radius: 10rpx;"></image>
						</view>
						<view class="u-flex u-row-center u-col-center upbtn" @click="uploadimg(index)">
							<text class="icon zhuti_color_hui u-font-42">&#xe620;</text>
						</view>
					</view>
				</view>
			</view>
			<!-- <view class="u-padding-20 u-flex u-col-center u-row-between" style="border-bottom: 1rpx solid #F4F5F7;">
				<view class="" style="width: 150rpx;">
					<text class="u-font-28 u-color-balck3">性别</text>
				</view>
				<picker @change="bindPickerChange" :value="index" :range="array">
					<text class="u-font-28 u-color-balck3">{{array[index]}}</text>
					<text class="icon u-color-balck9 u-font-26" style="margin-left: 430rpx;">&#xe623;</text>
				</picker>
			</view> -->
			<!-- <view class="u-padding-20 u-flex u-col-center" style="border-bottom: 1rpx solid #F4F5F7;">
				<view class="" style="width: 150rpx;">
					<text class="u-font-28 u-color-balck3">手机</text>
				</view>
				<view class="u-margin-left-20" style="width: 300rpx;">
					<input type="text" class="u-font-28 u-color-balck3" placeholder="请输入您的手机号"/>
				</view>
				<view class="" style="width: 150rpx;text-align: right;">
				</view>
			</view> -->
		</view>
		
		<view class="u-flex-col u-bg-white u-margin-top-20" style="border-radius: 20rpx;">
			<view class="u-padding-20 u-flex u-col-center u-row-between" style="border-bottom: 1rpx solid #F4F5F7;">
				<view>
					<text class="u-font-28 u-color-balck u-font-bold">已选择预约</text>
				</view>
			</view>
			<view class="u-padding-20 u-flex u-col-top">
				<image :src="list.img" style="width: 80rpx;height: 80rpx;" mode=""></image>
				<view class="u-flex-col u-margin-left-20" style="width: 580rpx;">
					<view class="u-flex u-row-between u-col-center">
						<text class="u-font-28 u-color-balck3 u-font-bold">{{list.title?list.title:''}}</text>
						<text class="u-font-28 u-color-balck6" v-if="list.money>0">￥{{list.money?list.money:'0'}}</text>
					</view>
				</view>
			</view>
		</view>
		
		<view class="u-flex-col u-bg-white u-margin-top-20" style="border-radius: 20rpx;" v-if="riqi && wek && time">
			<view class="u-padding-20 u-flex u-col-center u-row-between" style="border-bottom: 1rpx solid #F4F5F7;">
				<view class="" style="width: 150rpx;">
					<text class="u-font-28 u-color-balck3">到店时间</text>
				</view>
				<view class="u-margin-left-20" style="width: 300rpx;">
					<text class="u-font-26 u-color-balck6">{{riqi}}</text>
					<text class="u-font-26 u-color-balck6 u-margin-left-10">{{wek}}</text>
					<text class="u-font-26 u-color-balck6 u-margin-left-10">{{time}}</text>
				</view>
			</view>
		</view>
		
		<view style="height: 110rpx;"></view>
		<view class="bootbtn u-flex u-row-between" v-if="list.money>0">
			<view class="u-flex-col u-margin-left-20">
				<view class="u-flex">
					<text class="u-font-28 u-font-bold u-color-balck6">立即付款：</text>
					<text class="u-font-28 u-font-bold" style="color: #FF8244;">￥{{list.money?list.money:'0'}}</text>
				</view>
				<text class="u-font-24 u-color-balck6 u-margin-top-10">订单总价：￥{{list.money?list.money:'0'}}</text>
			</view>
			<view @click="add_order()" class="ribtn zhuti_bg u-flex u-row-center u-col-center u-margin-right-20">
				<text class="u-font-28 u-color-white">去支付</text>
			</view>
		</view>
		<view class="bootbtn u-flex u-row-center u-col-center" v-if="list.money==0">
			<view @click="add_order()" class="ribtn zhuti_bg u-flex u-row-center u-col-center" style="width: 600rpx;">
				<text class="u-font-28 u-color-white">立即预约</text>
			</view>
		</view>
		
	</view>
</template>

<script>
	export default {
		data() {
			return {
				form:[],
				list:[],
				token:'',
				riqi:'',
				time:'',
				wek:'',
				timeid:0,
				indexx:0
			}
		},
		onLoad(e) {
			if(e.id){
				this.id=e.id;
			}
			if(e.riqi){
				this.riqi=e.riqi;
			}
			if(e.time){
				this.time=e.time;
			}
			if(e.wek){
				this.wek=e.wek;
			}
			if(e.timeid){
				this.timeid=e.timeid;
			}
			this.GetToken();
		},
		onPullDownRefresh() {
			this.getData();
			setTimeout(function () {
			    uni.stopPullDownRefresh();
			}, 1000);
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
			danxuanval(val,index1,index2){
				this.form[index1].value=val;
			},
			bindPickerChange: function(e) {
				var ind=e.target.id;
				var va=this.form[ind].arr[e.target.value];
				this.form[ind].value=va;
				this.indexx = e.target.value;
			},
			checkboxChange: function (e) {
				var val=e.target.value;
				var ind=0;
				var str='';
				for (var i = 0, lenI = val.length; i < lenI; ++i) {
					var list=val[i];
					var strs=list.split("|");
					ind=strs[0];
					str+=strs[1];
					if(i<lenI-1){
						str+=',';
					}
				}
				this.form[ind].value=str;
			},
			uploadimg(index){
				var th = this;
				if(th.form[index].imageList.length>=5){
					uni.showToast({
						title: '最多只能上传5张图片',
						icon:"none",
						duration: 2000
					});
					return false;
				}
				uni.chooseImage({
					count:1,
					success: (chooseImageRes) => {
						const tempFilePaths = chooseImageRes.tempFilePaths;
						uni.uploadFile({
							url: th.$puburl+'resource/formupload', //仅为示例，非真实的接口地址
							filePath: tempFilePaths[0],
							name: 'file',
							fileType:"image",
							formData: {
								token: th.token
							}, 
							success: (res) => {
								var res = JSON.parse(res.data);
								if(res.code==200){
									var va=th.form[index].value;
									if(va){
										va=va+','+res.img;
										
									}else{
										va=res.img;
									}
									th.form[index].value=va;
									th.form[index].imageList = th.form[index].imageList.concat(res.savename);
									uni.showToast({
									    title: res.message,
										icon:'none',
									    duration: 2000
									});
								}else{
									uni.showToast({
									    title: res.message,
										icon:'none',
									    duration: 2000
									});
								}
							}
						});
					}
				});
			},
			getData(){
				var _this=this;
				this.$api.getyuyue({id:this.id,token:this.token}).then((res)=>{
					_this.list=res.data.data.arr;
					_this.form=res.data.data.form;
				})
			},
			add_order(){
				var _this=this;
				this.$api.add_yuyue({
					id:this.id,
					riqi:this.riqi,
					time:this.time,
					timeid:this.timeid,
					list:encodeURI(JSON.stringify(this.form)),
					token:this.token
				}).then((res)=>{
					if(res.data.code==200){
						var id=res.data.data.id;
						_this.getpay(id);
					}else if(res.data.code==300){
						var id=res.data.data.id;
						var money=res.data.data.money;
						uni.showToast({
							title: res.data.message,
							icon:"none",
							duration: 2000
						});
						setTimeout(() => {
							// 此为uView的方法，详见路由相关文档
							uni.navigateTo({
								url: '../pay/payok?id='+id+'&money='+money
							});
						},1500)
					}else{
						uni.showToast({
							title: res.data.message,
							icon:"none",
							duration: 2000
						});
					}
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
				this.onBridgeReady(id);
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
			GetToken(){
				var _this=this;
				uni.getStorage({
				    key: 'userData',
					success:function (res) {
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
	background-color: #EEEEEF;
}
.bootbtn{
	width: 750rpx;
	height: 100rpx;
	background-color: #FFFFFF;
	position: fixed;
	bottom: 0rpx;
	left: 0rpx;
}
.border-bootom{
	border-bottom: 1rpx solid #F4F5F7;
}
.ribtn{
	width: 300rpx;
	height: 80rpx;
	border-radius: 40rpx;
}
</style>
