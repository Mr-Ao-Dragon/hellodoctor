package datastruct

type AddReserve struct {
	DoctorID string `json:"doctor_id"  copier:"DoctorID"` // 医生ID
	Mobile   int64  `json:"mobile"     copier:"Mobile"`   // 手机号
	Name     string `json:"name"       copier:"Name"`     // 预约者名字
	Time     int64  `json:"time"       copier:"Time"`     // 预约unix时间戳
	Token    string `json:"token"`                        // 预约者本次登录凭证
}

type CancelReserve struct {
	ID    int64  `json:"id"`    // 目标单号id
	Token string `json:"token"` // 操作者token
}
type InfoReserve struct {
	Code *int64          `json:"code"` // 操作状态码
	Data []SingleReserve `json:"data"` // 切片数据
	Msg  *string         `json:"msg"`  // 操作消息
}

type SingleReserve struct {
	DocAvatar  string `json:"doc_avatar"` // 医生头像
	DoctorID   string `json:"doc_id"`     // 医生OpenID
	DoctorName string `json:"doc_name"`   // 医生名称
	ID         int32  `json:"id"`         // 订单id
	Mobile     int64  `json:"mobile"`     // 预约者手机号
	Name       string `json:"name"`       // 预约者名称
	Status     int8   `json:"status"`     // 状态
	Time       int64  `json:"time"`       // 预约目标时间戳
}
