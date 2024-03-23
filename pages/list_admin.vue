<template>
  <view>
    <view class="list">
      <view class="item" v-for="(item,index) in list" :key="index" @click="goto(item)">
        <view class="one">
          <view class="title">
            {{ item.status?item.doc_name:item.name }}
          </view>
          <view class="status">
            {{getStatus(item.status)}}
          </view>
        </view>
        <view class="three" v-if="item.status">
          <view>{{ item.name }}</view>
        </view>
      </view>
    </view>
  </view>
</template>
  <script>
import { getUserStatusById } from '@/api/dictionary';
import request from '../api/request';

export default {
  data() {
    return {
      list: [],
      ready: false
    }
  },
  methods: {
    goto(item) {
      uni.navigateTo({
        url: `/pages/info_user?id=${item.id}&name=${item.name}&doc_name=${item.doc_name}&doc_avatar=${item.doc_avatar}&status=${item.status}`
      })
    },
    getStatus(id) {
      return getUserStatusById(id)
    }
  },
  async onLoad() {
    // this.list = await request.getUserList()
    this.list = [{ id: 1, name: "张三", status: 0, doc_name: "", doc_avatar: "" },
    { id: 2, name: "李四", status: 1, doc_name: "王医生", doc_avatar: "" }]
    //为列表排序
    this.list.sort((a, b) => {
      if (a.status === 1 && b.status !== 1) {
        return -1; 
      } else if (a.status !== 1 && b.status === 1) {
        return 1; 
      } else if (a.status === 5 && b.status !== 5) {
        return -1; 
      } else if (a.status !== 5 && b.status === 5) {
        return 1; 
      } else {
        return 0; 
      }
    });
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