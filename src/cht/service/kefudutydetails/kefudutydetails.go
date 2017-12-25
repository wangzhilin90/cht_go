package kefudutydetails

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	kdd "cht/models/kefudutydetails"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	QUERY_KEFU_DUTY_DETAILS_SUCCESS = 1000
	QUERY_KEFU_DUTY_DETAILS_FAILED  = 1001
)

var Stat = map[int]string{
	QUERY_KEFU_DUTY_DETAILS_SUCCESS: "查询值班详情成功",
	QUERY_KEFU_DUTY_DETAILS_FAILED:  "查询值班详情失败",
}

type kefudutydetailsservice struct{}

func (kfdds *kefudutydetailsservice) GetKefuDutyDetails(requestObj *KefuDutyDetailsRequestStruct) (r *KefuDutyDetailsResponseStruct, err error) {
	kddr := new(kdd.KefuDutyDetailsRequest)
	kddr.ID = requestObj.GetID()
	kddr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, err := kdd.GetKefuDutyDetails(kddr)
	if err != nil {
		Logger.Errorf("GetKefuDutyDetails get duty details failed:%v", err)
		return &KefuDutyDetailsResponseStruct{
			Status: QUERY_KEFU_DUTY_DETAILS_FAILED,
			Msg:    Stat[QUERY_KEFU_DUTY_DETAILS_FAILED],
		}, nil
	}

	var response KefuDutyDetailsResponseStruct
	if res != nil {
		kdds := new(KefuDutyDetailsStruct)
		kdds.ID = res.ID
		kdds.Customer = res.Customer
		kdds.DutyTime = res.DutyTime
		kdds.HolidayUser = res.HolidayUser
		kdds.IsRest = res.IsRest
		kdds.Starttime = res.Starttime
		kdds.Endtime = res.Endtime
		response.KefuDutyDetails = kdds
	}
	response.Status = QUERY_KEFU_DUTY_DETAILS_SUCCESS
	response.Msg = Stat[QUERY_KEFU_DUTY_DETAILS_SUCCESS]
	Logger.Debugf("GetKefuDutyDetails response:%v", response)
	return &response, nil
}

/**
 * [StartKeFuDutyDetailsServer 客服值班---值班详情服务]
 * @DateTime 2017-10-28T14:15:54+0800
 */
func StartKeFuDutyDetailsServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30061"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/KefuDutyDetailsThriftService/providers"
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

	handler := &kefudutydetailsservice{}
	processor := NewKefuDutyDetailsThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
