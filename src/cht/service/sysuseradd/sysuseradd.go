package sysuseradd

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	useradd "cht/models/sysuseradd"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	INSERT_SYS_USER_SUCCESS = 1000
	INSERT_SYS_USER_FAILED  = 1001
)

var Stat = map[int]string{
	INSERT_SYS_USER_SUCCESS: "添加后台管理用户成功",
	INSERT_SYS_USER_FAILED:  "添加后台管理用户失败",
}

type sysuseraddservice struct{}

func NewSysUserAddRequest(requestObj *SysUserAddRequestStruct) *useradd.SysUserAddRequest {
	suars := new(useradd.SysUserAddRequest)
	suars.Account = requestObj.GetAccount()
	suars.Password = requestObj.GetPassword()
	suars.Realname = requestObj.GetRealname()
	suars.Mobile = requestObj.GetMobile()
	suars.Qq = requestObj.GetQq()
	suars.Status = requestObj.GetStatus()
	suars.RoleID = requestObj.GetRoleID()
	suars.CustomerType = requestObj.GetCustomerType()
	suars.CreateTime = requestObj.GetCreateTime()
	suars.Lastlogintime = requestObj.GetLastlogintime()
	suars.Views = requestObj.GetViews()
	suars.Lastloginip = requestObj.GetLastloginip()
	suars.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	return suars
}

func DealTempFunc(suars *useradd.SysUserAddRequest) *useradd.SysUserAddRequest {
	if suars.Account == "" {
		suars.Account = " "
	}

	if suars.Password == "" {
		suars.Password = " "
	}

	if suars.Realname == "" {
		suars.Realname = " "
	}

	if suars.Mobile == "" {
		suars.Mobile = " "
	}

	if suars.Qq == "" {
		suars.Qq = " "
	}

	if suars.Lastloginip == "" {
		suars.Lastloginip = " "
	}

	if suars.CustomerType == 0 {
		suars.CustomerType = 668
	}

	return suars
}

func (suas *sysuseraddservice) AddSysUser(requestObj *SysUserAddRequestStruct) (r *SysUserAddResponseStruct, err error) {
	Logger.Debugf("AddSysUser input param:%v", requestObj)
	suars := NewSysUserAddRequest(requestObj)
	suars = DealTempFunc(suars)
	b := useradd.AddSysUser(suars)
	if b == false {
		return &SysUserAddResponseStruct{
			Status: INSERT_SYS_USER_FAILED,
			Msg:    Stat[INSERT_SYS_USER_FAILED],
		}, nil
	}

	return &SysUserAddResponseStruct{
		Status: INSERT_SYS_USER_SUCCESS,
		Msg:    Stat[INSERT_SYS_USER_SUCCESS],
	}, nil
}

/**
 * [StartSysUserAddServer 添加后台管理用户服务]
 * @DateTime 2017-10-17T16:11:09+0800
 */
func StartSysUserAddServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30032"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/SysUserAddThriftService/providers"
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

	handler := &sysuseraddservice{}
	processor := NewSysUserAddThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
