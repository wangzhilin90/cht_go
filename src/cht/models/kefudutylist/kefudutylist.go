package kefudutylist

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	NEXT_DAY_TIME = 24 * 3600
)

type KefuDutyListRequest struct {
	StartTime            int32  `thrift:"start_time,1" db:"start_time" json:"start_time"`
	EndTime              int32  `thrift:"end_time,2" db:"end_time" json:"end_time"`
	Kefu                 int32  `thrift:"kefu,3" db:"kefu" json:"kefu"`
	IsExport             int32  `thrift:"is_export,4" db:"is_export" json:"is_export"`
	LimitOffset          int32  `thrift:"limitOffset,5" db:"limitOffset" json:"limitOffset"`
	LimitNum             int32  `thrift:"limitNum,6" db:"limitNum" json:"limitNum"`
	ChengHuiTongTraceLog string `thrift:"chengHuiTongTraceLog,7" db:"chengHuiTongTraceLog" json:"chengHuiTongTraceLog"`
}

type KefuDutyListResult struct {
	ID          int32  `orm:"column(id)"`
	Customer    string `orm:"column(customer)"`
	IsRest      int32  `orm:"column(is_rest)"`
	DutyTime    int32  `orm:"column(duty_time)"`
	HolidayUser string `orm:"column(holiday_user)"`
	Addtime     int32  `orm:"column(addtime)"`
	Starttime   int32  `orm:"column(start)"`
	Endtime     int32  `orm:"column(end)"`
}

/**
 * [GetKefuDutyListTotalNum 客服值班列表总条目数]
 * @param    kdlr *KefuDutyListRequest 请求入参
 * @return   int32 总条目数
 * @DateTime 2017-10-28T17:10:55+0800
 */
func GetKefuDutyListTotalNum(kdlr *KefuDutyListRequest) (int32, error) {
	Logger.Debugf("GetKefuDutyListTotalNum input param:%v", kdlr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("COUNT(1) FROM jl_customer_plan").
		Where("1=1")

	if kdlr.StartTime != 0 {
		qb.And(fmt.Sprintf("duty_time>=%v", kdlr.StartTime))
	}

	if kdlr.EndTime != 0 && kdlr.StartTime <= kdlr.EndTime {
		qb.And(fmt.Sprintf("duty_time<=%v", kdlr.EndTime))
	}

	if kdlr.Kefu != 0 {
		qb.And(fmt.Sprintf("customer=%d", kdlr.Kefu))
	}

	sql := qb.String()
	Logger.Debugf("GetKefuDutyListTotalNum sql:%v", sql)
	var totalNum int32
	err := o.Raw(sql).QueryRow(&totalNum)
	if err == orm.ErrNoRows {
		return 0, nil
	} else if err != nil {
		Logger.Errorf("GetKefuDutyListTotalNum query failed :%v", err)
		return 0, err
	}
	Logger.Debugf("GetKefuDutyListTotalNum return num:%v", totalNum)
	return totalNum, nil
}

/**
 * [GetKefuDutyList 客服值班列表详情]
 * @param    kdlr *KefuDutyListRequest 请求入参
 * @return   []KefuDutyListResult 值班列表详情
 * @DateTime 2017-10-28T17:10:01+0800
 */
func GetKefuDutyList(kdlr *KefuDutyListRequest) ([]KefuDutyListResult, error) {
	Logger.Debugf("GetKefuDutyList input param:%v", kdlr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_customer_plan").
		Where("1=1")

	if kdlr.StartTime != 0 {
		qb.And(fmt.Sprintf("duty_time>=%v", kdlr.StartTime))
	}

	if kdlr.EndTime != 0 && kdlr.StartTime <= kdlr.EndTime {
		qb.And(fmt.Sprintf("duty_time<=%v", kdlr.EndTime+NEXT_DAY_TIME))
	}

	if kdlr.Kefu != 0 {
		qb.And(fmt.Sprintf("customer=%d", kdlr.Kefu))
	}

	if kdlr.LimitNum != 0 {
		qb.Limit(int(kdlr.LimitNum))
	}

	if kdlr.LimitOffset != 0 {
		qb.Offset(int(kdlr.LimitOffset))
	}

	sql := qb.String()
	Logger.Debugf("GetKefuDutyList sql:%v", sql)

	var kdl []KefuDutyListResult
	_, err := o.Raw(sql).QueryRows(&kdl)
	if err != nil {
		Logger.Errorf("GetKefuDutyList query failed :%v", err)
		return nil, err
	}
	Logger.Debugf("GetKefuDutyList return value:%v", kdl)
	return kdl, nil
}
