package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jinzhu/copier"

	"github.com/Mr-Ao-Dragon/hellodoctor/database"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
)

type (
	eventStruct struct {
		QueryParameters struct {
			Token  string `json:"token" copier:"Token"`
			OpenID string `json:"id" copier:"OpenID"`
		} `json:"queryParameters"`
	}
	reposeJson struct {
		StatusCode int `json:"statusCode"`
		Headers    struct {
			ContentType string `json:"Content-Type"`
		} `json:"headers"`
		IsBase64Encoded string `json:"isBase64Encoded"`
		Body            string `json:"body"`
	}
)

func HandleHttpRequest(ctx context.Context, event eventStruct) (repose string, err error) {
	Login := new(datastruct.AuthStruct)
	err = copier.Copy(*Login, event.QueryParameters)
	var OpenID []string
	OpenID = append(OpenID, Login.OpenID)
	if err != nil {
		return "", err
	}
	LoginStatus, err := database.QueryLogin(Login)
	if !LoginStatus {
		reposeData := new(reposeJson)
		reposeData.StatusCode = http.StatusUnauthorized
		reposeData.Headers.ContentType = "application/json"
		reposeData.IsBase64Encoded = "false"
		reposeData.Body = "Unauthorized"
		reposeByte, _ := json.Marshal(&reposeData)
		return string(reposeByte), nil
	}
	queryResults, err := database.QueryReserve(OpenID)
	if err == nil {
		reposeData := new(reposeJson)
		reposeData.StatusCode = http.StatusInternalServerError
		reposeData.Headers.ContentType = "application/json"
		reposeData.IsBase64Encoded = "false"
		reposeData.Body = "查询失败"
		reposeByte, _ := json.Marshal(&reposeData)
		return string(reposeByte), nil
	}
	reposeData := new(reposeJson)
	reposeData.StatusCode = http.StatusOK
	reposeData.Headers.ContentType = "application/json"
	reposeData.IsBase64Encoded = "false"
	reposeByte, err := json.Marshal(&queryResults)
	if err != nil {
		reposeData.StatusCode = http.StatusInternalServerError
		reposeData.Headers.ContentType = "application/json"
		reposeData.IsBase64Encoded = "false"
		reposeData.Body = "查询失败"
		reposeByte, _ := json.Marshal(&reposeData)
		return string(reposeByte), nil
	}
	return string(reposeByte), nil
}
func main() {

}
