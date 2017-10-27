package advertupdate

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	adup "cht/models/advertupdate"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	UPDATE_ADVERT_SUCCESS = 1000
	UPDATE_ADVERT_FAILED  = 1001
)

var Stat = map[int]string{
	UPDATE_ADVERT_SUCCESS: "修改广告图片成功",
	UPDATE_ADVERT_FAILED:  "修改广告图片失败",
}

type advertupdateservice struct{}

func (aus *advertupdateservice) UpdateAdvert(requestObj *AdvertUpdateRequestStruct) (r *AdvertUpdateResponseStruct, err error) {
	aur := new(adup.AdvertUpdateRequest)
	aur.ID = requestObj.GetID()
	aur.Type = requestObj.GetType()
	aur.Img = requestObj.GetImg()
	aur.Adverturl = requestObj.GetAdverturl()
	aur.Title = requestObj.GetTitle()
	aur.Adduser = requestObj.GetAdduser()
	aur.Fid = requestObj.GetFid()
	aur.Starttime = requestObj.GetStarttime()
	aur.Endtime = requestObj.GetEndtime()
	aur.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	b := adup.UpdateAdvert(aur)
	if b == false {
		Logger.Errorf("UpdateAdvert failed")
		return &AdvertUpdateResponseStruct{
			Status: UPDATE_ADVERT_FAILED,
			Msg:    Stat[UPDATE_ADVERT_FAILED],
		}, nil
	}

	return &AdvertUpdateResponseStruct{
		Status: UPDATE_ADVERT_SUCCESS,
		Msg:    Stat[UPDATE_ADVERT_SUCCESS],
	}, nil
}

/**
 * [StartAdvertListServer 广告图片管理---修改广告图片服务]
 * @DateTime 2017-10-25T15:57:37+0800
 */
func StartAdvertUpdateServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30052"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/AdvertUpdateThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &advertupdateservice{}
	processor := NewAdvertUpdateThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
