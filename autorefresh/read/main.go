package refresh

type storeKey struct {
	PrimaryKey  int8
	AccessToken string
	expiresIn   int16
}

//func read() (AccessToken string, expiresIn int16) {
//	client := tablestore.NewClient(os.Getenv("Endpoint"), os.Getenv("AccessKeyId"), os.Getenv("AccessKeySecret"), os.Getenv("InstanceName"))
//	getRowRequest := new(tablestore.GetRowRequest)
//	criteria := new(tablestore.SingleRowQueryCriteria)
//	getPk := new(tablestore.PrimaryKey)
//	Key := storeKey{
//		PrimaryKey:  1,
//		AccessToken: AccessToken,
//		expiresIn:   expiresIn,
//	}
//	getPk.AddPrimaryKeyColumn("PrimaryKey", Key.PrimaryKey)
//	criteria.PrimaryKey = getPk
//	getRowRequest.SingleRowQueryCriteria = criteria
//	getRowRequest.SingleRowQueryCriteria.TableName = "AccessToken"
//	getRowRequest.SingleRowQueryCriteria.MaxVersion = 1
//	getRowRequest.SingleRowQueryCriteria.ColumnsToGet = []string{"AccessToken", "expiresIn"}
//	getResponse, err := client.GetRow(getRowRequest)
//
//	return AccessToken, expiresIn
//}
