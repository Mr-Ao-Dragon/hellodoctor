package database

import (
	"errors"
	"os"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

func QueryPermission(OpenID string) (PermissionLevel int8, err error) {
	client := tablestore.NewClient(
		os.Getenv("EndPoint"),
		os.Getenv("InstanceName"),
		os.Getenv("AccessKeyId"),
		os.Getenv("AccessKeySecret"),
	)
	getRowRequest := new(tablestore.GetRowRequest)
	criteria := new(tablestore.SingleRowQueryCriteria)
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("OpenID", OpenID)
	criteria.PrimaryKey = putPk
	criteria.TableName = "user"
	criteria.MaxVersion = 1
	criteria.AddColumnToGet("permission")
	getRowRequest.SingleRowQueryCriteria = criteria
	queryResult, err := client.GetRow(getRowRequest)
	if err != nil {
		return 0, err
	}
	if queryResult == nil || len(queryResult.Columns) == 0 {
		return 0, errors.New("用户不存在")
	}
	for _, col := range queryResult.Columns {
		switch col.ColumnName {
		case "permission":
			PermissionLevel = int8(col.Value.(int64))
		}
	}
	return PermissionLevel, nil
}

func SetPermission(PermLevel int, OpenID string) (isSusses bool, err error) {
	client := tablestore.NewClient(
		os.Getenv("EndPoint"),
		os.Getenv("InstanceName"),
		os.Getenv("AccessKeyId"),
		os.Getenv("AccessKeySecret"),
	)
	UpdateRowRequest := new(tablestore.UpdateRowRequest)
	UpdateRowChange := new(tablestore.UpdateRowChange)
	UpdateRowChange.TableName = "user"
	UpdatePk := new(tablestore.PrimaryKey)
	UpdatePk.AddPrimaryKeyColumn("OpenID", OpenID)
	UpdateRowChange.PrimaryKey = UpdatePk
	UpdateRowChange.PutColumn("permission", int64(PermLevel))
	UpdateRowRequest.UpdateRowChange = UpdateRowChange
	_, err = client.UpdateRow(UpdateRowRequest)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
