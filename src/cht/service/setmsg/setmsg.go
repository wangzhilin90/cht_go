package setmsg

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	sm "cht/models/setmsg"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	QUERY_SET_MSG_DETAILS_SUCCESS = 1000
	QUERY_SET_MSG_DETAILS_FAILED  = 1001
	QUERY_SET_MSG_DETAILS_EMPTY   = 1002
)

var Details_Stat = map[int]string{
	QUERY_SET_MSG_DETAILS_SUCCESS: "查询用户设置提醒表详情成功",
	QUERY_SET_MSG_DETAILS_FAILED:  "查询用户设置提醒表详情失败",
	QUERY_SET_MSG_DETAILS_EMPTY:   "查询用户设置提醒表详情为空",
}

const (
	UPDATE_SET_MSG_SUCCESS = 1000
	UPDATE_SET_MSG_FAILED  = 1001
)

var Update_Stat = map[int]string{
	UPDATE_SET_MSG_SUCCESS: "更新用户设置提醒表成功",
	UPDATE_SET_MSG_FAILED:  "更新用户设置提醒表失败",
}

const (
	INSERT_SET_MSG_SUCCESS = 1000
	INSERT_SET_MSG_FAILED  = 1001
)

var Insert_Stat = map[int]string{
	INSERT_SET_MSG_SUCCESS: "新增用户设置提醒表成功",
	INSERT_SET_MSG_FAILED:  "新增用户设置提醒表失败",
}

type setmsgservice struct{}

func (sms *setmsgservice) GetSetMsgDetails(requestObj *SetMsgDetailsRequestStruct) (r *SetMsgDetailsResponseStruct, err error) {
	Logger.Infof("GetSetMsgDetails requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*SetMsgDetailsRequestStruct)
	smdr := new(sm.SetMsgDetailsRequest)
	smdr.UserID = requestObj.GetUserID()
	smdr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, err := sm.GetSetMsgDetails(smdr)
	if err != nil {
		Logger.Errorf("GetSetMsgDetails failed:%v", err)
		return &SetMsgDetailsResponseStruct{
			Status: QUERY_SET_MSG_DETAILS_FAILED,
			Msg:    Details_Stat[QUERY_SET_MSG_DETAILS_FAILED],
		}, nil
	}
	if res == nil {
		Logger.Debugf("GetSetMsgDetails query empty")
		return &SetMsgDetailsResponseStruct{
			Status: QUERY_SET_MSG_DETAILS_EMPTY,
			Msg:    Details_Stat[QUERY_SET_MSG_DETAILS_EMPTY],
		}, nil
	}

	var response SetMsgDetailsResponseStruct
	if res != nil {
		smdrs := new(SetMsgDetailsStruct)
		smdrs.ID = res.ID
		smdrs.UserID = res.UserID
		smdrs.Addtime = res.Addtime
		smdrs.Status = res.Status
		smdrs.TenderStatus = res.TenderStatus
		smdrs.BorrowStatus = res.BorrowStatus
		smdrs.ProtoStatus = res.ProtoStatus
		smdrs.StationStatus = res.StationStatus
		smdrs.GuideStatus = res.GuideStatus
		smdrs.SoundStatus = res.SoundStatus
		response.SetMsgDetails = smdrs
	}

	response.Status = QUERY_SET_MSG_DETAILS_SUCCESS
	response.Msg = Details_Stat[QUERY_SET_MSG_DETAILS_SUCCESS]
	Logger.Debugf("GetSetMsgDetails response:%v", response)
	return &response, nil
}

func (sms *setmsgservice) UpdateSetMsgDetails(requestObj *SetMsgDealRequestStruct) (r *SetMsgDealResponseStruct, err error) {
	Logger.Infof("UpdateSetMsgDetails requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*SetMsgDealRequestStruct)
	smdr := new(sm.SetMsgDealRequest)
	smdr.UserID = requestObj.GetUserID()
	smdr.Addtime = requestObj.GetAddtime()
	smdr.Status = requestObj.GetStatus()
	smdr.TenderStatus = requestObj.GetTenderStatus()
	smdr.BorrowStatus = requestObj.GetBorrowStatus()
	smdr.ProtoStatus = requestObj.GetProtoStatus()
	smdr.StationStatus = requestObj.GetStationStatus()
	smdr.GuideStatus = requestObj.GetGuideStatus()
	smdr.SoundStatus = requestObj.GetSoundStatus()
	smdr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	b := sm.UpdateSetMsgDetails(smdr)
	if b == false {
		Logger.Errorf("UpdateSetMsgDetails failed:%v", err)
		return &SetMsgDealResponseStruct{
			Status: UPDATE_SET_MSG_FAILED,
			Msg:    Update_Stat[UPDATE_SET_MSG_FAILED],
		}, nil
	}

	return &SetMsgDealResponseStruct{
		Status: UPDATE_SET_MSG_SUCCESS,
		Msg:    Update_Stat[UPDATE_SET_MSG_SUCCESS],
	}, nil
}

func (sms *setmsgservice) InsertSetMsgDetails(requestObj *SetMsgDealRequestStruct) (r *SetMsgDealResponseStruct, err error) {
	Logger.Infof("InsertSetMsgDetails requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*SetMsgDealRequestStruct)
	smdr := new(sm.SetMsgDealRequest)
	smdr.UserID = requestObj.GetUserID()
	smdr.Addtime = requestObj.GetAddtime()
	smdr.Status = requestObj.GetStatus()
	smdr.TenderStatus = requestObj.GetTenderStatus()
	smdr.BorrowStatus = requestObj.GetBorrowStatus()
	smdr.ProtoStatus = requestObj.GetProtoStatus()
	smdr.StationStatus = requestObj.GetStationStatus()
	smdr.GuideStatus = requestObj.GetGuideStatus()
	smdr.SoundStatus = requestObj.GetSoundStatus()
	smdr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	b := sm.InsertSetMsgDetails(smdr)
	if b == false {
		Logger.Errorf("InsertSetMsgDetails failed:%v", err)
		return &SetMsgDealResponseStruct{
			Status: UPDATE_SET_MSG_FAILED,
			Msg:    Update_Stat[UPDATE_SET_MSG_FAILED],
		}, nil
	}

	return &SetMsgDealResponseStruct{
		Status: UPDATE_SET_MSG_SUCCESS,
		Msg:    Update_Stat[UPDATE_SET_MSG_SUCCESS],
	}, nil
}

/**
 * [StartSetMsgServer 开启用户设置提醒表服务]
 * @DateTime 2017-12-13T10:45:57+0800
 */
func StartSetMsgServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30068"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/SetMsgThriftService/providers"
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

	handler := &setmsgservice{}
	processor := NewSetMsgThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
