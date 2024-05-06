package main

import (
	"context"
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/aliyun/fc-runtime-go-sdk/fc"

	"github.com/Mr-Ao-Dragon/hellodoctor/database"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/commonData/ContentType"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
)

type QueryResult struct {
	Code int                                 `json:"code"`
	Msg  string                              `json:"msg"`
	Data []datastruct.SingleDoctorDataStruct `json:"data"`
}

func HandleHttpRequest(ctx context.Context, event datastruct.EventStruct) (repose *datastruct.UniversalRepose, err error) {
	repose = new(datastruct.UniversalRepose)
	repose.Init()
	Query, err := database.ListDoctor()
	defer runtime.GC()
	if err != nil {
		respBody := QueryResult{
			Code: http.StatusInternalServerError,
			Msg:  "Query failed",
			Data: nil,
		}
		respBodyJson, err := json.Marshal(respBody)
		repose.StatusCode = http.StatusInternalServerError
		repose.Body = string(respBodyJson)
		return repose, err
	}
	respBody := QueryResult{Code: http.StatusOK, Msg: "success", Data: Query}
	respBodyJson, _ := json.Marshal(respBody)
	repose.StatusCode = http.StatusOK
	repose.Headers["ContentType"] = ContentType.JsonUTF8
	repose.Headers["AccessControlAllowOrigin"] = "*"
	repose.Body = string(respBodyJson)
	return repose, nil
}

func main() {
	fc.Start(HandleHttpRequest)
}
