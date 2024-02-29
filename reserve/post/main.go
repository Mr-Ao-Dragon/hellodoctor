package main

import (
	"context"
	"encoding/json"

	"github.com/jinzhu/copier"

	"github.com/Mr-Ao-Dragon/hellodoctor/database"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
)

func main() {

}

type (
	ReqBody struct {
		DoctorID string `json:"doctor_id"   copier:"DoctorID"` // 医生ID
		OpenID   string `json:"openid"      copier:"OpenID"`   // 用户ID
		Mobile   int64  `json:"mobile"      copier:"Mobile"`   // 手机号
		Name     string `json:"name"        copier:"Name"`     // 预约者名字
		Time     int64  `json:"time"        copier:"Time"`     // 预约unix时间戳
		Token    string `json:"token"`
	}
	structEvent struct {
		QueryParameters ReqBody `json:"queryParameters"`
		Body            string  `json:"body"`
	}
	reposeJson struct {
		Code int16  `json:"code"`
		Body string `json:"body"`
	}
)

func handleRequest(ctx context.Context, event structEvent) (repose string, err error) {
	auth := &datastruct.AuthStruct{
		SystemToken: event.QueryParameters.Token,
		OpenID:      event.QueryParameters.OpenID,
	}
	loginStatus, err := database.QueryLogin(auth)
	if loginStatus || err != nil {
		reposeByte, _ := json.Marshal(reposeJson{Code: 401, Body: "Unauthorized"})
		return string(reposeByte), nil
	}
	reserveData := new(datastruct.AddReserve)
	err = copier.Copy(&reserveData, &event.QueryParameters)
	if err != nil {
		reposeByte, _ := json.Marshal(reposeJson{Code: 503, Body: "Copy data failed"})
		return string(reposeByte), nil
	}
	reserveID, err := database.AddReserve(reserveData, auth)
	newReserve := new(datastruct.SingleReserve)
	err = copier.Copy(&newReserve, &reserveData)
	newReserve.ID = reserveID
	newReserve.Status = 1
	ReserveByte, _ := json.Marshal(newReserve)
	reposeByte, _ := json.Marshal(reposeJson{Code: 200, Body: string(ReserveByte)})
	return string(reposeByte), nil
}
