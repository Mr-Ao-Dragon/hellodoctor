<template>
	<view>
		
		<view class="u-flex class_list">
			<view class="class_left">
				<scroll-view scroll-y="true" class="scroll-Y">
				<view v-for="(item,index) in classx" :key="index" :class="navtop==index?'activebg zhuti_color':'u-color-balck6'" class="left_li u-flex-col u-row-center u-col-center" @click="setnavtop(index)">
					<text class="u-font-24 u-font-dan-sheng" style="width: 200rpx;text-align: center;">{{item.day}}</text>
					<text class="u-font-24 u-margin-top-5">{{item.week}}</text>
				</view>
				<!--占位-->
				<view style="height: 110rpx;"></view>
				</scroll-view>
			</view>
			<view style="width: 200rpx;"></view>
			<view class="class_right">
				
				<view class="class_rightli u-flex-col">
					<view class="u-padding-20 u-flex u-row-between zhuti_bg" style="width: 510rpx;border-top-left-radius: 20rpx;border-top-right-radius: 20rpx;">
						<text class="u-font-30 u-font-bold u-color-white">{{riqi}}</text>
						<text class="u-font-30 u-font-bold u-color-white">{{wek}}</text>
					</view>
					<view class="u-bg-white" style="width: 550rpx;border-bottom-left-radius: 20rpx;border-bottom-right-radius: 20rpx;padding: 20rpx 0;">
						<text class="u-font-26 u-color-balck9 u-margin-left-20">请选择预约时间</text>
						<view class="u-flex u-margin-top-20 u-flex-wrap">
							<view v-for="(item,index) in list" :key="index" class="u-flex u-flex-wrap">
								<view class="timeli" v-for="(itemx,index2) in item.list" :key="index2" @click="tiaourl(itemx.status,itemx.time,item.id)">
									<view :class="itemx.status==1 && time!=itemx.time?'u-color-balck3':itemx.status==1 && time==itemx.time?'u-color-white xuan zhuti_bg':'u-color-balck9'">
										<text class="u-font-26">{{itemx.title}}</text>
									</view>
								</view>
							</view>
							
						</view>
					</view>
				</view>
				<!--占位-->
				<view style="height: 110rpx;"></view>
			</view>
		</view>
		
		<view class="bootbtn u-flex u-row-center u-col-center">
			<view class="ribtn zhuti_bg u-flex u-row-center u-col-center" @click="tiaox()">
				<text class="u-font-28 u-color-white">确定</text>
			</view>
		</view>
		
	</view>
</template>

<script>
	export default {
		data() {
			return {
				title: 'Hello',
				list:[],
				id:0,
				navtop:0,
				classx:[],
				time:'',
				riqi:'',
				wek:'',
				timeid:0
			}
		},
		onLoad(e) {
			if(e.id){
				this.id=e.id;
			}
			this.GetToken();
		},
		methods: {
			getData(){
				var _this=this;
				this.$api.getydate({id:this.id}).then((res)=>{
					_this.classx=res.data.data.date;
					_this.setnavtop(0);
					_this.getytime();
					console.log(res);
				})
			},
			getytime(){
				var _this=this;
				this.$api.getytime({id:this.id,riqi:this.riqi}).then((res)=>{
					_this.list=res.data.data;
					console.log(res);
				})
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
			tiaourl(st,time,timeid){
				if(st==1){
					this.time=time;
					this.timeid=timeid;
				}else if(st==2){
					uni.showToast({
						title: '预约已满，请选择其他时间',
						icon:"none",
						duration: 2000
					});
				}else if(st==3){
					uni.showToast({
						title: '该时间已过期，请重新选择',
						icon:"none",
						duration: 2000
					});
				}
				return false;
			},
			tiaox(){
				if(this.time){
					var url="order?id="+this.id+'&riqi='+this.riqi+"&time="+this.time+'&wek='+this.wek+'&timeid='+this.timeid;
					this.tiao(url);
				}else{
					uni.showToast({
						title: '请选择预约时间',
						icon:"none",
						duration: 2000
					});
				}
			},
			setnavtop(st){
				this.navtop=st;
				this.riqi=this.classx[st].riqi;
				this.wek=this.classx[st].week;
				this.time='';
				this.getytime();
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
	.activebg{
		background-color: #EEEEEF!important;
	}
	.xuan{
		padding: 5rpx 10rpx;
		border-radius: 16rpx;
	}
	.ribtn{
		width: 690rpx;
		height: 84rpx;
		border-radius: 40rpx;
	}
	.bootbtn{
		width: 750rpx;
		height: 110rpx;
		background-color: #FFFFFF;
		position: fixed;
		bottom: 0rpx;
		left: 0rpx;
		z-index: 8;
		border-top: 1rpx solid #F6F6F6;
	}
	.class_list{
		
	}
	.timeli{
		width: 136rpx;
		height: 70rpx;
		display: flex;
		justify-content: center;
		align-items: center;
	}
	.scroll-Y{
		width: 160rpx;
		white-space: nowrap;
		height: 100%;
	}
	.class_left{
		width: 160rpx;
		position: fixed;
		top: var(--window-top);
		height: 100%;
		background-color: #FFFFFF;
		z-index: 3;
	}
	.left_li{
		height: 90rpx;
		display: flex;
		align-items: center;
		border-bottom: 1px solid #F6F6F6;
	}
	.xiaoli{
		width: 6rpx;
		height: 60rpx;
	}
	.class_right{
		width: 570rpx;
		height: 100%;
	}
	.class_rightli{
		width: 590rpx;
		margin-top: 20rpx;
	}
	.yue_btn{
		width: 100px;
		height: 50px;
		display: flex;
		justify-content: flex-end;
		align-items: center;
	}
	.yuyue_btn{
		width: 70px;
		height: 30px;
		display: flex;
		justify-content: center;
		align-items: center;
		border-radius: 5px;
	}
</style>
