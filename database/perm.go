package database

import (
	"errors"
	"os"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

func QueryPermission(OpenID string) (PermissionLevel int8, err error) {
	client := tablestore.NewClientWithConfig(
		os.Getenv("AccessKeyId"),
		os.Getenv("AccessKeySecret"),
		os.Getenv("EndPoint"),
		os.Getenv("InstanceName"),
		"",
		nil,
	)
	getRowRequest := new(tablestore.GetRowRequest)
	criteria := new(tablestore.SingleRowQueryCriteria)
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("OpenID", OpenID)
	criteria.PrimaryKey = putPk
	criteria.TableName = "user"
	getRowRequest.SingleRowQueryCriteria = criteria
	criteria.ColumnsToGet = []string{"permission"}
	queryResult, err := client.GetRow(getRowRequest)
	if err != nil {
		return 0, err
	}
	if queryResult == nil {
		return 0, errors.New("用户不存在")
	}
	return queryResult.Columns[0].Value.(int8), nil
}

func SetPermission(PermLevel int, OpenID string) (isSusses bool, err error) {
	client := tablestore.NewClientWithConfig(
		os.Getenv("AccessKeyId"),
		os.Getenv("AccessKeySecret"),
		os.Getenv("EndPoint"),
		os.Getenv("InstanceName"),
		"",
		nil,
	)
	UpdateRowRequest := new(tablestore.UpdateRowRequest)
	UpdateRowChange := new(tablestore.UpdateRowChange)
	UpdateRowChange.TableName = "user"
	UpdatePk := new(tablestore.PrimaryKey)
	UpdatePk.AddPrimaryKeyColumn("OpenID", OpenID)
	UpdateRowChange.PrimaryKey = UpdatePk
	UpdateRowChange.PutColumn("permission", PermLevel)
	UpdateRowRequest.UpdateRowChange = UpdateRowChange
	_, err = client.UpdateRow(UpdateRowRequest)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
