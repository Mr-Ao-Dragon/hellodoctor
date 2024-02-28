package user

import (
	"encoding/json"

	"github.com/Mr-Ao-Dragon/hellodoctor/database"
)

func QueryDoctor(OpenID string) (jsonByte []byte, err error) {
	queryResult, err := database.QueryDoctor(OpenID)
	if err != nil {
		return nil, err
	}
	jsonByte, err = json.Marshal(queryResult)
	if err != nil {
		return nil, err
	}
	return jsonByte, nil
}

func UpToDoctor(AuthData *AuthStruct, dataStruct *SingleDoctorDataStruct) (err error) {
	_, _, err = Login(AuthData)
	if err != nil {
		return err
	}
	err = database.UpToDoctor(dataStruct)
	if err != nil {
		return err
	}
	return nil
}
