package main

import (
	"context"
	. "encoding/json"

	"github.com/aliyun/fc-runtime-go-sdk/fc"

	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
	"github.com/Mr-Ao-Dragon/hellodoctor/user"
)

type structEvent struct {
	Body string `json:"body"`
}
type Id struct {
	Id string `json:"id"`
}

func HandleHttpRequest(ctx context.Context, event structEvent) (repose string, err error) {
	var Open = new(Id)
	err = Unmarshal([]byte(event.Body), Open)
	if err != nil {
		return "", err
	}
	json, err := user.QueryDoctor(Open.Id)
	marshaledJson, err := Marshal(json)
	if err != nil {
		return "", err
	}
	var reposeStruct = new(datastruct.UniversalRepose)
	reposeStruct.Body = string(marshaledJson)
	reposeByte, err := Marshal(reposeStruct)
	if err != nil {
		return "", err
	}
	return string(reposeByte), nil
}
func main() {
	fc.Start(HandleHttpRequest)
}
