package database

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"

	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
	"github.com/Mr-Ao-Dragon/hellodoctor/tool/gen"
)

// QueryUserExist 根据OpenID查询用户是否存在
func QueryUserExist(OpenID string) (queryResult bool, err error) {
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
	criteria.TableName = "user"
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("OpenID", OpenID)
	criteria.PrimaryKey = putPk
	getRowRequest.SingleRowQueryCriteria = criteria
	Result, err := client.GetRow(getRowRequest)
	if err != nil {
		return false, err
	}
	if Result != nil {
		return true, nil
	} else {
		return false, nil
	}
}

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

func QueryLogin(authStruct *datastruct.AuthStruct) (isSusses bool, err error) {
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
	putPk.AddPrimaryKeyColumn("OpenID", authStruct.OpenID)
	criteria.PrimaryKey = putPk
	criteria.TableName = "user"
	getRowRequest.SingleRowQueryCriteria = criteria
	criteria.ColumnsToGet = []string{"ExpiresIn", "SystemToken"}
	Result, err := client.GetRow(getRowRequest)
	if Result == nil {
		err = errors.New("查询用户信息失败，查询结果为空")
		log.Fatalf(err.Error())
		return false, err
	}
	if err != nil {
		return false, err
	}
	Column := Result.Columns
	expiresIn := Column[0].Value.(int64)
	SystemToken := Column[1].Value.(string)
	if SystemToken != authStruct.SystemToken {
		err = errors.New("查询用户信息失败，OpenID为" + authStruct.OpenID + "的用户凭据无效")
		log.Fatalf(err.Error())
		return false, err
	}
	if expiresIn <= time.Now().Unix() {
		err = errors.New("查询用户信息失败，OpenID为" + authStruct.OpenID + "的用户登录态过期")
		log.Fatalf(err.Error())
		return false, err
	}
	return true, nil
}

func QueryDoctor(OpenID string) (result *datastruct.SingleDoctorDataStruct, err error) {
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
	getPk := new(tablestore.PrimaryKey)
	getPk.AddPrimaryKeyColumn("OpenID", OpenID)
	criteria.PrimaryKey = getPk
	criteria.TableName = "doctor"
	criteria.AddColumnToGet("name")
	criteria.AddColumnToGet("avatar")
	getRowRequest.SingleRowQueryCriteria = criteria
	queryResult, err := client.GetRow(getRowRequest)
	if err != nil {
		return nil, err
	}
	Column := queryResult.Columns
	var name string
	var avatar string
	var index int
	for ; index < len(Column); index++ {
		if Column[index].ColumnName == "name" {
			name = Column[index].Value.(string)
		} else if Column[index].ColumnName == "avatar" {
			avatar = Column[index].Value.(string)
		}
		if index >= 3 && (name == "" || avatar == "") {
			err = errors.New("查询用户信息失败，数据损坏，请联系管理员")
			log.Fatalf(err.Error())
			return nil, err
		}
	}
	result = &datastruct.SingleDoctorDataStruct{
		Name:   name,
		Avatar: avatar,
		Id:     OpenID,
	}
	return result, nil
}

// UserLogin 用户登录
func UserLogin(AuthData *datastruct.AuthStruct) (isSusses bool, expiresIn int64, expressed bool, err error) {
	checkUserExist, err := QueryUserExist(AuthData.OpenID)
	NowUnix := time.Now().Unix()
	expiresIn = NowUnix + (86400 * 30)
	if checkUserExist == false {
		return false, NowUnix, true, nil
	}
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
	putPk.AddPrimaryKeyColumn("OpenID", AuthData.OpenID)
	criteria.PrimaryKey = putPk
	getRowRequest.SingleRowQueryCriteria = criteria
	getRowRequest.SingleRowQueryCriteria.TableName = "user"
	queryExpiresIn, err := client.GetRow(getRowRequest)
	if err != nil {
		return false, NowUnix, true, err
	}
	desiredColumnNumber := 0
	queryExpiresInResult := queryExpiresIn.Columns
	for ; desiredColumnNumber < len(queryExpiresInResult); desiredColumnNumber++ {
		if queryExpiresInResult[desiredColumnNumber].ColumnName == "ExpiresIn" {
			expiresIn = queryExpiresInResult[desiredColumnNumber].Value.(int64)
		}
	}
	if queryExpiresInResult == nil {
		return false, NowUnix, true, nil
	}
	if expiresIn >= 86400 {
		return true, expiresIn, false, nil
	}
	UpdateRowRequest := new(tablestore.UpdateRowRequest)
	UpdateRowChange := new(tablestore.UpdateRowChange)
	UpdateRowChange.TableName = "user"
	UpdatePk := new(tablestore.PrimaryKey)
	UpdatePk.AddPrimaryKeyColumn("OpenID", AuthData.OpenID)
	UpdateRowChange.PrimaryKey = UpdatePk
	UpdateRowChange.PutColumn("LoginTime", NowUnix)
	UpdateRowChange.PutColumn("ExpiresIn", expiresIn)
	UpdateRowChange.PutColumn("SystemToken", AuthData.SystemToken)
	UpdateRowRequest.UpdateRowChange = UpdateRowChange
	_, err = client.UpdateRow(UpdateRowRequest)
	if err != nil {
		return false, NowUnix, true, err
	} else {
		return true, expiresIn, false, nil
	}
}

func AddUser(OpenID string, PermissionLevel int8, systemToken string) (isSusses bool, expiresIn int64, err error) {
	NowUnix := time.Now().Unix()
	expiresIn = NowUnix + (86400 * 30)
	client := tablestore.NewClientWithConfig(
		os.Getenv("AccessKeyId"),
		os.Getenv("AccessKeySecret"),
		os.Getenv("EndPoint"),
		os.Getenv("InstanceName"),
		"",
		nil,
	)
	putRowRequest := new(tablestore.PutRowRequest)
	putRowChange := new(tablestore.PutRowChange)
	putRowChange.TableName = "user"
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("OpenID", OpenID)
	putRowChange.AddColumn("LoginTime", NowUnix)
	putRowChange.AddColumn("ExpiresIn", expiresIn)
	putRowChange.AddColumn("SystemToken", systemToken)
	putRowChange.AddColumn("permission", PermissionLevel)
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	putRowRequest.PutRowChange = putRowChange
	putRowRequest.PutRowChange.PrimaryKey = putPk
	_, err = client.PutRow(putRowRequest)
	if err != nil {
		return false, time.Now().Unix(), err
	} else {
		return true, time.Now().Unix(), nil
	}
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

func UpToDoctor(dataStruct *datastruct.SingleDoctorDataStruct) (err error) {
	client := tablestore.NewClientWithConfig(
		os.Getenv("AccessKeyId"),
		os.Getenv("AccessKeySecret"),
		os.Getenv("EndPoint"),
		os.Getenv("InstanceName"),
		"",
		nil,
	)
	putRowRequest := new(tablestore.PutRowRequest)
	putRowChange := new(tablestore.PutRowChange)
	putRowChange.TableName = "doctor"
	putRowChange.PrimaryKey = new(tablestore.PrimaryKey)
	putRowChange.PrimaryKey.AddPrimaryKeyColumn("OpenID", dataStruct.Id)
	putRowChange.AddColumn("doctorAvatar", dataStruct.Avatar)
	putRowChange.AddColumn("doctorName", dataStruct.Name)
	putRowChange.AddColumn("doctorProfile", dataStruct.Profile)
	putRowRequest.PutRowChange = putRowChange
	_, err = client.PutRow(putRowRequest)
	if err != nil {
		return err
	}
	return nil
}

func ListDoctor() (queryResult []datastruct.SingleDoctorDataStruct, err error) {
	client := tablestore.NewClientWithConfig(
		os.Getenv("AccessKe yId"),
		os.Getenv("AccessKeySecret"),
		os.Getenv("EndPoint"),
		os.Getenv("InstanceName"),
		"",
		nil,
	)
	getRangeRequest := &tablestore.GetRangeRequest{}
	rangeRowQueryCriteria := &tablestore.RangeRowQueryCriteria{}
	rangeRowQueryCriteria.TableName = "doctor"
	rangeRowQueryCriteria.StartPrimaryKey = &tablestore.PrimaryKey{}
	rangeRowQueryCriteria.StartPrimaryKey.AddPrimaryKeyColumn("OpenID", tablestore.VT_INF_MIN)
	rangeRowQueryCriteria.EndPrimaryKey = &tablestore.PrimaryKey{}
	rangeRowQueryCriteria.EndPrimaryKey.AddPrimaryKeyColumn("OpenID", tablestore.VT_INF_MAX)
	rangeRowQueryCriteria.AddColumnToGet("Name")
	rangeRowQueryCriteria.AddColumnToGet("Avatar")
	rangeRowQueryCriteria.AddColumnToGet("id")
	rangeRowQueryCriteria.AddColumnToGet("Profile")
	rangeRowQueryCriteria.Direction = tablestore.FORWARD
	rangeRowQueryCriteria.MaxVersion = 1
	getRangeRequest.RangeRowQueryCriteria = rangeRowQueryCriteria
	getRangeResp, err := client.GetRange(getRangeRequest)
	rows := getRangeResp.Rows
	var resultData []datastruct.SingleDoctorDataStruct
	for _, row := range rows {
		var name, avatar, id, profile string
		for _, col := range row.Columns {
			switch col.ColumnName {
			case "Name":
				name = col.Value.(string)
				if name != "" { // 可以在这里添加条件来确保 Name 已找到
					break
				}
			case "Avatar":
				avatar = col.Value.(string)
				if avatar != "" { // 可以在这里添加条件来确保 Avatar 已找到
					break
				}
			case "Profile":
				profile = col.Value.(string)
			}
		}
		for _, PK := range row.PrimaryKey.PrimaryKeys {
			if PK.ColumnName == "OpenID" {
				id = PK.Value.(string)
				break
			}
		}
		if name != "" && avatar != "" && id != "" {
			singleData := datastruct.SingleDoctorDataStruct{
				Name:    name,
				Avatar:  avatar,
				Id:      id,
				Profile: profile,
			}
			resultData = append(resultData, singleData)
		}
	}
	return resultData, nil
}

// AddReserve 函数用于添加预约信息到表格存储中
func AddReserve(newReserve *datastruct.AddReserve, AuthData *datastruct.AuthStruct) (reserveID int32, err error) {
	client := tablestore.NewClientWithConfig(
		os.Getenv("AccessKeyId"),
		os.Getenv("AccessKeySecret"),
		os.Getenv("EndPoint"),
		os.Getenv("InstanceName"),
		"",
		nil,
	)
	putRowRequest := new(tablestore.PutRowRequest)
	putRowChange := new(tablestore.PutRowChange)
	putRowChange.TableName = "reserve"
	putRowChange.PrimaryKey = new(tablestore.PrimaryKey)
	reserveID = gen.UniqueReserveID()
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
	client := tablestore.NewClientWithConfig(
		os.Getenv("AccessKeyId"),
		os.Getenv("AccessKeySecret"),
		os.Getenv("EndPoint"),
		os.Getenv("InstanceName"),
		"",
		nil,
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
		idValue := row.PrimaryKey.PrimaryKeys[0].Value
		if !ok {
			return nil, errors.New("failed to parse ID from primary key")
		}
		intValue := idValue.(*tablestore.AttributeColumn).Value
		idInt64, err := strconv.ParseInt(intValue.(string), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse ID from primary key: %v", err)
		}
		singleReserve.ID = int32(idInt64)
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
				singleReserve.Status = col.Value.(int8)
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
func CancelReserve(reserveID int32) (err error) {
	client := tablestore.NewClientWithConfig(
		os.Getenv("AccessKeyId"),
		os.Getenv("AccessKeySecret"),
		os.Getenv("EndPoint"),
		os.Getenv("InstanceName"),
		"",
		nil,
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
