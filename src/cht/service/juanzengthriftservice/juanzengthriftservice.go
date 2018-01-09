package juanzengthriftservice

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/juanzeng"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"reflect"
	"regexp"
	"time"
)

type juanzengservice struct{}

func fiterSpecialCharacters(input *RequestStruct) *RequestStruct {
	str := `select|Update|and|or|delete|insert|trancate| \
			char|into|substr|ascii|declare|exec|count|master|into| \
			drop|execute|\"|'|%|;|\(|\)|&|\+`
	var re, _ = regexp.Compile(str)
	t := reflect.ValueOf(input).Elem()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		switch f.Kind() {
		case reflect.String:
			if f.CanInterface() {
				if str, ok := f.Interface().(string); ok {
					new1 := re.ReplaceAllString(str, "")
					f.SetString(new1)
				}
			}
		}
	}

	return input
}

func (js *juanzengservice) GetInfo(requestObj *RequestStruct) (r *JuanzengResponseStruct, err error) {
	Logger.Infof("GetInfo requestObj:%v", requestObj)
	requestObj = fiterSpecialCharacters(requestObj)
	rs := new(juanzeng.RequestStruct)
	rs.UserID = requestObj.GetUserID()
	rs.Content = requestObj.GetContent()
	rs.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	var response JuanzengResponseStruct
	messlistResult, err := juanzeng.GetMesslistResult(rs)
	if err != nil {
		response.Messlist = nil
	} else {
		for _, v := range messlistResult {
			mrs := new(MesslistResultStruct)
			mrs.Username = v.Username
			mrs.Avatar = v.Avatar
			mrs.Addtime = v.Addtime
			mrs.Content = v.Content
			mrs.Reply = v.Reply
			mrs.UpContent = v.UpContent
			mrs.UpReply = v.UpReply
			response.Messlist = append(response.Messlist, mrs)
		}
	}

	fundlistResult, err := juanzeng.GetFundlistResult(rs)
	if err != nil {
		response.Fundlist = nil
	} else {
		for _, v := range fundlistResult {
			frs := new(FundlistResultStruct)
			frs.Type = v.Type
			frs.Addtime = v.Addtime
			frs.Username = v.Username
			frs.Money = v.Money
			response.Fundlist = append(response.Fundlist, frs)
		}
	}

	numlistResult, _ := juanzeng.GetNumlistResult(rs)
	if numlistResult != nil {
		nrs := new(NumlistResultStruct)
		nrs.Num = numlistResult.Num
		nrs.Money = numlistResult.Money
		response.Numlist = nrs
	}

	totalJunaNum, _ := juanzeng.GetTotalJuanNum(rs)
	response.Tzr = totalJunaNum
	return &response, nil
}

/**
 * [func 添加留言]
 * @param    AddMess(requestObj *RequestStruct 请求入参
 * @return   int32 返回最新插入的一条留言ID，插入失败返回0
 * @return   {[type]}                    [description]
 * @DateTime 2017-09-20T16:15:48+0800
 */
func (js *juanzengservice) AddMess(requestObj *RequestStruct) (r int32, err error) {
	requestObj = fiterSpecialCharacters(requestObj)
	rs := new(juanzeng.RequestStruct)
	rs.UserID = requestObj.GetUserID()
	rs.Content = requestObj.GetContent()
	rs.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	lastInsertNUm, err := juanzeng.AddMess(rs)
	if err != nil {
		Logger.Errorf("AddMess return failed :%v", err)
		return 0, nil
	}
	return lastInsertNUm, nil
}

/*开启捐赠服务*/
func StartJuanzengServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30019"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/JuanzengThriftService/providers"
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

	handler := &juanzengservice{}
	processor := NewJuanzengThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
