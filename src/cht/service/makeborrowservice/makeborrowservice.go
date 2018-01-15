package makeborrowservice

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/makeborrow"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"strconv"
	"time"
)

type borrowservice struct{}

const (
	NOT_DEPOSIT_ACCOUNT         = 1001
	EXCEED_REDIT_LIMIT          = 1002
	GET_REDIT_LIMIT_FAILED      = 1003
	GET_BORROW_MONEY_FAILED     = 1004
	GET_FEE_RATE_FAILED         = 1005
	ISSURE_FAILD                = 1006
	QUERY_REVIEW_ACCOUNT_FAILED = 1007
	ISSURE_SUCCESS              = 1000
)

var Status = map[int]string{
	NOT_DEPOSIT_ACCOUNT:         "未开通存管账户",
	GET_REDIT_LIMIT_FAILED:      "获取信用额度失败",
	GET_BORROW_MONEY_FAILED:     "获取借款金额失败",
	EXCEED_REDIT_LIMIT:          "信誉额度不够",
	GET_FEE_RATE_FAILED:         "获取利率失败",
	ISSURE_FAILD:                "发标插入失败",
	QUERY_REVIEW_ACCOUNT_FAILED: "查询用户发标待审金额失败",
	ISSURE_SUCCESS:              "发标成功",
}

func checkAddCredit(borrow_type int32) bool {
	if borrow_type == 5 {
		return true
	}
	return false
}

func NewMakeBorrowRequest(requestObj *MakeBorrowRequestStruct) *makeborrow.Borrow {
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*MakeBorrowRequestStruct)
	mbr := new(makeborrow.Borrow)
	mbr.ID, _ = makeborrow.GetLatestBorrowID()
	mbr.BorrowType = requestObj.GetBorrowType()
	mbr.UserID = requestObj.GetUserID()
	mbr.Title = requestObj.GetTitle()
	mbr.Content = requestObj.GetContent()
	mbr.Litpic = requestObj.GetLitpic()
	mbr.BorrowUse = requestObj.GetBorrowUse()
	mbr.IsDatetype = requestObj.GetIsDatetype()
	mbr.TimeLimit = requestObj.GetTimeLimit()
	mbr.Style = requestObj.GetStyle()
	mbr.Account = requestObj.GetAccount()
	mbr.AccountTender = requestObj.GetAccountTender()
	mbr.Apr = requestObj.GetApr()
	mbr.AprAdd = requestObj.GetAprAdd()
	mbr.MortgageFile = requestObj.GetMortgageFile()
	mbr.IsDxb = requestObj.GetIsDxb()
	mbr.Pwd = requestObj.GetPwd()
	mbr.LowestAccount = requestObj.GetLowestAccount()
	mbr.MostAccount = requestObj.GetMostAccount()
	mbr.ValidTime = requestObj.GetValidTime()
	mbr.Award = requestObj.GetAward()
	mbr.Bonus = requestObj.GetBonus()
	mbr.IsFalse = requestObj.GetIsFalse()
	mbr.OpenAccount = requestObj.GetOpenAccount()
	mbr.OpenBorrow = requestObj.GetOpenBorrow()
	mbr.OpenTender = requestObj.GetOpenTender()
	mbr.OpenCredit = requestObj.GetOpenCredit()
	mbr.OpenZiliao = requestObj.GetOpenZiliao()
	mbr.Material = requestObj.GetMaterial()
	mbr.Addtime = requestObj.GetAddtime()
	mbr.Addip = requestObj.GetAddip()
	mbr.Status = requestObj.GetStatus()
	mbr.RutenAllnumber = requestObj.GetRutenAllnumber()
	mbr.RutenNumber = requestObj.GetRutenNumber()
	mbr.VerifyUser = requestObj.GetVerifyUser()
	mbr.VerifyTime = requestObj.GetVerifyTime()
	mbr.VerifyRemark = requestObj.GetVerifyRemark()
	mbr.ReviewUser = requestObj.GetReviewUser()
	mbr.ReviewTimeLocal = requestObj.GetReviewTimeLocal()
	mbr.ReviewTime = requestObj.GetReviewTime()
	mbr.Secured = requestObj.GetSecured()
	mbr.Zhuanrangren = requestObj.GetZhuanrangren()
	mbr.Huodong = requestObj.GetHuodong()
	mbr.SignDate = requestObj.GetSignDate()
	mbr.Subledger = requestObj.GetSubledger()
	mbr.RepaySign = requestObj.GetRepaySign()
	mbr.AutoTenderLock = requestObj.GetAutoTenderLock()
	mbr.IsAuto = requestObj.GetIsAuto()
	mbr.IsCheck = requestObj.GetIsCheck()
	mbr.ReviewLock = requestObj.GetReviewLock()
	mbr.FeeRate = requestObj.GetFeeRate()
	mbr.BorrowName = requestObj.GetBorrowName()
	mbr.VipLevelLimit = requestObj.GetVipLevelLimit()
	return mbr
}

func (bs *borrowservice) MakeBorrow(requestObj *MakeBorrowRequestStruct) (r *MakeBorrowResponseStruct, err error) {
	mbr := NewMakeBorrowRequest(requestObj)
	Logger.Debug("requestObj:", requestObj)
	Logger.Debug("NewMakeBorrowRequest:", mbr)
	bIsDeposit := makeborrow.CheckDepositAccount(mbr.UserID)
	if bIsDeposit == false {
		return &MakeBorrowResponseStruct{
			Status: NOT_DEPOSIT_ACCOUNT,
			Msg:    Status[NOT_DEPOSIT_ACCOUNT],
		}, nil
	}

	bIsAddCredit := checkAddCredit(mbr.BorrowType)

	if bIsAddCredit {
		mbr.SignDate = ""
	} else if mbr.Secured != "" {
		_, err := strconv.ParseFloat(mbr.Secured, 64)
		if err == nil { //担保方是数字字符串，需要过滤掉
			mbr.Secured = ""
		}
	}

	limit, err := makeborrow.GetCreditLimit(mbr.UserID)
	if err != nil {
		Logger.Error("get credit limit failed", err)
		return &MakeBorrowResponseStruct{
			Status: GET_REDIT_LIMIT_FAILED,
			Msg:    Status[GET_REDIT_LIMIT_FAILED],
		}, nil
	}
	formatLimit, _ := strconv.ParseFloat(limit, 64)

	borrow := requestObj.GetAccount()
	if borrow == "" {
		Logger.Error("get borrow money failed")
		return &MakeBorrowResponseStruct{
			Status: GET_BORROW_MONEY_FAILED,
			Msg:    Status[GET_BORROW_MONEY_FAILED],
		}, nil
	}
	formatBorrow, _ := strconv.ParseFloat(borrow, 64)

	review_account, err := makeborrow.GetReviewAccount(mbr.UserID)
	if err != nil {
		Logger.Errorf("makeBorrow query review_account failed:", err)
		return &MakeBorrowResponseStruct{
			Status: QUERY_REVIEW_ACCOUNT_FAILED,
			Msg:    Status[QUERY_REVIEW_ACCOUNT_FAILED],
		}, nil
	}
	format_account, _ := strconv.ParseFloat(review_account, 64)
	Logger.Debugf("formatLimit %v,formatBorrow %v,format_account %v", formatLimit, formatBorrow, format_account)

	//如果可用信用额度减去用户发标待审中的金额小于做标金额，返回额度不够错误
	if formatLimit-format_account-formatBorrow < 0 {
		Logger.Errorf("exceed_redit_limit,limit:%v,borrow:%v", limit, borrow)
		return &MakeBorrowResponseStruct{
			Status: EXCEED_REDIT_LIMIT,
			Msg:    Status[EXCEED_REDIT_LIMIT],
		}, nil
	}

	_, err = makeborrow.InsertBorrowTbl(mbr)
	if err != nil {
		Logger.Error("InsertBorrowTbl failed", err)
		return &MakeBorrowResponseStruct{
			Status: ISSURE_FAILD,
			Msg:    Status[ISSURE_FAILD],
		}, err
	}
	Logger.Debugf("InsertBorrowTbl success status: %v msg:%v", ISSURE_SUCCESS, Status[ISSURE_SUCCESS])
	return &MakeBorrowResponseStruct{
		Status: ISSURE_SUCCESS,
		Msg:    Status[ISSURE_SUCCESS],
	}, nil
}

/**
 * [StartUpdatePasswdsServer 开启发标服务]
 * @DateTime 2017-08-24T15:19:45+0800
 */
func StartMakeBorrowServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30006"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/MakeBorrowThriftService/providers"
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

	handler := &borrowservice{}
	processor := NewMakeBorrowThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
