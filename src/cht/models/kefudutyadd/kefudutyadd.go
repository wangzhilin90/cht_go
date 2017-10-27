package kefudutyadd

import (
	"bytes"
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type KefuDutyAddRequest struct {
	Customer             string
	DutyTime             int32
	HolidayUser          string
	IsRest               int32
	Addtime              int32
	Starttime            int32
	Endtime              int32
	ChengHuiTongTraceLog string
}

/**
 * [AddKefuDuty 客服值班---新增值班]
 * @param    kdar *KefuDutyAddRequest 请求入参
 * @return   bool true:添加成功  false:添加失败
 * @DateTime 2017-10-27T15:23:34+0800
 */
func AddKefuDuty(kdar *KefuDutyAddRequest) bool {
	Logger.Debugf("AddKefuDuty input param: %v", kdar)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("INSERT INTO jl_customer_plan (id,customer,duty_time,holiday_user,is_rest,addtime,start,end) ")
	buf.WriteString("VALUES (next VALUE FOR MYCATSEQ_CUSTOMER_PLAN,?,?,?,?,?,?,?)")
	sql := buf.String()

	Logger.Debugf("AddKefuDuty sql: %v", sql)

	res, err := o.Raw(sql,
		kdar.Customer,
		kdar.DutyTime,
		kdar.HolidayUser,
		kdar.IsRest,
		time.Now().Unix(),
		kdar.Starttime,
		kdar.Endtime,
	).Exec()
	if err != nil {
		Logger.Errorf("AddKefuDuty insert failed:%v", err)
		return false
	}
	num, _ := res.RowsAffected()
	if num == 0 {
		return false
	}
	return true
}
