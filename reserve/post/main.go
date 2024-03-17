package main

// 导入依赖包
import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/jinzhu/copier"

	"github.com/Mr-Ao-Dragon/hellodoctor/database"
	"github.com/Mr-Ao-Dragon/hellodoctor/reserve"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/commonData/ContentType"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
)

// 定义请求体结构体
type (
	ReqBody struct {
		DoctorID string `json:"doctor_id"   copier:"DoctorID"` // 医生ID
		OpenID   string `json:"openid"      copier:"OpenID"`   // 用户ID
		Mobile   int64  `json:"mobile"      copier:"Mobile"`   // 手机号
		Name     string `json:"name"        copier:"Name"`     // 预约者名字
		Time     int64  `json:"time"        copier:"Time"`     // 预约unix时间戳
		Token    string `json:"token"`
	}
	structEvent struct {
		QueryParameters ReqBody `json:"queryParameters"`
		Body            string  `json:"body"`
	}
)

// 处理请求的函数
// ctx: 上下文信息
// event: 包含请求参数的结构体
// 返回值: 处理结果的结构体和可能发生的错误
func handleRequest(ctx context.Context, event structEvent) (repose *datastruct.UniversalRepose, err error) {
	// 验证登录状态
	auth := &datastruct.AuthStruct{
		SystemToken: event.QueryParameters.Token,
		OpenID:      event.QueryParameters.OpenID,
	}
	loginStatus, err := database.QueryLogin(auth)
	if loginStatus || err != nil {
		// 登录失败或存在错误，返回未授权响应
		repose.StatusCode = http.StatusUnauthorized
		repose.Body = "Unauthorized"
		repose.Headers.ContentType = ContentType.TextUTF8
		repose.IsBase64Encoded = false
		err = nil
		return
	}

	// 复制请求数据到储备结构体
	reserveData := new(datastruct.AddReserve)
	err = copier.Copy(&reserveData, &event.QueryParameters)
	if err != nil {
		// 数据复制失败，返回内部服务器错误
		repose.StatusCode = http.StatusInternalServerError
		repose.Body = "Copy data failed"
		repose.Headers.ContentType = ContentType.TextUTF8
		repose.IsBase64Encoded = false
		err = nil
		return
	}

	// 提交预约信息
	newReserve, err := reserve.PostReserve(reserveData, auth)
	ReserveByte, _ := json.Marshal(newReserve)
	// 预约成功，返回结果
	repose.StatusCode = http.StatusOK
	repose.Body = string(ReserveByte)
	repose.Headers.ContentType = ContentType.JsonUTF8
	repose.IsBase64Encoded = false
	return repose, nil
}

// 主函数，启动服务
func main() {
	fc.Start(handleRequest)
}
