package accessconfig

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	ac "cht/models/accessconfig"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
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
	acr := new(ac.AccessConfigRequest)
	acr.Source = requestObj.GetSource()
	acr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, err := ac.GetAccessConfig(acr)
	if err != nil {
		return &AccessConfigResponseStruct{
			Status: QUERY_ACCESS_CONFIG_FAILED,
			Msg:    Stat[QUERY_ACCESS_CONFIG_FAILED],
		}, nil
	}

	acst := new(AccessConfigStruct)
	acst.ID = res.ID
	acst.Name = res.Name
	acst.Source = res.Source
	acst.Addtime = res.Addtime

	return &AccessConfigResponseStruct{
		Status:       QUERY_ACCESS_CONFIG_SUCCESS,
		Msg:          Stat[QUERY_ACCESS_CONFIG_SUCCESS],
		AccessConfig: acst,
	}, nil
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

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &accessconfigservice{}
	processor := NewAccessConfigThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
