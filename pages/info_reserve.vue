<template>
  <view class="body">
    <view class="block">
      <view class="title">
        <text class="icon u-font-42 zhuti_color" style="margin-right:12px">&#xe85a;</text>
        <text class="text">
          订单信息
        </text>
      </view>
      <view class="list">
        <view class="item">
          预约ID: {{info.id}}
        </view>
        <view class="item">
          预约项目: {{info.doc_name}}
        </view>
        <view class="item">
          提交时间: {{info.createTime}}
        </view>
        <view class="item">
          订单状态: {{info.status}}
        </view>
      </view>
    </view>
    <view class="block">
      <view class="title" >
        <text class="icon u-font-42 zhuti_color" style="margin-right:12px">&#xe85a;</text>
        <text class="text">
          预约信息
        </text>
      </view>
      <view class="list" > 
        <view class="item">
          联系人姓名: {{info.name}}
        </view>
        <view class="item">
          联系人手机号: {{info.mobile}}
        </view>
        <view class="item">
          预约时间: {{info.date}}
        </view>
      </view>
    </view>
    <view class="block2">
      <view class="title">
        <text class="icon u-font-42 zhuti_color" style="margin-right:12px">&#xe85a;</text>
        <text class="text">
          订单进度
        </text>
      </view>
      <view class="jindu">
        <view class="radio">
          <view class="icon" :style="{backgroundColor:info.status>0?'#FA3534':''}">
            1
          </view>
          预约
        </view>
        <view class="line">
        </view>
        <view class="radio">
          <view class="icon" :style="{backgroundColor:info.status>2?'#FA3534':''}">
            2
          </view>
          进行中
        </view>
        <view class="line">
        </view>
        <view class="radio" :style="{backgroundColor:info.status>3?'#FA3534':''}">
          <view class="icon">
            3
          </view>
          已完成
        </view>
      </view>
    </view>
  </view>
</template>
  
  <script>
import { getReserveStatusById } from '../api/dictionary'
import request from '../api/request'

export default {
  data() {
    return {
      info: {
        status:1,
      },
    }
  },
  methods: {
    formSubmit: function (e) {
      e.detail.value.isDoc
    },
    getStatus(id) {
      return getReserveStatusById(id)
    },
    cannelOrder() {
      request.updateReserve({ id: this.info.id, status: 6 })
    },
    cannelOrderByDoc() {
      request.updateReserve({ id: this.info.id, status: 7 })
    },
    confirmOrderByDoc() {
      request.updateReserve({ id: this.info.id, status: 2 })
    }
  },
  async onLoad(options) {
    this.info = options
    console.log(options)
  }
}
  </script>
  
  <style lang="scss" scoped>
.body {
  margin: 16px;
  .title {
    margin: 12px;
    padding-bottom: 12px;
    border-bottom: 1px solid #f5f5f5;
  }
  .block2{
.jindu{
  display: flex;
  justify-content: space-between;
  align-items: center;
  text-align: center;
  .line{
    flex-basis: 20%;
    height: 2px;
    background-color: #909399;
  }
  .radio{
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    text-align: center;
    flex-basis: 15%;
    .icon{
      background-color: #999999;
      border-radius: 50%;
      width: 36px;
      height: 36px;
      line-height: 36px;
      color: #fff;
    }
  }
}    
  }
  .list {
    margin: 12px;
    display: flex;
    flex-direction: column;
    font-size: 18px;
    .item {
      margin: 12px 0;
    }
  }
}
</style>