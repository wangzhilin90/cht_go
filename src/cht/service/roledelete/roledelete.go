package roledelete

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/roledelete"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type roledeleteservice struct{}

const (
	ROLE_DELETE_SUCCESS = 1000
	ROLE_DELETE_FAILED  = 1001
)

var Status = map[int]string{
	ROLE_DELETE_SUCCESS: "角色删除成功",
	ROLE_DELETE_FAILED:  "角色删除失败",
}

func (rds *roledeleteservice) DeleteRole(requestObj *RoleDeleteRequestStruct) (r *RoleDeleteResponseStruct, err error) {
	rdrs := new(roledelete.RoleDeleteRequestStruct)
	rdrs.RoleIDStr = requestObj.GetRoleIDStr()
	b := roledelete.DeleteRole(rdrs)
	if b == false {
		Logger.Debugf("DeleteRole delete failed")
		return &RoleDeleteResponseStruct{
			Status: ROLE_DELETE_FAILED,
			Msg:    Status[ROLE_DELETE_FAILED],
		}, nil
	}

	Logger.Debugf("DeleteRole delete success")
	return &RoleDeleteResponseStruct{
		Status: ROLE_DELETE_SUCCESS,
		Msg:    Status[ROLE_DELETE_SUCCESS],
	}, nil
}

func StartRoleDeleteServer() {
	zkServers := []string{"192.168.8.208:2181"}
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30024"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/RoleDeleteThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &roledeleteservice{}
	processor := NewRoleDeleteThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
