package usercashrecordList

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/cashrecord"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type cashrecordservice struct{}

const (
	QUERY_CASHRECORD_FAILED  = 1001
	QUERY_CASHRECORD_SUCCESS = 1000
)

var Status = map[int]string{
	QUERY_CASHRECORD_FAILED:  "查询提现记录失败",
	QUERY_CASHRECORD_SUCCESS: "查询提现记录成功",
}

func (cs *cashrecordservice) GetUserCashRecordList(requestObj *UserCashRecordListRequestStruct) (r *UserCashRecordListResponseStruct, err error) {
	crrs := new(cashrecord.CashRecordRequestStruct)
	crrs.UserID = requestObj.GetUserID()
	crrs.StartTime = requestObj.GetStartTime()
	crrs.EndTime = requestObj.GetEndTime()
	crrs.QueryTime = requestObj.GetQueryTime()
	crrs.RechargeStatus = requestObj.GetRechargeStatus()
	crrs.LimitOffset = requestObj.GetLimitOffset()
	crrs.LimitNum = requestObj.GetLimitNum()
	crrs.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, CashStat, num, err := cashrecord.GetCashRecord(crrs)
	if err != nil {
		return &UserCashRecordListResponseStruct{
			Status: QUERY_CASHRECORD_FAILED,
			Msg:    Status[QUERY_CASHRECORD_FAILED],
		}, err
	}

	var response UserCashRecordListResponseStruct
	for _, v := range res {
		crs := new(UserCashRecordDetailsStruct)
		crs.ID = v.ID
		crs.UserID = v.UserID
		crs.OrderSn = v.OrderSn
		crs.Money = v.Money
		crs.Credited = v.Credited
		crs.Fee = v.Fee
		crs.UseReturnMoney = v.UseReturnMoney
		crs.UseFreeNum = v.UseFreeNum
		crs.Addtime = v.Addtime
		crs.Status = v.Status
		crs.PayWay = v.PayWay
		crs.DealTime = v.DealTime
		crs.FailResult = v.FailResult
		response.UserCashRecordList = append(response.UserCashRecordList, crs)
	}

	if CashStat != nil {
		css := new(UserCashStatsStruct)
		css.Fee = CashStat.Fee
		css.Money = CashStat.Money
		response.UserCashStruct = css
	}
	response.Status = QUERY_CASHRECORD_SUCCESS
	response.Msg = Status[QUERY_CASHRECORD_SUCCESS]
	response.Totalnum = num

	Logger.Debug("getUserCashRecordList res:", response)
	return &response, nil
}

/**
 * [StartUpdatePasswdsServer 开启充值提现，获取提现记录服务]
 * @DateTime 2017-08-24T15:19:45+0800
 */
func StartCashRecordServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30009"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/UserCashRecordListThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode %v failed", servicename, err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &cashrecordservice{}
	processor := NewUserCashRecordListThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
