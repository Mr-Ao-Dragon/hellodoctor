package main

import (
	"encoding/json"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"golang.org/x/net/context"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func handleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	// 设置响应的状态码为200
	w.WriteHeader(http.StatusOK)
	// 设置响应头的Content-Type为text/json
	w.Header().Add("Content-Type", "application/json")
	// 组织回复内容
	var response = Response{
		Message: "ok",
	} // 向响应中写入心跳响应
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(response.Message))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	fc.StartHttp(handleHttpRequest)
}
