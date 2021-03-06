// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	. "cht/common/logger"
	"cht/service/userloginservice"
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"net/url"
	"os"
	"strings"
)

func main() {
	var host string
	var port int
	var protocol string
	var urlString string
	var trans thrift.TTransport

	flag.StringVar(&host, "h", "192.168.8.208", "Specify host and port")
	flag.IntVar(&port, "p", 30002, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Println("Error parsing URL: ", err)
			os.Exit(1)
		}
		host = parsedUrl.Host
	}

	var err error
	portStr := fmt.Sprint(port)
	if strings.Contains(host, ":") {
		host, portStr, err = net.SplitHostPort(host)
		if err != nil {
			fmt.Println("error with host:", err)
			os.Exit(1)
		}
	}
	trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
	if err != nil {
		fmt.Println("error resolving address:", err)
		os.Exit(1)
	}
	defer trans.Close()

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	client := userloginservice.NewUserLoginThriftServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Println("Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}
	Logger.Debugf("start request")

	ulrs := userloginservice.NewUserlLoginRequestStruct("July", "9f7add09b41ac15889441e467ff208bf", "", "")
	res, err := client.GetUserLoginInfo(ulrs)
	fmt.Println("GetUserLoginInfo res", res)
	Logger.Debug("GetUserLoginInfo res", res)
}
