package advertmanagethriftservice

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/advertmanage"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	QUERY_ADVERT_MANAGE_SUCCESS = 1000
	QUERY_ADVERT_MANAGE_FAILED  = 1001
)

var Status = map[int]string{
	QUERY_ADVERT_MANAGE_SUCCESS: "查询广告成功",
	QUERY_ADVERT_MANAGE_FAILED:  "查询广告失败",
}

type advertmanageservice struct{}

func (ams *advertmanageservice) GetAdvertManage(requestObj *AdvertManageRequestStruct) (r *AdvertManageResponseStruct, err error) {
	amr := new(advertmanage.AdvertManageRequest)
	amr.Type = requestObj.GetType()
	amr.Limit = requestObj.GetLimit()
	AdvertManage, err := advertmanage.GetAdvertManage(amr)
	if err != nil {
		Logger.Errorf("GetAdvertManage return failed:%v", err)
		return &AdvertManageResponseStruct{
			Status: QUERY_ADVERT_MANAGE_FAILED,
			Msg:    Status[QUERY_ADVERT_MANAGE_FAILED],
		}, nil
	}
	var response AdvertManageResponseStruct
	for _, v := range AdvertManage {
		ams := new(AdvertManageStruct)
		ams.ID = v.ID
		ams.Type = v.Type
		ams.Img = v.Img
		ams.Adverturl = v.Adverturl
		ams.Title = v.Title
		ams.Addtime = v.Addtime
		ams.Adduser = v.Adduser
		ams.Status = v.Status
		ams.Fid = v.Fid
		ams.Starttime = v.Starttime
		ams.Endtime = v.Endtime
		response.AdvertManageList = append(response.AdvertManageList, ams)
	}
	response.Status = QUERY_ADVERT_MANAGE_SUCCESS
	response.Msg = Status[QUERY_ADVERT_MANAGE_SUCCESS]
	return &response, nil
}

/**
 * [StartCashRecordServer 开启广告管理]
 * @DateTime 2017-09-20T11:19:45+0800
 */
func StartAdvertManageServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30018"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/AdvertManageThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &advertmanageservice{}
	processor := NewAdvertManageThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
