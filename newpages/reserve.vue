<template>
    <view class="body">
  
      <form @submit="formSubmit">
        <view class="one">
          <view>预约医生</view>
          <view>王医生</view>
        </view>
        <view class="three">
          <view>姓名</view>
          <input name="input" placeholder="输入个人姓名" />
        </view>
        <view class="three">
          <view>联系电话</view>
          <input name="input" placeholder="输入联系电话" />
        </view>
        <view class="three">
            <view>预约时间</view>
            <input name="input" placeholder="选择预约时间" />
        </view>
        <view class="four">
          <button form-type="submit">提交</button>
        </view>
      </form>
    </view>
  </template>
  
  <script>
  export default {
    methods:{
      formSubmit: function(e) {
        e.detail.value.isDoc
      }
    }
  }
  </script>
  
  <style lang="scss" scoped>
  .body{
    margin: 8px;
    padding: 8px;
    box-shadow: 0 0 2px 0 rgba(0,0,0,0.2);
    line-height: 50px;
    .one{
      display: flex;
      align-items:center ;
      justify-content: space-around;
      margin: 8px 0;
      view{
        flex:1
      }
      .swi{
        scale: 0.7;
      }
    }
    .two{
      @extend .one;
    }
    .three{
      @extend .one;
      view{
        flex: 1;
      }
      input{
        flex: 1;
        border: 1px solid #dcdcdc;
        text-indent: 2em;
        padding:8px 12px;
        border-radius: 4px;
      }
    }
    button{
      width: 100%;
      background: #007aff;
      color: #fff;
      border-radius: 4px;
      border: none;
      margin: 8px 0;
    }
  }
  
  </style>