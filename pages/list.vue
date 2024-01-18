<template>
  <view>
    <view class="list">
      <view class="item" v-for="(item,index) in list" :key="index" @click="goto(item)">
        <view class="one">
          <view class="title">
            预约 {{ list.doc_name }}
          </view>
          <view class="status">
            状态：{{ getStatus(item.status)}} >
          </view>
        </view>
        <view class="two">
          <view class="time">
            预约时间：{{ day(item.time)}}
          </view>
        </view>
        <view class="three">
          <view class="user">
            预约用户：{{item.name}}
          </view>
          <view class="mobile">
            预留电话：{{ item.mobile }}
          </view>
        </view>
      </view>
    </view>
  </view>
</template>
<script>
import request from '../api/request'
import dayjs from 'dayjs'
import {getReserveStatusById} from '../api/dictionary'
export default {
  data() {
    return {
      list: [],
    }
  },
  methods: {
    goto(item) {
      uni.navigateTo({
        url: `/pages/info_reserve?by=user&id=${item.id}&doc_name=${item.doc_name}&doc_avatar=${item.doc_avatar}&doc_id=${item.doc_id}&name=${item.name}&mobile=${item.mobile}&time=${item.time}&status=${item.status}`,
      })
    },
    getStatus(status){
      return getReserveStatusById(status)
    },
    day(time){
      return dayjs.unix(time).format('MM月DD日 HH点')
    }
  },
  async onLoad() {
    // this.list = await request.getReserve()
    this.list = [
      {
      "id": 1,
      "doc_name": "王医生",
      "doc_avatar": "http://www.baidu.com",
      "doc_id": 1,
      "name": "张三",
      "mobile": 1388888888,
      "time": 1615945600,
      "status": 1
    },{
      "id": 1,
      "doc_name": "王医生",
      "doc_avatar": "http://www.baidu.com",
      "doc_id": 1,
      "name": "张三",
      "mobile": 1388888888,
      "time": 1615945600,
      "status": 1
    }
    ]
    await uni.hideLoading()
    this.ready = true
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