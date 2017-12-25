package sysuserdelete

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	userdelete "cht/models/sysuserdelete"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	DELETE_SYS_USER_SUCCESS = 1000
	DELETE_SYS_USER_FAILED  = 1001
)

var Stat = map[int]string{
	DELETE_SYS_USER_SUCCESS: "删除后台系统用户成功",
	DELETE_SYS_USER_FAILED:  "删除后台系统用户失败",
}

type sysuserdeleteservice struct{}

func (suds *sysuserdeleteservice) DeleteSysUser(requestObj *SysUserDeleteRequestStruct) (r *SysUserDeleteResponseStruct, err error) {
	sudr := new(userdelete.SysUserDeleteRequest)
	sudr.UserIDStr = requestObj.GetUserIDStr()
	sudr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	b := userdelete.DeleteSysUser(sudr)
	if b == false {
		return &SysUserDeleteResponseStruct{
			Status: DELETE_SYS_USER_FAILED,
			Msg:    Stat[DELETE_SYS_USER_FAILED],
		}, nil
	}

	return &SysUserDeleteResponseStruct{
		Status: DELETE_SYS_USER_SUCCESS,
		Msg:    Stat[DELETE_SYS_USER_SUCCESS],
	}, nil
}

/**
 * [StartSysUserDeleteServer 删除后台管理用户服务]
 * @DateTime 2017-10-17T17:29:09+0800
 */
func StartSysUserDeleteServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30033"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/SysUserDeleteThriftService/providers"
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

	handler := &sysuserdeleteservice{}
	processor := NewSysUserDeleteThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
