package phoneattestationthriftservice

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/phoneattestation"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type phoneattestationservice struct{}

const (
	PHONE_USED   = "1001"
	PHONE_UNUSED = "1000"
)

func (pts *phoneattestationservice) CheckPhoneByPhone(requestObj *CheckPhoneUseRequestStruct) (string, error) {
	cpur := new(phoneattestation.CheckPhoneUseRequest)
	cpur.Phone = requestObj.GetPhone()
	cpur.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	b := phoneattestation.CheckPhoneUse(cpur)
	if b {
		Logger.Debugf("CheckPhoneByPhone phone %v is used status :%v", cpur.Phone, PHONE_USED)
		return PHONE_USED, nil
	} else {
		Logger.Debugf("CheckPhoneByPhone phone %v is not used status :%v", cpur.Phone, PHONE_UNUSED)
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

func (pts *phoneattestationservice) UpdatePhoneByTransaction(requestObj *UpdatePhoneRequestStruct) (string, error) {
	upr := new(phoneattestation.UpdatePhoneRequest)
	upr.Phone = requestObj.GetPhone()
	upr.UserID = requestObj.GetUserID()
	status := phoneattestation.UpdatePhone(upr)

	Logger.Debugf("UpdatePhoneByTransaction status:%v", status)
	return status, nil
}

/**
 * [StartPhoneAttestationServer ]
 * @DateTime 2017-09-14T11:41:45+0800
 */
func StartPhoneAttestationServer() {
	zkServers := []string{"192.168.8.208:2181"}
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
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &phoneattestationservice{}
	processor := NewPhoneAttestationThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
