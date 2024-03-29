package main

import (
	"context"
	. "encoding/json"
	"net/http"

	"github.com/aliyun/fc-runtime-go-sdk/fc"

	"github.com/Mr-Ao-Dragon/hellodoctor/tool/commonData/ContentType"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
	"github.com/Mr-Ao-Dragon/hellodoctor/user"
)

type Id struct {
	Id string `json:"id"`
}

func HandleHttpRequest(ctx context.Context, event datastruct.EventStruct) (repose *datastruct.UniversalRepose, err error) {
	Open := new(Id)
	repose = new(datastruct.UniversalRepose)
	_ = Unmarshal([]byte(event.Body), Open)
	json, err := user.QueryDoctor(Open.Id)
	if err != nil {
		repose.StatusCode = http.StatusNotFound
		repose.Headers.ContentType = ContentType.JsonUTF8
		repose.IsBase64Encoded = false
		repose.Body = "not found"
		return repose, err
	}
	marshaledJson, _ := Marshal(json)
	repose.Headers.ContentType = ContentType.JsonUTF8
	repose.IsBase64Encoded = false
	repose.StatusCode = http.StatusOK
	repose.Body = string(marshaledJson)
	return repose, nil
}

func main() {
	fc.Start(HandleHttpRequest)
}
