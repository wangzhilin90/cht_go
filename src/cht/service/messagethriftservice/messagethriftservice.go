package messagethriftservice

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/message"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type MessageRequest struct {
	Smsid                int32  //短信id
	Phone                string //手机号
	Addtime              string //添加时间
	Type                 int32  //类型
	ChengHuiTongTraceLog string
}

type messageservice struct{}

const (
	QUERY_MES_INFO_SUCCESS  = 1000
	QUERY_MES_INFO_FAILED   = 1001
	QUERY_MES_COUNT_SUCCESS = 1002
	QUERY_MES_COUNT_FAILED  = 1003
)

var Status = map[int]string{
	QUERY_MES_INFO_SUCCESS:  "查询短信详情成功",
	QUERY_MES_INFO_FAILED:   "查询短信详情失败",
	QUERY_MES_COUNT_SUCCESS: "查询短信记录数成功",
	QUERY_MES_COUNT_FAILED:  "查询短信记录数失败",
}

/*获取短信详情*/
func (ms *messageservice) GetMessageInfo(requestObj *MessageRequestStruct) (r *MessageInfoResponseStruct, err error) {
	mr := new(message.MessageRequest)
	mr.Smsid = requestObj.GetSmsid()
	mr.Phone = requestObj.GetPhone()
	mr.Addtime = requestObj.GetAddtime()
	mr.Type = requestObj.GetType()
	mes, err := message.GetMessageInfo(mr)
	if err != nil {
		Logger.Debugf("GetMessageInfo query failed", err)
		return &MessageInfoResponseStruct{
			Status: QUERY_MES_INFO_FAILED,
			Msg:    Status[QUERY_MES_INFO_FAILED],
		}, nil
	}

	mis := new(MessageInfoStruct)
	mis.ID = mes.ID
	mis.Type = mes.Type
	mis.UserID = mes.UserID
	mis.SendTo = mes.SendTo
	mis.Subject = mes.Subject
	mis.Content = mes.Content
	mis.Attachment = mes.Attachment
	mis.Addtime = mes.Addtime
	mis.IP = mes.IP
	mis.Posttime = mes.Posttime
	mis.Status = mes.Status

	response := new(MessageInfoResponseStruct)
	response.MessageInfo = mis
	response.Status = QUERY_MES_INFO_SUCCESS
	response.Msg = Status[QUERY_MES_INFO_SUCCESS]
	Logger.Debugf("GetMessageInfo res %v", response)
	return response, nil
}

/*获取短信记录数*/
func (ms *messageservice) GetMessageCount(requestObj *MessageRequestStruct) (*MessageCountResponseStruct, error) {
	mr := new(message.MessageRequest)
	mr.Smsid = requestObj.GetSmsid()
	mr.Phone = requestObj.GetPhone()
	mr.Addtime = requestObj.GetAddtime()
	mr.Type = requestObj.GetType()
	num, err := message.GetMessageCount(mr)
	if err != nil {
		Logger.Debugf("GetMessageCount query failed", err)
		return &MessageCountResponseStruct{
			Status: QUERY_MES_COUNT_FAILED,
			Msg:    Status[QUERY_MES_COUNT_FAILED],
		}, nil
	}

	mcr := new(MessageCountResponseStruct)
	mcr.Status = QUERY_MES_INFO_SUCCESS
	mcr.Msg = Status[QUERY_MES_INFO_SUCCESS]
	mcr.Count = num
	return mcr, nil
}

/**
 * [StartUpdatePasswdsServer 开启短信服务]
 * @DateTime 2017-09-11T15:19:45+0800
 */
func StartMessageServer() {
	zkServers := []string{"192.168.8.212:2181", "192.168.8.213:2181", "192.168.8.214:2181"}
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30011"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/MessageThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &messageservice{}
	processor := NewMessageThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}