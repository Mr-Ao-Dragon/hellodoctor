package datastruct

type SingleDoctorDataStruct struct {
	Name    string `json:"name"     copier:"DoctorName"`
	Avatar  string `json:"avatar"   copier:"DoctorAvatar"`
	Id      string `json:"id"       copier:"DoctorID"`
	Profile string `json:"profile"`
}
