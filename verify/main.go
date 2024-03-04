package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"log"
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
	log.Printf("strSlice: %v", strSlice)
	log.Printf("Token: %s", Token)
	log.Printf("Timestamp: %s", event.QueryParameters.Timestamp)
	log.Printf("EchoStr: %s", event.QueryParameters.EchoStr)
	sort.Strings(strSlice)
	sortedStr := strings.Join(strSlice, "")
	log.Printf("sortedStr: %s", sortedStr)
	hashed := sha1.New()
	hashed.Write([]byte(sortedStr))
	log.Printf("hashed: %s", hashed.Sum(nil))
	localSignature := hex.EncodeToString(hashed.Sum(nil))
	log.Printf("localSignature: %s", localSignature)
	if localSignature == event.QueryParameters.Signature {
		log.Printf("Signature is correct!")
		log.Printf("remote Signture is: %s", event.QueryParameters.Signature)
		log.Printf("local Signature is: %s", localSignature)
		// reposeStruct := reposeJson{
		// 	Code: http.StatusOK,
		// 	Body: event.QueryParameters.EchoStr,
		// }
		// reposeByte, _ := json.Marshal(reposeStruct)
		repose = localSignature
		err = nil
		return
	} else {
		log.Printf("Signature is not correct")
		log.Printf("remote Signture is: %s", event.QueryParameters.Signature)
		log.Printf("local Signature is: %s", localSignature)
		// reposeStruct := reposeJson{
		// 	Code: http.StatusBadRequest,
		// 	Body: "",
		// }
		// reposeByte, _ := json.Marshal(reposeStruct)
		repose = ""
		err = nil
		return
	}
}
func main() {
	fc.Start(HandleHttpRequest)
}
