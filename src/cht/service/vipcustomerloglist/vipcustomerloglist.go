package vipcustomerloglist

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	vcll "cht/models/vipcustomerloglist"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	QUERY_VIP_CUSTOMER_LOG_LIST_SUCCESS          = 1000
	QUERY_VIP_CUSTOMER_LOG_LIST_TOTAL_NUM_FAILED = 1001
	QUERY_VIP_CUSTOMER_LOG_LIST_FAILED           = 1002
)

var Stat = map[int]string{
	QUERY_VIP_CUSTOMER_LOG_LIST_SUCCESS:          "专属客服日志记录查询成功",
	QUERY_VIP_CUSTOMER_LOG_LIST_TOTAL_NUM_FAILED: "专属客服日志记录总数查询失败",
	QUERY_VIP_CUSTOMER_LOG_LIST_FAILED:           "专属客服日志记录列表查询失败",
}

type vipcustomerloglistservice struct{}

func (vclls *vipcustomerloglistservice) GetVipCustomerLogList(requestObj *VipCustomerLogListRequestStruct) (r *VipCustomerLogListResponseStruct, err error) {
	vcllr := new(vcll.VipCustomerLogListRequest)
	vcllr.StartTime = requestObj.GetStartTime()
	vcllr.EndTime = requestObj.GetEndTime()
	vcllr.Keywords = requestObj.GetKeywords()
	vcllr.Type = requestObj.GetType()
	vcllr.LimitOffset = requestObj.GetLimitOffset()
	vcllr.LimitNum = requestObj.GetLimitNum()
	vcllr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	totalNum, err := vcll.GetVipCustomerLogListTotalNum(vcllr)
	if err != nil {
		return &VipCustomerLogListResponseStruct{
			Status: QUERY_VIP_CUSTOMER_LOG_LIST_TOTAL_NUM_FAILED,
			Msg:    Stat[QUERY_VIP_CUSTOMER_LOG_LIST_TOTAL_NUM_FAILED],
		}, nil
	}

	res, err := vcll.GetVipCustomerLogList(vcllr)
	if err != nil {
		return &VipCustomerLogListResponseStruct{
			Status: QUERY_VIP_CUSTOMER_LOG_LIST_FAILED,
			Msg:    Stat[QUERY_VIP_CUSTOMER_LOG_LIST_FAILED],
		}, nil
	}

	var response VipCustomerLogListResponseStruct
	for _, v := range res {
		vcds := new(VipCustomerDetailsStruct)
		vcds.ID = v.ID
		vcds.UserID = v.UserID
		vcds.Username = v.Username
		vcds.Email = v.Email
		vcds.Realname = v.Realname
		vcds.Phone = v.Phone
		vcds.ScenePasstime = v.ScenePasstime
		vcds.VipStatus = v.VipStatus
		vcds.VipPasstime = v.VipPasstime
		vcds.VipVerifytime = v.VipVerifytime
		vcds.OldCustomer = v.OldCustomer
		vcds.NewCustomer_ = v.NewCustomer_
		vcds.Updatetime = v.Updatetime
		vcds.Remark = v.Remark
		response.VipCustomerLogList = append(response.VipCustomerLogList, vcds)
	}
	response.TotalNum = totalNum
	response.Status = QUERY_VIP_CUSTOMER_LOG_LIST_SUCCESS
	response.Msg = Stat[QUERY_VIP_CUSTOMER_LOG_LIST_SUCCESS]
	Logger.Debugf("GetVipCustomerLogList response:%v", response)
	return &response, nil
}

/**
 * [StartVipCustomerLogListServer 专属客服日志记录服务]
 * @DateTime 2017-10-26T17:45:08+0800
 */
func StartVipCustomerLogListServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30056"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/VipCustomerLogListThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &vipcustomerloglistservice{}
	processor := NewVipCustomerLogListThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
