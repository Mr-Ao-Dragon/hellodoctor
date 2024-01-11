<template>
  <view>

    <view class="usertop zhuti_bg"  style="overflow: hidden;border-bottom-left-radius:20px;border-bottom-right-radius:20px;">
      <view class="u-flex u-col-center u-row-center" style="width:100%;margin-top:25px;">
        <image src="../static/img/user/mo.png" style="width: 50px;height:50px;border-radius:50%;border:3px solid #FFFFFF;" />
          <view class="u-flex-col u-margin-left-10" style="width:65%;">
            <text class="u-font-28 u-font-bold u-color-white">{{name?name:'请登录'}}</text>
          </view>
      </view>
    </view>

    <view class="userlist u-flex-col u-margin-top-20" >
      <view class="userli u-flex u-row-between u-bg-white" @click="goto(0)">
        <view class="iconlist u-flex u-row-center u-col-center">
          <i class="icon u-font-44 zhuti_color">&#xe63c;</i>
        </view>
        <view class="user_cent" >
          <span class="u-font-28 u-color-balck3">我的预约</span>
        </view>
        <text class="icon u-color-balck9">&#xe623;</text>
      </view>
      <view class="userli u-flex u-row-between" @click="goto(1)">
        <view class="iconlist u-flex u-row-center u-col-center">
          <text class="icon u-font-38 zhuti_color">&#xe60b;</text>
        </view>
        <view class="user_cent">
          <text class="u-font-28 u-color-balck3">预约管理(医生端)</text>
        </view>
        <text class="icon u-color-balck9">&#xe623;</text>
      </view>
      <view class="userli u-flex u-row-between" @click="goto(2)">
        <view class="iconlist u-flex u-row-center u-col-center">
          <text class="icon u-font-38 zhuti_color">&#xe6fd;</text>
        </view>
        <view class="user_cent">
          <text class="u-font-28 u-color-balck3">用户管理（管理端）</text>
        </view>
        <text class="icon u-color-balck9">&#xe623;</text>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      title: 'Hello',
      headimg: '',
      name: '用户',
      token: '',
      role: ""//角色，1医生2患者3管理员
    }
  },
  onShow() {
    // this.GetToken();
  },
  methods: {
    goto(i){
      const url = ['list','list_doc','list_admin']
      uni.navigateTo({
        url:`/newpages/${url[i]}`
      })
    },
    GetToken() {
      var _this = this;
      uni.getStorage({
        key: 'userData',
        success: (res) => {
          _this.name = res.data.vuex_user.name;
          _this.token = res.data.vuex_token;
          _this.getData();
        }
      });
    },
    getData() {
      var _this = this;
      this.$api.user_info({ token: this.token }).then(res => {
        //获取个人信息
      })
    },
    tiao(url) {
      if (!this.token) {
        //清理本地
        uni.removeStorage({
          key: 'userData'
        });
        uni.navigateTo({
          url: '/pages/login/login'
        })
        return false;
      }
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
.quan {
  width: 36rpx;
  height: 36rpx;
  border-radius: 20rpx;
  top: -16rpx;
  right: 40rpx;
  background-color: #ff5722;
}
.xfvie {
  width: 720rpx;
  height: 160rpx;
  background: #ffffff;
  position: absolute;
  left: 15rpx;
  top: 200rpx;
  border-radius: 20rpx;
}
.zxian {
  width: 4rpx;
  height: 80rpx;
  background: #f6f6f6;
}
.viptop {
  height: 100rpx;
  background-color: #ffffff;
}
.topcenttop {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  padding: 20rpx 20rpx;
  border-bottom: 1rpx solid #f6f6f6;
}
.topcent {
  width: 710rpx;
  margin-top: 20rpx;
  margin-bottom: 20rpx;
  margin-left: 20rpx;
  background-color: #ffffff;
  border-radius: 10rpx;
}

.usertop {
  width: 750rpx;
  height: 280rpx;
  border-bottom-left-radius: 20rpx;
  border-bottom-right-radius: 20rpx;
  background-image: url("/static/img/user/zbg.png");
  background-repeat: repeat-x;
}
.userli {
  height: 60px;
  border-bottom: 1rpx solid #f6f6f6;
  padding: 0 10px;
}
.userlist {
  width: 710rpx;
  margin-left: 20rpx;
  background-color: #ffffff;
  border-radius: 10rpx;
}
.iconlist {
  width: 30px;
}
.user_cent {
  width: 86%;
}
</style>
