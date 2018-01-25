package advertdetails

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	adDetails "cht/models/advertdetails"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	GET_ADVERT_DETAILS_SUCCESS = 1000
	GET_ADVERT_DETAILS_FAILED  = 1001
	GET_ADVERT_DETAILS_EMPTY   = 1002
)

var Stat = map[int]string{
	GET_ADVERT_DETAILS_SUCCESS: "获取广告详情成功",
	GET_ADVERT_DETAILS_FAILED:  "获取广告详情失败",
	GET_ADVERT_DETAILS_EMPTY:   "获取广告详情为空",
}

type advertdetailsservice struct{}

func (ads *advertdetailsservice) GetAdvertDetails(requestObj *AdvertDetailsRequestStruct) (r *AdvertDetailsReponseStruct, err error) {
	Logger.Infof("GetAdvertDetails requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*AdvertDetailsRequestStruct)
	adr := new(adDetails.AdvertDetailsRequest)
	adr.ID = requestObj.GetID()
	adr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, err := adDetails.GetAdvertDetails(adr)
	if err != nil {
		Logger.Errorf("GetAdvertDetails get details failed:%v", err)
		return &AdvertDetailsReponseStruct{
			Status: GET_ADVERT_DETAILS_FAILED,
			Msg:    Stat[GET_ADVERT_DETAILS_FAILED],
		}, nil
	}

	if res == nil {
		Logger.Debugf("GetAdvertDetails query empty")
		return &AdvertDetailsReponseStruct{
			Status: GET_ADVERT_DETAILS_EMPTY,
			Msg:    Stat[GET_ADVERT_DETAILS_EMPTY],
		}, nil
	}

	var response AdvertDetailsReponseStruct
	if res != nil {
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
		response.AdvertDetails = adst
	}
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
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30050"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/AdvertDetailsThriftService/providers"
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

	handler := &advertdetailsservice{}
	processor := NewAdvertDetailsThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
