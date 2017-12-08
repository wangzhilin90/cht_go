package sysconfigthriftservice

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/sysconfig"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	QUERY_SYS_CONFIG_SUCCESS = 1000
	QUERY_SYS_CONFIG_FAILED  = 1001
)

var Status = map[int]string{
	QUERY_SYS_CONFIG_SUCCESS: "查询系统配置成功",
	QUERY_SYS_CONFIG_FAILED:  "查询系统配置失败",
}

type sysconfigservice struct{}

func (scs *sysconfigservice) GetSysConfig(requestObj *SysConfigRequestStruct) (r *SysConfigResponseStruct, err error) {
	Logger.Info("GetSysConfig input param:", requestObj)
	scrs := new(sysconfig.SysConfigRequestStruct)
	scrs.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	SysConfig, err := sysconfig.GetSysConfig(scrs)
	if err != nil {
		Logger.Debugf("GetSysConfig query failed :%v", err)
		return &SysConfigResponseStruct{
			Status: QUERY_SYS_CONFIG_FAILED,
			Msg:    Status[QUERY_SYS_CONFIG_FAILED],
		}, nil
	}

	var response SysConfigResponseStruct
	for _, v := range SysConfig {
		scs := new(SysConfigStruct)
		scs.ID = v.ID
		scs.Nid = v.Nid
		scs.Value = v.Value
		scs.Name = v.Name
		response.SysConfigList = append(response.SysConfigList, scs)
	}
	response.Status = QUERY_SYS_CONFIG_SUCCESS
	response.Msg = Status[QUERY_SYS_CONFIG_SUCCESS]
	Logger.Debugf("GetSysConfig res:%v", response)
	return &response, nil
}

/**
 * [StartSysConfigServer 开启查询系统配置服务]
 * @DateTime 2017-09-19T17:58:45+0800
 */
func StartSysConfigServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30017"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/SysConfigThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode %v failed", servicename, err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &sysconfigservice{}
	processor := NewSysConfigThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
