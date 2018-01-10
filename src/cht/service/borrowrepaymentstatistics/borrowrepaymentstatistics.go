package borrowrepaymentstatistics

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	brs "cht/models/borrowrepaymentstatistics"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

type borrowrepaymentstatisticsservice struct{}

const (
	QUERY_REPAYMENT_SUCCESS = 1000
	QUERY_REPAYMENT_FAILED  = 1001
)

var Borrow_Stat = map[int]string{
	QUERY_REPAYMENT_SUCCESS: "查询标分期还款记录表成功",
	QUERY_REPAYMENT_FAILED:  "查询标分期还款记录表失败",
}

const (
	QUERY_TOTAL_REPLAY_SUCCESS = 1000
	QUERY_TOTAL_REPLAY_FAILED  = 1001
)

var Total_Stat = map[int]string{
	QUERY_TOTAL_REPLAY_SUCCESS: "获取总还款金额成功",
	QUERY_TOTAL_REPLAY_FAILED:  "获取总还款金额失败",
}

func (brss *borrowrepaymentstatisticsservice) GetRepaymentStatisticsList(requestObj *RepaymentStatisticsRequestStruct) (r *RepaymentStatisticsResponseStruct, err error) {
	Logger.Infof("GetRepaymentStatisticsList requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*RepaymentStatisticsRequestStruct)
	rsr := new(brs.RepaymentStatisticsRequest)
	rsr.UserID = requestObj.GetUserID()
	rsr.Status = requestObj.GetStatus()
	rsr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, err := brs.GetRepaymentStatisticsList(rsr)
	if err != nil {
		Logger.Errorf("GetRepaymentStatisticsList query failed:%v", res)
		return &RepaymentStatisticsResponseStruct{
			Status: QUERY_REPAYMENT_FAILED,
			Msg:    Borrow_Stat[QUERY_REPAYMENT_FAILED],
		}, nil
	}

	var response RepaymentStatisticsResponseStruct
	for _, v := range res {
		rsds := new(RepaymentStatisticsDetailsStruct)
		rsds.BorrowID = v.BorrowID
		rsds.WillMoney = v.WillMoney
		rsds.ReplaymentMoney = v.ReplaymentMoney
		rsds.NoreplaymentMoney = v.NoreplaymentMoney
		response.RepaymentStatisticsList = append(response.RepaymentStatisticsList, rsds)
	}
	response.Status = QUERY_REPAYMENT_SUCCESS
	response.Msg = Borrow_Stat[QUERY_REPAYMENT_SUCCESS]
	Logger.Info("GetRepaymentStatisticsList response:%v", response)
	return &response, nil
}

func (brss *borrowrepaymentstatisticsservice) GetTotalReplaymentMoney(requestObj *RepaymentStatisticsRequestStruct) (r *TotalReplaymentMoneyResponseStruct, err error) {
	Logger.Info("GetTotalReplaymentMoney requestObj:", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*RepaymentStatisticsRequestStruct)
	rsr := new(brs.RepaymentStatisticsRequest)
	rsr.UserID = requestObj.GetUserID()
	rsr.Status = requestObj.GetStatus()
	rsr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, err := brs.GetTotalReplaymentMoney(rsr)
	if err != nil {
		Logger.Errorf("GetTotalReplaymentMoney query failed:%v", res)
		return &TotalReplaymentMoneyResponseStruct{
			Status: QUERY_TOTAL_REPLAY_FAILED,
			Msg:    Total_Stat[QUERY_TOTAL_REPLAY_FAILED],
		}, nil
	}

	var response TotalReplaymentMoneyResponseStruct
	response.TotalReplaymentMoney = res
	response.Status = QUERY_TOTAL_REPLAY_SUCCESS
	response.Msg = Total_Stat[QUERY_TOTAL_REPLAY_SUCCESS]
	Logger.Debugf("GetTotalReplaymentMoney response:%v", response)
	return &response, nil
}

/**
 * [StartBorrowRepaymentStatisticsServer 开启标分期还款记录表服务]
 * @DateTime 2017-12-11T12:06:31+0800
 */
func StartBorrowRepaymentStatisticsServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30067"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/BorrowRepaymentStatisticsThriftService/providers"
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

	handler := &borrowrepaymentstatisticsservice{}
	processor := NewBorrowRepaymentStatisticsThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
