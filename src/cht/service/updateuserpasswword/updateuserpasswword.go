package updateuserpasswword

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/updatepasswd"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type updatepasswdservice struct{}

const (
	UPDATEPASSWDFAILED  = 1001
	UPDATEPASSWDSUCCESS = 1000
)

var Status = map[int]string{
	UPDATEPASSWDFAILED:  "更新密码失败",
	UPDATEPASSWDSUCCESS: "更新密码成功",
}

func (ups *updatepasswdservice) UpdateUserPassword(requestObj *UpdateUserPasswWordRequestStruct) (r *UpdateUserPasswWordResponseStruct, err error) {
	upr := new(updatepasswd.UpdatePasswdRequest)
	upr.ID = requestObj.GetID()
	upr.Password = requestObj.GetPassword()

	b := updatepasswd.UpdatePasswd(upr)
	if b == false {
		return &UpdateUserPasswWordResponseStruct{
			Status: UPDATEPASSWDFAILED,
			Msg:    Status[UPDATEPASSWDFAILED],
		}, nil
	}
	return &UpdateUserPasswWordResponseStruct{
		Status: UPDATEPASSWDSUCCESS,
		Msg:    Status[UPDATEPASSWDSUCCESS],
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
