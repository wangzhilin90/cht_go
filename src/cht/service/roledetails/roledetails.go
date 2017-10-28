package roledetails

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/roledetails"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type roledetailsservice struct{}

const (
	ROLE_DETAILS_SUCCESS = 1000
	ROLE_DETAILS_FAILED  = 1001
)

var Status = map[int]string{
	ROLE_DETAILS_SUCCESS: "角色添加成功",
	ROLE_DETAILS_FAILED:  "角色添加失败",
}

func (rdss *roledetailsservice) GetRoleDetails(requestObj *RoleDetailsRequestStruct) (r *RoleDetailsResponseStruct, err error) {
	rdrs := new(roledetails.RoleDetailsRequestStruct)
	rdrs.RoleID = requestObj.GetRoleID()
	res, err := roledetails.GetRoleDetails(rdrs)
	if err != nil {
		Logger.Debugf("GetRoleDetails query failed", err)
		return &RoleDetailsResponseStruct{
			Status: ROLE_DETAILS_FAILED,
			Msg:    Status[ROLE_DETAILS_FAILED],
		}, nil
	}

	rds := new(RoleDetailsStruct)
	rds.ID = res.ID
	rds.RoleName = res.RoleName
	rds.Remark = res.Remark
	rds.PowerConfig = res.PowerConfig
	rds.CreateTime = res.CreateTime

	var response RoleDetailsResponseStruct
	response.Status = ROLE_DETAILS_SUCCESS
	response.Msg = Status[ROLE_DETAILS_SUCCESS]
	response.RoleDetails = rds
	Logger.Debugf("GetRoleDetails return value %v", response)
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

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &roledetailsservice{}
	processor := NewRoleDetailsThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
