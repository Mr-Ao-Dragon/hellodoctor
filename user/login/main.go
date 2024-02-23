package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/Mr-Ao-Dragon/hellodoctor/database"
	"github.com/Mr-Ao-Dragon/hellodoctor/user"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

// StructEvent represents the JSON structure of the event received by the function.
type StructEvent struct {
	QueryParameters struct {
		Code         string `json:"code"`
		InitPassword string `json:"init_password"`
	} `json:"query_parameters"`
}

// ReposeBody represents the JSON structure of the response body.
type ReposeBody struct {
	Token string `json:"token"`

	ExpiresIn int64 `json:"expires_in"`
}

// ReposeStruct represents the JSON structure of the response.
type ReposeStruct struct {
	StatusCode int16      `json:"statusCode"`
	Body       ReposeBody `json:"body"`
}

// HandleHttpRequest TODO: 整理此处代码，目前版本代码可读性较差
func HandleHttpRequest(ctx context.Context, event StructEvent) (Repose string, err error) {
	// Convert code to OpenID
	OpenID, err := user.CodeToOpenID(event.QueryParameters.Code)
	if err != nil {
		return "", err
	}

	// Query user existence in the database
	var QueryResult ReposeBody
	Result, err := database.QueryUserExist(OpenID)
	if err != nil {
		return "", err
	}
	Perm := int8(0)
	if event.QueryParameters.InitPassword == os.Getenv("InitPass") {
		Perm = 5
	}
	// If user exists, login, otherwise register
	if Result {
		QueryResult.Token, QueryResult.ExpiresIn, err = user.Login(OpenID)
	} else {
		QueryResult.Token, QueryResult.ExpiresIn, err = user.Register(OpenID, Perm)
	}
	if err != nil {
		return "{\"statusCode\": 401,\"body\": \"{\"error\": \"User does not exist\"}\"}", err
	}

	// Create response structure
	reposeStruct := ReposeStruct{
		StatusCode: 200,
		Body:       QueryResult,
	}

	// Marshal response structure to JSON
	ReposeByte, err := json.Marshal(reposeStruct)
	if err != nil {
		return "", err
	}

	// Convert JSON to string
	Repose = string(ReposeByte)
	return
}

func main() {
	fc.Start(HandleHttpRequest)
}
