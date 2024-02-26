package user

type QueryResult struct {
	Code int                `json:"code"`
	Msg  string             `json:"msg"`
	Data []SingleDataStruct `json:"data"`
}

type SingleDataStruct struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Id     int    `json:"id"`
}

// func ListDoctor() (err error) {
// 	client := tablestore.NewClientWithConfig(
// 		os.Getenv("AccessKeyId"),
// 		os.Getenv("AccessKeySecret"),
// 		os.Getenv("EndPoint"),
// 		os.Getenv("InstanceName"),
// 		"",
// 		nil,
// 	)
// 	getRangeRequest := &tablestore.GetRangeRequest{}
// 	rangeRowQueryCriteria := &tablestore.RangeRowQueryCriteria{}
// 	rangeRowQueryCriteria.TableName = "doctor"
// 	rangeRowQueryCriteria.StartPrimaryKey = &tablestore.PrimaryKey{}
// 	rangeRowQueryCriteria.StartPrimaryKey.AddPrimaryKeyColumn("OpenID", tablestore.VT_INF_MIN)
// 	rangeRowQueryCriteria.EndPrimaryKey = &tablestore.PrimaryKey{}
// 	rangeRowQueryCriteria.EndPrimaryKey.AddPrimaryKeyColumn("OpenID", tablestore.VT_INF_MAX)
// 	rangeRowQueryCriteria.Direction = tablestore.FORWARD
// 	rangeRowQueryCriteria.MaxVersion = 1
// 	getRangeRequest.RangeRowQueryCriteria = rangeRowQueryCriteria
// 	getRangeResp, err := client.GetRange(getRangeRequest)
// 	rows := getRangeResp.Rows
// 	OpenIDs := rows. // 这里我要怎么写呢？
// 	return nil
// }
