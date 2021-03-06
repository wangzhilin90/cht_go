package operationaldata

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/operationaldata"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	QUERY_OPERATIONAL_DATA_SUCCESS     = 1000
	QUERY_THIRTYDAYS_RESULT_FAILED     = 1001
	QUERY_TWELVE_MONTH_RESULT_FAILED   = 1002
	QUERY_ONE_MONTH_RESULT_FAILED      = 1003
	QUERY_PERIOD_RESULT_FAILED         = 1004
	QUERY_INVEST_RESULT_FAILED         = 1005
	QUERY_BID_RESULT_FAILED            = 1006
	QUERY_WAITRESULT_FAILED            = 1007
	QUERY_TWELVE_MONTH_TOTALNUM_FAILED = 1008
	QUERY_TOTALREPAYMENT_FAILED        = 1009
	QUERY_TENDER_FAILED                = 1010
	QUERY_TENDER_TODAY_FAILED          = 1011
	QUERY_PROFIT_FAILED                = 1012
	QUERY_USER_TENDER_ACCOUNT_FAILED   = 1013
)

var Stat = map[int]string{
	QUERY_OPERATIONAL_DATA_SUCCESS:     "查询运营数据成功",
	QUERY_THIRTYDAYS_RESULT_FAILED:     "查询最近30天投标排行失败",
	QUERY_TWELVE_MONTH_RESULT_FAILED:   "查询最近12个月每月成交量失败",
	QUERY_ONE_MONTH_RESULT_FAILED:      "查询最近1个月每月成交量失败",
	QUERY_PERIOD_RESULT_FAILED:         "查询借款周期占比失败",
	QUERY_INVEST_RESULT_FAILED:         "查询投资金额占比失败",
	QUERY_BID_RESULT_FAILED:            "查询标的比例失败",
	QUERY_WAITRESULT_FAILED:            "查询实时待收排行榜失败",
	QUERY_TWELVE_MONTH_TOTALNUM_FAILED: "查询12个月之前成交总量失败",
	QUERY_TOTALREPAYMENT_FAILED:        "查询目前累计成功还款失败",
	QUERY_TENDER_FAILED:                "获取今天前投标且已经审核通过的投资总额失败",
	QUERY_TENDER_TODAY_FAILED:          "获取昨日满标复审通过的标的总额失败",
	QUERY_PROFIT_FAILED:                "获取今天前已还利息总额失败",
	QUERY_USER_TENDER_ACCOUNT_FAILED:   "获取出借人数失败",
}

type operationaldataservice struct{}

func (ods *operationaldataservice) GetOperationalData(requestObj *OperationalDataRequestStruct) (r *OperationalDataResponseStruct, err error) {
	Logger.Infof("GetOperationalData requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*OperationalDataRequestStruct)
	odrs := new(operationaldata.OperationalDataRequestStruct)
	odrs.Start = requestObj.GetStart()
	odrs.StartMonth = requestObj.GetStartMonth()
	odrs.TodayTime = requestObj.GetTodayTime()
	odrs.YesterdayTime = requestObj.GetYesterdayTime()
	odrs.TomorrowTime = requestObj.GetTomorrowTime()
	odrs.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	var response OperationalDataResponseStruct
	/*得到最近30天投标排行结果*/
	tdr, err := operationaldata.GetThirtyDaysResult(odrs)
	if err != nil {
		return &OperationalDataResponseStruct{
			Status: QUERY_THIRTYDAYS_RESULT_FAILED,
			Msg:    Stat[QUERY_THIRTYDAYS_RESULT_FAILED],
		}, nil
	}

	for _, v := range tdr {
		tdrs := new(ThirtyDaysResultStruct)
		tdrs.Money = v.Money
		tdrs.Username = v.Username
		response.ThirtyDaysList = append(response.ThirtyDaysList, tdrs)
	}

	/*得到最近12个月每月成交量结果*/
	tmr, err := operationaldata.GetTwelveMonthResult(odrs)
	if err != nil {
		Logger.Errorf("GetOperationalData get twelve month failed:%v", err)
		return &OperationalDataResponseStruct{
			Status: QUERY_TWELVE_MONTH_RESULT_FAILED,
			Msg:    Stat[QUERY_TWELVE_MONTH_RESULT_FAILED],
		}, nil
	}

	for _, v := range tmr {
		tmrs := new(TwelveMonthResultStruct)
		tmrs.Category = v.Category
		tmrs.Account = v.Account
		response.TwelveMonthList = append(response.TwelveMonthList, tmrs)
	}

	/*得到最近1个月每月成交量结果*/
	omr, err := operationaldata.GetOneMonthResult(odrs)
	if err != nil {
		Logger.Errorf("GetOperationalData get one month failed:%v", err)
		return &OperationalDataResponseStruct{
			Status: QUERY_TWELVE_MONTH_RESULT_FAILED,
			Msg:    Stat[QUERY_TWELVE_MONTH_RESULT_FAILED],
		}, nil
	}

	for _, v := range omr {
		omrs := new(OneMonthResultStruct)
		omrs.Account = v.Account
		omrs.Category = v.Category
		response.OneMonthList = append(response.OneMonthList, omrs)
	}

	/*得到借款周期占比结果*/
	pr, err := operationaldata.GetPeriodResult(odrs)
	if err != nil {
		Logger.Errorf("GetOperationalData get period failed:%v", err)
		return &OperationalDataResponseStruct{
			Status: QUERY_PERIOD_RESULT_FAILED,
			Msg:    Stat[QUERY_PERIOD_RESULT_FAILED],
		}, nil
	}

	for _, v := range pr {
		prs := new(PeriodResultStruct)
		prs.Category = v.Category
		prs.Column_1 = v.Column_1
		response.PeriodList = append(response.PeriodList, prs)
	}

	/*得到投资金额占比结果*/
	ir, err := operationaldata.GetInvestResult(odrs)
	if err != nil {
		Logger.Errorf("GetOperationalData get invest failed:%v", err)
		return &OperationalDataResponseStruct{
			Status: QUERY_INVEST_RESULT_FAILED,
			Msg:    Stat[QUERY_INVEST_RESULT_FAILED],
		}, nil
	}

	if ir != nil {
		irs := new(InvestResultStruct)
		irs.A1 = ir.A1
		irs.A2 = ir.A2
		irs.A3 = ir.A3
		irs.A4 = ir.A4
		irs.A5 = ir.A5
		response.InvestAccount = irs
	}

	/*得到标的比例结果*/
	br, err := operationaldata.GetBidResult(odrs)
	if err != nil {
		Logger.Errorf("GetOperationalData git bid failed:%v", err)
		return &OperationalDataResponseStruct{
			Status: QUERY_BID_RESULT_FAILED,
			Msg:    Stat[QUERY_BID_RESULT_FAILED],
		}, nil
	}

	for _, v := range br {
		brs := new(BidResultStruct)
		brs.BorrowType = v.BorrowType
		brs.Number = v.Number
		response.BidList = append(response.BidList, brs)
	}

	/*得到实时待收排行榜结果*/
	wr, err := operationaldata.GetWaitResult(odrs)
	if err != nil {
		Logger.Errorf("GetOperationalData get wait failed:%v", err)
		return &OperationalDataResponseStruct{
			Status: QUERY_WAITRESULT_FAILED,
			Msg:    Stat[QUERY_WAITRESULT_FAILED],
		}, nil
	}

	for _, v := range wr {
		wrs := new(WaitResultStruct)
		wrs.Money = v.Money
		wrs.Username = v.Username
		response.WaitList = append(response.WaitList, wrs)
	}

	/*得到12个月之前成交总量*/
	tmtn, err := operationaldata.GetTwelveMonthTotalNum(odrs)
	if err != nil {
		Logger.Errorf("GetOperationalData get twelve month total num failed", err)
		return &OperationalDataResponseStruct{
			Status: QUERY_TWELVE_MONTH_TOTALNUM_FAILED,
			Msg:    Stat[QUERY_TWELVE_MONTH_TOTALNUM_FAILED],
		}, nil
	}
	response.OldSum = tmtn

	/*得到目前累计成功还款金额*/
	tr, err := operationaldata.GetTotalRepayment(odrs)
	if err != nil {
		Logger.Errorf("GetOperationalData get total repatyment failed:%v", err)
		return &OperationalDataResponseStruct{
			Status: QUERY_TOTALREPAYMENT_FAILED,
			Msg:    Stat[QUERY_TOTALREPAYMENT_FAILED],
		}, nil
	}
	response.Repayment = tr

	tender, err := operationaldata.GetTender(odrs)
	if err != nil {
		Logger.Errorf("GetOperationalData get tender failed:%v", err)
		return &OperationalDataResponseStruct{
			Status: QUERY_TENDER_FAILED,
			Msg:    Stat[QUERY_TENDER_FAILED],
		}, nil
	}

	tenderToday, err := operationaldata.GetTenderToday(odrs)
	if err != nil {
		Logger.Errorf("GetOperationalData get tender today failed:%v", err)
		return &OperationalDataResponseStruct{
			Status: QUERY_TENDER_TODAY_FAILED,
			Msg:    Stat[QUERY_TENDER_TODAY_FAILED],
		}, nil
	}

	profit, err := operationaldata.GetProfit(odrs)
	if err != nil {
		Logger.Errorf("GetOperationalData get profit failed:%v", err)
		return &OperationalDataResponseStruct{
			Status: QUERY_PROFIT_FAILED,
			Msg:    Stat[QUERY_PROFIT_FAILED],
		}, nil
	}

	tenderUserCount, err := operationaldata.GetTenderUserCount(odrs)
	if err != nil {
		Logger.Errorf("GetOperationalData get tenderUserCount failed:%v", err)
		return &OperationalDataResponseStruct{
			Status: QUERY_USER_TENDER_ACCOUNT_FAILED,
			Msg:    Stat[QUERY_USER_TENDER_ACCOUNT_FAILED],
		}, nil
	}

	srs := new(SumResultStruct)
	srs.Tender = tender
	srs.TenderToday = tenderToday
	srs.Profit = profit
	response.Sum = srs

	response.TenderUserCount = tenderUserCount
	response.Status = QUERY_OPERATIONAL_DATA_SUCCESS
	response.Msg = Stat[QUERY_OPERATIONAL_DATA_SUCCESS]
	Logger.Debugf("GetOperationalData return value:%v", response)
	return &response, nil
}

/**
 * [StartUpdatePasswdsServer 开启[后台]运营数据服务]
 */
func StartOperationalDataServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30029"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/OperationalDataThriftService/providers"
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

	handler := &operationaldataservice{}
	processor := NewOperationalDataThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
