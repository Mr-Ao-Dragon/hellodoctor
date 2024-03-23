<template>
  <view class="content">
    <view class="u-flex u-col-center u-row-between u-padding-20">
      <view class="imglist">
        <swiper class="swiper" :indicator-dots="true" :autoplay="true">
          <swiper-item v-for="(item,index) in lunbo" :key="index">
            <image :src="item" mode="" class="swiimg" style="border-radius: 20rpx;"></image>
          </swiper-item>
        </swiper>
      </view>
    </view>

    <!--公告-->
    <view class="notilist u-border-ra20" v-if="noti && noti.length>0">
      <view class="u-flex u-row-between u-row-center u-padding-20 u-col-center">
        <text class="icon zhuti_color u-font-28">&#xe852;</text>
        <view class="noticent">
          <swiper :vertical="true" :autoplay="true" style="height: 40rpx;line-height: 40rpx;">
            <swiper-item class="u-font-dan-sheng" v-for="(item,index) in noti" :key="index">
              <text class="u-font-28 u-color-balck3 u-font-dan-sheng" style="width: 620rpx;">{{item}}</text>
            </swiper-item>
          </swiper>
        </view>

      </view>
    </view>

    <!--列表-->
    <view class="u-flex-col gxcent u-margin-bottom-20 u-border-ra20 u-bg-white u-margin-top-20" v-if="list && list.length>0">
      <view class="u-padding-20 u-flex border-bootom-1">
        <text class="icon zhuti_color u-font-28">
          &#xe831;</text>
        <span class="u-font-28 u-color-balck3 u-margin-left-10">门店医生</span>
      </view>
      <view v-for="(item,index) in list" :key="index" class="u-padding-20 u-flex border-bootom-1" @click="gotoReserve(item)">
        <image :src="item.avatar" style="width: 200rpx;height: 140rpx;border-radius: 20rpx;" mode=""></image>
        <view class="u-flex-col u-margin-left-20 clix u-row-between">
          <text class="u-font-28 u-color-balck3">{{item.name}}</text>
          <view class="u-flex u-row-between">
            <!-- <text class="u-font-24 zhuti_color" style="color: #19be6b;">电话:{{item.mobile}}</text> -->
            <view class="lxbtn">
              <text class="u-font-24 zhuti_color" style="color: #19be6b;">立即预约</text>
            </view>
          </view>
        </view>
      </view>
    </view>
  </view>
</template>

<script>
import request from '../api/request';

export default {
  data() {
    return {
      lunbo: [],
      noti: [],
      list: [],
      ready: false,
      wxLoginData: {
        appid: 'wx277005e156d46f0a',//填写appid
        redirect_uri: '',//填写这个网页的地址
        response_type: 'code',
        scope: 'snsapi_userinfo',
      }

    }
  },
  async onLoad(options) {
    await this.getData();
    this.ready = true;
    if (options && options.code) {
      const res = await request.postLogin(code)
      uni.setStorageSync('token', res.access_token)
      uni.setStorageSync('expries_in', res.expries_in)
    }
    else {
      const loginLink = `https://open.weixin.qq.com/connect/oauth2/authorize?appid=${wxLoginData.appid}&redirect_uri=${wxLoginData.redirect_uri}&response_type=${wxLoginData.response_type}&scope=${wxLoginData.scope}#wechat_redirect`
        window.location.href = loginLink
    }
  },
  methods: {
    gotoReserve: function (item) {
      uni.navigateTo({
        url: `/pages/info_doc?id=${item.id}&name=${item.name}`
      })
    },
    async getData() {
      const doclist = await request.getDocList()
      const notify = request.getNotice()
      this.lunbo = notice.lunbotu
      this.noti = notice.notice
      this.list = doclist
    },
  }
}
</script>

<style scoped>
.icon {
  margin-right: 8px;
}
page {
  background-color: #f6f6f6;
}
.kefuvie {
  position: fixed;
  right: 40rpx;
  bottom: 200rpx;
  width: 78rpx;
  height: 78rpx;
  background: #ffffff;
  border-radius: 40rpx;
  z-index: 9;
  display: flex;
  justify-content: center;
  align-items: center;
}
.faguang {
  box-shadow: 0px 0px 16rpx #cccccc;
  border-radius: 40rpx;
}
.putx {
  width: 500rpx;
  height: 70rpx;
  background-color: #f0f0f0;
  margin-bottom: 15rpx;
  margin-top: 8rpx;
  text-indent: 20rpx;
  font-size: 28rpx;
}
.popbox {
  padding: 20rpx;
  background-color: #ffffff;
  border-radius: 10rpx;
}
.clix {
  width: 490rpx;
  height: 130rpx;
}
.lxbtn {
  width: 140rpx;
  height: 48rpx;
  border: 1rpx solid #19be6b;
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 10rpx;
}
.gxcent {
  width: 710rpx;
}
.xiaoli {
  width: 6rpx;
  height: 36rpx;
}
.noticent {
  width: 620rpx;
  height: 40rpx;
}
.notilist {
  width: 710rpx;
  background-color: #ffffff;
}
.content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
.an_li {
  width: 177.5rpx;
}
.imglist {
  width: 710rpx;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.swiimg {
  width: 710rpx;
  height: 300rpx;
}
</style>
