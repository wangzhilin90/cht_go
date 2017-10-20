package userattestioncardinfosave

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	uacis "cht/models/userattestioncardinfosave"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	SAVE_USER_ATTESTION_CARDINFO_SUCCESS = 1000
	SAVE_USER_ATTESTION_CARDINFO_FAILED  = 1001
)

var Stat = map[int]string{
	SAVE_USER_ATTESTION_CARDINFO_SUCCESS: "保存用户认证卡证信息成功",
	SAVE_USER_ATTESTION_CARDINFO_FAILED:  "保存用户认证卡证卡证失败",
}

type userattestioncardinfosaveservice struct{}

func (uis *userattestioncardinfosaveservice) SaveUserAttestionCardInfo(requestObj *UserAttestionCardInfoSaveRequestStruct) (r *UserAttestionCardInfoSaveResponseStruct, err error) {
	uacisr := new(uacis.UserAttestionCardInfoSaveRequest)
	uacisr.UserID = requestObj.GetUserID()
	uacisr.CardType = requestObj.GetCardType()
	uacisr.CardID = requestObj.GetCardID()
	uacisr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	b := uacis.SaveUserAttestionCardInfo(uacisr)
	if b == false {
		return &UserAttestionCardInfoSaveResponseStruct{
			Status: SAVE_USER_ATTESTION_CARDINFO_FAILED,
			Msg:    Stat[SAVE_USER_ATTESTION_CARDINFO_FAILED],
		}, nil
	}

	return &UserAttestionCardInfoSaveResponseStruct{
		Status: SAVE_USER_ATTESTION_CARDINFO_SUCCESS,
		Msg:    Stat[SAVE_USER_ATTESTION_CARDINFO_SUCCESS],
	}, nil
}

/**
 * [StartUserAttestionCardInfoSaveServer 保存用户认证卡证信息服务]
 * @DateTime 2017-10-20T09:47:01+0800
 */
func StartUserAttestionCardInfoSaveServer() {
	zkServers := []string{"192.168.8.208:2181"}
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30039"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/UserAttestionCardInfoSaveThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &userattestioncardinfosaveservice{}
	processor := NewUserAttestionCardInfoSaveThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
