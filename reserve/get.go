package reserve

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Mr-Ao-Dragon/hellodoctor/database"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
)

func GetReserve(OpenID string) (reserveData *datastruct.InfoReserve, err error) {
	var OpenIDStr []string
	OpenIDStr = append(OpenIDStr, OpenID)
	reserveData = new(datastruct.InfoReserve)
	reserveData.Data, err = database.QueryReserve(OpenIDStr)
	if err != nil {
		log.Printf("查询失败,数据库错误")
		log.Printf("Error info: %#v", err)
		return nil, err
	}
	*reserveData.Code = http.StatusOK
	*reserveData.Msg = "success"
	return reserveData, nil
}
func GetReserveByID(reserveID int32) (reserveData *datastruct.SingleReserve, err error) {
	reserveData, err = database.QueryReserveSingle(reserveID)
	if err != nil {
		log.Printf("查询失败,数据库错误")
		log.Printf("Error info: %#v", err)
		return nil, err
	}
	return reserveData, nil
}
func GetReserveListWithDoctor(doctorID string) (reserveData *datastruct.InfoReserve, err error) {
	reserveData = new(datastruct.InfoReserve)
	reserveData.Data, err = database.QueryReserveListByDoctor(doctorID)
	if err != nil {
		reserveData.Data = nil
		*reserveData.Msg = fmt.Sprintf("query error: %#v", err)
		log.Print(*reserveData.Msg)
		return reserveData, err
	}
	*reserveData.Code = 200
	*reserveData.Msg = "success"
	return
}
