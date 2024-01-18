import axios from 'axios';

axios.interceptors.request.use((config)=>{
    uni.showLoading({title:"请求中"})
    const token = uni.getStorageSync('token')?uni.getStorageSync('token'):""
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
    postLogin:async(code)=>{
        return await axios.post('/api/login',{code:code})
    },
    getNotify:async()=>{
        return await axios.get('/api/notify')
    },
    getDocList:async()=>{
        return await axios.get('/api/doclist')
    },
    /**
     * 
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
    getReserve:async()=>{
        return await axios.get(`/api/reserve`)
    },
    deleteReserve:async(id)=>{
        return await axios.delete(`/api/reserve/?id=${id}`)
    },
    getDocReserve:async()=>{
        return await axios.get(`/api/reserve/doc`)
    },
    updateReserve:async(id,status)=>{
        return await axios.put(`/api/reserve`,{
            id:id,
            status:status
        })
    },
    getUserList:async()=>{
        return await axios.get(`/api/userlist`)
    },
    /**
     * 
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