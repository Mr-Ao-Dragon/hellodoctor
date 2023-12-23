package main

import (
	"context"
	"encoding/json"
	"github.com/Mr-Ao-Dragon/hellodoctor/dbtool"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/go-pg/pg/v10"
	"net/http"
	"os"
)

func NewOrder(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	db := pg.Connect(&pg.Options{
		Addr:     os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		User:     "ReadAndWrite",
		Database: os.Getenv("DB_DATABASE"),
		Password: os.Getenv("DB_PASSWORD_RAW"),
		//ConnectTimeout: 600,
		MinIdleConns: 0,
		ReadTimeout:  180,
		WriteTimeout: 180,
	})
	ctx := context.Background()
	for db.Ping(ctx) == nil {

	}
	dbtool.ReadRole(`Lottery/Lottery`)
	json.Decoder{}
	error := dbtool.Write()
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(error))
		return error
	}
}
func main() {
	fc.StartHttp(NewOrder)
}
