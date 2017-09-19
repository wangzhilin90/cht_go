// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package sysconfigthriftservice

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

// Attributes:
//  - ID
//  - Nid
//  - Value
//  - Name
type SysConfigStruct struct {
  ID int32 `thrift:"id,1" db:"id" json:"id"`
  Nid string `thrift:"nid,2" db:"nid" json:"nid"`
  Value string `thrift:"value,3" db:"value" json:"value"`
  Name string `thrift:"name,4" db:"name" json:"name"`
}

func NewSysConfigStruct() *SysConfigStruct {
  return &SysConfigStruct{}
}


func (p *SysConfigStruct) GetID() int32 {
  return p.ID
}

func (p *SysConfigStruct) GetNid() string {
  return p.Nid
}

func (p *SysConfigStruct) GetValue() string {
  return p.Value
}

func (p *SysConfigStruct) GetName() string {
  return p.Name
}
func (p *SysConfigStruct) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    case 3:
      if err := p.ReadField3(iprot); err != nil {
        return err
      }
    case 4:
      if err := p.ReadField4(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *SysConfigStruct)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.ID = v
}
  return nil
}

func (p *SysConfigStruct)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Nid = v
}
  return nil
}

func (p *SysConfigStruct)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.Value = v
}
  return nil
}

func (p *SysConfigStruct)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.Name = v
}
  return nil
}

func (p *SysConfigStruct) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("SysConfigStruct"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
    if err := p.writeField4(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *SysConfigStruct) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("id", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:id: ", p), err) }
  if err := oprot.WriteI32(int32(p.ID)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.id (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:id: ", p), err) }
  return err
}

func (p *SysConfigStruct) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("nid", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:nid: ", p), err) }
  if err := oprot.WriteString(string(p.Nid)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.nid (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:nid: ", p), err) }
  return err
}

func (p *SysConfigStruct) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("value", thrift.STRING, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:value: ", p), err) }
  if err := oprot.WriteString(string(p.Value)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.value (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:value: ", p), err) }
  return err
}

func (p *SysConfigStruct) writeField4(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("name", thrift.STRING, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:name: ", p), err) }
  if err := oprot.WriteString(string(p.Name)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.name (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:name: ", p), err) }
  return err
}

func (p *SysConfigStruct) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("SysConfigStruct(%+v)", *p)
}

// Attributes:
//  - ChengHuiTongTraceLog
type SysConfigRequestStruct struct {
  ChengHuiTongTraceLog string `thrift:"chengHuiTongTraceLog,1" db:"chengHuiTongTraceLog" json:"chengHuiTongTraceLog"`
}

func NewSysConfigRequestStruct() *SysConfigRequestStruct {
  return &SysConfigRequestStruct{}
}


func (p *SysConfigRequestStruct) GetChengHuiTongTraceLog() string {
  return p.ChengHuiTongTraceLog
}
func (p *SysConfigRequestStruct) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *SysConfigRequestStruct)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.ChengHuiTongTraceLog = v
}
  return nil
}

func (p *SysConfigRequestStruct) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("SysConfigRequestStruct"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *SysConfigRequestStruct) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("chengHuiTongTraceLog", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:chengHuiTongTraceLog: ", p), err) }
  if err := oprot.WriteString(string(p.ChengHuiTongTraceLog)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.chengHuiTongTraceLog (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:chengHuiTongTraceLog: ", p), err) }
  return err
}

func (p *SysConfigRequestStruct) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("SysConfigRequestStruct(%+v)", *p)
}

// Attributes:
//  - Status
//  - Msg
//  - SysConfigList
type SysConfigResponseStruct struct {
  Status int32 `thrift:"status,1" db:"status" json:"status"`
  Msg string `thrift:"msg,2" db:"msg" json:"msg"`
  SysConfigList []*SysConfigStruct `thrift:"sysConfigList,3" db:"sysConfigList" json:"sysConfigList"`
}

func NewSysConfigResponseStruct() *SysConfigResponseStruct {
  return &SysConfigResponseStruct{}
}


func (p *SysConfigResponseStruct) GetStatus() int32 {
  return p.Status
}

func (p *SysConfigResponseStruct) GetMsg() string {
  return p.Msg
}

func (p *SysConfigResponseStruct) GetSysConfigList() []*SysConfigStruct {
  return p.SysConfigList
}
func (p *SysConfigResponseStruct) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    case 3:
      if err := p.ReadField3(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *SysConfigResponseStruct)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Status = v
}
  return nil
}

func (p *SysConfigResponseStruct)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Msg = v
}
  return nil
}

func (p *SysConfigResponseStruct)  ReadField3(iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]*SysConfigStruct, 0, size)
  p.SysConfigList =  tSlice
  for i := 0; i < size; i ++ {
    _elem0 := &SysConfigStruct{}
    if err := _elem0.Read(iprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem0), err)
    }
    p.SysConfigList = append(p.SysConfigList, _elem0)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *SysConfigResponseStruct) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("SysConfigResponseStruct"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *SysConfigResponseStruct) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("status", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:status: ", p), err) }
  if err := oprot.WriteI32(int32(p.Status)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.status (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:status: ", p), err) }
  return err
}

func (p *SysConfigResponseStruct) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("msg", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:msg: ", p), err) }
  if err := oprot.WriteString(string(p.Msg)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.msg (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:msg: ", p), err) }
  return err
}

func (p *SysConfigResponseStruct) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("sysConfigList", thrift.LIST, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:sysConfigList: ", p), err) }
  if err := oprot.WriteListBegin(thrift.STRUCT, len(p.SysConfigList)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
  }
  for _, v := range p.SysConfigList {
    if err := v.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
    }
  }
  if err := oprot.WriteListEnd(); err != nil {
    return thrift.PrependError("error writing list end: ", err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:sysConfigList: ", p), err) }
  return err
}

func (p *SysConfigResponseStruct) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("SysConfigResponseStruct(%+v)", *p)
}

type SysConfigThriftService interface {
  // Parameters:
  //  - RequestObj
  GetSysConfig(requestObj *SysConfigRequestStruct) (r *SysConfigResponseStruct, err error)
}

type SysConfigThriftServiceClient struct {
  Transport thrift.TTransport
  ProtocolFactory thrift.TProtocolFactory
  InputProtocol thrift.TProtocol
  OutputProtocol thrift.TProtocol
  SeqId int32
}

func NewSysConfigThriftServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *SysConfigThriftServiceClient {
  return &SysConfigThriftServiceClient{Transport: t,
    ProtocolFactory: f,
    InputProtocol: f.GetProtocol(t),
    OutputProtocol: f.GetProtocol(t),
    SeqId: 0,
  }
}

func NewSysConfigThriftServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *SysConfigThriftServiceClient {
  return &SysConfigThriftServiceClient{Transport: t,
    ProtocolFactory: nil,
    InputProtocol: iprot,
    OutputProtocol: oprot,
    SeqId: 0,
  }
}

// Parameters:
//  - RequestObj
func (p *SysConfigThriftServiceClient) GetSysConfig(requestObj *SysConfigRequestStruct) (r *SysConfigResponseStruct, err error) {
  if err = p.sendGetSysConfig(requestObj); err != nil { return }
  return p.recvGetSysConfig()
}

func (p *SysConfigThriftServiceClient) sendGetSysConfig(requestObj *SysConfigRequestStruct)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("getSysConfig", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := SysConfigThriftServiceGetSysConfigArgs{
  RequestObj : requestObj,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *SysConfigThriftServiceClient) recvGetSysConfig() (value *SysConfigResponseStruct, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "getSysConfig" {
    err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getSysConfig failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getSysConfig failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error1 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error2 error
    error2, err = error1.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error2
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getSysConfig failed: invalid message type")
    return
  }
  result := SysConfigThriftServiceGetSysConfigResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}


type SysConfigThriftServiceProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler SysConfigThriftService
}

func (p *SysConfigThriftServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *SysConfigThriftServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *SysConfigThriftServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewSysConfigThriftServiceProcessor(handler SysConfigThriftService) *SysConfigThriftServiceProcessor {

  self3 := &SysConfigThriftServiceProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self3.processorMap["getSysConfig"] = &sysConfigThriftServiceProcessorGetSysConfig{handler:handler}
return self3
}

func (p *SysConfigThriftServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x4 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x4.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush()
  return false, x4

}

type sysConfigThriftServiceProcessorGetSysConfig struct {
  handler SysConfigThriftService
}

func (p *sysConfigThriftServiceProcessorGetSysConfig) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := SysConfigThriftServiceGetSysConfigArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("getSysConfig", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := SysConfigThriftServiceGetSysConfigResult{}
var retval *SysConfigResponseStruct
  var err2 error
  if retval, err2 = p.handler.GetSysConfig(args.RequestObj); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getSysConfig: " + err2.Error())
    oprot.WriteMessageBegin("getSysConfig", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  } else {
    result.Success = retval
}
  if err2 = oprot.WriteMessageBegin("getSysConfig", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - RequestObj
type SysConfigThriftServiceGetSysConfigArgs struct {
  RequestObj *SysConfigRequestStruct `thrift:"requestObj,1" db:"requestObj" json:"requestObj"`
}

func NewSysConfigThriftServiceGetSysConfigArgs() *SysConfigThriftServiceGetSysConfigArgs {
  return &SysConfigThriftServiceGetSysConfigArgs{}
}

var SysConfigThriftServiceGetSysConfigArgs_RequestObj_DEFAULT *SysConfigRequestStruct
func (p *SysConfigThriftServiceGetSysConfigArgs) GetRequestObj() *SysConfigRequestStruct {
  if !p.IsSetRequestObj() {
    return SysConfigThriftServiceGetSysConfigArgs_RequestObj_DEFAULT
  }
return p.RequestObj
}
func (p *SysConfigThriftServiceGetSysConfigArgs) IsSetRequestObj() bool {
  return p.RequestObj != nil
}

func (p *SysConfigThriftServiceGetSysConfigArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *SysConfigThriftServiceGetSysConfigArgs)  ReadField1(iprot thrift.TProtocol) error {
  p.RequestObj = &SysConfigRequestStruct{}
  if err := p.RequestObj.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.RequestObj), err)
  }
  return nil
}

func (p *SysConfigThriftServiceGetSysConfigArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("getSysConfig_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *SysConfigThriftServiceGetSysConfigArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("requestObj", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:requestObj: ", p), err) }
  if err := p.RequestObj.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.RequestObj), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:requestObj: ", p), err) }
  return err
}

func (p *SysConfigThriftServiceGetSysConfigArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("SysConfigThriftServiceGetSysConfigArgs(%+v)", *p)
}

// Attributes:
//  - Success
type SysConfigThriftServiceGetSysConfigResult struct {
  Success *SysConfigResponseStruct `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewSysConfigThriftServiceGetSysConfigResult() *SysConfigThriftServiceGetSysConfigResult {
  return &SysConfigThriftServiceGetSysConfigResult{}
}

var SysConfigThriftServiceGetSysConfigResult_Success_DEFAULT *SysConfigResponseStruct
func (p *SysConfigThriftServiceGetSysConfigResult) GetSuccess() *SysConfigResponseStruct {
  if !p.IsSetSuccess() {
    return SysConfigThriftServiceGetSysConfigResult_Success_DEFAULT
  }
return p.Success
}
func (p *SysConfigThriftServiceGetSysConfigResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *SysConfigThriftServiceGetSysConfigResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if err := p.ReadField0(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *SysConfigThriftServiceGetSysConfigResult)  ReadField0(iprot thrift.TProtocol) error {
  p.Success = &SysConfigResponseStruct{}
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *SysConfigThriftServiceGetSysConfigResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("getSysConfig_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *SysConfigThriftServiceGetSysConfigResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *SysConfigThriftServiceGetSysConfigResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("SysConfigThriftServiceGetSysConfigResult(%+v)", *p)
}


