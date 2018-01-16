package usercollectionlist

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/collection"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
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

func (cs *collectionservice) GetUserCollectionList(requestObj *UserCollectionListRequestStruct) (r *UserCollectionListResponseStruct, err error) {
	Logger.Infof("GetUserCollectionList requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*UserCollectionListRequestStruct)
	cr := new(collection.UserCollectionListRequest)
	cr.UserID = requestObj.GetUserID()
	cr.Starttime = requestObj.GetStarttime()
	cr.Endtime = requestObj.GetEndtime()
	cr.SearchTime = requestObj.GetSearchTime()
	cr.State = requestObj.GetState()
	cr.LimitOffset = requestObj.GetLimitOffset()
	cr.LimitNum = requestObj.GetLimitNum()
	cr.Borrowid = requestObj.GetBorrowid()
	cr.CheckZhuanrangren = requestObj.GetCheckZhuanrangren()
	cr.TenderID = requestObj.GetTenderID()
	cr.CheckOldUserID = requestObj.GetCheckOldUserID()
	cr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, num, err := collection.GetCollectionInfo(cr)
	if err != nil {
		Logger.Errorf("GetUserCollectionList get collection failed:%v", err)
		return &UserCollectionListResponseStruct{
			Status:   QUERY_COLLECTION_FAILED,
			Msg:      Status[QUERY_COLLECTION_FAILED],
			TotalNum: 0,
		}, nil
	}

	var response UserCollectionListResponseStruct
	for _, v := range res {
		cis := new(UserCollectionDetailsStruct)
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
		response.UserCollectionList = append(response.UserCollectionList, cis)
	}

	response.Status = QUERY_COLLECTION_SUCCESS
	response.Msg = Status[QUERY_COLLECTION_SUCCESS]
	response.TotalNum = num
	Logger.Debugf("GetUserCollectionList response:%v", response)
	return &response, nil
}

/*获取我的账户回款明细信息*/
func StartGetCollectionListServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30010"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/UserCollectionListThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode %v failed", servicename, err)
	}

	go func() {
		time.Sleep(time.Second * 2)
		err = zkclient.WatchNode(conn, servicename, listenAddr)
		if err != nil {
			Logger.Fatalf("WatchNode %v failed:%v", servicename, err)
		}
	}()

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &collectionservice{}
	processor := NewUserCollectionListThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
