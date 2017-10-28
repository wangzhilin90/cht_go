package memberhelperlist

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	mhl "cht/models/memberhelperlist"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	QUERY_MEMBER_HELP_LIST_SUCCESS          = 1000
	QUERY_MEMBER_HELP_LIST_TOTAL_NUM_FAILED = 1001
	QUERY_MEMBER_HELP_LIST_FAILED           = 1002
)

var Stat = map[int]string{
	QUERY_MEMBER_HELP_LIST_SUCCESS:          "查询会员紧急联系人成功",
	QUERY_MEMBER_HELP_LIST_TOTAL_NUM_FAILED: "查询会员紧急联系人总记录数失败",
	QUERY_MEMBER_HELP_LIST_FAILED:           "查询会员紧急联系人列表失败",
}

type memberhelperlistservice struct{}

func (mhls *memberhelperlistservice) GetMemberHelperList(requestObj *MemberHelperListRequestStruct) (r *MemberHelperListResponseStruct, err error) {
	mhlr := new(mhl.MemberHelperListRequest)
	mhlr.Type = requestObj.GetType()
	mhlr.Keywords = requestObj.GetKeywords()
	mhlr.LimitNum = requestObj.GetLimitNum()
	mhlr.LimitOffset = requestObj.GetLimitOffset()
	mhlr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	totalNum, err := mhl.GetMemberHelperListTotalNum(mhlr)
	if err != nil {
		Logger.Errorf("GetMemberHelperList query failed:%v", err)
		return &MemberHelperListResponseStruct{
			Status: QUERY_MEMBER_HELP_LIST_TOTAL_NUM_FAILED,
			Msg:    Stat[QUERY_MEMBER_HELP_LIST_TOTAL_NUM_FAILED],
		}, nil
	}

	res, err := mhl.GetMemberHelperList(mhlr)
	if err != nil {
		Logger.Errorf("GetMemberHelperList query help list failed:%v", err)
		return &MemberHelperListResponseStruct{
			Status: QUERY_MEMBER_HELP_LIST_FAILED,
			Msg:    Stat[QUERY_MEMBER_HELP_LIST_FAILED],
		}, nil
	}

	var response MemberHelperListResponseStruct
	for _, v := range res {
		mhds := new(MemberHelperDetailsStruct)
		mhds.Linkman = v.Linkman
		mhds.Linkrelation = v.Linkrelation
		mhds.Linkphone = v.Linkphone
		mhds.Updatetime = v.Updatetime
		mhds.ID = v.ID
		mhds.Username = v.Username
		mhds.Realname = v.Realname
		mhds.Phone = v.Phone
		mhds.Customer = v.Customer
		response.LinkManList = append(response.LinkManList, mhds)
	}
	response.TotalNum = totalNum
	response.Status = QUERY_MEMBER_HELP_LIST_SUCCESS
	response.Msg = Stat[QUERY_MEMBER_HELP_LIST_SUCCESS]
	response.TotalNum = totalNum
	Logger.Debugf("GetMemberHelperList response:%v", response)
	return &response, nil
}

/**
 * [StartMemberHelperListServer 客户管理---会员紧急联系人服务]
 * @DateTime 2017-10-23T14:34:48+0800
 */
func StartMemberHelperListServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30041"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/MemberHelperListThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode %v failed", servicename, err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &memberhelperlistservice{}
	processor := NewMemberHelperListThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
