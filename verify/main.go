package main

// 导入依赖包
import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"github.com/aliyun/fc-runtime-go-sdk/events"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"

	"github.com/Mr-Ao-Dragon/hellodoctor/tool/commonData/ContentType"
)

// HandleHttpRequest 处理HTTP请求的函数
// ctx: 上下文信息，用于传递请求生命周期的值和取消信号
// event: 包含HTTP请求信息的结构体
// 返回值：repose为返回的HTTP响应信息，err为错误信息
func HandleHttpRequest(ctx context.Context, event events.HTTPTriggerEvent) (repose *events.HTTPTriggerResponse, err error) {
	repose = new(events.HTTPTriggerResponse)
	// 从上下文中获取FC（函数计算）的环境信息
	fcContext, _ := fccontext.FromContext(ctx)
	log.Printf("fcContext: %v", fcContext)

	// 从环境变量中获取Token
	Token := os.Getenv("Token")

	// 组合验证字符串
	strSlice := []string{(*event.QueryParameters)["timestamp"], (*event.QueryParameters)["nonce"], Token}
	log.Printf("Timestamp: %s", (*event.QueryParameters)["timestamp"])
	log.Printf("EchoStr: %s", (*event.QueryParameters)["echostr"])
	sort.Strings(strSlice)
	sortedStr := strings.Join(strSlice, "")

	// 计算签名
	hashed := sha1.New()
	hashed.Write([]byte(sortedStr))
	log.Printf("hashed: %s", hashed.Sum(nil))
	localSignature := hex.EncodeToString(hashed.Sum(nil))
	log.Printf("localSignature: %s", localSignature)

	// 验证签名
	if localSignature == (*event.QueryParameters)["signature"] {
		log.Printf("Signature is correct!")
		log.Printf("remote Signture is: %s", (*event.QueryParameters)["signature"])
		log.Printf("local Signature is: %s", localSignature)

		// 验证通过，返回成功响应
		repose.StatusCode = http.StatusOK
		repose.Body = (*event.QueryParameters)["echostr"]
		repose.Headers["ContentType"] = ContentType.Text
		repose.IsBase64Encoded = false
		err = nil
		return
	} else {
		// 签名不匹配，返回错误响应
		log.Printf("Signature is not correct")
		log.Printf("remote Signture is: %s", (*event.QueryParameters)["signature"])
		log.Printf("local Signature is: %s", localSignature)
		repose.StatusCode = http.StatusBadRequest
		repose.Body = ""
		repose.Headers["ContentType"] = ContentType.TextUTF8
		err = nil
		return
	}
}

// main 函数，启动函数计算服务
func main() {
	fc.Start(HandleHttpRequest)
}
