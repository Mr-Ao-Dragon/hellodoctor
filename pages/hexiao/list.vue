<template>
	<view>
		
		<!--#ifdef MP-WEIXIN-->
		<view class="u-flex u-padding-20 u-row-center u-col-center" @click="saoma()">
			<text class="icon u-font-42 zhuti_color">&#xe62d;</text>
			<text class="u-font-28 u-margin-left-20 zhuti_color">扫码核销</text>
		</view>
		<view class="jianju-5"></view>
		<!--#endif-->
		
		<!--统计-->
		<div class="u-flex" style="border-bottom:1px solid #f8f8f8;">
			<div class="u-flex-col u-row-center u-col-center" style="flex:1;height:60px;">
				<span class="u-font-26 u-color-balck6">今日核销</span>
				<span class="u-font-32 u-color-balck6 u-font-bold">{{zno1}}次</span>
			</div>
			<div class="u-flex-col u-row-center u-col-center" style="flex:1;height:60px;">
				<span class="u-font-26 u-color-balck6">昨日核销</span>
				<span class="u-font-32 u-color-balck6 u-font-bold">{{zno2}}次</span>
			</div>
			<div class="u-flex-col u-row-center u-col-center" style="flex:1;height:60px;">
				<span class="u-font-26 u-color-balck6">总核销</span>
				<span class="u-font-32 u-color-balck6 u-font-bold">{{res_count}}次</span>
			</div>
		</div>
		<view class="jianju-5"></view>
		
		<view class="u-flex-col"  v-for="(item,index) in list" :key="index" @click="tiao('../user/hexiao?id='+item.order_id)">
			<div class="listli u-flex-col u-padding-20" style="border-bottom:1rpx solid #F6F6F6;line-height: 40rpx;">
				<div class="u-flex u-row-between">
					<text class="u-font-28 zhuti_color_lan u-font-bold">订单ID：{{item.order_id}}</text>
					<view class="u-flex" style="font-size:13px;color:#666666;">
						<image :src="item.headimg" style="width:40rpx;border-radius:10px;margin-right:4px;height: 40rpx;"/>
						<text class="u-font-26 u-color-balck6">{{item.nick}}</text>
					</view>
				</div>
				<div class="u-flex u-row-between u-margin-top-10">
					<span class="u-font-14 u-color-black6" style="font-size:13px;color:#666666;">姓名：{{item.xmname}}</span>
					<span class="u-font-14 u-color-black6" style="font-size:12px;color:#666666;">联系电话：{{item.xmmobile}}</span>
				</div>
				<div class="u-flex u-row-between u-margin-top-10">
					<span class="u-font-14 u-color-black6" style="font-size:13px;color:#666666;">预约日期：{{item.riqi}}</span>
					<span class="u-font-14 u-color-black6" style="font-size:12px;color:#666666;">预约时间：{{item.time}}</span>
				</div>
				<div class="u-flex u-row-between u-margin-top-10">
					<span class="u-font-14 u-color-black6" style="font-size:13px;color:#666666;">核销时间：{{item.addtime}}</span>
					<text class="u-font-14 zhuti_color_huang u-font-bold" style="font-size:12px;">核销项目：{{item.name}}</text>
				</div>
				<view class="u-flex u-row-between u-margin-top-10">
					<text class="u-font-14 u-color-black6" style="font-size:13px;color:#666666;">总核销：{{item.heno}}次</text>
					<text class="u-font-14 u-color-black6" style="font-size:12px;color:#666666;">已核销：{{item.usehe}}次</text>
				</view>
			</div>
		</view>
		
		<uni-load-more :status="status"></uni-load-more>
		
	</view>
</template>

<script>
export default {
	data() {
		return {
			token:'',
			tanst:1,
			showtit:'',
			bgcolor:'',
			address:'',
			titcolor:'#FFFFFF',
			id:0,
			navst:0,
			list:[],
			status: 'more',
			type:0,
			count: 10,
			res_count:0,
			page: 1,
			orderid:'',
			zno1:0,
			zno2:0
		}
	},
	onShow(){
		this.list=[];
		this.page=1;
		this.status='more';
		this.GetToken();
	},
	onPullDownRefresh() {
		this.list=[];
		this.page=1;
		this.status='more';
		this.getData2();
		setTimeout(function () {
			uni.stopPullDownRefresh();
		}, 1000);
	},
	onReachBottom() {
		if(this.count >= this.res_count) return ;
		this.status = 'loading';
		this.page = ++ this.page;
		setTimeout(() => {
			this.count += 10;
			this.getData2();
		}, 2000)
	},
	onBackPress() {
		uni.redirectTo({
			url:'../user/index'
		});
		return true;
	},
	methods: {
		saoma(){
			if(!this.token){
				uni.navigateTo({
					url:'/pages/login/login'
				})
				return false;
			}
			 var str='281&1';
			// //拆分id和type出来 跳转
			let arr = str.split('&');
			var id=arr[0];
			uni.navigateTo({
				url:'../order/myorder_info?id='+id
			})
			return false;
			// if(arr[1]==1){
			// 	uni.navigateTo({
			// 		url:'../user/hexiao?id='+id
			// 	})
			// }
			// return false;
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
							url:'../user/hexiao?id='+id
						})
					}
			    }
			});
		},
		GetToken(){
			var _this=this;
			uni.getStorage({
			    key: 'userData',
			    success: function (res) {
					_this.token=res.data.vuex_token;
					_this.getData2();
			    }
			});
		},
		getData2(){
			var _this=this;
			this.$api.hexiaolist({'token':this.token,type:this.navst,page:this.page}).then((res)=>{
				_this.list=res.data.data.list;
				_this.res_count=res.data.data.count;
				_this.zno1=res.data.data.zno1;
				_this.zno2=res.data.data.zno2;
				if(_this.count >= _this.res_count) _this.status = 'nomore';
				else _this.status = 'loading';
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
</style>
