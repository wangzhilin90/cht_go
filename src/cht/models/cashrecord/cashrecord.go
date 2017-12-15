package cashrecord

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

type CashStats struct {
	Money string //成功的总提现金额
	Fee   string //成功的总提现手续费
}

/**
 * [GetCashStats 得到提现成功的总金额和手续费]
 * @param    crrs *CashRecordRequestStruct 请求入参
 * @return   CashStats 返回提现的总金额和总手续费
 * @DateTime 2017-09-11T14:43:41+0800
 */
func GetCashStats(crrs *CashRecordRequestStruct) (*CashStats, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("GetCashStats input param:", crrs)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("SUM(money) AS money,SUM(fee) AS fee").
		From("jl_hs_cash").
		Where(fmt.Sprintf("user_id=%d", crrs.UserID)).
		And(fmt.Sprintf("status=1"))

	sql := qb.String()
	Logger.Debug("GetCashStats sql:", sql)
	var cs CashStats
	err := o.Raw(sql).QueryRow(&cs)
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err != nil {
		Logger.Error("GetCashStats query failed:", err)
		return nil, err
	}
	return &cs, nil
}

/**
 * [GetRechargeRecord 查询提现记录]
 * @param    rrr *RechargeRecordRequest请求入参 (
 * @return   int32 返回不带limit的查询记录总数
 * @return   *CashRecordStruct 返回充值查询记录信息
 * @DateTime 2017-09-04T17:02:40+0800
 */
func GetCashRecord(crrs *CashRecordRequestStruct) ([]CashRecordStruct, *CashStats, int32, error) {
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_hs_cash").
		Where(fmt.Sprintf("user_id=%d", crrs.UserID))

	//查询天数, 0:查全部  1：查最近7天 2：查一个月 3：查两个月
	switch {
	/*时间查全部，开始时间和结束时间才有效*/
	case crrs.QueryTime == 0:
		if crrs.StartTime != 0 {
			qb.And(fmt.Sprintf("addtime >= %d", crrs.StartTime))
		}
		if crrs.EndTime != 0 {
			qb.And(fmt.Sprintf("addtime <= %d", crrs.EndTime))
		}
	case crrs.QueryTime == 1:
		/*查最近七天充值记录*/
		qb.And(fmt.Sprintf("addtime >=%d", time.Now().Unix()-EVEN_DAYS_QUANTUM)).
			And(fmt.Sprintf("addtime <=%d", localtime.GetLocalZeroTime()))
	case crrs.QueryTime == 2:
		/*查最近一个月充值记录*/
		qb.And(fmt.Sprintf("addtime >=%d", time.Now().Unix()-ONE_MONTH_QUANTUM)).
			And(fmt.Sprintf("addtime <=%d", localtime.GetLocalZeroTime()))
	case crrs.QueryTime == 3:
		/*查最近两个月充值记录*/
		qb.And(fmt.Sprintf("addtime >=%d", time.Now().Unix()-TWO_MONTH_QUANTUM)).
			And(fmt.Sprintf("addtime <=%d", localtime.GetLocalZeroTime()))
	}

	//充值状态,0:查全部 1:已成功 2:审核中  3:审核失败
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

	qb.OrderBy("addtime").Desc()
	sql := qb.String()
	Logger.Debugf("GetCashRecord origin sql:", sql)

	var crs1 []CashRecordStruct
	totalnum, err := o.Raw(sql).QueryRows(&crs1)
	if err != nil {
		Logger.Errorf("GetCashRecord query failed:%v", err)
		return nil, nil, 0, err
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
		Logger.Errorf("GetCashRecord queryrows failed:%v", err)
		return nil, nil, 0, err
	}

	cs, err := GetCashStats(crrs)
	if err != nil {
		Logger.Error("GetCashStats failed", err)
		return nil, nil, 0, err
	}

	Logger.Debugf("GetCashRecord res:%v %v %d", crs, cs, totalnum)
	return crs, cs, int32(totalnum), nil
}
