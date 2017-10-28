package goodsedit

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	ge "cht/models/goodsedit"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	EDIT_GOODS_SUCCESS = 1000
	EDIT_GOODS_FAILED  = 1001
)

var Stat = map[int]string{
	EDIT_GOODS_SUCCESS: "编辑商品成功",
	EDIT_GOODS_FAILED:  "编辑商品失败",
}

type goodseditservice struct{}

func (ges *goodseditservice) EditGoods(requestObj *GoodsEditRequestStruct) (r *GoodsEditResponseStruct, err error) {
	ger := new(ge.GoodsEditRequest)
	ger.ID = requestObj.GetID()
	ger.ShowTime = requestObj.GetShowTime()
	ger.CloseTime = requestObj.GetCloseTime()
	ger.IsTimer = requestObj.GetIsTimer()
	ger.Litpic = requestObj.GetLitpic()
	ger.Name = requestObj.GetName()
	ger.Category = requestObj.GetCategory()
	ger.RedbagMoney = requestObj.GetRedbagMoney()
	ger.OriginalPoint = requestObj.GetOriginalPoint()
	ger.CurrentPoint = requestObj.GetCurrentPoint()
	ger.TotalNum = requestObj.GetTotalNum()
	ger.SingleNum = requestObj.GetSingleNum()
	ger.Content = requestObj.GetContent()
	ger.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	b := ge.EditGoods(ger)
	if b == false {
		return &GoodsEditResponseStruct{
			Status: EDIT_GOODS_FAILED,
			Msg:    Stat[EDIT_GOODS_FAILED],
		}, nil
	}

	return &GoodsEditResponseStruct{
		Status: EDIT_GOODS_SUCCESS,
		Msg:    Stat[EDIT_GOODS_SUCCESS],
	}, nil
}

/**
 * [StartGoodsEditServer [后台服务]编辑商品服务]
 * @DateTime 2017-10-24T14:57:42+0800
 */
func StartGoodsEditServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30045"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/GoodsEditThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode %v failed", servicename, err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &goodseditservice{}
	processor := NewGoodsEditThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
