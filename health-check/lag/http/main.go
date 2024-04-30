package main

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/aliyun/fc-runtime-go-sdk/fc"

	"github.com/Mr-Ao-Dragon/hellodoctor/tool/commonData/ContentType"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
)

func HandleHttpRequest(ctx context.Context, event datastruct.EventStruct) (repose *datastruct.UniversalRepose, err error) {
	repose = new(datastruct.UniversalRepose)
	repose.Init()
	repose.StatusCode = http.StatusOK
	repose.IsBase64Encoded = false
	repose.Headers["ContentType"] = ContentType.TextUTF8
	repose.Headers["AccessControlAllowOrigin"] = "*"
	repose.Body = strconv.FormatInt(time.Now().UnixMilli(), 10)
	err = nil
	return
}

func main() {
	fc.Start(HandleHttpRequest)
}
