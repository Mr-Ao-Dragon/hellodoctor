import http from './interface'

//公告详情
export const indeximg = (data) => {
    return http.request({
        url: 'resource/zong_imglist',
        method: 'GET', 
        data,
		// handle:true
    })
}
export const noti_info = (data) => {
    return http.request({
        url: 'resource/noti_info',
        method: 'GET', 
        data,
		// handle:true
    })
}
//系统设置
export const login_user = (data) => {
    return http.request({
        url: 'resource/login_user',
        method: 'POST', 
        data,
		// handle:true
    })
}
export const searchdata = (data) => {
    return http.request({
        url: 'resource/searchdata',
        method: 'POST', 
        data,
		// handle:true
    })
}

export const add_order = (data) => {
    return http.request({
        url: 'resource/addorder',
        method: 'POST', 
        data,
		// handle:true
    })
}
export const changepwd = (data) => {
    return http.request({
        url: 'resource/changepwd',
        method: 'POST', 
        data,
    })
}
export const classlist = (data) => {
    return http.request({
        url: 'resource/classlist',
        method: 'POST', 
        data,
    })
}
export const myorderlist = (data) => {
    return http.request({
        url: 'resource/myorderlist',
        method: 'POST', 
        data,
    })
}
export const getydate = (data) => {
    return http.request({
        url: 'resource/getydate',
        method: 'GET', 
        data,
		// handle:true
    })
}
export const getytime = (data) => {
    return http.request({
        url: 'resource/getytime',
        method: 'POST', 
        data,
    })
}
export const yuyue_info = (data) => {
    return http.request({
        url: 'resource/yuyue_info',
        method: 'GET', 
        data,
    })
}
export const getyuyue = (data) => {
    return http.request({
        url: 'resource/getyuyue',
        method: 'GET', 
        data,
    })
}
//登录  发送验证码
export const smssend = (data) => {
    return http.request({
        url: 'resource/sendmsg',
        method: 'POST', 
        data,
		// handle:true
    })
}
//登录
export const phonelogin = (data) => {
    return http.request({
        url: 'resource/phonelogin',
        method: 'POST', 
        data,
    })
}
export const getform = (data) => {
    return http.request({
        url: 'resource/getform',
        method: 'GET', 
        data,
    })
}
export const add_yuyue = (data) => {
    return http.request({
        url: 'resource/add_yuyue',
        method: 'POST', 
        data,
    })
}
export const my_yuyue = (data) => {
    return http.request({
        url: 'resource/my_yuyue',
        method: 'POST', 
        data,
		// handle:true
    })
}
export const quxiao_yuyue = (data) => {
    return http.request({
        url: 'resource/quxiao_yuyue',
        method: 'POST', 
        data,
		// handle:true
    })
}
export const myorderinfo = (data) => {
    return http.request({
        url: 'resource/myorderinfo',
        method: 'POST', 
        data,
    })
}
export const xiaodata = (data) => {
    return http.request({
        url: 'resource/xiaodata',
        method: 'POST', 
        data,
    })
}
export const hexiao = (data) => {
    return http.request({
        url: 'resource/hexiao',
        method: 'POST', 
        data,
    })
}
export const user_info = (data) => {
    return http.request({
        url: 'resource/user_info',
        method: 'POST', 
        data,
    })
}
export const getsite = (data) => {
    return http.request({
        url: 'resource/getsite',
        method: 'POST', 
        data,
    })
}
export const setpaylist = (data) => {
    return http.request({
        url: 'resource/setpaylist',
        method: 'POST', 
        data,
    })
}
export const yuelist = (data) => {
    return http.request({
        url: 'resource/yuelist',
        method: 'POST', 
        data,
    })
}

/**邀请**/
export const getinvitation = (data) => {
    return http.request({
        url: 'resource/getinvitation',
        method: 'POST', 
        data
    })
}
export const addcode = (data) => {
    return http.request({
        url: 'resource/addcode',
        method: 'POST', 
        data
    })
}
export const invitalist = (data) => {
    return http.request({
        url: 'resource/invitalist',
        method: 'POST', 
        data
    })
}
export const addinvata = (data) => {
    return http.request({
        url: 'resource/addinvata',
        method: 'POST', 
        data
    })
}
export const getinvata = (data) => {
    return http.request({
        url: 'resource/getinvata',
        method: 'POST', 
        data
    })
}

// 默认全部导出  import api from '@/common/vmeitime-http/'
export default {
	indeximg,
	noti_info,
	login_user,
	searchdata,
	add_order,
	changepwd,
	classlist,
	myorderlist,
	getydate,
	getytime,
	yuyue_info,
	getyuyue,
	smssend,
	phonelogin,
	getform,
	add_yuyue,
	my_yuyue,
	quxiao_yuyue,
	myorderinfo,
	xiaodata,
	hexiao,
	user_info,
	getsite,
	setpaylist,
	yuelist,
	getinvitation,addcode,invitalist,addinvata,getinvata
}