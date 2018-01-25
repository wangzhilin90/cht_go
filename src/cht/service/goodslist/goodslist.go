package goodslist

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	gl "cht/models/goodslist"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	GET_GOODS_LIST_SUCCESS          = 1000
	GET_GOODS_LIST_TOTAL_NUM_FAILED = 1001
	GET_GOODS_LIST_FAILED           = 1002
	GET_GOODS_LIST_EMPTY            = 1003
)

var Stat = map[int]string{
	GET_GOODS_LIST_SUCCESS:          "获取商品管理列表成功",
	GET_GOODS_LIST_TOTAL_NUM_FAILED: "获取商品管理列表总数失败",
	GET_GOODS_LIST_FAILED:           "获取商品管理列表失败",
	GET_GOODS_LIST_EMPTY:            "获取商品管理列表为空",
}

type goodslistservice struct{}

func (gls *goodslistservice) GetGoodsList(requestObj *GoodsListRequestStruct) (r *GoodsListReponseStruct, err error) {
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*GoodsListRequestStruct)
	glr := new(gl.GoodsListRequest)
	glr.Name = requestObj.GetName()
	glr.Category = requestObj.GetCategory()
	glr.IsExport = requestObj.GetIsExport()
	glr.LimitOffset = requestObj.GetLimitOffset()
	glr.LimitNum = requestObj.GetLimitNum()
	glr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	totalNum, err := gl.GetGoodsListTotalNum(glr)
	if err != nil {
		Logger.Errorf("GetGoodsList get totalNum failed:%v", err)
		return &GoodsListReponseStruct{
			Status: GET_GOODS_LIST_TOTAL_NUM_FAILED,
			Msg:    Stat[GET_GOODS_LIST_TOTAL_NUM_FAILED],
		}, nil
	}

	res, err := gl.GetGoodsList(glr)
	if err != nil {
		Logger.Errorf("GetGoodsList get goodslist failed:%v", err)
		return &GoodsListReponseStruct{
			Status: GET_GOODS_LIST_FAILED,
			Msg:    Stat[GET_GOODS_LIST_FAILED],
		}, nil
	}

	if res == nil {
		Logger.Debugf("GetGoodsList query empty")
		return &GoodsListReponseStruct{
			Status: GET_GOODS_LIST_EMPTY,
			Msg:    Stat[GET_GOODS_LIST_EMPTY],
		}, nil
	}

	var response GoodsListReponseStruct
	for _, v := range res {
		glrs := new(GoodsListResultStruct)
		glrs.ID = v.ID
		glrs.Addtime = v.Addtime
		glrs.ShowTime = v.ShowTime
		glrs.CloseTime = v.CloseTime
		glrs.IsTimer = v.IsTimer
		glrs.Category = v.Category
		glrs.RedbagMoney = v.RedbagMoney
		glrs.OriginalPoint = v.OriginalPoint
		glrs.CurrentPoint = v.CurrentPoint
		glrs.TotalNum = v.TotalNum
		glrs.SoldNum = v.SoldNum
		glrs.SingleNum = v.SingleNum
		glrs.Name = v.Name
		glrs.Litpic = v.Litpic
		glrs.Content = v.Content
		response.GoodsList = append(response.GoodsList, glrs)
	}
	response.TotalNum = totalNum
	response.Status = GET_GOODS_LIST_SUCCESS
	response.Msg = Stat[GET_GOODS_LIST_SUCCESS]
	Logger.Debugf("GetGoodsList response:%v", response)
	return &response, nil
}

/**
 * [StartGoodsListServer [后台服务]商品管理列表服务]
 * @DateTime 2017-10-24T15:51:27+0800
 */
func StartGoodsListServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30046"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/GoodsListThriftService/providers"
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

	handler := &goodslistservice{}
	processor := NewGoodsListThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
