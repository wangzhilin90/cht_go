package rechargerecord

import (
	"cht/common/localtime"
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
 * [GetRechargeTotalMoney 得到用户充值成功的总金额]
 * @param    rrr *RechargeRecordRequest请求入参
 * @return   string 返回用户充值成功的总金额
 * @DateTime 2017-09-08T15:36:16+0800
 */
func GetRechargeTotalMoney(rrr *RechargeRecordRequest) (string, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("GetRechargeTotalMoney input param:", rrr)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("SUM(money)").
		From("jl_hs_recharge").
		Where(fmt.Sprintf("user_id=%d", rrr.UserID)).
		And(fmt.Sprintf("status=1"))

	sql := qb.String()
	Logger.Debug("GetRechargeTotalMoney sql:", sql)
	var TotalMoney string
	err := o.Raw(sql).QueryRow(&TotalMoney)
	if err != nil {
		Logger.Error("GetRechargeTotalMoney query failed:", err)
		return "", err
	}
	Logger.Debugf("GetRechargeTotalMoney res ", TotalMoney)
	return TotalMoney, nil
}

/**
 * [GetRechargeRecord 查询充值记录]
 * @param    rrr *RechargeRecordRequest请求入参 (
 * @return   *RechargeRecordStruct 返回充值查询记录信息
 * @return   int32 充值总记录数
 * @return   string 充值总金额
 * @DateTime 2017-09-04T17:02:40+0800
 */
func GetRechargeRecord(rrr *RechargeRecordRequest) ([]RechargeRecordStruct, int32, string, error) {
	money, err := GetRechargeTotalMoney(rrr)
	if err != nil || money == "" {
		money = "0.00"
	}

	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_hs_recharge").
		Where(fmt.Sprintf("user_id=%d", rrr.UserID))

	//查询天数 0:查全部  1：查最近7天 2：查一个月 3：查两个月
	switch {
	/*时间查全部，开始时间和结束时间才有效*/
	case rrr.QueryTime == 0:
		if rrr.StartTime != 0 {
			qb.And(fmt.Sprintf("addtime >= %d", rrr.StartTime))
		}
		if rrr.EndTime != 0 {
			qb.And(fmt.Sprintf("addtime <= %d", rrr.EndTime))
		}
	case rrr.QueryTime == 1:
		/*查最近七天充值记录*/
		qb.And(fmt.Sprintf("addtime >=%d", time.Now().Unix()-EVEN_DAYS_QUANTUM)).
			And(fmt.Sprintf("addtime <=%d", localtime.GetLocalZeroTime()))
	case rrr.QueryTime == 2:
		/*查最近一个月充值记录*/
		qb.And(fmt.Sprintf("addtime >=%d", time.Now().Unix()-ONE_MONTH_QUANTUM)).
			And(fmt.Sprintf("addtime <=%d", localtime.GetLocalZeroTime()))
	case rrr.QueryTime == 3:
		/*查最近两个月充值记录*/
		qb.And(fmt.Sprintf("addtime >=%d", time.Now().Unix()-TWO_MONTH_QUANTUM)).
			And(fmt.Sprintf("addtime <=%d", localtime.GetLocalZeroTime()))
	}

	//充值状态,0:查全部 1:已成功 2:审核中  3:审核失败
	switch {
	/*查已成功充值记录*/
	case rrr.RechargeStatus == 1:
		qb.And(fmt.Sprintf("status=%d", 1))
	/*查审核中充值记录*/
	case rrr.RechargeStatus == 2:
		qb.And(fmt.Sprintf("status=%d", 0))
	/*查审核失败*/
	case rrr.RechargeStatus == 3:
		qb.And(fmt.Sprintf("status=%d", 2))
	}

	qb.OrderBy("addtime").Desc()
	sql := qb.String()
	Logger.Debugf("GetRechargeRecord origin sql", sql)

	var rrs1 []RechargeRecordStruct
	totalnum, err := o.Raw(sql).QueryRows(&rrs1)
	if err != nil {
		Logger.Error("GetRechargeRecord query failed:", err)
		return nil, 0, "0.00", err
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
		return nil, 0, "0.00", err
	}
	Logger.Debugf("GetRechargeRecord res:%v %d", rrs, totalnum)
	return rrs, int32(totalnum), money, nil
}
