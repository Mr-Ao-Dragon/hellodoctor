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
        <view @click="showPicker=!showPicker">{{ timeword?timeword:"选择预约时间" }}</view>
      </view>

      <view class="four">
        <button form-type="submit">提交</button>
      </view>

      <ziohtimepicker :start-time="'8:30'" :end-time="'17:30'" :step="15" :after-hours="2" :show="showPicker" @cancel="showPicker=!showPicker" @confirm="confirmTime" />
      <view :class="{shadowClass:showPicker}" @click="showPicker=!showPicker">
      </view>
    </form>
  </view>
</template>
  
  <script>
import request from '../api/request'
import ziohtimepicker from '../components/ziohtimepicker.vue'

export default {
  components: { ziohtimepicker },
  data() {
    return {
      name: "",
      id: null,
      time: null,
      timeword: null,
      showPicker: false
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
      } else {
        request.postReserve({
          name: e.detail.value.name,
          mobile: e.detail.value.mobile,
          id: this.id,
          time: this.time
        })
      }
    },
    confirmTime(e) {
      this.showPicker = !this.showPicker
      this.timeword = e.substring(5, 16)
      this.time = new Date(e).getTime()

      console.log({ e, word: this.timeword, time: this.time })
    }
  },
  onLoad(options) {
    this.nowDate = new Date().getTime()
    console.log(options)
    this.id = options.id
    this.name = options.name
  }
}
  </script>
  
  <style lang="scss" scoped>
.shadowClass {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  bottom: 0;
  right: 0;
  background: rgba(0, 0, 0, 0.5);
}
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