package database

import (
	"errors"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"os"

	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/gen"
)

// AddReserve 函数用于添加预约信息到表格存储中
func AddReserve(newReserve *datastruct.AddReserve, AuthData *datastruct.AuthStruct) (reserveID int64, err error) {
	client := tablestore.NewClient(
		os.Getenv("EndPoint"),
		os.Getenv("InstanceName"),
		os.Getenv("AccessKeyId"),
		os.Getenv("AccessKeySecret"),
	)
	putRowRequest := new(tablestore.PutRowRequest)
	putRowChange := new(tablestore.PutRowChange)
	putRowChange.TableName = "reserve"
	putRowChange.PrimaryKey = new(tablestore.PrimaryKey)
	reserveID = int64(gen.UniqueReserveID())
	putRowChange.PrimaryKey.AddPrimaryKeyColumn("ID", reserveID)
	putRowChange.AddColumn("DoctorID", newReserve.DoctorID)
	putRowChange.AddColumn("Mobile", newReserve.Mobile)
	putRowChange.AddColumn("Name", newReserve.Name)
	putRowChange.AddColumn("Time", newReserve.Time)
	putRowChange.AddColumn("Token", AuthData.OpenID)
	putRowChange.AddColumn("Status", 1)
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	putRowRequest.PutRowChange = putRowChange
	_, err = client.PutRow(putRowRequest)
	return reserveID, nil
}

// QueryReserve 函数用于查询预约信息
func QueryReserve(OpenID []string) (queryResult []datastruct.SingleReserve, err error) {
	client := tablestore.NewClient(
		os.Getenv("EndPoint"),
		os.Getenv("InstanceName"),
		os.Getenv("AccessKeyId"),
		os.Getenv("AccessKeySecret"),
	)
	batchGetReq := new(tablestore.BatchGetRowRequest)
	mqCriteria := new(tablestore.MultiRowQueryCriteria)
	mqCriteria.TableName = "reserve"
	var readedDataConst int
	for length := len(OpenID); readedDataConst < length; readedDataConst++ {
		pkToGet := new(tablestore.PrimaryKey)
		pkToGet.AddPrimaryKeyColumn("OpenID", OpenID[readedDataConst])
		mqCriteria.AddRow(pkToGet)
		mqCriteria.MaxVersion = 1
	}
	batchGetReq.MultiRowQueryCriteria = append(batchGetReq.MultiRowQueryCriteria, mqCriteria)
	batchGetResp, err := client.BatchGetRow(batchGetReq)
	if err != nil {
		return nil, err
	}
	rows, ok := batchGetResp.TableToRowsResult["reserve"]
	if !ok {
		return nil, errors.New("no results found for table")
	}
	for _, row := range rows {
		var singleReserve datastruct.SingleReserve
		// 从主键列获取ID
		for _, idValuePointer := range row.PrimaryKey.PrimaryKeys {
			switch idValuePointer.ColumnName {
			case "ID":
				singleReserve.ID = int32(idValuePointer.Value.(int64))
			default:
				break
			}

		}
		for _, col := range row.Columns {
			switch col.ColumnName {
			case "DoctorID":
				singleReserve.DoctorID = col.Value.(string)
			case "Mobile":
				singleReserve.Mobile = col.Value.(int64)
			case "Name":
				singleReserve.Name = col.Value.(string)
			case "Time":
				singleReserve.Time = col.Value.(int64)
			case "DoctorAvatar":
				singleReserve.DocAvatar = col.Value.(string)
			case "DoctorName":
				singleReserve.DoctorName = col.Value.(string)
			case "Status":
				singleReserve.Status = int8(col.Value.(int64))
			case "OpenID":
				singleReserve.OpenID = col.Value.(string)
			}
		}
		if singleReserve.DoctorID != "" && singleReserve.Mobile != 0 && singleReserve.Name != "" && singleReserve.Time != 0 && singleReserve.DocAvatar != "" && singleReserve.DoctorName != "" && singleReserve.Status != 0 && singleReserve.OpenID != "" {
			queryResult = append(queryResult, singleReserve)
		}
	}
	return queryResult, nil
}
func QueryReserveSingle(reserveID int32) (reserveData *datastruct.SingleReserve, err error) {
	client := tablestore.NewClient(
		os.Getenv("EndPoint"),
		os.Getenv("InstanceName"),
		os.Getenv("AccessKeyId"),
		os.Getenv("AccessKeySecret"),
	)
	getRowRequest := new(tablestore.GetRowRequest)
	criteria := new(tablestore.SingleRowQueryCriteria)
	criteria.TableName = "reserve"
	criteria.PrimaryKey = new(tablestore.PrimaryKey)
	criteria.PrimaryKey.AddPrimaryKeyColumn("ID", reserveID)
	criteria.MaxVersion = 1
	getRowRequest.SingleRowQueryCriteria = criteria
	getRowResp, err := client.GetRow(getRowRequest)
	if err != nil {
		return nil, err
	}
	reserveData = new(datastruct.SingleReserve)
	index := 0
	for _, reserveIDS := range getRowResp.PrimaryKey.PrimaryKeys {
		// 在循环开始之前确保Column Name是预期的
		if reserveIDS.ColumnName != "ID" {
			continue
		}
		// 因为已知是"ID"，所以可以直接转换并赋值
		reserveData.ID = int32(reserveIDS.Value.(int64))
	}
	for range getRowResp.Columns {
		switch {
		case getRowResp.Columns[index].ColumnName == "DoctorID":
			reserveData.DoctorID = getRowResp.Columns[index].Value.(string)
		case getRowResp.Columns[index].ColumnName == "Mobile":
			reserveData.Mobile = getRowResp.Columns[index].Value.(int64)
		case getRowResp.Columns[index].ColumnName == "Name":
			reserveData.Name = getRowResp.Columns[index].Value.(string)
		case getRowResp.Columns[index].ColumnName == "Time":
			reserveData.Time = getRowResp.Columns[index].Value.(int64)
		case getRowResp.Columns[index].ColumnName == "DoctorAvatar":
			reserveData.DocAvatar = getRowResp.Columns[index].Value.(string)
		case getRowResp.Columns[index].ColumnName == "DoctorName":
			reserveData.DoctorName = getRowResp.Columns[index].Value.(string)
		case getRowResp.Columns[index].ColumnName == "Status":
			reserveData.Status = getRowResp.Columns[index].Value.(int8)
		case getRowResp.Columns[index].ColumnName == "OpenID":
			reserveData.OpenID = getRowResp.Columns[index].Value.(string)
		default:
			continue
		}
		index++
	}
	return reserveData, nil
}
func CancelReserve(reserveID int32) (err error) {
	client := tablestore.NewClient(
		os.Getenv("EndPoint"),
		os.Getenv("InstanceName"),
		os.Getenv("AccessKeyId"),
		os.Getenv("AccessKeySecret"),
	)
	updateRowRequest := new(tablestore.UpdateRowRequest)
	updateRowChange := new(tablestore.UpdateRowChange)
	updateRowChange.TableName = "reserve"
	updateRowChange.PrimaryKey = new(tablestore.PrimaryKey)
	updateRowChange.PrimaryKey.AddPrimaryKeyColumn("ID", reserveID)
	updateRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	updateRowChange.PutColumn("Status", 0)
	_, err = client.UpdateRow(updateRowRequest)
	if err != nil {
		return err
	}
	return nil
}
