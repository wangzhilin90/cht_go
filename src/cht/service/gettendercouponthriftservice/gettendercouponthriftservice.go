package gettendercouponthriftservice

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/gettendercoupon"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type gettendercouponservice struct{}

const (
	QUERY_COUPON_FAILED  = 1001
	QUERY_COUPON_SUCCESS = 1002
)

var Status = map[int]string{
	QUERY_COUPON_FAILED:  "查询加息值出错",
	QUERY_COUPON_SUCCESS: "查询加息值成功",
}

func (gts *gettendercouponservice) GetCouponInfo(requestObj *TenderCouponRequestStruct) (r *TenderCouponResponseStruct, err error) {
	tcr := new(gettendercoupon.TenderCouponRequest)
	tcr.UserId = requestObj.GetUserId()
	tcr.TenderId = requestObj.GetTenderId()
	tcr.CouponId = requestObj.GetCouponId()
	tcr.TenderMoney = requestObj.GetTenderMoney()
	tcr.TimeLimit = requestObj.GetTimeLimit()
	tcr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, err := gettendercoupon.GetTenderCoupon(tcr)
	if err != nil {
		return &TenderCouponResponseStruct{
			Status: QUERY_COUPON_FAILED,
			Msg:    Status[QUERY_COUPON_SUCCESS],
		}, nil
	}

	Logger.Debug("GetCouponInfo res ", res)

	return &TenderCouponResponseStruct{
		Status: QUERY_COUPON_SUCCESS,
		Msg:    Status[QUERY_COUPON_SUCCESS],
		Coupon: res,
	}, nil
}

func StartGetCouponServer() {
	zkServers := []string{"192.168.8.212:2181", "192.168.8.213:2181", "192.168.8.214:2181"}
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30008"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/GetTenderCouponThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &gettendercouponservice{}
	processor := NewGetTenderCouponThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}