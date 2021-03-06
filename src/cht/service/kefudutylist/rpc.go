// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package kefudutylist

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
//  - StartTime
//  - EndTime
//  - Kefu
//  - IsExport
//  - LimitOffset
//  - LimitNum
//  - ChengHuiTongTraceLog
type KefuDutyListRequestStruct struct {
	StartTime            int32  `thrift:"start_time,1" db:"start_time" json:"start_time"`
	EndTime              int32  `thrift:"end_time,2" db:"end_time" json:"end_time"`
	Kefu                 int32  `thrift:"kefu,3" db:"kefu" json:"kefu"`
	IsExport             int32  `thrift:"is_export,4" db:"is_export" json:"is_export"`
	LimitOffset          int32  `thrift:"limitOffset,5" db:"limitOffset" json:"limitOffset"`
	LimitNum             int32  `thrift:"limitNum,6" db:"limitNum" json:"limitNum"`
	ChengHuiTongTraceLog string `thrift:"chengHuiTongTraceLog,7" db:"chengHuiTongTraceLog" json:"chengHuiTongTraceLog"`
}

// func NewKefuDutyListRequestStruct() *KefuDutyListRequestStruct {
//   return &KefuDutyListRequestStruct{}
// }

func (p *KefuDutyListRequestStruct) GetStartTime() int32 {
	return p.StartTime
}

func (p *KefuDutyListRequestStruct) GetEndTime() int32 {
	return p.EndTime
}

func (p *KefuDutyListRequestStruct) GetKefu() int32 {
	return p.Kefu
}

func (p *KefuDutyListRequestStruct) GetIsExport() int32 {
	return p.IsExport
}

func (p *KefuDutyListRequestStruct) GetLimitOffset() int32 {
	return p.LimitOffset
}

func (p *KefuDutyListRequestStruct) GetLimitNum() int32 {
	return p.LimitNum
}

func (p *KefuDutyListRequestStruct) GetChengHuiTongTraceLog() string {
	return p.ChengHuiTongTraceLog
}
func (p *KefuDutyListRequestStruct) Read(iprot thrift.TProtocol) error {
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

func (p *KefuDutyListRequestStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.StartTime = v
	}
	return nil
}

func (p *KefuDutyListRequestStruct) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.EndTime = v
	}
	return nil
}

func (p *KefuDutyListRequestStruct) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Kefu = v
	}
	return nil
}

func (p *KefuDutyListRequestStruct) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.IsExport = v
	}
	return nil
}

func (p *KefuDutyListRequestStruct) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.LimitOffset = v
	}
	return nil
}

func (p *KefuDutyListRequestStruct) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.LimitNum = v
	}
	return nil
}

func (p *KefuDutyListRequestStruct) ReadField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.ChengHuiTongTraceLog = v
	}
	return nil
}

func (p *KefuDutyListRequestStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("KefuDutyListRequestStruct"); err != nil {
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
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *KefuDutyListRequestStruct) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("start_time", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:start_time: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.StartTime)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.start_time (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:start_time: ", p), err)
	}
	return err
}

func (p *KefuDutyListRequestStruct) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("end_time", thrift.I32, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:end_time: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.EndTime)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.end_time (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:end_time: ", p), err)
	}
	return err
}

func (p *KefuDutyListRequestStruct) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("kefu", thrift.I32, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:kefu: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Kefu)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.kefu (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:kefu: ", p), err)
	}
	return err
}

func (p *KefuDutyListRequestStruct) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("is_export", thrift.I32, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:is_export: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.IsExport)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.is_export (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:is_export: ", p), err)
	}
	return err
}

func (p *KefuDutyListRequestStruct) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("limitOffset", thrift.I32, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:limitOffset: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.LimitOffset)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.limitOffset (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:limitOffset: ", p), err)
	}
	return err
}

func (p *KefuDutyListRequestStruct) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("limitNum", thrift.I32, 6); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:limitNum: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.LimitNum)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.limitNum (6) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 6:limitNum: ", p), err)
	}
	return err
}

func (p *KefuDutyListRequestStruct) writeField7(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("chengHuiTongTraceLog", thrift.STRING, 7); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:chengHuiTongTraceLog: ", p), err)
	}
	if err := oprot.WriteString(string(p.ChengHuiTongTraceLog)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.chengHuiTongTraceLog (7) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 7:chengHuiTongTraceLog: ", p), err)
	}
	return err
}

func (p *KefuDutyListRequestStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("KefuDutyListRequestStruct(%+v)", *p)
}

// Attributes:
//  - ID
//  - Customer
//  - IsRest
//  - DutyTime
//  - HolidayUser
//  - Addtime
//  - Starttime
//  - Endtime
type KefuDutyListResultStruct struct {
	ID          int32  `thrift:"id,1" db:"id" json:"id"`
	Customer    string `thrift:"customer,2" db:"customer" json:"customer"`
	IsRest      int32  `thrift:"is_rest,3" db:"is_rest" json:"is_rest"`
	DutyTime    int32  `thrift:"duty_time,4" db:"duty_time" json:"duty_time"`
	HolidayUser string `thrift:"holiday_user,5" db:"holiday_user" json:"holiday_user"`
	Addtime     int32  `thrift:"addtime,6" db:"addtime" json:"addtime"`
	Starttime   int32  `thrift:"starttime,7" db:"starttime" json:"starttime"`
	Endtime     int32  `thrift:"endtime,8" db:"endtime" json:"endtime"`
}

func NewKefuDutyListResultStruct() *KefuDutyListResultStruct {
	return &KefuDutyListResultStruct{}
}

func (p *KefuDutyListResultStruct) GetID() int32 {
	return p.ID
}

func (p *KefuDutyListResultStruct) GetCustomer() string {
	return p.Customer
}

func (p *KefuDutyListResultStruct) GetIsRest() int32 {
	return p.IsRest
}

func (p *KefuDutyListResultStruct) GetDutyTime() int32 {
	return p.DutyTime
}

func (p *KefuDutyListResultStruct) GetHolidayUser() string {
	return p.HolidayUser
}

func (p *KefuDutyListResultStruct) GetAddtime() int32 {
	return p.Addtime
}

func (p *KefuDutyListResultStruct) GetStarttime() int32 {
	return p.Starttime
}

func (p *KefuDutyListResultStruct) GetEndtime() int32 {
	return p.Endtime
}
func (p *KefuDutyListResultStruct) Read(iprot thrift.TProtocol) error {
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

func (p *KefuDutyListResultStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ID = v
	}
	return nil
}

func (p *KefuDutyListResultStruct) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Customer = v
	}
	return nil
}

func (p *KefuDutyListResultStruct) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.IsRest = v
	}
	return nil
}

func (p *KefuDutyListResultStruct) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.DutyTime = v
	}
	return nil
}

func (p *KefuDutyListResultStruct) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.HolidayUser = v
	}
	return nil
}

func (p *KefuDutyListResultStruct) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.Addtime = v
	}
	return nil
}

func (p *KefuDutyListResultStruct) ReadField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.Starttime = v
	}
	return nil
}

func (p *KefuDutyListResultStruct) ReadField8(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 8: ", err)
	} else {
		p.Endtime = v
	}
	return nil
}

func (p *KefuDutyListResultStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("KefuDutyListResultStruct"); err != nil {
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

func (p *KefuDutyListResultStruct) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *KefuDutyListResultStruct) writeField2(oprot thrift.TProtocol) (err error) {
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

func (p *KefuDutyListResultStruct) writeField3(oprot thrift.TProtocol) (err error) {
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

func (p *KefuDutyListResultStruct) writeField4(oprot thrift.TProtocol) (err error) {
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

func (p *KefuDutyListResultStruct) writeField5(oprot thrift.TProtocol) (err error) {
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

func (p *KefuDutyListResultStruct) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("addtime", thrift.I32, 6); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:addtime: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Addtime)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.addtime (6) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 6:addtime: ", p), err)
	}
	return err
}

func (p *KefuDutyListResultStruct) writeField7(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("starttime", thrift.I32, 7); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:starttime: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Starttime)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.starttime (7) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 7:starttime: ", p), err)
	}
	return err
}

func (p *KefuDutyListResultStruct) writeField8(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("endtime", thrift.I32, 8); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:endtime: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Endtime)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.endtime (8) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 8:endtime: ", p), err)
	}
	return err
}

func (p *KefuDutyListResultStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("KefuDutyListResultStruct(%+v)", *p)
}

// Attributes:
//  - Status
//  - Msg
//  - KefuDutyList
//  - TotalNum
type KefuDutyListResponseStruct struct {
	Status       int32                       `thrift:"status,1" db:"status" json:"status"`
	Msg          string                      `thrift:"msg,2" db:"msg" json:"msg"`
	KefuDutyList []*KefuDutyListResultStruct `thrift:"KefuDutyList,3" db:"KefuDutyList" json:"KefuDutyList"`
	TotalNum     int32                       `thrift:"total_num,4" db:"total_num" json:"total_num"`
}

func NewKefuDutyListResponseStruct() *KefuDutyListResponseStruct {
	return &KefuDutyListResponseStruct{}
}

func (p *KefuDutyListResponseStruct) GetStatus() int32 {
	return p.Status
}

func (p *KefuDutyListResponseStruct) GetMsg() string {
	return p.Msg
}

func (p *KefuDutyListResponseStruct) GetKefuDutyList() []*KefuDutyListResultStruct {
	return p.KefuDutyList
}

func (p *KefuDutyListResponseStruct) GetTotalNum() int32 {
	return p.TotalNum
}
func (p *KefuDutyListResponseStruct) Read(iprot thrift.TProtocol) error {
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

func (p *KefuDutyListResponseStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Status = v
	}
	return nil
}

func (p *KefuDutyListResponseStruct) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Msg = v
	}
	return nil
}

func (p *KefuDutyListResponseStruct) ReadField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*KefuDutyListResultStruct, 0, size)
	p.KefuDutyList = tSlice
	for i := 0; i < size; i++ {
		_elem0 := &KefuDutyListResultStruct{}
		if err := _elem0.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem0), err)
		}
		p.KefuDutyList = append(p.KefuDutyList, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *KefuDutyListResponseStruct) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.TotalNum = v
	}
	return nil
}

func (p *KefuDutyListResponseStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("KefuDutyListResponseStruct"); err != nil {
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
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *KefuDutyListResponseStruct) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *KefuDutyListResponseStruct) writeField2(oprot thrift.TProtocol) (err error) {
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

func (p *KefuDutyListResponseStruct) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("KefuDutyList", thrift.LIST, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:KefuDutyList: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.KefuDutyList)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.KefuDutyList {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:KefuDutyList: ", p), err)
	}
	return err
}

func (p *KefuDutyListResponseStruct) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("total_num", thrift.I32, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:total_num: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.TotalNum)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.total_num (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:total_num: ", p), err)
	}
	return err
}

func (p *KefuDutyListResponseStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("KefuDutyListResponseStruct(%+v)", *p)
}

type KefuDutyListThriftService interface {
	// Parameters:
	//  - RequestObj
	GetKefuDutyList(requestObj *KefuDutyListRequestStruct) (r *KefuDutyListResponseStruct, err error)
}

type KefuDutyListThriftServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewKefuDutyListThriftServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *KefuDutyListThriftServiceClient {
	return &KefuDutyListThriftServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewKefuDutyListThriftServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *KefuDutyListThriftServiceClient {
	return &KefuDutyListThriftServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - RequestObj
func (p *KefuDutyListThriftServiceClient) GetKefuDutyList(requestObj *KefuDutyListRequestStruct) (r *KefuDutyListResponseStruct, err error) {
	if err = p.sendGetKefuDutyList(requestObj); err != nil {
		return
	}
	return p.recvGetKefuDutyList()
}

func (p *KefuDutyListThriftServiceClient) sendGetKefuDutyList(requestObj *KefuDutyListRequestStruct) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getKefuDutyList", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := KefuDutyListThriftServiceGetKefuDutyListArgs{
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

func (p *KefuDutyListThriftServiceClient) recvGetKefuDutyList() (value *KefuDutyListResponseStruct, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "getKefuDutyList" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getKefuDutyList failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getKefuDutyList failed: out of sequence response")
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
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getKefuDutyList failed: invalid message type")
		return
	}
	result := KefuDutyListThriftServiceGetKefuDutyListResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type KefuDutyListThriftServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      KefuDutyListThriftService
}

func (p *KefuDutyListThriftServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *KefuDutyListThriftServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *KefuDutyListThriftServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewKefuDutyListThriftServiceProcessor(handler KefuDutyListThriftService) *KefuDutyListThriftServiceProcessor {

	self3 := &KefuDutyListThriftServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self3.processorMap["getKefuDutyList"] = &kefuDutyListThriftServiceProcessorGetKefuDutyList{handler: handler}
	return self3
}

func (p *KefuDutyListThriftServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
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

type kefuDutyListThriftServiceProcessorGetKefuDutyList struct {
	handler KefuDutyListThriftService
}

func (p *kefuDutyListThriftServiceProcessorGetKefuDutyList) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := KefuDutyListThriftServiceGetKefuDutyListArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getKefuDutyList", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := KefuDutyListThriftServiceGetKefuDutyListResult{}
	var retval *KefuDutyListResponseStruct
	var err2 error
	if retval, err2 = p.handler.GetKefuDutyList(args.RequestObj); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getKefuDutyList: "+err2.Error())
		oprot.WriteMessageBegin("getKefuDutyList", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getKefuDutyList", thrift.REPLY, seqId); err2 != nil {
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
type KefuDutyListThriftServiceGetKefuDutyListArgs struct {
	RequestObj *KefuDutyListRequestStruct `thrift:"requestObj,1" db:"requestObj" json:"requestObj"`
}

func NewKefuDutyListThriftServiceGetKefuDutyListArgs() *KefuDutyListThriftServiceGetKefuDutyListArgs {
	return &KefuDutyListThriftServiceGetKefuDutyListArgs{}
}

var KefuDutyListThriftServiceGetKefuDutyListArgs_RequestObj_DEFAULT *KefuDutyListRequestStruct

func (p *KefuDutyListThriftServiceGetKefuDutyListArgs) GetRequestObj() *KefuDutyListRequestStruct {
	if !p.IsSetRequestObj() {
		return KefuDutyListThriftServiceGetKefuDutyListArgs_RequestObj_DEFAULT
	}
	return p.RequestObj
}
func (p *KefuDutyListThriftServiceGetKefuDutyListArgs) IsSetRequestObj() bool {
	return p.RequestObj != nil
}

func (p *KefuDutyListThriftServiceGetKefuDutyListArgs) Read(iprot thrift.TProtocol) error {
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

func (p *KefuDutyListThriftServiceGetKefuDutyListArgs) ReadField1(iprot thrift.TProtocol) error {
	p.RequestObj = &KefuDutyListRequestStruct{}
	if err := p.RequestObj.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.RequestObj), err)
	}
	return nil
}

func (p *KefuDutyListThriftServiceGetKefuDutyListArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getKefuDutyList_args"); err != nil {
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

func (p *KefuDutyListThriftServiceGetKefuDutyListArgs) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *KefuDutyListThriftServiceGetKefuDutyListArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("KefuDutyListThriftServiceGetKefuDutyListArgs(%+v)", *p)
}

// Attributes:
//  - Success
type KefuDutyListThriftServiceGetKefuDutyListResult struct {
	Success *KefuDutyListResponseStruct `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewKefuDutyListThriftServiceGetKefuDutyListResult() *KefuDutyListThriftServiceGetKefuDutyListResult {
	return &KefuDutyListThriftServiceGetKefuDutyListResult{}
}

var KefuDutyListThriftServiceGetKefuDutyListResult_Success_DEFAULT *KefuDutyListResponseStruct

func (p *KefuDutyListThriftServiceGetKefuDutyListResult) GetSuccess() *KefuDutyListResponseStruct {
	if !p.IsSetSuccess() {
		return KefuDutyListThriftServiceGetKefuDutyListResult_Success_DEFAULT
	}
	return p.Success
}
func (p *KefuDutyListThriftServiceGetKefuDutyListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *KefuDutyListThriftServiceGetKefuDutyListResult) Read(iprot thrift.TProtocol) error {
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

func (p *KefuDutyListThriftServiceGetKefuDutyListResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &KefuDutyListResponseStruct{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *KefuDutyListThriftServiceGetKefuDutyListResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getKefuDutyList_result"); err != nil {
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

func (p *KefuDutyListThriftServiceGetKefuDutyListResult) writeField0(oprot thrift.TProtocol) (err error) {
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

func (p *KefuDutyListThriftServiceGetKefuDutyListResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("KefuDutyListThriftServiceGetKefuDutyListResult(%+v)", *p)
}
