package borrowerthriftservice

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/borrower"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
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

func (bs *borrowerservice) GetBorrowerInfo(requestObj *BorrowerInfoRequestStruct) (r *BorrowerInfoResponseStruct, err error) {
	bir := new(borrower.BorrowerInfoRequest)
	bir.Name = requestObj.GetName()
	bir.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	borrowInfo, err := borrower.GetBorrowerInfo(bir)
	if err != nil {
		Logger.Errorf("GetBorrowerInfo failed %v", err)
		return &BorrowerInfoResponseStruct{
			Status: QUERY_USER_ID_FAILED,
			Msg:    Status[QUERY_USER_ID_FAILED],
		}, nil
	}

	bi := new(BorrowerInfoStruct)
	bi.ID = borrowInfo.ID
	bi.Realname = borrowInfo.Realname
	bi.IsBorrower = borrowInfo.IsBorrower
	bi.Credit = borrowInfo.Credit
	bi.Guarantor = borrowInfo.Guarantor

	mi := new(MaterialInfoStruct)
	mi.ID = borrowInfo.MaterialInfo.ID
	mi.Name = borrowInfo.MaterialInfo.Name

	bi.MaterialInfo = mi
	response := new(BorrowerInfoResponseStruct)
	response.BorrowerInfo = bi
	response.Status = QUERY_BORROW_INFO_SUCCESS
	response.Msg = Status[QUERY_BORROW_INFO_SUCCESS]
	Logger.Debugf("GetBorrowerInfo res:%v", response)
	return response, nil
}

/**
 * [StartCashRecordServer 开启做标服务---借款人服务]
 * @DateTime 2017-08-24T15:19:45+0800
 */
func StartCashRecordServer() {
	zkServers := []string{"192.168.8.208:2181"}
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30014"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/BorrowerThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &borrowerservice{}
	processor := NewBorrowerThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
