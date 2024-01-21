<template>
  <view class="content">
    <view class="u-flex u-col-center u-row-between u-padding-20">
      <view class="imglist">
        <swiper class="swiper" :indicator-dots="true" :autoplay="true">
          <swiper-item v-for="(item,index) in lunbo" :key="index">
            <image :src="item.img" mode="" class="swiimg" style="border-radius: 20rpx;"></image>
          </swiper-item>
        </swiper>
      </view>
    </view>

    <!--公告-->
    <view class="notilist u-border-ra20" v-if="noti && noti.length>0">
      <view class="u-flex u-row-between u-row-center u-padding-20 u-col-center">
        <!-- 图标 -->
        <text class="icon zhuti_color u-font-28">&#xe609;</text>
        <view class="noticent">
          <swiper :vertical="true" :autoplay="true" style="height: 40rpx;line-height: 40rpx;">
            <swiper-item class="u-font-dan-sheng" @click="tiao('../texter/noti_info?id='+item.id)" v-for="(item,index) in noti" :key="index">
              <text class="u-font-28 u-color-balck3 u-font-dan-sheng" style="width: 620rpx;">{{item.title}}</text>
            </swiper-item>
          </swiper>
        </view>
        <text class="icon u-color-balck9 u-font-28">&#xe605;</text>
      </view>
    </view>

    <!--列表-->
    <view class="u-flex-col gxcent u-margin-bottom-20 u-border-ra20 u-bg-white u-margin-top-20" v-if="list && list.length>0">
      <view class="u-padding-20 u-flex border-bootom-1">
        <text class="icon zhuti_color u-font-28">&#xe6be;</text>
        <span class="u-font-28 u-color-balck3 u-margin-left-10">热门服务</span>
      </view>
      <view v-for="(item,index) in list" :key="index" class="u-padding-20 u-flex border-bootom-1" @click="gotoReserve(item.id)">
        <image :src="item.img" style="width: 200rpx;height: 140rpx;border-radius: 20rpx;" mode=""></image>
        <view class="u-flex-col u-margin-left-20 clix u-row-between">
          <text class="u-font-28 u-color-balck3">{{item.title}}</text>
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
//#ifdef H5
var jweixin = require('@/components/jweixin-module/index.js');
//#endif
export default {
  data() {
    return {
      name: '',
      mobile: '',
      title: 'Hello',
      lunbo: [],
      kuai: [],
      token: '',
      list: [],
      noti: [],
      guang: [],
      classr: [],
      status: 'loadmore',
      page: 1,
      count: 10,
      res_count: 0,
      wxcode: '',
      appid: '',
      kfurl: '',
      iswx: 1,
      uuid: '',
      index: 0,
      sharedata: [],
      stores: [],
      identity: [],
      fenxiang: {
        path: '/pages/index/index'
      }
    }
  },
  onLoad(option) {

    // //微信公众号登录
    // let href = window.location.href;
    // if (href.includes("cn/?code")) {  //url包括 com/?code 证明为从微信跳转回来的
    //   var url = href.substring(0, href.length - 2); //vue自动在末尾加了 #/ 符号，截取去掉
    //   var jingPosit = url.indexOf("cn/") + 3; //获取域名结束的位置
    //   var urlLeft = url.substring(0, jingPosit);//url左侧部分
    //   var urlRight = url.substring(jingPosit, url.length); //url右侧部分
    //   var url = urlLeft + "#/" + urlRight;
    //   window.location = url;//拼接跳转
    //   console.log(url); return false;
    // }
    // if (href.includes("com/?code")) {  //url包括 com/?code 证明为从微信跳转回来的
    //   var url = href.substring(0, href.length - 2); //vue自动在末尾加了 #/ 符号，截取去掉
    //   var jingPosit = url.indexOf("com/") + 4; //获取域名结束的位置
    //   var urlLeft = url.substring(0, jingPosit);//url左侧部分
    //   var urlRight = url.substring(jingPosit, url.length); //url右侧部分
    //   var url = urlLeft + "#/" + urlRight;
    //   window.location = url;//拼接跳转
    //   console.log(url); return false;
    // }
    // if (href.includes("net/?code")) {  //url包括 net/?code 证明为从微信跳转回来的
    //   var url = href.substring(0, href.length - 2); //vue自动在末尾加了 #/ 符号，截取去掉
    //   var jingPosit = url.indexOf("net/") + 4; //获取域名结束的位置
    //   var urlLeft = url.substring(0, jingPosit);//url左侧部分
    //   var urlRight = url.substring(jingPosit, url.length); //url右侧部分
    //   var url = urlLeft + "#/" + urlRight;
    //   window.location = url;//拼接跳转
    //   console.log(url); return false;
    // }
    // //todo 疑似是邀请码 可以去掉
    // if (option.code) {
    //   this.wxcode = option.code;
    //   this.getopenid();
    //   return false;
    // }

    // //todo 疑似登录相关
    // this.iswx = this.$iswx;
    // var that = this;
    // if (this.iswx == 2) {
    //   var apiUrl = location.href.split("#")[0];
    //   uni.request({
    //     url: that.$puburl + 'resource/Wxlogin/getSignPackage',
    //     data: {
    //       url: encodeURIComponent(apiUrl),//当前页面的域名
    //       api: ['scanQRCode', 'checkJsApi'],//调用的方法去接口列表里找
    //     },
    //     success: function (res) {
    //       var wxData = res.data.data.data;
    //       that.wx_sanCode(wxData)
    //     }
    //   })
    // }
    //获取初始资料
    this.getData();
    //是否为微信，可以先留着
    // this.iswx = this.$iswx;//1网页 2微信
    //获取token
    // this.GetToken();

  },
  onPullDownRefresh() {
    //下拉刷新
    this.list = [];
    this.page = 1;
    this.status = 'loadmore';
    this.getData();
    setTimeout(function () {
      uni.stopPullDownRefresh();
    }, 1000);
  },
  methods: {
    gotoReserve:function(id){
      uni.navigateTo({
        url:"/newpages/reserve?id="+id
      })
    },
    //扫码
    //todo 扫码登录 不知何用，后续可删除
    wx_sanCode: function (wxData) {
      var _this = this;
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
      jweixin.ready(() => {
        jweixin.updateAppMessageShareData({
          title: _this.sharedata.title, // 分享标题
          desc: _this.sharedata.desc, // 分享描述
          link: apiUrl, // 分享链接，该链接域名或路径必须与当前页面对应的公众号JS安全域名一致
          imgUrl: _this.sharedata.imgUrl, // 分享图标									
          success: function (data) {
            // 设置成功
            console.log('updateAppMessageShareData success:', data);
          },
          fail: function (error) {
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
      jweixin.error(function (res) {
        // config信息验证失败会执行error函数，如签名过期导致验证失败，具体错误信息可以打开config的debug模式查看，也可以在返回的res参数中查看，对于SPA可以在这里更新签名。
        console.log(res, '接口验证失败')
      });
    },
    GetToken() {
      var _this = this;
      uni.getStorage({
        key: 'userData',
        success: function (res) {
          _this.token = res.data.vuex_token;
          _this.uid = value.vuex_user.id;
          _this.fenxiang.path = '/pages/index/index?uid=' + _this.uid;
        },
        fail() {
          if (_this.iswx == 2) {
            _this.getUserInfo();
          }
        }
      });
    },
    // 三个用于获取数据的接口
    getUserInfo() {
      if (this.$wxurl) {
        var xurl = this.$wxurl + '?type=2';
      } else {
        var xurl = this.$puburl;
      }
      var url = encodeURI('https://open.weixin.qq.com/connect/oauth2/authorize?appid=' + this.$wxappid + '&redirect_uri=' + xurl + '&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect');
      window.location.href = url;
      if (this.code) {
        //将onLoad中的获取token方法放进来即可
      } else {
        //没有code就去获取code
      }
    },
    getopenid() {
      var _this = this;
      if (this.wxcode) {
        uni.showLoading({
          title: '登录中...'
        });
        var uid = 0;
        //获取师傅id
        const value = uni.getStorageSync('invitation');
        if (value) {
          if (value.uid) {
            uid = value.uid;
          }
        }
        uni.request({
          url: _this.$puburl + 'resource/wxh5register',
          data: {
            code: this.wxcode,
            uid: uid
          },
          method: 'GET',
          header: {
            'content-type': 'application/json'
          },
          success: (res) => {
            console.log(JSON.stringify(res));
            if (res.data.code == 1) {
              _this.token = res.data.data.token;
              var token = res.data.data.token;
              var userlist = res.data.data.userlist;
              //console.log('uu==='+userlist);
              //登录成功  跳转到首页
              uni.setStorage({
                key: 'userData',
                data: {
                  'vuex_token': token,
                  'vuex_user': userlist
                },
                success: function () {
                  console.log('success');
                }
              });
              window.location = _this.$puburl;//跳转
              _this.getData();
              //_this.wxGetUserInfo(); //需要手动确认
            } else {
              uni.showToast({
                title: res.data.message,
                icon: "none",
                duration: 2000
              });
            }
            //openId、或SessionKdy存储//隐藏loading
            uni.hideLoading();
          }, fail(e) {
            console.log(e);
            uni.hideLoading();
          }
        });
      } else {
        uni.showToast({
          icon: 'none',
          title: '请先登录'
        })
        //没有code在走一次流程去获取code
      }
    },

    getData() {
      //获取首页数据
      // this.$api.indeximg().then((res) => {
      //   //todo:轮播图和list的获取

      // })
      this.lunbo = [
        {
          "id": 2,
          "title": "轮播图1",
          "img": "../static/lunbo1.jpg",
          "type": 1,
          "url": "#",
          "style": 1
        }
      ]
      this.list = [
        {
          "id": 16,
          "title": "王医生",
          "mobile": "13688885555",
          "img": "../static/avatar.jpg",
          "address": "15分钟一个患者",
          "lat": "1",
          "lng": "1",
          "is_info": 1,
          "recommended": "",
          "val": 0,
          "opst": 1
        }
      ]
    },
    tiao(url) {
      uni.navigateTo({
        url: url
      })
    }
  }
}
</script>

<style>
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
