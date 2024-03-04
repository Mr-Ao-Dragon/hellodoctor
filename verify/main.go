package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

type queryParameters struct {
	Signature string `json:"signature"`
	Timestamp string `json:"timestamp"`
	Nonce     string `json:"nonce"`
	EchoStr   string `json:"echostr"`
}
type reposeJson struct {
	Code int16  `json:"code"`
	Body string `json:"body"`
}
type StructEvent struct {
	QueryParameters queryParameters `json:"queryParameters"`
}

func HandleHttpRequest(ctx context.Context, event StructEvent) (repose string, err error) {
	Token := os.Getenv("Token")
	strSlice := []string{event.QueryParameters.Timestamp, event.QueryParameters.Nonce, Token}
	sort.Strings(strSlice)
	sortedStr := strings.Join(strSlice, "")
	hashed := sha1.New()
	hashed.Write([]byte(sortedStr))
	localSignature := hex.EncodeToString(hashed.Sum(nil))
	if localSignature == event.QueryParameters.Signature {
		reposeStruct := reposeJson{
			Code: http.StatusOK,
			Body: event.QueryParameters.EchoStr,
		}
		reposeByte, _ := json.Marshal(reposeStruct)
		repose = string(reposeByte)
		err = nil
		return
	} else {
		reposeStruct := reposeJson{
			Code: http.StatusBadRequest,
			Body: "",
		}
		reposeByte, _ := json.Marshal(reposeStruct)
		repose = string(reposeByte)
		err = nil
		return
	}
}
func main() {
	fc.Start(HandleHttpRequest)
}
