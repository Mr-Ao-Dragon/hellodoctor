package main

import (
	"context"
	"encoding/json"

	"github.com/aliyun/fc-runtime-go-sdk/fc"

	"github.com/Mr-Ao-Dragon/hellodoctor/database"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
	"github.com/Mr-Ao-Dragon/hellodoctor/user"
)

type StructEvent struct {
	Body string `json:"body"`
}
type RequestBody struct {
	TargetOpenID     string `json:"target_openid"`
	OperatorOpenID   string `json:"operator_openid"`
	SystemToken      string `json:"system_token"`
	TargetPermission int16  `json:"target_permission"`
}
type ReposeErrBody struct {
	Code int16  `json:"code"`
	Body string `json:"body"`
}
type ReposeBody struct {
	Code int16  `json:"code"`
	Body string `json:"body"`
}

func HandleHttpRequest(ctx context.Context, event StructEvent) (repose string, err error) {
	var Request RequestBody
	err = json.Unmarshal([]byte(event.Body), &Request)
	if err != nil {
		return "", err
	}
	authData := &datastruct.AuthStruct{
		SystemToken: Request.SystemToken,
		OpenID:      Request.OperatorOpenID,
	}
	isSusses, err := database.QueryLogin(authData)
	if isSusses == false {
		ReposeErr := ReposeErrBody{Code: 403, Body: "未登录，拒绝访问"}
		ReposeErrBody, _ := json.Marshal(ReposeErr)
		return string(ReposeErrBody), nil
	}
	if err != nil {
		ReposeErr := ReposeErrBody{Code: 503, Body: "查询失败，服务器内部错误"}
		ReposeErrBody, _ := json.Marshal(ReposeErr)
		return string(ReposeErrBody), nil
	}
	PermissionLevel, err := database.QueryPermission(Request.OperatorOpenID)
	if PermissionLevel <= 3 {
		ReposeErr := ReposeErrBody{Code: 403, Body: "权限不足，禁止操作"}
		ReposeErrBody, _ := json.Marshal(ReposeErr)
		return string(ReposeErrBody), nil
	}
	if err != nil {
		ReposeErr := ReposeErrBody{Code: 503, Body: "查询失败，服务器内部错误"}
		ReposeErrBody, _ := json.Marshal(ReposeErr)
		return string(ReposeErrBody), nil
	}
	var Permission = new(user.PermLevel)
	switch Request.TargetPermission {
	case 0:
		Permission.Set().ToBaned(Request.TargetOpenID)
	case 1:
		Permission.Set().ToUser(Request.TargetOpenID)
	case 2:
		Permission.Set().ToAssistant(Request.TargetOpenID)
	case 3:
		Permission.Set().ToDoctor(Request.TargetOpenID)
	case 4:
		Permission.Set().ToSystemAdmin(Request.TargetOpenID)
	default:
		ReposeErr := ReposeErrBody{Code: 400, Body: "错误的权限，请联系管理员。"}
		ReposeErrBody, _ := json.Marshal(ReposeErr)
		return string(ReposeErrBody), nil
	}
	ReposeStruct := ReposeBody{Code: 200, Body: "成功！"}
	ReposeBody, err := json.Marshal(ReposeStruct)
	return string(ReposeBody), nil

}
func main() {
	fc.Start(HandleHttpRequest)
}
