package customerlist

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type CustomerListRequest struct {
	Customer             int32  `thrift:"customer,1" db:"customer" json:"customer"`
	StartTime            int32  `thrift:"start_time,2" db:"start_time" json:"start_time"`
	EndTime              int32  `thrift:"end_time,3" db:"end_time" json:"end_time"`
	Islock               int32  `thrift:"islock,4" db:"islock" json:"islock"`
	Username             string `thrift:"username,5" db:"username" json:"username"`
	IsExport             int32  `thrift:"is_export,6" db:"is_export" json:"is_export"`
	LimitOffset          int32  `thrift:"limitOffset,7" db:"limitOffset" json:"limitOffset"`
	LimitNum             int32  `thrift:"limitNum,8" db:"limitNum" json:"limitNum"`
	ChengHuiTongTraceLog string `thrift:"chengHuiTongTraceLog,9" db:"chengHuiTongTraceLog" json:"chengHuiTongTraceLog"`
}

type CustomerListResult struct {
	ID               int32  `orm:"column(id)"`
	Username         string `orm:"column(username)"`
	Password         string `orm:"column(password)"`
	Paypassword      string `orm:"column(paypassword)"`
	Point            int32  `orm:"column(point)"`
	Email            string `orm:"column(email)"`
	Avatar           string `orm:"column(avatar)"`
	Sex              int32  `orm:"column(sex)"`
	Realname         string `orm:"column(realname)"`
	Phone            string `orm:"column(phone)"`
	Tel              string `orm:"column(tel)"`
	Birthday         int32  `orm:"column(birthday)"`
	Nation           int32  `orm:"column(nation)"`
	Province         int32  `orm:"column(province)"`
	City             int32  `orm:"column(city)"`
	Area             int32  `orm:"column(area)"`
	Address          string `orm:"column(address)"`
	Customer         int32  `orm:"column(customer)"`
	Logintime        int32  `orm:"column(logintime)"`
	Loginip          string `orm:"column(loginip)"`
	Addtime          int32  `orm:"column(addtime)"`
	Addip            string `orm:"column(addip)"`
	Islock           int32  `orm:"column(islock)"`
	Isvest           int32  `orm:"column(isvest)"`
	OsType           int32  `orm:"column(os_type)"`
	DeviceToken      string `orm:"column(device_token)"`
	WeinxinID        string `orm:"column(weinxin_id)"`
	BindTime         int32  `orm:"column(bind_time)"`
	InvitationCode   string `orm:"column(invitation_code)"`
	Source           string `orm:"column(source)"`
	Hsid             string `orm:"column(hsid)"`
	GStatus          int32  `orm:"column(g_status)"`
	GPassword        string `orm:"column(g_password)"`
	AutoProtocolCode string `orm:"column(auto_protocol_code)"`
	IsBorrower       int32  `orm:"column(is_borrower)"`
	IsWorker         int32  `orm:"column(is_worker)"`
	Hswaitactivate   int32  `orm:"column(hswaitactivate)"`
}

/**
 * [GetCustomerListTotalNum 专属客服列表总记录数]
 * @param    clr *CustomerListRequest 请求入参
 * @return   int32 总记录数
 * @DateTime 2017-10-26T13:53:53+0800
 */
func GetCustomerListTotalNum(clr *CustomerListRequest) (int32, error) {
	Logger.Debugf("GetCustomerListTotalNum input param:%v", clr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("COUNT(1)").
		From("jl_user").
		Where("1=1")

	if clr.Customer != 0 {
		qb.And(fmt.Sprintf("customer = %d", clr.Customer))
	}

	if clr.Username != "" {
		qb.And(fmt.Sprintf("username like \"%%%s%%\"", clr.Username))
	}

	if clr.StartTime != 0 {
		qb.And(fmt.Sprintf("addtime>%d", clr.StartTime))
	}

	if clr.EndTime != 0 && clr.StartTime <= clr.EndTime {
		qb.And(fmt.Sprintf("addtime<%d", clr.EndTime))
	}

	if clr.Islock != -1 {
		qb.And(fmt.Sprintf("islock=%d", clr.Islock))
	}

	sql := qb.String()

	Logger.Debug("GetCustomerListTotalNum sql:", sql)

	var totalNum int32
	err := o.Raw(sql).QueryRow(&totalNum)
	if err == orm.ErrNoRows {
		return 0, nil
	} else if err != nil {
		Logger.Errorf("GetCustomerListTotalNum query failed %v", err)
		return 0, err
	}
	Logger.Debugf("GetCustomerListTotalNum res %v", totalNum)
	return totalNum, nil
}

/**
 * [GetCustomerList 专属客服列表]
 * @param    clr *CustomerListRequest 请求入参
 * @return   []CustomerListResult 专属客服列表结果
 * @DateTime 2017-10-26T14:44:35+0800
 */
func GetCustomerList(clr *CustomerListRequest) ([]CustomerListResult, error) {
	Logger.Debugf("GetCustomerList input param:%v", clr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_user").
		Where("1=1")

	if clr.Customer != 0 {
		qb.And(fmt.Sprintf("customer = %d", clr.Customer))
	}

	if clr.Username != "" {
		qb.And(fmt.Sprintf("username like \"%%%s%%\"", clr.Username))
	}

	if clr.StartTime != 0 {
		qb.And(fmt.Sprintf("addtime>%d", clr.StartTime))
	}

	if clr.EndTime != 0 && clr.StartTime <= clr.EndTime {
		qb.And(fmt.Sprintf("addtime<%d", clr.EndTime))
	}

	if clr.Islock != -1 {
		qb.And(fmt.Sprintf("islock=%d", clr.Islock))
	}

	qb.OrderBy("id").Desc()

	//是否导出，默认为0不导出,此时limitnum和limitoffset生效
	if clr.IsExport == 0 {
		if clr.LimitNum != 0 {
			qb.Limit(int(clr.LimitNum))
		}

		if clr.LimitOffset != 0 {
			qb.Offset(int(clr.LimitOffset))
		}
	}

	sql := qb.String()
	Logger.Debug("GetCustomerList sql:", sql)

	var clrt []CustomerListResult
	_, err := o.Raw(sql).QueryRows(&clrt)
	if err != nil {
		Logger.Errorf("GetCustomerList query failed %v", err)
		return nil, err
	}
	Logger.Debugf("GetCustomerList res:%v", clrt)
	return clrt, nil
}
