package main

import (
	"context"
	"net/http"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

// 主函数，用于启动一个HTTP服务器
func main() {
	fc.StartHttp(HandleHttpRequest)
}

// HandleHttpRequest 处理 HTTP 请求的函数
func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) (err error) {
	// 设置响应的状态码为200
	w.WriteHeader(http.StatusOK)
	// 设置响应头的Content-Type为text/plain
	w.Header().Add("Content-Type", "text/plain")
	// 向响应中写入"hello, world!\n"的内容
	w.Write([]byte("hello, world!\n"))
	return nil
}
