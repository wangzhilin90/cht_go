package borrowuserdetails

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/borrower"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"strconv"
	"time"
)

type borrowerservice struct{}

const (
	QUERY_BORROW_INFO_SUCCESS = 1000
	QUERY_USER_ID_FAILED      = 1001
)

var Status = map[int]string{
	QUERY_BORROW_INFO_SUCCESS: "查询借款人信息成功",
	QUERY_USER_ID_FAILED:      "无此借款人!",
}

type BorrowerInfoRequest struct {
	Name                 string
	ChengHuiTongTraceLog string
}

func (bs *borrowerservice) GetBorrowUserDetails(requestObj *BorrowUserDetailsRequestStruct) (r *BorrowUserDetailsResponseStruct, err error) {
	bir := new(borrower.BorrowerInfoRequest)
	bir.Name = requestObj.GetName()
	bir.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	borrowInfo, err := borrower.GetBorrowerUID(bir)
	if err != nil || borrowInfo == nil {
		Logger.Errorf("GetBorrowUserDetails query failed", err)
		return &BorrowUserDetailsResponseStruct{
			Status: QUERY_USER_ID_FAILED,
			Msg:    Status[QUERY_USER_ID_FAILED],
		}, nil
	}

	card_id, _ := borrower.GetCardID(borrowInfo.ID)
	credit_use, _ := borrower.GetCreditUse(borrowInfo.ID)
	review_account, _ := borrower.GetReviewAccount(borrowInfo.ID)
	guarantor, _ := borrower.GetGuarantor(borrowInfo.ID)
	material, _ := borrower.GetMaterialInfo(borrowInfo.ID)
	format_credit_use, _ := strconv.ParseFloat(credit_use, 64)
	format_review_account, _ := strconv.ParseFloat(review_account, 64)

	bis := new(BorrowUserDetailsStruct)
	bis.ID = borrowInfo.ID
	bis.Realname = borrowInfo.Realname
	bis.IsBorrower = borrowInfo.IsBorrower
	bis.CardID = card_id
	bis.Credit = fmt.Sprintf("%.6f", (format_credit_use - format_review_account))
	bis.Guarantor = guarantor

	for _, v := range material {
		m := new(MaterialInfoStruct)
		m.ID = v.ID
		m.Name = v.Name
		bis.MaterialList = append(bis.MaterialList, m)
	}

	response := new(BorrowUserDetailsResponseStruct)
	response.BorrowUserDetails = bis
	response.Status = QUERY_BORROW_INFO_SUCCESS
	response.Msg = Status[QUERY_BORROW_INFO_SUCCESS]
	Logger.Debugf("GetBorrowUserDetails response:%v", response)
	return response, nil
}

/**
 * [StartCashRecordServer 开启做标服务---借款人服务]
 * @DateTime 2017-08-24T15:19:45+0800
 */
func StartBorrowerServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30014"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/BorrowUserDetailsThriftService/providers"
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

	handler := &borrowerservice{}
	processor := NewBorrowUserDetailsThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
