package advertadd

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	adAdd "cht/models/advertadd"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	ADD_ADVERT_SUCCESS = 1000
	ADD_ADVERT_FAILED  = 1001
)

var Stat = map[int]string{
	ADD_ADVERT_SUCCESS: "添加广告成功",
	ADD_ADVERT_FAILED:  "添加广告失败",
}

type advertaddservice struct{}

func (aas *advertaddservice) AddAdvert(requestObj *AdvertAddRequestStruct) (r *AdvertAddResponseStruct, err error) {
	aar := new(adAdd.AdvertAddRequest)
	aar.Type = requestObj.GetType()
	aar.Img = requestObj.GetImg()
	aar.Adverturl = requestObj.GetAdverturl()
	aar.Title = requestObj.GetTitle()
	aar.Adduser = requestObj.GetAdduser()
	aar.Fid = requestObj.GetFid()
	aar.Starttime = requestObj.GetStarttime()
	aar.Endtime = requestObj.GetEndtime()
	aar.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	b := adAdd.AddAdvert(aar)
	if b == false {
		Logger.Errorf("AddAdvert add failed:%v", err)
		return &AdvertAddResponseStruct{
			Status: ADD_ADVERT_FAILED,
			Msg:    Stat[ADD_ADVERT_FAILED],
		}, nil
	}

	return &AdvertAddResponseStruct{
		Status: ADD_ADVERT_SUCCESS,
		Msg:    Stat[ADD_ADVERT_SUCCESS],
	}, nil
}

/**
 * [StartAdvertAddServer 广告图片管理---添加广告图片服务]
 * @DateTime 2017-10-25T11:21:43+0800
 */
func StartAdvertAddServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30048"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/AdvertAddThriftService/providers"
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

	handler := &advertaddservice{}
	processor := NewAdvertAddThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
