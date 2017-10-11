package dutydetails

import (
	"cht/common/localtime"
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type DutyDetailsRequestStruct struct {
	ChengHuiTongTraceLog string
}

type DutyDetailsStruct struct {
	ID          int32  `orm:column(id)`
	Customer    string `orm:column(customer)`
	IsRest      int32  `orm:column(is_rest)`
	DutyTime    int32  `orm:column(duty_time)`
	HolidayUser string `orm:column(holiday_user)`
	StartTime   int32  `orm:column(start)`
	EndTime     int32  `orm:column(end)`
	Addtime     int32  `orm:column(addtime)`
}

func GetDutyDetails(ddrs *DutyDetailsRequestStruct) (*DutyDetailsStruct, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("CheckEmailUse input param:", ddrs)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_customer_plan").
		Where(fmt.Sprintf("duty_time >= %d", localtime.GetLocalZeroTime())).
		And(fmt.Sprintf("duty_time <= %d", localtime.GetLocal24Time())).
		And(fmt.Sprintf("start <= %d", time.Now().Unix())).
		And(fmt.Sprintf("end >= %d", time.Now().Unix())).
		Limit(1)

	sql := qb.String()
	Logger.Debug("GetDutyDetails sql:", sql)
	var dds DutyDetailsStruct
	err := o.Raw(sql).QueryRow(&dds)
	if err != nil {
		Logger.Debugf("GetDutyDetails query failed %v", err)
		return nil, err
	}
	Logger.Debugf("GetDutyDetails res :%v", dds)
	return &dds, nil
}
