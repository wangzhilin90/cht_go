package advertlist

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	adList "cht/models/advertlist"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	QUERY_ADVERT_LIST_SUCCESS          = 1000
	QUERY_ADVERT_LIST_TOTAL_NUM_FAILED = 1001
	QUERY_ADVERT_LIST_FAILED           = 1002
	QUERY_ADVERT_LIST_EMPTY            = 1003
)

var Stat = map[int]string{
	QUERY_ADVERT_LIST_SUCCESS:          "广告图片列表查询成功",
	QUERY_ADVERT_LIST_TOTAL_NUM_FAILED: "广告图片列表总记录数查询失败",
	QUERY_ADVERT_LIST_FAILED:           "广告图片列表查询失败",
	QUERY_ADVERT_LIST_EMPTY:            "广告图片列表查询为空",
}

type advertlistservice struct{}

func (als *advertlistservice) GetAdvertList(requestObj *AdvertListRequestStruct) (r *AdvertListResponseStruct, err error) {
	Logger.Infof("GetAdvertList requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*AdvertListRequestStruct)
	alr := new(adList.AdvertListRequest)
	alr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	tatalNum, err := adList.GetAdvertListTatalNum(alr)
	if err != nil {
		Logger.Errorf("GetAdvertList get totalNUm failed:%v", err)
		return &AdvertListResponseStruct{
			Status: QUERY_ADVERT_LIST_TOTAL_NUM_FAILED,
			Msg:    Stat[QUERY_ADVERT_LIST_TOTAL_NUM_FAILED],
		}, nil
	}

	res, err := adList.GetAdvertList(alr)
	if err != nil {
		Logger.Errorf("GetAdvertList get advert list failed:%v", err)
		return &AdvertListResponseStruct{
			Status: QUERY_ADVERT_LIST_FAILED,
			Msg:    Stat[QUERY_ADVERT_LIST_FAILED],
		}, nil
	}

	if res == nil {
		Logger.Debugf("GetAdvertList query empty")
		return &AdvertListResponseStruct{
			Status: QUERY_ADVERT_LIST_EMPTY,
			Msg:    Stat[QUERY_ADVERT_LIST_EMPTY],
		}, nil
	}

	var response AdvertListResponseStruct
	for _, v := range res {
		als := new(AdvertListStruct)
		als.ID = v.ID
		als.Type = v.Type
		als.Img = v.Img
		als.Adverturl = v.Adverturl
		als.Title = v.Title
		als.Addtime = v.Addtime
		als.Adduser = v.Adduser
		als.Status = v.Status
		als.Fid = v.Fid
		als.Starttime = v.Starttime
		als.Endtime = v.Endtime
		response.AdvertList = append(response.AdvertList, als)
	}
	response.TotalNum = tatalNum
	response.Status = QUERY_ADVERT_LIST_SUCCESS
	response.Msg = Stat[QUERY_ADVERT_LIST_SUCCESS]
	Logger.Debugf("GetAdvertList response:%v", response)
	return &response, nil
}

/**
 * [StartAdvertDetailsServer 广告图片管理---列表服务]
 * @DateTime 2017-10-25T15:00:12+0800
 */
func StartAdvertListServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30051"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/AdvertListThriftService/providers"
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

	handler := &advertlistservice{}
	processor := NewAdvertListThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
