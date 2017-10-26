package updateuserpasswword

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	up "cht/models/updatepasswd"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type updatepasswdservice struct{}

const (
	UPDATE_PASSWORD_SUCCESS  = 1000
	GET_OLD_PASSWORD_FAILED  = 1001
	OLD_PASSWORD_INPUT_ERROR = 1002
	UPDATE_PASSWORD_FAILED   = 1003
)

var Status = map[int]string{
	UPDATE_PASSWORD_SUCCESS:  "更新密码成功",
	GET_OLD_PASSWORD_FAILED:  "获取数据库登录密码失败",
	OLD_PASSWORD_INPUT_ERROR: "旧密码输入不正确",
	UPDATE_PASSWORD_FAILED:   "更新密码失败",
}

func (ups *updatepasswdservice) UpdateUserPasswWord(requestObj *UpdateUserPasswWordRequestStruct) (r *UpdateUserPasswWordResponseStruct, err error) {
	upr := new(up.UpdatePasswdRequest)
	upr.ID = requestObj.GetID()
	upr.NewPassword_ = requestObj.GetNewPassword_()
	upr.OldPassword = requestObj.GetOldPassword()

	oldPw, err := up.GetDBPasswd(upr)
	if err != nil {
		Logger.Errorf("UpdateUserPasswWord get old password failed:%v", err)
		return &UpdateUserPasswWordResponseStruct{
			Status: GET_OLD_PASSWORD_FAILED,
			Msg:    Status[GET_OLD_PASSWORD_FAILED],
		}, nil
	}

	if oldPw != requestObj.GetOldPassword() {
		Logger.Errorf("UpdateUserPasswWord old password input error")
		return &UpdateUserPasswWordResponseStruct{
			Status: OLD_PASSWORD_INPUT_ERROR,
			Msg:    Status[OLD_PASSWORD_INPUT_ERROR],
		}, nil
	}

	b := up.UpdatePasswd(upr)
	if b == false {
		Logger.Errorf("UpdateUserPasswWord update password failed")
		return &UpdateUserPasswWordResponseStruct{
			Status: UPDATE_PASSWORD_FAILED,
			Msg:    Status[UPDATE_PASSWORD_FAILED],
		}, nil
	}

	Logger.Debugf("UpdateUserPasswWord update password success")
	return &UpdateUserPasswWordResponseStruct{
		Status: UPDATE_PASSWORD_SUCCESS,
		Msg:    Status[UPDATE_PASSWORD_SUCCESS],
	}, nil
}

/**
 * [StartUpdatePasswdsServer 开启忘记密码重置密码服务]
 * @DateTime 2017-08-24T15:19:45+0800
 */
func StartUpdatePasswdsServer() {
	zkServers := []string{"192.168.8.208:2181"}
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30004"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/UpdateUserPasswWordThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &updatepasswdservice{}
	processor := NewUpdateUserPasswWordThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
