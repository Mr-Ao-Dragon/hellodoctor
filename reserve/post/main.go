package main

// 导入依赖包
import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aliyun/fc-runtime-go-sdk/fc"

	"github.com/Mr-Ao-Dragon/hellodoctor/database"
	"github.com/Mr-Ao-Dragon/hellodoctor/reserve"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/commonData/ContentType"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/dataProcess"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
)

// 处理请求的函数
// ctx: 上下文信息
// event: 包含请求参数的结构体
// 返回值: 处理结果的结构体和可能发生的错误
func handleRequest(ctx context.Context, event datastruct.EventStruct) (repose *datastruct.UniversalRepose, err error) {
	// 验证登录状态
	token, err := event.ReadToken()
	if err != nil {
		repose.StatusCode = http.StatusBadRequest
		repose.Headers["ContentType"] = ContentType.TextUTF8
		repose.IsBase64Encoded = false
		repose.Body = "无法读取 Token"
		return repose, nil
	}
	OpenID, err := dataProcess.CutOpenID(token)
	auth := &datastruct.AuthStruct{
		SystemToken: token,
		OpenID:      OpenID,
	}
	loginStatus, err := database.QueryLogin(auth)
	if loginStatus || err != nil {
		// 登录失败或存在错误，返回未授权响应
		repose.StatusCode = http.StatusUnauthorized
		repose.Body = "Unauthorized"
		repose.Headers["ContentType"] = ContentType.TextUTF8
		repose.IsBase64Encoded = false
		err = nil
		return
	}
	// 复制请求数据到储备结构体
	reserveData := new(datastruct.AddReserve)
	reserveData.Token, err = event.ReadToken()
	reserveData.DoctorID = event.QueryParameters["doctor_id"].(string)
	reserveData.Mobile = event.QueryParameters["mobile"].(int64)
	reserveData.Name = event.QueryParameters["name"].(string)
	reserveData.Time = event.QueryParameters["time"].(int64)
	// 提交预约信息
	newReserve, err := reserve.PostReserve(reserveData, auth)
	ReserveByte, _ := json.Marshal(newReserve)
	// 预约成功，返回结果
	repose.StatusCode = http.StatusOK
	repose.Body = string(ReserveByte)
	repose.Headers["ContentType"] = ContentType.JsonUTF8
	repose.IsBase64Encoded = false
	return repose, nil
}

// 主函数，启动服务
func main() {
	fc.Start(handleRequest)
}
