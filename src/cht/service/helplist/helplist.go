package helplist

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/helplist"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

type helplistservice struct{}

func (hs *helplistservice) GetHelpList(requestObj *HelpListRequestStruct) (r *HelpListResponseStrcut, err error) {
	Logger.Infof("GetHelpList requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*HelpListRequestStruct)
	hr := new(helplist.HelpListRequest)
	hr.Status = requestObj.GetStatus()
	hr.Cateid = requestObj.GetCateid()
	hr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	var response HelpListResponseStrcut
	res, _ := helplist.GetHelpList(hr)
	for _, v := range res {
		hlrs := new(HelpListResultStruct)
		hlrs.Title = v.Title
		hlrs.Content = v.Content
		response.HelpList = append(response.HelpList, hlrs)
	}
	Logger.Debugf("GetHelpList return value:%v", response)
	return &response, nil
}

/*获取帮助中心文章列表服务*/
func StartHelpListServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30030"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/HelpListThriftService/providers"
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

	handler := &helplistservice{}
	processor := NewHelpListThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
