package usertimes

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	ut "cht/models/usertimes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	QUERY_USER_TIMES_SUCCESS = 1000
	QUERY_USER_TIMES_FAILED  = 1001

	UPDATE_USER_TIMES_SUCCESS = 1000
	UPDATE_USER_TIMES_FAILED  = 1001

	INSERT_USER_TIMES_SUCCESS = 1000
	INSERT_USER_TIEMS_FAILED  = 1001
)

var Query_Stat = map[int]string{
	QUERY_USER_TIMES_SUCCESS: "查询会员登陆次数限制表成功",
	QUERY_USER_TIMES_FAILED:  "查询会员登陆次数限制表失败",
}

var Update_Stat = map[int]string{
	UPDATE_USER_TIMES_SUCCESS: "更新会员登陆次数限制表成功",
	QUERY_USER_TIMES_FAILED:   "更新会员登陆次数限制表失败",
}

var Insert_Stat = map[int]string{
	INSERT_USER_TIMES_SUCCESS: "新增会员登陆次数限制表成功",
	INSERT_USER_TIEMS_FAILED:  "新增会员登陆次数限制表失败",
}

type usertimesservice struct{}

func (uts *usertimesservice) GetUserTimesDetails(requestObj *UserTimesDetailsRequestStruct) (r *UserTimesDetailsResponseStruct, err error) {
	utdr := new(ut.UserTimesDetailsRequest)
	utdr.Username = requestObj.GetUsername()
	utdr.Isadmin = requestObj.GetIsadmin()
	utdr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, err := ut.GetUserTimesDetails(utdr)
	if err != nil {
		Logger.Errorf("GetUserTimesDetails failed:%v", err)
		return &UserTimesDetailsResponseStruct{
			Status: QUERY_USER_TIMES_FAILED,
			Msg:    Query_Stat[QUERY_USER_TIMES_FAILED],
		}, nil
	}

	var response UserTimesDetailsResponseStruct
	if res != nil {
		utds := new(UserTimesDetailsStruct)
		utds.Username = res.Username
		utds.IP = res.IP
		utds.Logintime = res.Logintime
		utds.Times = res.Times
		utds.Isadmin = res.Isadmin
		response.UserTimesDetails = utds
	}
	response.Status = QUERY_USER_TIMES_SUCCESS
	response.Msg = Query_Stat[QUERY_USER_TIMES_SUCCESS]
	Logger.Debugf("GetUserTimesDetails response :%v", response)
	return &response, nil
}

func (uts *usertimesservice) UpdateUserTimes(requestObj *UserTimesUpdateRequestStruct) (r *UserTimesUpdateResponseStruct, err error) {
	utur := new(ut.UserTimesUpdateRequest)
	utur.Username = requestObj.GetUsername()
	utur.IP = requestObj.GetIP()
	utur.Logintime = requestObj.GetLogintime()
	utur.Times = requestObj.GetTimes()
	utur.Isadmin = requestObj.GetIsadmin()
	utur.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	b := ut.UpdateUserTimes(utur)
	if b == false {
		Logger.Errorf("UpdateUserTimes failed")
		return &UserTimesUpdateResponseStruct{
			Status: UPDATE_USER_TIMES_FAILED,
			Msg:    Update_Stat[UPDATE_USER_TIMES_FAILED],
		}, nil
	}

	return &UserTimesUpdateResponseStruct{
		Status: UPDATE_USER_TIMES_SUCCESS,
		Msg:    Update_Stat[UPDATE_USER_TIMES_SUCCESS],
	}, nil
}

func (uts *usertimesservice) InsertUserTimes(requestObj *UserTimesInsertRequestStruct) (r *UserTimesInsertResponseStruct, err error) {
	utir := new(ut.UserTimesInsertRequest)
	utir.Username = requestObj.GetUsername()
	utir.IP = requestObj.GetIP()
	utir.Logintime = requestObj.GetLogintime()
	utir.Times = requestObj.GetTimes()
	utir.Isadmin = requestObj.GetIsadmin()
	utir.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	b := ut.InsertUserTimes(utir)
	if b == false {
		Logger.Errorf("InsertUserTimes failed")
		return &UserTimesInsertResponseStruct{
			Status: INSERT_USER_TIEMS_FAILED,
			Msg:    Insert_Stat[INSERT_USER_TIEMS_FAILED],
		}, nil
	}

	return &UserTimesInsertResponseStruct{
		Status: INSERT_USER_TIMES_SUCCESS,
		Msg:    Insert_Stat[INSERT_USER_TIMES_SUCCESS],
	}, nil
}

/**
 * [StartUserTimesServer 开启会员登陆次数限制表服务]
 * @DateTime 2017-12-14T16:37:21+0800
 */
func StartUserTimesServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30070"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/UserTimesThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode %v failed", servicename, err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &usertimesservice{}
	processor := NewUserTimesThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
