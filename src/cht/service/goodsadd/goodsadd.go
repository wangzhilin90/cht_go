package goodsadd

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	ga "cht/models/goodsadd"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	ADD_GOODS_SUCCESS = 1000
	ADD_GOODS_FAILED  = 1001
)

var Stat = map[int]string{
	ADD_GOODS_SUCCESS: "添加商品成功",
	ADD_GOODS_FAILED:  "添加商品失败",
}

type goodsaddservice struct{}

func NewGoodsAddRequest(requestObj *GoodsAddRequestStruct) *ga.GoodsAddRequest {
	gar := new(ga.GoodsAddRequest)
	gar.ShowTime = requestObj.GetShowTime()
	gar.CloseTime = requestObj.GetCloseTime()
	gar.IsTimer = requestObj.GetIsTimer()
	gar.Litpic = requestObj.GetLitpic()
	gar.Name = requestObj.GetName()
	gar.Category = requestObj.GetCategory()
	gar.RedbagMoney = requestObj.GetRedbagMoney()
	gar.OriginalPoint = requestObj.GetOriginalPoint()
	gar.CurrentPoint = requestObj.GetCurrentPoint()
	gar.TotalNum = requestObj.GetTotalNum()
	gar.SingleNum = requestObj.GetSingleNum()
	gar.Content = requestObj.GetContent()
	gar.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	return gar
}

func initGoodsDefaultValue(req *ga.GoodsAddRequest) *ga.GoodsAddRequest {
	if req.Name == "" {
		req.Name = " "
	}

	if req.Litpic == "" {
		req.Litpic = " "
	}

	if req.RedbagMoney == "" {
		req.RedbagMoney = " "
	}

	if req.Content == "" {
		req.Content = " "
	}

	return req
}

func (gas *goodsaddservice) AddGoods(requestObj *GoodsAddRequestStruct) (r *GoodsAddResponseStruct, err error) {
	req := NewGoodsAddRequest(requestObj)
	req = initGoodsDefaultValue(req)
	b := ga.AddGoods(req)
	if b == false {
		Logger.Errorf("AddGoods insert failed")
		return &GoodsAddResponseStruct{
			Status: ADD_GOODS_FAILED,
			Msg:    Stat[ADD_GOODS_FAILED],
		}, nil
	}

	return &GoodsAddResponseStruct{
		Status: ADD_GOODS_SUCCESS,
		Msg:    Stat[ADD_GOODS_SUCCESS],
	}, nil
}

/**
 * [StartAddGoodsServer [后台]商品管理---添加商品服务]
 * @DateTime 2017-10-23T16:38:58+0800
 */
func StartAddGoodsServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30042"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/GoodsAddThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &goodsaddservice{}
	processor := NewGoodsAddThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
