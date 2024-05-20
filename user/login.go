package user

import (
	"errors"
	"log"
	"os"

	"github.com/Mr-Ao-Dragon/hellodoctor/database"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/gen"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
)

func Login(authStruct *datastruct.AuthStruct) (systemAccessToken string, expiresIn int64, err error) {
	retry := 0
	var expressed bool
	isFinish := false
	for !isFinish || err == nil || retry <= 3 || authStruct == nil || expressed == true {
		isFinish, expiresIn, expressed, err = database.UserLogin(authStruct)
		retry++
		if retry == 4 {
			return "", 0, errors.New("重试次数大于 " + string(rune(retry)))
		}
	}
	systemAccessToken, err = gen.Token(16)
	systemAccessToken = systemAccessToken + authStruct.OpenID
	return
}

func CodeToOpenID(code string) (OpenID string, err error) {
	OfficialAccountApp, err := officialAccount.NewOfficialAccount(&officialAccount.UserConfig{
		AppID:  os.Getenv("AppId"),          // 公众号、小程序的appid
		Secret: os.Getenv("AppSecret"),      // 公众号、小程序的appsecret
		Token:  os.Getenv("token"),          // 不需要使用但要求填写
		AESKey: os.Getenv("EncodingAESKey"), // 不需要使用但要求填写
		Log: officialAccount.Log{
			Level: "debug",
		},
		HttpDebug: false,
		Debug:     false,
	})
	if err != nil {
		log.Printf("err msg: %#v", err)
		return "", err
	}
	UserInterface, err := OfficialAccountApp.OAuth.UserFromCode(code)
	if err != nil {
		log.Fatalf("err: %#v", err)
		return "", err
	}
	if UserInterface == nil {
		err = errors.New("userInterface is nil")
		return "", err
	}
	OpenID = UserInterface.GetOpenID()
	err = nil
	return
}
