package talking

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/talking"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

type talkingservcie struct{}

func (ts *talkingservcie) GetTalkingList(requestObj *TalkingRequestStruct) (r *TalkingResponseStruct, err error) {
	Logger.Infof("GetTalkingList requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*TalkingRequestStruct)
	tr := new(talking.TalkingRequest)
	tr.Cateid = requestObj.GetCateid()
	tr.Status = requestObj.GetStatus()
	tr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	var response TalkingResponseStruct
	talklist, _ := talking.GetTalkinglist(tr)
	for _, v := range talklist {
		tlrs := new(TalkListResultStruct)
		tlrs.ID = v.ID
		tlrs.Title = v.Title
		tlrs.ImgURL = v.ImgURL
		tlrs.Content = v.Content
		response.TalkList = append(response.TalkList, tlrs)
	}

	onelist, _ := talking.GetOnelist(tr)
	for _, v := range onelist {
		tlrs := new(TalkListResultStruct)
		tlrs.ID = v.ID
		tlrs.Title = v.Title
		tlrs.ImgURL = v.ImgURL
		tlrs.Content = v.Content
		response.OneList = append(response.OneList, tlrs)
	}
	Logger.Debugf("GetTalkingList return value:%v", response)
	return &response, nil
}

/**
 * [StartSysConfigServer 开启[后台]小城交流日服务]
 * @DateTime 2017-10-11T14:58:45+0800
 */
func StartTalkingServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30028"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/TalkingThriftService/providers"
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

	handler := &talkingservcie{}
	processor := NewTalkingThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
