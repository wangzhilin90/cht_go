package userattestionlist

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	ual "cht/models/userattestionlist"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	QUERY_USER_ATTESTION_SUCCESS          = 1000
	QUERY_USER_ATTESTION_TOTAL_NUM_FAILED = 1001
	QUERY_USER_ATTESTION_LIST_FAILED      = 1002
)

var Stat = map[int]string{
	QUERY_USER_ATTESTION_SUCCESS:          "基础认证信息查找成功",
	QUERY_USER_ATTESTION_TOTAL_NUM_FAILED: "基础认证信息记录数获取失败",
	QUERY_USER_ATTESTION_LIST_FAILED:      "基础认证信息列表获取失败",
}

type userattestionlistservice struct{}

func (uals *userattestionlistservice) UserAttestionList(requestObj *UserAttestionListRequestStruct) (r *UserAttestionListResponseStruct, err error) {
	Logger.Infof("UserAttestionList requestObj:%v", requestObj)
	ualr := new(ual.UserAttestionListRequest)
	ualr.Username = requestObj.GetUsername()
	ualr.Realname = requestObj.GetRealname()
	ualr.RealStatus = requestObj.GetRealStatus()
	ualr.EmailStatus = requestObj.GetEmailStatus()
	ualr.PhoneStatus = requestObj.GetPhoneStatus()
	ualr.VideoStatus = requestObj.GetVideoStatus()
	ualr.SceneStatus = requestObj.GetSceneStatus()
	ualr.LimitNum = requestObj.GetLimitNum()
	ualr.LimitOffset = requestObj.GetLimitOffset()
	ualr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	tatalNum, err := ual.GetUserAttestionTatalNum(ualr)
	if err != nil {
		Logger.Errorf("UserAttestionList GetUserAttestionTatalNum failed:%v", err)
		return &UserAttestionListResponseStruct{
			Status: QUERY_USER_ATTESTION_TOTAL_NUM_FAILED,
			Msg:    Stat[QUERY_USER_ATTESTION_TOTAL_NUM_FAILED],
		}, nil
	}

	res, err := ual.GetUserAttestionList(ualr)
	if err != nil {
		Logger.Errorf("UserAttestionList UserAttestationDetails failed:%v", err)
		return &UserAttestionListResponseStruct{
			Status: QUERY_USER_ATTESTION_LIST_FAILED,
			Msg:    Stat[QUERY_USER_ATTESTION_LIST_FAILED],
		}, nil
	}

	var response UserAttestionListResponseStruct
	for _, v := range res {
		uads := new(UserAttestationDetailsStruct)
		uads.UserID = v.UserID
		uads.CardType = v.CardType
		uads.HsCardType = v.HsCardType
		uads.CardID = v.CardID
		uads.CardPic1 = v.CardPic1
		uads.CardPic2 = v.CardPic2
		uads.VideoPic = v.VideoPic
		uads.RealStatus = v.RealStatus
		uads.RealPasstime = v.RealPasstime
		uads.EmailStatus = v.EmailStatus
		uads.EmailPasstime = v.EmailPasstime
		uads.PhoneStatus = v.PhoneStatus
		uads.PhonePasstime = v.PhonePasstime
		uads.SceneStatus = v.SceneStatus
		uads.ScenePasstime = v.ScenePasstime
		uads.VipStatus = v.VipStatus
		uads.VipPasstime = v.VipPasstime
		uads.VipVerifytime = v.VipVerifytime
		uads.Username = v.Username
		uads.Realname = v.Realname
		uads.Phone = v.Phone
		uads.Email = v.Email
		uads.Name = v.Name
		response.UserAttestionList = append(response.UserAttestionList, uads)
	}
	response.Total = tatalNum
	response.Status = QUERY_USER_ATTESTION_SUCCESS
	response.Msg = Stat[QUERY_USER_ATTESTION_SUCCESS]
	Logger.Debugf("UserAttestionList response:%v", response)
	return &response, nil
}

/**
 * [StartUpdatePasswdsServer 基础认证列表服务]
 * @DateTime 2017-10-19T16:35:39+0800
 */
func StartUserAttestionListServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30037"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/UserAttestionListThriftService/providers"
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

	handler := &userattestionlistservice{}
	processor := NewUserAttestionListThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
