// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package sysuseredit

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
//  - Account
//  - Password
//  - Realname
//  - Mobile
//  - Qq
//  - Status
//  - RoleID
//  - CustomerType
//  - CreateTime
//  - UserID
//  - ChengHuiTongTraceLog
type SysUserEditRequestStruct struct {
	Account              string `thrift:"account,1" db:"account" json:"account"`
	Password             string `thrift:"password,2" db:"password" json:"password"`
	Realname             string `thrift:"realname,3" db:"realname" json:"realname"`
	Mobile               string `thrift:"mobile,4" db:"mobile" json:"mobile"`
	Qq                   string `thrift:"qq,5" db:"qq" json:"qq"`
	Status               int32  `thrift:"status,6" db:"status" json:"status"`
	RoleID               int32  `thrift:"role_id,7" db:"role_id" json:"role_id"`
	CustomerType         int32  `thrift:"customer_type,8" db:"customer_type" json:"customer_type"`
	CreateTime           string `thrift:"create_time,9" db:"create_time" json:"create_time"`
	UserID               int32  `thrift:"user_id,10" db:"user_id" json:"user_id"`
	ChengHuiTongTraceLog string `thrift:"chengHuiTongTraceLog,11" db:"chengHuiTongTraceLog" json:"chengHuiTongTraceLog"`
}

// func NewSysUserEditRequestStruct() *SysUserEditRequestStruct {
//   return &SysUserEditRequestStruct{}
// }

func (p *SysUserEditRequestStruct) GetAccount() string {
	return p.Account
}

func (p *SysUserEditRequestStruct) GetPassword() string {
	return p.Password
}

func (p *SysUserEditRequestStruct) GetRealname() string {
	return p.Realname
}

func (p *SysUserEditRequestStruct) GetMobile() string {
	return p.Mobile
}

func (p *SysUserEditRequestStruct) GetQq() string {
	return p.Qq
}

func (p *SysUserEditRequestStruct) GetStatus() int32 {
	return p.Status
}

func (p *SysUserEditRequestStruct) GetRoleID() int32 {
	return p.RoleID
}

func (p *SysUserEditRequestStruct) GetCustomerType() int32 {
	return p.CustomerType
}

func (p *SysUserEditRequestStruct) GetCreateTime() string {
	return p.CreateTime
}

func (p *SysUserEditRequestStruct) GetUserID() int32 {
	return p.UserID
}

func (p *SysUserEditRequestStruct) GetChengHuiTongTraceLog() string {
	return p.ChengHuiTongTraceLog
}
func (p *SysUserEditRequestStruct) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
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
		case 5:
			if err := p.ReadField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.ReadField6(iprot); err != nil {
				return err
			}
		case 7:
			if err := p.ReadField7(iprot); err != nil {
				return err
			}
		case 8:
			if err := p.ReadField8(iprot); err != nil {
				return err
			}
		case 9:
			if err := p.ReadField9(iprot); err != nil {
				return err
			}
		case 10:
			if err := p.ReadField10(iprot); err != nil {
				return err
			}
		case 11:
			if err := p.ReadField11(iprot); err != nil {
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

func (p *SysUserEditRequestStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Account = v
	}
	return nil
}

func (p *SysUserEditRequestStruct) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Password = v
	}
	return nil
}

func (p *SysUserEditRequestStruct) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Realname = v
	}
	return nil
}

func (p *SysUserEditRequestStruct) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.Mobile = v
	}
	return nil
}

func (p *SysUserEditRequestStruct) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.Qq = v
	}
	return nil
}

func (p *SysUserEditRequestStruct) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.Status = v
	}
	return nil
}

func (p *SysUserEditRequestStruct) ReadField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.RoleID = v
	}
	return nil
}

func (p *SysUserEditRequestStruct) ReadField8(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 8: ", err)
	} else {
		p.CustomerType = v
	}
	return nil
}

func (p *SysUserEditRequestStruct) ReadField9(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 9: ", err)
	} else {
		p.CreateTime = v
	}
	return nil
}

func (p *SysUserEditRequestStruct) ReadField10(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 10: ", err)
	} else {
		p.UserID = v
	}
	return nil
}

func (p *SysUserEditRequestStruct) ReadField11(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 11: ", err)
	} else {
		p.ChengHuiTongTraceLog = v
	}
	return nil
}

func (p *SysUserEditRequestStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("SysUserEditRequestStruct"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
		if err := p.writeField3(oprot); err != nil {
			return err
		}
		if err := p.writeField4(oprot); err != nil {
			return err
		}
		if err := p.writeField5(oprot); err != nil {
			return err
		}
		if err := p.writeField6(oprot); err != nil {
			return err
		}
		if err := p.writeField7(oprot); err != nil {
			return err
		}
		if err := p.writeField8(oprot); err != nil {
			return err
		}
		if err := p.writeField9(oprot); err != nil {
			return err
		}
		if err := p.writeField10(oprot); err != nil {
			return err
		}
		if err := p.writeField11(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *SysUserEditRequestStruct) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("account", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:account: ", p), err)
	}
	if err := oprot.WriteString(string(p.Account)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.account (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:account: ", p), err)
	}
	return err
}

func (p *SysUserEditRequestStruct) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("password", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:password: ", p), err)
	}
	if err := oprot.WriteString(string(p.Password)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.password (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:password: ", p), err)
	}
	return err
}

func (p *SysUserEditRequestStruct) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("realname", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:realname: ", p), err)
	}
	if err := oprot.WriteString(string(p.Realname)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.realname (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:realname: ", p), err)
	}
	return err
}

func (p *SysUserEditRequestStruct) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("mobile", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:mobile: ", p), err)
	}
	if err := oprot.WriteString(string(p.Mobile)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.mobile (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:mobile: ", p), err)
	}
	return err
}

func (p *SysUserEditRequestStruct) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("qq", thrift.STRING, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:qq: ", p), err)
	}
	if err := oprot.WriteString(string(p.Qq)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.qq (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:qq: ", p), err)
	}
	return err
}

func (p *SysUserEditRequestStruct) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("status", thrift.I32, 6); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:status: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Status)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.status (6) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 6:status: ", p), err)
	}
	return err
}

func (p *SysUserEditRequestStruct) writeField7(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("role_id", thrift.I32, 7); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:role_id: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.RoleID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.role_id (7) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 7:role_id: ", p), err)
	}
	return err
}

func (p *SysUserEditRequestStruct) writeField8(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("customer_type", thrift.I32, 8); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:customer_type: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.CustomerType)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.customer_type (8) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 8:customer_type: ", p), err)
	}
	return err
}

func (p *SysUserEditRequestStruct) writeField9(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("create_time", thrift.STRING, 9); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 9:create_time: ", p), err)
	}
	if err := oprot.WriteString(string(p.CreateTime)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.create_time (9) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 9:create_time: ", p), err)
	}
	return err
}

func (p *SysUserEditRequestStruct) writeField10(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("user_id", thrift.I32, 10); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 10:user_id: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.UserID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.user_id (10) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 10:user_id: ", p), err)
	}
	return err
}

func (p *SysUserEditRequestStruct) writeField11(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("chengHuiTongTraceLog", thrift.STRING, 11); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 11:chengHuiTongTraceLog: ", p), err)
	}
	if err := oprot.WriteString(string(p.ChengHuiTongTraceLog)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.chengHuiTongTraceLog (11) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 11:chengHuiTongTraceLog: ", p), err)
	}
	return err
}

func (p *SysUserEditRequestStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SysUserEditRequestStruct(%+v)", *p)
}

// Attributes:
//  - Status
//  - Msg
type SysUserEditResponseStruct struct {
	Status int32 `thrift:"status,1" db:"status" json:"status"`
	// unused field # 2
	Msg string `thrift:"msg,3" db:"msg" json:"msg"`
}

func NewSysUserEditResponseStruct() *SysUserEditResponseStruct {
	return &SysUserEditResponseStruct{}
}

func (p *SysUserEditResponseStruct) GetStatus() int32 {
	return p.Status
}

func (p *SysUserEditResponseStruct) GetMsg() string {
	return p.Msg
}
func (p *SysUserEditResponseStruct) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
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

func (p *SysUserEditResponseStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Status = v
	}
	return nil
}

func (p *SysUserEditResponseStruct) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Msg = v
	}
	return nil
}

func (p *SysUserEditResponseStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("SysUserEditResponseStruct"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField3(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *SysUserEditResponseStruct) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("status", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:status: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Status)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.status (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:status: ", p), err)
	}
	return err
}

func (p *SysUserEditResponseStruct) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("msg", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:msg: ", p), err)
	}
	if err := oprot.WriteString(string(p.Msg)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.msg (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:msg: ", p), err)
	}
	return err
}

func (p *SysUserEditResponseStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SysUserEditResponseStruct(%+v)", *p)
}

type SysUserEditThriftService interface {
	// Parameters:
	//  - RequestObj
	EditSysUser(requestObj *SysUserEditRequestStruct) (r *SysUserEditResponseStruct, err error)
}

type SysUserEditThriftServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewSysUserEditThriftServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *SysUserEditThriftServiceClient {
	return &SysUserEditThriftServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewSysUserEditThriftServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *SysUserEditThriftServiceClient {
	return &SysUserEditThriftServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - RequestObj
func (p *SysUserEditThriftServiceClient) EditSysUser(requestObj *SysUserEditRequestStruct) (r *SysUserEditResponseStruct, err error) {
	if err = p.sendEditSysUser(requestObj); err != nil {
		return
	}
	return p.recvEditSysUser()
}

func (p *SysUserEditThriftServiceClient) sendEditSysUser(requestObj *SysUserEditRequestStruct) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("editSysUser", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := SysUserEditThriftServiceEditSysUserArgs{
		RequestObj: requestObj,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *SysUserEditThriftServiceClient) recvEditSysUser() (value *SysUserEditResponseStruct, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "editSysUser" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "editSysUser failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "editSysUser failed: out of sequence response")
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
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "editSysUser failed: invalid message type")
		return
	}
	result := SysUserEditThriftServiceEditSysUserResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type SysUserEditThriftServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      SysUserEditThriftService
}

func (p *SysUserEditThriftServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *SysUserEditThriftServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *SysUserEditThriftServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewSysUserEditThriftServiceProcessor(handler SysUserEditThriftService) *SysUserEditThriftServiceProcessor {

	self2 := &SysUserEditThriftServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self2.processorMap["editSysUser"] = &sysUserEditThriftServiceProcessorEditSysUser{handler: handler}
	return self2
}

func (p *SysUserEditThriftServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x3.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x3

}

type sysUserEditThriftServiceProcessorEditSysUser struct {
	handler SysUserEditThriftService
}

func (p *sysUserEditThriftServiceProcessorEditSysUser) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := SysUserEditThriftServiceEditSysUserArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("editSysUser", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := SysUserEditThriftServiceEditSysUserResult{}
	var retval *SysUserEditResponseStruct
	var err2 error
	if retval, err2 = p.handler.EditSysUser(args.RequestObj); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing editSysUser: "+err2.Error())
		oprot.WriteMessageBegin("editSysUser", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("editSysUser", thrift.REPLY, seqId); err2 != nil {
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
type SysUserEditThriftServiceEditSysUserArgs struct {
	RequestObj *SysUserEditRequestStruct `thrift:"requestObj,1" db:"requestObj" json:"requestObj"`
}

func NewSysUserEditThriftServiceEditSysUserArgs() *SysUserEditThriftServiceEditSysUserArgs {
	return &SysUserEditThriftServiceEditSysUserArgs{}
}

var SysUserEditThriftServiceEditSysUserArgs_RequestObj_DEFAULT *SysUserEditRequestStruct

func (p *SysUserEditThriftServiceEditSysUserArgs) GetRequestObj() *SysUserEditRequestStruct {
	if !p.IsSetRequestObj() {
		return SysUserEditThriftServiceEditSysUserArgs_RequestObj_DEFAULT
	}
	return p.RequestObj
}
func (p *SysUserEditThriftServiceEditSysUserArgs) IsSetRequestObj() bool {
	return p.RequestObj != nil
}

func (p *SysUserEditThriftServiceEditSysUserArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
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

func (p *SysUserEditThriftServiceEditSysUserArgs) ReadField1(iprot thrift.TProtocol) error {
	p.RequestObj = &SysUserEditRequestStruct{}
	if err := p.RequestObj.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.RequestObj), err)
	}
	return nil
}

func (p *SysUserEditThriftServiceEditSysUserArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("editSysUser_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *SysUserEditThriftServiceEditSysUserArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("requestObj", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:requestObj: ", p), err)
	}
	if err := p.RequestObj.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.RequestObj), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:requestObj: ", p), err)
	}
	return err
}

func (p *SysUserEditThriftServiceEditSysUserArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SysUserEditThriftServiceEditSysUserArgs(%+v)", *p)
}

// Attributes:
//  - Success
type SysUserEditThriftServiceEditSysUserResult struct {
	Success *SysUserEditResponseStruct `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewSysUserEditThriftServiceEditSysUserResult() *SysUserEditThriftServiceEditSysUserResult {
	return &SysUserEditThriftServiceEditSysUserResult{}
}

var SysUserEditThriftServiceEditSysUserResult_Success_DEFAULT *SysUserEditResponseStruct

func (p *SysUserEditThriftServiceEditSysUserResult) GetSuccess() *SysUserEditResponseStruct {
	if !p.IsSetSuccess() {
		return SysUserEditThriftServiceEditSysUserResult_Success_DEFAULT
	}
	return p.Success
}
func (p *SysUserEditThriftServiceEditSysUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SysUserEditThriftServiceEditSysUserResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
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

func (p *SysUserEditThriftServiceEditSysUserResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &SysUserEditResponseStruct{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *SysUserEditThriftServiceEditSysUserResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("editSysUser_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *SysUserEditThriftServiceEditSysUserResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *SysUserEditThriftServiceEditSysUserResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SysUserEditThriftServiceEditSysUserResult(%+v)", *p)
}