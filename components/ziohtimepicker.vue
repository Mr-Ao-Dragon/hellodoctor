<template>

  <view class="w-time-picker" :class="{'show':show}" :style="{display:show?'block':'none'}">
    <view class="w-time-days">
      <scroll-view scroll-x>
        <view class="w-time-scroll">
          <view class="w-time-day" :style="{'color':tabIndex==index?theme:'#333','border-color':tabIndex==index?theme:'#ddd'}" v-for="(day,index) in dayList" :key="index" @click="toggleTab(day,index)">
            <view class="w-time-week">{{day.week}}</view>
            <view class="w-time-date">{{day.month}}/{{day.day}}</view>
          </view>
        </view>
      </scroll-view>
    </view>
    <view class="w-time-body">
      <scroll-view scroll-y class="w-time-list-scroll">
        <view class="w-time-list">
          <view class="w-time-item" :style="{'color':itemIndex==index?theme:'#333','border-color':itemIndex==index?theme:'#ddd'}" :class="{'w-time-item-disable':item.disabled}" v-for="(item,index) in timeList" :key="index" @click="toggleItem(item,index)">
            <view class="time">

              {{item.label}}
            </view>
            <view class="text">空位:1</view>
          </view>
        </view>
      </scroll-view>
    </view>
    <view class="w-time-footer">
      <view class="w-time-cancel w-time-btn" @click="cancel">取消</view>
      <view class="w-time-sure w-time-btn" :style="{'background-color':theme}" @click="confirm">确定</view>
    </view>
  </view>

</template>

<script>
export default {
  props: {
    afterDays: {
      type: [String, Number],
      default: 7
    },
    startTime: {
      type: String,
      validator: function (value) {
        return /^\d+:\d+$/.test(value);
      },
      default: '07:00'
    },
    endTime: {
      type: String,
      validator: function (value) {
        return /^\d+:\d+$/.test(value);
      },
      default: '19:00'
    },
    step: {
      type: [String, Number],
      default: 30
    },
    afterHours: {
      type: [String, Number],
      default: 2
    },
    theme: {
      type: String,
      default: "#f5a200"
    },
    show: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      dayList: [],
      timeList: [],
      tabIndex: 0,
      itemIndex: -1,
      time: ""
    }
  },
  computed: {
    propsChange() {
      return [this.afterDays, this.startTime, this.endTime, this.step, this.afterHours];
    }
  },
  watch: {
    propsChange() {
      this.init()
    }
  },
  mounted() {
    this.init();
  },
  methods: {
    cancel() {
      this.$emit("cancel");
    },
    confirm() {
      if (this.time != "") {
        let tabItem = this.dayList[this.tabIndex];
        let result = tabItem.year + "-" + tabItem.month + "-" + tabItem.day + " " + this.time.label + ":00";
        this.$emit("confirm", result);
      } else {
        uni.showToast({
          title: "请选择时间",
          icon: "none"
        })
      }
    },
    toggleTab(item, index) {
      this.tabIndex = index;
      this.itemIndex = -1;
      this.time = "";
      this.initTime(item.isToday);
    },
    toggleItem(item, index) {
      if (!item.disabled) {
        this.itemIndex = index;
        this.time = item;
      }
    },
    forMatNumber(n) {
      return n < 10 ? '0' + n : n;
    },
    strtotime(str) {
      const strArr = str.split(':');
      const date = new Date();
      date.setHours(strArr[0]);
      date.setMinutes(strArr[1]);
      date.setSeconds(0);
      date.setMilliseconds(0);
      return date.getTime();
    },
    init() {
      this.initDay();
      this.initTime();
    },
    initDay() {
      let aDate = new Date();
      let weekList = ["周日", "周一", "周二", "周三", "周四", "周五", "周六"];
      this.dayList = [];
      for (let i = 0; i < this.afterDays * 1; i++) {
        i > 0 && aDate.setDate(aDate.getDate() + 1);
        this.dayList.push({
          year: aDate.getFullYear(),
          month: this.forMatNumber(aDate.getMonth() + 1),
          day: this.forMatNumber(aDate.getDate()),
          week: weekList[aDate.getDay()],
          isToday: i === 0
        })
      }
    },
    initTime(isToday = true) {
      let afterSeconds = this.afterHours * 3600 * 1000;
      let stepSeconds = this.step * 60 * 1000;
      let curTimestamp = new Date().getTime();
      let startTimestamp = this.strtotime(this.startTime);
      let endTimestamp = this.strtotime(this.endTime);
      this.timeList = [];
      while (startTimestamp <= endTimestamp) {
        let timestamp = new Date(startTimestamp);
        if (this.forMatNumber(timestamp.getHours()) === 12
          && this.forMatNumber(timestamp.getMinutes()) !== '00') {

        } else if (this.forMatNumber(timestamp.getHours()) === 13) {
        } else {
          this.timeList.push({
            label: this.forMatNumber(timestamp.getHours()) + ":" + this.forMatNumber(timestamp.getMinutes()),
            disabled: isToday ? (curTimestamp + afterSeconds) > startTimestamp : false
          });
        }
        startTimestamp += stepSeconds;
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.w-time-picker.show {
  transform: translate3d(0, 0, 0);
}

.w-time-picker {
  position: sticky;
  top: 0;
  bottom: 0;
  border-radius: 16px 16px 0 0;
  overflow: hidden;
  width: 100%;
  z-index: 9999;
  background-color: #f5f5f5;
  display: flex;
  flex-direction: column;
  transform: translate3d(0, 100%, 0);
  transition: all 0.3s ease;

  .w-time-days {
    overflow: hidden;
    padding: 20upx 0;
    background-color: #fff;

    .w-time-scroll {
      white-space: nowrap;
    }

    .w-time-day {
      display: inline-block;
      width: 120upx;
      text-align: center;
      border: solid 1px #ddd;
      margin: 0 20upx;
      border-radius: 6upx;
      padding: 10upx 0;
      color: #333;

      .w-time-week {
        font-size: 28upx;
        line-height: 1;
      }

      .w-time-date {
        font-size: 24upx;
        line-height: 1;
        margin-top: 10upx;
      }
    }

    .w-time-day-active {
      color: #f00;
      border-color: #f00;
    }
  }

  .w-time-body {
    flex: 1;
    overflow: hidden;

    .w-time-list-scroll {
      height: 100%;
    }

    .w-time-list {
      display: flex;
      flex-wrap: wrap;
      padding: 20upx 10upx;
      justify-content: start;
      
    }

    .w-time-item {
      flex-basis: 23%;
      width: 120upx;
      margin: 5px 2px;
      height: 64upx;
      line-height: 32upx;
      text-align: center;
      border: solid 1px #ddd;
      border-radius: 6upx;
      font-size: 28upx;
      transition: all 0.3s ease;
      background-color: #fff;
      display: flex;
      flex-direction: column;
    }

    .w-time-item-disable {
      color: #ddd !important;
      background-color: #f5f5f5;
    }
  }

  .w-time-footer {
    height: 88upx;
    display: flex;
    background-color: #fff;

    .w-time-btn {
      flex: 1;
      text-align: center;
      line-height: 88upx;
      font-size: 30upx;
    }

    .w-time-sure {
      background-color: #f00;
      color: #fff;
    }
  }
}
</style>