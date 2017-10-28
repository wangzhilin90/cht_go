package sysuserdetails

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	userdetails "cht/models/sysuserdetails"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	QUERY_SYS_USER_DETAILS_SUCCESS = 1000
	QUERY_SYS_USER_DETAILS_FAILED  = 1001
)

var Stat = map[int]string{
	QUERY_SYS_USER_DETAILS_SUCCESS: "查询后台管理员详情成功",
	QUERY_SYS_USER_DETAILS_FAILED:  "查询后台管理员详情失败",
}

type sysuserdetailsservice struct{}

func (suds *sysuserdetailsservice) GetSysUserDetails(requestObj *SysUserDetailsRequestStruct) (r *SysUserDetailsResponseStruct, err error) {
	sudr := new(userdetails.SysUserDetailsRequest)
	sudr.UserID = requestObj.GetUserID()
	sudr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	res, err := userdetails.GetSysUserDetails(sudr)
	if err != nil {
		return &SysUserDetailsResponseStruct{
			Status: QUERY_SYS_USER_DETAILS_FAILED,
			Msg:    Stat[QUERY_SYS_USER_DETAILS_FAILED],
		}, nil
	}

	var response SysUserDetailsResponseStruct
	sud := new(SysUserDetailsStruct)
	sud.ID = res.ID
	sud.RoleID = res.RoleID
	sud.Account = res.Account
	sud.Realname = res.Realname
	sud.Password = res.Password
	sud.Mobile = res.Mobile
	sud.Qq = res.Qq
	sud.Lastloginip = res.Lastloginip
	sud.Lastlogintime = res.Lastlogintime
	sud.CreateTime = res.CreateTime
	sud.Status = res.Status
	sud.Views = res.Views
	sud.CustomerType = res.CustomerType

	response.SysUserDetails = sud
	response.Status = QUERY_SYS_USER_DETAILS_SUCCESS
	response.Msg = Stat[QUERY_SYS_USER_DETAILS_SUCCESS]
	Logger.Debugf("GetSysUserDetails return value:%v", response)
	return &response, nil
}

/**
 * [StartSysUserDeleteServer 后台管理员详情服务]
 * @DateTime 2017-10-18T14:58:03+0800
 */
func StartSysUserDetailsServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30034"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/SysUserDetailsThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode %v failed", servicename, err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &sysuserdetailsservice{}
	processor := NewSysUserDetailsThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
