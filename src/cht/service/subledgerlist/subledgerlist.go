package subledgerlist

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/subledger"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	QUERY_SUBLEDGER_SUCCESS = 1000
	QUERY_SUBLEDGER_FAILED  = 1001
)

var Status = map[int]string{
	QUERY_SUBLEDGER_FAILED:  "查询分账人信息失败",
	QUERY_SUBLEDGER_SUCCESS: "查询分账人信息成功",
}

type subledgerservice struct{}

func (ss *subledgerservice) GetSubledgerList(requestObj *SubledgerListRequestStruct) (r *SubledgerListResponseStruct, err error) {
	Logger.Infof("GetSubledgerList requestObj:%v", requestObj)
	sr := new(subledger.SubledgerRequest)
	sr.HsZhuanrangrenStr = requestObj.GetHsZhuanrangrenStr()
	sr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	sublederList, err := subledger.GetSubledgerList(sr)
	if err != nil {
		Logger.Errorf("GetSubledgerList query failed %v", err)
		return &SubledgerListResponseStruct{
			Status: QUERY_SUBLEDGER_FAILED,
			Msg:    Status[QUERY_SUBLEDGER_FAILED],
		}, nil
	}

	var response SubledgerListResponseStruct
	for _, v := range sublederList {
		si := new(SubledgerDetailsStruct)
		si.UserID = v.UserID
		si.Realname = v.Realname
		si.CardID = v.CardID
		response.SubledgerList = append(response.SubledgerList, si)
	}
	response.Status = QUERY_SUBLEDGER_SUCCESS
	response.Msg = Status[QUERY_SUBLEDGER_SUCCESS]
	Logger.Debugf("GetSubledgerList res:%v", response)
	return &response, nil
}

/**
 * [StartLogUserLoginServer 开启做标服务---分账人服务]
 * @DateTime 2017-09-13T17:58:45+0800
 */
func StartsubledgerServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30012"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/SubledgerListThriftService/providers"
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

	handler := &subledgerservice{}
	processor := NewSubledgerListThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
