package securedlist

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

func (ss *securedservice) GetSecuredList(requestObj *SecuredListRequestStruct) (r *SecuredListResponseStruct, err error) {
	srs := new(secured.SecuredRequestStruct)
	srs.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	securedList, err := secured.GetSecuredList(srs)
	if err != nil {
		Logger.Errorf("GetSecuredList query failed %v", err)
		return &SecuredListResponseStruct{
			Status: QUERY_SECURED_FAILED,
			Msg:    Status[QUERY_SECURED_FAILED],
		}, nil
	}

	response := new(SecuredListResponseStruct)
	for _, v := range securedList {
		sis := new(SecuredDetailsStruct)
		sis.Secured = v.Secured
		response.SecuredList = append(response.SecuredList, sis)
	}

	//固定担保人
	var PermanentSecured = []SecuredDetailsStruct{
		{"贵州喜年华装饰工程有限公司/贵州联宇置业有限公司"},
		{"贵州行成道企业管理有限公司/贵州联宇置业有限公司"},
		{"贵州联宇同行汽车销售服务有限公司/贵州联宇置业有限公司"},
		{"深圳市合泰典当有限公司"},
		{"贵州保胜信用管理有限公司"},
	}

	for _, v := range PermanentSecured {
		response.SecuredList = append(response.SecuredList, &v)
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

	servicename := "/cht/SecuredListThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &securedservice{}
	processor := NewSecuredListThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
