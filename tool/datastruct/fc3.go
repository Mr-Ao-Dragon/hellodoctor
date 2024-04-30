package datastruct

import (
	"errors"
	"log"
	"strings"
	"time"
)

type (
	UniversalRepose struct {
		StatusCode      int16          `json:"statusCode"`
		Headers         map[string]any `json:"headers"`
		IsBase64Encoded bool           `json:"isBase64Encoded"`
		Body            string         `json:"body"`
	}
	EventStruct struct {
		Version         string            `json:"version"`
		RawPath         string            `json:"rawPath"`
		Body            string            `json:"body"`
		IsBase64Encoded bool              `json:"isBase64Encoded"`
		Headers         map[string]string `json:"headers"`
		QueryParameters map[string]string `json:"queryParameters"`
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
)

func (source EventStruct) ReadToken() (string, error) {
	prefix := "Bearer "
	log.Printf("raw token is: %s", source.QueryParameters["token"])
	strIndex := strings.Index(source.QueryParameters["token"], prefix)
	if strIndex == -1 {
		return "", errors.New(`wong format`)
	}
	return source.QueryParameters["token"][strIndex+len(prefix):], nil
}
func (repose UniversalRepose) Init() {
	repose.Headers = make(map[string]any)
}
