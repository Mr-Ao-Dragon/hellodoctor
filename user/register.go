package user

import (
	"errors"
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
	for isFinish := false; !isFinish || err != nil || OpenID == ""; {
		isFinish, expiresIn, err = database.AddUser(OpenID, PermissionLevel, systemAccessToken)
		retry++
		switch {
		case retry > 3:
			return "", 0, errors.New("重试次数过多")
		default:
			continue
		}
	}
	return
}
