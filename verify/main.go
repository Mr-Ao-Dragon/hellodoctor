package main

// 导入依赖包
import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"

	"github.com/Mr-Ao-Dragon/hellodoctor/tool/commonData/ContentType"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
)

// queryParameters 定义请求参数结构体
type queryParameters struct {
	Signature string `json:"signature"`
	Timestamp string `json:"timestamp"`
	Nonce     string `json:"nonce"`
	EchoStr   string `json:"echostr"`
}

// StructEvent 定义事件结构体，用于解析接收到的HTTP请求信息
type StructEvent struct {
	Version         string          `json:"version"`
	RawPath         string          `json:"rawPath"`
	HttpMethod      string          `json:"httpMethod"`
	QueryParameters queryParameters `json:"queryParameters"`
	Body            string          `json:"body"`
}

// HandleHttpRequest 处理HTTP请求的函数
// ctx: 上下文信息，用于传递请求生命周期的值和取消信号
// event: 包含HTTP请求信息的结构体
// 返回值：repose为返回的HTTP响应信息，err为错误信息
func HandleHttpRequest(ctx context.Context, event StructEvent) (repose *datastruct.UniversalRepose, err error) {
	repose = new(datastruct.UniversalRepose)
	// 从上下文中获取FC（函数计算）的环境信息
	fcContext, _ := fccontext.FromContext(ctx)
	log.Printf("fcContext: %v", fcContext)

	// 从环境变量中获取Token
	Token := os.Getenv("Token")

	// 组合验证字符串
	strSlice := []string{event.QueryParameters.Timestamp, event.QueryParameters.Nonce, Token}
	log.Printf("Timestamp: %s", event.QueryParameters.Timestamp)
	log.Printf("EchoStr: %s", event.QueryParameters.EchoStr)
	sort.Strings(strSlice)
	sortedStr := strings.Join(strSlice, "")

	// 计算签名
	hashed := sha1.New()
	hashed.Write([]byte(sortedStr))
	log.Printf("hashed: %s", hashed.Sum(nil))
	localSignature := hex.EncodeToString(hashed.Sum(nil))
	log.Printf("localSignature: %s", localSignature)

	// 验证签名
	if localSignature == event.QueryParameters.Signature {
		log.Printf("Signature is correct!")
		log.Printf("remote Signture is: %s", event.QueryParameters.Signature)
		log.Printf("local Signature is: %s", localSignature)

		// 验证通过，返回成功响应
		repose.StatusCode = http.StatusOK
		repose.Body = event.QueryParameters.EchoStr
		repose.Headers["ContentType"] = ContentType.Text
		repose.IsBase64Encoded = false
		err = nil
		return
	} else {
		// 签名不匹配，返回错误响应
		log.Printf("Signature is not correct")
		log.Printf("remote Signture is: %s", event.QueryParameters.Signature)
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
