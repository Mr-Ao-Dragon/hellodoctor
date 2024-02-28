package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Mr-Ao-Dragon/hellodoctor/database"
	"github.com/Mr-Ao-Dragon/hellodoctor/user"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

type QueryResult struct {
	Code int                           `json:"code"`
	Msg  string                        `json:"msg"`
	Data []user.SingleDoctorDataStruct `json:"data"`
}
type structEvent struct {
}
type reposeJson struct {
	Code int16       `json:"code"`
	Body QueryResult `json:"body"`
}

func HandleHttpRequest(ctx context.Context, event structEvent) (repose string, err error) {
	Query, err := database.ListDoctor()
	if err != nil {
		reposeByte, _ := json.Marshal(reposeJson{
			Code: http.StatusInternalServerError,
			Body: QueryResult{
				Code: http.StatusInternalServerError,
				Msg:  "DatabaseQueryFailed",
				Data: nil,
			},
		})
		return string(reposeByte), err
	}
	reposeBody := QueryResult{Code: 200, Msg: "success", Data: Query}
	reposeStruct := reposeJson{Code: 200, Body: reposeBody}
	reposeByte, err := json.Marshal(reposeStruct)
	return string(reposeByte), nil
}
func main() {
	fc.Start(HandleHttpRequest)
}
