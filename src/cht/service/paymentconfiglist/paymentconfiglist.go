package paymentconfiglist

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	pcl "cht/models/paymentconfiglist"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	QUERY_PAYMENT_CONFIG_LIST_SUCCESS = 1000
	QUERY_PAYMENT_CONFIG_LIST_FAILED  = 1001
)

var Stat = map[int]string{
	QUERY_PAYMENT_CONFIG_LIST_SUCCESS: "查询第三方支付列表成功",
	QUERY_PAYMENT_CONFIG_LIST_FAILED:  "查询第三方支付列表失败",
}

type paymentconfiglistservice struct{}

func (pcls *paymentconfiglistservice) GetPaymentConfigList(requestObj *PaymentConfigListRequestStruct) (r *PaymentConfigListResponseStruct, err error) {
	Logger.Infof("GetPaymentConfigList requestObj:%v", requestObj)
	pclr := new(pcl.PaymentConfigListRequest)
	pclr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, err := pcl.GetPaymentConfigList(pclr)
	if err != nil {
		Logger.Errorf("GetPaymentConfigList get config list failed:%v", err)
		return &PaymentConfigListResponseStruct{
			Status: QUERY_PAYMENT_CONFIG_LIST_FAILED,
			Msg:    Stat[QUERY_PAYMENT_CONFIG_LIST_FAILED],
		}, nil
	}

	var response PaymentConfigListResponseStruct
	for _, v := range res {
		pcds := new(PaymentConfigDetailsStruct)
		pcds.ID = v.ID
		pcds.Type = v.Type
		pcds.Nid = v.Nid
		pcds.Name = v.Name
		pcds.Logo = v.Logo
		pcds.Config = v.Config
		pcds.Fee = v.Fee
		pcds.Status = v.Status
		pcds.Remark = v.Remark
		pcds.Sort = v.Sort
		response.PaymentConfigList = append(response.PaymentConfigList, pcds)
	}
	response.Status = QUERY_PAYMENT_CONFIG_LIST_SUCCESS
	response.Msg = Stat[QUERY_PAYMENT_CONFIG_LIST_SUCCESS]
	Logger.Debugf("GetPaymentConfigList response:%v", response)
	return &response, nil
}

/**
 * [StartOperationalDataServer 第三方支付方式列表服务]
 * @DateTime 2017-10-27T14:30:19+0800
 */
func StartPaymentConfigListServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30058"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/PaymentConfigListThriftService/providers"
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

	handler := &paymentconfiglistservice{}
	processor := NewPaymentConfigListThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
