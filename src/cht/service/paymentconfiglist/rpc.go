// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package paymentconfiglist

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
type PaymentConfigListRequestStruct struct {
	ChengHuiTongTraceLog string `thrift:"chengHuiTongTraceLog,1" db:"chengHuiTongTraceLog" json:"chengHuiTongTraceLog"`
}

// func NewPaymentConfigListRequestStruct() *PaymentConfigListRequestStruct {
//   return &PaymentConfigListRequestStruct{}
// }

func (p *PaymentConfigListRequestStruct) GetChengHuiTongTraceLog() string {
	return p.ChengHuiTongTraceLog
}
func (p *PaymentConfigListRequestStruct) Read(iprot thrift.TProtocol) error {
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

func (p *PaymentConfigListRequestStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ChengHuiTongTraceLog = v
	}
	return nil
}

func (p *PaymentConfigListRequestStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("PaymentConfigListRequestStruct"); err != nil {
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

func (p *PaymentConfigListRequestStruct) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *PaymentConfigListRequestStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PaymentConfigListRequestStruct(%+v)", *p)
}

// Attributes:
//  - ID
//  - Type
//  - Nid
//  - Name
//  - Logo
//  - Config
//  - Fee
//  - Status
//  - Remark
//  - Sort
type PaymentConfigDetailsStruct struct {
	ID     int32  `thrift:"id,1" db:"id" json:"id"`
	Type   int32  `thrift:"type,2" db:"type" json:"type"`
	Nid    string `thrift:"nid,3" db:"nid" json:"nid"`
	Name   string `thrift:"name,4" db:"name" json:"name"`
	Logo   string `thrift:"logo,5" db:"logo" json:"logo"`
	Config string `thrift:"config,6" db:"config" json:"config"`
	Fee    string `thrift:"fee,7" db:"fee" json:"fee"`
	Status int32  `thrift:"status,8" db:"status" json:"status"`
	Remark string `thrift:"remark,9" db:"remark" json:"remark"`
	Sort   int32  `thrift:"sort,10" db:"sort" json:"sort"`
}

func NewPaymentConfigDetailsStruct() *PaymentConfigDetailsStruct {
	return &PaymentConfigDetailsStruct{}
}

func (p *PaymentConfigDetailsStruct) GetID() int32 {
	return p.ID
}

func (p *PaymentConfigDetailsStruct) GetType() int32 {
	return p.Type
}

func (p *PaymentConfigDetailsStruct) GetNid() string {
	return p.Nid
}

func (p *PaymentConfigDetailsStruct) GetName() string {
	return p.Name
}

func (p *PaymentConfigDetailsStruct) GetLogo() string {
	return p.Logo
}

func (p *PaymentConfigDetailsStruct) GetConfig() string {
	return p.Config
}

func (p *PaymentConfigDetailsStruct) GetFee() string {
	return p.Fee
}

func (p *PaymentConfigDetailsStruct) GetStatus() int32 {
	return p.Status
}

func (p *PaymentConfigDetailsStruct) GetRemark() string {
	return p.Remark
}

func (p *PaymentConfigDetailsStruct) GetSort() int32 {
	return p.Sort
}
func (p *PaymentConfigDetailsStruct) Read(iprot thrift.TProtocol) error {
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

func (p *PaymentConfigDetailsStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ID = v
	}
	return nil
}

func (p *PaymentConfigDetailsStruct) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Type = v
	}
	return nil
}

func (p *PaymentConfigDetailsStruct) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Nid = v
	}
	return nil
}

func (p *PaymentConfigDetailsStruct) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.Name = v
	}
	return nil
}

func (p *PaymentConfigDetailsStruct) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.Logo = v
	}
	return nil
}

func (p *PaymentConfigDetailsStruct) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.Config = v
	}
	return nil
}

func (p *PaymentConfigDetailsStruct) ReadField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.Fee = v
	}
	return nil
}

func (p *PaymentConfigDetailsStruct) ReadField8(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 8: ", err)
	} else {
		p.Status = v
	}
	return nil
}

func (p *PaymentConfigDetailsStruct) ReadField9(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 9: ", err)
	} else {
		p.Remark = v
	}
	return nil
}

func (p *PaymentConfigDetailsStruct) ReadField10(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 10: ", err)
	} else {
		p.Sort = v
	}
	return nil
}

func (p *PaymentConfigDetailsStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("PaymentConfigDetailsStruct"); err != nil {
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

func (p *PaymentConfigDetailsStruct) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *PaymentConfigDetailsStruct) writeField2(oprot thrift.TProtocol) (err error) {
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

func (p *PaymentConfigDetailsStruct) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("nid", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:nid: ", p), err)
	}
	if err := oprot.WriteString(string(p.Nid)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.nid (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:nid: ", p), err)
	}
	return err
}

func (p *PaymentConfigDetailsStruct) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("name", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:name: ", p), err)
	}
	if err := oprot.WriteString(string(p.Name)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.name (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:name: ", p), err)
	}
	return err
}

func (p *PaymentConfigDetailsStruct) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("logo", thrift.STRING, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:logo: ", p), err)
	}
	if err := oprot.WriteString(string(p.Logo)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.logo (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:logo: ", p), err)
	}
	return err
}

func (p *PaymentConfigDetailsStruct) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("config", thrift.STRING, 6); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:config: ", p), err)
	}
	if err := oprot.WriteString(string(p.Config)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.config (6) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 6:config: ", p), err)
	}
	return err
}

func (p *PaymentConfigDetailsStruct) writeField7(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("fee", thrift.STRING, 7); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:fee: ", p), err)
	}
	if err := oprot.WriteString(string(p.Fee)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.fee (7) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 7:fee: ", p), err)
	}
	return err
}

func (p *PaymentConfigDetailsStruct) writeField8(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("status", thrift.I32, 8); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:status: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Status)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.status (8) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 8:status: ", p), err)
	}
	return err
}

func (p *PaymentConfigDetailsStruct) writeField9(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("remark", thrift.STRING, 9); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 9:remark: ", p), err)
	}
	if err := oprot.WriteString(string(p.Remark)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.remark (9) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 9:remark: ", p), err)
	}
	return err
}

func (p *PaymentConfigDetailsStruct) writeField10(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("sort", thrift.I32, 10); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 10:sort: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Sort)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.sort (10) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 10:sort: ", p), err)
	}
	return err
}

func (p *PaymentConfigDetailsStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PaymentConfigDetailsStruct(%+v)", *p)
}

// Attributes:
//  - Status
//  - Msg
//  - PaymentConfigList
type PaymentConfigListResponseStruct struct {
	Status            int32                         `thrift:"status,1" db:"status" json:"status"`
	Msg               string                        `thrift:"msg,2" db:"msg" json:"msg"`
	PaymentConfigList []*PaymentConfigDetailsStruct `thrift:"PaymentConfigList,3" db:"PaymentConfigList" json:"PaymentConfigList"`
}

func NewPaymentConfigListResponseStruct() *PaymentConfigListResponseStruct {
	return &PaymentConfigListResponseStruct{}
}

func (p *PaymentConfigListResponseStruct) GetStatus() int32 {
	return p.Status
}

func (p *PaymentConfigListResponseStruct) GetMsg() string {
	return p.Msg
}

func (p *PaymentConfigListResponseStruct) GetPaymentConfigList() []*PaymentConfigDetailsStruct {
	return p.PaymentConfigList
}
func (p *PaymentConfigListResponseStruct) Read(iprot thrift.TProtocol) error {
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

func (p *PaymentConfigListResponseStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Status = v
	}
	return nil
}

func (p *PaymentConfigListResponseStruct) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Msg = v
	}
	return nil
}

func (p *PaymentConfigListResponseStruct) ReadField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*PaymentConfigDetailsStruct, 0, size)
	p.PaymentConfigList = tSlice
	for i := 0; i < size; i++ {
		_elem0 := &PaymentConfigDetailsStruct{}
		if err := _elem0.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem0), err)
		}
		p.PaymentConfigList = append(p.PaymentConfigList, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *PaymentConfigListResponseStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("PaymentConfigListResponseStruct"); err != nil {
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

func (p *PaymentConfigListResponseStruct) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *PaymentConfigListResponseStruct) writeField2(oprot thrift.TProtocol) (err error) {
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

func (p *PaymentConfigListResponseStruct) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("PaymentConfigList", thrift.LIST, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:PaymentConfigList: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.PaymentConfigList)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.PaymentConfigList {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:PaymentConfigList: ", p), err)
	}
	return err
}

func (p *PaymentConfigListResponseStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PaymentConfigListResponseStruct(%+v)", *p)
}

type PaymentConfigListThriftService interface {
	// Parameters:
	//  - RequestObj
	GetPaymentConfigList(requestObj *PaymentConfigListRequestStruct) (r *PaymentConfigListResponseStruct, err error)
}

type PaymentConfigListThriftServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewPaymentConfigListThriftServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *PaymentConfigListThriftServiceClient {
	return &PaymentConfigListThriftServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewPaymentConfigListThriftServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *PaymentConfigListThriftServiceClient {
	return &PaymentConfigListThriftServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - RequestObj
func (p *PaymentConfigListThriftServiceClient) GetPaymentConfigList(requestObj *PaymentConfigListRequestStruct) (r *PaymentConfigListResponseStruct, err error) {
	if err = p.sendGetPaymentConfigList(requestObj); err != nil {
		return
	}
	return p.recvGetPaymentConfigList()
}

func (p *PaymentConfigListThriftServiceClient) sendGetPaymentConfigList(requestObj *PaymentConfigListRequestStruct) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getPaymentConfigList", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := PaymentConfigListThriftServiceGetPaymentConfigListArgs{
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

func (p *PaymentConfigListThriftServiceClient) recvGetPaymentConfigList() (value *PaymentConfigListResponseStruct, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "getPaymentConfigList" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getPaymentConfigList failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getPaymentConfigList failed: out of sequence response")
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
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getPaymentConfigList failed: invalid message type")
		return
	}
	result := PaymentConfigListThriftServiceGetPaymentConfigListResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type PaymentConfigListThriftServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      PaymentConfigListThriftService
}

func (p *PaymentConfigListThriftServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *PaymentConfigListThriftServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *PaymentConfigListThriftServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewPaymentConfigListThriftServiceProcessor(handler PaymentConfigListThriftService) *PaymentConfigListThriftServiceProcessor {

	self3 := &PaymentConfigListThriftServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self3.processorMap["getPaymentConfigList"] = &paymentConfigListThriftServiceProcessorGetPaymentConfigList{handler: handler}
	return self3
}

func (p *PaymentConfigListThriftServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
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

type paymentConfigListThriftServiceProcessorGetPaymentConfigList struct {
	handler PaymentConfigListThriftService
}

func (p *paymentConfigListThriftServiceProcessorGetPaymentConfigList) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := PaymentConfigListThriftServiceGetPaymentConfigListArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getPaymentConfigList", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := PaymentConfigListThriftServiceGetPaymentConfigListResult{}
	var retval *PaymentConfigListResponseStruct
	var err2 error
	if retval, err2 = p.handler.GetPaymentConfigList(args.RequestObj); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getPaymentConfigList: "+err2.Error())
		oprot.WriteMessageBegin("getPaymentConfigList", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getPaymentConfigList", thrift.REPLY, seqId); err2 != nil {
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
type PaymentConfigListThriftServiceGetPaymentConfigListArgs struct {
	RequestObj *PaymentConfigListRequestStruct `thrift:"requestObj,1" db:"requestObj" json:"requestObj"`
}

func NewPaymentConfigListThriftServiceGetPaymentConfigListArgs() *PaymentConfigListThriftServiceGetPaymentConfigListArgs {
	return &PaymentConfigListThriftServiceGetPaymentConfigListArgs{}
}

var PaymentConfigListThriftServiceGetPaymentConfigListArgs_RequestObj_DEFAULT *PaymentConfigListRequestStruct

func (p *PaymentConfigListThriftServiceGetPaymentConfigListArgs) GetRequestObj() *PaymentConfigListRequestStruct {
	if !p.IsSetRequestObj() {
		return PaymentConfigListThriftServiceGetPaymentConfigListArgs_RequestObj_DEFAULT
	}
	return p.RequestObj
}
func (p *PaymentConfigListThriftServiceGetPaymentConfigListArgs) IsSetRequestObj() bool {
	return p.RequestObj != nil
}

func (p *PaymentConfigListThriftServiceGetPaymentConfigListArgs) Read(iprot thrift.TProtocol) error {
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

func (p *PaymentConfigListThriftServiceGetPaymentConfigListArgs) ReadField1(iprot thrift.TProtocol) error {
	p.RequestObj = &PaymentConfigListRequestStruct{}
	if err := p.RequestObj.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.RequestObj), err)
	}
	return nil
}

func (p *PaymentConfigListThriftServiceGetPaymentConfigListArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getPaymentConfigList_args"); err != nil {
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

func (p *PaymentConfigListThriftServiceGetPaymentConfigListArgs) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *PaymentConfigListThriftServiceGetPaymentConfigListArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PaymentConfigListThriftServiceGetPaymentConfigListArgs(%+v)", *p)
}

// Attributes:
//  - Success
type PaymentConfigListThriftServiceGetPaymentConfigListResult struct {
	Success *PaymentConfigListResponseStruct `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewPaymentConfigListThriftServiceGetPaymentConfigListResult() *PaymentConfigListThriftServiceGetPaymentConfigListResult {
	return &PaymentConfigListThriftServiceGetPaymentConfigListResult{}
}

var PaymentConfigListThriftServiceGetPaymentConfigListResult_Success_DEFAULT *PaymentConfigListResponseStruct

func (p *PaymentConfigListThriftServiceGetPaymentConfigListResult) GetSuccess() *PaymentConfigListResponseStruct {
	if !p.IsSetSuccess() {
		return PaymentConfigListThriftServiceGetPaymentConfigListResult_Success_DEFAULT
	}
	return p.Success
}
func (p *PaymentConfigListThriftServiceGetPaymentConfigListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PaymentConfigListThriftServiceGetPaymentConfigListResult) Read(iprot thrift.TProtocol) error {
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

func (p *PaymentConfigListThriftServiceGetPaymentConfigListResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &PaymentConfigListResponseStruct{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PaymentConfigListThriftServiceGetPaymentConfigListResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getPaymentConfigList_result"); err != nil {
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

func (p *PaymentConfigListThriftServiceGetPaymentConfigListResult) writeField0(oprot thrift.TProtocol) (err error) {
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

func (p *PaymentConfigListThriftServiceGetPaymentConfigListResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PaymentConfigListThriftServiceGetPaymentConfigListResult(%+v)", *p)
}
