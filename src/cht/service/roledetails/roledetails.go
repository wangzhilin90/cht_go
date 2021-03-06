package roledetails

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/roledetails"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

type roledetailsservice struct{}

const (
	ROLE_DETAILS_SUCCESS = 1000
	ROLE_DETAILS_FAILED  = 1001
	ROLE_DETAILS_EMPTY   = 1002
)

var Status = map[int]string{
	ROLE_DETAILS_SUCCESS: "角色详情获取成功",
	ROLE_DETAILS_FAILED:  "角色详情获取失败",
	ROLE_DETAILS_EMPTY:   "角色详情获取为空",
}

func (rdss *roledetailsservice) GetRoleDetails(requestObj *RoleDetailsRequestStruct) (r *RoleDetailsResponseStruct, err error) {
	Logger.Infof("GetRoleDetails requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*RoleDetailsRequestStruct)
	rdrs := new(roledetails.RoleDetailsRequestStruct)
	rdrs.RoleID = requestObj.GetRoleID()
	res, err := roledetails.GetRoleDetails(rdrs)
	if err != nil {
		Logger.Errorf("GetRoleDetails query failed:%v", err)
		return &RoleDetailsResponseStruct{
			Status: ROLE_DETAILS_FAILED,
			Msg:    Status[ROLE_DETAILS_FAILED],
		}, nil
	}

	if res == nil {
		Logger.Debugf("GetRoleDetails query empty")
		return &RoleDetailsResponseStruct{
			Status: ROLE_DETAILS_EMPTY,
			Msg:    Status[ROLE_DETAILS_EMPTY],
		}, nil
	}

	var response RoleDetailsResponseStruct
	if res != nil {
		rds := new(RoleDetailsStruct)
		rds.ID = res.ID
		rds.RoleName = res.RoleName
		rds.Remark = res.Remark
		rds.PowerConfig = res.PowerConfig
		rds.CreateTime = res.CreateTime
		response.RoleDetails = rds
	}

	response.Status = ROLE_DETAILS_SUCCESS
	response.Msg = Status[ROLE_DETAILS_SUCCESS]

	Logger.Debugf("GetRoleDetails response:%v", response)
	return &response, nil
}

func StartRoleDetailsServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30025"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/RoleDetailsThriftService/providers"
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

	handler := &roledetailsservice{}
	processor := NewRoleDetailsThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
