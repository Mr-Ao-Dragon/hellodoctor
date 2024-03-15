package main

import (
	"context"
	"time"

	"github.com/aliyun/fc-runtime-go-sdk/fc"

	"github.com/Mr-Ao-Dragon/hellodoctor/health-check/lag"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
)

type structEvent struct {
	Version         string `json:"version"`
	RawPath         string `json:"rawPath"`
	Body            string `json:"body"`
	IsBase64Encoded bool   `json:"isBase64Encoded"`
	RequestContext  struct {
		AccountId    string `json:"accountId"`
		DomainName   string `json:"domainName"`
		DomainPrefix string `json:"domainPrefix"`
		Http         struct {
			Method    string `json:"method"`
			Path      string `json:"path"`
			Protocol  string `json:"protocol"`
			SourceIp  string `json:"sourceIp"`
			UserAgent string `json:"userAgent"`
		} `json:"http"`
		RequestId string    `json:"requestId"`
		Time      time.Time `json:"time"`
		TimeEpoch string    `json:"timeEpoch"`
	} `json:"requestContext"`
}

func HandleHttpRequest(ctx context.Context, event structEvent) (repose *datastruct.UniversalRepose, err error) {
	return lag.HttpLag(), nil
}

func main() {
	fc.Start(HandleHttpRequest)
}
