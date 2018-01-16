package accessconfig

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	ac "cht/models/accessconfig"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	QUERY_ACCESS_CONFIG_SUCCESS = 1000
	QUERY_ACCESS_CONFIG_FAILED  = 1001
)

var Stat = map[int]string{
	QUERY_ACCESS_CONFIG_SUCCESS: "查询推广名称记录表成功",
	QUERY_ACCESS_CONFIG_FAILED:  "查询推广名称记录表失败",
}

type accessconfigservice struct{}

func (acs *accessconfigservice) GetAccessConfig(requestObj *AccessConfigRequestStruct) (r *AccessConfigResponseStruct, err error) {
	Logger.Infof("GetAccessConfig requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*AccessConfigRequestStruct)
	acr := new(ac.AccessConfigRequest)
	acr.Source = requestObj.GetSource()
	acr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, err := ac.GetAccessConfig(acr)
	if err != nil {
		Logger.Errorf("GetAccessConfig failed:%v", err)
		return &AccessConfigResponseStruct{
			Status: QUERY_ACCESS_CONFIG_FAILED,
			Msg:    Stat[QUERY_ACCESS_CONFIG_FAILED],
		}, nil
	}

	var response AccessConfigResponseStruct
	if res != nil {
		acst := new(AccessConfigStruct)
		acst.ID = res.ID
		acst.Name = res.Name
		acst.Source = res.Source
		acst.Addtime = res.Addtime
		response.AccessConfig = acst
	}
	response.Status = QUERY_ACCESS_CONFIG_SUCCESS
	response.Msg = Stat[QUERY_ACCESS_CONFIG_SUCCESS]
	Logger.Debugf("GetAccessConfig response:%v", response)
	return &response, nil
}

/**
 * [StartAccessConfigServer [后台]推广名称记录表服务]
 * @DateTime 2017-10-30T16:42:17+0800
 */
func StartAccessConfigServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30064"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/AccessConfigThriftService/providers"
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

	handler := &accessconfigservice{}
	processor := NewAccessConfigThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
