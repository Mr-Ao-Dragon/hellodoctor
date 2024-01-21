<template>
  <view>
    <view class="list">
      <view class="item" v-for="(item,index) in 2" :key="index" @click="goto(item)">
        <view class="one">
          <view class="title">
            预约 王医生
          </view>
          <view class="status">
            状态：已核销 >
          </view>
        </view>
        <view class="two">
          <view class="time">
            时间：2020-05-0{{index+1}} 12:00
          </view>
        </view>
        <view class="three">
          <view class="user">
            预约用户：张三
          </view>
          <view class="mobile">
            电话：13812345678
          </view>
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
        url: '/newpages/info_reserve?id='+item
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