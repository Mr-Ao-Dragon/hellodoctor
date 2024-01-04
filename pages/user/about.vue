<template>
	<view>
		
		<div class="u-flex u-row-center u-col-center u-flex-col" style="height: 300px;">
			<img v-if="img" :src="img" style="width: 80px;height: 80px;">
			<span class="u-margin-top-20 u-font-28 u-color-balck3" v-if="name">{{name}}</span>
			<span class="u-margin-top-20 u-font-28 u-color-balck3" v-if="version">{{version}}</span>
			<span class="u-margin-top-20 u-font-24 u-color-balck6" v-if="company">{{company}}版权所有</span>
		</div>
		
		<div class="userli u-flex u-row-between u-bg-white"  @click="tiao('../texter/text_xieyi?id=1')">
			<div class="user_cent">
				<span class="u-font-28 u-color-balck3">用户协议</span>
			</div>
			<i class="icon u-color-balck9">&#xe605;</i>
		</div>
		<div class="userli u-flex u-row-between u-bg-white"  @click="tiao('../texter/text_xieyi?id=2')">
			<div class="user_cent">
				<span class="u-font-28 u-color-balck3">隐私政策</span>
			</div>
			<i class="icon u-color-balck9">&#xe605;</i>
		</div>
		<div class="userli u-flex u-row-between u-bg-white"  @click="tiao('../texter/text_xieyi?id=3')">
			<div class="user_cent">
				<span class="u-font-28 u-color-balck3">联系我们</span>
			</div>
			<i class="icon u-color-balck9">&#xe605;</i>
		</div>
		<div class="userli u-flex u-row-between u-bg-white"  @click="tiao('more')">
			<div class="user_cent">
				<span class="u-font-28 u-color-balck3">更多资料</span>
			</div>
			<i class="icon u-color-balck9">&#xe605;</i>
		</div>
		
	</view>
</template>

<script>
	export default {
		data() {
			return {
				name:'',
				version:'',
				company:'',
				img:''
			}
		},
		onLoad() {
			this.getData();
		},
		methods: {
			getData() {
				var _this=this;
				this.$api.getsite().then(res => {
					_this.name=res.data.data.name;
					_this.company=res.data.data.company;
					_this.version=res.data.data.version;
					_this.img=res.data.data.img;
				})
			},
			GetToken(){
				var _this=this;
				uni.getStorage({
				    key: 'userData',
				    success: function (res) {
				        _this.id=res.data.vuex_user.id;
						_this.name=res.data.vuex_user.name;
						_this.headimg=res.data.vuex_user.headimg;
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
	.userli{
		height: 110rpx;
		border-bottom: 1px solid #F6F6F6;
		padding: 0 20rpx;
	}
	.user_cent{
		width: 645rpx;
	}
</style>
