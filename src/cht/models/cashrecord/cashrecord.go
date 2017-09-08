package cashrecord

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	EVEN_DAYS_QUANTUM = int64(24 * 3600 * 7)
	ONE_MONTH_QUANTUM = int64(24 * 3600 * 30)
	TWO_MONTH_QUANTUM = int64(24 * 3600 * 60)
)

type CashRecordRequestStruct struct {
	UserID               int32
	StartTime            int32
	EndTime              int32
	QueryTime            int32
	RechargeStatus       int32
	LimitOffset          int32
	LimitNum             int32
	ChengHuiTongTraceLog string
}

type CashRecordStruct struct {
	ID             int32  `orm:"column(id)"`
	UserID         int32  `orm:"column(user_id)"`
	OrderSn        string `orm:"column(order_sn)"`
	Money          string `orm:"column(money)"`
	Credited       string `orm:"column(credited)"`
	Fee            string `orm:"column(fee)"`
	UseReturnMoney string `orm:"column(use_return_money)"`
	UseFreeNum     int32  `orm:"column(use_free_num)"`
	Addtime        int32  `orm:"column(addtime)"`
	Status         int32  `orm:"column(status)"`
	PayWay         int32  `orm:"column(pay_way)"`
	DealTime       int32  `orm:"column(deal_time)"`
	FailResult     string `orm:"column(fail_result)"`
}

/**
 * [GetRechargeRecord 查询提现记录]
 * @param    rrr *RechargeRecordRequest请求入参 (
 * @return   int32 返回不带limit的查询记录总数
 * @return   *CashRecordStruct 返回充值查询记录信息
 * @DateTime 2017-09-04T17:02:40+0800
 */
func GetCashRecord(crrs *CashRecordRequestStruct) ([]CashRecordStruct, int32, error) {
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_hs_cash").
		Where(fmt.Sprintf("user_id=%d", crrs.UserID))

	/*两个都查全部，开始时间和结束时间才有效*/
	if crrs.RechargeStatus == 0 && crrs.QueryTime == 0 {
		if crrs.StartTime != 0 {
			qb.And(fmt.Sprintf("addtime >= %d", crrs.StartTime))
		}
		if crrs.EndTime != 0 {
			qb.And(fmt.Sprintf("addtime <= %d", crrs.EndTime))
		}
	} else {
		switch {
		case crrs.QueryTime == 1:
			/*查最近七天充值记录*/
			qb.And(fmt.Sprintf("addtime >=%d", time.Now().Unix()-EVEN_DAYS_QUANTUM))
		case crrs.QueryTime == 2:
			/*查最近一个月充值记录*/
			qb.And(fmt.Sprintf("addtime >=%d", time.Now().Unix()-ONE_MONTH_QUANTUM))
		case crrs.QueryTime == 3:
			/*查最近两个月充值记录*/
			qb.And(fmt.Sprintf("addtime >=%d", time.Now().Unix()-TWO_MONTH_QUANTUM))
		}

		switch {
		case crrs.RechargeStatus == 1:
			/*查已成功充值记录*/
			qb.And(fmt.Sprintf("status=%d", 1))
		case crrs.RechargeStatus == 2:
			/*查审核中充值记录*/
			qb.And(fmt.Sprintf("status=%d", 0))
		case crrs.RechargeStatus == 3:
			/*查审核失败*/
			qb.And(fmt.Sprintf("status=%d", 2))
		}
	}
	qb.OrderBy("addtime").Desc()
	sql := qb.String()
	Logger.Debugf("GetCashRecord origin sql:", sql)

	var crs1 []CashRecordStruct
	totalnum, err := o.Raw(sql).QueryRows(&crs1)
	if err != nil {
		Logger.Debug("GetCashRecord query failed:", err)
		return nil, 0, err
	}
	/*得到总的查询数*/
	Logger.Debug("GetCashRecord query totalnum:", totalnum)

	/*带limit查询得到提现记录数据*/
	if crrs.LimitNum != 0 {
		qb.Limit(int(crrs.LimitNum))
	}
	if crrs.LimitOffset != 0 {
		qb.Offset(int(crrs.LimitOffset))
	}
	sql = qb.String()
	Logger.Debugf("GetCashRecord sql with limit:", sql)

	var crs []CashRecordStruct
	_, err = o.Raw(sql).QueryRows(&crs)
	if err != nil {
		Logger.Debug("GetCashRecord queryrows failed")
		return nil, 0, err
	}
	Logger.Debugf("GetCashRecord res:%v %d", crs, totalnum)
	return crs, int32(totalnum), nil
}
