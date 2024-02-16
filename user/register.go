package user

import (
	"encoding/hex"
	"github.com/Mr-Ao-Dragon/hellodoctor/database"
	"io"
	"log"
	"os"
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
	systemAccessToken, err = GenToken(16)
	return
}
func GenToken(byteDataLength int16) (Token string, err error) {
	file, err := os.Open("/dev/random")
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	buffer := make([]byte, byteDataLength)

	// 使用io.ReadFull确保读取指定数量的字节
	_, err = io.ReadFull(file, buffer)
	if err != nil && err != io.EOF {
		return "", err
	}
	return hex.EncodeToString(buffer), nil
}
