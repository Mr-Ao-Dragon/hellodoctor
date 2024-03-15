package lag

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
)

func HttpLag() (repose *datastruct.UniversalRepose) {
	repose = new(datastruct.UniversalRepose)
	repose.StatusCode = http.StatusOK
	repose.IsBase64Encoded = false
	repose.Body = strconv.FormatInt(time.Now().UnixMilli(), 10)
	return
}
