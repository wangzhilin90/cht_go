package emailattestationthriftservice

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/emailattestation"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

type emailattestationservice struct{}

func (eas *emailattestationservice) CheckEmailUse(requestObj *CheckEmailUseRequestStruct) (r int32, err error) {
	eurs := new(emailattestation.CheckEmailUseRequestStruct)
	eurs.Email = requestObj.GetEmail()
	eurs.UserID = requestObj.GetUserID()
	eurs.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	res := emailattestation.CheckEmailUse(eurs)
	return res, nil
}

func (eas *emailattestationservice) UserAttestationSave(requestObj *UserAttestationSaveStruct) (r int32, err error) {
	uass := new(emailattestation.UserAttestationSaveStruct)
	uass.UserID = requestObj.GetUserID()
	uass.EmailStatus = requestObj.GetEmailStatus()
	uass.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	res := emailattestation.UserAttestationSave(uass)
	return res, nil
}

func (eas *emailattestationservice) UserEmailSave(requestObj *UserEmailSaveRequestStruct) (r int32, err error) {
	uesrs := new(emailattestation.UserEmailSaveRequestStruct)
	uesrs.Email = requestObj.GetEmail()
	uesrs.UserID = requestObj.GetUserID()
	uesrs.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	res := emailattestation.UserEmailSave(uesrs)
	return res, nil
}

func (eas *emailattestationservice) SendEmail(requestObj *SendEmailRequestStruct) (r int32, err error) {
	sers := new(emailattestation.SendEmailRequestStruct)
	sers.UserID = requestObj.GetUserID()
	sers.SendTo = requestObj.GetSendTo()
	sers.Subject = requestObj.GetSubject()
	sers.Content = requestObj.GetContent()
	sers.IP = requestObj.GetIP()
	sers.Addtime = requestObj.GetAddtime()
	sers.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	res := emailattestation.SendEmail(sers)
	return res, nil
}

/*开启邮箱认证服务*/
func StartEmailAttestationServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30016"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/EmailAttestationThriftService/providers"
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

	handler := &emailattestationservice{}
	processor := NewEmailAttestationThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
