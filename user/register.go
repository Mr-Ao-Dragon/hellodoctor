package user

import (
	"github.com/Mr-Ao-Dragon/hellodoctor/database"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/gen"
)

func Register(OpenID string, PermissionLevel int8) (systemAccessToken string, expiresIn int64, err error) {
	retry := 0
	TokenRaw, err := gen.Token(16)
	if err != nil {
		return "", 0, err
	}
	systemAccessToken = TokenRaw + OpenID
	for isFinish := false; !isFinish || err == nil || retry >= 3 || OpenID == ""; {
		isFinish, expiresIn, err = database.AddUser(OpenID, PermissionLevel, systemAccessToken)
		retry++
	}
	if retry == 3 && err != nil {
		return "", 0, err
	}

	return
}
