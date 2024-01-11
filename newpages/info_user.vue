<template>
  <view class="body">

    <form @submit="formSubmit">
      <view class="one">
        <view>是否为医生</view>
          <switch class="swi" name="isDoc"/>
      </view>
      <view class="two">
        <view>医生状态</view>
        <radio-group name="radio">
          <label>
            <radio value="1" /><text>启用</text>
          </label>
          <label>
            <radio value="2" /><text>停用</text>
          </label>
        </radio-group>
      </view>
      <view class="three">
        <view>医生姓名</view>
        <input name="input" placeholder="输入医生姓名" />
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