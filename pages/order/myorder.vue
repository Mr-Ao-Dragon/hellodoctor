<template>
	<view>
		
		<div class="topvie u-flex u-bg-white">
			<div class="topli u-flex u-flex-1 u-row-center u-col-center u-font-28" :class="type==0?'zhuti_border_bottom-1 zhuti_color':'u-color-balck3'" @click="setnav(0)">全部</div>
			<div class="topli u-flex u-flex-1 u-row-center u-col-center u-font-28" :class="type==1?'zhuti_border_bottom-1 zhuti_color':'u-color-balck3'" @click="setnav(1)">待付款</div>
			<div class="topli u-flex u-flex-1 u-row-center u-col-center u-font-28" :class="type==2?'zhuti_border_bottom-1 zhuti_color':'u-color-balck3'" @click="setnav(2)">待核销</div>
			<div class="topli u-flex u-flex-1 u-row-center u-col-center u-font-28" :class="type==3?'zhuti_border_bottom-1 zhuti_color':'u-color-balck3'" @click="setnav(3)">已完成</div>
			<div class="topli u-flex u-flex-1 u-row-center u-col-center u-font-28" :class="type==4?'zhuti_border_bottom-1 zhuti_color':'u-color-balck3'" @click="setnav(4)">已取消</div>
		</div>
		<div class="u-flex-col u-padding-20">
			<div class="u-flex-col u-padding-20  u-bg-white" style="border-bottom: 10rpx solid #F6F6F6;" v-for="(item,index) in list" :key="index" @click="tiao('myorder_info?id='+item.id)">
				<view class="u-flex u-row-between u-color-balck6">
					<span class="u-font-26">ID：{{item.id}}</span>
					<span class="u-font-26" v-if="item.status==1">待支付</span>
					<span class="u-font-26" v-if="item.status==2">待核销</span>
					<span class="u-font-26" v-if="item.status==3">已完成</span>
					<span class="u-font-26" v-if="item.status==4">已取消</span>
					<span class="u-font-26" v-if="item.status==5">已退款</span>
				</view>
				<div class="u-margin-top-20 u-flex" style="border-bottom: 1px solid #F6F6F6;padding-bottom: 10px;justify-content: space-between;">
					<div class="u-flex">
						<image :src="item.img" mode="" style="width: 80rpx;height: 80rpx;"></image>
					</div>
					<div class="u-flex-col" style="width: 570rpx;">
						<text class="u-font-26 u-color-balck6 u-font-dan-sheng" style="width: 570rpx;">{{item.title}}</text>
						<text class="u-font-26 zhuti_color_hong u-font-bold u-margin-top-10" v-if="item.money>0">{{item.money}}</text>
					</div>
				</div>
				<div class="u-flex u-row-between" style="margin-top: 10rpx;">
					<text class="u-font-24 u-color-balck6"  style="width:610rpx;">预约时间：{{item.y_data}} {{item.y_time}}</text>
					<div class="u-font-28 zhuti_color">查看</div>
				</div>
			</div>
		</div>
		
		<uni-load-more :status="status"></uni-load-more>
		
	</view>
</template>

<script>
	export default {
		data() {
			return {
				navst:0,
				list:[],
				title: 'Hello',
				token:'',
				status: 'loadmore',
				type:0,
				count: 10,
				res_count:0,
				page: 1
			}
		},
		onLoad(e){
			if(e.type){
				this.type=e.type;
			}
			this.GetToken();
		},
		onPullDownRefresh() {
			this.getData();
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
				this.getData();
			}, 2000)
		},
		onBackPress() {
			uni.switchTab({
				url:'../user/index'
			});
			return true;
		},
		methods: {
			setnav(st){
				this.list=[];
				this.status = 'loading';
				this.navst=st;
				this.type=st;
				this.page=0;
				this.getData();
			},
			getData(){
				var _this=this;
				this.$api.my_yuyue({'token':this.token,type:this.type,page:this.page}).then((res)=>{
					res.data.data.list.forEach(item=>{
						_this.list.push(item);
					});
					_this.res_count=res.data.data.count;
					if(_this.count >= _this.res_count) _this.status = 'nomore';
					else _this.status = 'loading';
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
	page{
		background-color: #F6F6F6;
	}
	.topvie{
		width: 750rpx;
		position: -webkit-sticky;
		position: sticky;
		top: var(--window-top);
		z-index: 99;
		height: 80rpx;
	}
	.topli{
		height: 80rpx;
	}
</style>
