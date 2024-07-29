package main

import (
	"context"
	"encoding/json"
	"github.com/Mr-Ao-Dragon/hellodoctor/database"
	"github.com/Mr-Ao-Dragon/hellodoctor/reserve"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/dataProcess"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
	"github.com/aliyun/fc-runtime-go-sdk/events"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"log"
	"net/http"
)

func handleRequest(ctx context.Context, event events.HTTPTriggerEvent) (repose *events.HTTPTriggerResponse, err error) {
	// 验证登录状态
	token := (*event.QueryParameters)["token"]
	OpenID, err := dataProcess.CutOpenID(token)
	auth := &datastruct.AuthStruct{
		SystemToken: token,
		OpenID:      OpenID,
	}
	if loginStatus, err := database.QueryLogin(auth); !loginStatus || err != nil {
		(*repose).StatusCode = http.StatusUnauthorized
		(*repose).Body = "Unauthorized"
		(*repose).IsBase64Encoded = false
		(*repose).Headers["Access-Control-Allow-Origin"] = "*"
		(*repose).Headers["Content-Type"] = "text/plain"
		err = nil
		return
	}
	if perm, err := database.QueryPermission(OpenID); perm < 2 || err != nil {
		(*repose).StatusCode = http.StatusForbidden
		(*repose).Body = "you do not have perm to do this action"
		(*repose).IsBase64Encoded = false
		(*repose).Headers["Access-Control-Allow-Origin"] = "*"
		(*repose).Headers["Content-Type"] = "text/plain"
		log.Printf("err: %#v", err)
		err = nil
		return
	}
	queryData, err := reserve.GetReserveListWithDoctor(OpenID)
	if err != nil {
		(*repose).StatusCode = http.StatusInternalServerError
		(*repose).Body = "internal server error"
		(*repose).IsBase64Encoded = false
		(*repose).Headers["Access-Control-Allow-Origin"] = "*"
		(*repose).Headers["Content-Type"] = "text/plain"
	}
	reposeByte, _ := json.Marshal(queryData)
	(*repose).StatusCode = http.StatusOK
	(*repose).Body = string(reposeByte)
	(*repose).IsBase64Encoded = false
	(*repose).Headers["Access-Control-Allow-Origin"] = "*"
	(*repose).Headers["Content-Type"] = "application/json"
	err = nil
	return
}
func main() {
	fc.Start(handleRequest)

}
