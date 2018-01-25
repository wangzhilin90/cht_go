package sysuserlist

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	userlist "cht/models/sysuserlist"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	QUERY_SYS_USER_LIST_SUCCESS = 1000
	QUERY_SYS_USER_LIST_FAILED  = 1001
	QUERY_SYS_USER_LIST_EMPTY   = 1002
)

var Stat = map[int]string{
	QUERY_SYS_USER_LIST_SUCCESS: "获取后台管理员列表成功",
	QUERY_SYS_USER_LIST_FAILED:  "获取后台管理员列表失败",
	QUERY_SYS_USER_LIST_EMPTY:   "获取后台管理员列表为空",
}

type sysuserlistservice struct{}

func (suls *sysuserlistservice) GetSysUserList(requestObj *SysUserListRequestStruct) (r *SysUserListResponseStruct, err error) {
	Logger.Infof("GetSysUserList requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*SysUserListRequestStruct)
	sulr := new(userlist.SysUserListRequest)
	sulr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	res, err := userlist.GetSysUserList(sulr)
	if err != nil {
		Logger.Errorf("GetSysUserList get sys user list failed:%v", err)
		return &SysUserListResponseStruct{
			Status: QUERY_SYS_USER_LIST_FAILED,
			Msg:    Stat[QUERY_SYS_USER_LIST_FAILED],
		}, nil
	}

	if res == nil {
		Logger.Debugf("GetSysUserList query empty")
		return &SysUserListResponseStruct{
			Status: QUERY_SYS_USER_LIST_EMPTY,
			Msg:    Stat[QUERY_SYS_USER_LIST_EMPTY],
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

	handler := &sysuserlistservice{}
	processor := NewSysUserListThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
