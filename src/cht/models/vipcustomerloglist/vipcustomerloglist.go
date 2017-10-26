package vipcustomerloglist

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type VipCustomerLogListRequest struct {
	StartTime            int32  `thrift:"start_time,1" db:"start_time" json:"start_time"`
	EndTime              int32  `thrift:"end_time,2" db:"end_time" json:"end_time"`
	Keywords             string `thrift:"keywords,3" db:"keywords" json:"keywords"`
	Type                 int32  `thrift:"type,4" db:"type" json:"type"`
	LimitOffset          int32  `thrift:"limitOffset,5" db:"limitOffset" json:"limitOffset"`
	LimitNum             int32  `thrift:"limitNum,6" db:"limitNum" json:"limitNum"`
	ChengHuiTongTraceLog string `thrift:"chengHuiTongTraceLog,7" db:"chengHuiTongTraceLog" json:"chengHuiTongTraceLog"`
}

type VipCustomerDetails struct {
	ID            int32  `orm:"column(id)"`
	UserID        int32  `orm:"column(user_id)"`
	Username      string `orm:"column(username)"`
	Email         string `orm:"column(email)"`
	Realname      string `orm:"column(realname)"`
	Phone         string `orm:"column(phone)"`
	ScenePasstime int32  `orm:"column(scene_passtime)"`
	VipStatus     int32  `orm:"column(vip_status)"`
	VipPasstime   int32  `orm:"column(vip_passtime)"`
	VipVerifytime int32  `orm:"column(vip_verifytime)"`
	OldCustomer   int32  `orm:"column(old_customer)"`
	NewCustomer_  int32  `orm:"column(new_customer)"`
	Updatetime    int32  `orm:"column(updatetime)"`
	Remark        string `orm:"column(remark)"`
}

func GetVipCustomerLogListTotalNum(vclr *VipCustomerLogListRequest) (int32, error) {
	Logger.Debugf("GetVipCustomerLogListTotalNum input param:%v", vclr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("COUNT(1)").
		From("jl_change_vip").
		Where("1=1")

	if vclr.StartTime != 0 {
		qb.And(fmt.Sprintf("vip_passtime>=%d", vclr.StartTime))
	}

	if vclr.EndTime != 0 && vclr.StartTime <= vclr.EndTime {
		qb.And(fmt.Sprintf("vip_passtime<=%d", vclr.EndTime))
	}

	if vclr.Type != 0 && vclr.Keywords != "" {
		if vclr.Type == 1 {
			qb.And(fmt.Sprintf("username like \"%%%s%%\"", vclr.Keywords))
		} else if vclr.Type == 2 {
			qb.And(fmt.Sprintf("user_id=\"%s\"", vclr.Keywords))
		} else if vclr.Type == 3 {
			qb.And(fmt.Sprintf("realname like \"%%%s%%\"", vclr.Keywords))
		}
	}
	qb.OrderBy("id").Desc()

	sql := qb.String()
	Logger.Debugf("GetVipCustomerLogListTotalNum sql:%v", sql)

	var totalNum int32
	err := o.Raw(sql).QueryRow(&totalNum)
	if err != nil {
		Logger.Errorf("GetVipCustomerLogListTotalNum query failed :%v", err)
		return 0, err
	}
	Logger.Debugf("GetVipCustomerLogListTotalNum return num:%v", totalNum)
	return totalNum, nil
}

func GetVipCustomerLogList(vclr *VipCustomerLogListRequest) ([]VipCustomerDetails, error) {
	Logger.Debugf("GetVipCustomerLogList input param:%v", vclr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_change_vip").
		Where("1=1")

	if vclr.StartTime != 0 {
		qb.And(fmt.Sprintf("vip_passtime>=%d", vclr.StartTime))
	}

	if vclr.EndTime != 0 && vclr.StartTime <= vclr.EndTime {
		qb.And(fmt.Sprintf("vip_passtime<=%d", vclr.EndTime))
	}

	if vclr.Type != 0 && vclr.Keywords != "" {
		if vclr.Type == 1 {
			qb.And(fmt.Sprintf("username like \"%%%s%%\"", vclr.Keywords))
		} else if vclr.Type == 2 {
			qb.And(fmt.Sprintf("user_id=\"%s\"", vclr.Keywords))
		} else if vclr.Type == 3 {
			qb.And(fmt.Sprintf("realname like \"%%%s%%\"", vclr.Keywords))
		}
	}
	qb.OrderBy("id").Desc()

	if vclr.LimitNum != 0 {
		qb.Limit(int(vclr.LimitNum))
	}
	if vclr.LimitOffset != 0 {
		qb.Offset(int(vclr.LimitOffset))
	}

	sql := qb.String()
	Logger.Debugf("GetVipCustomerLogList sql:%v", sql)

	var vcd []VipCustomerDetails
	_, err := o.Raw(sql).QueryRows(&vcd)
	if err != nil {
		Logger.Errorf("GetVipCustomerLogList query failed :%v", err)
		return nil, err
	}
	Logger.Debugf("GetVipCustomerLogList return value:%v", vcd)
	return vcd, nil
}
