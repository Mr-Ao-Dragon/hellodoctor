package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/aliyun/fc-runtime-go-sdk/fc"

	"github.com/Mr-Ao-Dragon/hellodoctor/tool/commonData/ContentType"

	"github.com/Mr-Ao-Dragon/hellodoctor/database"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/gen"
	"github.com/Mr-Ao-Dragon/hellodoctor/user"
)

// ReposeBody represents the JSON structure of the response body.
type ReposeBody struct {
	Token     string `json:"token"`
	ExpiresIn int64  `json:"expires_in"`
}

// HandleHttpRequest TODO: 整理此处代码，目前版本代码可读性较差
func HandleHttpRequest(ctx context.Context, event datastruct.EventStruct) (Repose *datastruct.UniversalRepose, err error) {
	// Convert code to OpenID
	Repose = new(datastruct.UniversalRepose)
	OpenID, err := user.CodeToOpenID(event.QueryParameters["code"].(string))
	if err != nil {
		log.Printf("Query Failed By: %s", event.QueryParameters["code"])
		Repose.StatusCode = http.StatusUnauthorized
		Repose.Headers.ContentType = ContentType.JsonUTF8
		respJson := struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		}
		respJsonByte, _ := json.Marshal(respJson)
		Repose.Body = string(respJsonByte)
		return Repose, nil
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
		Repose.StatusCode = http.StatusServiceUnavailable
		Repose.Headers.ContentType = ContentType.JsonUTF8
		respJson := struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusServiceUnavailable,
			Message: "Database query failed",
		}
		respJsonByte, err := json.Marshal(respJson)
		Repose.Body = string(respJsonByte)
		return Repose, err
	}
	Perm := int8(0)
	if event.QueryParameters["init_password"] == os.Getenv("InitPass") {
		Perm = 5
	}
	// If user exists, login, otherwise register
	if Result {
		log.Printf("%s this user is logined", OpenID)
		QueryResult.Token, QueryResult.ExpiresIn, err = user.Login(AuthData)
	} else {
		log.Printf("This user is %s reated", OpenID)
		QueryResult.Token, QueryResult.ExpiresIn, err = user.Register(OpenID, Perm)
	}
	if err != nil {
		Repose.StatusCode = http.StatusBadRequest
		Repose.Headers.ContentType = ContentType.JsonUTF8
		respJson := struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusBadRequest,
			Message: "Unknown data",
		}
		respJsonByte, err := json.Marshal(respJson)
		Repose.Body = string(respJsonByte)
		return Repose, err
	}
	QueryResultJson, _ := json.Marshal(QueryResult)
	Repose.StatusCode = http.StatusOK
	Repose.Headers.ContentType = ContentType.JsonUTF8
	Repose.IsBase64Encoded = false
	Repose.Body = string(QueryResultJson)
	return
}

func main() {
	fc.Start(HandleHttpRequest)
}
