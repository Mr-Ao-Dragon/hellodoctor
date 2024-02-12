package database

import (
	"os"
	"time"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
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

// UserLogin 用户登录
func UserLogin(OpenID string, systemToken string) (isSusses bool, expiresIn int64, expressed bool, err error) {
	checkUserExist, err := QueryUserExist(OpenID)
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
	putPk.AddPrimaryKeyColumn("OpenID", OpenID)
	criteria.PrimaryKey = putPk
	getRowRequest.SingleRowQueryCriteria = criteria
	getRowRequest.SingleRowQueryCriteria.TableName = "user"
	queryExpiresIn, err := client.GetRow(getRowRequest)
	if err != nil {
		return false, NowUnix, true, err
	}
	desiredColumnNumber := 0
	queryExpiresInResult := queryExpiresIn.Columns[desiredColumnNumber]
	if queryExpiresInResult == nil {
		return false, NowUnix, true, nil
	}
	expiresIn = queryExpiresInResult.Value.(int64)
	if expiresIn >= 86400 {
		return true, expiresIn, false, nil
	}
	UpdateRowRequest := new(tablestore.UpdateRowRequest)
	UpdateRowChange := new(tablestore.UpdateRowChange)
	UpdateRowChange.TableName = "user"
	UpdatePk := new(tablestore.PrimaryKey)
	UpdatePk.AddPrimaryKeyColumn("OpenID", OpenID)
	UpdateRowChange.PrimaryKey = UpdatePk
	UpdateRowChange.PutColumn("LoginTime", NowUnix)
	UpdateRowChange.PutColumn("ExpiresIn", expiresIn)
	UpdateRowRequest.UpdateRowChange = UpdateRowChange
	_, err = client.UpdateRow(UpdateRowRequest)
	if err != nil {
		return false, NowUnix, true, err
	} else {
		return true, expiresIn, false, nil
	}
}

// AddUser 添加用户
func AddUser(OpenID string, systemToken string) (isSusses bool, expiresIn int64, err error) {
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
