<template>
  <view class="body">

    <form @submit="formSubmit">
      <view class="one">
        <view>预约医生</view>
        <view>{{ name }}</view>
      </view>
      <view class="three">
        <view>姓名</view>
        <input name="name" placeholder="输入个人姓名" />
      </view>
      <view class="three">
        <view>联系电话</view>
        <input name="mobile" type="number" maxlength="11" placeholder="输入联系电话" />
      </view>
      <view class="three">
        <view>预约时间</view>
        <uni-datetime-picker name="time" returnType="timestamp" :start=nowDate :end="nowDate + (3600 * 24 * 1000*7)" v-model="time" :clear-icon="false" />
      </view>
      <view class="four">
        <button form-type="submit">提交</button>
      </view>
    </form>
  </view>
</template>
  
  <script>
import request from '../api/request'

export default {
  data() {
    return {
      name: "",
      id: null,
      time: null,
      nowDate: null,
    }
  },
  methods: {
    formSubmit(e) {
      console.log(e.detail.value)
      // e.detail.value.id = this.id
      if (e.detail.value.name == "") {
        uni.showToast({
          title: '请输入姓名',
          icon: 'none'
        })
      } else if (e.detail.value.mobile == "" || e.detail.value.mobile.length != 11) {
        uni.showToast({
          title: '请输入正确的手机号',
          icon: 'none'
        })
      } else{
        request.postReserve({
          name:e.detail.value.name,
          mobile:e.detail.value.mobile,
          id:this.id,
          time:this.time
        })
      }
    },
    onLoad(options) {
      this.nowDate = new Date().getTime()
      console.log(options)
      this.id = options.id
      this.name = options.name
    }
  }
}
  </script>
  
  <style lang="scss" scoped>
.body {
  margin: 8px;
  padding: 8px;
  box-shadow: 0 0 2px 0 rgba(0, 0, 0, 0.2);
  line-height: 50px;
  .one {
    display: flex;
    align-items: center;
    justify-content: space-around;
    margin: 8px 0;
    view {
      flex: 1;
    }
    .swi {
      scale: 0.7;
    }
  }
  .two {
    @extend .one;
  }
  .three {
    @extend .one;
    view {
      flex: 1;
    }
    input {
      flex: 1;
      border: 1px solid #dcdcdc;
      text-indent: 2em;
      padding: 8px 12px;
      border-radius: 4px;
    }
  }
  button {
    width: 100%;
    background: #007aff;
    color: #fff;
    border-radius: 4px;
    border: none;
    margin: 8px 0;
  }
}
</style>