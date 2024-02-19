package user

import (
	"github.com/Mr-Ao-Dragon/hellodoctor/database"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/gen"
)

func Register(OpenID string) (systemAccessToken string, expiresIn int64, err error) {
	retry := 0
	for isFinish := false; !isFinish || err == nil || retry >= 3 || OpenID == ""; {
		isFinish, expiresIn, err = database.AddUser(OpenID, systemAccessToken)
		retry++
	}
	if retry == 3 && err != nil {
		return "", 0, err
	}
	systemAccessToken, err = gen.Token(16)
	return
}
