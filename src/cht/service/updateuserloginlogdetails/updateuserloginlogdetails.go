package updateuserloginlogdetails

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/loguserlogin"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
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

	QUERY_USER_LOGIN_LOG_SUCCESS = 1000
	QUERY_USER_LOGIN_LOG_FAILED  = 1001
)

var Stat = map[int]string{
	UPDATE_LOG_FAILED:  "更新登录日志失败",
	UPDATE_LOG_SUCCESS: "更新登录日志成功",
}

var Query_Stat = map[int]string{
	QUERY_USER_LOGIN_LOG_SUCCESS: "查询登录日志详情成功",
	QUERY_USER_LOGIN_LOG_FAILED:  "查询登录日志详情失败",
}

func (luls *LogUserLoginService) UpdateUserLoginLogDetails(requestObj *UpdateUserLoginLogDetailsRequestStruct) (r *UpdateUserLoginLogDetailsResponseStruct, err error) {
	Logger.Infof("UpdateUserLoginLogDetails requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*UpdateUserLoginLogDetailsRequestStruct)
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

	return &UpdateUserLoginLogDetailsResponseStruct{
		Status: UPDATE_LOG_SUCCESS,
		Msg:    Stat[UPDATE_LOG_SUCCESS],
	}, nil
}

func (luls *LogUserLoginService) GetUserLoginLogDetails(requestObj *UserLoginLogDetailsRequestStruct) (r *UserLoginLogDetailsResponseStruct, err error) {
	Logger.Infof("GetUserLoginLogDetails requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*UserLoginLogDetailsRequestStruct)
	ulldr := new(loguserlogin.UserLoginLogDetailsRequest)
	ulldr.UserID = requestObj.GetUserID()
	ulldr.LoginStyle = requestObj.GetLoginStyle()
	ulldr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, err := loguserlogin.GetUserLoginLogDetails(ulldr)
	if err != nil {
		Logger.Errorf("GetUserLoginLogDetails failed:%v", err)
		return &UserLoginLogDetailsResponseStruct{
			Status: QUERY_USER_LOGIN_LOG_FAILED,
			Msg:    Query_Stat[QUERY_USER_LOGIN_LOG_FAILED],
		}, nil
	}

	var response UserLoginLogDetailsResponseStruct
	if res != nil {
		ullds := new(UserLoginLogDetailsStruct)
		ullds.ID = res.ID
		ullds.UserID = res.UserID
		ullds.LoginTime = res.LoginTime
		ullds.LoginStyle = res.LoginStyle
		ullds.LoginIP = res.LoginIP
		ullds.TenderMoney = res.TenderMoney
		ullds.TenderTime = res.TenderTime
		response.UserLoginLogDetails = ullds
	}
	response.Status = QUERY_USER_LOGIN_LOG_SUCCESS
	response.Msg = Query_Stat[QUERY_USER_LOGIN_LOG_SUCCESS]
	Logger.Debugf("GetUserLoginLogDetails response:%v", response)
	return &response, nil
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

	handler := &LogUserLoginService{}
	processor := NewUpdateUserLoginLogDetailsThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
