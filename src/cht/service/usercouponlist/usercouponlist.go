package usercouponlist

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/rateroupon"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

type CouponService struct{}

func (cps *CouponService) GetUserCouponList(requestObj *UserCouponListRequestStruct) (r *UserCouponListResponseStruct, err error) {
	Logger.Info("GetUserCouponList requestObj:", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*UserCouponListRequestStruct)
	req := new(rateroupon.CouponRequest)
	req.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	req.Limit = requestObj.GetLimit()
	req.OrderBy = requestObj.GetOrderBy()
	req.Status = requestObj.GetStatus()
	req.UserID = requestObj.GetUserID()

	res, err := rateroupon.GetRateRoupon(req)
	if err != nil {
		Logger.Errorf("GetRateRoupon failed:%v", err)
		return nil, nil
	}
	var crs UserCouponListResponseStruct
	for _, v := range res {
		cs := new(UserCouponDetailsStruct)
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
		crs.UserCouponList = append(crs.UserCouponList, cs)
	}
	Logger.Debug(crs.UserCouponList)
	return &crs, nil
}

/**
 * [StartCouponServer 开启加息券服务API]
 * @DateTime 2017-08-24T15:19:45+0800
 */
func StartCouponServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30001"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/UserCouponListThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode %v failed %v", servicename, err)
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

	handler := &CouponService{}
	processor := NewUserCouponListThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
