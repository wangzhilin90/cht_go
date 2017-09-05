package makeborrowservice

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/makeborrow"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"strconv"
)

type borrowservice struct{}

const (
	NOT_DEPOSIT_ACCOUNT     = 1001
	EXCEED_REDIT_LIMIT      = 1002
	GET_REDIT_LIMIT_FAILED  = 1003
	GET_BORROW_MONEY_FAILED = 1004
	GET_FEE_RATE_FAILED     = 1005
	ISSURE_FAILD            = 1006
	ISSURE_SUCCESS          = 0
)

var Status = map[int]string{
	NOT_DEPOSIT_ACCOUNT:     "未开通存管账户",
	GET_REDIT_LIMIT_FAILED:  "获取信用额度失败",
	GET_BORROW_MONEY_FAILED: "获取借款金额失败",
	EXCEED_REDIT_LIMIT:      "信誉额度不够",
	GET_FEE_RATE_FAILED:     "获取利率失败",
	ISSURE_FAILD:            "发标插入失败",
	ISSURE_SUCCESS:          "发标成功",
}

func checkAddCredit(borrow_type int32) bool {
	if borrow_type == 5 {
		return true
	}
	return false
}

func NewMakeBorrowRequest(requestObj *MakeBorrowRequestStruct) *makeborrow.MakeBorrowRequest {
	mbr := new(makeborrow.MakeBorrowRequest)
	mbr.ID = requestObj.GetID()
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
	return mbr
}

func DealTempFunc(mbr *makeborrow.MakeBorrowRequest) *makeborrow.MakeBorrowRequest {
	if mbr.Title == "" {
		mbr.Title = ","
	}
	if mbr.Content == "" {
		mbr.Content = ","
	}
	if mbr.Litpic == "" {
		mbr.Litpic = ","
	}
	if mbr.TimeLimit == 0 {
		mbr.TimeLimit = 1
	}
	if mbr.Account == "" {
		mbr.Account = "3000000.00"
	}
	if mbr.AccountTender == "" {
		mbr.AccountTender = "0.00"
	}
	if mbr.Apr == "" {
		mbr.Apr = "0.0000"
	}
	if mbr.AprAdd == "" {
		mbr.AprAdd = "0.0000"
	}
	if mbr.MortgageFile == "" {
		mbr.MortgageFile = ","
	}
	if mbr.VerifyRemark == "" {
		mbr.VerifyRemark = ","
	}
	if mbr.Pwd == "" {
		mbr.Pwd = ","
	}
	if mbr.LowestAccount == "" {
		mbr.LowestAccount = "50.00"
	}
	if mbr.MostAccount == "" {
		mbr.MostAccount = "0.00"
	}
	if mbr.ValidTime == 0 {
		mbr.ValidTime = 1
	}
	if mbr.Bonus == "" {
		mbr.Bonus = "0.00"
	}
	if mbr.OpenAccount == 0 {
		mbr.OpenAccount = 1
	}
	if mbr.OpenBorrow == 0 {
		mbr.OpenBorrow = 1
	}
	if mbr.OpenTender == 0 {
		mbr.OpenTender = 1
	}
	if mbr.OpenCredit == 0 {
		mbr.OpenCredit = 1
	}
	if mbr.OpenZiliao == 0 {
		mbr.OpenZiliao = 1
	}
	if mbr.Addip == "" {
		mbr.Addip = ","
	}
	if mbr.Secured == "" {
		mbr.Secured = ","
	}
	if mbr.Zhuanrangren == "" {
		mbr.Zhuanrangren = ","
	}
	if mbr.SignDate == "" {
		mbr.SignDate = ","
	}
	if mbr.FeeRate == "" {
		mbr.FeeRate = ","
	}
	if mbr.BorrowName == "" {
		mbr.BorrowName = ","
	}
	return mbr
}

func (bs *borrowservice) makeBorrow(requestObj *MakeBorrowRequestStruct) (r *MakeBorrowResponseStruct, err error) {
	mbr := NewMakeBorrowRequest(requestObj)
	mbr = DealTempFunc(mbr)
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
	guarantor := makeborrow.GetGuarantor(mbr.UserID)
	if guarantor == "" {
		guarantor = ","
	}
	if bIsAddCredit {
		mbr.Secured = guarantor
		mbr.SignDate = ","
		feeRate := requestObj.GetFeeRate()
		if feeRate == "" {
			Logger.Error("get fee rate failed")
			return &MakeBorrowResponseStruct{
				Status: GET_FEE_RATE_FAILED,
				Msg:    Status[GET_FEE_RATE_FAILED],
			}, nil
		}
		formatFeeRate, _ := strconv.ParseFloat(feeRate, 64)
		Logger.Debug("formatFeeRate:", formatFeeRate)
		mbr.FeeRate = fmt.Sprintf("%.4f", formatFeeRate/100.0) //后面整改，保留4位小数
		Logger.Debug("mbr.FeeRate:", mbr.FeeRate)
	} else if _, err := strconv.ParseFloat(guarantor, 64); err != nil {
		mbr.Secured = guarantor
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
	Logger.Debugf("formatLimit %v,formatBorrow %v", formatLimit, formatBorrow)

	if formatLimit-formatBorrow < 0 {
		Logger.Errorf("exceed_redit_limit,limit:%v,borrow:%v", limit, borrow)
		return &MakeBorrowResponseStruct{
			Status: EXCEED_REDIT_LIMIT,
			Msg:    Status[EXCEED_REDIT_LIMIT],
		}, nil
	}

	err = makeborrow.InsertBorrowTbl(mbr)
	if err != nil {
		Logger.Error("InsertBorrowTbl failed", err)
		return &MakeBorrowResponseStruct{
			Status: ISSURE_FAILD,
			Msg:    Status[ISSURE_FAILD],
		}, err
	}
	Logger.Debug("InsertBorrowTbl success")
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
	zkServers := []string{"192.168.8.212:2181", "192.168.8.213:2181", "192.168.8.214:2181"}
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
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &borrowservice{}
	processor := NewMakeBorrowThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
