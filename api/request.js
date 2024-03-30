import axios from 'axios';
//axios的请求配置

axios.interceptors.request.use((config)=>{
    uni.showLoading({title:"请求中"})
    const token = uni.getStorageSync('token')?uni.getStorageSync('token'):""
    //需要改动的就是这里填写后端的基础地址
    config.baseURL= 'v1'
    config.headers.Authorization = 'Bearer ' + token
    config.headers['Content-Type']= 'application/json'
    return config
},(error)=>{return Promise.reject(error)})
axios.interceptors.response.use((response)=>{
    uni.hideLoading()
    return response.data.data
},async(error)=>{
    uni.hideLoading()
    await uni.showToast({title:error,icon:"none"})
    return Promise.reject(error)})
const request = {
    //登录
    postLogin:async(code)=>{
        return await axios.get(`/v1/user/login?token=${code}`)
    },
    //获取通知
    getNotify:async()=>{
        return await axios.get('/v1/api/notify')
    },
    //获取医生列表
    getDocList:async()=>{
        return await axios.get('/v1/doctor/index/list')
    },
    /**
     * 预约申请
     * @param {Object} data 
     * @param {string} data.name
     * @param {number} data.id
     * @param {number} data.doc_id
     * @param {number} data.mobile
     * @param {string} data.date
     */
    postReserve:async(data)=>{
        return await axios.post('/v1/rese/new',data)
    },
    //获取订单
    getReserve:async()=>{
        return await axios.get(`/v1/rese/list`)
    },
    /*
    * 删除订单
    * @param {number} id
    * @param {string} token
    * */
    deleteReserve:async(id)=>{
        return await axios.delete(`/v1/reve/cancel?id=${id}`)
    },
    //获取订单(医生)
    getDocReserve:async()=>{
        return await axios.get(`/v1/api/reserve/doc`)
    },
    //更新订单状态
    updateReserve:async(id,status)=>{
        return await axios.put(`/v1/rese/update`,{
            id:id,
            status:status
        })
    },
    //获取用户列表(管理员)
    getUserList:async()=>{
        return await axios.get(`/v1/admin/user_list`)
    },
    /**
     * 更新用户状态
     * @param {number} id 
     * @param {Object} data 
     * @param {string} doc_name
     * @param {number} status
     */
    UpdateUser:async(id,data)=>{
        return await axios.put(`/v1/admin/mod_user_perm${id}`,data)
    }
}
export default request