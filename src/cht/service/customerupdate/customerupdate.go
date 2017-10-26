package customerupdate

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	cu "cht/models/customerupdate"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	UPDATE_CUSTOMER_LOCK_SUCCESS = 1000
	UPDATE_CUSTOMER_LOCK_FAILED  = 1001
)

var Stat = map[int]string{
	UPDATE_CUSTOMER_LOCK_SUCCESS: "更新用户是否锁定成功",
	UPDATE_CUSTOMER_LOCK_FAILED:  "更新用户是否锁定失败",
}

type customerupdateservice struct{}

func (cus *customerupdateservice) UpdateCustomer(requestObj *CustomerUpdateRequestStruct) (r *CustomerUpdateResponseStruct, err error) {
	cur := new(cu.CustomerUpdateRequest)
	cur.ID = requestObj.GetID()
	cur.Islock = requestObj.GetIslock()
	cur.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	b := cu.UpdateCustomer(cur)
	if b == false {
		return &CustomerUpdateResponseStruct{
			Status: UPDATE_CUSTOMER_LOCK_FAILED,
			Msg:    Stat[UPDATE_CUSTOMER_LOCK_FAILED],
		}, nil
	}

	return &CustomerUpdateResponseStruct{
		Status: UPDATE_CUSTOMER_LOCK_SUCCESS,
		Msg:    Stat[UPDATE_CUSTOMER_LOCK_SUCCESS],
	}, nil
}

/**
 * [StartCustomerUpdateServer 专属客服---会员启用、禁用服务]
 * @DateTime 2017-10-26T16:35:32+0800
 */
func StartCustomerUpdateServer() {
	zkServers := []string{"192.168.8.208:2181"}
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30055"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/CustomerUpdateThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &customerupdateservice{}
	processor := NewCustomerUpdateThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
