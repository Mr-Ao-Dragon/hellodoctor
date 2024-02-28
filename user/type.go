package user

type AuthStruct struct {
	OpenID      string `json:"id"`
	SystemToken string `json:"token"`
}

type SingleDoctorDataStruct struct {
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
	Id      string `json:"id"`
	Profile string `json:"profile"`
}
