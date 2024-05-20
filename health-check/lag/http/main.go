package main

import (
	"context"
	"github.com/aliyun/fc-runtime-go-sdk/events"
	"net/http"
	"strconv"
	"time"

	"github.com/aliyun/fc-runtime-go-sdk/fc"

	"github.com/Mr-Ao-Dragon/hellodoctor/tool/commonData/ContentType"
)

func HandleHttpRequest(ctx context.Context, event events.HTTPTriggerEvent) (repose *events.HTTPTriggerResponse, err error) {
	repose = new(events.HTTPTriggerResponse)
	repose.StatusCode = http.StatusOK
	repose.IsBase64Encoded = false
	repose.Headers = make(map[string]string)
	repose.Headers["ContentType"] = ContentType.TextUTF8
	repose.Headers["AccessControlAllowOrigin"] = "*"
	repose.Body = strconv.FormatInt(time.Now().UnixMilli(), 10)
	err = nil
	return
}

func main() {
	fc.Start(HandleHttpRequest)
}
