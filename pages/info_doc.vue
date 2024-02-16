<script>
import ziohtimepicker from '../components/ziohtimepicker.vue'
export default {
  components: { ziohtimepicker },
  data() {
    return {
      showPicker: false,
      nowDate:null,
      id:null,
      name:null
    }
  },
  methods: {
    Confirm(value) {
      console.log(value)
      uni.navigateTo({
        url: `/pages/reserve?id=${this.id}&name=${this.name}&time=${value}`
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
</script>

</script>
<template>
    <div class="container">
        <scroll-view scroll-y class="content">
            <div class="cover"><img src="../static/avatar.jpg" alt="医生头像" class="img"></div>
            <div class="bar">
                <div class="top">
                    王医生
                </div>
            </div>
            <div class="bar">
                <div class="title">服务介绍</div>
                <div class="text">这是介绍内容这是介绍内容这是介绍内容这是介绍
                    内容这是介绍内容这是介绍内容这是介绍内容这是介绍内容这是介绍内容这是介绍内容
                    内容这是介绍内容这是介绍内容这是介绍内容这是介绍内容这是介绍内容这是介绍内容
                    内容这是介绍内容这是介绍内容这是介绍内容这是介绍内容这是介绍内容这是介绍内容
                    内容这是介绍内容这是介绍内容这是介绍内容这是介绍内容这是介绍内容这是介绍内容
                    内容这是介绍内容这是介绍内容这是介绍内容这是介绍内容这是介绍内容这是介绍内容
                    内容这是介绍内容这是介绍内容这是介绍内容这是介绍内容这是介绍内容这是介绍内容
                </div>
            </div>
        </scroll-view>
        <div class="foot">
            <button class="btn" @click="showPicker=!showPicker">立即预约</button>
        </div>
        <ziohtimepicker 
        :start-time="'8:30'" 
        :end-time="'17:30'" 
        :step="15" 
        :after-hours="2" 
        :show="showPicker"
        @cancel="showPicker=!showPicker" 
        @confirm="Confirm($event)" />
    </div>
</template>
<style lang="scss" scoped>
.container {
  background-color: #f4f4f5;
  height: 100vh;
  .content {
    display: flex;
    flex-direction: column;
    height: 100%;
    .cover {
      width: 100%;
      .img {
        width: 100%;
      }
    }
    .bar {
      background-color: #ffffff;
      margin: 12px 0;
      padding: 12px;
      .top {
        display: flex;
        font-size: 18px;
        font-weight: 500;
      }
      > .title {
        font-weight: 700;
        font-size: 20px;
        color: #db3635;
        text-align: center;
        padding-bottom: 16px;
        border-bottom: #db3635 1px solid;
      }
      > .text {
        text-indent: 2em;
      }
    }
  }
  .foot {
    position: sticky;
    bottom: 0;
    width: 100%;
    background-color: #fff;
    padding: 6px 0;
    height: 60px;
    .btn {
      height: 48px;
      line-height: 48px;
      width: 90%;
      background-color: #db3635;
      color: #fff;
      border-radius: 12px;
    }
  }
}
</style>
