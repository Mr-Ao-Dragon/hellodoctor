package datastruct

import (
	"errors"
	"time"
)

type (
	UniversalRepose struct {
		StatusCode      int16   `json:"statusCode"`
		Headers         Headers `json:"headers"`
		IsBase64Encoded bool    `json:"isBase64Encoded"`
		Body            string  `json:"body"`
	}
	Headers struct {
		AccessControlAllowOrigin      string `json:"Access-Control-Allow-Origin"`
		AccessControlAllowHeaders     string `json:"Access-Control-Allow-Headers"`
		AccessControlAllowMethods     string `json:"Access-Control-Allow-Methods"`
		AccessControlMaxAge           string `json:"Access-Control-Max-Age"`
		AccessControlAllowCredentials string `json:"Access-Control-Allow-Credentials"`
		AccessControlExposeHeaders    string `json:"Access-Control-Expose-Headers"`
		AccessControlRequestHeaders   string `json:"Access-Control-Request-Headers"`
		AccessControlRequestMethod    string `json:"Access-Control-Request-Method"`
		ContentType                   string `json:"Content-Type"`
	}
	EventStruct struct {
		Version         string         `json:"version"`
		RawPath         string         `json:"rawPath"`
		Body            string         `json:"body"`
		IsBase64Encoded bool           `json:"isBase64Encoded"`
		Headers         map[string]any `json:"headers"`
		QueryParameters map[string]any `json:"queryParameters"`
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
	if source.Headers["token"] != nil {
		return source.Headers["token"].(string), nil
	}
	return "", errors.New("read failed")
}
