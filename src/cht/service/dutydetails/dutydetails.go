package dutydetails

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/dutydetails"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

type dutydetailsservice struct{}

const (
	QUERY_DUTY_DETAILS_FAILED  = 1001
	QUERY_DUTY_DETAILS_SUCCESS = 1000
)

var Status = map[int]string{
	QUERY_DUTY_DETAILS_FAILED:  "查询值班人失败",
	QUERY_DUTY_DETAILS_SUCCESS: "查询值班人成功",
}

func (ddss *dutydetailsservice) GetDutyDetails(requestObj *DutyDetailsRequestStruct) (r *DutyDetailsResponseStruct, err error) {
	ddrs := new(dutydetails.DutyDetailsRequestStruct)
	ddrs.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	res, err := dutydetails.GetDutyDetails(ddrs)
	if err != nil {
		Logger.Errorf("GetDutyDetails query failed %v", err)
		return &DutyDetailsResponseStruct{
			Status: QUERY_DUTY_DETAILS_FAILED,
			Msg:    Status[QUERY_DUTY_DETAILS_FAILED],
		}, nil
	}

	var response DutyDetailsResponseStruct
	if res != nil {
		dds := new(DutyDetailsStruct)
		dds.ID = res.ID
		dds.Customer = res.Customer
		dds.IsRest = res.IsRest
		dds.DutyTime = res.DutyTime
		dds.HolidayUser = res.HolidayUser
		dds.StartTime = res.StartTime
		dds.EndTime = res.EndTime
		dds.Addtime = res.Addtime
		response.DutyDetails = dds
	}

	response.Status = QUERY_DUTY_DETAILS_SUCCESS
	response.Msg = Status[QUERY_DUTY_DETAILS_SUCCESS]
	Logger.Debugf("GetDutyDetails response:%v", response)
	return &response, nil
}

/**
 * [StartDutyDetailServer 开启后台服务---值班人详情服务]
 * @DateTime 2017-08-24T15:19:45+0800
 */
func StartDutyDetailServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30022"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/DutyDetailsThriftService/providers"
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

	handler := &dutydetailsservice{}
	processor := NewDutyDetailsThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
