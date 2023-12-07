package main

import (
	"context"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"net/http"
	"os"
)

func register(openId string, accessToken string) {
	client := tablestore.NewClient(
		os.Getenv("endPoint"),
		os.Getenv("instanceName"),
		os.Getenv("accessKeyId"),
		os.Getenv("accessKeySecret"),
		nil,
	)
	putRowRequest := new(tablestore.PutRowRequest)
	PutRowChange := new(tablestore.PutRowChange)
	PutRowChange.TableName = "user"
	// 设置主键并设置数据
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("openID", openId)
	PutRowChange.PrimaryKey = putPk
	PutRowChange.AddColumn("accessToken", accessToken)
	PutRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	putRowRequest.PutRowChange = PutRowChange
	_, err := client.PutRow(putRowRequest)
	if err != nil {
		panic(err)
	}
}
func checkRegister(openId string) bool {
	client := tablestore.NewClient(
		os.Getenv("endPoint"),
		os.Getenv("instanceName"),
		os.Getenv("accessKeyId"),
		os.Getenv("accessKeySecret"),
		nil,
	)
	request := &tablestore.SQLQueryRequest{Query: "SELECT openID FORM user WHERE" + openId}
	QueryResult, err := client.SQLQuery(request)
	if err != nil {
		panic(err)
	}
	var Result bool
	if QueryResult != nil {
		Result = true
	} else {
		Result = false
	}
	return Result
}
func handleRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	type errorResponse struct {
		Error string `json:"error"`
	}
	return nil

}
func main() {
	fc.StartHttp(handleRequest)
}
