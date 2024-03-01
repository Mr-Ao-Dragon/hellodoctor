package reserve

import (
	"github.com/Mr-Ao-Dragon/hellodoctor/database"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
)

func GetReserve(OpenID string) (reserveData *datastruct.InfoReserve, err error) {
	var OpenIDStr []string
	OpenIDStr = append(OpenIDStr, OpenID)
	reserveData.Data, err = database.QueryReserve(OpenIDStr)
	if err != nil {
		return nil, err
	}
	*reserveData.Code = 200
	*reserveData.Msg = "success"
	return reserveData, nil
}
