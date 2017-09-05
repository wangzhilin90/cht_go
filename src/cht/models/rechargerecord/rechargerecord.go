package rechargerecord

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

type RechargeRecordRequest struct {
	UserID               int32
	StartTime            int32
	EndTime              int32
	QueryTime            int32
	RechargeStatus       int32
	LimitOffset          int32
	LimitNum             int32
	ChengHuiTongTraceLog string
}

type RechargeRecordStruct struct {
	ID         int32  `orm:"column(id)"`
	UserID     int32  `orm:"column(user_id)"`
	OrderSn    string `orm:"column(order_sn)"`
	Money      string `orm:"column(money)"`
	Addtime    int32  `orm:"column(addtime)"`
	Status     int32  `orm:"column(status)"`
	DealTime   int32  `orm:"column(deal_time)"`
	PayType    int32  `orm:"column(pay_type)"`
	PayWay     int32  `orm:"column(pay_way)"`
	FailResult string `orm:"column(fail_result)"`
}

/**
 * [GetRechargeRecord 查询充值记录]
 * @param    rrr *RechargeRecordRequest请求入参 (
 * @return   int32 返回不带limit的查询记录总数
 * @return   *RechargeRecordStruct 返回充值查询记录信息
 * @DateTime 2017-09-04T17:02:40+0800
 */
func GetRechargeRecord(rrr *RechargeRecordRequest) ([]RechargeRecordStruct, int32, error) {
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_hs_recharge").
		Where(fmt.Sprintf("user_id=%d", rrr.UserID))

	/*两个都查全部，开始时间和结束时间才有效*/
	if rrr.RechargeStatus == 0 && rrr.QueryTime == 0 {
		if rrr.StartTime != 0 {
			qb.And(fmt.Sprintf("addtime >= %d", rrr.StartTime))
		}
		if rrr.EndTime != 0 {
			qb.And(fmt.Sprintf("addtime <= %d", rrr.EndTime))
		}
	} else {
		switch {
		case rrr.QueryTime == 1:
			/*查最近七天充值记录*/
			Logger.Debug(time.Now().Unix())
			Logger.Debug(EVEN_DAYS_QUANTUM)
			qb.And(fmt.Sprintf("addtime >=%d", time.Now().Unix()-EVEN_DAYS_QUANTUM))
		case rrr.QueryTime == 2:
			/*查最近一个月充值记录*/
			qb.And(fmt.Sprintf("addtime >=%d", time.Now().Unix()-ONE_MONTH_QUANTUM))
		case rrr.QueryTime == 3:
			/*查最近两个月充值记录*/
			qb.And(fmt.Sprintf("addtime >=%d", time.Now().Unix()-TWO_MONTH_QUANTUM))
		}

		switch {
		case rrr.RechargeStatus == 1:
			/*查已成功充值记录*/
			qb.And(fmt.Sprintf("status=%d", 1))
		case rrr.RechargeStatus == 2:
			/*查审核中充值记录*/
			qb.And(fmt.Sprintf("status=%d", 0))
		case rrr.RechargeStatus == 3:
			/*查审核失败*/
			qb.And(fmt.Sprintf("status=%d", 2))
		}
	}
	qb.OrderBy("addtime").Desc()
	sql := qb.String()
	Logger.Debugf("GetRechargeRecord origin sql", sql)

	var rrs1 []RechargeRecordStruct
	totalnum, err := o.Raw(sql).QueryRows(&rrs1)
	if err != nil {
		Logger.Error("GetRechargeRecord query failed:", err)
		return nil, 0, err
	}
	/*得到总的查询数*/
	Logger.Debug("GetRechargeRecord query totalnum:", totalnum)

	/*带limit查询得到充值记录数据*/
	if rrr.LimitNum != 0 {
		qb.Limit(int(rrr.LimitNum))
	}
	if rrr.LimitOffset != 0 {
		qb.Offset(int(rrr.LimitOffset))
	}
	sql = qb.String()
	Logger.Debugf("GetRechargeRecord sql with limit:", sql)

	var rrs []RechargeRecordStruct
	_, err = o.Raw(sql).QueryRows(&rrs)
	if err != nil {
		Logger.Debug("GetRechargeRecord queryrows failed")
		return nil, 0, err
	}
	Logger.Debugf("GetRechargeRecord res:%v %d", rrs, totalnum)
	return rrs, int32(totalnum), nil
}