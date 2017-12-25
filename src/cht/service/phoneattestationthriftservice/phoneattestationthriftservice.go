package phoneattestationthriftservice

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/phoneattestation"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

type phoneattestationservice struct{}

const (
	PHONE_USED   = "1001" //没查到记录返回1000
	PHONE_UNUSED = "1000" //查到记录返回1001
)

func (pts *phoneattestationservice) CheckPhoneUse(requestObj *CheckPhoneUseRequestStruct) (string, error) {
	cpur := new(phoneattestation.CheckPhoneUseRequest)
	cpur.Phone = requestObj.GetPhone()
	cpur.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	b := phoneattestation.CheckPhoneUse(cpur)
	if b {
		Logger.Debugf("CheckPhoneUse phone %v is used status :%v", cpur.Phone, PHONE_USED)
		return PHONE_USED, nil
	} else {
		Logger.Debugf("CheckPhoneUse phone %v is not used status :%v", cpur.Phone, PHONE_UNUSED)
		return PHONE_UNUSED, nil
	}
}

func (pts *phoneattestationservice) GetUserIdByhsid(requestObj *GetUserIdByhsidRequestStruct) (int32, error) {
	gibr := new(phoneattestation.GetUserIdByhsidRequest)
	gibr.Hsid = requestObj.GetHsid()
	gibr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	user_id, err := phoneattestation.GetUserIdByhsid(gibr)
	if err != nil {
		Logger.Errorf("GetUserIdByhsid query failed", err)
	}
	Logger.Debugf("GetUserIdByhsid res %v", user_id)
	return user_id, nil
}

func (pts *phoneattestationservice) UpdatePhone(requestObj *UpdatePhoneRequestStruct) (string, error) {
	upr := new(phoneattestation.UpdatePhoneRequest)
	upr.Phone = requestObj.GetPhone()
	upr.UserID = requestObj.GetUserID()
	status := phoneattestation.UpdatePhone(upr)

	Logger.Debugf("UpdatePhone status:%v", status)
	return status, nil
}

/**
 * [StartPhoneAttestationServer ]
 * @DateTime 2017-09-14T11:41:45+0800
 */
func StartPhoneAttestationServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30015"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/PhoneAttestationThriftService/providers"
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

	handler := &phoneattestationservice{}
	processor := NewPhoneAttestationThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
