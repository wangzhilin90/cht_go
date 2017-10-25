package advertdetails

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	adDetails "cht/models/advertdetails"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	GET_ADVERT_DETAILS_SUCCESS = 1000
	GET_ADVERT_DETAILS_FAILED  = 1001
)

var Stat = map[int]string{
	GET_ADVERT_DETAILS_SUCCESS: "获取广告详情成功",
	GET_ADVERT_DETAILS_FAILED:  "获取广告详情失败",
}

type advertdetailsservice struct{}

func (ads *advertdetailsservice) GetAdvertDetails(requestObj *AdvertDetailsRequestStruct) (r *AdvertDetailsReponseStruct, err error) {
	adr := new(adDetails.AdvertDetailsRequest)
	adr.ID = requestObj.GetID()
	adr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, err := adDetails.GetAdvertDetails(adr)
	if err != nil {
		Logger.Debugf("GetAdvertDetails get details failed:%v", err)
		return &AdvertDetailsReponseStruct{
			Status: GET_ADVERT_DETAILS_FAILED,
			Msg:    Stat[GET_ADVERT_DETAILS_FAILED],
		}, nil
	}

	adst := new(AdvertDetailsStruct)
	adst.ID = res.ID
	adst.Type = res.Type
	adst.Img = res.Img
	adst.Adverturl = res.Adverturl
	adst.Title = res.Title
	adst.Addtime = res.Addtime
	adst.Adduser = res.Adduser
	adst.Status = res.Status
	adst.Fid = res.Fid
	adst.Starttime = res.Starttime
	adst.Endtime = res.Endtime

	var response AdvertDetailsReponseStruct
	response.AdvertDetails = adst
	response.Status = GET_ADVERT_DETAILS_SUCCESS
	response.Msg = Stat[GET_ADVERT_DETAILS_SUCCESS]
	Logger.Debugf("GetAdvertDetails response:%v", response)
	return &response, nil
}

/**
 * [StartAdvertDetailsServer 广告图片管理---图片详情服务]
 * @DateTime 2017-10-25T14:24:15+0800
 */
func StartAdvertDetailsServer() {
	zkServers := []string{"192.168.8.208:2181"}
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30049"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/AdvertAddThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &advertdetailsservice{}
	processor := NewAdvertAddThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
