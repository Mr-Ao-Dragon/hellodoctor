<template>
    <view class="body">
  
      <form @submit="formSubmit">
        <view class="one">
          <view>预约状态</view>
            <view>{{getStatus(info.status)}}</view>
        </view>
        <view class="two">
          <view>预约人 {{ info.name }}</view>
            <view>预约人信息</view>
        </view>
        <view class="three">
            这里补充预约单的详情信息，看看要返回什么，用户在这里查看，可取消，医生在这里可核销
        </view>
      </form>
    </view>
  </template>
  
  <script>
import { getReserveStatusById } from '../api/dictionary'
import request from '../api/request'

  export default {
    data(){
      return{
        info:{},
      }
    },
    methods:{
      formSubmit: function(e) {
        e.detail.value.isDoc
      },
      getStatus(id){
        return getReserveStatusById(id)
      },
      cannelOrder(){
        request.updateReserve({id:this.info.id,status:6})
      },
      cannelOrderByDoc(){
        request.updateReserve({id:this.info.id,status:7})
      },
      confirmOrderByDoc(){
        request.updateReserve({id:this.info.id,status:2})
      }
    },
    async onLoad(options){
      this.info = options
      console.log(options)
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