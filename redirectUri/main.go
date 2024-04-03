package main

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/aliyun/fc-runtime-go-sdk/fc"

	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
)

func handleHttpRequest(ctx context.Context, event datastruct.EventStruct) (repose *datastruct.UniversalRepose, err error) {
	repose = new(datastruct.UniversalRepose)
	repose.Body = os.Getenv("RUVerify")
	if repose.Body == "" {
		repose.StatusCode = http.StatusNotFound
		repose.Headers.ContentType = "text/plain"
		repose.IsBase64Encoded = false
		err = errors.New("未正确配置环境变量: RUVerify")
		return repose, err
	}
	repose.StatusCode = http.StatusOK
	repose.Headers.ContentType = "text/plain"
	repose.IsBase64Encoded = false
	return repose, err
}
func main() {
	fc.Start(handleHttpRequest)
}