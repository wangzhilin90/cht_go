// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package goodsdel

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
//  - ChengHuiTongTraceLog
type GoodsDeLRequestStruct struct {
	ID                   int32  `thrift:"id,1" db:"id" json:"id"`
	ChengHuiTongTraceLog string `thrift:"chengHuiTongTraceLog,2" db:"chengHuiTongTraceLog" json:"chengHuiTongTraceLog"`
}

// func NewGoodsDeLRequestStruct() *GoodsDeLRequestStruct {
//   return &GoodsDeLRequestStruct{}
// }

func (p *GoodsDeLRequestStruct) GetID() int32 {
	return p.ID
}

func (p *GoodsDeLRequestStruct) GetChengHuiTongTraceLog() string {
	return p.ChengHuiTongTraceLog
}
func (p *GoodsDeLRequestStruct) Read(iprot thrift.TProtocol) error {
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

func (p *GoodsDeLRequestStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ID = v
	}
	return nil
}

func (p *GoodsDeLRequestStruct) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.ChengHuiTongTraceLog = v
	}
	return nil
}

func (p *GoodsDeLRequestStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GoodsDeLRequestStruct"); err != nil {
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

func (p *GoodsDeLRequestStruct) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("id", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:id: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.ID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.id (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:id: ", p), err)
	}
	return err
}

func (p *GoodsDeLRequestStruct) writeField2(oprot thrift.TProtocol) (err error) {
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

func (p *GoodsDeLRequestStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GoodsDeLRequestStruct(%+v)", *p)
}

// Attributes:
//  - Status
//  - Msg
type GoodsDeLResponseStruct struct {
	Status int32  `thrift:"status,1" db:"status" json:"status"`
	Msg    string `thrift:"msg,2" db:"msg" json:"msg"`
}

func NewGoodsDeLResponseStruct() *GoodsDeLResponseStruct {
	return &GoodsDeLResponseStruct{}
}

func (p *GoodsDeLResponseStruct) GetStatus() int32 {
	return p.Status
}

func (p *GoodsDeLResponseStruct) GetMsg() string {
	return p.Msg
}
func (p *GoodsDeLResponseStruct) Read(iprot thrift.TProtocol) error {
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

func (p *GoodsDeLResponseStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Status = v
	}
	return nil
}

func (p *GoodsDeLResponseStruct) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Msg = v
	}
	return nil
}

func (p *GoodsDeLResponseStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GoodsDeLResponseStruct"); err != nil {
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

func (p *GoodsDeLResponseStruct) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *GoodsDeLResponseStruct) writeField2(oprot thrift.TProtocol) (err error) {
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

func (p *GoodsDeLResponseStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GoodsDeLResponseStruct(%+v)", *p)
}

type GoodsDeLThriftService interface {
	// Parameters:
	//  - RequestObj
	DelGoods(requestObj *GoodsDeLRequestStruct) (r *GoodsDeLResponseStruct, err error)
}

type GoodsDeLThriftServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewGoodsDeLThriftServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *GoodsDeLThriftServiceClient {
	return &GoodsDeLThriftServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewGoodsDeLThriftServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *GoodsDeLThriftServiceClient {
	return &GoodsDeLThriftServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - RequestObj
func (p *GoodsDeLThriftServiceClient) DelGoods(requestObj *GoodsDeLRequestStruct) (r *GoodsDeLResponseStruct, err error) {
	if err = p.sendDelGoods(requestObj); err != nil {
		return
	}
	return p.recvDelGoods()
}

func (p *GoodsDeLThriftServiceClient) sendDelGoods(requestObj *GoodsDeLRequestStruct) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("delGoods", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := GoodsDeLThriftServiceDelGoodsArgs{
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

func (p *GoodsDeLThriftServiceClient) recvDelGoods() (value *GoodsDeLResponseStruct, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "delGoods" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "delGoods failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "delGoods failed: out of sequence response")
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
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "delGoods failed: invalid message type")
		return
	}
	result := GoodsDeLThriftServiceDelGoodsResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type GoodsDeLThriftServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      GoodsDeLThriftService
}

func (p *GoodsDeLThriftServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *GoodsDeLThriftServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *GoodsDeLThriftServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewGoodsDeLThriftServiceProcessor(handler GoodsDeLThriftService) *GoodsDeLThriftServiceProcessor {

	self2 := &GoodsDeLThriftServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self2.processorMap["delGoods"] = &goodsDeLThriftServiceProcessorDelGoods{handler: handler}
	return self2
}

func (p *GoodsDeLThriftServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
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

type goodsDeLThriftServiceProcessorDelGoods struct {
	handler GoodsDeLThriftService
}

func (p *goodsDeLThriftServiceProcessorDelGoods) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := GoodsDeLThriftServiceDelGoodsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("delGoods", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := GoodsDeLThriftServiceDelGoodsResult{}
	var retval *GoodsDeLResponseStruct
	var err2 error
	if retval, err2 = p.handler.DelGoods(args.RequestObj); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing delGoods: "+err2.Error())
		oprot.WriteMessageBegin("delGoods", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("delGoods", thrift.REPLY, seqId); err2 != nil {
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
type GoodsDeLThriftServiceDelGoodsArgs struct {
	RequestObj *GoodsDeLRequestStruct `thrift:"requestObj,1" db:"requestObj" json:"requestObj"`
}

func NewGoodsDeLThriftServiceDelGoodsArgs() *GoodsDeLThriftServiceDelGoodsArgs {
	return &GoodsDeLThriftServiceDelGoodsArgs{}
}

var GoodsDeLThriftServiceDelGoodsArgs_RequestObj_DEFAULT *GoodsDeLRequestStruct

func (p *GoodsDeLThriftServiceDelGoodsArgs) GetRequestObj() *GoodsDeLRequestStruct {
	if !p.IsSetRequestObj() {
		return GoodsDeLThriftServiceDelGoodsArgs_RequestObj_DEFAULT
	}
	return p.RequestObj
}
func (p *GoodsDeLThriftServiceDelGoodsArgs) IsSetRequestObj() bool {
	return p.RequestObj != nil
}

func (p *GoodsDeLThriftServiceDelGoodsArgs) Read(iprot thrift.TProtocol) error {
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

func (p *GoodsDeLThriftServiceDelGoodsArgs) ReadField1(iprot thrift.TProtocol) error {
	p.RequestObj = &GoodsDeLRequestStruct{}
	if err := p.RequestObj.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.RequestObj), err)
	}
	return nil
}

func (p *GoodsDeLThriftServiceDelGoodsArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("delGoods_args"); err != nil {
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

func (p *GoodsDeLThriftServiceDelGoodsArgs) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *GoodsDeLThriftServiceDelGoodsArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GoodsDeLThriftServiceDelGoodsArgs(%+v)", *p)
}

// Attributes:
//  - Success
type GoodsDeLThriftServiceDelGoodsResult struct {
	Success *GoodsDeLResponseStruct `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewGoodsDeLThriftServiceDelGoodsResult() *GoodsDeLThriftServiceDelGoodsResult {
	return &GoodsDeLThriftServiceDelGoodsResult{}
}

var GoodsDeLThriftServiceDelGoodsResult_Success_DEFAULT *GoodsDeLResponseStruct

func (p *GoodsDeLThriftServiceDelGoodsResult) GetSuccess() *GoodsDeLResponseStruct {
	if !p.IsSetSuccess() {
		return GoodsDeLThriftServiceDelGoodsResult_Success_DEFAULT
	}
	return p.Success
}
func (p *GoodsDeLThriftServiceDelGoodsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GoodsDeLThriftServiceDelGoodsResult) Read(iprot thrift.TProtocol) error {
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

func (p *GoodsDeLThriftServiceDelGoodsResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &GoodsDeLResponseStruct{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *GoodsDeLThriftServiceDelGoodsResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("delGoods_result"); err != nil {
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

func (p *GoodsDeLThriftServiceDelGoodsResult) writeField0(oprot thrift.TProtocol) (err error) {
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

func (p *GoodsDeLThriftServiceDelGoodsResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GoodsDeLThriftServiceDelGoodsResult(%+v)", *p)
}