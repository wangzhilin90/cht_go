package kefudutylist

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	kdl "cht/models/kefudutylist"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	QUERY_KEFU_DUTY_LIST_SUCCESS          = 1000
	QUERY_KEFU_DUTY_LIST_TOTAL_NUM_FAILED = 1001
	QUERY_KEFU_DUTY_LIST_FAILED           = 1002
)

var Stat = map[int]string{
	QUERY_KEFU_DUTY_LIST_SUCCESS:          "客服值班列表查询成功",
	QUERY_KEFU_DUTY_LIST_TOTAL_NUM_FAILED: "客服值班列表总条目数查询失败",
	QUERY_KEFU_DUTY_LIST_FAILED:           "客服值班列表查询失败",
}

type kefudutylistservice struct{}

func (kdls *kefudutylistservice) GetKefuDutyList(requestObj *KefuDutyListRequestStruct) (r *KefuDutyListResponseStruct, err error) {
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*KefuDutyListRequestStruct)
	kdlr := new(kdl.KefuDutyListRequest)
	kdlr.StartTime = requestObj.GetStartTime()
	kdlr.EndTime = requestObj.GetEndTime()
	kdlr.Kefu = requestObj.GetKefu()
	kdlr.IsExport = requestObj.GetIsExport()
	kdlr.LimitOffset = requestObj.GetLimitOffset()
	kdlr.LimitNum = requestObj.GetLimitNum()
	kdlr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	totalNum, err := kdl.GetKefuDutyListTotalNum(kdlr)
	if err != nil {
		Logger.Errorf("GetKefuDutyList get total num failed:%v", err)
		return &KefuDutyListResponseStruct{
			Status: QUERY_KEFU_DUTY_LIST_TOTAL_NUM_FAILED,
			Msg:    Stat[QUERY_KEFU_DUTY_LIST_TOTAL_NUM_FAILED],
		}, nil
	}

	res, err := kdl.GetKefuDutyList(kdlr)
	if err != nil {
		Logger.Errorf("GetKefuDutyList get list failed:%v", err)
		return &KefuDutyListResponseStruct{
			Status: QUERY_KEFU_DUTY_LIST_FAILED,
			Msg:    Stat[QUERY_KEFU_DUTY_LIST_FAILED],
		}, nil
	}

	var response KefuDutyListResponseStruct
	for _, v := range res {
		kdlrs := new(KefuDutyListResultStruct)
		kdlrs.ID = v.ID
		kdlrs.Customer = v.Customer
		kdlrs.IsRest = v.IsRest
		kdlrs.DutyTime = v.DutyTime
		kdlrs.HolidayUser = v.HolidayUser
		kdlrs.Addtime = v.Addtime
		kdlrs.Starttime = v.Starttime
		kdlrs.Endtime = v.Endtime
		response.KefuDutyList = append(response.KefuDutyList, kdlrs)
	}
	response.TotalNum = totalNum
	response.Status = QUERY_KEFU_DUTY_LIST_SUCCESS
	response.Msg = Stat[QUERY_KEFU_DUTY_LIST_SUCCESS]
	Logger.Debugf("GetKefuDutyList response:%v", response)
	return &response, nil
}

/**
 * [StartKeFuDutyListServer 客服值班---列表服务]
 * @DateTime 2017-10-30T13:46:03+0800
 */
func StartKeFuDutyListServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30062"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/KefuDutyListThriftService/providers"
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

	handler := &kefudutylistservice{}
	processor := NewKefuDutyListThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
