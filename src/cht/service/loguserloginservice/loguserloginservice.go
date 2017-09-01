package loguserloginservice

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/loguserlogin"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type LogUserlLoginRequest struct {
	UserID               int32
	LoginIP              string
	LoginStyle           int32
	ChengHuiTongTraceLog string
}

type LogUserLoginService struct{}

func (luls *LogUserLoginService) UpdateLogUserlLogin(requestObj *LogUserlLoginRequestStruct) (r *LogUserLoginResponseStruct, err error) {
	llr := new(loguserlogin.LogUserlLoginRequest)
	llr.UserID = requestObj.GetUserID()
	llr.LoginIP = requestObj.GetLoginIP()
	llr.LoginStyle = requestObj.GetLoginStyle()
	llr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	b, err := loguserlogin.UpdateLogUserlLogin(llr)
	if b == false {
		Logger.Error("UpdateLogUserlLogin failed")
		return nil, err
	}

	var llrs LogUserLoginResponseStruct
	llrs.UserID = llr.UserID
	return &llrs, nil
}

/**
 * [StartLogUserLoginServer 开启登录日志服务]
 * @DateTime 2017-08-30T17:58:45+0800
 */
func StartLogUserLoginServer() {
	zkServers := []string{"192.168.8.212:2181", "192.168.8.213:2181", "192.168.8.214:2181"}
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30003"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/LogUserLoginThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &LogUserLoginService{}
	processor := NewLogUserLoginThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
