package user

import "github.com/Mr-Ao-Dragon/hellodoctor/database"

type PermLevel struct {
	Baned       int8
	User        int8
	Assistant   int8
	Doctor      int8
	SystemAdmin int8
}

func (p *PermLevel) ToBaned(OpenID string) {
	p.Baned = 0
	finished, err := database.SetPermission(int(p.Baned), OpenID)
	switch {
	case finished == false:
		return
	case err != nil:
		return
	}
}

func (p *PermLevel) ToUser(OpenID string) {
	p.User = 1
	finished, err := database.SetPermission(int(p.User), OpenID)
	switch {
	case finished == false:
		return
	case err != nil:
		return
	}
}

func (p *PermLevel) ToAssistant(OpenID string) {
	p.Assistant = 2
	finished, err := database.SetPermission(int(p.Assistant), OpenID)
	switch {
	case finished == false:
		return
	case err != nil:
		return
	}
}

func (p *PermLevel) ToDoctor(OpenID string) {
	p.Doctor = 3
	finished, err := database.SetPermission(int(p.Doctor), OpenID)
	switch {
	case finished == false:
		return
	case err != nil:
		return
	}
}

func (p *PermLevel) ToSystemAdmin(OpenID string) {
	p.SystemAdmin = 4
	finished, err := database.SetPermission(int(p.SystemAdmin), OpenID)
	switch {
	case finished == false:
		return
	case err != nil:
		return
	}
}

func (p *PermLevel) Set() (level *PermLevel) {
	return p
}
