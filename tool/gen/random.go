package gen

import (
	"crypto/rand"
	"encoding/hex"
)

// Token 从 /dev/random 生成指定长度的随机字节，并返回其十六进制表示的令牌。
func Token(length int) (token string, err error) {

	buffer := make([]byte, length)

	// 使用 io.ReadFull 确保读取指定数量的字节
	_, err = rand.Read(buffer)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(buffer), nil
}
