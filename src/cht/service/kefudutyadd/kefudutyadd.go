package kefudutyadd

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	kda "cht/models/kefudutyadd"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	KE_FU_DUTY_ADD_SUCCESS = 1000
	KE_FU_DUTY_ADD_FAILED  = 1001
)

var Stat = map[int]string{
	KE_FU_DUTY_ADD_SUCCESS: "客服值班新增成功",
	KE_FU_DUTY_ADD_FAILED:  "客服值班新增失败",
}

type kefudutyaddservice struct{}

func NewKefuDutyAddRequest(requestObj *KefuDutyAddRequestStruct) *kda.KefuDutyAddRequest {
	kdar := new(kda.KefuDutyAddRequest)
	kdar.Customer = requestObj.GetCustomer()
	kdar.DutyTime = requestObj.GetDutyTime()
	kdar.HolidayUser = requestObj.GetHolidayUser()
	kdar.IsRest = requestObj.GetIsRest()
	kdar.Addtime = requestObj.GetAddtime()
	kdar.Starttime = requestObj.GetStarttime()
	kdar.Endtime = requestObj.GetEndtime()
	kdar.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	return kdar
}

/*分配默认值*/
func AssignDefaultValue(req *kda.KefuDutyAddRequest) *kda.KefuDutyAddRequest {
	if req.Customer == "" {
		req.Customer = " "
	}

	if req.HolidayUser == "" {
		req.HolidayUser = " "
	}
	return req
}

func (kdas *kefudutyaddservice) AddKefuDuty(requestObj *KefuDutyAddRequestStruct) (r *KefuDutyAddResponseStruct, err error) {
	kfdars := NewKefuDutyAddRequest(requestObj)
	kfdars = AssignDefaultValue(kfdars)

	b := kda.AddKefuDuty(kfdars)
	if b == false {
		Logger.Debugf("AddKefuDuty add failed")
		return &KefuDutyAddResponseStruct{
			Status: KE_FU_DUTY_ADD_FAILED,
			Msg:    Stat[KE_FU_DUTY_ADD_FAILED],
		}, nil
	}

	return &KefuDutyAddResponseStruct{
		Status: KE_FU_DUTY_ADD_SUCCESS,
		Msg:    Stat[KE_FU_DUTY_ADD_SUCCESS],
	}, nil
}

/**
 * [StartKeFuDutyAddServer 客服值班---新增值班服务]
 * @DateTime 2017-10-27T15:58:21+0800
 */
func StartKeFuDutyAddServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30059"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/KefuDutyAddThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &kefudutyaddservice{}
	processor := NewKefuDutyAddThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
