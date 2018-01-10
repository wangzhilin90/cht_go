package kefudutydelete

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	kdd "cht/models/kefudutydelete"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	KEFU_DUTY_DELETE_SUCCESS = 1000
	KEFU_DUTY_DELETE_FAILED  = 1001
)

var Stat = map[int]string{
	KEFU_DUTY_DELETE_SUCCESS: "客服值班-删除值班成功",
	KEFU_DUTY_DELETE_FAILED:  "客服值班-删除值班失败",
}

type kefudutydeleteservice struct{}

func (kdds *kefudutydeleteservice) DeleteKefuDuty(requestObj *KefuDutyDeleteRequestStruct) (r *KefuDutyDeleteResponseStruct, err error) {
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*KefuDutyDeleteRequestStruct)
	kddr := new(kdd.KefuDutyDeleteRequest)
	kddr.Idstr = requestObj.GetIdstr()
	kddr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	b := kdd.DeleteKefuDuty(kddr)
	if b == false {
		Logger.Debugf("DeleteKefuDuty delete failed")
		return &KefuDutyDeleteResponseStruct{
			Status: KEFU_DUTY_DELETE_FAILED,
			Msg:    Stat[KEFU_DUTY_DELETE_FAILED],
		}, nil
	}

	return &KefuDutyDeleteResponseStruct{
		Status: KEFU_DUTY_DELETE_SUCCESS,
		Msg:    Stat[KEFU_DUTY_DELETE_SUCCESS],
	}, nil
}

/**
 * [StartKeFuDutyDeleteServer 客服值班---删除值班服务]
 * @DateTime 2017-10-28T11:31:49+0800
 */
func StartKeFuDutyDeleteServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30060"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/KefuDutyDeleteThriftService/providers"
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

	handler := &kefudutydeleteservice{}
	processor := NewKefuDutyDeleteThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
