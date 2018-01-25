package usertenderredbagdestails

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/gettenderredbag"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

type gettenderredservice struct{}

const (
	QUERY_RED_BAG_SUCCESS = 1000
	QUERY_RED_BAG_FAILED  = 1001
)

var Status = map[int]string{
	QUERY_RED_BAG_SUCCESS: "查询红包金额成功",
	QUERY_RED_BAG_FAILED:  "查询红包金额出错",
}

func (gts *gettenderredservice) GetUserTenderRedbagDestails(requestObj *UserTenderRedbagDestailsRequestStruct) (r *UserTenderRedbagDestailsResponseStruct, err error) {
	Logger.Info("GetUserTenderRedbagDestails requestObj:", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*UserTenderRedbagDestailsRequestStruct)
	trr := new(gettenderredbag.TenderRedbagRequest)
	trr.UserId = requestObj.GetUserId()
	trr.TenderId = requestObj.GetTenderId()
	trr.RedId = requestObj.GetRedId()
	trr.TenderMoney = requestObj.GetTenderMoney()
	trr.TimeLimit = requestObj.GetTimeLimit()
	trr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, err := gettenderredbag.GetRedBagMoney(trr)
	if err != nil {
		Logger.Errorf("GetUserTenderRedbagDestails qeury failed:%v", err)
		return &UserTenderRedbagDestailsResponseStruct{
			Status: QUERY_RED_BAG_FAILED,
			Msg:    Status[QUERY_RED_BAG_FAILED],
		}, nil
	}

	return &UserTenderRedbagDestailsResponseStruct{
		Status:      QUERY_RED_BAG_SUCCESS,
		Msg:         Status[QUERY_RED_BAG_SUCCESS],
		RedbagMoney: res,
	}, nil
}

func StartGetTenderRedBagServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30007"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/UserTenderRedbagDestailsThriftService/providers"
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

	handler := &gettenderredservice{}
	processor := NewUserTenderRedbagDestailsThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
