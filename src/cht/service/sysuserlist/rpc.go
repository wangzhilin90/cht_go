// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package sysuserlist

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
type SysUserListRequestStruct struct {
	ChengHuiTongTraceLog string `thrift:"chengHuiTongTraceLog,1" db:"chengHuiTongTraceLog" json:"chengHuiTongTraceLog"`
}

// func NewSysUserListRequestStruct() *SysUserListRequestStruct {
//   return &SysUserListRequestStruct{}
// }

func (p *SysUserListRequestStruct) GetChengHuiTongTraceLog() string {
	return p.ChengHuiTongTraceLog
}
func (p *SysUserListRequestStruct) Read(iprot thrift.TProtocol) error {
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

func (p *SysUserListRequestStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ChengHuiTongTraceLog = v
	}
	return nil
}

func (p *SysUserListRequestStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("SysUserListRequestStruct"); err != nil {
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

func (p *SysUserListRequestStruct) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *SysUserListRequestStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SysUserListRequestStruct(%+v)", *p)
}

// Attributes:
//  - ID
//  - RoleID
//  - Account
//  - Realname
//  - Password
//  - Mobile
//  - Qq
//  - Lastloginip
//  - Lastlogintime
//  - CreateTime
//  - Status
//  - Views
//  - CustomerType
type SysUserDetailsStruct struct {
	ID            int32  `thrift:"id,1" db:"id" json:"id"`
	RoleID        int32  `thrift:"role_id,2" db:"role_id" json:"role_id"`
	Account       string `thrift:"account,3" db:"account" json:"account"`
	Realname      string `thrift:"realname,4" db:"realname" json:"realname"`
	Password      string `thrift:"password,5" db:"password" json:"password"`
	Mobile        string `thrift:"mobile,6" db:"mobile" json:"mobile"`
	Qq            string `thrift:"qq,7" db:"qq" json:"qq"`
	Lastloginip   string `thrift:"lastloginip,8" db:"lastloginip" json:"lastloginip"`
	Lastlogintime int32  `thrift:"lastlogintime,9" db:"lastlogintime" json:"lastlogintime"`
	CreateTime    int32  `thrift:"create_time,10" db:"create_time" json:"create_time"`
	Status        int32  `thrift:"status,11" db:"status" json:"status"`
	Views         int32  `thrift:"views,12" db:"views" json:"views"`
	CustomerType  int32  `thrift:"customer_type,13" db:"customer_type" json:"customer_type"`
}

func NewSysUserDetailsStruct() *SysUserDetailsStruct {
	return &SysUserDetailsStruct{}
}

func (p *SysUserDetailsStruct) GetID() int32 {
	return p.ID
}

func (p *SysUserDetailsStruct) GetRoleID() int32 {
	return p.RoleID
}

func (p *SysUserDetailsStruct) GetAccount() string {
	return p.Account
}

func (p *SysUserDetailsStruct) GetRealname() string {
	return p.Realname
}

func (p *SysUserDetailsStruct) GetPassword() string {
	return p.Password
}

func (p *SysUserDetailsStruct) GetMobile() string {
	return p.Mobile
}

func (p *SysUserDetailsStruct) GetQq() string {
	return p.Qq
}

func (p *SysUserDetailsStruct) GetLastloginip() string {
	return p.Lastloginip
}

func (p *SysUserDetailsStruct) GetLastlogintime() int32 {
	return p.Lastlogintime
}

func (p *SysUserDetailsStruct) GetCreateTime() int32 {
	return p.CreateTime
}

func (p *SysUserDetailsStruct) GetStatus() int32 {
	return p.Status
}

func (p *SysUserDetailsStruct) GetViews() int32 {
	return p.Views
}

func (p *SysUserDetailsStruct) GetCustomerType() int32 {
	return p.CustomerType
}
func (p *SysUserDetailsStruct) Read(iprot thrift.TProtocol) error {
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
		case 12:
			if err := p.ReadField12(iprot); err != nil {
				return err
			}
		case 13:
			if err := p.ReadField13(iprot); err != nil {
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

func (p *SysUserDetailsStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ID = v
	}
	return nil
}

func (p *SysUserDetailsStruct) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.RoleID = v
	}
	return nil
}

func (p *SysUserDetailsStruct) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Account = v
	}
	return nil
}

func (p *SysUserDetailsStruct) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.Realname = v
	}
	return nil
}

func (p *SysUserDetailsStruct) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.Password = v
	}
	return nil
}

func (p *SysUserDetailsStruct) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.Mobile = v
	}
	return nil
}

func (p *SysUserDetailsStruct) ReadField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.Qq = v
	}
	return nil
}

func (p *SysUserDetailsStruct) ReadField8(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 8: ", err)
	} else {
		p.Lastloginip = v
	}
	return nil
}

func (p *SysUserDetailsStruct) ReadField9(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 9: ", err)
	} else {
		p.Lastlogintime = v
	}
	return nil
}

func (p *SysUserDetailsStruct) ReadField10(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 10: ", err)
	} else {
		p.CreateTime = v
	}
	return nil
}

func (p *SysUserDetailsStruct) ReadField11(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 11: ", err)
	} else {
		p.Status = v
	}
	return nil
}

func (p *SysUserDetailsStruct) ReadField12(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 12: ", err)
	} else {
		p.Views = v
	}
	return nil
}

func (p *SysUserDetailsStruct) ReadField13(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 13: ", err)
	} else {
		p.CustomerType = v
	}
	return nil
}

func (p *SysUserDetailsStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("SysUserDetailsStruct"); err != nil {
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
		if err := p.writeField12(oprot); err != nil {
			return err
		}
		if err := p.writeField13(oprot); err != nil {
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

func (p *SysUserDetailsStruct) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *SysUserDetailsStruct) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("role_id", thrift.I32, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:role_id: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.RoleID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.role_id (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:role_id: ", p), err)
	}
	return err
}

func (p *SysUserDetailsStruct) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("account", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:account: ", p), err)
	}
	if err := oprot.WriteString(string(p.Account)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.account (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:account: ", p), err)
	}
	return err
}

func (p *SysUserDetailsStruct) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("realname", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:realname: ", p), err)
	}
	if err := oprot.WriteString(string(p.Realname)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.realname (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:realname: ", p), err)
	}
	return err
}

func (p *SysUserDetailsStruct) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("password", thrift.STRING, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:password: ", p), err)
	}
	if err := oprot.WriteString(string(p.Password)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.password (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:password: ", p), err)
	}
	return err
}

func (p *SysUserDetailsStruct) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("mobile", thrift.STRING, 6); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:mobile: ", p), err)
	}
	if err := oprot.WriteString(string(p.Mobile)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.mobile (6) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 6:mobile: ", p), err)
	}
	return err
}

func (p *SysUserDetailsStruct) writeField7(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("qq", thrift.STRING, 7); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:qq: ", p), err)
	}
	if err := oprot.WriteString(string(p.Qq)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.qq (7) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 7:qq: ", p), err)
	}
	return err
}

func (p *SysUserDetailsStruct) writeField8(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("lastloginip", thrift.STRING, 8); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:lastloginip: ", p), err)
	}
	if err := oprot.WriteString(string(p.Lastloginip)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.lastloginip (8) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 8:lastloginip: ", p), err)
	}
	return err
}

func (p *SysUserDetailsStruct) writeField9(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("lastlogintime", thrift.I32, 9); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 9:lastlogintime: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Lastlogintime)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.lastlogintime (9) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 9:lastlogintime: ", p), err)
	}
	return err
}

func (p *SysUserDetailsStruct) writeField10(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("create_time", thrift.I32, 10); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 10:create_time: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.CreateTime)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.create_time (10) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 10:create_time: ", p), err)
	}
	return err
}

func (p *SysUserDetailsStruct) writeField11(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("status", thrift.I32, 11); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 11:status: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Status)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.status (11) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 11:status: ", p), err)
	}
	return err
}

func (p *SysUserDetailsStruct) writeField12(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("views", thrift.I32, 12); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 12:views: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Views)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.views (12) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 12:views: ", p), err)
	}
	return err
}

func (p *SysUserDetailsStruct) writeField13(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("customer_type", thrift.I32, 13); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 13:customer_type: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.CustomerType)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.customer_type (13) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 13:customer_type: ", p), err)
	}
	return err
}

func (p *SysUserDetailsStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SysUserDetailsStruct(%+v)", *p)
}

// Attributes:
//  - Status
//  - SysUserList
//  - Msg
type SysUserListResponseStruct struct {
	Status      int32                   `thrift:"status,1" db:"status" json:"status"`
	SysUserList []*SysUserDetailsStruct `thrift:"SysUserList,2" db:"SysUserList" json:"SysUserList"`
	Msg         string                  `thrift:"msg,3" db:"msg" json:"msg"`
}

func NewSysUserListResponseStruct() *SysUserListResponseStruct {
	return &SysUserListResponseStruct{}
}

func (p *SysUserListResponseStruct) GetStatus() int32 {
	return p.Status
}

func (p *SysUserListResponseStruct) GetSysUserList() []*SysUserDetailsStruct {
	return p.SysUserList
}

func (p *SysUserListResponseStruct) GetMsg() string {
	return p.Msg
}
func (p *SysUserListResponseStruct) Read(iprot thrift.TProtocol) error {
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

func (p *SysUserListResponseStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Status = v
	}
	return nil
}

func (p *SysUserListResponseStruct) ReadField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*SysUserDetailsStruct, 0, size)
	p.SysUserList = tSlice
	for i := 0; i < size; i++ {
		_elem0 := &SysUserDetailsStruct{}
		if err := _elem0.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem0), err)
		}
		p.SysUserList = append(p.SysUserList, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *SysUserListResponseStruct) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Msg = v
	}
	return nil
}

func (p *SysUserListResponseStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("SysUserListResponseStruct"); err != nil {
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

func (p *SysUserListResponseStruct) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *SysUserListResponseStruct) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("SysUserList", thrift.LIST, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:SysUserList: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.SysUserList)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.SysUserList {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:SysUserList: ", p), err)
	}
	return err
}

func (p *SysUserListResponseStruct) writeField3(oprot thrift.TProtocol) (err error) {
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

func (p *SysUserListResponseStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SysUserListResponseStruct(%+v)", *p)
}

type SysUserListThriftService interface {
	// Parameters:
	//  - RequestObj
	GetSysUserList(requestObj *SysUserListRequestStruct) (r *SysUserListResponseStruct, err error)
}

type SysUserListThriftServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewSysUserListThriftServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *SysUserListThriftServiceClient {
	return &SysUserListThriftServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewSysUserListThriftServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *SysUserListThriftServiceClient {
	return &SysUserListThriftServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - RequestObj
func (p *SysUserListThriftServiceClient) GetSysUserList(requestObj *SysUserListRequestStruct) (r *SysUserListResponseStruct, err error) {
	if err = p.sendGetSysUserList(requestObj); err != nil {
		return
	}
	return p.recvGetSysUserList()
}

func (p *SysUserListThriftServiceClient) sendGetSysUserList(requestObj *SysUserListRequestStruct) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getSysUserList", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := SysUserListThriftServiceGetSysUserListArgs{
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

func (p *SysUserListThriftServiceClient) recvGetSysUserList() (value *SysUserListResponseStruct, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "getSysUserList" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getSysUserList failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getSysUserList failed: out of sequence response")
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
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getSysUserList failed: invalid message type")
		return
	}
	result := SysUserListThriftServiceGetSysUserListResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type SysUserListThriftServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      SysUserListThriftService
}

func (p *SysUserListThriftServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *SysUserListThriftServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *SysUserListThriftServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewSysUserListThriftServiceProcessor(handler SysUserListThriftService) *SysUserListThriftServiceProcessor {

	self3 := &SysUserListThriftServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self3.processorMap["getSysUserList"] = &sysUserListThriftServiceProcessorGetSysUserList{handler: handler}
	return self3
}

func (p *SysUserListThriftServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
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

type sysUserListThriftServiceProcessorGetSysUserList struct {
	handler SysUserListThriftService
}

func (p *sysUserListThriftServiceProcessorGetSysUserList) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := SysUserListThriftServiceGetSysUserListArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getSysUserList", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := SysUserListThriftServiceGetSysUserListResult{}
	var retval *SysUserListResponseStruct
	var err2 error
	if retval, err2 = p.handler.GetSysUserList(args.RequestObj); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getSysUserList: "+err2.Error())
		oprot.WriteMessageBegin("getSysUserList", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getSysUserList", thrift.REPLY, seqId); err2 != nil {
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
type SysUserListThriftServiceGetSysUserListArgs struct {
	RequestObj *SysUserListRequestStruct `thrift:"requestObj,1" db:"requestObj" json:"requestObj"`
}

func NewSysUserListThriftServiceGetSysUserListArgs() *SysUserListThriftServiceGetSysUserListArgs {
	return &SysUserListThriftServiceGetSysUserListArgs{}
}

var SysUserListThriftServiceGetSysUserListArgs_RequestObj_DEFAULT *SysUserListRequestStruct

func (p *SysUserListThriftServiceGetSysUserListArgs) GetRequestObj() *SysUserListRequestStruct {
	if !p.IsSetRequestObj() {
		return SysUserListThriftServiceGetSysUserListArgs_RequestObj_DEFAULT
	}
	return p.RequestObj
}
func (p *SysUserListThriftServiceGetSysUserListArgs) IsSetRequestObj() bool {
	return p.RequestObj != nil
}

func (p *SysUserListThriftServiceGetSysUserListArgs) Read(iprot thrift.TProtocol) error {
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

func (p *SysUserListThriftServiceGetSysUserListArgs) ReadField1(iprot thrift.TProtocol) error {
	p.RequestObj = &SysUserListRequestStruct{}
	if err := p.RequestObj.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.RequestObj), err)
	}
	return nil
}

func (p *SysUserListThriftServiceGetSysUserListArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getSysUserList_args"); err != nil {
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

func (p *SysUserListThriftServiceGetSysUserListArgs) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *SysUserListThriftServiceGetSysUserListArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SysUserListThriftServiceGetSysUserListArgs(%+v)", *p)
}

// Attributes:
//  - Success
type SysUserListThriftServiceGetSysUserListResult struct {
	Success *SysUserListResponseStruct `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewSysUserListThriftServiceGetSysUserListResult() *SysUserListThriftServiceGetSysUserListResult {
	return &SysUserListThriftServiceGetSysUserListResult{}
}

var SysUserListThriftServiceGetSysUserListResult_Success_DEFAULT *SysUserListResponseStruct

func (p *SysUserListThriftServiceGetSysUserListResult) GetSuccess() *SysUserListResponseStruct {
	if !p.IsSetSuccess() {
		return SysUserListThriftServiceGetSysUserListResult_Success_DEFAULT
	}
	return p.Success
}
func (p *SysUserListThriftServiceGetSysUserListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SysUserListThriftServiceGetSysUserListResult) Read(iprot thrift.TProtocol) error {
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

func (p *SysUserListThriftServiceGetSysUserListResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &SysUserListResponseStruct{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *SysUserListThriftServiceGetSysUserListResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getSysUserList_result"); err != nil {
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

func (p *SysUserListThriftServiceGetSysUserListResult) writeField0(oprot thrift.TProtocol) (err error) {
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

func (p *SysUserListThriftServiceGetSysUserListResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SysUserListThriftServiceGetSysUserListResult(%+v)", *p)
}