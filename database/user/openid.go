package database

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"os"
)

func CreateOpenId(openId string) error {
	// 初始化客户端
	client := tablestore.NewClient(
		os.Getenv("endPoint"),
		os.Getenv("instanceName"),
		os.Getenv("accessKeyId"),
		os.Getenv("accessKeySecret"),
		nil,
	)

	// 创建PutRowChange对象
	putRowRequest := new(tablestore.PutRowRequest)
	PutRowChange := new(tablestore.PutRowChange)

	// 设置主键并设置数据
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("token", openId)

	// 设置目标表
	PutRowChange.TableName = "token"
	PutRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	putRowRequest.PutRowChange = PutRowChange

	_, err := client.PutRow(putRowRequest)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// UpdateOpenId 更新OpenID
func UpdateOpenId(openId string) (err error) {
	// 初始化客户端
	client := tablestore.NewClient(
		os.Getenv("endPoint"),        // 获取环境变量中的endPoint
		os.Getenv("instanceName"),    // 获取环境变量中的instanceName
		os.Getenv("accessKeyId"),     // 获取环境变量中的accessKeyId
		os.Getenv("accessKeySecret"), // 获取环境变量中的accessKeySecret
		nil,
	)
	UpdateRowRequest := new(tablestore.UpdateRowRequest)
	UpdateRowChange := new(tablestore.UpdateRowChange)
	UpdateRowChange.TableName = "token" // 设置要更新的表名为"token"
	updatePk := new(tablestore.PrimaryKey)
	updatePk.AddPrimaryKeyColumn("token", openId) // 添加要更新的数据的主键列"token"和对应的值openId
	UpdateRowChange.PrimaryKey = updatePk
	UpdateRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE) // 设置更新条件为忽略该行的存在性
	UpdateRowRequest.UpdateRowChange = UpdateRowChange                      // 将UpdateRowChange赋值给UpdateRowRequest的UpdateRowChange字段
	_, err = client.UpdateRow(UpdateRowRequest)                             // 调用UpdateRow方法更新数据
	if err != nil {
		fmt.Println("update failed with error:") // 如果更新失败，打印错误信息
		return err
	} else {
		fmt.Println("update row finished") // 如果更新成功，打印更新完成信息
		return nil
	}
}
