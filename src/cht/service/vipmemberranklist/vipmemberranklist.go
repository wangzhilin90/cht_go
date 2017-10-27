package vipmemberranklist

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	vmrl "cht/models/vipmemberranklist"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	QUERY_VIPMEMBERRANKLIST_SUCCESS          = 1000
	QUERY_VIPMEMBERRANKLIST_TOTAL_NUM_FAILED = 1001
	QUERY_VIPMEMBERRANKLIST_FAILED           = 1002
)

var Stat = map[int]string{
	QUERY_VIPMEMBERRANKLIST_SUCCESS:          "查询VIP会员等级成功",
	QUERY_VIPMEMBERRANKLIST_TOTAL_NUM_FAILED: "查询VIP会员等级总条目数失败",
	QUERY_VIPMEMBERRANKLIST_FAILED:           "查询VIP会员等级列表失败",
}

type vipmemberranklistservice struct{}

func (vmrls *vipmemberranklistservice) GetVipMemberRankList(requestObj *VipMemberRankListRequestStruct) (r *VipMemberRankListReponseStruct, err error) {
	vmrlr := new(vmrl.VipMemberRankListRequest)
	vmrlr.Type = requestObj.GetType()
	vmrlr.Keywords = requestObj.GetKeywords()
	vmrlr.LimitOffset = requestObj.GetLimitOffset()
	vmrlr.LimitNum = requestObj.GetLimitNum()
	vmrlr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	totalNum, err := vmrl.GetVipMemberRankListTotalNum(vmrlr)
	if err != nil {
		Logger.Errorf("GetVipMemberRankList get total num failed:%v", err)
		return &VipMemberRankListReponseStruct{
			Status: QUERY_VIPMEMBERRANKLIST_TOTAL_NUM_FAILED,
			Msg:    Stat[QUERY_VIPMEMBERRANKLIST_TOTAL_NUM_FAILED],
		}, nil
	}

	res, err := vmrl.GetVipMemberRankList(vmrlr)
	if err != nil {
		Logger.Errorf("GetVipMemberRankList get vip list failed:%v", err)
		return &VipMemberRankListReponseStruct{
			Status: QUERY_VIPMEMBERRANKLIST_FAILED,
			Msg:    Stat[QUERY_VIPMEMBERRANKLIST_FAILED],
		}, nil
	}

	var response VipMemberRankListReponseStruct
	for _, v := range res {
		vmrds := new(VipMemberRankDetailsStruct)
		vmrds.UserID = v.UserID
		vmrds.MoneyFreeze = v.MoneyFreeze
		vmrds.MoneyUsable = v.MoneyUsable
		vmrds.CashTime = v.CashTime
		vmrds.ReturnMoney = v.ReturnMoney
		vmrds.VipLevel = v.VipLevel
		vmrds.VipWaitMoney = v.VipWaitMoney
		vmrds.HsreturnMoney = v.HsreturnMoney
		vmrds.HsmoneyFreeze = v.HsmoneyFreeze
		vmrds.HsmoneyWait = v.HsmoneyWait
		vmrds.Username = v.Username
		vmrds.Realname = v.Realname
		vmrds.Addtime = v.Addtime
		response.VipMemberRankList = append(response.VipMemberRankList, vmrds)
	}

	response.TotalNum = totalNum
	response.Status = QUERY_VIPMEMBERRANKLIST_SUCCESS
	response.Msg = Stat[QUERY_VIPMEMBERRANKLIST_SUCCESS]
	Logger.Debugf("GetVipMemberRankList response:%v", response)
	return &response, nil
}

/**
 * [StartVipMemberRankListServer VIP会员等级服务]
 * @DateTime 2017-10-27T14:02:02+0800
 */
func StartVipMemberRankListServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30057"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/VipMemberRankListThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &vipmemberranklistservice{}
	processor := NewVipMemberRankListThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
