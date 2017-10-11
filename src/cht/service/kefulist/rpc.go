// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package kefulist

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
//  - RoleID
//  - Status
//  - CustomerType
//  - ChengHuiTongTraceLog
type KeFuListRequestStruct struct {
	RoleID               int32  `thrift:"role_id,1" db:"role_id" json:"role_id"`
	Status               int32  `thrift:"status,2" db:"status" json:"status"`
	CustomerType         string `thrift:"customer_type,3" db:"customer_type" json:"customer_type"`
	ChengHuiTongTraceLog string `thrift:"chengHuiTongTraceLog,4" db:"chengHuiTongTraceLog" json:"chengHuiTongTraceLog"`
}

// func NewKeFuListRequestStruct() *KeFuListRequestStruct {
//   return &KeFuListRequestStruct{}
// }

func (p *KeFuListRequestStruct) GetRoleID() int32 {
	return p.RoleID
}

func (p *KeFuListRequestStruct) GetStatus() int32 {
	return p.Status
}

func (p *KeFuListRequestStruct) GetCustomerType() string {
	return p.CustomerType
}

func (p *KeFuListRequestStruct) GetChengHuiTongTraceLog() string {
	return p.ChengHuiTongTraceLog
}
func (p *KeFuListRequestStruct) Read(iprot thrift.TProtocol) error {
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

func (p *KeFuListRequestStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.RoleID = v
	}
	return nil
}

func (p *KeFuListRequestStruct) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Status = v
	}
	return nil
}

func (p *KeFuListRequestStruct) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.CustomerType = v
	}
	return nil
}

func (p *KeFuListRequestStruct) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.ChengHuiTongTraceLog = v
	}
	return nil
}

func (p *KeFuListRequestStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("KeFuListRequestStruct"); err != nil {
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

func (p *KeFuListRequestStruct) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("role_id", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:role_id: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.RoleID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.role_id (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:role_id: ", p), err)
	}
	return err
}

func (p *KeFuListRequestStruct) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("status", thrift.I32, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:status: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Status)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.status (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:status: ", p), err)
	}
	return err
}

func (p *KeFuListRequestStruct) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("customer_type", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:customer_type: ", p), err)
	}
	if err := oprot.WriteString(string(p.CustomerType)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.customer_type (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:customer_type: ", p), err)
	}
	return err
}

func (p *KeFuListRequestStruct) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("chengHuiTongTraceLog", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:chengHuiTongTraceLog: ", p), err)
	}
	if err := oprot.WriteString(string(p.ChengHuiTongTraceLog)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.chengHuiTongTraceLog (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:chengHuiTongTraceLog: ", p), err)
	}
	return err
}

func (p *KeFuListRequestStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("KeFuListRequestStruct(%+v)", *p)
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
type KeFuDetailsStruct struct {
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

func NewKeFuDetailsStruct() *KeFuDetailsStruct {
	return &KeFuDetailsStruct{}
}

func (p *KeFuDetailsStruct) GetID() int32 {
	return p.ID
}

func (p *KeFuDetailsStruct) GetRoleID() int32 {
	return p.RoleID
}

func (p *KeFuDetailsStruct) GetAccount() string {
	return p.Account
}

func (p *KeFuDetailsStruct) GetRealname() string {
	return p.Realname
}

func (p *KeFuDetailsStruct) GetPassword() string {
	return p.Password
}

func (p *KeFuDetailsStruct) GetMobile() string {
	return p.Mobile
}

func (p *KeFuDetailsStruct) GetQq() string {
	return p.Qq
}

func (p *KeFuDetailsStruct) GetLastloginip() string {
	return p.Lastloginip
}

func (p *KeFuDetailsStruct) GetLastlogintime() int32 {
	return p.Lastlogintime
}

func (p *KeFuDetailsStruct) GetCreateTime() int32 {
	return p.CreateTime
}

func (p *KeFuDetailsStruct) GetStatus() int32 {
	return p.Status
}

func (p *KeFuDetailsStruct) GetViews() int32 {
	return p.Views
}

func (p *KeFuDetailsStruct) GetCustomerType() int32 {
	return p.CustomerType
}
func (p *KeFuDetailsStruct) Read(iprot thrift.TProtocol) error {
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

func (p *KeFuDetailsStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ID = v
	}
	return nil
}

func (p *KeFuDetailsStruct) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.RoleID = v
	}
	return nil
}

func (p *KeFuDetailsStruct) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Account = v
	}
	return nil
}

func (p *KeFuDetailsStruct) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.Realname = v
	}
	return nil
}

func (p *KeFuDetailsStruct) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.Password = v
	}
	return nil
}

func (p *KeFuDetailsStruct) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.Mobile = v
	}
	return nil
}

func (p *KeFuDetailsStruct) ReadField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.Qq = v
	}
	return nil
}

func (p *KeFuDetailsStruct) ReadField8(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 8: ", err)
	} else {
		p.Lastloginip = v
	}
	return nil
}

func (p *KeFuDetailsStruct) ReadField9(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 9: ", err)
	} else {
		p.Lastlogintime = v
	}
	return nil
}

func (p *KeFuDetailsStruct) ReadField10(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 10: ", err)
	} else {
		p.CreateTime = v
	}
	return nil
}

func (p *KeFuDetailsStruct) ReadField11(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 11: ", err)
	} else {
		p.Status = v
	}
	return nil
}

func (p *KeFuDetailsStruct) ReadField12(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 12: ", err)
	} else {
		p.Views = v
	}
	return nil
}

func (p *KeFuDetailsStruct) ReadField13(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 13: ", err)
	} else {
		p.CustomerType = v
	}
	return nil
}

func (p *KeFuDetailsStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("KeFuDetailsStruct"); err != nil {
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

func (p *KeFuDetailsStruct) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *KeFuDetailsStruct) writeField2(oprot thrift.TProtocol) (err error) {
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

func (p *KeFuDetailsStruct) writeField3(oprot thrift.TProtocol) (err error) {
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

func (p *KeFuDetailsStruct) writeField4(oprot thrift.TProtocol) (err error) {
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

func (p *KeFuDetailsStruct) writeField5(oprot thrift.TProtocol) (err error) {
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

func (p *KeFuDetailsStruct) writeField6(oprot thrift.TProtocol) (err error) {
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

func (p *KeFuDetailsStruct) writeField7(oprot thrift.TProtocol) (err error) {
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

func (p *KeFuDetailsStruct) writeField8(oprot thrift.TProtocol) (err error) {
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

func (p *KeFuDetailsStruct) writeField9(oprot thrift.TProtocol) (err error) {
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

func (p *KeFuDetailsStruct) writeField10(oprot thrift.TProtocol) (err error) {
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

func (p *KeFuDetailsStruct) writeField11(oprot thrift.TProtocol) (err error) {
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

func (p *KeFuDetailsStruct) writeField12(oprot thrift.TProtocol) (err error) {
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

func (p *KeFuDetailsStruct) writeField13(oprot thrift.TProtocol) (err error) {
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

func (p *KeFuDetailsStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("KeFuDetailsStruct(%+v)", *p)
}

// Attributes:
//  - Status
//  - KeFuList
//  - Msg
type KeFuListResponseStruct struct {
	Status   int32                `thrift:"status,1" db:"status" json:"status"`
	KeFuList []*KeFuDetailsStruct `thrift:"KeFuList,2" db:"KeFuList" json:"KeFuList"`
	Msg      string               `thrift:"msg,3" db:"msg" json:"msg"`
}

func NewKeFuListResponseStruct() *KeFuListResponseStruct {
	return &KeFuListResponseStruct{}
}

func (p *KeFuListResponseStruct) GetStatus() int32 {
	return p.Status
}

func (p *KeFuListResponseStruct) GetKeFuList() []*KeFuDetailsStruct {
	return p.KeFuList
}

func (p *KeFuListResponseStruct) GetMsg() string {
	return p.Msg
}
func (p *KeFuListResponseStruct) Read(iprot thrift.TProtocol) error {
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

func (p *KeFuListResponseStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Status = v
	}
	return nil
}

func (p *KeFuListResponseStruct) ReadField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*KeFuDetailsStruct, 0, size)
	p.KeFuList = tSlice
	for i := 0; i < size; i++ {
		_elem0 := &KeFuDetailsStruct{}
		if err := _elem0.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem0), err)
		}
		p.KeFuList = append(p.KeFuList, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *KeFuListResponseStruct) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Msg = v
	}
	return nil
}

func (p *KeFuListResponseStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("KeFuListResponseStruct"); err != nil {
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

func (p *KeFuListResponseStruct) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *KeFuListResponseStruct) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("KeFuList", thrift.LIST, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:KeFuList: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.KeFuList)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.KeFuList {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:KeFuList: ", p), err)
	}
	return err
}

func (p *KeFuListResponseStruct) writeField3(oprot thrift.TProtocol) (err error) {
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

func (p *KeFuListResponseStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("KeFuListResponseStruct(%+v)", *p)
}

type KeFuListThriftService interface {
	// Parameters:
	//  - RequestObj
	GetKeFuList(requestObj *KeFuListRequestStruct) (r *KeFuListResponseStruct, err error)
}

type KeFuListThriftServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewKeFuListThriftServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *KeFuListThriftServiceClient {
	return &KeFuListThriftServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewKeFuListThriftServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *KeFuListThriftServiceClient {
	return &KeFuListThriftServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - RequestObj
func (p *KeFuListThriftServiceClient) GetKeFuList(requestObj *KeFuListRequestStruct) (r *KeFuListResponseStruct, err error) {
	if err = p.sendGetKeFuList(requestObj); err != nil {
		return
	}
	return p.recvGetKeFuList()
}

func (p *KeFuListThriftServiceClient) sendGetKeFuList(requestObj *KeFuListRequestStruct) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getKeFuList", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := KeFuListThriftServiceGetKeFuListArgs{
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

func (p *KeFuListThriftServiceClient) recvGetKeFuList() (value *KeFuListResponseStruct, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "getKeFuList" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getKeFuList failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getKeFuList failed: out of sequence response")
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
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getKeFuList failed: invalid message type")
		return
	}
	result := KeFuListThriftServiceGetKeFuListResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type KeFuListThriftServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      KeFuListThriftService
}

func (p *KeFuListThriftServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *KeFuListThriftServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *KeFuListThriftServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewKeFuListThriftServiceProcessor(handler KeFuListThriftService) *KeFuListThriftServiceProcessor {

	self3 := &KeFuListThriftServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self3.processorMap["getKeFuList"] = &keFuListThriftServiceProcessorGetKeFuList{handler: handler}
	return self3
}

func (p *KeFuListThriftServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
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

type keFuListThriftServiceProcessorGetKeFuList struct {
	handler KeFuListThriftService
}

func (p *keFuListThriftServiceProcessorGetKeFuList) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := KeFuListThriftServiceGetKeFuListArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getKeFuList", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := KeFuListThriftServiceGetKeFuListResult{}
	var retval *KeFuListResponseStruct
	var err2 error
	if retval, err2 = p.handler.GetKeFuList(args.RequestObj); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getKeFuList: "+err2.Error())
		oprot.WriteMessageBegin("getKeFuList", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getKeFuList", thrift.REPLY, seqId); err2 != nil {
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
type KeFuListThriftServiceGetKeFuListArgs struct {
	RequestObj *KeFuListRequestStruct `thrift:"requestObj,1" db:"requestObj" json:"requestObj"`
}

func NewKeFuListThriftServiceGetKeFuListArgs() *KeFuListThriftServiceGetKeFuListArgs {
	return &KeFuListThriftServiceGetKeFuListArgs{}
}

var KeFuListThriftServiceGetKeFuListArgs_RequestObj_DEFAULT *KeFuListRequestStruct

func (p *KeFuListThriftServiceGetKeFuListArgs) GetRequestObj() *KeFuListRequestStruct {
	if !p.IsSetRequestObj() {
		return KeFuListThriftServiceGetKeFuListArgs_RequestObj_DEFAULT
	}
	return p.RequestObj
}
func (p *KeFuListThriftServiceGetKeFuListArgs) IsSetRequestObj() bool {
	return p.RequestObj != nil
}

func (p *KeFuListThriftServiceGetKeFuListArgs) Read(iprot thrift.TProtocol) error {
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

func (p *KeFuListThriftServiceGetKeFuListArgs) ReadField1(iprot thrift.TProtocol) error {
	p.RequestObj = &KeFuListRequestStruct{}
	if err := p.RequestObj.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.RequestObj), err)
	}
	return nil
}

func (p *KeFuListThriftServiceGetKeFuListArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getKeFuList_args"); err != nil {
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

func (p *KeFuListThriftServiceGetKeFuListArgs) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *KeFuListThriftServiceGetKeFuListArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("KeFuListThriftServiceGetKeFuListArgs(%+v)", *p)
}

// Attributes:
//  - Success
type KeFuListThriftServiceGetKeFuListResult struct {
	Success *KeFuListResponseStruct `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewKeFuListThriftServiceGetKeFuListResult() *KeFuListThriftServiceGetKeFuListResult {
	return &KeFuListThriftServiceGetKeFuListResult{}
}

var KeFuListThriftServiceGetKeFuListResult_Success_DEFAULT *KeFuListResponseStruct

func (p *KeFuListThriftServiceGetKeFuListResult) GetSuccess() *KeFuListResponseStruct {
	if !p.IsSetSuccess() {
		return KeFuListThriftServiceGetKeFuListResult_Success_DEFAULT
	}
	return p.Success
}
func (p *KeFuListThriftServiceGetKeFuListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *KeFuListThriftServiceGetKeFuListResult) Read(iprot thrift.TProtocol) error {
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

func (p *KeFuListThriftServiceGetKeFuListResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &KeFuListResponseStruct{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *KeFuListThriftServiceGetKeFuListResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getKeFuList_result"); err != nil {
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

func (p *KeFuListThriftServiceGetKeFuListResult) writeField0(oprot thrift.TProtocol) (err error) {
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

func (p *KeFuListThriftServiceGetKeFuListResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("KeFuListThriftServiceGetKeFuListResult(%+v)", *p)
}