package reserve

import (
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
