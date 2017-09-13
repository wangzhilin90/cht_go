package rechargerecordthriftservice

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/rechargerecord"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type rechargerecordservice struct{}

const (
	QUERY_RECHAGE_FAILED  = 1001
	QUERY_RECHAGE_SUCCESS = 0
)

var Status = map[int]string{
	QUERY_RECHAGE_FAILED:  "查询充值记录失败",
	QUERY_RECHAGE_SUCCESS: "查询充值记录成功",
}

func (rrs *rechargerecordservice) GetRechargeRecord(requestObj *RechargeRecordRequestStruct) (r *RechargeRecordResponseStruct, err error) {
	rrr := new(rechargerecord.RechargeRecordRequest)
	rrr.UserID = requestObj.GetUserID()
	rrr.StartTime = requestObj.GetStartTime()
	rrr.EndTime = requestObj.GetEndTime()
	rrr.QueryTime = requestObj.GetQueryTime()
	rrr.RechargeStatus = requestObj.GetRechargeStatus()
	rrr.LimitOffset = requestObj.GetLimitNum()
	rrr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, num, money, err := rechargerecord.GetRechargeRecord(rrr)
	if err != nil {
		return &RechargeRecordResponseStruct{
			Status:               QUERY_RECHAGE_FAILED,
			Msg:                  Status[QUERY_RECHAGE_FAILED],
			Totalnum:             0,
			TotalHsRechargeMoney: "0.00",
		}, err
	}
	var rrrs RechargeRecordResponseStruct
	for _, v := range res {
		rrs := new(RechargeRecordStruct)
		rrs.ID = v.ID
		rrs.UserID = v.UserID
		rrs.OrderSn = v.OrderSn
		rrs.Money = v.Money
		rrs.Addtime = v.Addtime
		rrs.Status = v.Status
		rrs.DealTime = v.DealTime
		rrs.PayType = v.PayType
		rrs.PayWay = v.PayWay
		rrs.FailResult = v.FailResult
		rrrs.RechargeRecordList = append(rrrs.RechargeRecordList, rrs)
	}
	rrrs.Status = QUERY_RECHAGE_SUCCESS
	rrrs.Msg = Status[QUERY_RECHAGE_SUCCESS]
	rrrs.Totalnum = num
	rrrs.TotalHsRechargeMoney = money
	Logger.Debug("GetRechargeRecord res:", rrrs)
	return &rrrs, nil
}

/**
 * [StartUpdatePasswdsServer 开启查询充值记录服务]
 * @DateTime 2017-08-24T15:19:45+0800
 */
func StartRechargeRecordServer() {
	zkServers := []string{"192.168.8.208:2181"}
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30005"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/RechargeRecordThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &rechargerecordservice{}
	processor := NewRechargeRecordThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
