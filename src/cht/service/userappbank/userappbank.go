package userappbank

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	uab "cht/models/userappbank"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type userappbankservice struct{}

const (
	QUERY_USER_BANK_DETAILS_SUCCESS = 1000
	QUERY_USER_BANK_DETAILS_FAILED  = 1001
)

var Details_Stat = map[int]string{
	QUERY_USER_BANK_DETAILS_SUCCESS: "查询app会员银行信息表详情成功",
	QUERY_USER_BANK_DETAILS_FAILED:  "查询app会员银行信息表详情失败",
}

const (
	UPDATE_USER_APP_BANK_SUCCESS = 1000
	UPDATE_USER_APP_BANK_FAILED  = 1001
)

var Update_Stat = map[int]string{
	UPDATE_USER_APP_BANK_SUCCESS: "更新app会员银行信息表详情成功",
	UPDATE_USER_APP_BANK_FAILED:  "更新app会员银行信息表详情失败",
}

const (
	INSERT_USER_APP_BANK_SUCCESS = 1000
	INSERT_USER_APP_BANK_FAILED  = 1001
)

var Insert_Stat = map[int]string{
	INSERT_USER_APP_BANK_SUCCESS: "插入app会员银行信息表详情成功",
	INSERT_USER_APP_BANK_FAILED:  "插入app会员银行信息表详情失败",
}

const (
	DELETE_USER_APP_BANK_SUCCESS = 1000
	DELETE_USER_APP_BANK_FAILED  = 1001
)

var Delete_Stat = map[int]string{
	DELETE_USER_APP_BANK_SUCCESS: "删除app会员银行信息表详情成功",
	DELETE_USER_APP_BANK_FAILED:  "删除app会员银行信息表详情失败",
}

func (uabs *userappbankservice) GetUserAppBankDetails(requestObj *UserAppBankDetailsRequestStruct) (r *UserAppBankDetailsResponseStruct, err error) {
	Logger.Info("GetUserAppBankDetails input param:", requestObj)
	uabdr := new(uab.UserAppBankDetailsRequest)
	uabdr.UserID = requestObj.GetUserID()
	uabdr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, err := uab.GetUserAppBankDetails(uabdr)
	if err != nil {
		return &UserAppBankDetailsResponseStruct{
			Status: QUERY_USER_BANK_DETAILS_FAILED,
			Msg:    Details_Stat[QUERY_USER_BANK_DETAILS_FAILED],
		}, nil
	}

	var response UserAppBankDetailsResponseStruct
	uabd := new(UserAppBankDetailsStruct)
	uabd.ID = res.ID
	uabd.UserID = res.UserID
	uabd.Name = res.Name
	uabd.Account = res.Account
	uabd.Bank = res.Bank
	uabd.Branch = res.Branch
	uabd.Province = res.Province
	uabd.City = res.City
	uabd.Area = res.Area
	uabd.Addtime = res.Addtime
	uabd.Addip = res.Addip
	response.UserAppBankDetails = uabd

	response.Status = QUERY_USER_BANK_DETAILS_SUCCESS
	response.Msg = Details_Stat[QUERY_USER_BANK_DETAILS_SUCCESS]
	Logger.Debugf("GetUserAppBankDetails response:%v", response)
	return &response, nil
}

func (uabs *userappbankservice) UpdateUserAppBank(requestObj *UserAppBankUpdateRequestStruct) (r *UserAppBankUpdateResponseStruct, err error) {
	uabur := new(uab.UserAppBankUpdateRequest)
	uabur.ID = requestObj.GetID()
	uabur.UserID = requestObj.GetUserID()
	uabur.Name = requestObj.GetName()
	uabur.Account = requestObj.GetAccount()
	uabur.Bank = requestObj.GetBank()
	uabur.Branch = requestObj.GetBranch()
	uabur.Province = requestObj.GetProvince()
	uabur.City = requestObj.GetCity()
	uabur.Area = requestObj.GetArea()
	uabur.Addtime = requestObj.GetAddtime()
	uabur.Addip = requestObj.GetAddip()
	uabur.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	b := uab.UpdateUserAppBank(uabur)
	if b == false {
		return &UserAppBankUpdateResponseStruct{
			Status: UPDATE_USER_APP_BANK_FAILED,
			Msg:    Update_Stat[UPDATE_USER_APP_BANK_FAILED],
		}, nil
	}

	return &UserAppBankUpdateResponseStruct{
		Status: UPDATE_USER_APP_BANK_SUCCESS,
		Msg:    Update_Stat[UPDATE_USER_APP_BANK_SUCCESS],
	}, nil
}

func (uabs *userappbankservice) InsertUserAppBank(requestObj *UserAppBankInsertRequestStruct) (r *UserAppBankInsertResponseStruct, err error) {
	Logger.Info("InsertUserAppBank input param:", requestObj)
	uabir := new(uab.UserAppBankInsertRequest)
	uabir.ID = requestObj.GetID()
	uabir.UserID = requestObj.GetUserID()
	uabir.Name = requestObj.GetName()
	uabir.Account = requestObj.GetAccount()
	uabir.Bank = requestObj.GetBank()
	uabir.Branch = requestObj.GetBranch()
	uabir.Province = requestObj.GetProvince()
	uabir.City = requestObj.GetCity()
	uabir.Area = requestObj.GetArea()
	uabir.Addtime = requestObj.GetAddtime()
	uabir.Addip = requestObj.GetAddip()
	uabir.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	b := uab.InsertUserAppBank(uabir)
	if b == false {
		return &UserAppBankInsertResponseStruct{
			Status: INSERT_USER_APP_BANK_FAILED,
			Msg:    Insert_Stat[INSERT_USER_APP_BANK_FAILED],
		}, nil
	}

	return &UserAppBankInsertResponseStruct{
		Status: INSERT_USER_APP_BANK_SUCCESS,
		Msg:    Insert_Stat[INSERT_USER_APP_BANK_SUCCESS],
	}, nil
}

func (uabs *userappbankservice) DeletetUserAppBank(requestObj *UserAppBankDeleteRequestStruct) (r *UserAppBankDeleteResponseStruct, err error) {
	Logger.Info("DeletetUserAppBank input param:", requestObj)
	uabdr := new(uab.UserAppBankDeleteRequest)
	uabdr.UserID = requestObj.GetUserID()
	uabdr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	b := uab.DeletetUserAppBank(uabdr)
	if b == false {
		return &UserAppBankDeleteResponseStruct{
			Status: DELETE_USER_APP_BANK_FAILED,
			Msg:    Delete_Stat[DELETE_USER_APP_BANK_FAILED],
		}, nil
	}

	return &UserAppBankDeleteResponseStruct{
		Status: DELETE_USER_APP_BANK_SUCCESS,
		Msg:    Delete_Stat[DELETE_USER_APP_BANK_SUCCESS],
	}, nil
}

/**
 * [StartUserAppBankServer 开启app会员银行信息表详情服务]
 * @DateTime 2017-12-06T16:55:32+0800
 */
func StartUserAppBankServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30065"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/UserAppBankThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode %v failed", servicename, err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &userappbankservice{}
	processor := NewUserAppBankThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
