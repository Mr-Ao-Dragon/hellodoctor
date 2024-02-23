package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"log"
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
	tempArray := []string{event.QueryParameters.Timestamp, event.QueryParameters.Nonce, event.QueryParameters.EchoStr}
	sort.Strings(tempArray)
	sha1String := strings.Join(tempArray, "")
	log.Printf("sha1String: %s", sha1String)
	// HashedInput := sha1.Sum([]byte(sha1String))
	sha1Hasher := sha1.New()
	sha1Hasher.Write([]byte(sha1String))
	log.Printf("sha1String: %s", sha1Hasher)
	sha1String = hex.EncodeToString(sha1Hasher.Sum([]byte("")))
	log.Printf("sha1String: %s", sha1String)
	if sha1String == event.QueryParameters.Signature {
		body := reposeJson{
			Code: 200,
			Body: sha1String,
		}
		reposeBody, err := json.Marshal(body)
		log.Printf("reposeBody: %s", reposeBody)
		if err != nil {
			return "", err
		}
		return string(reposeBody), nil
	} else {
		body := reposeJson{
			Code: 400,
			Body: "error",
		}
		reposeBody, err := json.Marshal(body)
		if err != nil {
			return "", err
		}
		return string(reposeBody), nil
	}
}
func main() {
	fc.Start(HandleHttpRequest)

}
