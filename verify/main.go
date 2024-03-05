package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

type queryParameters struct {
	Signature string `json:"signature"`
	Timestamp string `json:"timestamp"`
	Nonce     string `json:"nonce"`
	EchoStr   string `json:"echostr"`
}
type StructEvent struct {
	Version         string          `json:"version"`
	RawPath         string          `json:"rawPath"`
	HttpMethod      string          `json:"httpMethod"`
	QueryParameters queryParameters `json:"queryParameters"`
	Body            string          `json:"body"`
}
type ReposeBody struct {
	StatusCode int16   `json:"statusCode"`
	Body       int     `json:"body"`
	Headers    headers `json:"headers"`
}
type headers struct {
	ContentType string `json:"Content-Type"`
}

func HandleHttpRequest(ctx context.Context, event StructEvent) (repose string, err error) {
	fcContext, _ := fccontext.FromContext(ctx)
	log.Printf("fcContext: %v", fcContext)
	log.Printf("event: %v", event)
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
		reposeStr := new(ReposeBody)
		reposeStr.StatusCode = http.StatusCreated
		reposeStr.Body, _ = strconv.Atoi(event.QueryParameters.EchoStr)
		reposeStr.Headers.ContentType = "text/html; charset=utf-8"
		reposeByte, _ := json.Marshal(*reposeStr)
		repose = string(reposeByte)
		err = nil
		log.Printf("repose: %s", repose)
		return
	} else {
		log.Printf("Signature is not correct")
		log.Printf("remote Signture is: %s", event.QueryParameters.Signature)
		log.Printf("local Signature is: %s", localSignature)
		reposeStr := new(ReposeBody)
		reposeStr.StatusCode = http.StatusBadRequest
		reposeStr.Body = 0
		reposeStr.Headers.ContentType = "text/html; charset=utf-8"
		reposeByte, _ := json.Marshal(*reposeStr)
		repose = string(reposeByte)
		err = nil
		return
	}
}
func main() {
	fc.Start(HandleHttpRequest)
}
