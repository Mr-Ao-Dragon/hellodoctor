package main

import (
	"encoding/json"
	"github.com/Mr-Ao-Dragon/hellodoctor/mrao_wechat"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"os"

	gr "github.com/awesome-fc/golang-runtime"
)

type storeKey struct {
	PrimaryKey  int8
	AccessToken string
	expiresIn   int16
}

func initialize(ctx *gr.FCContext) error {
	ctx.GetLogger().Infoln("")
	return nil
}

func handler(ctx *gr.FCContext, event []byte) ([]byte, error) {
	b, err := json.Marshal(ctx)
	if err != nil {
		ctx.GetLogger().Error("error:", err)
	}
	AccessToken, expiresIn, err := mrao_wechat.ForceRefresh(os.Getenv("AppId"), os.Getenv("AppSecret"))
	Key := storeKey{
		PrimaryKey:  1,
		AccessToken: AccessToken,
		expiresIn:   expiresIn,
	}
	client := tablestore.NewClient(os.Getenv("Endpoint"), os.Getenv("AccessKeyId"), os.Getenv("AccessKeySecret"), os.Getenv("InstanceName"))
	updateRowRequest := new(tablestore.UpdateRowRequest)
	updateRowChange := new(tablestore.UpdateRowChange)
	updateRowChange.TableName = "AccessToken"
	updatePk := new(tablestore.PrimaryKey)
	updatePk.AddPrimaryKeyColumn("PrimaryKey", Key.PrimaryKey)
	updateRowChange.PrimaryKey = updatePk
	updateRowChange.PutColumn("AccessToken", Key.AccessToken)
	updateRowChange.PutColumn("ExpiresIn", Key.expiresIn)
	updateRowRequest.UpdateRowChange = updateRowChange
	_, err = client.UpdateRow(updateRowRequest)
	if err != nil {
		ctx.GetLogger().Error("error:", err)
		return nil, err
	} else {
		ctx.GetLogger().Infof("更新成功\n 上下文： %s", string(b))
		return event, nil
	}
}
func main() {
	gr.Start(handler, initialize)
}
