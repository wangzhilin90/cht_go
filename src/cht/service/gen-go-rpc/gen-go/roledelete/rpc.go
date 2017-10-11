// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package roledelete

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
//  - RoleIDStr
//  - ChengHuiTongTraceLog
type RoleDeleteRequestStruct struct {
  RoleIDStr string `thrift:"role_id_str,1" db:"role_id_str" json:"role_id_str"`
  ChengHuiTongTraceLog string `thrift:"chengHuiTongTraceLog,2" db:"chengHuiTongTraceLog" json:"chengHuiTongTraceLog"`
}

func NewRoleDeleteRequestStruct() *RoleDeleteRequestStruct {
  return &RoleDeleteRequestStruct{}
}


func (p *RoleDeleteRequestStruct) GetRoleIDStr() string {
  return p.RoleIDStr
}

func (p *RoleDeleteRequestStruct) GetChengHuiTongTraceLog() string {
  return p.ChengHuiTongTraceLog
}
func (p *RoleDeleteRequestStruct) Read(iprot thrift.TProtocol) error {
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

func (p *RoleDeleteRequestStruct)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.RoleIDStr = v
}
  return nil
}

func (p *RoleDeleteRequestStruct)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.ChengHuiTongTraceLog = v
}
  return nil
}

func (p *RoleDeleteRequestStruct) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("RoleDeleteRequestStruct"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *RoleDeleteRequestStruct) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("role_id_str", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:role_id_str: ", p), err) }
  if err := oprot.WriteString(string(p.RoleIDStr)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.role_id_str (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:role_id_str: ", p), err) }
  return err
}

func (p *RoleDeleteRequestStruct) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("chengHuiTongTraceLog", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:chengHuiTongTraceLog: ", p), err) }
  if err := oprot.WriteString(string(p.ChengHuiTongTraceLog)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.chengHuiTongTraceLog (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:chengHuiTongTraceLog: ", p), err) }
  return err
}

func (p *RoleDeleteRequestStruct) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("RoleDeleteRequestStruct(%+v)", *p)
}

// Attributes:
//  - Status
//  - Msg
type RoleDeleteResponseStruct struct {
  Status int32 `thrift:"status,1" db:"status" json:"status"`
  Msg string `thrift:"msg,2" db:"msg" json:"msg"`
}

func NewRoleDeleteResponseStruct() *RoleDeleteResponseStruct {
  return &RoleDeleteResponseStruct{}
}


func (p *RoleDeleteResponseStruct) GetStatus() int32 {
  return p.Status
}

func (p *RoleDeleteResponseStruct) GetMsg() string {
  return p.Msg
}
func (p *RoleDeleteResponseStruct) Read(iprot thrift.TProtocol) error {
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

func (p *RoleDeleteResponseStruct)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Status = v
}
  return nil
}

func (p *RoleDeleteResponseStruct)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Msg = v
}
  return nil
}

func (p *RoleDeleteResponseStruct) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("RoleDeleteResponseStruct"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *RoleDeleteResponseStruct) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("status", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:status: ", p), err) }
  if err := oprot.WriteI32(int32(p.Status)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.status (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:status: ", p), err) }
  return err
}

func (p *RoleDeleteResponseStruct) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("msg", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:msg: ", p), err) }
  if err := oprot.WriteString(string(p.Msg)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.msg (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:msg: ", p), err) }
  return err
}

func (p *RoleDeleteResponseStruct) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("RoleDeleteResponseStruct(%+v)", *p)
}

type RoleDeleteThriftService interface {
  // Parameters:
  //  - RequestObj
  DeleteRole(requestObj *RoleDeleteRequestStruct) (r *RoleDeleteResponseStruct, err error)
}

type RoleDeleteThriftServiceClient struct {
  Transport thrift.TTransport
  ProtocolFactory thrift.TProtocolFactory
  InputProtocol thrift.TProtocol
  OutputProtocol thrift.TProtocol
  SeqId int32
}

func NewRoleDeleteThriftServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *RoleDeleteThriftServiceClient {
  return &RoleDeleteThriftServiceClient{Transport: t,
    ProtocolFactory: f,
    InputProtocol: f.GetProtocol(t),
    OutputProtocol: f.GetProtocol(t),
    SeqId: 0,
  }
}

func NewRoleDeleteThriftServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *RoleDeleteThriftServiceClient {
  return &RoleDeleteThriftServiceClient{Transport: t,
    ProtocolFactory: nil,
    InputProtocol: iprot,
    OutputProtocol: oprot,
    SeqId: 0,
  }
}

// Parameters:
//  - RequestObj
func (p *RoleDeleteThriftServiceClient) DeleteRole(requestObj *RoleDeleteRequestStruct) (r *RoleDeleteResponseStruct, err error) {
  if err = p.sendDeleteRole(requestObj); err != nil { return }
  return p.recvDeleteRole()
}

func (p *RoleDeleteThriftServiceClient) sendDeleteRole(requestObj *RoleDeleteRequestStruct)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("deleteRole", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := RoleDeleteThriftServiceDeleteRoleArgs{
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


func (p *RoleDeleteThriftServiceClient) recvDeleteRole() (value *RoleDeleteResponseStruct, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "deleteRole" {
    err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "deleteRole failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "deleteRole failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error1 error
    error1, err = error0.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error1
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "deleteRole failed: invalid message type")
    return
  }
  result := RoleDeleteThriftServiceDeleteRoleResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}


type RoleDeleteThriftServiceProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler RoleDeleteThriftService
}

func (p *RoleDeleteThriftServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *RoleDeleteThriftServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *RoleDeleteThriftServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewRoleDeleteThriftServiceProcessor(handler RoleDeleteThriftService) *RoleDeleteThriftServiceProcessor {

  self2 := &RoleDeleteThriftServiceProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self2.processorMap["deleteRole"] = &roleDeleteThriftServiceProcessorDeleteRole{handler:handler}
return self2
}

func (p *RoleDeleteThriftServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x3.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush()
  return false, x3

}

type roleDeleteThriftServiceProcessorDeleteRole struct {
  handler RoleDeleteThriftService
}

func (p *roleDeleteThriftServiceProcessorDeleteRole) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := RoleDeleteThriftServiceDeleteRoleArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("deleteRole", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := RoleDeleteThriftServiceDeleteRoleResult{}
var retval *RoleDeleteResponseStruct
  var err2 error
  if retval, err2 = p.handler.DeleteRole(args.RequestObj); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing deleteRole: " + err2.Error())
    oprot.WriteMessageBegin("deleteRole", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  } else {
    result.Success = retval
}
  if err2 = oprot.WriteMessageBegin("deleteRole", thrift.REPLY, seqId); err2 != nil {
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
type RoleDeleteThriftServiceDeleteRoleArgs struct {
  RequestObj *RoleDeleteRequestStruct `thrift:"requestObj,1" db:"requestObj" json:"requestObj"`
}

func NewRoleDeleteThriftServiceDeleteRoleArgs() *RoleDeleteThriftServiceDeleteRoleArgs {
  return &RoleDeleteThriftServiceDeleteRoleArgs{}
}

var RoleDeleteThriftServiceDeleteRoleArgs_RequestObj_DEFAULT *RoleDeleteRequestStruct
func (p *RoleDeleteThriftServiceDeleteRoleArgs) GetRequestObj() *RoleDeleteRequestStruct {
  if !p.IsSetRequestObj() {
    return RoleDeleteThriftServiceDeleteRoleArgs_RequestObj_DEFAULT
  }
return p.RequestObj
}
func (p *RoleDeleteThriftServiceDeleteRoleArgs) IsSetRequestObj() bool {
  return p.RequestObj != nil
}

func (p *RoleDeleteThriftServiceDeleteRoleArgs) Read(iprot thrift.TProtocol) error {
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

func (p *RoleDeleteThriftServiceDeleteRoleArgs)  ReadField1(iprot thrift.TProtocol) error {
  p.RequestObj = &RoleDeleteRequestStruct{}
  if err := p.RequestObj.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.RequestObj), err)
  }
  return nil
}

func (p *RoleDeleteThriftServiceDeleteRoleArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("deleteRole_args"); err != nil {
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

func (p *RoleDeleteThriftServiceDeleteRoleArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("requestObj", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:requestObj: ", p), err) }
  if err := p.RequestObj.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.RequestObj), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:requestObj: ", p), err) }
  return err
}

func (p *RoleDeleteThriftServiceDeleteRoleArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("RoleDeleteThriftServiceDeleteRoleArgs(%+v)", *p)
}

// Attributes:
//  - Success
type RoleDeleteThriftServiceDeleteRoleResult struct {
  Success *RoleDeleteResponseStruct `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewRoleDeleteThriftServiceDeleteRoleResult() *RoleDeleteThriftServiceDeleteRoleResult {
  return &RoleDeleteThriftServiceDeleteRoleResult{}
}

var RoleDeleteThriftServiceDeleteRoleResult_Success_DEFAULT *RoleDeleteResponseStruct
func (p *RoleDeleteThriftServiceDeleteRoleResult) GetSuccess() *RoleDeleteResponseStruct {
  if !p.IsSetSuccess() {
    return RoleDeleteThriftServiceDeleteRoleResult_Success_DEFAULT
  }
return p.Success
}
func (p *RoleDeleteThriftServiceDeleteRoleResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *RoleDeleteThriftServiceDeleteRoleResult) Read(iprot thrift.TProtocol) error {
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

func (p *RoleDeleteThriftServiceDeleteRoleResult)  ReadField0(iprot thrift.TProtocol) error {
  p.Success = &RoleDeleteResponseStruct{}
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *RoleDeleteThriftServiceDeleteRoleResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("deleteRole_result"); err != nil {
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

func (p *RoleDeleteThriftServiceDeleteRoleResult) writeField0(oprot thrift.TProtocol) (err error) {
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

func (p *RoleDeleteThriftServiceDeleteRoleResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("RoleDeleteThriftServiceDeleteRoleResult(%+v)", *p)
}

