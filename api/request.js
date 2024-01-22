import axios from 'axios';
//axios的请求配置

axios.interceptors.request.use((config)=>{
    uni.showLoading({title:"请求中"})
    const token = uni.getStorageSync('token')?uni.getStorageSync('token'):""
    //需要改动的就是这里填写后端的基础地址
    config.baseURL= ''
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
        return await axios.post('/api/login',{code:code})
    },
    //获取通知
    getNotify:async()=>{
        return await axios.get('/api/notify')
    },
    //获取医生列表
    getDocList:async()=>{
        return await axios.get('/api/doclist')
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
        return await axios.post('/api/reserve',data)
    },
    //获取订单
    getReserve:async()=>{
        return await axios.get(`/api/reserve`)
    },
    //删除订单
    deleteReserve:async(id)=>{
        return await axios.delete(`/api/reserve/?id=${id}`)
    },
    //获取订单(医生)
    getDocReserve:async()=>{
        return await axios.get(`/api/reserve/doc`)
    },
    //更新订单状态
    updateReserve:async(id,status)=>{
        return await axios.put(`/api/reserve`,{
            id:id,
            status:status
        })
    },
    //获取用户列表(管理员)
    getUserList:async()=>{
        return await axios.get(`/api/userlist`)
    },
    /**
     * 更新用户状态
     * @param {number} id 
     * @param {Object} data 
     * @param {string} doc_name
     * @param {number} status
     */
    UpdateUser:async(id,data)=>{
        return await axios.put(`/api/userlist/${id}`,data)
    }
}
export default request