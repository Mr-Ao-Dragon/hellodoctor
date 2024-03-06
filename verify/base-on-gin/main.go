package main

import (
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		signature, _ := c.Params.Get("signature")
		timestamp, _ := c.Params.Get("timestamp")
		nonce, _ := c.Params.Get("nonce")
		echostr, _ := c.Params.Get("echostr")
		token := os.Getenv("token")
		strSlice := []string{token, timestamp, nonce}
		sort.Strings(strSlice)
		sortedStr := strings.Join(strSlice, "")
		hashed := sha1.New()
		hashed.Write([]byte(sortedStr))
		localSignature := hex.EncodeToString(hashed.Sum(nil))
		if localSignature == signature {
			c.Data(http.StatusOK, "text/html", []byte(echostr))
		} else {
			c.Data(http.StatusBadRequest, "text/html", []byte("failed"))
		}
	})
	r.Run("0.0.0.0:80") // 监听并在 0.0.0.0:8080 上启动服务
}
