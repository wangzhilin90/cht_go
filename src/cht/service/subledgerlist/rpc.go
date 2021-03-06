// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package subledgerlist

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
//  - HsZhuanrangrenStr
//  - ChengHuiTongTraceLog
type SubledgerListRequestStruct struct {
	HsZhuanrangrenStr    string `thrift:"hs_zhuanrangren_str,1" db:"hs_zhuanrangren_str" json:"hs_zhuanrangren_str"`
	ChengHuiTongTraceLog string `thrift:"chengHuiTongTraceLog,2" db:"chengHuiTongTraceLog" json:"chengHuiTongTraceLog"`
}

// func NewSubledgerListRequestStruct() *SubledgerListRequestStruct {
// 	return &SubledgerListRequestStruct{}
// }

func (p *SubledgerListRequestStruct) GetHsZhuanrangrenStr() string {
	return p.HsZhuanrangrenStr
}

func (p *SubledgerListRequestStruct) GetChengHuiTongTraceLog() string {
	return p.ChengHuiTongTraceLog
}
func (p *SubledgerListRequestStruct) Read(iprot thrift.TProtocol) error {
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

func (p *SubledgerListRequestStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.HsZhuanrangrenStr = v
	}
	return nil
}

func (p *SubledgerListRequestStruct) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.ChengHuiTongTraceLog = v
	}
	return nil
}

func (p *SubledgerListRequestStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("SubledgerListRequestStruct"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
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

func (p *SubledgerListRequestStruct) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("hs_zhuanrangren_str", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:hs_zhuanrangren_str: ", p), err)
	}
	if err := oprot.WriteString(string(p.HsZhuanrangrenStr)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.hs_zhuanrangren_str (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:hs_zhuanrangren_str: ", p), err)
	}
	return err
}

func (p *SubledgerListRequestStruct) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("chengHuiTongTraceLog", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:chengHuiTongTraceLog: ", p), err)
	}
	if err := oprot.WriteString(string(p.ChengHuiTongTraceLog)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.chengHuiTongTraceLog (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:chengHuiTongTraceLog: ", p), err)
	}
	return err
}

func (p *SubledgerListRequestStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SubledgerListRequestStruct(%+v)", *p)
}

// Attributes:
//  - UserID
//  - Realname
//  - CardID
type SubledgerDetailsStruct struct {
	UserID   int32  `thrift:"user_id,1" db:"user_id" json:"user_id"`
	Realname string `thrift:"realname,2" db:"realname" json:"realname"`
	CardID   string `thrift:"card_id,3" db:"card_id" json:"card_id"`
}

func NewSubledgerDetailsStruct() *SubledgerDetailsStruct {
	return &SubledgerDetailsStruct{}
}

func (p *SubledgerDetailsStruct) GetUserID() int32 {
	return p.UserID
}

func (p *SubledgerDetailsStruct) GetRealname() string {
	return p.Realname
}

func (p *SubledgerDetailsStruct) GetCardID() string {
	return p.CardID
}
func (p *SubledgerDetailsStruct) Read(iprot thrift.TProtocol) error {
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

func (p *SubledgerDetailsStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.UserID = v
	}
	return nil
}

func (p *SubledgerDetailsStruct) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Realname = v
	}
	return nil
}

func (p *SubledgerDetailsStruct) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.CardID = v
	}
	return nil
}

func (p *SubledgerDetailsStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("SubledgerDetailsStruct"); err != nil {
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
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *SubledgerDetailsStruct) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("user_id", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:user_id: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.UserID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.user_id (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:user_id: ", p), err)
	}
	return err
}

func (p *SubledgerDetailsStruct) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("realname", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:realname: ", p), err)
	}
	if err := oprot.WriteString(string(p.Realname)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.realname (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:realname: ", p), err)
	}
	return err
}

func (p *SubledgerDetailsStruct) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("card_id", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:card_id: ", p), err)
	}
	if err := oprot.WriteString(string(p.CardID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.card_id (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:card_id: ", p), err)
	}
	return err
}

func (p *SubledgerDetailsStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SubledgerDetailsStruct(%+v)", *p)
}

// Attributes:
//  - Status
//  - Msg
//  - SubledgerList
type SubledgerListResponseStruct struct {
	Status        int32                     `thrift:"status,1" db:"status" json:"status"`
	Msg           string                    `thrift:"msg,2" db:"msg" json:"msg"`
	SubledgerList []*SubledgerDetailsStruct `thrift:"SubledgerList,3" db:"SubledgerList" json:"SubledgerList"`
}

func NewSubledgerListResponseStruct() *SubledgerListResponseStruct {
	return &SubledgerListResponseStruct{}
}

func (p *SubledgerListResponseStruct) GetStatus() int32 {
	return p.Status
}

func (p *SubledgerListResponseStruct) GetMsg() string {
	return p.Msg
}

func (p *SubledgerListResponseStruct) GetSubledgerList() []*SubledgerDetailsStruct {
	return p.SubledgerList
}
func (p *SubledgerListResponseStruct) Read(iprot thrift.TProtocol) error {
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

func (p *SubledgerListResponseStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Status = v
	}
	return nil
}

func (p *SubledgerListResponseStruct) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Msg = v
	}
	return nil
}

func (p *SubledgerListResponseStruct) ReadField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*SubledgerDetailsStruct, 0, size)
	p.SubledgerList = tSlice
	for i := 0; i < size; i++ {
		_elem0 := &SubledgerDetailsStruct{}
		if err := _elem0.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem0), err)
		}
		p.SubledgerList = append(p.SubledgerList, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *SubledgerListResponseStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("SubledgerListResponseStruct"); err != nil {
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
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *SubledgerListResponseStruct) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *SubledgerListResponseStruct) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("msg", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:msg: ", p), err)
	}
	if err := oprot.WriteString(string(p.Msg)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.msg (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:msg: ", p), err)
	}
	return err
}

func (p *SubledgerListResponseStruct) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("SubledgerList", thrift.LIST, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:SubledgerList: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.SubledgerList)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.SubledgerList {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:SubledgerList: ", p), err)
	}
	return err
}

func (p *SubledgerListResponseStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SubledgerListResponseStruct(%+v)", *p)
}

type SubledgerListThriftService interface {
	// Parameters:
	//  - RequestObj
	GetSubledgerList(requestObj *SubledgerListRequestStruct) (r *SubledgerListResponseStruct, err error)
}

type SubledgerListThriftServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewSubledgerListThriftServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *SubledgerListThriftServiceClient {
	return &SubledgerListThriftServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewSubledgerListThriftServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *SubledgerListThriftServiceClient {
	return &SubledgerListThriftServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - RequestObj
func (p *SubledgerListThriftServiceClient) GetSubledgerList(requestObj *SubledgerListRequestStruct) (r *SubledgerListResponseStruct, err error) {
	if err = p.sendGetSubledgerList(requestObj); err != nil {
		return
	}
	return p.recvGetSubledgerList()
}

func (p *SubledgerListThriftServiceClient) sendGetSubledgerList(requestObj *SubledgerListRequestStruct) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getSubledgerList", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := SubledgerListThriftServiceGetSubledgerListArgs{
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

func (p *SubledgerListThriftServiceClient) recvGetSubledgerList() (value *SubledgerListResponseStruct, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "getSubledgerList" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getSubledgerList failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getSubledgerList failed: out of sequence response")
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
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getSubledgerList failed: invalid message type")
		return
	}
	result := SubledgerListThriftServiceGetSubledgerListResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type SubledgerListThriftServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      SubledgerListThriftService
}

func (p *SubledgerListThriftServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *SubledgerListThriftServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *SubledgerListThriftServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewSubledgerListThriftServiceProcessor(handler SubledgerListThriftService) *SubledgerListThriftServiceProcessor {

	self3 := &SubledgerListThriftServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self3.processorMap["getSubledgerList"] = &subledgerListThriftServiceProcessorGetSubledgerList{handler: handler}
	return self3
}

func (p *SubledgerListThriftServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x4 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x4.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x4

}

type subledgerListThriftServiceProcessorGetSubledgerList struct {
	handler SubledgerListThriftService
}

func (p *subledgerListThriftServiceProcessorGetSubledgerList) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := SubledgerListThriftServiceGetSubledgerListArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getSubledgerList", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := SubledgerListThriftServiceGetSubledgerListResult{}
	var retval *SubledgerListResponseStruct
	var err2 error
	if retval, err2 = p.handler.GetSubledgerList(args.RequestObj); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getSubledgerList: "+err2.Error())
		oprot.WriteMessageBegin("getSubledgerList", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getSubledgerList", thrift.REPLY, seqId); err2 != nil {
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
type SubledgerListThriftServiceGetSubledgerListArgs struct {
	RequestObj *SubledgerListRequestStruct `thrift:"requestObj,1" db:"requestObj" json:"requestObj"`
}

func NewSubledgerListThriftServiceGetSubledgerListArgs() *SubledgerListThriftServiceGetSubledgerListArgs {
	return &SubledgerListThriftServiceGetSubledgerListArgs{}
}

var SubledgerListThriftServiceGetSubledgerListArgs_RequestObj_DEFAULT *SubledgerListRequestStruct

func (p *SubledgerListThriftServiceGetSubledgerListArgs) GetRequestObj() *SubledgerListRequestStruct {
	if !p.IsSetRequestObj() {
		return SubledgerListThriftServiceGetSubledgerListArgs_RequestObj_DEFAULT
	}
	return p.RequestObj
}
func (p *SubledgerListThriftServiceGetSubledgerListArgs) IsSetRequestObj() bool {
	return p.RequestObj != nil
}

func (p *SubledgerListThriftServiceGetSubledgerListArgs) Read(iprot thrift.TProtocol) error {
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

func (p *SubledgerListThriftServiceGetSubledgerListArgs) ReadField1(iprot thrift.TProtocol) error {
	p.RequestObj = &SubledgerListRequestStruct{}
	if err := p.RequestObj.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.RequestObj), err)
	}
	return nil
}

func (p *SubledgerListThriftServiceGetSubledgerListArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getSubledgerList_args"); err != nil {
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

func (p *SubledgerListThriftServiceGetSubledgerListArgs) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *SubledgerListThriftServiceGetSubledgerListArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SubledgerListThriftServiceGetSubledgerListArgs(%+v)", *p)
}

// Attributes:
//  - Success
type SubledgerListThriftServiceGetSubledgerListResult struct {
	Success *SubledgerListResponseStruct `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewSubledgerListThriftServiceGetSubledgerListResult() *SubledgerListThriftServiceGetSubledgerListResult {
	return &SubledgerListThriftServiceGetSubledgerListResult{}
}

var SubledgerListThriftServiceGetSubledgerListResult_Success_DEFAULT *SubledgerListResponseStruct

func (p *SubledgerListThriftServiceGetSubledgerListResult) GetSuccess() *SubledgerListResponseStruct {
	if !p.IsSetSuccess() {
		return SubledgerListThriftServiceGetSubledgerListResult_Success_DEFAULT
	}
	return p.Success
}
func (p *SubledgerListThriftServiceGetSubledgerListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SubledgerListThriftServiceGetSubledgerListResult) Read(iprot thrift.TProtocol) error {
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

func (p *SubledgerListThriftServiceGetSubledgerListResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &SubledgerListResponseStruct{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *SubledgerListThriftServiceGetSubledgerListResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getSubledgerList_result"); err != nil {
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

func (p *SubledgerListThriftServiceGetSubledgerListResult) writeField0(oprot thrift.TProtocol) (err error) {
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

func (p *SubledgerListThriftServiceGetSubledgerListResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SubledgerListThriftServiceGetSubledgerListResult(%+v)", *p)
}
