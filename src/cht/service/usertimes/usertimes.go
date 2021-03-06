package usertimes

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	ut "cht/models/usertimes"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	QUERY_USER_TIMES_SUCCESS = 1000
	QUERY_USER_TIMES_FAILED  = 1001
	QUERY_USER_TIMES_EMPTY   = 1002

	UPDATE_USER_TIMES_SUCCESS = 1000
	UPDATE_USER_TIMES_FAILED  = 1001

	INSERT_USER_TIMES_SUCCESS       = 1000
	QUERY_USER_TIMES_DETAILS_FAILED = 1001
	UPDATE_USER_TIMES_TABLE_FAILED  = 1002
	INSERT_USER_TIEMS_FAILED        = 1003

	DELETE_USER_TIMES_SUCCESS = 1000
	DELETE_USER_TIMES_FAILED  = 1001
)

var Query_Stat = map[int]string{
	QUERY_USER_TIMES_SUCCESS: "查询会员登陆次数限制表成功",
	QUERY_USER_TIMES_FAILED:  "查询会员登陆次数限制表失败",
	QUERY_USER_TIMES_EMPTY:   "查询会员登陆次数限制表为空",
}

var Update_Stat = map[int]string{
	UPDATE_USER_TIMES_SUCCESS: "更新会员登陆次数限制表成功",
	QUERY_USER_TIMES_FAILED:   "更新会员登陆次数限制表失败",
}

var Insert_Stat = map[int]string{
	INSERT_USER_TIMES_SUCCESS:       "新增会员登陆次数限制表成功",
	QUERY_USER_TIMES_DETAILS_FAILED: "查询会员登陆次数限制表失败",
	UPDATE_USER_TIMES_TABLE_FAILED:  "更新会员登陆次数限制表失败",
	INSERT_USER_TIEMS_FAILED:        "新增会员登陆次数限制表失败",
}

var Delete_Stat = map[int]string{
	DELETE_USER_TIMES_SUCCESS: "删除会员登陆次数限制表成功",
	DELETE_USER_TIMES_FAILED:  "删除会员登陆次数限制表失败",
}

type usertimesservice struct{}

func (uts *usertimesservice) GetUserTimesDetails(requestObj *UserTimesDetailsRequestStruct) (r *UserTimesDetailsResponseStruct, err error) {
	Logger.Info("GetUserTimesDetails requestObj:", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*UserTimesDetailsRequestStruct)
	utdr := new(ut.UserTimesDetailsRequest)
	utdr.Username = requestObj.GetUsername()
	utdr.Isadmin = requestObj.GetIsadmin()
	utdr.Type = requestObj.GetType()
	utdr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, err := ut.GetUserTimesDetails(utdr)
	if err != nil {
		Logger.Errorf("GetUserTimesDetails failed:%v", err)
		return &UserTimesDetailsResponseStruct{
			Status: QUERY_USER_TIMES_FAILED,
			Msg:    Query_Stat[QUERY_USER_TIMES_FAILED],
		}, nil
	}

	if res == nil {
		Logger.Debugf("GetUserTimesDetails query empty")
		return &UserTimesDetailsResponseStruct{
			Status: QUERY_USER_TIMES_EMPTY,
			Msg:    Query_Stat[QUERY_USER_TIMES_EMPTY],
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
		utds.Type = res.Type
		response.UserTimesDetails = utds
	}
	response.Status = QUERY_USER_TIMES_SUCCESS
	response.Msg = Query_Stat[QUERY_USER_TIMES_SUCCESS]
	Logger.Debugf("GetUserTimesDetails response :%v", response)
	return &response, nil
}

func (uts *usertimesservice) UpdateUserTimes(requestObj *UserTimesUpdateRequestStruct) (r *UserTimesUpdateResponseStruct, err error) {
	Logger.Info("UpdateUserTimes requestObj:", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*UserTimesUpdateRequestStruct)
	utur := new(ut.UserTimesUpdateRequest)
	utur.Username = requestObj.GetUsername()
	utur.IP = requestObj.GetIP()
	utur.Type = requestObj.GetType()
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
	Logger.Infof("InsertUserTimes requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*UserTimesInsertRequestStruct)
	utdr := new(ut.UserTimesDetailsRequest)
	utdr.Username = requestObj.GetUsername()
	utdr.Isadmin = requestObj.GetIsadmin()
	utdr.Type = requestObj.GetType()
	utdr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, err := ut.GetUserTimesDetails(utdr)
	if err != nil {
		Logger.Errorf("InsertUserTimes query failed:%v", err)
		return &UserTimesInsertResponseStruct{
			Status: QUERY_USER_TIMES_DETAILS_FAILED,
			Msg:    Insert_Stat[QUERY_USER_TIMES_DETAILS_FAILED],
		}, nil
	}

	if res != nil {
		utur := new(ut.UserTimesUpdateRequest)
		utur.Username = requestObj.GetUsername()
		utur.IP = requestObj.GetIP()
		utur.Type = requestObj.GetType()
		utur.Isadmin = requestObj.GetIsadmin()
		utur.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
		b := ut.UpdateUserTimes(utur)
		if b == false {
			return &UserTimesInsertResponseStruct{
				Status: UPDATE_USER_TIMES_TABLE_FAILED,
				Msg:    Insert_Stat[UPDATE_USER_TIMES_TABLE_FAILED],
			}, nil
		}
	} else {
		utir := new(ut.UserTimesInsertRequest)
		utir.Username = requestObj.GetUsername()
		utir.IP = requestObj.GetIP()
		utir.Type = requestObj.GetType()
		utir.Isadmin = requestObj.GetIsadmin()
		utir.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
		b := ut.InsertUserTimes(utir)
		if b == false {
			return &UserTimesInsertResponseStruct{
				Status: INSERT_USER_TIEMS_FAILED,
				Msg:    Insert_Stat[INSERT_USER_TIEMS_FAILED],
			}, nil
		}
	}

	return &UserTimesInsertResponseStruct{
		Status: INSERT_USER_TIMES_SUCCESS,
		Msg:    Insert_Stat[INSERT_USER_TIMES_SUCCESS],
	}, nil
}

func (uts *usertimesservice) DeleteUserTimes(requestObj *UserTimesDeleteRequestStruct) (r *UserTimesDeleteResponseStruct, err error) {
	Logger.Info("DeleteUserTimes requestObj:", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*UserTimesDeleteRequestStruct)
	utdr := new(ut.UserTimesDeleteRequest)
	utdr.Username = requestObj.GetUsername()
	utdr.Type = requestObj.GetType()
	utdr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	b := ut.DeleteUserTimes(utdr)
	if b == false {
		Logger.Errorf("DeleteUserTimes failed")
		return &UserTimesDeleteResponseStruct{
			Status: DELETE_USER_TIMES_FAILED,
			Msg:    Delete_Stat[DELETE_USER_TIMES_FAILED],
		}, nil
	}

	return &UserTimesDeleteResponseStruct{
		Status: DELETE_USER_TIMES_SUCCESS,
		Msg:    Delete_Stat[DELETE_USER_TIMES_SUCCESS],
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

	handler := &usertimesservice{}
	processor := NewUserTimesThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
