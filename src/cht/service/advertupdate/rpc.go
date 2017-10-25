// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package advertupdate

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
//  - Type
//  - Img
//  - Adverturl
//  - Title
//  - Adduser
//  - Fid
//  - Starttime
//  - Endtime
//  - ChengHuiTongTraceLog
type AdvertUpdateRequestStruct struct {
	ID                   int32  `thrift:"id,1" db:"id" json:"id"`
	Type                 int32  `thrift:"type,2" db:"type" json:"type"`
	Img                  string `thrift:"img,3" db:"img" json:"img"`
	Adverturl            string `thrift:"adverturl,4" db:"adverturl" json:"adverturl"`
	Title                string `thrift:"title,5" db:"title" json:"title"`
	Adduser              int32  `thrift:"adduser,6" db:"adduser" json:"adduser"`
	Fid                  int32  `thrift:"fid,7" db:"fid" json:"fid"`
	Starttime            int32  `thrift:"starttime,8" db:"starttime" json:"starttime"`
	Endtime              int32  `thrift:"endtime,9" db:"endtime" json:"endtime"`
	ChengHuiTongTraceLog string `thrift:"chengHuiTongTraceLog,10" db:"chengHuiTongTraceLog" json:"chengHuiTongTraceLog"`
}

// func NewAdvertUpdateRequestStruct() *AdvertUpdateRequestStruct {
//   return &AdvertUpdateRequestStruct{}
// }

func (p *AdvertUpdateRequestStruct) GetID() int32 {
	return p.ID
}

func (p *AdvertUpdateRequestStruct) GetType() int32 {
	return p.Type
}

func (p *AdvertUpdateRequestStruct) GetImg() string {
	return p.Img
}

func (p *AdvertUpdateRequestStruct) GetAdverturl() string {
	return p.Adverturl
}

func (p *AdvertUpdateRequestStruct) GetTitle() string {
	return p.Title
}

func (p *AdvertUpdateRequestStruct) GetAdduser() int32 {
	return p.Adduser
}

func (p *AdvertUpdateRequestStruct) GetFid() int32 {
	return p.Fid
}

func (p *AdvertUpdateRequestStruct) GetStarttime() int32 {
	return p.Starttime
}

func (p *AdvertUpdateRequestStruct) GetEndtime() int32 {
	return p.Endtime
}

func (p *AdvertUpdateRequestStruct) GetChengHuiTongTraceLog() string {
	return p.ChengHuiTongTraceLog
}
func (p *AdvertUpdateRequestStruct) Read(iprot thrift.TProtocol) error {
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

func (p *AdvertUpdateRequestStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ID = v
	}
	return nil
}

func (p *AdvertUpdateRequestStruct) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Type = v
	}
	return nil
}

func (p *AdvertUpdateRequestStruct) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Img = v
	}
	return nil
}

func (p *AdvertUpdateRequestStruct) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.Adverturl = v
	}
	return nil
}

func (p *AdvertUpdateRequestStruct) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.Title = v
	}
	return nil
}

func (p *AdvertUpdateRequestStruct) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.Adduser = v
	}
	return nil
}

func (p *AdvertUpdateRequestStruct) ReadField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.Fid = v
	}
	return nil
}

func (p *AdvertUpdateRequestStruct) ReadField8(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 8: ", err)
	} else {
		p.Starttime = v
	}
	return nil
}

func (p *AdvertUpdateRequestStruct) ReadField9(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 9: ", err)
	} else {
		p.Endtime = v
	}
	return nil
}

func (p *AdvertUpdateRequestStruct) ReadField10(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 10: ", err)
	} else {
		p.ChengHuiTongTraceLog = v
	}
	return nil
}

func (p *AdvertUpdateRequestStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("AdvertUpdateRequestStruct"); err != nil {
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
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AdvertUpdateRequestStruct) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *AdvertUpdateRequestStruct) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("type", thrift.I32, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:type: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Type)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.type (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:type: ", p), err)
	}
	return err
}

func (p *AdvertUpdateRequestStruct) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("img", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:img: ", p), err)
	}
	if err := oprot.WriteString(string(p.Img)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.img (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:img: ", p), err)
	}
	return err
}

func (p *AdvertUpdateRequestStruct) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("adverturl", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:adverturl: ", p), err)
	}
	if err := oprot.WriteString(string(p.Adverturl)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.adverturl (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:adverturl: ", p), err)
	}
	return err
}

func (p *AdvertUpdateRequestStruct) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("title", thrift.STRING, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:title: ", p), err)
	}
	if err := oprot.WriteString(string(p.Title)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.title (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:title: ", p), err)
	}
	return err
}

func (p *AdvertUpdateRequestStruct) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("adduser", thrift.I32, 6); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:adduser: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Adduser)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.adduser (6) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 6:adduser: ", p), err)
	}
	return err
}

func (p *AdvertUpdateRequestStruct) writeField7(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("fid", thrift.I32, 7); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:fid: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Fid)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.fid (7) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 7:fid: ", p), err)
	}
	return err
}

func (p *AdvertUpdateRequestStruct) writeField8(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("starttime", thrift.I32, 8); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:starttime: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Starttime)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.starttime (8) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 8:starttime: ", p), err)
	}
	return err
}

func (p *AdvertUpdateRequestStruct) writeField9(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("endtime", thrift.I32, 9); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 9:endtime: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Endtime)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.endtime (9) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 9:endtime: ", p), err)
	}
	return err
}

func (p *AdvertUpdateRequestStruct) writeField10(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("chengHuiTongTraceLog", thrift.STRING, 10); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 10:chengHuiTongTraceLog: ", p), err)
	}
	if err := oprot.WriteString(string(p.ChengHuiTongTraceLog)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.chengHuiTongTraceLog (10) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 10:chengHuiTongTraceLog: ", p), err)
	}
	return err
}

func (p *AdvertUpdateRequestStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AdvertUpdateRequestStruct(%+v)", *p)
}

// Attributes:
//  - Status
//  - Msg
type AdvertUpdateResponseStruct struct {
	Status int32  `thrift:"status,1" db:"status" json:"status"`
	Msg    string `thrift:"msg,2" db:"msg" json:"msg"`
}

func NewAdvertUpdateResponseStruct() *AdvertUpdateResponseStruct {
	return &AdvertUpdateResponseStruct{}
}

func (p *AdvertUpdateResponseStruct) GetStatus() int32 {
	return p.Status
}

func (p *AdvertUpdateResponseStruct) GetMsg() string {
	return p.Msg
}
func (p *AdvertUpdateResponseStruct) Read(iprot thrift.TProtocol) error {
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

func (p *AdvertUpdateResponseStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Status = v
	}
	return nil
}

func (p *AdvertUpdateResponseStruct) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Msg = v
	}
	return nil
}

func (p *AdvertUpdateResponseStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("AdvertUpdateResponseStruct"); err != nil {
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

func (p *AdvertUpdateResponseStruct) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *AdvertUpdateResponseStruct) writeField2(oprot thrift.TProtocol) (err error) {
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

func (p *AdvertUpdateResponseStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AdvertUpdateResponseStruct(%+v)", *p)
}

type AdvertUpdateThriftService interface {
	// Parameters:
	//  - RequestObj
	UpdateAdvert(requestObj *AdvertUpdateRequestStruct) (r *AdvertUpdateResponseStruct, err error)
}

type AdvertUpdateThriftServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewAdvertUpdateThriftServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *AdvertUpdateThriftServiceClient {
	return &AdvertUpdateThriftServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewAdvertUpdateThriftServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *AdvertUpdateThriftServiceClient {
	return &AdvertUpdateThriftServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - RequestObj
func (p *AdvertUpdateThriftServiceClient) UpdateAdvert(requestObj *AdvertUpdateRequestStruct) (r *AdvertUpdateResponseStruct, err error) {
	if err = p.sendUpdateAdvert(requestObj); err != nil {
		return
	}
	return p.recvUpdateAdvert()
}

func (p *AdvertUpdateThriftServiceClient) sendUpdateAdvert(requestObj *AdvertUpdateRequestStruct) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("updateAdvert", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := AdvertUpdateThriftServiceUpdateAdvertArgs{
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

func (p *AdvertUpdateThriftServiceClient) recvUpdateAdvert() (value *AdvertUpdateResponseStruct, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "updateAdvert" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "updateAdvert failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "updateAdvert failed: out of sequence response")
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
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "updateAdvert failed: invalid message type")
		return
	}
	result := AdvertUpdateThriftServiceUpdateAdvertResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type AdvertUpdateThriftServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      AdvertUpdateThriftService
}

func (p *AdvertUpdateThriftServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *AdvertUpdateThriftServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *AdvertUpdateThriftServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewAdvertUpdateThriftServiceProcessor(handler AdvertUpdateThriftService) *AdvertUpdateThriftServiceProcessor {

	self2 := &AdvertUpdateThriftServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self2.processorMap["updateAdvert"] = &advertUpdateThriftServiceProcessorUpdateAdvert{handler: handler}
	return self2
}

func (p *AdvertUpdateThriftServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
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

type advertUpdateThriftServiceProcessorUpdateAdvert struct {
	handler AdvertUpdateThriftService
}

func (p *advertUpdateThriftServiceProcessorUpdateAdvert) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AdvertUpdateThriftServiceUpdateAdvertArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("updateAdvert", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := AdvertUpdateThriftServiceUpdateAdvertResult{}
	var retval *AdvertUpdateResponseStruct
	var err2 error
	if retval, err2 = p.handler.UpdateAdvert(args.RequestObj); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing updateAdvert: "+err2.Error())
		oprot.WriteMessageBegin("updateAdvert", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("updateAdvert", thrift.REPLY, seqId); err2 != nil {
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
type AdvertUpdateThriftServiceUpdateAdvertArgs struct {
	RequestObj *AdvertUpdateRequestStruct `thrift:"requestObj,1" db:"requestObj" json:"requestObj"`
}

func NewAdvertUpdateThriftServiceUpdateAdvertArgs() *AdvertUpdateThriftServiceUpdateAdvertArgs {
	return &AdvertUpdateThriftServiceUpdateAdvertArgs{}
}

var AdvertUpdateThriftServiceUpdateAdvertArgs_RequestObj_DEFAULT *AdvertUpdateRequestStruct

func (p *AdvertUpdateThriftServiceUpdateAdvertArgs) GetRequestObj() *AdvertUpdateRequestStruct {
	if !p.IsSetRequestObj() {
		return AdvertUpdateThriftServiceUpdateAdvertArgs_RequestObj_DEFAULT
	}
	return p.RequestObj
}
func (p *AdvertUpdateThriftServiceUpdateAdvertArgs) IsSetRequestObj() bool {
	return p.RequestObj != nil
}

func (p *AdvertUpdateThriftServiceUpdateAdvertArgs) Read(iprot thrift.TProtocol) error {
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

func (p *AdvertUpdateThriftServiceUpdateAdvertArgs) ReadField1(iprot thrift.TProtocol) error {
	p.RequestObj = &AdvertUpdateRequestStruct{}
	if err := p.RequestObj.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.RequestObj), err)
	}
	return nil
}

func (p *AdvertUpdateThriftServiceUpdateAdvertArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("updateAdvert_args"); err != nil {
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

func (p *AdvertUpdateThriftServiceUpdateAdvertArgs) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *AdvertUpdateThriftServiceUpdateAdvertArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AdvertUpdateThriftServiceUpdateAdvertArgs(%+v)", *p)
}

// Attributes:
//  - Success
type AdvertUpdateThriftServiceUpdateAdvertResult struct {
	Success *AdvertUpdateResponseStruct `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewAdvertUpdateThriftServiceUpdateAdvertResult() *AdvertUpdateThriftServiceUpdateAdvertResult {
	return &AdvertUpdateThriftServiceUpdateAdvertResult{}
}

var AdvertUpdateThriftServiceUpdateAdvertResult_Success_DEFAULT *AdvertUpdateResponseStruct

func (p *AdvertUpdateThriftServiceUpdateAdvertResult) GetSuccess() *AdvertUpdateResponseStruct {
	if !p.IsSetSuccess() {
		return AdvertUpdateThriftServiceUpdateAdvertResult_Success_DEFAULT
	}
	return p.Success
}
func (p *AdvertUpdateThriftServiceUpdateAdvertResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AdvertUpdateThriftServiceUpdateAdvertResult) Read(iprot thrift.TProtocol) error {
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

func (p *AdvertUpdateThriftServiceUpdateAdvertResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &AdvertUpdateResponseStruct{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *AdvertUpdateThriftServiceUpdateAdvertResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("updateAdvert_result"); err != nil {
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

func (p *AdvertUpdateThriftServiceUpdateAdvertResult) writeField0(oprot thrift.TProtocol) (err error) {
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

func (p *AdvertUpdateThriftServiceUpdateAdvertResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AdvertUpdateThriftServiceUpdateAdvertResult(%+v)", *p)
}
