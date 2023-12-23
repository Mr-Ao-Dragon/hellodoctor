package main

import (
	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

func main() {

	fc.StartHttp(DelOrder)
	fc.StartHttp(QueryOrder)
}

func DelOrder() {

}
func QueryOrder() {

}
