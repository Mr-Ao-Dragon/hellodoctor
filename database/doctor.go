package database

import (
	"errors"
	"log"
	"os"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"

	"github.com/Mr-Ao-Dragon/hellodoctor/tool/datastruct"
)

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
