package collectionthriftservice

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/collection"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type collectionservice struct{}

const (
	QUERY_COLLECTION_FAILED  = 1001
	QUERY_COLLECTION_SUCCESS = 1000
)

var Status = map[int]string{
	QUERY_COLLECTION_FAILED:  "查询回款明细出错",
	QUERY_COLLECTION_SUCCESS: "查询回款明细成功成功",
}

func (cs *collectionservice) GetCollectionList(requestObj *CollectionRequestStruct) (r *CollectionListResponseStruct, err error) {
	cr := new(collection.CollectionRequest)
	cr.UserID = requestObj.GetUserID()
	cr.Starttime = requestObj.GetStarttime()
	cr.Endtime = requestObj.GetEndtime()
	cr.SearchTime = requestObj.GetSearchTime()
	cr.State = requestObj.GetState()
	cr.LimitOffset = requestObj.GetLimitOffset()
	cr.LimitNum = requestObj.GetLimitNum()
	cr.Borrowid = requestObj.GetBorrowid()
	cr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, num, err := collection.GetCollectionInfo(cr)
	if err != nil {
		return &CollectionListResponseStruct{
			Status:   QUERY_COLLECTION_FAILED,
			Msg:      Status[QUERY_COLLECTION_FAILED],
			TotalNum: 0,
		}, nil
	}

	var response CollectionListResponseStruct
	for _, v := range res {
		cis := new(CollectionInfoStruct)
		cis.Username = v.Username
		cis.Title = v.Title
		cis.IsDatetype = v.IsDatetype
		cis.TimeLimit = v.TimeLimit
		cis.Zhuanrangren = v.Zhuanrangren
		cis.RepayTime = v.RepayTime
		cis.BorrowID = v.BorrowID
		cis.Periods = v.Periods
		cis.RepayYestime = v.RepayYestime
		cis.RepayYesaccount = v.RepayYesaccount
		cis.RepayAccount = v.RepayAccount
		cis.Capital = v.Capital
		cis.Interest = v.Interest
		cis.LateInterest = v.LateInterest
		cis.LateDays = v.LateDays
		cis.Status = v.Status
		cis.InterestAdd = v.InterestAdd
		cis.OldUserID = v.OldUserID
		cis.Style = v.Style
		response.CollectionInfo = append(response.CollectionInfo, cis)
	}

	Logger.Debug("GetCollectionList res ", res)
	response.Status = QUERY_COLLECTION_SUCCESS
	response.Msg = Status[QUERY_COLLECTION_SUCCESS]
	response.TotalNum = num
	return &response, nil
}

/*获取我的账户回款明细信息*/
func StartGetCollectionListServer() {
	zkServers := []string{"192.168.8.208:2181"}
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30010"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/CollectionThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &collectionservice{}
	processor := NewCollectionThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
