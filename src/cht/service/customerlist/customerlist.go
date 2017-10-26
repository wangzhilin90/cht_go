package customerlist

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	cl "cht/models/customerlist"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	QUERY_CUSTOMER_LIST_SUCCESS         = 1000
	QUERY_CUSTOMER_LIST_TOTALNUM_FAILED = 1001
	QUERY_CUSTOMER_LIST_FAILED          = 1002
)

var Stat = map[int]string{
	QUERY_CUSTOMER_LIST_SUCCESS:         "查询专属客服列表成功",
	QUERY_CUSTOMER_LIST_TOTALNUM_FAILED: "查询专属客服列表总数失败",
	QUERY_CUSTOMER_LIST_FAILED:          "查询专属客服列表详情失败",
}

type customerlistservice struct{}

func (cls *customerlistservice) GetCustomerList(requestObj *CustomerListRequestStruct) (r *CustomerListResponseStruct, err error) {
	clr := new(cl.CustomerListRequest)
	clr.Customer = requestObj.GetCustomer()
	clr.StartTime = requestObj.GetStartTime()
	clr.EndTime = requestObj.GetEndTime()
	clr.Islock = requestObj.GetIslock()
	clr.Username = requestObj.GetUsername()
	clr.IsExport = requestObj.GetIsExport()
	clr.LimitOffset = requestObj.GetLimitOffset()
	clr.LimitNum = requestObj.GetLimitNum()
	clr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	totalNum, err := cl.GetCustomerListTotalNum(clr)
	if err != nil {
		Logger.Errorf("GetCustomerList get total num failed:%v", err)
		return &CustomerListResponseStruct{
			Status: QUERY_CUSTOMER_LIST_TOTALNUM_FAILED,
			Msg:    Stat[QUERY_CUSTOMER_LIST_TOTALNUM_FAILED],
		}, nil
	}

	res, err := cl.GetCustomerList(clr)
	if err != nil {
		Logger.Errorf("GetCustomerList get customer list failed:%v", err)
		return &CustomerListResponseStruct{
			Status: QUERY_CUSTOMER_LIST_FAILED,
			Msg:    Stat[QUERY_CUSTOMER_LIST_FAILED],
		}, nil
	}

	var response CustomerListResponseStruct
	for _, v := range res {
		clrs := new(CustomerListResultStruct)
		clrs.ID = v.ID
		clrs.Username = v.Username
		clrs.Password = v.Password
		clrs.Paypassword = v.Paypassword
		clrs.Point = v.Point
		clrs.Email = v.Email
		clrs.Avatar = v.Avatar
		clrs.Sex = v.Sex
		clrs.Realname = v.Realname
		clrs.Phone = v.Phone
		clrs.Tel = v.Tel
		clrs.Birthday = v.Birthday
		clrs.Nation = v.Nation
		clrs.Province = v.Province
		clrs.City = v.City
		clrs.Area = v.Area
		clrs.Address = v.Address
		clrs.Customer = v.Customer
		clrs.Logintime = v.Logintime
		clrs.Loginip = v.Loginip
		clrs.Addtime = v.Addtime
		clrs.Addip = v.Addip
		clrs.Islock = v.Islock
		clrs.Isvest = v.Isvest
		clrs.OsType = v.OsType
		clrs.DeviceToken = v.DeviceToken
		clrs.WeinxinID = v.WeinxinID
		clrs.BindTime = v.BindTime
		clrs.InvitationCode = v.InvitationCode
		clrs.Source = v.Source
		clrs.Hsid = v.Hsid
		clrs.GStatus = v.GStatus
		clrs.GPassword = v.GPassword
		clrs.AutoProtocolCode = v.AutoProtocolCode
		clrs.IsBorrower = v.IsBorrower
		clrs.IsWorker = v.IsWorker
		clrs.Hswaitactivate = v.Hswaitactivate
		response.CustomerList = append(response.CustomerList, clrs)
	}
	response.TotalNum = totalNum
	response.Status = QUERY_CUSTOMER_LIST_SUCCESS
	response.Msg = Stat[QUERY_CUSTOMER_LIST_SUCCESS]
	Logger.Debugf("GetCustomerList response:%v", response)
	return &response, nil
}

/**
 * [StartCustomerListServer 专属客服---列表服务]
 * @DateTime 2017-10-26T15:15:52+0800
 */
func StartCustomerListServer() {
	zkServers := []string{"192.168.8.208:2181"}
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30054"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/CustomerListThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &customerlistservice{}
	processor := NewCustomerListThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
