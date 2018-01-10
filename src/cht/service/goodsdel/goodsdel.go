package goodsdel

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	gd "cht/models/goodsdel"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	DELETE_GOODS_SUCCESS = 1000
	DELETE_GOODS_FAILED  = 1001
)

var Stat = map[int]string{
	DELETE_GOODS_SUCCESS: "删除商品成功",
	DELETE_GOODS_FAILED:  "删除商品失败",
}

type goodsdelservice struct{}

func (gds *goodsdelservice) DelGoods(requestObj *GoodsDeLRequestStruct) (r *GoodsDeLResponseStruct, err error) {
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*GoodsDeLRequestStruct)
	gdr := new(gd.GoodsDeLRequest)
	gdr.ID = requestObj.GetID()
	gdr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	b := gd.DelGoods(gdr)
	if b == false {
		return &GoodsDeLResponseStruct{
			Status: DELETE_GOODS_FAILED,
			Msg:    Stat[DELETE_GOODS_FAILED],
		}, nil
	}

	return &GoodsDeLResponseStruct{
		Status: DELETE_GOODS_SUCCESS,
		Msg:    Stat[DELETE_GOODS_SUCCESS],
	}, nil
}

/**
 * [StartDelGoodsServer 开启[后台]商品管理---删除商品服务]
 * @DateTime 2017-10-23T17:11:09+0800
 */
func StartDelGoodsServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30043"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/GoodsDeLThriftService/providers"
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

	handler := &goodsdelservice{}
	processor := NewGoodsDeLThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
