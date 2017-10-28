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
	QUERY_MES_INFO_SUCCESS = 1000
	QUERY_MES_INFO_FAILED  = 1001

	QUERY_MES_COUNT_SUCCESS = 1000
	QUERY_MES_COUNT_FAILED  = 1002

	QUERY_USER_INFO_SUCCESS = 1000
	QUERY_USER_INFO_FAILED  = 1003
)

var MesInfoStatus = map[int]string{
	QUERY_MES_INFO_SUCCESS: "查询短信详情成功",
	QUERY_MES_INFO_FAILED:  "查询短信详情失败",
}

var MesCountStatus = map[int]string{
	QUERY_MES_COUNT_SUCCESS: "查询短信记录数成功",
	QUERY_MES_COUNT_FAILED:  "查询短信记录数失败",
}

var UserInfoStatus = map[int]string{
	QUERY_USER_INFO_SUCCESS: "查询用户信息成功",
	QUERY_USER_INFO_FAILED:  "查询用户信息失败",
}

/*获取短信详情*/
func (ms *messageservice) GetMessageDetails(requestObj *MessageRequestStruct) (r *MessageDetailsResponseStruct, err error) {
	mr := new(message.MessageRequest)
	mr.Smsid = requestObj.GetSmsid()
	mr.Phone = requestObj.GetPhone()
	mr.Addtime = requestObj.GetAddtime()
	mr.Type = requestObj.GetType()
	mes, err := message.GetMessageInfo(mr)
	if err != nil {
		Logger.Debugf("GetMessageDetails query failed", err)
		return &MessageDetailsResponseStruct{
			Status: QUERY_MES_INFO_FAILED,
			Msg:    MesInfoStatus[QUERY_MES_INFO_FAILED],
		}, nil
	}

	mis := new(MessageDetailsStruct)
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

	response := new(MessageDetailsResponseStruct)
	response.MessageDetails = mis
	response.Status = QUERY_MES_INFO_SUCCESS
	response.Msg = MesInfoStatus[QUERY_MES_INFO_SUCCESS]
	Logger.Debugf("GetMessageDetails res %v", response)
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
			Msg:    MesCountStatus[QUERY_MES_COUNT_FAILED],
		}, nil
	}

	mcr := new(MessageCountResponseStruct)
	mcr.Status = QUERY_MES_INFO_SUCCESS
	mcr.Msg = MesCountStatus[QUERY_MES_INFO_SUCCESS]
	mcr.Count = num
	return mcr, nil
}

/*根据phone获取用户id和手机号*/
func (ms *messageservice) GetUserDetials(requestObj *MessageRequestStruct) (r *UserDetailsResponseStruct, err error) {
	mr := new(message.MessageRequest)
	mr.Smsid = requestObj.GetSmsid()
	mr.Phone = requestObj.GetPhone()
	mr.Addtime = requestObj.GetAddtime()
	mr.Type = requestObj.GetType()
	userInfo, err := message.GetUserInfo(mr)
	if err != nil {
		Logger.Debugf("GetUserInfo query failed", err)
		return &UserDetailsResponseStruct{
			Status: QUERY_USER_INFO_FAILED,
			Msg:    UserInfoStatus[QUERY_USER_INFO_FAILED],
		}, nil
	}

	uis := new(UserDetailsStruct)
	uis.ID = userInfo.ID
	uis.Phone = userInfo.Phone

	uirs := new(UserDetailsResponseStruct)
	uirs.Status = QUERY_USER_INFO_SUCCESS
	uirs.Msg = UserInfoStatus[QUERY_USER_INFO_SUCCESS]
	uirs.UserDetails = uis
	return uirs, nil
}

/**
 * [StartUpdatePasswdsServer 开启短信服务]
 * @DateTime 2017-09-11T15:19:45+0800
 */
func StartMessageServer() {
	zkServers := zkclient.ZkServerAddress
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
		Logger.Fatalf("RegisterNode %v failed", servicename, err)
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
