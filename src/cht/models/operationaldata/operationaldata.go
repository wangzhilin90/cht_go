package operationaldata

import (
	"bytes"
	"cht/common/localtime"
	. "cht/common/logger"
	"cht/utils"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	ONE_MONTH_QUANTUM = int64(24 * 3600 * 30)
)

type OperationalDataRequestStruct struct {
	StartMonth           int32 //12个月前时间搓
	Start                int32 //1个月前时间搓
	TodayTime            int32
	YesterdayTime        int32
	TomorrowTime         int32
	ChengHuiTongTraceLog string
}

//最近30天投标排行
type ThirtyDaysResult struct {
	Money    string `orm:"column(money)"`
	Username string `orm:"column(username)"`
}

//最近12个月每月成交量
type TwelveMonthResult struct {
	Category string `orm:"column(category)"`
	Account  string `orm:"column(account)"`
}

//最近1个月每月成交量
type OneMonthResult struct {
	Category string `orm:"column(category)"`
	Account  string `orm:"column(account)"`
}

//借款周期占比
type PeriodResult struct {
	Category string `orm:"column(category)"`
	Column_1 string `orm:"column(column_1)"`
}

//投资金额占比
type InvestResult struct {
	A1 string `orm:"column(a1)"`
	A2 string `orm:"column(a2)"`
	A3 string `orm:"column(a3)"`
	A4 string `orm:"column(a4)"`
	A5 string `orm:"column(a5)"`
}

//标的比例
type BidResult struct {
	BorrowType int32  `orm:"column(borrow_type)"`
	Number     string `orm:"column(number)"`
}

//实时待收排行榜
type WaitResult struct {
	Money    string `orm:"column(money)"`
	Username string `orm:"column(username)"`
}

/*得到最近30天投标排行结果*/
func GetThirtyDaysResult(odrs *OperationalDataRequestStruct) ([]ThirtyDaysResult, error) {
	Logger.Debugf("GetThirtyDaysResult input param: %v", odrs)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("SELECT SUM(BT.account_act) AS money,U.username FROM jl_borrow_tender BT LEFT JOIN jl_user U ON BT.user_id=U.id ")
	buf.WriteString("WHERE BT.status=1 AND BT.addtime>=? AND BT.addtime<=? AND U.isvest=0 GROUP BY BT.user_id ORDER BY money DESC LIMIT 10")
	sql := buf.String()

	Logger.Debugf("GetThirtyDaysResult sql: %v", sql)

	var tdr []ThirtyDaysResult
	_, err := o.Raw(sql, time.Now().Unix()-ONE_MONTH_QUANTUM, localtime.GetLocal24Time()).QueryRows(&tdr)
	if err != nil {
		Logger.Errorf("GetThirtyDaysResult query failed %v", err)
		return nil, err
	}
	Logger.Debugf("GetThirtyDaysResult return value:%v", tdr)
	return tdr, nil
}

/*得到最近12个月每月成交量结果*/
func GetTwelveMonthResult(odrs *OperationalDataRequestStruct) ([]TwelveMonthResult, error) {
	Logger.Debugf("GetTwelveMonthResult input param: %v", odrs)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("SELECT FROM_UNIXTIME(`review_time`,'%Y-%m') AS category,SUM(account) AS account FROM jl_borrow ")
	buf.WriteString("WHERE review_time>=? AND status IN (2,3,6,7) GROUP BY category ORDER BY category ")
	sql := buf.String()

	Logger.Debugf("GetTwelveMonthResult sql: %v", sql)

	var tmr []TwelveMonthResult
	_, err := o.Raw(sql, odrs.StartMonth).QueryRows(&tmr)
	if err != nil {
		Logger.Errorf("GetTwelveMonthResult query failed %v", err)
		return nil, err
	}
	Logger.Debugf("GetTwelveMonthResult return value:%v", tmr)
	return tmr, nil
}

/*得到最近1个月每月成交量结果*/
func GetOneMonthResult(odrs *OperationalDataRequestStruct) ([]OneMonthResult, error) {
	Logger.Debugf("GetOneMonthResultStruct input param: %v", odrs)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("SELECT FROM_UNIXTIME(`review_time`,'%Y-%m') AS category,SUM(account) AS account FROM jl_borrow ")
	buf.WriteString("WHERE review_time>=? AND status IN (2,3,6,7) GROUP BY category ORDER BY category LIMIT 1 ")
	sql := buf.String()

	Logger.Debugf("GetOneMonthResultStruct sql: %v", sql)

	var omrs []OneMonthResult
	_, err := o.Raw(sql, odrs.Start).QueryRows(&omrs)
	if err != nil {
		Logger.Errorf("GetOneMonthResultStruct query failed %v", err)
		return nil, err
	}
	Logger.Debugf("GetOneMonthResultStruct return value:%v", omrs)
	return omrs, nil
}

/*得到借款周期占比结果*/
func GetPeriodResult(odrs *OperationalDataRequestStruct) ([]PeriodResult, error) {
	Logger.Debugf("GetPeriodResultStruct input param: %v", odrs)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("SELECT time_limit AS category,COUNT(1) AS 'column_1' FROM jl_borrow ")
	buf.WriteString("WHERE status IN (2,3,6,7) AND addtime>1420041600 GROUP BY category ORDER BY category")
	sql := buf.String()

	Logger.Debugf("GetPeriodResultStruct sql: %v", sql)

	var pr []PeriodResult
	_, err := o.Raw(sql).QueryRows(&pr)
	if err != nil {
		Logger.Errorf("GetPeriodResultStruct query failed %v", err)
		return nil, err
	}
	Logger.Debugf("GetPeriodResultStruct return value:%v", pr)
	return pr, nil
}

/*得到投资金额占比结果*/
func GetInvestResult(odrs *OperationalDataRequestStruct) (*InvestResult, error) {
	Logger.Debugf("GetInvestResult input param: %v", odrs)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("SELECT SUM(IF(account_act<10000,`account_act`,0))/10000 AS a1,")
	buf.WriteString("SUM(IF(account_act>=10000 AND account_act<100000,`account_act`,0))/10000 AS a2,")
	buf.WriteString("SUM(IF(account_act>=100000 AND account_act<500000,`account_act`,0))/10000 AS a3,")
	buf.WriteString("SUM(IF(account_act>=500000 AND account_act<1000000,`account_act`,0))/10000 AS a4,")
	buf.WriteString("SUM(IF(account_act>=1000000,`account_act`,0))/10000 AS a5 FROM jl_borrow_tender WHERE status=1 LIMIT 1 ")
	sql := buf.String()

	Logger.Debugf("GetInvestResult sql: %v", sql)

	var ir InvestResult
	err := o.Raw(sql).QueryRow(&ir)
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err != nil {
		Logger.Errorf("GetInvestResult query failed %v", err)
		return nil, err
	}
	Logger.Debugf("GetInvestResult return value:%v", ir)
	return &ir, nil
}

/*得到标的比例结果*/
func GetBidResult(odrs *OperationalDataRequestStruct) ([]BidResult, error) {
	Logger.Debugf("GetBidResult input param: %v", odrs)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("SELECT `borrow_type`,COUNT(1) AS number FROM jl_borrow WHERE status IN (2,3,6,7) GROUP BY borrow_type")
	sql := buf.String()

	Logger.Debugf("GetBidResult sql: %v", sql)

	var br []BidResult
	_, err := o.Raw(sql).QueryRows(&br)
	if err != nil {
		Logger.Errorf("GetBidResult query failed %v", err)
		return nil, err
	}
	Logger.Debugf("GetBidResult return value:%v", br)
	return br, nil
}

/*得到实时待收排行榜结果*/
func GetWaitResult(odrs *OperationalDataRequestStruct) ([]WaitResult, error) {
	Logger.Debugf("GetWaitResult input param: %v", odrs)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("SELECT U.username,A.hsmoney_wait AS money FROM jl_user U LEFT JOIN jl_account A ON A.user_id=U.id ")
	buf.WriteString("WHERE U.isvest=0 ORDER BY A.hsmoney_wait DESC LIMIT 10")
	sql := buf.String()

	Logger.Debugf("GetWaitResult sql: %v", sql)

	var wr []WaitResult
	_, err := o.Raw(sql).QueryRows(&wr)
	if err != nil {
		Logger.Errorf("GetWaitResult query failed %v", err)
		return nil, err
	}
	Logger.Debugf("GetWaitResult return value:%v", wr)
	return wr, nil
}

/*得到12个月之前成交总量*/
func GetTwelveMonthTotalNum(odrs *OperationalDataRequestStruct) (string, error) {
	Logger.Debugf("GetTwelveMonthTotalNum input param: %v", odrs)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("SELECT SUM(account) FROM jl_borrow WHERE review_time<? AND status IN (2,3,6,7) LIMIT 1")
	sql := buf.String()

	Logger.Debugf("GetTwelveMonthTotalNum sql: %v", sql)

	var totalNum string
	err := o.Raw(sql, odrs.StartMonth).QueryRow(&totalNum)
	if err != nil {
		Logger.Debugf("GetTwelveMonthTotalNum query failed %v", err)
		return "0", nil
	}
	Logger.Debugf("GetTwelveMonthTotalNum return value:%v", totalNum)
	return totalNum, nil
}

/*得到目前累计成功还款金额*/
func GetTotalRepayment(odrs *OperationalDataRequestStruct) (string, error) {
	Logger.Debugf("GetTotalRepayment input param: %v", odrs)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("SELECT SUM(replayment_money) FROM jl_borrow_repayment WHERE status=1 LIMIT 1")
	sql := buf.String()

	Logger.Debugf("GetTotalRepayment sql: %v", sql)

	var totalRepayment string
	err := o.Raw(sql).QueryRow(&totalRepayment)
	if err != nil {
		Logger.Debugf("GetTotalRepayment query failed %v", err)
		return "0", nil
	}
	Logger.Debugf("GetTotalRepayment return value:%v", totalRepayment)
	return totalRepayment, nil
}

/**
 * [GetTender 获取今天前投标且已经审核通过的投资总额]
 * @param    odrs *OperationalDataRequestStruct 请求入参
 * @return   string 投资总额
 * @DateTime 2018-01-22T17:40:03+0800
 */
func GetTender(odrs *OperationalDataRequestStruct) (string, error) {
	Logger.Debugf("GetTender input param: %v", odrs)

	var tender string
	err := utils.GetCache("borrow:tender:account_act", &tender)
	if err != nil {
		o := orm.NewOrm()
		o.Using("default")

		buf := bytes.Buffer{}
		buf.WriteString("SELECT SUM(BT.account_act) FROM jl_borrow_tender BT JOIN jl_borrow B ON BT.borrow_id=B.id ")
		buf.WriteString("WHERE BT.status=1 AND B.review_time<?")
		sql := buf.String()
		Logger.Debugf("GetTender sql: %v", sql)

		err := o.Raw(sql, odrs.TodayTime).QueryRow(&tender)
		if err == orm.ErrNoRows {
			return "0.00", nil
		} else if err != nil {
			Logger.Errorf("GetTender query failed:%v", err)
			return "0.00", err
		}
		utils.SetCache("borrow:tender", tender, 600)
	}

	Logger.Debugf("GetTender return value:%v", tender)
	return tender, nil
}

/**
 * [GetTenderToday 获取昨日满标复审通过的标的总额]
 * @param    odrs *OperationalDataRequestStruct 请求入参
 * @return   string 总额
 * @DateTime 2018-01-22T17:40:58+0800
 */
func GetTenderToday(odrs *OperationalDataRequestStruct) (string, error) {
	Logger.Debugf("GetTenderToday input param: %v", odrs)

	var tenderToday string
	err := utils.GetCache("borrow:account", &tenderToday)
	if err != nil {
		o := orm.NewOrm()
		o.Using("default")

		buf := bytes.Buffer{}
		buf.WriteString("SELECT SUM(account) FROM jl_borrow ")
		buf.WriteString("WHERE status=6 AND review_time>=? AND review_time<?")
		sql := buf.String()
		Logger.Debugf("GetTenderToday sql: %v", sql)

		err := o.Raw(sql, odrs.YesterdayTime, odrs.TodayTime).QueryRow(&tenderToday)
		if err == orm.ErrNoRows {
			return "0.00", nil
		} else if err != nil {
			Logger.Errorf("GetTenderToday query failed:%v", err)
			return "0.00", err
		}
		utils.SetCache("borrow:account", tenderToday, 600)
	}

	Logger.Debugf("GetTenderToday return value:%v", tenderToday)
	return tenderToday, nil
}

/**
 * [GetProfit 获取今天前已还利息总额]
 * @param    odrs *OperationalDataRequestStruct 请求入参
 * @return   string 已还利息总额
 * @DateTime 2018-01-22T17:41:38+0800
 */
func GetProfit(odrs *OperationalDataRequestStruct) (string, error) {
	Logger.Debugf("GetProfit input param: %v", odrs)

	var profit string
	err := utils.GetCache("borrow:collection:interest", &profit)
	if err != nil {
		o := orm.NewOrm()
		o.Using("default")

		buf := bytes.Buffer{}
		buf.WriteString("SELECT SUM(interest) FROM jl_borrow_collection ")
		buf.WriteString("WHERE status=1 AND repay_yestime<?")
		sql := buf.String()

		Logger.Debugf("GetProfit sql: %v", sql)

		err := o.Raw(sql, odrs.TodayTime).QueryRow(&profit)
		if err == orm.ErrNoRows {
			return "0.00", nil
		} else if err != nil {
			Logger.Errorf("GetProfit query failed:%v", err)
			return "0.00", err
		}
		utils.SetCache("borrow:collection:interest", profit, 600)
	}

	Logger.Debugf("GetProfit return value:%v", profit)
	return profit, nil
}

/**
 * [GetTenderUserCount 获取出借人数]
 * @param    odrs *OperationalDataRequestStruct 请求入参
 * @return   string 获取出借人数
 * @DateTime 2018-01-22T17:43:30+0800
 */
func GetTenderUserCount(odrs *OperationalDataRequestStruct) (int32, error) {
	Logger.Debugf("GetTenderUserCount input param: %v", odrs)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("SELECT COUNT(DISTINCT user_id) AS user_count FROM jl_borrow_tender")
	sql := buf.String()

	Logger.Debugf("GetTenderUserCount sql: %v", sql)

	var tenderUserCount int32
	err := o.Raw(sql).QueryRow(&tenderUserCount)
	if err == orm.ErrNoRows {
		return 0, nil
	} else if err != nil {
		Logger.Errorf("GetTenderUserCount query failed:%v", err)
		return 0, err
	}
	Logger.Debugf("GetTenderUserCount return value:%v", tenderUserCount)
	return tenderUserCount, nil
}
