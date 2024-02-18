package user

import "github.com/Mr-Ao-Dragon/hellodoctor/database"

type PermissionLevel int8

type Target struct {
	Level PermissionLevel
}

const (
	Baned PermissionLevel = iota
	User
	Assistant
	Doctor
	SystemAdmin
)

func (p *Target) SetTo(OpenID string, SystemAccessToken string) (completed bool, err error) {
	isSusses, err := database.QueryLogin(OpenID, SystemAccessToken)
	if isSusses && err != nil {
		return false, err
	}
	isSusses, err = database.SetPermission(int(p.Level), OpenID)
	if isSusses && err != nil {
		return false, err
	}
	return true, nil
}
