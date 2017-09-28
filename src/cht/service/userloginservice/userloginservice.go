package userloginservice

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/userlogin"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	RETRY_TOO_MUCH    = 1001
	ACCOUNT_NOT_EXIST = 1002
	ACCOUNT_LOCKED    = 1003
	VERIFY_FAILED     = 1004
	VERIFY_PASS       = 1000
)

var Status = map[int]string{
	RETRY_TOO_MUCH:    "密码重试次数太多",
	ACCOUNT_NOT_EXIST: "帐号不存在，请重新输入",
	ACCOUNT_LOCKED:    "您的帐号被锁定了，请联系我们",
	VERIFY_FAILED:     "密码错误",
	VERIFY_PASS:       "密码验证通过",
}

type UserLoginService struct{}

func (uls *UserLoginService) GetUserLoginInfo(requestObj *UserlLoginRequestStruct) (r *UserLoginResponseStruct, err error) {
	ulr := new(userlogin.UserlLoginRequest)
	ulr.Username = requestObj.GetUsername()
	ulr.Password = requestObj.GetPassword()
	ulr.IP = requestObj.GetIP()
	ulr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	Logger.Debug("GetUserLoginInfo input param:", requestObj)

	var v UserLoginResponseStruct
	var max_retry_times int32 = 5
	times, err := userlogin.GetLoginFailedTimes(ulr)
	if times >= max_retry_times {
		v = UserLoginResponseStruct{
			Status: RETRY_TOO_MUCH,
			Msg:    Status[RETRY_TOO_MUCH],
		}
		return &v, nil
	}

	res, bExists, _ := userlogin.CheckLoginUserExists(ulr)
	if bExists == false {
		v = UserLoginResponseStruct{
			Status: ACCOUNT_NOT_EXIST,
			Msg:    Status[ACCOUNT_NOT_EXIST],
		}
		return &v, nil
	}

	if res.Islock == true {
		v = UserLoginResponseStruct{
			UserID: res.ID,
			Status: ACCOUNT_LOCKED,
			Msg:    Status[ACCOUNT_LOCKED],
		}
		return &v, nil
	}
	bl := userlogin.Checkpassword(ulr)
	if bl == false {
		b2, _ := userlogin.CheckUserTimesTbExist(ulr)
		if b2 == false {
			userlogin.InsertUserTimesTb(ulr)
		} else {
			userlogin.UpdateUserTimesTb(ulr)
		}
		v = UserLoginResponseStruct{
			UserID: res.ID,
			Status: VERIFY_FAILED,
			Msg:    Status[VERIFY_FAILED],
			Flag:   max_retry_times - times - 1,
		}
	} else {
		userlogin.DeleteUserTimesTb(ulr)
		v = UserLoginResponseStruct{
			UserID: res.ID,
			Status: VERIFY_PASS,
			Msg:    Status[VERIFY_PASS],
			Flag:   max_retry_times,
		}
	}
	Logger.Debug("GetUserLoginInfo return value:", v)
	return &v, nil
}

/**
 * [StartUserLoginServer 开启用户登录服务]
 * @param    listenAddr string [监听ip和端口]30002
 * @DateTime 2017-08-30T10:38:44+0800
 */
func StartUserLoginServer() {
	zkServers := []string{"192.168.8.208:2181"}
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30002"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/UserLoginThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &UserLoginService{}
	processor := NewUserLoginThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
