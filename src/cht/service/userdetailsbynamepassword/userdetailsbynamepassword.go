package userdetailsbynamepassword

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/userdetailsbynamepassword"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

type userdetailsbynamepasswordservice struct{}

const (
	QUERY_USER_DETAILS_SUCCESS = 1000
	QUERY_USER_DETAILS_FAILED  = 1001
	QUERY_USER_DETAILS_EMPTY   = 1002
)

var Status = map[int]string{
	QUERY_USER_DETAILS_SUCCESS: "查询后台用户详情成功",
	QUERY_USER_DETAILS_FAILED:  "查询后台用户详情失败",
	QUERY_USER_DETAILS_EMPTY:   "查询后台用户详情为空",
}

func (uds *userdetailsbynamepasswordservice) GetUseDetailsrByNamePassword(requestObj *UserDetailsByNamePasswordRequestStruct) (r *UserDetailsByNamePasswordResponseStruct, err error) {
	Logger.Info("GetUseDetailsrByNamePassword requestObj:", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*UserDetailsByNamePasswordRequestStruct)
	udbr := new(userdetailsbynamepassword.UserDetailsByNamePasswordRequest)
	udbr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	udbr.Name = requestObj.GetName()
	udbr.Password = requestObj.GetPassword()
	userDetails, err := userdetailsbynamepassword.GetUseDetailsrByNamePassword(udbr)
	if err != nil {
		Logger.Errorf("GetUseDetailsrByNamePassword failed %v", err)
		return &UserDetailsByNamePasswordResponseStruct{
			Status: QUERY_USER_DETAILS_FAILED,
			Msg:    Status[QUERY_USER_DETAILS_FAILED],
		}, nil
	}

	if userDetails == nil {
		Logger.Debugf("GetUseDetailsrByNamePassword query empty")
		return &UserDetailsByNamePasswordResponseStruct{
			Status: QUERY_USER_DETAILS_EMPTY,
			Msg:    Status[QUERY_USER_DETAILS_EMPTY],
		}, nil
	}

	var response UserDetailsByNamePasswordResponseStruct
	if userDetails != nil {
		sud := new(SysUserDetailsStruct)
		sud.ID = userDetails.ID
		sud.RoleID = userDetails.RoleID
		sud.Account = userDetails.Account
		sud.Realname = userDetails.Realname
		sud.Password = userDetails.Password
		sud.Mobile = userDetails.Mobile
		sud.Qq = userDetails.Qq
		sud.Lastloginip = userDetails.Lastloginip
		sud.Lastlogintime = userDetails.Lastlogintime
		sud.CreateTime = userDetails.CreateTime
		sud.Status = userDetails.Status
		sud.Views = userDetails.Views
		sud.CustomerType = userDetails.CustomerType
		response.SysUserDetails = sud
	}
	response.Status = QUERY_USER_DETAILS_SUCCESS
	response.Msg = Status[QUERY_USER_DETAILS_SUCCESS]

	Logger.Debugf("GetUseDetailsrByNamePassword query res:%v", userDetails)
	return &response, nil
}

func StartUseDetailsrServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30020"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/UserDetailsByNamePasswordThriftService/providers"
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

	handler := &userdetailsbynamepasswordservice{}
	processor := NewUserDetailsByNamePasswordThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
