package roleadd

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/roleadd"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

type roleaddservice struct{}

const (
	ROLE_ADD_SUCCESS = 1000
	ROLE_ADD_FAILED  = 1001
)

var Status = map[int]string{
	ROLE_ADD_SUCCESS: "角色添加成功",
	ROLE_ADD_FAILED:  "角色添加失败",
}

func (rs *roleaddservice) AddRole(requestObj *RoleAddRequestStruct) (r *RoleAddResponseStruct, err error) {
	rars := new(roleadd.RoleAddRequestStruct)
	rars.Remark = requestObj.GetRemark()
	rars.RoleName = requestObj.GetRoleName()

	b := roleadd.AddRole(rars)
	if b == false {
		Logger.Debugf("AddRole add failed")
		return &RoleAddResponseStruct{
			Status: ROLE_ADD_FAILED,
			Msg:    Status[ROLE_ADD_FAILED],
		}, nil
	}

	Logger.Debugf("AddRole add success")
	return &RoleAddResponseStruct{
		Status: ROLE_ADD_SUCCESS,
		Msg:    Status[ROLE_ADD_SUCCESS],
	}, nil
}

func StartRoleAddServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30023"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/RoleAddThriftService/providers"
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

	handler := &roleaddservice{}
	processor := NewRoleAddThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
