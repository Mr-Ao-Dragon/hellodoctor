package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/jinzhu/copier"

	"github.com/Mr-Ao-Dragon/hellodoctor/database"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/commonData/ContentType"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
)

type (
	eventStruct struct {
		QueryParameters struct {
			Token  string `json:"token" copier:"Token"`
			OpenID string `json:"id" copier:"OpenID"`
		} `json:"queryParameters"`
	}
)

func HandleHttpRequest(ctx context.Context, event eventStruct) (repose *datastruct.UniversalRepose, err error) {
	Login := new(datastruct.AuthStruct)
	repose = new(datastruct.UniversalRepose)
	err = copier.Copy(*Login, event.QueryParameters)

	var OpenID []string
	OpenID = append(OpenID, Login.OpenID)
	if err != nil {
		repose.StatusCode = http.StatusInternalServerError
		repose.Headers.ContentType = ContentType.TextUTF8
		repose.IsBase64Encoded = false
		repose.Body = "数据深拷贝失败"
		return repose, nil
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
		repose.Headers.ContentType = ContentType.TextUTF8
		repose.IsBase64Encoded = false
		repose.Body = "Unauthorized"
		return repose, nil
	}
	log.Printf("用户 %v 成功登录！", Login.OpenID)
	queryResults, err := database.QueryReserve(OpenID)
	if err != nil {
		log.Printf("查询失败,数据库错误")
		log.Printf("Error info: %#v", err)
		repose.StatusCode = http.StatusInternalServerError
		repose.Headers.ContentType = ContentType.TextUTF8
		repose.IsBase64Encoded = false
		repose.Body = "查询失败,数据库错误"
		return repose, nil
	}
	log.Printf("——————————————————————————————————————————————————")
	log.Printf("查询结果: %#v", queryResults)
	repose.StatusCode = http.StatusOK
	repose.Headers.ContentType = ContentType.JsonUTF8
	repose.IsBase64Encoded = false
	respBody, err := json.Marshal(queryResults)
	if err != nil {
		repose.StatusCode = http.StatusInternalServerError
		repose.Headers.ContentType = "application/json"
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
