package user

import (
	"encoding/hex"
	"io"
	"os"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
	database "github.com/Mr-Ao-Dragon/hellodoctor/database/register"
)

func Register(code string) (systemAccessToken string, expiresIn int64, err error) {
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
	for isFinish := false; !isFinish || err == nil || retry >= 3 || OpenID == ""; {
		isFinish, expiresIn, err = database.AddUser(OpenID, systemAccessToken)
		retry++
	}
	if retry == 3 && err != nil {
		return "", 0, err
	}
	systemAccessToken, err = GenToken(16)
	return
}
func GenToken(byteDataLength int16) (Token string, err error) {
	file, err := os.Open("/dev/random")
	if err != nil {
		return "", err
	}
	defer file.Close()

	buffer := make([]byte, byteDataLength)

	// 使用io.ReadFull确保读取指定数量的字节
	_, err = io.ReadFull(file, buffer)
	if err != nil && err != io.EOF {
		return "", err
	}
	return hex.EncodeToString(buffer), nil
}
