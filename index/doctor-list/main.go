package main

import (
	"context"
	"encoding/json"
	"github.com/aliyun/fc-runtime-go-sdk/events"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"log"
	"net/http"

	"github.com/Mr-Ao-Dragon/hellodoctor/database"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/commonData/ContentType"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
)

type QueryResult struct {
	Code int                                 `json:"code"`
	Msg  string                              `json:"msg"`
	Data []datastruct.SingleDoctorDataStruct `json:"data"`
}

func HandleHttpRequest(ctx context.Context, event events.HTTPTriggerEvent) (repose *events.HTTPTriggerResponse, err error) {
	repose = new(events.HTTPTriggerResponse)
	Query, err := database.ListDoctor()
	if err != nil {
		log.Printf("error is: %+v", err)
		respBody := QueryResult{
			Code: http.StatusInternalServerError,
			Msg:  "Query failed",
			Data: nil,
		}
		respBodyJson, _ := json.Marshal(respBody)
		repose.StatusCode = http.StatusInternalServerError
		repose.Body = string(respBodyJson)
		return repose, err
	}
	respBody := QueryResult{Code: http.StatusOK, Msg: "success", Data: Query}
	respBodyJson, _ := json.Marshal(respBody)
	repose.StatusCode = http.StatusOK
	repose.Headers = make(map[string]string)
	repose.Headers["ContentType"] = ContentType.JsonUTF8
	repose.Headers["AccessControlAllowOrigin"] = "*"
	repose.Body = string(respBodyJson)
	return repose, nil
}

func main() {
	fc.Start(HandleHttpRequest)
}
