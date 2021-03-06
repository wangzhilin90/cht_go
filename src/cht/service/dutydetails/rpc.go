// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package dutydetails

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
//  - ChengHuiTongTraceLog
type DutyDetailsRequestStruct struct {
	ChengHuiTongTraceLog string `thrift:"chengHuiTongTraceLog,1" db:"chengHuiTongTraceLog" json:"chengHuiTongTraceLog"`
}

// func NewDutyDetailsRequestStruct() *DutyDetailsRequestStruct {
//   return &DutyDetailsRequestStruct{}
// }

func (p *DutyDetailsRequestStruct) GetChengHuiTongTraceLog() string {
	return p.ChengHuiTongTraceLog
}
func (p *DutyDetailsRequestStruct) Read(iprot thrift.TProtocol) error {
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

func (p *DutyDetailsRequestStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ChengHuiTongTraceLog = v
	}
	return nil
}

func (p *DutyDetailsRequestStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("DutyDetailsRequestStruct"); err != nil {
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

func (p *DutyDetailsRequestStruct) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("chengHuiTongTraceLog", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:chengHuiTongTraceLog: ", p), err)
	}
	if err := oprot.WriteString(string(p.ChengHuiTongTraceLog)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.chengHuiTongTraceLog (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:chengHuiTongTraceLog: ", p), err)
	}
	return err
}

func (p *DutyDetailsRequestStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DutyDetailsRequestStruct(%+v)", *p)
}

// Attributes:
//  - ID
//  - Customer
//  - IsRest
//  - DutyTime
//  - HolidayUser
//  - StartTime
//  - EndTime
//  - Addtime
type DutyDetailsStruct struct {
	ID          int32  `thrift:"id,1" db:"id" json:"id"`
	Customer    string `thrift:"customer,2" db:"customer" json:"customer"`
	IsRest      int32  `thrift:"is_rest,3" db:"is_rest" json:"is_rest"`
	DutyTime    int32  `thrift:"duty_time,4" db:"duty_time" json:"duty_time"`
	HolidayUser string `thrift:"holiday_user,5" db:"holiday_user" json:"holiday_user"`
	StartTime   int32  `thrift:"start_time,6" db:"start_time" json:"start_time"`
	EndTime     int32  `thrift:"end_time,7" db:"end_time" json:"end_time"`
	Addtime     int32  `thrift:"addtime,8" db:"addtime" json:"addtime"`
}

func NewDutyDetailsStruct() *DutyDetailsStruct {
	return &DutyDetailsStruct{}
}

func (p *DutyDetailsStruct) GetID() int32 {
	return p.ID
}

func (p *DutyDetailsStruct) GetCustomer() string {
	return p.Customer
}

func (p *DutyDetailsStruct) GetIsRest() int32 {
	return p.IsRest
}

func (p *DutyDetailsStruct) GetDutyTime() int32 {
	return p.DutyTime
}

func (p *DutyDetailsStruct) GetHolidayUser() string {
	return p.HolidayUser
}

func (p *DutyDetailsStruct) GetStartTime() int32 {
	return p.StartTime
}

func (p *DutyDetailsStruct) GetEndTime() int32 {
	return p.EndTime
}

func (p *DutyDetailsStruct) GetAddtime() int32 {
	return p.Addtime
}
func (p *DutyDetailsStruct) Read(iprot thrift.TProtocol) error {
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

func (p *DutyDetailsStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ID = v
	}
	return nil
}

func (p *DutyDetailsStruct) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Customer = v
	}
	return nil
}

func (p *DutyDetailsStruct) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.IsRest = v
	}
	return nil
}

func (p *DutyDetailsStruct) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.DutyTime = v
	}
	return nil
}

func (p *DutyDetailsStruct) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.HolidayUser = v
	}
	return nil
}

func (p *DutyDetailsStruct) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.StartTime = v
	}
	return nil
}

func (p *DutyDetailsStruct) ReadField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.EndTime = v
	}
	return nil
}

func (p *DutyDetailsStruct) ReadField8(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 8: ", err)
	} else {
		p.Addtime = v
	}
	return nil
}

func (p *DutyDetailsStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("DutyDetailsStruct"); err != nil {
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
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *DutyDetailsStruct) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *DutyDetailsStruct) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("customer", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:customer: ", p), err)
	}
	if err := oprot.WriteString(string(p.Customer)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.customer (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:customer: ", p), err)
	}
	return err
}

func (p *DutyDetailsStruct) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("is_rest", thrift.I32, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:is_rest: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.IsRest)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.is_rest (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:is_rest: ", p), err)
	}
	return err
}

func (p *DutyDetailsStruct) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("duty_time", thrift.I32, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:duty_time: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.DutyTime)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.duty_time (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:duty_time: ", p), err)
	}
	return err
}

func (p *DutyDetailsStruct) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("holiday_user", thrift.STRING, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:holiday_user: ", p), err)
	}
	if err := oprot.WriteString(string(p.HolidayUser)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.holiday_user (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:holiday_user: ", p), err)
	}
	return err
}

func (p *DutyDetailsStruct) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("start_time", thrift.I32, 6); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:start_time: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.StartTime)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.start_time (6) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 6:start_time: ", p), err)
	}
	return err
}

func (p *DutyDetailsStruct) writeField7(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("end_time", thrift.I32, 7); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:end_time: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.EndTime)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.end_time (7) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 7:end_time: ", p), err)
	}
	return err
}

func (p *DutyDetailsStruct) writeField8(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("addtime", thrift.I32, 8); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:addtime: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Addtime)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.addtime (8) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 8:addtime: ", p), err)
	}
	return err
}

func (p *DutyDetailsStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DutyDetailsStruct(%+v)", *p)
}

// Attributes:
//  - Status
//  - DutyDetails
//  - Msg
type DutyDetailsResponseStruct struct {
	Status      int32              `thrift:"status,1" db:"status" json:"status"`
	DutyDetails *DutyDetailsStruct `thrift:"DutyDetails,2" db:"DutyDetails" json:"DutyDetails"`
	Msg         string             `thrift:"msg,3" db:"msg" json:"msg"`
}

func NewDutyDetailsResponseStruct() *DutyDetailsResponseStruct {
	return &DutyDetailsResponseStruct{}
}

func (p *DutyDetailsResponseStruct) GetStatus() int32 {
	return p.Status
}

var DutyDetailsResponseStruct_DutyDetails_DEFAULT *DutyDetailsStruct

func (p *DutyDetailsResponseStruct) GetDutyDetails() *DutyDetailsStruct {
	if !p.IsSetDutyDetails() {
		return DutyDetailsResponseStruct_DutyDetails_DEFAULT
	}
	return p.DutyDetails
}

func (p *DutyDetailsResponseStruct) GetMsg() string {
	return p.Msg
}
func (p *DutyDetailsResponseStruct) IsSetDutyDetails() bool {
	return p.DutyDetails != nil
}

func (p *DutyDetailsResponseStruct) Read(iprot thrift.TProtocol) error {
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

func (p *DutyDetailsResponseStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Status = v
	}
	return nil
}

func (p *DutyDetailsResponseStruct) ReadField2(iprot thrift.TProtocol) error {
	p.DutyDetails = &DutyDetailsStruct{}
	if err := p.DutyDetails.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.DutyDetails), err)
	}
	return nil
}

func (p *DutyDetailsResponseStruct) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Msg = v
	}
	return nil
}

func (p *DutyDetailsResponseStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("DutyDetailsResponseStruct"); err != nil {
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

func (p *DutyDetailsResponseStruct) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *DutyDetailsResponseStruct) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("DutyDetails", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:DutyDetails: ", p), err)
	}
	if err := p.DutyDetails.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.DutyDetails), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:DutyDetails: ", p), err)
	}
	return err
}

func (p *DutyDetailsResponseStruct) writeField3(oprot thrift.TProtocol) (err error) {
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

func (p *DutyDetailsResponseStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DutyDetailsResponseStruct(%+v)", *p)
}

type DutyDetailsThriftService interface {
	// Parameters:
	//  - RequestObj
	GetDutyDetails(requestObj *DutyDetailsRequestStruct) (r *DutyDetailsResponseStruct, err error)
}

type DutyDetailsThriftServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewDutyDetailsThriftServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *DutyDetailsThriftServiceClient {
	return &DutyDetailsThriftServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewDutyDetailsThriftServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *DutyDetailsThriftServiceClient {
	return &DutyDetailsThriftServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - RequestObj
func (p *DutyDetailsThriftServiceClient) GetDutyDetails(requestObj *DutyDetailsRequestStruct) (r *DutyDetailsResponseStruct, err error) {
	if err = p.sendGetDutyDetails(requestObj); err != nil {
		return
	}
	return p.recvGetDutyDetails()
}

func (p *DutyDetailsThriftServiceClient) sendGetDutyDetails(requestObj *DutyDetailsRequestStruct) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getDutyDetails", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := DutyDetailsThriftServiceGetDutyDetailsArgs{
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

func (p *DutyDetailsThriftServiceClient) recvGetDutyDetails() (value *DutyDetailsResponseStruct, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "getDutyDetails" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getDutyDetails failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getDutyDetails failed: out of sequence response")
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
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getDutyDetails failed: invalid message type")
		return
	}
	result := DutyDetailsThriftServiceGetDutyDetailsResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type DutyDetailsThriftServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      DutyDetailsThriftService
}

func (p *DutyDetailsThriftServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *DutyDetailsThriftServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *DutyDetailsThriftServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewDutyDetailsThriftServiceProcessor(handler DutyDetailsThriftService) *DutyDetailsThriftServiceProcessor {

	self2 := &DutyDetailsThriftServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self2.processorMap["getDutyDetails"] = &dutyDetailsThriftServiceProcessorGetDutyDetails{handler: handler}
	return self2
}

func (p *DutyDetailsThriftServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
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

type dutyDetailsThriftServiceProcessorGetDutyDetails struct {
	handler DutyDetailsThriftService
}

func (p *dutyDetailsThriftServiceProcessorGetDutyDetails) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := DutyDetailsThriftServiceGetDutyDetailsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getDutyDetails", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := DutyDetailsThriftServiceGetDutyDetailsResult{}
	var retval *DutyDetailsResponseStruct
	var err2 error
	if retval, err2 = p.handler.GetDutyDetails(args.RequestObj); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getDutyDetails: "+err2.Error())
		oprot.WriteMessageBegin("getDutyDetails", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getDutyDetails", thrift.REPLY, seqId); err2 != nil {
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
type DutyDetailsThriftServiceGetDutyDetailsArgs struct {
	RequestObj *DutyDetailsRequestStruct `thrift:"requestObj,1" db:"requestObj" json:"requestObj"`
}

func NewDutyDetailsThriftServiceGetDutyDetailsArgs() *DutyDetailsThriftServiceGetDutyDetailsArgs {
	return &DutyDetailsThriftServiceGetDutyDetailsArgs{}
}

var DutyDetailsThriftServiceGetDutyDetailsArgs_RequestObj_DEFAULT *DutyDetailsRequestStruct

func (p *DutyDetailsThriftServiceGetDutyDetailsArgs) GetRequestObj() *DutyDetailsRequestStruct {
	if !p.IsSetRequestObj() {
		return DutyDetailsThriftServiceGetDutyDetailsArgs_RequestObj_DEFAULT
	}
	return p.RequestObj
}
func (p *DutyDetailsThriftServiceGetDutyDetailsArgs) IsSetRequestObj() bool {
	return p.RequestObj != nil
}

func (p *DutyDetailsThriftServiceGetDutyDetailsArgs) Read(iprot thrift.TProtocol) error {
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

func (p *DutyDetailsThriftServiceGetDutyDetailsArgs) ReadField1(iprot thrift.TProtocol) error {
	p.RequestObj = &DutyDetailsRequestStruct{}
	if err := p.RequestObj.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.RequestObj), err)
	}
	return nil
}

func (p *DutyDetailsThriftServiceGetDutyDetailsArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getDutyDetails_args"); err != nil {
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

func (p *DutyDetailsThriftServiceGetDutyDetailsArgs) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *DutyDetailsThriftServiceGetDutyDetailsArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DutyDetailsThriftServiceGetDutyDetailsArgs(%+v)", *p)
}

// Attributes:
//  - Success
type DutyDetailsThriftServiceGetDutyDetailsResult struct {
	Success *DutyDetailsResponseStruct `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewDutyDetailsThriftServiceGetDutyDetailsResult() *DutyDetailsThriftServiceGetDutyDetailsResult {
	return &DutyDetailsThriftServiceGetDutyDetailsResult{}
}

var DutyDetailsThriftServiceGetDutyDetailsResult_Success_DEFAULT *DutyDetailsResponseStruct

func (p *DutyDetailsThriftServiceGetDutyDetailsResult) GetSuccess() *DutyDetailsResponseStruct {
	if !p.IsSetSuccess() {
		return DutyDetailsThriftServiceGetDutyDetailsResult_Success_DEFAULT
	}
	return p.Success
}
func (p *DutyDetailsThriftServiceGetDutyDetailsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DutyDetailsThriftServiceGetDutyDetailsResult) Read(iprot thrift.TProtocol) error {
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

func (p *DutyDetailsThriftServiceGetDutyDetailsResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &DutyDetailsResponseStruct{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *DutyDetailsThriftServiceGetDutyDetailsResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getDutyDetails_result"); err != nil {
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

func (p *DutyDetailsThriftServiceGetDutyDetailsResult) writeField0(oprot thrift.TProtocol) (err error) {
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

func (p *DutyDetailsThriftServiceGetDutyDetailsResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DutyDetailsThriftServiceGetDutyDetailsResult(%+v)", *p)
}
