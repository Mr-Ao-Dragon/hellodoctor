(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-info_reserve"],{"390e":function(t,e,i){"use strict";i.r(e);var a=i("3a0b"),n=i("62c5");for(var s in n)["default"].indexOf(s)<0&&function(t){i.d(e,t,(function(){return n[t]}))}(s);i("4682");var u=i("f0c5"),r=Object(u["a"])(n["default"],a["b"],a["c"],!1,null,"22ae3ad7",null,!1,a["a"],void 0);e["default"]=r.exports},"3a0b":function(t,e,i){"use strict";i.d(e,"b",(function(){return a})),i.d(e,"c",(function(){return n})),i.d(e,"a",(function(){}));var a=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("v-uni-view",{staticClass:"body"},[i("v-uni-view",{staticClass:"block"},[i("v-uni-view",{staticClass:"title"},[i("v-uni-text",{staticClass:"icon u-font-42 zhuti_color",staticStyle:{"margin-right":"12px"}},[t._v("")]),i("v-uni-text",{staticClass:"text"},[t._v("订单信息")])],1),i("v-uni-view",{staticClass:"list"},[i("v-uni-view",{staticClass:"item"},[t._v("预约ID: "+t._s(t.info.id))]),i("v-uni-view",{staticClass:"item"},[t._v("预约项目: "+t._s(t.info.doc_name))]),i("v-uni-view",{staticClass:"item"},[t._v("提交时间: "+t._s(t.info.createTime))]),i("v-uni-view",{staticClass:"item"},[t._v("订单状态: "+t._s(t.info.status))])],1)],1),i("v-uni-view",{staticClass:"block"},[i("v-uni-view",{staticClass:"title"},[i("v-uni-text",{staticClass:"icon u-font-42 zhuti_color",staticStyle:{"margin-right":"12px"}},[t._v("")]),i("v-uni-text",{staticClass:"text"},[t._v("预约信息")])],1),i("v-uni-view",{staticClass:"list"},[i("v-uni-view",{staticClass:"item"},[t._v("联系人姓名: "+t._s(t.info.name))]),i("v-uni-view",{staticClass:"item"},[t._v("联系人手机号: "+t._s(t.info.mobile))]),i("v-uni-view",{staticClass:"item"},[t._v("预约时间: "+t._s(t.info.date))])],1)],1),i("v-uni-view",{staticClass:"block2"},[i("v-uni-view",{staticClass:"title"},[i("v-uni-text",{staticClass:"icon u-font-42 zhuti_color",staticStyle:{"margin-right":"12px"}},[t._v("")]),i("v-uni-text",{staticClass:"text"},[t._v("订单进度")])],1),i("v-uni-view",{staticClass:"jindu"},[i("v-uni-view",{staticClass:"radio"},[i("v-uni-view",{staticClass:"icon",style:{backgroundColor:t.info.status>0?"#FA3534":""}},[t._v("1")]),t._v("预约")],1),i("v-uni-view",{staticClass:"line"}),i("v-uni-view",{staticClass:"radio"},[i("v-uni-view",{staticClass:"icon",style:{backgroundColor:t.info.status>2?"#FA3534":""}},[t._v("2")]),t._v("进行中")],1),i("v-uni-view",{staticClass:"line"}),i("v-uni-view",{staticClass:"radio",style:{backgroundColor:t.info.status>3?"#FA3534":""}},[i("v-uni-view",{staticClass:"icon"},[t._v("3")]),t._v("已完成")],1)],1)],1)],1)},n=[]},"3a11":function(t,e,i){"use strict";i("7a82"),Object.defineProperty(e,"__esModule",{value:!0}),e.getUserStatusByText=e.getUserStatusById=e.getReserveStatusByText=e.getReserveStatusById=void 0,i("4de4"),i("d3b7"),i("a9e3");var a=[{value:0,text:"用户"},{value:1,text:"医生，已启用"},{value:5,text:"已停用"}],n=[{value:0,text:"错误"},{value:1,text:"已发起"},{value:2,text:"已核销"},{value:5,text:"过期未核销"},{value:6,text:"用户已取消"},{value:7,text:"医生已取消"}];e.getUserStatusById=function(t){return a.filter((function(e){return e.value===Number(t)}))[0].text};e.getUserStatusByText=function(t){return a.filter((function(e){return e.value===Number(t)}))[0].value};e.getReserveStatusById=function(t){return n.filter((function(e){return e.value===Number(t)}))[0].text};e.getReserveStatusByText=function(t){return n.filter((function(e){return e.value===Number(t)}))[0].value}},4682:function(t,e,i){"use strict";var a=i("4990"),n=i.n(a);n.a},4990:function(t,e,i){var a=i("e021");a.__esModule&&(a=a.default),"string"===typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);var n=i("4f06").default;n("7cd1d22c",a,!0,{sourceMap:!1,shadowMode:!1})},"62c5":function(t,e,i){"use strict";i.r(e);var a=i("da7f"),n=i.n(a);for(var s in a)["default"].indexOf(s)<0&&function(t){i.d(e,t,(function(){return a[t]}))}(s);e["default"]=n.a},da7f:function(t,e,i){"use strict";i("7a82");var a=i("4ea4").default;Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var n=a(i("c7eb")),s=a(i("1da1")),u=i("3a11"),r=a(i("861a")),o={data:function(){return{info:{status:1}}},methods:{formSubmit:function(t){t.detail.value.isDoc},getStatus:function(t){return(0,u.getReserveStatusById)(t)},cannelOrder:function(){r.default.updateReserve({id:this.info.id,status:6})},cannelOrderByDoc:function(){r.default.updateReserve({id:this.info.id,status:7})},confirmOrderByDoc:function(){r.default.updateReserve({id:this.info.id,status:2})}},onLoad:function(t){var e=this;return(0,s.default)((0,n.default)().mark((function i(){return(0,n.default)().wrap((function(i){while(1)switch(i.prev=i.next){case 0:e.info=t,console.log(t);case 2:case"end":return i.stop()}}),i)})))()}};e.default=o},e021:function(t,e,i){var a=i("24fb");e=a(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */.body[data-v-22ae3ad7]{margin:16px}.body .title[data-v-22ae3ad7]{margin:12px;padding-bottom:12px;border-bottom:1px solid #f5f5f5}.body .block2 .jindu[data-v-22ae3ad7]{display:flex;justify-content:space-between;align-items:center;text-align:center}.body .block2 .jindu .line[data-v-22ae3ad7]{flex-basis:20%;height:2px;background-color:#909399}.body .block2 .jindu .radio[data-v-22ae3ad7]{display:flex;flex-direction:column;justify-content:center;align-items:center;text-align:center;flex-basis:15%}.body .block2 .jindu .radio .icon[data-v-22ae3ad7]{background-color:#999;border-radius:50%;width:36px;height:36px;line-height:36px;color:#fff}.body .list[data-v-22ae3ad7]{margin:12px;display:flex;flex-direction:column;font-size:18px}.body .list .item[data-v-22ae3ad7]{margin:12px 0}',""]),t.exports=e}}]);