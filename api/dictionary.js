const userStatus = [{
    value: 0,
    text: "用户"
}, {
    value: 1,
    text: "医生，已启用"
}, {
    value: 5,
    text: "已停用"
}]
const reserveStatus = [{
    value: 0,
    text: "错误"
}, {
    value: 1,
    text: "已发起"
}, {
    value: 2,
    text: "已核销"
}, {
    value: 5,
    text: "过期未核销"
}, {
    value: 6,
    text: "用户已取消"
}, {
    value: 7,
    text: "医生已取消"
},]

export const getUserStatusById = (id) => {
    return userStatus.filter(item => item.value === Number(id))[0].text
}
export const getUserStatusByText = (id) => {
    return userStatus.filter(item => item.value === Number(id))[0].value
}
export const getReserveStatusById = (id) => {
    return reserveStatus.filter(item=>item.value===Number(id))[0].text
}
export const getReserveStatusByText = (id) => {
    return reserveStatus.filter(item => item.value === Number(id))[0].value
}