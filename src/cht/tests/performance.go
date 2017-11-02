package main

import (
	"cht/service/userloginservice"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"os"
	"sync"
	"time"
)

var WaitGroup sync.WaitGroup

func NewUserlLoginRequestStruct(username string, password string, loginip string, log string) *userloginservice.UserLoginRequestStruct {
	ulr := userloginservice.UserLoginRequestStruct{
		Username:             username,
		Password:             password,
		IP:                   loginip,
		ChengHuiTongTraceLog: log,
	}
	return &ulr
}

func DealJob(i int) {
	host := "192.168.8.220"
	port := "30002"
	trans, err := thrift.NewTSocket(net.JoinHostPort(host, port))
	if err != nil {
		fmt.Sprintf("failed to new socket:", err)
		os.Exit(1)
	}
	defer trans.Close()

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	client := userloginservice.NewUserLoginThriftServiceClientFactory(trans, protocolFactory)

	if err := trans.Open(); err != nil {
		fmt.Println("Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	value0 := NewUserlLoginRequestStruct("July", "9f7add09b41ac15889441e467ff208bf", "", "")
	start := time.Now().Unix()
	_, err = client.GetUserLoginInfo(value0)
	if err != nil {
		fmt.Println("get resposne failed:", err)
		os.Exit(1)
	}
	end := time.Now().Unix()
	fmt.Printf("ROUTINE %v cost time:%vs\n", i, (end - start))

	WaitGroup.Done()
	return
}

func main() {
	var goroutineNum = 1000

	for i := 0; i < goroutineNum; i++ {
		WaitGroup.Add(1)
		go func(i int) {
			DealJob(i)
		}(i)
	}

	WaitGroup.Wait()
}
