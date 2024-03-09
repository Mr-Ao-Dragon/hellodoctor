package lag

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
)

func HttpLag() (repose string) {
	reposeStr := new(datastruct.UniversalRepose)
	type jsonStr struct {
		ReposeTime int64 `jsonStr:"repose_time"`
	}
	bodyStr := new(jsonStr)
	bodyStr.ReposeTime = time.Now().UnixMilli()
	bodyByte, _ := json.Marshal(*bodyStr)
	reposeStr.StatusCode = http.StatusOK
	reposeStr.IsBase64Encoded = "false"
	reposeStr.Body = string(bodyByte)
	reposeByte, _ := json.Marshal(*reposeStr)
	return string(reposeByte)
}
