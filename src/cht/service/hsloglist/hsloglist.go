package hsloglist

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	hl "cht/models/hsloglist"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	QUERY_HS_LOG_SUCCESS          = 1000
	QUERY_HS_LOG_TOTAL_NUM_FAILED = 1001
	QUERY_HS_LOG_DETAILS_FAILED   = 1002
)

var Stat = map[int]string{
	QUERY_HS_LOG_SUCCESS:          "查询徽商日志成功",
	QUERY_HS_LOG_TOTAL_NUM_FAILED: "查询徽商总数失败",
	QUERY_HS_LOG_DETAILS_FAILED:   "查询徽商日志详情失败",
}

type hsloglistservice struct{}

func (hlls *hsloglistservice) GetHslogList(requestObj *HsLogListRequestStruct) (r *HsLogListReponseStruct, err error) {
	hlr := new(hl.HsLogListRequest)
	hlr.StartTime = requestObj.GetStartTime()
	hlr.EndTime = requestObj.GetEndTime()
	hlr.Type = requestObj.GetType()
	hlr.Type2 = requestObj.GetType2()
	hlr.Kws = requestObj.GetKws()
	hlr.Utype = requestObj.GetUtype()
	hlr.IsExport = requestObj.GetIsExport()
	hlr.LimitOffset = requestObj.GetLimitOffset()
	hlr.LimitNum = requestObj.GetLimitNum()
	hlr.BorrowID = requestObj.GetBorrowID()
	hlr.BorrowID = requestObj.GetBorrowID()
	hlr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	totalNum, err := hl.GetHsLogTotalNum(hlr)
	if err != nil {
		Logger.Errorf("GetHslogList query total num failed:%v", err)
		return &HsLogListReponseStruct{
			Status: QUERY_HS_LOG_TOTAL_NUM_FAILED,
			Msg:    Stat[QUERY_HS_LOG_TOTAL_NUM_FAILED],
		}, nil
	}

	res, err := hl.GetHsLog(hlr)
	if err != nil {
		Logger.Errorf("GetHslogList get hs log failed:%v", err)
		return &HsLogListReponseStruct{
			Status: QUERY_HS_LOG_DETAILS_FAILED,
			Msg:    Stat[QUERY_HS_LOG_DETAILS_FAILED],
		}, nil
	}

	var response HsLogListReponseStruct
	for _, v := range res {
		hlds := new(HsLogDetailsStruct)
		hlds.ID = v.ID
		hlds.UserID = v.UserID
		hlds.Orderno = v.Orderno
		hlds.Type = v.Type
		hlds.Money = v.Money
		hlds.FreezeMoney = v.FreezeMoney
		hlds.WaitMoney = v.WaitMoney
		hlds.Addtime = v.Addtime
		hlds.Toid = v.Toid
		hlds.Remark = v.Remark
		hlds.Username = v.Username
		hlds.Realname = v.Realname
		hlds.Regtime = v.Regtime
		response.HslogList = append(response.HslogList, hlds)
	}
	response.TotalNum = totalNum
	response.Status = QUERY_HS_LOG_SUCCESS
	response.Msg = Stat[QUERY_HS_LOG_SUCCESS]
	Logger.Debugf("GetHslogList response:%v", response)
	return &response, nil
}

/**
 * [StartHSLogListServer 徽商日志明细服务]
 * @DateTime 2017-10-20T16:44:01+0800
 */
func StartHSLogListServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30047"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/HsLogListThriftService/providers"
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

	handler := &hsloglistservice{}
	processor := NewHsLogListThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
