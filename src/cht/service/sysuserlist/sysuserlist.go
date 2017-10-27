package sysuserlist

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	userlist "cht/models/sysuserlist"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	QUERY_SYS_USER_LIST_SUCCESS = 1000
	QUERY_SYS_USER_LIST_FAILED  = 1001
)

var Stat = map[int]string{
	QUERY_SYS_USER_LIST_SUCCESS: "获取后台管理员列表成功",
	QUERY_SYS_USER_LIST_FAILED:  "获取后台管理员列表失败",
}

type sysuserlistservice struct{}

func (suls *sysuserlistservice) GetSysUserList(requestObj *SysUserListRequestStruct) (r *SysUserListResponseStruct, err error) {
	sulr := new(userlist.SysUserListRequest)
	sulr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	res, err := userlist.GetSysUserList(sulr)
	if err != nil {
		return &SysUserListResponseStruct{
			Status: QUERY_SYS_USER_LIST_FAILED,
			Msg:    Stat[QUERY_SYS_USER_LIST_FAILED],
		}, nil
	}

	var response SysUserListResponseStruct
	for _, v := range res {
		suds := new(SysUserDetailsStruct)
		suds.ID = v.ID
		suds.RoleID = v.RoleID
		suds.Account = v.Account
		suds.Realname = v.Realname
		suds.Password = v.Password
		suds.Mobile = v.Mobile
		suds.Qq = v.Qq
		suds.Lastloginip = v.Lastloginip
		suds.Lastlogintime = v.Lastlogintime
		suds.CreateTime = v.CreateTime
		suds.Status = v.Status
		suds.Views = v.Views
		suds.CustomerType = v.CustomerType
		response.SysUserList = append(response.SysUserList, suds)
	}
	response.Status = QUERY_SYS_USER_LIST_SUCCESS
	response.Msg = Stat[QUERY_SYS_USER_LIST_SUCCESS]
	Logger.Debugf("GetSysUserList response:%v", response)
	return &response, nil
}

/**
 * [StartSysUserListServer 后台管理员列表服务]
 * @DateTime 2017-10-19T10:41:58+0800
 */
func StartSysUserListServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30036"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/SysUserListThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &sysuserlistservice{}
	processor := NewSysUserListThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
