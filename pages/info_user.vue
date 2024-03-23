<template>
  <view class="body">

    <form @submit="formSubmit">
      <view class="two">
        <view>医生状态</view>
        <radio-group @change="changeStatus($event)" name="radio" v-for="item in status" :key="item.value">
          <label>
            <radio  :checked="item.checked" :value="item.value" /><text class="tag-name">{{item.name}}</text>
          </label>
        </radio-group>
      </view>
      <view class="three">
        <view>医生姓名</view>
        <input name="input" :disabled="status[0].checked" v-model="info.doc_name" placeholder="输入医生姓名" />
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
      info: {},
      status: [{
        value: "1",
        name: '用户',
        checked: false
      }, {
        value: "3",
        name: '医生',
        checked: false
      }, {
        value: "0",
        name: '停用',
        checked: false
      },]
    }
  },
  methods: {
    getIsDoc() { return this.isDoc },
    formSubmit: function (e) {
      const data ={
        id:this.info.id,
        doc_name:this.info.name,
      }
      data.status = Number(this.status.find(item => item.checked === true)[0].value)
      request.UpdateUser(this.id,data)
    },
    changeStatus(e) {
      console.log(e.detail)
      this.status.forEach(item => {

        item.checked=false
        if (item.value === e.detail.value) {
          console.log(item)
          item.checked = true
        }
      })
    }

  },
  onLoad(option) {
    console.log(option)
    console.log(Number(option.status))
    this.info = option
    this.status.forEach(item => { if (item.value === option.status) { item.checked = true } })
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
    .tag-name {
      margin-right: 8px;
    }
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