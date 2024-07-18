package main

import (
	"context"
	"encoding/json"
	"github.com/Mr-Ao-Dragon/hellodoctor/database"
	"github.com/Mr-Ao-Dragon/hellodoctor/reserve"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/commonData/ContentType"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/dataProcess"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
	"github.com/aliyun/fc-runtime-go-sdk/events"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"log"
	"net/http"
)

func HandleHttpRequest(ctx context.Context, event events.HTTPTriggerEvent) (repose *events.HTTPTriggerResponse, err error) {
	token := (*event.QueryParameters)["token"]
	OpenID, err := dataProcess.CutOpenID(token)
	repose = new(events.HTTPTriggerResponse)
	if err != nil {
		repose.StatusCode = http.StatusInternalServerError
		repose.Headers["ContentType"] = ContentType.TextUTF8
		repose.Headers["AccessControlAllowOrigin"] = "*"
		repose.IsBase64Encoded = false
		repose.Body = "剪切 OpenID 失败！"
		return repose, nil
	}
	Login := &datastruct.AuthStruct{
		SystemToken: token,
		OpenID:      OpenID,
	}

	log.Printf("数据已拷贝！")
	log.Printf("form: %#v", Login)
	log.Printf("to: %#v", repose)
	log.Printf("——————————————————————————————————————————————————")
	log.Printf("用户 %v 尝试登录……", Login.OpenID)
	LoginStatus, err := database.QueryLogin(Login)
	if !LoginStatus {
		log.Printf("用户 %v 未登录！", Login.OpenID)
		repose.StatusCode = http.StatusUnauthorized
		repose.Headers["ContentType"] = ContentType.TextUTF8
		repose.Headers["AccessControlAllowOrigin"] = "*"
		repose.IsBase64Encoded = false
		repose.Body = "Unauthorized"
		return repose, nil
	}
	log.Printf("用户 %v 成功登录！", Login.OpenID)
	queryResult, err := reserve.GetReserve(Login.OpenID)
	if err != nil {
		repose.StatusCode = http.StatusInternalServerError
		repose.Headers["ContentType"] = ContentType.TextUTF8
		repose.Headers["AccessControlAllowOrigin"] = "*"
		repose.IsBase64Encoded = false
		repose.Body = "查询失败,数据库错误"
		return repose, nil
	}
	log.Printf("——————————————————————————————————————————————————")
	log.Printf("查询结果: %#v", queryResult)
	repose.StatusCode = http.StatusOK
	repose.Headers["ContentType"] = ContentType.JsonUTF8
	repose.Headers["AccessControlAllowOrigin"] = "*"
	repose.IsBase64Encoded = false
	respBody, err := json.Marshal(queryResult)
	if err != nil {
		repose.StatusCode = http.StatusInternalServerError
		repose.Headers["ContentType"] = "application/json"
		repose.IsBase64Encoded = false
		repose.Body = "组装 json 失败"
		return repose, nil
	}
	repose.Body = string(respBody)
	return repose, nil
}

func main() {
	fc.Start(HandleHttpRequest)
}
