package rolerightset

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/rolerightset"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

type rolerightsetservice struct{}

const (
	ROLE_RIGHT_SUCCESS = 1000
	ROLE_RIGHT_FAILED  = 1001
)

var Status = map[int]string{
	ROLE_RIGHT_SUCCESS: "角色权限修改成功",
	ROLE_RIGHT_FAILED:  "角色权限修改失败",
}

func (rrss *rolerightsetservice) SetRoleRight(requestObj *RoleRightSetRequestStruct) (r *RoleRightSetResponseStruct, err error) {
	Logger.Infof("SetRoleRight requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*RoleRightSetRequestStruct)
	rrsrs := new(rolerightset.RoleRightSetRequestStruct)
	rrsrs.RoleID = requestObj.GetRoleID()
	rrsrs.PowerConfig = requestObj.GetPowerConfig()
	rrsrs.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	b := rolerightset.SetRoleRight(rrsrs)
	if b == false {
		Logger.Debugf("SetRoleRight set failed")
		return &RoleRightSetResponseStruct{
			Status: ROLE_RIGHT_FAILED,
			Msg:    Status[ROLE_RIGHT_FAILED],
		}, nil
	}

	Logger.Debugf("SetRoleRight set success")
	return &RoleRightSetResponseStruct{
		Status: ROLE_RIGHT_SUCCESS,
		Msg:    Status[ROLE_RIGHT_SUCCESS],
	}, nil
}

func StartRoleEditServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30027"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/RoleRightSetThriftService/providers"
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

	handler := &rolerightsetservice{}
	processor := NewRoleRightSetThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
