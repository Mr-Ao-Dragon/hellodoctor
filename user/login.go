package user

import (
	"os"

	"github.com/Mr-Ao-Dragon/hellodoctor/database"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/gen"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
)

func Login(authStruct *datastruct.AuthStruct) (systemAccessToken string, expiresIn int64, err error) {

	retry := 0
	var expressed bool
	for isFinish := false; !isFinish || err == nil || retry >= 3 || authStruct == nil || expressed == true; {
		isFinish, expiresIn, expressed, err = database.UserLogin(authStruct)
		retry++
	}
	if retry == 3 && err != nil {
		return "", 0, err
	}
	systemAccessToken, err = gen.Token(16)
	systemAccessToken = systemAccessToken + authStruct.OpenID
	return
}
func CodeToOpenID(code string) (OpenID string, err error) {
	OfficialAccountApp, err := officialAccount.NewOfficialAccount(&officialAccount.UserConfig{
		AppID:  os.Getenv("AppId"),     // 公众号、小程序的appid
		Secret: os.Getenv("AppSecret"), // 公众号、小程序的appsecret
		Token:  "0000",                 // 不需要使用但要求填写
		AESKey: "0000",                 // 不需要使用但要求填写
		Log: officialAccount.Log{
			Level: "debug",
		},
		HttpDebug: true,
		Debug:     true,
	})
	if err != nil {
		return "", err
	}
	UserInterface, err := OfficialAccountApp.OAuth.UserFromCode(code)
	OpenID = UserInterface.GetOpenID()
	err = nil
	return
}
