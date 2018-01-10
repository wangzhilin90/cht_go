package kefulist

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/kefulist"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

type kefulistservice struct{}

const (
	QUERY_KEFU_LIST_SUCCESS = 1000
	QUERY_KEFU_LIST_FAILED  = 1001
)

var Status = map[int]string{
	QUERY_KEFU_LIST_SUCCESS: "查询客服列表成功",
	QUERY_KEFU_LIST_FAILED:  "查询客服列表失败",
}

func (kls *kefulistservice) GetKeFuList(requestObj *KeFuListRequestStruct) (r *KeFuListResponseStruct, err error) {
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*KeFuListRequestStruct)
	kfrs := new(kefulist.KeFuListRequestStruct)
	kfrs.Status = requestObj.GetStatus()
	kfrs.RoleID = requestObj.GetRoleID()
	kfrs.CustomerType = requestObj.GetCustomerType()
	KeFuDetailslist, err := kefulist.GetKeFuList(kfrs)
	if err != nil {
		Logger.Errorf("GetKeFuList query failed", err)
		return &KeFuListResponseStruct{
			Status: QUERY_KEFU_LIST_FAILED,
			Msg:    Status[QUERY_KEFU_LIST_FAILED],
		}, nil
	}

	var response KeFuListResponseStruct
	for _, v := range KeFuDetailslist {
		kfds := new(KeFuDetailsStruct)
		kfds.ID = v.ID
		kfds.RoleID = v.RoleID
		kfds.Account = v.Account
		kfds.Realname = v.Realname
		kfds.Password = v.Password
		kfds.Mobile = v.Mobile
		kfds.Qq = v.Qq
		kfds.Lastloginip = v.Lastloginip
		kfds.Lastlogintime = v.Lastlogintime
		kfds.CreateTime = v.CreateTime
		kfds.Status = v.Status
		kfds.Views = v.Views
		kfds.CustomerType = v.CustomerType
		response.KeFuList = append(response.KeFuList, kfds)
	}
	response.Status = QUERY_KEFU_LIST_SUCCESS
	response.Msg = Status[QUERY_KEFU_LIST_SUCCESS]
	Logger.Debugf("GetKeFuList res :%v", response)
	return &response, nil
}

func StartKeFuListsServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30021"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/KeFuListThriftService/providers"
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

	handler := &kefulistservice{}
	processor := NewKeFuListThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
