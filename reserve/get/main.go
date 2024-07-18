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
	"net/http"
	"strconv"
)

type respBody struct {
	Code int                      `json:"code"`
	Msg  string                   `json:"msg"`
	Data datastruct.SingleReserve `json:"data"`
}

func HandleHttpRequest(ctx context.Context, event events.HTTPTriggerEvent) (repose *events.HTTPTriggerResponse, err error) {
	token := (*event.QueryParameters)["token"]
	reserveID, _ := (strconv.Atoi((*event.QueryParameters)["id"]))
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
	reserveData, err := reserve.GetReserveByID(int32(reserveID))
	if err != nil {
		repose.StatusCode = http.StatusNotFound
		repose.Body = "Not Found"
		repose.Headers["ContentType"] = ContentType.TextUTF8
		repose.IsBase64Encoded = false
		err = nil
		return
	}
	respJson := respBody{
		Code: 200,
		Msg:  "操作成功",
		Data: *reserveData,
	}
	respBytes, _ := json.Marshal(respJson)
	repose.StatusCode = http.StatusOK
	repose.Body = string(respBytes)
	repose.Headers["ContentType"] = ContentType.JsonUTF8
	repose.Headers["Access-Control-Allow-Origin"] = "*"
	repose.IsBase64Encoded = false
	err = nil
	return repose, nil
}

func main() {
	fc.Start(HandleHttpRequest)
}
