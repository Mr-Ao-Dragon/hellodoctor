package main

import (
	"crypto/sha1"
	"encoding/hex"
	"log"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/zhshch2002/goreq"
)

func GenerateRandomNDigits(n int) int {
	minRange := int(math.Pow10(n - 1))
	maxRange := int(math.Pow10(n)) - 1
	return rand.Intn(maxRange-minRange) + minRange
}
func main() {
	token := os.Getenv("token")
	if token == "" {
		panic("token is empty")
	}
	timestamp := GenerateRandomNDigits(10)
	echoStr := GenerateRandomNDigits(19)
	nonce := GenerateRandomNDigits(9)
	strSlice := []string{strconv.Itoa(timestamp), strconv.Itoa(nonce), token}
	sort.Strings(strSlice)
	sortedStr := strings.Join(strSlice, "")
	hashed := sha1.New()
	hashed.Write([]byte(sortedStr))
	localSignature := hex.EncodeToString(hashed.Sum(nil))
	c := goreq.NewClient(goreq.WithRandomUA())
	resp := goreq.Get(os.Getenv("verifyAddress")).
		AddParam("signature", localSignature).
		AddParam("timestamp", strconv.Itoa(timestamp)).
		AddParam("nonce", strconv.Itoa(nonce)).
		AddParam("echostr", strconv.Itoa(echoStr)).
		SetClient(c).
		Do()
	bodyTxt, _ := resp.Txt()
	statusCode := strconv.Itoa(resp.StatusCode)
	log.Printf("status code: %s", statusCode)
	log.Printf("body: %s", bodyTxt)
	// log.Printf("local sign is: %s", localSignature)
}
