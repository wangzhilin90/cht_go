package sysconfigthriftservice

import (
	_ "cht/initial"
	"testing"
)

func NewSysConfigRequestStruct() *SysConfigRequestStruct {
	return &SysConfigRequestStruct{}
}

func TestGetSysConfig(t *testing.T) {
	scrs := NewSysConfigRequestStruct()
	ss := sysconfigservice{}
	res, err := ss.GetSysConfig(scrs)
	if err != nil {
		t.Fatalf("TestGetSysConfig failed :%v", err)
	}
	t.Logf("TestGetSysConfig res:%v", res)
}

/**
 * [StartsubledgerServer 开启查询系统配置服务]
 * @DateTime 2017-09-19T17:58:45+0800
 */
func StartsubledgerServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30017"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/subledgerThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode %v failed", servicename, err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &subledgerservice{}
	processor := NewSubledgerThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
