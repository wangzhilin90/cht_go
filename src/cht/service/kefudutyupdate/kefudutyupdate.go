package kefudutyupdate

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	kdu "cht/models/kefudutyupdate"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	UPDATE_KEFU_DUTY_SUCCESS = 1000
	UPDATE_KEFU_DUTY_FAILED  = 1001
)

var Stat = map[int]string{
	UPDATE_KEFU_DUTY_SUCCESS: "客服值班---修改值班成功",
	UPDATE_KEFU_DUTY_FAILED:  "客服值班---修改值班失败",
}

type kefudutyupdateservice struct{}

func (kdus *kefudutyupdateservice) UpdateKefuDuty(requestObj *KefuDutyUpdateRequestStruct) (r *KefuDutyUpdateResponseStruct, err error) {
	kdur := new(kdu.KefuDutyUpdateRequest)
	kdur.ID = requestObj.GetID()
	kdur.Customer = requestObj.GetCustomer()
	kdur.DutyTime = requestObj.GetDutyTime()
	kdur.HolidayUser = requestObj.GetHolidayUser()
	kdur.IsRest = requestObj.GetIsRest()
	kdur.Starttime = requestObj.GetStarttime()
	kdur.Endtime = requestObj.GetEndtime()
	kdur.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	b := kdu.UpdateKefuDuty(kdur)
	if b == false {
		Logger.Debugf("UpdateKefuDuty update failed")
		return &KefuDutyUpdateResponseStruct{
			Status: UPDATE_KEFU_DUTY_FAILED,
			Msg:    Stat[UPDATE_KEFU_DUTY_FAILED],
		}, nil
	}

	return &KefuDutyUpdateResponseStruct{
		Status: UPDATE_KEFU_DUTY_SUCCESS,
		Msg:    Stat[UPDATE_KEFU_DUTY_SUCCESS],
	}, nil
}

/**
 * [StartKeFuDutyUpdateServer 客服值班---修改值班服务]
 * @DateTime 2017-10-30T15:07:59+0800
 */
func StartKeFuDutyUpdateServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30063"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/KefuDutyUpdateThriftService/providers"
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

	handler := &kefudutyupdateservice{}
	processor := NewKefuDutyUpdateThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
