<template>
	<view>
		
		<view class="u-flex">
			<view class="u-flex u-flex-1 u-row-center u-col-center u-padding-20" :class="type==1?'zhuti_color zhuti-border-bootom-1':'u-color-balck3'" @tap="setnav(1)">
				<text class="u-font-28">我邀请</text>
			</view>
			<view class="u-flex u-flex-1 u-row-center u-col-center u-padding-20" :class="type==2?'zhuti_color zhuti-border-bootom-1':'u-color-balck3'" @tap="setnav(2)">
				<text class="u-font-28">下级邀请</text>
			</view>
		</view>
		
		<view class="u-flex-col">
			<view class="u-flex u-padding-20 u-row-between u-border-bottom-1" v-for="(item, index) in list" :key="index">
				<view class="u-flex-col">
					<text class="u-font-28 u-color-balck6">{{item.nick}}</text>
					<text class="u-font-28 u-color-balck6 u-margin-top-10">{{item.reg_time}}</text>
				</view>
				<view class="u-flex-col u-row-right u-col-bottom">
					<text class="u-font-28 zhuti_color_hong" v-if="type==1">{{item.rel1_money}}元</text>
					<text class="u-font-28 zhuti_color_hong" v-if="type==2">{{item.rel2_money}}元</text>
					<text class="u-font-28 zhuti_color_lv u-margin-top-10" v-if="item.status==1">正常</text>
					<text class="u-font-28 u-color-balck9 u-margin-top-10" v-if="item.status!=1">无效</text>
				</view>
			</view>
		</view>
		
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
				type:1,
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
				this.$api.invitalist({'token':this.token,type:this.type,page:this.page}).then((res)=>{
					_this.list=res.data.data.list;
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
		background: #FFFFFF;
	}
</style>
