package updateuserloginlogdetails

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

const (
	UPDATE_LOG_FAILED  = 1001
	UPDATE_LOG_SUCCESS = 1000
)

var Stat = map[int]string{
	UPDATE_LOG_FAILED:  "更新登录日志失败",
	UPDATE_LOG_SUCCESS: "更新登录日志成功",
}

func (luls *LogUserLoginService) UpdateUserLoginLogDetails(requestObj *UpdateUserLoginLogDetailsRequestStruct) (r *UpdateUserLoginLogDetailsResponseStruct, err error) {
	llr := new(loguserlogin.LogUserlLoginRequest)
	llr.UserID = requestObj.GetUserID()
	llr.LoginIP = requestObj.GetLoginIP()
	llr.LoginStyle = requestObj.GetLoginStyle()
	llr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	b, err := loguserlogin.UpdateLogUserlLogin(llr)
	if b == false {
		Logger.Error("UpdateUserLoginLogDetails failed", err)
		return &UpdateUserLoginLogDetailsResponseStruct{
			Status: UPDATE_LOG_FAILED,
			Msg:    Stat[UPDATE_LOG_FAILED],
		}, nil
	}
	Logger.Debug("UpdateUserLoginLogDetails success")
	return &UpdateUserLoginLogDetailsResponseStruct{
		Status: UPDATE_LOG_SUCCESS,
		Msg:    Stat[UPDATE_LOG_SUCCESS],
	}, nil
}

/**
 * [StartLogUserLoginServer 开启登录日志服务]
 * @DateTime 2017-08-30T17:58:45+0800
 */
func StartLogUserLoginServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30003"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/UpdateUserLoginLogDetailsThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &LogUserLoginService{}
	processor := NewUpdateUserLoginLogDetailsThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
