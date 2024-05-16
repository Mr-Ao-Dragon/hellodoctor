package database

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"os"

	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/gen"
)

func AddAd(single *datastruct.AdDataSingle) (err error) {
	// 初始化数据库
	client := tablestore.NewClient(
		os.Getenv("EndPoint"),
		os.Getenv("InstanceName"),
		os.Getenv("AccessKeyId"),
		os.Getenv("AccessKeySecret"),
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
	putPk.AddPrimaryKeyColumn("AdID", single.AdID)
	putRowChange.AddColumn("TimeOut", single.TimeOut)
	putRowChange.AddColumn("Message", single.Message)
	// 数据发送
	putRowChange.PrimaryKey = putPk
	putRowRequest.PutRowChange = putRowChange
	_, err = client.PutRow(putRowRequest)
	// 错误处理
	if err != nil {
		return err
	}
	return nil
}
func ReadAd() (Ads []datastruct.AdDataSingle, err error) {
	client := tablestore.NewClient(
		os.Getenv("EndPoint"),
		os.Getenv("InstanceName"),
		os.Getenv("AccessKeyId"),
		os.Getenv("AccessKeySecret"),
	)
	queryRequest := new(tablestore.GetRangeRequest)
	rangeQuery := new(tablestore.RangeRowQueryCriteria)
	startPK := new(tablestore.PrimaryKey)
	startPK.AddPrimaryKeyColumnWithMinValue("AdID")
	rangeQuery.StartPrimaryKey = startPK
	endPK := new(tablestore.PrimaryKey)
	endPK.AddPrimaryKeyColumnWithMinValue("AdID")
	rangeQuery.EndPrimaryKey = endPK
	queryRequest.RangeRowQueryCriteria = rangeQuery
	result, _ := client.GetRange(queryRequest)
	Ads = make([]datastruct.AdDataSingle, 0)
	Rows := result.Rows
	for _, Row := range Rows {
		ad := datastruct.AdDataSingle{}
		// 假设Row.PrimaryKey为主键数据的访问方式，且AdDataSingle结构体有相应字段
		for _, pk := range Row.PrimaryKey.PrimaryKeys {
			switch pk.ColumnName {
			case "AdID":
				ad.AdID = pk.Value.(string)
			}
		}
		// 假设Row.Columns为列数据的访问方式
		for _, column := range Row.Columns {
			switch column.ColumnName {
			case "Message":
				ad.Message = column.Value.(string)
			case "TimeOut":
				ad.TimeOut = column.Value.(int)
			}
		}
		Ads = append(Ads, ad)
	}
	return Ads, nil
}
