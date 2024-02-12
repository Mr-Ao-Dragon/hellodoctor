package user

import (
	"os"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
	database "github.com/Mr-Ao-Dragon/hellodoctor/database/register"
	user "github.com/Mr-Ao-Dragon/hellodoctor/user/action/register"
)

func Login(code string) (systemAccessToken string, expiresIn int64, err error) {
	OfficialAccountApp, err := officialAccount.NewOfficialAccount(&officialAccount.UserConfig{
		AppID:  os.Getenv("AppId"), // 公众号、小程序的appid
		Secret: os.Getenv("AppSecret"),
		Token:  os.Getenv("Token"),
		AESKey: os.Getenv("AESKey"),
		Log: officialAccount.Log{
			Level: "debug",
		},
		HttpDebug: true,
		Debug:     true,
	})
	UserInterface, err := OfficialAccountApp.OAuth.UserFromCode(code)
	OpenID := UserInterface.GetOpenID()
	retry := 0
	var expressed bool
	for isFinish := false; !isFinish || err == nil || retry >= 3 || OpenID == "" || expressed == true; {
		isFinish, expiresIn, expressed, err = database.UserLogin(OpenID, systemAccessToken)
		retry++
	}
	if retry == 3 && err != nil {
		return "", 0, err
	}
	systemAccessToken, err = user.GenToken(16)
	return
}