package userattestionbaseinfosave

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	uabs "cht/models/userattestionbaseinfosave"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	SAVE_USER_ATTESTION_BASE_INFO_SUCCESS = 1000
	SAVE_USER_ATTESTION_BASE_INFO_FAILED  = 1001
)

var Stat = map[int]string{
	SAVE_USER_ATTESTION_BASE_INFO_SUCCESS: "保存用户认证信息成功",
	SAVE_USER_ATTESTION_BASE_INFO_FAILED:  "保存用户认证信息失败",
}

type userattestionbaseinfosaveservice struct{}

/*保存用户认证信息*/
func (uabss *userattestionbaseinfosaveservice) SaveUserAttestionBaseInfo(requestObj *UserAttestionBaseInfoSaveRequestStruct) (r *UserAttestionBaseInfoSaveResponseStruct, err error) {
	usr := new(uabs.UserAttestionBaseInfoSaveRequest)
	usr.UserID = requestObj.GetUserID()
	usr.VideoPic = requestObj.GetVideoPic()
	usr.RealStatus = requestObj.GetRealStatus()
	usr.EmailStatus = requestObj.GetEmailStatus()
	usr.PhoneStatus = requestObj.GetPhoneStatus()
	usr.VideoStatus = requestObj.GetVideoStatus()
	usr.SceneStatus = requestObj.GetSceneStatus()
	usr.RealPasstime = requestObj.GetRealPasstime()
	usr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	b := uabs.SaveUserAttestionBaseInfo(usr)
	if b == false {
		return &UserAttestionBaseInfoSaveResponseStruct{
			Status: SAVE_USER_ATTESTION_BASE_INFO_FAILED,
			Msg:    Stat[SAVE_USER_ATTESTION_BASE_INFO_FAILED],
		}, nil
	}

	return &UserAttestionBaseInfoSaveResponseStruct{
		Status: SAVE_USER_ATTESTION_BASE_INFO_SUCCESS,
		Msg:    Stat[SAVE_USER_ATTESTION_BASE_INFO_SUCCESS],
	}, nil
}

/**
 * [StartUserAttestionBaseInfoSaveServer 保持用户认证信息服务]
 * @DateTime 2017-10-19T17:14:16+0800
 */
func StartUserAttestionBaseInfoSaveServer() {
	zkServers := []string{"192.168.8.208:2181"}
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30038"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/UserAttestionBaseInfoSaveThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &userattestionbaseinfosaveservice{}
	processor := NewUserAttestionBaseInfoSaveThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
