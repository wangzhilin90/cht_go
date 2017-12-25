package sysuseredit

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	useredit "cht/models/sysuseredit"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	SYS_USER_EDIT_SUCCESS = 1000
	SYS_USER_EDIT_FAILED  = 1001
)

var Stat = map[int]string{
	SYS_USER_EDIT_SUCCESS: "编辑后台管理用户成功",
	SYS_USER_EDIT_FAILED:  "编辑后台管理用户失败",
}

type sysusereditservice struct{}

func (sues *sysusereditservice) EditSysUser(requestObj *SysUserEditRequestStruct) (r *SysUserEditResponseStruct, err error) {
	suer := new(useredit.SysUserEditRequest)
	suer.Account = requestObj.GetAccount()
	suer.Password = requestObj.GetPassword()
	suer.Realname = requestObj.GetRealname()
	suer.Mobile = requestObj.GetMobile()
	suer.Qq = requestObj.GetQq()
	suer.Status = requestObj.GetStatus()
	suer.Status = requestObj.GetStatus()
	suer.RoleID = requestObj.GetRoleID()
	suer.CustomerType = requestObj.GetCustomerType()
	suer.CreateTime = requestObj.GetCreateTime()
	suer.UserID = requestObj.GetUserID()
	suer.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	b := useredit.EditSysUser(suer)
	if b == false {
		return &SysUserEditResponseStruct{
			Status: SYS_USER_EDIT_FAILED,
			Msg:    Stat[SYS_USER_EDIT_FAILED],
		}, nil
	}

	return &SysUserEditResponseStruct{
		Status: SYS_USER_EDIT_SUCCESS,
		Msg:    Stat[SYS_USER_EDIT_SUCCESS],
	}, nil
}

/**
 * [StartSysUserEditServer 编辑后台管理用户服务]
 * @DateTime 2017-10-18T17:24:39+0800
 */
func StartSysUserEditServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30035"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/SysUserEditThriftService/providers"
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

	handler := &sysusereditservice{}
	processor := NewSysUserEditThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
