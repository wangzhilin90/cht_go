package hscashlist

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	hcl "cht/models/hscashlist"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	QUERY_HS_CASH_LIST_SUCCESS     = 1000
	QUERY_HS_CASH_TOTAL_NUM_FAILED = 1001
	QUERY_HS_CASH_LIST_FAILED      = 1002
)

var Stat = map[int]string{
	QUERY_HS_CASH_LIST_SUCCESS:     "获取徽商提现记录成功",
	QUERY_HS_CASH_TOTAL_NUM_FAILED: "获取徽商提现记录总数失败",
	QUERY_HS_CASH_LIST_FAILED:      "获取徽商提现记录列表失败",
}

type hscashlistservice struct{}

/**
 * [GetHsCashList [后台]获取徽商提现记录]
 * @param    GetHsCashList 请求入参
 * @return   HsCashListResponseStruct 返回徽商提现记录
 * @DateTime 2017-10-23T11:20:59+0800
 */
func (hcls *hscashlistservice) GetHsCashList(requestObj *HsCashListRequestStruct) (r *HsCashListResponseStruct, err error) {
	Logger.Infof("GetHsCashList requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*HsCashListRequestStruct)
	hclr := new(hcl.HsCashListRequest)
	hclr.StartTime = requestObj.GetStartTime()
	hclr.EndTime = requestObj.GetEndTime()
	hclr.Timetype = requestObj.GetTimetype()
	hclr.Utype = requestObj.GetUtype()
	hclr.Type = requestObj.GetType()
	hclr.Keywords = requestObj.GetKeywords()
	hclr.PayWay = requestObj.GetPayWay()
	hclr.Status = requestObj.GetStatus()
	hclr.IsExport = requestObj.GetIsExport()
	hclr.LimitOffset = requestObj.GetLimitOffset()
	hclr.LimitNum = requestObj.GetLimitNum()
	hclr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	totalNum, err := hcl.GetHsCashListTotalNum(hclr)
	if err != nil {
		Logger.Errorf("GetHsCashList get total num failed:%v", err)
		return &HsCashListResponseStruct{
			Status: QUERY_HS_CASH_TOTAL_NUM_FAILED,
			Msg:    Stat[QUERY_HS_CASH_TOTAL_NUM_FAILED],
		}, nil
	}

	res, err := hcl.GetHsCashList(hclr)
	if err != nil {
		Logger.Errorf("GetHsCashList get cash list failed:%v", err)
		return &HsCashListResponseStruct{
			Status: QUERY_HS_CASH_LIST_FAILED,
			Msg:    Stat[QUERY_HS_CASH_LIST_FAILED],
		}, nil
	}

	var response HsCashListResponseStruct
	for _, v := range res {
		hclrs := new(HsCashListResultStruct)
		hclrs.ID = v.ID
		hclrs.UserID = v.UserID
		hclrs.OrderSn = v.OrderSn
		hclrs.Money = v.Money
		hclrs.Credited = v.Credited
		hclrs.Fee = v.Fee
		hclrs.UseReturnMoney = v.UseReturnMoney
		hclrs.UseFreeNum = v.UseFreeNum
		hclrs.Addtime = v.Addtime
		hclrs.Status = v.Status
		hclrs.PayWay = v.PayWay
		hclrs.DealTime = v.DealTime
		hclrs.DealTime = v.DealTime
		hclrs.FailResult_ = v.FailResult_
		hclrs.Username = v.Username
		hclrs.Realname = v.Realname
		hclrs.Regtime = v.Regtime
		response.HsCashList = append(response.HsCashList, hclrs)
	}
	response.TotalNum = totalNum
	response.Status = QUERY_HS_CASH_LIST_SUCCESS
	response.Msg = Stat[QUERY_HS_CASH_LIST_SUCCESS]
	Logger.Debugf("GetHsCashList response:%v", response)
	return &response, nil
}

/**
 * [StartHSCashListServer [后台]徽商提现记录服务]
 * @DateTime 2017-10-23T11:34:48+0800
 */
func StartHSCashListServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30040"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/HsCashListThriftService/providers"
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

	handler := &hscashlistservice{}
	processor := NewHsCashListThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
