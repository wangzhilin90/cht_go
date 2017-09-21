// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        "git.apache.org/thrift.git/lib/go/thrift"
        "phoneattestationthriftservice"
)


func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  string checkPhoneUse(CheckPhoneUseRequestStruct requestObj)")
  fmt.Fprintln(os.Stderr, "  i32 getUserIdByhsid(GetUserIdByhsidRequestStruct requestObj)")
  fmt.Fprintln(os.Stderr, "  string updatePhone(UpdatePhoneRequestStruct requestObj)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  var parsedUrl url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Parse()
  
  if len(urlString) > 0 {
    parsedUrl, err := url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  client := phoneattestationthriftservice.NewPhoneAttestationThriftServiceClientFactory(trans, protocolFactory)
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "checkPhoneUse":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CheckPhoneUse requires 1 args")
      flag.Usage()
    }
    arg8 := flag.Arg(1)
    mbTrans9 := thrift.NewTMemoryBufferLen(len(arg8))
    defer mbTrans9.Close()
    _, err10 := mbTrans9.WriteString(arg8)
    if err10 != nil {
      Usage()
      return
    }
    factory11 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt12 := factory11.GetProtocol(mbTrans9)
    argvalue0 := phoneattestationthriftservice.NewCheckPhoneUseRequestStruct()
    err13 := argvalue0.Read(jsProt12)
    if err13 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CheckPhoneUse(value0))
    fmt.Print("\n")
    break
  case "getUserIdByhsid":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetUserIdByhsid requires 1 args")
      flag.Usage()
    }
    arg14 := flag.Arg(1)
    mbTrans15 := thrift.NewTMemoryBufferLen(len(arg14))
    defer mbTrans15.Close()
    _, err16 := mbTrans15.WriteString(arg14)
    if err16 != nil {
      Usage()
      return
    }
    factory17 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt18 := factory17.GetProtocol(mbTrans15)
    argvalue0 := phoneattestationthriftservice.NewGetUserIdByhsidRequestStruct()
    err19 := argvalue0.Read(jsProt18)
    if err19 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetUserIdByhsid(value0))
    fmt.Print("\n")
    break
  case "updatePhone":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "UpdatePhone requires 1 args")
      flag.Usage()
    }
    arg20 := flag.Arg(1)
    mbTrans21 := thrift.NewTMemoryBufferLen(len(arg20))
    defer mbTrans21.Close()
    _, err22 := mbTrans21.WriteString(arg20)
    if err22 != nil {
      Usage()
      return
    }
    factory23 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt24 := factory23.GetProtocol(mbTrans21)
    argvalue0 := phoneattestationthriftservice.NewUpdatePhoneRequestStruct()
    err25 := argvalue0.Read(jsProt24)
    if err25 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.UpdatePhone(value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}