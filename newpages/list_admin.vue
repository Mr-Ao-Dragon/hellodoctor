<template>
  <view>
    <view class="list">
      <view class="item" v-for="(item,index) in 2" :key="index" @click="goto(item)">
        <view class="one">
          <view class="title">
            用户{{index}}
          </view>
          <view class="status">
            {{index===0?"普通用户":"医生"}}
          </view>
        </view>
        <view class="two">
            注册时间：2024年1月11日 13:34:23
        </view>
        <view class="three">
        </view>
      </view>
    </view>
  </view>
</template>
  <script>
export default {
  data() {
    return {

    }
  },
  methods: {
    goto(item) {
      uni.navigateTo({
        url: '/newpages/info_user?id=' + item
      })
    }
  }
}
  </script>
  <style lang="scss" scoped>
.list {
  .item {
    display: flex;
    flex-direction: column;
    background: #fff;
    margin: 12px;
    box-shadow: 0 0 2px rgba(0, 0, 0, 0.2);
    line-height: 72rpx;
    text-align: center;
    border-radius: 4px;
    .one {
      display: flex;
      justify-content: space-between;
      padding: 0 12px;
      .title {
        font-weight: 700;
      }
      .status {
        color: green;
      }
    }
    .two {
      @extend .one;
      .time {
      }
    }
    .three {
      @extend .one;
      flex-direction: column;
      font-size: 12px;
      align-items: start;
      line-height: 25px;
      color: #909399;
      .name {
      }
      .mobile {
      }
    }
  }
}
</style>