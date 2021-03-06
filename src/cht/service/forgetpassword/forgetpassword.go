package forgetpassword

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	fp "cht/models/forgetpassword"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	UPDATE_FORGET_PASSWORD_SUCCESS = 1000
	UPDATE_FORGET_PASSWORD_FAILED  = 1001
)

var Stat = map[int]string{
	UPDATE_FORGET_PASSWORD_SUCCESS: "忘记密码更新密码成功",
	UPDATE_FORGET_PASSWORD_FAILED:  "忘记密码更新密码失败",
}

type forgetpasswordservice struct{}

func (fps *forgetpasswordservice) ForgetPassword(requestObj *ForgetPasswordRequestStruct) (r *ForgetPasswordResponseStruct, err error) {
	Logger.Infof("ForgetPassword requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*ForgetPasswordRequestStruct)
	fpr := new(fp.ForgetPasswordRequest)
	fpr.ID = requestObj.GetID()
	fpr.Password = requestObj.GetPassword()

	b := fp.ForgetPassword(fpr)
	if b == false {
		return &ForgetPasswordResponseStruct{
			Status: UPDATE_FORGET_PASSWORD_FAILED,
			Msg:    Stat[UPDATE_FORGET_PASSWORD_FAILED],
		}, nil
	}

	return &ForgetPasswordResponseStruct{
		Status: UPDATE_FORGET_PASSWORD_SUCCESS,
		Msg:    Stat[UPDATE_FORGET_PASSWORD_SUCCESS],
	}, nil
}

/**
 * [StartForgetPasswordServer 忘记密码重置密码服务]
 * @DateTime 2017-10-26T10:06:15+0800
 */
func StartForgetPasswordServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30053"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/ForgetPasswordThriftService/providers"
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

	handler := &forgetpasswordservice{}
	processor := NewForgetPasswordThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
