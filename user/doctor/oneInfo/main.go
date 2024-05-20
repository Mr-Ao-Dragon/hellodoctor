package main

import (
	"context"
	"encoding/json"
	"github.com/aliyun/fc-runtime-go-sdk/events"
	"net/http"
	"runtime"

	"github.com/aliyun/fc-runtime-go-sdk/fc"

	"github.com/Mr-Ao-Dragon/hellodoctor/tool/commonData/ContentType"
	"github.com/Mr-Ao-Dragon/hellodoctor/user"
)

type Id struct {
	Id string `json:"id"`
}

func HandleHttpRequest(ctx context.Context, event events.HTTPTriggerEvent) (repose *events.HTTPTriggerResponse, err error) {
	Open := new(Id)
	repose = new(events.HTTPTriggerResponse)
	_ = json.Unmarshal([]byte(*event.Body), Open)
	jsonByte, err := user.QueryDoctor(Open.Id)
	if err != nil {
		repose.StatusCode = http.StatusNotFound
		repose.Headers["ContentType"] = ContentType.JsonUTF8
		repose.Headers["AccessControlAllowOrigin"] = "*"
		repose.IsBase64Encoded = false
		repose.Body = "not found"
		return repose, err
	}
	marshaledJson, _ := json.Marshal(jsonByte)
	repose.Headers["ContentType"] = ContentType.JsonUTF8
	repose.Headers["AccessControlAllowOrigin"] = "*"
	repose.IsBase64Encoded = false
	repose.StatusCode = http.StatusOK
	repose.Body = string(marshaledJson)
	defer runtime.GC()
	return repose, nil
}

func main() {
	fc.Start(HandleHttpRequest)
}
