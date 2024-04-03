package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aliyun/fc-runtime-go-sdk/fc"

	"github.com/Mr-Ao-Dragon/hellodoctor/database"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/commonData/ContentType"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
	"github.com/Mr-Ao-Dragon/hellodoctor/user"
)

type RequestBody struct {
	TargetOpenID     string `json:"target_openid"`
	OperatorOpenID   string `json:"operator_openid"`
	SystemToken      string `json:"system_token"`
	TargetPermission int16  `json:"target_permission"`
}

func HandleHttpRequest(ctx context.Context, event datastruct.EventStruct) (repose *datastruct.UniversalRepose, err error) {
	var Request RequestBody
	repose = new(datastruct.UniversalRepose)
	_ = json.Unmarshal([]byte(event.Body), &Request)

	authData := &datastruct.AuthStruct{
		SystemToken: Request.SystemToken,
		OpenID:      Request.OperatorOpenID,
	}
	isSusses, err := database.QueryLogin(authData)
	if isSusses == false {
		repose.StatusCode = http.StatusUnauthorized
		repose.Body = "未登录，拒绝访问"
		repose.Headers["ContentType"] = ContentType.JsonUTF8
		repose.Headers["AccessControlAllowOrigin"] = "*"
		repose.IsBase64Encoded = false
		return repose, nil
	}
	if err != nil {
		repose.StatusCode = http.StatusForbidden
		repose.Body = "认证失败，拒绝访问"
		repose.Headers["ContentType"] = ContentType.JsonUTF8
		repose.Headers["AccessControlAllowOrigin"] = "*"
		repose.IsBase64Encoded = false
		return repose, nil
	}
	PermissionLevel, err := database.QueryPermission(Request.OperatorOpenID)
	if PermissionLevel <= 3 {
		repose.StatusCode = http.StatusForbidden
		repose.Body = "未登录，拒绝访问"
		repose.Headers["ContentType"] = ContentType.JsonUTF8
		repose.Headers["AccessControlAllowOrigin"] = "*"
		repose.IsBase64Encoded = false
		return repose, nil
	}
	if err != nil {
		repose.StatusCode = http.StatusServiceUnavailable
		repose.Body = "查询失败，数据库错误"
		repose.Headers["ContentType"] = ContentType.JsonUTF8
		repose.Headers["AccessControlAllowOrigin"] = "*"
		repose.IsBase64Encoded = false
		return repose, nil
	}
	Permission := new(user.PermLevel)
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
		repose.StatusCode = http.StatusBadRequest
		repose.Body = "错误的权限，请检查输入。"
		repose.Headers["ContentType"] = ContentType.JsonUTF8
		repose.Headers["AccessControlAllowOrigin"] = "*"
		repose.IsBase64Encoded = false
	}
	repose.StatusCode = http.StatusOK
	repose.Body = "成功！"
	repose.Headers["ContentType"] = ContentType.JsonUTF8
	repose.Headers["AccessControlAllowOrigin"] = "*"
	repose.IsBase64Encoded = false
	return repose, nil
}

func main() {
	fc.Start(HandleHttpRequest)
}
