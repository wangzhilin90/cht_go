package advertdel

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	adDel "cht/models/advertdel"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	DELETE_ADVERT_SUCCESS = 1000
	DELETE_ADVERT_FAILED  = 1001
	GET_ADVERT_FID_FAILED = 1002
)

var Stat = map[int]string{
	DELETE_ADVERT_SUCCESS: "删除广告图片成功",
	DELETE_ADVERT_FAILED:  "删除广告图片失败",
	GET_ADVERT_FID_FAILED: "获取广告图片地址失败",
}

type advertdelservice struct{}

func (ads *advertdelservice) DelAdvert(requestObj *AdvertDelRequestStruct) (r *AdvertDelResponseStruct, err error) {
	Logger.Infof("DelAdvert requestObj:%v", requestObj)
	adr := new(adDel.AdvertDelRequest)
	adr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	adr.ID = requestObj.GetID()

	fid, err := adDel.GetAdvertFid(adr)
	if err != nil {
		Logger.Debugf("DelAdvert get fid failed:%v", err)
		return &AdvertDelResponseStruct{
			Status: GET_ADVERT_FID_FAILED,
			Msg:    Stat[GET_ADVERT_FID_FAILED],
		}, nil
	}

	b := adDel.DelAdvert(adr)
	if b == false {
		return &AdvertDelResponseStruct{
			Status: DELETE_ADVERT_FAILED,
			Msg:    Stat[DELETE_ADVERT_FAILED],
		}, nil
	}

	return &AdvertDelResponseStruct{
		Status: DELETE_ADVERT_SUCCESS,
		Msg:    Stat[DELETE_ADVERT_SUCCESS],
		Fid:    fid,
	}, nil
}

/**
 * [StartAdvertDelServer 广告图片管理---删除广告图片服务]
 * @DateTime 2017-10-25T13:48:57+0800
 */
func StartAdvertDelServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30049"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/AdvertDelThriftService/providers"
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

	handler := &advertdelservice{}
	processor := NewAdvertDelThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
