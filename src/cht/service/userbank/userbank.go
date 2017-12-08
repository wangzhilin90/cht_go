package userbank

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	ub "cht/models/userbank"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type userbankservice struct{}

const (
	QUERY_USER_BANK_DETAILS_SUCCESS = 1000
	QUERY_USER_BANK_DETAILS_FAILED  = 1001
)

var Details_Stat = map[int]string{
	QUERY_USER_BANK_DETAILS_SUCCESS: "查询会员银行信息表详情成功",
	QUERY_USER_BANK_DETAILS_FAILED:  "查询会员银行信息表详情失败",
}

const (
	UPDATE_USER_BANK_SUCCESS = 1000
	UPDATE_USER_BANK_FAILED  = 1001
)

var Update_Stat = map[int]string{
	UPDATE_USER_BANK_SUCCESS: "更新会员银行信息表详情成功",
	UPDATE_USER_BANK_FAILED:  "更新会员银行信息表详情失败",
}

const (
	INSERT_USER_BANK_SUCCESS = 1000
	INSERT_USER_BANK_FAILED  = 1001
)

var Insert_Stat = map[int]string{
	INSERT_USER_BANK_SUCCESS: "插入会员银行信息表详情成功",
	INSERT_USER_BANK_FAILED:  "插入会员银行信息表详情失败",
}

const (
	COUNT_USER_BANK_SUCCESS = 1000
	COUNT_USER_BANK_FAILED  = 1001
)

var Count_Stat = map[int]string{
	COUNT_USER_BANK_SUCCESS: "获取会员银行会员信息条目数成功",
	COUNT_USER_BANK_FAILED:  "获取会员银行会员信息条目数失败",
}

func (ubs *userbankservice) GetUserBankDetails(requestObj *UserBankDetailsRequestStruct) (r *UserBankDetailsResponseStruct, err error) {
	Logger.Info("GetUserBankDetails input param:", requestObj)
	ubd := new(ub.UserBankDetailsRequest)
	ubd.UserID = requestObj.GetUserID()
	ubd.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, err := ub.GetUserBankDetails(ubd)
	if err != nil {
		return &UserBankDetailsResponseStruct{
			Status: QUERY_USER_BANK_DETAILS_FAILED,
			Msg:    Details_Stat[QUERY_USER_BANK_DETAILS_FAILED],
		}, nil
	}

	var response UserBankDetailsResponseStruct
	ubds := new(UserBankDetailsStruct)
	ubds.ID = res.ID
	ubds.UserID = res.UserID
	ubds.Name = res.Name
	ubds.Account = res.Account
	ubds.Bank = res.Bank
	ubds.Branch = res.Branch
	ubds.Province = res.Province
	ubds.City = res.City
	ubds.Area = res.Area
	ubds.Addtime = res.Addtime
	ubds.Addip = res.Addip
	response.UserBankDetails = ubds

	response.Status = QUERY_USER_BANK_DETAILS_SUCCESS
	response.Msg = Details_Stat[QUERY_USER_BANK_DETAILS_SUCCESS]
	Logger.Debugf("GetUserBankDetails response:%v", response)
	return &response, nil
}

func (ubs *userbankservice) UpdateUserBank(requestObj *UserBankUpdateRequestStruct) (r *UserBankUpdateResponseStruct, err error) {
	ubur := new(ub.UserBankUpdateRequest)
	ubur.ID = requestObj.GetID()
	ubur.UserID = requestObj.GetUserID()
	ubur.Name = requestObj.GetName()
	ubur.Account = requestObj.GetAccount()
	ubur.Bank = requestObj.GetBank()
	ubur.Branch = requestObj.GetBranch()
	ubur.Province = requestObj.GetProvince()
	ubur.City = requestObj.GetCity()
	ubur.Area = requestObj.GetArea()
	ubur.Addtime = requestObj.GetAddtime()
	ubur.Addip = requestObj.GetAddip()
	ubur.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	b := ub.UpdateUserBank(ubur)
	if b == false {
		return &UserBankUpdateResponseStruct{
			Status: UPDATE_USER_BANK_FAILED,
			Msg:    Update_Stat[UPDATE_USER_BANK_FAILED],
		}, nil
	}

	return &UserBankUpdateResponseStruct{
		Status: UPDATE_USER_BANK_SUCCESS,
		Msg:    Update_Stat[UPDATE_USER_BANK_SUCCESS],
	}, nil
}

func (ubs *userbankservice) InsertUserBank(requestObj *UserBankInsertRequestStruct) (r *UserBankInsertResponseStruct, err error) {
	Logger.Info("InsertUserBank input param:", requestObj)
	ubir := new(ub.UserBankInsertRequest)
	ubir.ID = requestObj.GetID()
	ubir.UserID = requestObj.GetUserID()
	ubir.Name = requestObj.GetName()
	ubir.Account = requestObj.GetAccount()
	ubir.Bank = requestObj.GetBank()
	ubir.Branch = requestObj.GetBranch()
	ubir.Province = requestObj.GetProvince()
	ubir.City = requestObj.GetCity()
	ubir.Area = requestObj.GetArea()
	ubir.Addtime = requestObj.GetAddtime()
	ubir.Addip = requestObj.GetAddip()
	ubir.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	b := ub.InsertUserBank(ubir)
	if b == false {
		return &UserBankInsertResponseStruct{
			Status: INSERT_USER_BANK_FAILED,
			Msg:    Insert_Stat[INSERT_USER_BANK_FAILED],
		}, nil
	}

	return &UserBankInsertResponseStruct{
		Status: INSERT_USER_BANK_SUCCESS,
		Msg:    Insert_Stat[INSERT_USER_BANK_SUCCESS],
	}, nil
}

func (ubs *userbankservice) GetUserBankNum(requestObj *UserBankCountRequestStruct) (r *UserBankCountResponseStruct, err error) {
	Logger.Info("GetUserBankNum input param:", requestObj)
	ubcr := new(ub.UserBankCountRequest)
	ubcr.UserID = requestObj.GetUserID()
	ubcr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	num, err := ub.GetUserBankNum(ubcr)
	if err != nil {
		return &UserBankCountResponseStruct{
			Status: COUNT_USER_BANK_FAILED,
			Msg:    Count_Stat[COUNT_USER_BANK_FAILED],
		}, nil
	}

	return &UserBankCountResponseStruct{
		Status:   COUNT_USER_BANK_SUCCESS,
		Msg:      Count_Stat[COUNT_USER_BANK_SUCCESS],
		TotalNum: num,
	}, nil
}

/**
 * [StartUserAppBankServer 开启会员银行信息表详情服务]
 * @DateTime 2017-12-08T16:55:32+0800
 */
func StartUserBankServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30066"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/UserBankThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode %v failed", servicename, err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &userbankservice{}
	processor := NewUserBankThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
