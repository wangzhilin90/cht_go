package goodsdetails

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	gd "cht/models/goodsdetails"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	QUERY_GOODS_DETAILS_SUCCESS = 1000
	QUERY_GOODS_DETAILS_FAILED  = 1001
	QUERY_GOODS_DETAILS_EMPTY   = 1002
)

var Stat = map[int]string{
	QUERY_GOODS_DETAILS_SUCCESS: "获取商品详情成功",
	QUERY_GOODS_DETAILS_FAILED:  "获取商品详情失败",
	QUERY_GOODS_DETAILS_EMPTY:   "获取商品详情为空",
}

type goodsdetailsservice struct{}

func (gds *goodsdetailsservice) GetGoodsDetails(requestObj *GoodsDetailsRequestStruct) (r *GoodsDetailsResponseStruct, err error) {
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*GoodsDetailsRequestStruct)
	gdr := new(gd.GoodsDetailsRequest)
	gdr.ID = requestObj.GetID()
	gdr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	res, err := gd.GetGoodsDetails(gdr)
	if err != nil {
		Logger.Errorf("GetGoodsDetails query failed:v", err)
		return &GoodsDetailsResponseStruct{
			Status: QUERY_GOODS_DETAILS_FAILED,
			Msg:    Stat[QUERY_GOODS_DETAILS_FAILED],
		}, nil
	}

	if res == nil {
		Logger.Debugf("GetGoodsDetails query empty")
		return &GoodsDetailsResponseStruct{
			Status: QUERY_GOODS_DETAILS_EMPTY,
			Msg:    Stat[QUERY_GOODS_DETAILS_EMPTY],
		}, nil
	}

	var response GoodsDetailsResponseStruct
	if res != nil {
		gdss := new(GoodsDetailsStruct)
		gdss.ID = res.ID
		gdss.Addtime = res.Addtime
		gdss.ShowTime = res.ShowTime
		gdss.CloseTime = res.CloseTime
		gdss.IsTimer = res.IsTimer
		gdss.Category = res.Category
		gdss.RedbagMoney = res.RedbagMoney
		gdss.OriginalPoint = res.OriginalPoint
		gdss.CurrentPoint = res.CurrentPoint
		gdss.TotalNum = res.TotalNum
		gdss.SoldNum = res.SoldNum
		gdss.SingleNum = res.SingleNum
		gdss.Name = res.Name
		gdss.Litpic = res.Litpic
		gdss.Content = res.Content
		response.GoodsDetails = gdss
	}
	response.Status = QUERY_GOODS_DETAILS_SUCCESS
	response.Msg = Stat[QUERY_GOODS_DETAILS_SUCCESS]
	Logger.Debugf("GetGoodsDetails response:%v", response)
	return &response, nil
}

/**
 * [StartGoodsDetailsServer [后台]商品详情服务]
 * @DateTime 2017-10-24T14:28:24+0800
 */
func StartGoodsDetailsServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30044"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/GoodsDetailsThriftService/providers"
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

	handler := &goodsdetailsservice{}
	processor := NewGoodsDetailsThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
