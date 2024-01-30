
package main

import (
	"encoding/json"
	"github.com/Mr-Ao-Dragon/hellodoctor/mrao_wechat"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"net/http"
	"os"
)

type AK struct {
	AK        string `json:"access_token"`
	ExpiresIn int16  `json:"expires_in"`
}
type errorContext struct {
	Err  string `json:"Err"`
	Code int8   `json:"Code"`
}

func HandleHttpRequest(w http.ResponseWriter) error {
	AKReturn, EIReturn, err := mrao_wechat.GetAccessToken(os.Getenv("AppId"), os.Getenv("AppSecret"), false)
	AK := AK{
		AK:        AKReturn,
		ExpiresIn: EIReturn,
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-Type", "application/json")
		reposeBody, err := json.Marshal(errorContext{
			Err:  err.Error(),
			Code: 001,
		})
		_, err = w.Write(reposeBody)
		if err != nil {
			return err
		}
		return err
	}
	reposeBody, err := json.Marshal(AK)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-Type", "application/json")
		reposeBody, err := json.Marshal(errorContext{
			Err:  err.Error(),
			Code: 001,
		})
		_, err = w.Write(reposeBody)
		if err != nil {
			return err
		}
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(reposeBody)
	if err != nil {
		return err
	}
	return nil
}
func main() {
	fc.StartHttp(HandleHttpRequest)
}
