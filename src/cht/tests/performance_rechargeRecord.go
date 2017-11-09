package main

import (
	. "cht/service/userrechargerecordlist"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"os"
	"sync"
	"time"
)

func NewRechargeRecordRequestStruct(user_id, start_time, end_time, query_time, status, offset, limitnum int32, log string) *UserRechargeRecordListRequestStruct {
	return &UserRechargeRecordListRequestStruct{
		UserID:               user_id,
		StartTime:            start_time,
		EndTime:              end_time,
		QueryTime:            query_time,
		RechargeStatus:       status,
		LimitOffset:          offset,
		LimitNum:             limitnum,
		ChengHuiTongTraceLog: log,
	}
}

var WaitGroup sync.WaitGroup

func DealJob(i int) {
	host := "192.168.8.209"
	port := "30005"
	trans, err := thrift.NewTSocket(net.JoinHostPort(host, port))
	if err != nil {
		fmt.Sprintf("failed to new socket:", err)
		os.Exit(1)
	}
	defer trans.Close()

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	client := NewUserRechargeRecordListThriftServiceClientFactory(trans, protocolFactory)

	if err := trans.Open(); err != nil {
		fmt.Println("Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	value0 := NewRechargeRecordRequestStruct(242972, 1472486400, 1506441600, 0, 0, 2, 10, "")
	start := time.Now().Unix()
	_, err = client.GetUserRechargeRecordList(value0)
	if err != nil {
		fmt.Println("get resposne failed:", err)
		os.Exit(1)
	}
	end := time.Now().Unix()
	fmt.Printf("ROUTINE %v cost time:%vs\n", i, (end - start))
	// fmt.Printf("response %v", res)
	WaitGroup.Done()
	return
}

func main() {
	var goroutineNum = 10

	for i := 0; i < goroutineNum; i++ {
		WaitGroup.Add(1)
		go func(i int) {
			DealJob(i)
		}(i)
	}

	WaitGroup.Wait()
}
