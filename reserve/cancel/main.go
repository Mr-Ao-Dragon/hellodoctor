package main

import (
	"context"
	"github.com/aliyun/fc-runtime-go-sdk/events"
	"log"
	"net/http"
	"strconv"

	"github.com/aliyun/fc-runtime-go-sdk/fc"

	"github.com/Mr-Ao-Dragon/hellodoctor/database"
	"github.com/Mr-Ao-Dragon/hellodoctor/reserve"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/commonData/ContentType"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/dataProcess"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
)

/*
HandleHttpRequest 处理HTTP请求，根据请求参数执行相应的逻辑。
*/
func HandleHttpRequest(ctx context.Context, event events.HTTPTriggerEvent) (response *events.HTTPTriggerResponse, err error) {
	response = new(events.HTTPTriggerResponse) // 初始化响应结构体
	auth := new(datastruct.AuthStruct)         // 初始化认证信息结构体

	// 从请求参数中获取token，并检查其类型

	auth.SystemToken = (*event.QueryParameters)["token"]  // 保存token到认证信息中
	auth.OpenID, err = dataProcess.CutOpenID(auth.OpenID) // 处理OpenID
	if err != nil {
		// 如果处理OpenID失败，返回内部服务器错误
		response.StatusCode = http.StatusInternalServerError
		response.IsBase64Encoded = false
		response.Body = "Cut OpenID fail!"
		response.Headers["ContentType"] = ContentType.TextUTF8
		log.Printf("Process data failed: %#v", auth)
		err = nil
		return
	}

	// 根据认证信息查询登录状态
	loginStat, err := database.QueryLogin(auth)
	if !loginStat {
		// 如果用户未登录，返回未授权响应
		response.StatusCode = http.StatusUnauthorized
		response.IsBase64Encoded = false
		response.Body = "未登录，拒绝访问"
		response.Headers["ContentType"] = ContentType.TextUTF8
		log.Printf("No login: %#v", response)
		err = nil
		return
	}

	// 查询用户权限级别
	permLevel, err := database.QueryPermission(auth.OpenID)
	reserveID := (*event.QueryParameters)["id"]
	atoi, _ := strconv.Atoi(reserveID)
	// 根据权限级别执行相应操作
	switch {
	case permLevel <= 1:
		// 权限等级为1及以下：查询预约信息，并检查是否有权限取消预约

		queryResult, err := reserve.GetReserveByID(int32(atoi))
		if queryResult.OpenID != auth.OpenID {
			// 如果OpenID不匹配，返回未授权响应
			response.StatusCode = http.StatusForbidden
			response.Body = "权限不足，拒绝访问"
			response.Headers["ContentType"] = ContentType.TextUTF8
			response.Headers["AccessControlAllowOrigin"] = "*"
			log.Printf("No permission: %#v", response)
			err = nil
			return
		}
		// 执行成功逻辑，包括设置响应状态和正文，以及尝试取消预约
		response.StatusCode = http.StatusOK
		response.IsBase64Encoded = false
		response.Body = "success"
		response.Headers["ContentType"] = ContentType.TextUTF8
		response.Headers["AccessControlAllowOrigin"] = "*"
		log.Printf("success: %#v", response)
		err = reserve.CancelReserve(int32(atoi))
		if err != nil {
			// 如果取消预约失败，返回内部服务器错误
			response.StatusCode = http.StatusInternalServerError
			response.Body = "cancel reserve fail"
			response.Headers["ContentType"] = ContentType.TextUTF8
			log.Printf("cancel reserve fail: %#v", err)
			err = nil
			return
		}
		return
		//goland:noinspection GoDfaConstantCondition
	default:
		// 权限等级为2及以上：直接尝试取消预约
		err = reserve.CancelReserve(int32(atoi))
	}
	return
}

// main 函数作为程序入口，启动HTTP请求处理逻辑。
func main() {
	fc.Start(HandleHttpRequest)
}
