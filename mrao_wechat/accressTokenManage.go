package mrao_wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type AppRequest struct {
	AppId        string `json:"appid"`         // 应用ID
	AppSecret    string `json:"secret"`        // 应用密钥
	GrantType    string `json:"grant_type"`    // 授权类型
	ForceRefresh bool   `json:"force_refresh"` // 强制刷新
}
type AccessTokenResponse struct {
	AccessToken string `json:"AccessToken"` // 授权访问_token
	ExpiresIn   int16  `json:"ExpiresIn"`   // 过期时间
}

func GetAccessToken(AppId, AppSecret string, ForceRefresh bool) (AccessToken string, expiresIn int16, err error) {
	// 微信OAuth2接口地址
	url := "https://api.weixin.qq.com/cgi-bin/stable_token"
	// 创建IDRequest结构体对象，并设置参数
	Request := AppRequest{
		AppId:        "\"" + AppId + "\"",
		AppSecret:    "\"" + AppSecret + "\"",
		GrantType:    "authorization_code", // 授权类型
		ForceRefresh: ForceRefresh,
	}
	log.Printf("Request:%v", Request)
	// 将Request对象转换为JSON格式的请求体
	RequestBody, err := json.Marshal(Request)
	if err != nil {
		return "", 0, err
	}
	// 创建HTTP客户端对象
	Client := &http.Client{}
	// 创建HTTP GET请求对象，并设置请求URL和请求体
	req, err := http.NewRequest(http.MethodGet, url, bytes.NewReader(RequestBody))
	if err != nil {
		return "", 0, err
	}
	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	// 发送HTTP请求并获取响应
	// const maxRetries = 5
	// var resp *http.Response
	// for retry := 0; retry <= maxRetries; retry++ {
	//	resp, err = Client.Do(req)
	//	if err == nil {
	//		break
	//	} else {
	//		return "", 0, err
	//	}
	// }
	resp, err := Client.Do(req)
	if err != nil {
		return "", 0, err
	}
	// 延迟关闭响应体
	defer func(Body io.ReadCloser) (err error) {
		err = Body.Close()
		if err != nil {
			return err
		}
		return nil
	}(resp.Body)
	// 判断响应状态码是否为200，若不是则抛出异常
	if resp.StatusCode != 200 {
		return "", 0, fmt.Errorf(resp.Status)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}
	var tokenResp AccessTokenResponse
	err = json.Unmarshal(bodyBytes, &tokenResp)
	if err != nil {
		return "", 0, err
	}
	return tokenResp.AccessToken, tokenResp.ExpiresIn, nil
}

func ForceRefresh(AppID string, AppSecret string) (AccessToken string, expiresIn int16, err error) {
	const refreshed = 7200
	if expiresIn <= 0 || expiresIn >= refreshed {
		expiresIn = 0
	}
	for expiresIn < refreshed || err != nil {
		AccessToken, expiresIn, err = GetAccessToken(AppID, AppSecret, true)
	}
	if err != nil {
		return "", 0, err
	}
	return AccessToken, expiresIn, nil
}
