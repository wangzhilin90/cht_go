package kefudutydetails

import (
	"bytes"
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type KefuDutyDetailsRequest struct {
	ID                   int32
	ChengHuiTongTraceLog string
}

type KefuDutyDetails struct {
	ID          int32  `orm:"column(id)"`
	Customer    string `orm:"column(customer)"`
	DutyTime    int32  `orm:"column(duty_time)"`
	HolidayUser string `orm:"column(holiday_user)"`
	IsRest      int32  `orm:"column(is_rest)"`
	Starttime   int32  `orm:"column(start)"`
	Endtime     int32  `orm:"column(end)"`
}

/**
 * [GetKefuDutyDetails 值班详情]
 * @param    kddr *KefuDutyDetailsRequest 请求入参
 * @return   *KefuDutyDetails 值班详情
 * @DateTime 2017-10-28T14:06:04+0800
 */
func GetKefuDutyDetails(kddr *KefuDutyDetailsRequest) (*KefuDutyDetails, error) {
	Logger.Debugf("GetKefuDutyDetails input param: %v", kddr)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("SELECT * FROM jl_customer_plan WHERE id = ? LIMIT 1")
	sql := buf.String()

	Logger.Debugf("GetKefuDutyDetails sql: %v", sql)

	var kdd KefuDutyDetails
	err := o.Raw(sql, kddr.ID).QueryRow(&kdd)
	if err != nil {
		Logger.Debugf("GetKefuDutyDetails query failed %v", err)
		return nil, err
	}
	Logger.Debugf("GetKefuDutyDetails return value %v", kdd)
	return &kdd, nil
}
