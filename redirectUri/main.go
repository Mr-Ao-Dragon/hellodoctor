package main

import (
	"context"
	"errors"
	"github.com/aliyun/fc-runtime-go-sdk/events"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"net/http"
	"os"
)

func handleHttpRequest(ctx context.Context, event events.HTTPTriggerEvent) (repose *events.HTTPTriggerResponse, err error) {
	repose = new(events.HTTPTriggerResponse)
	repose.Body = os.Getenv("RUVerify")
	repose.Headers = make(map[string]string)
	if repose.Body == "" {
		repose.StatusCode = http.StatusNotFound
		repose.Headers["ContentType"] = "text/plain"
		repose.IsBase64Encoded = false
		err = errors.New("未正确配置环境变量: RUVerify")
		return repose, err
	}
	repose.StatusCode = http.StatusOK
	repose.Headers["AccessControlAllowOrigin"] = "text/plain"
	repose.IsBase64Encoded = false
	return repose, err
}
func main() {
	fc.Start(handleHttpRequest)
}
