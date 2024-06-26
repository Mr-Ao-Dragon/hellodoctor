package main

import (
	"context"
	"encoding/json"
	"github.com/aliyun/fc-runtime-go-sdk/events"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"log"
	"net/http"
	"os"

	"github.com/Mr-Ao-Dragon/hellodoctor/tool/commonData/ContentType"

	"github.com/Mr-Ao-Dragon/hellodoctor/database"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/gen"
	"github.com/Mr-Ao-Dragon/hellodoctor/user"
)

// ReposeBody represents the JSON structure of the response body.
type ReposeBody struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

// HandleHttpRequest TODO: 整理此处代码，目前版本代码可读性较差
func HandleHttpRequest(ctx context.Context, event events.HTTPTriggerEvent) (repose events.HTTPTriggerResponse, err error) {
	// Convert code to OpenID
	code := (*event.QueryParameters)["code"]

	OpenID, err := user.CodeToOpenID(code)
	if err != nil {
		log.Printf("Query Failed By: %s", code)
		repose.StatusCode = http.StatusUnauthorized
		repose.Headers["ContentType"] = ContentType.JsonUTF8
		repose.Headers["AccessControlAllowOrigin"] = "*"
		respJson := struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		}
		respJsonByte, _ := json.Marshal(respJson)
		repose.Body = string(respJsonByte)
		return repose, nil
	}
	TokenRaw, err := gen.Token(16)
	Token := TokenRaw + "-" + OpenID
	AuthData := &datastruct.AuthStruct{
		OpenID:      OpenID,
		SystemToken: Token,
	}
	// Query user existence in the database
	QueryResult := new(ReposeBody)
	Result, err := database.QueryUserExist(OpenID)
	if err != nil {
		log.Printf("Database down!")
		log.Printf("error message: %+v", err)
		repose.StatusCode = http.StatusServiceUnavailable
		repose.Headers["ContentType"] = ContentType.JsonUTF8
		repose.Headers["AccessControlAllowOrigin"] = "*"
		respJson := struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusServiceUnavailable,
			Message: "Database query failed",
		}
		respJsonByte, err := json.Marshal(respJson)
		repose.Body = string(respJsonByte)
		return repose, err
	}
	Perm := int8(0)
	if (*event.QueryParameters)["init_password"] == os.Getenv("InitPass") {
		Perm = 5
	}
	// If a user exists, login, otherwise register
	if Result {
		log.Printf("this user %s is logined", OpenID)
		QueryResult.Token, QueryResult.ExpiresIn, err = user.Login(AuthData)
	} else {
		log.Printf("This user %s is created", OpenID)
		QueryResult.Token, QueryResult.ExpiresIn, err = user.Register(OpenID, Perm)
	}
	log.Printf("the error is %+v", err)
	if err != nil {
		repose.StatusCode = http.StatusBadRequest
		repose.Headers = make(map[string]string)
		repose.Headers["ContentType"] = ContentType.JsonUTF8
		respJson := struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusBadRequest,
			Message: "Unknown data",
		}
		respJsonByte, err := json.Marshal(respJson)
		repose.Body = string(respJsonByte)
		return repose, err
	}
	QueryResultJson, _ := json.Marshal(QueryResult)
	repose.StatusCode = http.StatusOK
	repose.Headers = make(map[string]string)
	repose.Headers["ContentType"] = ContentType.JsonUTF8
	repose.Headers["AccessControlAllowOrigin"] = "*"
	repose.IsBase64Encoded = false
	log.Printf("body is: %s", string(QueryResultJson))
	repose.Body = string(QueryResultJson)
	return
}

func main() {
	fc.Start(HandleHttpRequest)
}
