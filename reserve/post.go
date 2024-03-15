package reserve

import (
	"github.com/jinzhu/copier"

	"github.com/Mr-Ao-Dragon/hellodoctor/database"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
)

func PostReserve(reserve *datastruct.AddReserve, authStruct *datastruct.AuthStruct) (NewReserve *datastruct.SingleReserve, err error) {
	reserveID, err := database.AddReserve(reserve, authStruct)
	if err != nil {
		return nil, err
	}
	DoctorInfo, err := database.QueryDoctor(reserve.DoctorID)
	err = copier.Copy(reserve, &NewReserve)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&DoctorInfo, &NewReserve)
	if err != nil {
		return nil, err
	}
	NewReserve.ID = reserveID
	NewReserve.Status = 1
	return NewReserve, nil
}
