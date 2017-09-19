package securedthriftservice

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/secured"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type securedservice struct{}

const (
	QUERY_SECURED_SUCCESS = 1000
	QUERY_SECURED_FAILED  = 1001
)

var Status = map[int]string{
	QUERY_SECURED_SUCCESS: "查询担保人成功",
	QUERY_SECURED_FAILED:  "查询担保人失败",
}

func (ss *securedservice) GetSecuredList(requestObj *SecuredRequestStruct) (r *SecuredResponseStruct, err error) {
	srs := new(secured.SecuredRequestStruct)
	srs.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	securedList, err := secured.GetSecuredList(srs)
	if err != nil {
		Logger.Errorf("GetSecuredList query failed %v", err)
		return &SecuredResponseStruct{
			Status: QUERY_SECURED_FAILED,
			Msg:    Status[QUERY_SECURED_FAILED],
		}, nil
	}

	response := new(SecuredResponseStruct)
	for _, v := range securedList {
		sis := new(SecuredInfoStruct)
		sis.Secured = v.Secured
		response.SecuredInfo = append(response.SecuredInfo, sis)
	}
	response.Status = QUERY_SECURED_SUCCESS
	response.Msg = Status[QUERY_SECURED_SUCCESS]
	Logger.Debugf("GetSecuredList res:%v", response)
	return response, nil
}

/*开启做标服务---担保人服务*/
func StartSecuredServer() {
	zkServers := []string{"192.168.8.208:2181"}
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30013"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/securedThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &securedservice{}
	processor := NewSecuredThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
