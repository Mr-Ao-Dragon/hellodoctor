package database

import (
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/gen"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"os"
)

func AddAd(single *datastruct.AdDataSingle) (AdDataSingle *datastruct.AdDataSingle, err error) {
	// 初始化数据库
	client := tablestore.NewClientWithConfig(
		os.Getenv("AccessKeyId"),
		os.Getenv("AccessKeySecret"),
		os.Getenv("EndPoint"),
		os.Getenv("InstanceName"),
		"",
		nil,
	)
	// 接收函数
	putRowRequest := new(tablestore.PutRowRequest)
	putRowChange := new(tablestore.PutRowChange)
	putPk := new(tablestore.PrimaryKey)
	single.AdID, err = gen.Token(16)
	putRowChange.TableName = "AdData"
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	putRowChange.PrimaryKey = putPk
	putRowRequest.PutRowChange = putRowChange
	AdDataSingle = single
	putPk.AddPrimaryKeyColumn("AdID", single.AdID)
	putRowChange.AddColumn("TimeOut", single.TimeOut)
	putRowChange.AddColumn("Message", single.Message)
	// 数据发送
	putRowChange.PrimaryKey = putPk
	putRowRequest.PutRowChange = putRowChange
	_, err = client.PutRow(putRowRequest)
	// 错误处理
	if err != nil {
		return nil, err
	}
	err = nil
	return
}
