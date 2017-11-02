package kefudutyupdate

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type KefuDutyUpdateRequest struct {
	ID                   int32
	Customer             string
	DutyTime             int32
	HolidayUser          string
	IsRest               int32
	Starttime            int32
	Endtime              int32
	ChengHuiTongTraceLog string
}

/**
 * [updateKefuDuty 客服值班---修改值班服务]
 * @param    kdur *KefuDutyUpdateRequest 请求入参
 * @return   bool true:修改成功  false:修改失败
 * @DateTime 2017-10-30T14:25:08+0800
 */
func UpdateKefuDuty(kdur *KefuDutyUpdateRequest) bool {
	Logger.Debugf("updateKefuDuty input param:%v", kdur)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Update("jl_customer_plan")

	if kdur.Customer == "" || kdur.DutyTime == 0 {
		Logger.Errorf("UpdateKefuDuty customer or holidayuser is null")
		return false
	}

	var str string
	if kdur.Customer != "" {
		str += fmt.Sprintf("customer=\"%s\",", kdur.Customer)
	}

	if kdur.HolidayUser != "" {
		str += fmt.Sprintf("holiday_user=\"%s\",", kdur.HolidayUser)
	}

	if kdur.DutyTime != 0 {
		str += fmt.Sprintf("duty_time=%d,", kdur.DutyTime)
	}

	if kdur.IsRest != 0 {
		str += fmt.Sprintf("is_rest=%d,", kdur.IsRest)
	}

	if kdur.Starttime != 0 {
		str += fmt.Sprintf("start=%d,", kdur.Starttime)
	}

	if kdur.Endtime != 0 {
		str += fmt.Sprintf("end=%d,", kdur.Endtime)
	}

	str = strings.TrimSuffix(str, ",")
	qb.Set(str)
	qb.Where(fmt.Sprintf("id=%d", kdur.ID))
	sql := qb.String()

	Logger.Debug("updateKefuDuty sql:", sql)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		Logger.Debugf("updateKefuDuty update failed :%v", err)
		return false
	}
	num, _ := res.RowsAffected()
	Logger.Debugf("updateKefuDuty change num :%v", num)
	if num == 0 {
		return false
	}
	return true
}
