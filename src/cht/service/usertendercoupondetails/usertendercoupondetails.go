package usertendercoupondetails

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/gettendercoupon"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

type gettendercouponservice struct{}

const (
	QUERY_COUPON_FAILED  = 1001
	QUERY_COUPON_SUCCESS = 1000
)

var Status = map[int]string{
	QUERY_COUPON_FAILED:  "查询加息值出错",
	QUERY_COUPON_SUCCESS: "查询加息值成功",
}

func (gts *gettendercouponservice) GetUserTenderCouponDetails(requestObj *UserTenderCouponDetailsRequestStruct) (r *UserTenderCouponDetailsResponseStruct, err error) {
	tcr := new(gettendercoupon.TenderCouponRequest)
	tcr.UserId = requestObj.GetUserId()
	tcr.TenderId = requestObj.GetTenderId()
	tcr.CouponId = requestObj.GetCouponId()
	tcr.TenderMoney = requestObj.GetTenderMoney()
	tcr.TimeLimit = requestObj.GetTimeLimit()
	tcr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, err := gettendercoupon.GetTenderCoupon(tcr)
	if err != nil {
		Logger.Errorf("GetUserTenderCouponDetails query failed:%v", err)
		return &UserTenderCouponDetailsResponseStruct{
			Status: QUERY_COUPON_FAILED,
			Msg:    Status[QUERY_COUPON_SUCCESS],
		}, nil
	}

	Logger.Debug("GetUserTenderCouponDetails res ", res)

	return &UserTenderCouponDetailsResponseStruct{
		Status: QUERY_COUPON_SUCCESS,
		Msg:    Status[QUERY_COUPON_SUCCESS],
		Coupon: res,
	}, nil
}

/*开启立即投资，获取投标加息值服务*/
func StartGetCouponServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30008"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/UserTenderCouponDetailsThriftService/providers"
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

	handler := &gettendercouponservice{}
	processor := NewUserTenderCouponDetailsThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
