package couponlistthriftservice

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/rateroupon"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type CouponService struct{}

func (cps *CouponService) GetCoupon(requestObj *CouponRequestStruct) (r *CouponResponseStruct, err error) {
	Logger.Debugf("GetCoupon start")
	if requestObj == nil {
		Logger.Fatal("input param nil")
	}

	req := new(rateroupon.CouponRequest)
	req.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	req.Limit = requestObj.GetLimit()
	req.OrderBy = requestObj.GetOrderBy()
	req.Status = requestObj.GetStatus()
	req.UserID = requestObj.GetUserID()
	Logger.Debug("GetCoupon input param", req)

	res, err := rateroupon.GetRateRoupon(req)
	if err != nil {
		Logger.Fatalf("GetRateRoupon failed", err)
	}
	var crs CouponResponseStruct
	for _, v := range res {
		cs := new(CouponStruct)
		cs.ID = v.ID
		cs.UserID = v.UserID
		cs.Addtime = v.Addtime
		cs.StartTime = v.StartTime
		cs.EndTime = v.EndTime
		cs.UseTime = v.UseTime
		cs.Status = v.Status
		cs.TenderID = v.TenderID
		cs.Apr = v.Apr
		cs.AppAdd = v.AppAdd
		cs.MinTender = v.MinTender
		cs.MaxTender = v.MaxTender
		cs.TimeLimit = v.TimeLimit
		cs.BorrowType = v.BorrowType
		cs.Name = v.Name
		cs.Remark = v.Remark
		cs.ActivityName = v.ActivityName
		crs.CouponList = append(crs.CouponList, cs)
	}
	Logger.Debug(crs.CouponList)
	return &crs, nil
}

/**
 * [StartCouponServer 开启加息券服务API]
 * @DateTime 2017-08-24T15:19:45+0800
 */
func StartCouponServer() {
	zkServers := []string{"192.168.8.212:2181", "192.168.8.213:2181", "192.168.8.214:2181"}
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30001"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/CouponListThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &CouponService{}
	processor := NewCouponListThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
