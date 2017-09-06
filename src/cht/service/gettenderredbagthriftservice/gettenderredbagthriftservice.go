package gettenderredbagthriftservice

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/gettenderredbag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type gettenderredservice struct{}

const (
	QUERY_RED_BAG_FAILED  = 1001
	QUERY_RED_BAG_SUCCESS = 1002
)

var Status = map[int]string{
	QUERY_RED_BAG_FAILED:  "查询红包金额出错",
	QUERY_RED_BAG_SUCCESS: "查询红包金额成功",
}

func (gts *gettenderredservice) GetRedbagInfo(requestObj *TenderRedbagRequestStruct) (r *TenderRedbagResponseStruct, err error) {
	trr := new(gettenderredbag.TenderRedbagRequest)
	trr.UserId = requestObj.GetUserId()
	trr.TenderId = requestObj.GetTenderId()
	trr.RedId = requestObj.GetRedId()
	trr.TenderMoney = requestObj.GetTenderMoney()
	trr.TimeLimit = requestObj.GetTimeLimit()
	trr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, err := gettenderredbag.GetRedBagMoney(trr)
	if err != nil {
		return &TenderRedbagResponseStruct{
			Status: QUERY_RED_BAG_FAILED,
			Msg:    Status[QUERY_RED_BAG_FAILED],
		}, nil
	}

	Logger.Debug("GetRedbagInfo red ", res)

	return &TenderRedbagResponseStruct{
		Status:      QUERY_RED_BAG_SUCCESS,
		Msg:         Status[QUERY_RED_BAG_SUCCESS],
		RedbagMoney: res,
	}, nil
}

func StartGetTenderRedBagServer() {
	zkServers := []string{"192.168.8.212:2181", "192.168.8.213:2181", "192.168.8.214:2181"}
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30007"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/GetTenderRedbagThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &gettenderredservice{}
	processor := NewGetTenderRedbagThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
