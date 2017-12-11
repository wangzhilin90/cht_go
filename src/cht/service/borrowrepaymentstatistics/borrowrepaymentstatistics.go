package borrowrepaymentstatistics

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	brs "cht/models/borrowrepaymentstatistics"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
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

func (brss *borrowrepaymentstatisticsservice) GetRepaymentStatisticsDetails(requestObj *RepaymentStatisticsRequestStruct) (r *RepaymentStatisticsResponseStruct, err error) {
	Logger.Info("GetRepaymentStatisticsDetails request param:", requestObj)
	rsr := new(brs.RepaymentStatisticsRequest)
	rsr.UserID = requestObj.GetUserID()
	rsr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, err := brs.GetRepaymentStatisticsDetails(rsr)
	if err != nil {
		Logger.Errorf("GetRepaymentStatisticsDetails query failed:%v", res)
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
	Logger.Info("GetRepaymentStatisticsDetails response:%v", response)
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

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &borrowrepaymentstatisticsservice{}
	processor := NewBorrowRepaymentStatisticsThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
