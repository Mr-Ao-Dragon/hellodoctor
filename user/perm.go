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

// Baned INFO: 护士、规培和前台统一归类为 Assistant
func (p *Target) Baned() *Target {
	p.Level = Baned
	return p
}
func (p *Target) User() *Target {
	p.Level = User
	return p
}
func (p *Target) Assistant() *Target {
	p.Level = Assistant
	return p
}
func (p *Target) Doctor() *Target {
	p.Level = Doctor
	return p
}
func (p *Target) SystemAdmin() *Target {
	p.Level = SystemAdmin
	return p
}

// SetTo WARN:这里不需要符合开闭原则，因为该操作单次执行一次，所以不需要继承
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
