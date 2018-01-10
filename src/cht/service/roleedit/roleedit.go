package roleedit

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/roleedit"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

type roleeditservice struct{}

const (
	ROLE_EDIT_SUCCESS = 1000
	ROLE_EDIT_FAILED  = 1001
)

var Status = map[int]string{
	ROLE_EDIT_SUCCESS: "角色编辑成功",
	ROLE_EDIT_FAILED:  "角色编辑失败",
}

func (res *roleeditservice) EditRole(requestObj *RoleEditRequestStruct) (r *RoleEditResponseStruct, err error) {
	Logger.Infof("EditRole requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*RoleEditRequestStruct)
	rers := new(roleedit.RoleEditRequestStruct)
	rers.Remark = requestObj.GetRemark()
	rers.RoleID = requestObj.GetRoleID()
	rers.RoleName = requestObj.GetRoleName()
	rers.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	b := roleedit.EditRole(rers)
	if b == false {
		Logger.Debugf("EditRole edit failed")
		return &RoleEditResponseStruct{
			Status: ROLE_EDIT_FAILED,
			Msg:    Status[ROLE_EDIT_FAILED],
		}, nil
	}

	Logger.Debugf("EditRole edit success")
	return &RoleEditResponseStruct{
		Status: ROLE_EDIT_SUCCESS,
		Msg:    Status[ROLE_EDIT_SUCCESS],
	}, nil
}

func StartRoleEditServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30026"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/RoleEditThriftService/providers"
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

	handler := &roleeditservice{}
	processor := NewRoleEditThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
