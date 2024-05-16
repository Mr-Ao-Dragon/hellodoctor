package database

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"

	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
)

// QueryUserExist 根据OpenID查询用户是否存在
func QueryUserExist(OpenID string) (queryResult bool, err error) {
	client := tablestore.NewClient(
		os.Getenv("EndPoint"),
		os.Getenv("InstanceName"),
		os.Getenv("AccessKeyId"),
		os.Getenv("AccessKeySecret"),
	)
	getRowRequest := new(tablestore.GetRowRequest)
	criteria := new(tablestore.SingleRowQueryCriteria)
	criteria.TableName = "user"
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("OpenID", OpenID)
	criteria.PrimaryKey = putPk
	criteria.MaxVersion = 1
	getRowRequest.SingleRowQueryCriteria = criteria
	Result, err := client.GetRow(getRowRequest)
	if Result.Columns == nil {
		return false, nil
	}
	if err != nil {
		log.Fatalf("数据库读写出错，错误信息：%#v\n请求内容为：%+v", err, client)
		return false, err
	}
	return true, nil

}

func QueryLogin(authStruct *datastruct.AuthStruct) (loginStat bool, err error) {
	client := tablestore.NewClient(
		os.Getenv("EndPoint"),
		os.Getenv("InstanceName"),
		os.Getenv("AccessKeyId"),
		os.Getenv("AccessKeySecret"),
	)
	getRowRequest := new(tablestore.GetRowRequest)
	criteria := new(tablestore.SingleRowQueryCriteria)
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("OpenID", authStruct.OpenID)
	criteria.PrimaryKey = putPk
	criteria.TableName = "user"
	criteria.MaxVersion = 1
	getRowRequest.SingleRowQueryCriteria = criteria
	criteria.ColumnsToGet = []string{"ExpiresIn", "SystemToken"}
	Result, err := client.GetRow(getRowRequest)
	if Result == nil {
		err = errors.New("查询用户信息失败，查询结果为空")
		log.Fatalf(err.Error())
		return false, err
	}
	if err != nil {
		log.Fatalf("数据库读写出错，错误信息：%#v请求id为：%s", err, Result.RequestId)
		return false, err
	}
	Column := Result.Columns
	var SystemToken string
	var expiresIn int64
	index := 0
	for range Column {
		switch Column[index].ColumnName {
		case "SystemToken":
			SystemToken = Column[index].Value.(string)
		case "ExpiresIn":
			expiresIn = Column[index].Value.(int64)
		}
		index++
		if SystemToken != "" && expiresIn != 0 {
			break
		}
	}
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

// UserLogin 用户登录
func UserLogin(AuthData *datastruct.AuthStruct) (isSusses bool, expiresIn int64, expressed bool, err error) {
	checkUserExist, err := QueryUserExist(AuthData.OpenID)
	NowUnix := time.Now().Unix()
	expiresIn = NowUnix + (86400 * 30)
	if checkUserExist == false {
		return false, NowUnix, true, nil
	}
	client := tablestore.NewClient(
		os.Getenv("EndPoint"),
		os.Getenv("InstanceName"),
		os.Getenv("AccessKeyId"),
		os.Getenv("AccessKeySecret"),
	)
	getRowRequest := new(tablestore.GetRowRequest)
	criteria := new(tablestore.SingleRowQueryCriteria)
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("OpenID", AuthData.OpenID)
	criteria.PrimaryKey = putPk
	criteria.MaxVersion = 1
	criteria.TableName = "user"
	getRowRequest.SingleRowQueryCriteria = criteria
	queryExpiresIn, err := client.GetRow(getRowRequest)
	if err != nil {
		log.Fatalf("数据库读写出错，错误信息：%#v", err)
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
	UpdateRowChange.PutColumn("LoginTime", int(NowUnix))
	UpdateRowChange.PutColumn("ExpiresIn", int(expiresIn))
	UpdateRowChange.PutColumn("SystemToken", AuthData.SystemToken)
	UpdateRowRequest.UpdateRowChange = UpdateRowChange
	_, err = client.UpdateRow(UpdateRowRequest)
	if err != nil {
		log.Fatalf("数据库读写出错，错误信息：%#v", err)
		return false, NowUnix, true, err
	}
	return true, expiresIn, false, nil
}

func AddUser(OpenID string, PermissionLevel int8, systemToken string) (isSusses bool, expiresIn int64, err error) {
	NowUnix := time.Now().Unix()
	expiresIn = NowUnix + (86400 * 30)
	client := tablestore.NewClient(
		os.Getenv("EndPoint"),
		os.Getenv("InstanceName"),
		os.Getenv("AccessKeyId"),
		os.Getenv("AccessKeySecret"),
	)
	putRowRequest := new(tablestore.PutRowRequest)
	putRowChange := new(tablestore.PutRowChange)
	putRowChange.TableName = "user"
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("OpenID", OpenID)
	putRowChange.AddColumn("LoginTime", int(NowUnix))
	putRowChange.AddColumn("ExpiresIn", int(expiresIn))
	putRowChange.AddColumn("SystemToken", systemToken)
	putRowChange.AddColumn("permission", int(PermissionLevel))
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	putRowRequest.PutRowChange = putRowChange
	putRowRequest.PutRowChange.PrimaryKey = putPk
	defer log.Printf("the error is: %+v", err)
	_, err = client.PutRow(putRowRequest)
	defer log.Printf("the error is: %+v", err)
	if err != nil {
		log.Printf("数据库读写出错，错误信息：%+v", err)
		return false, time.Now().Unix(), err
	}
	return true, time.Now().Unix(), nil
}
