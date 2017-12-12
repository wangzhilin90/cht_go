package borrowrepaymentstatistics

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type RepaymentStatisticsRequest struct {
	UserID               int32
	Status               int32
	ChengHuiTongTraceLog string
}

type RepaymentStatisticsDetails struct {
	BorrowID          int32  `orm:"column(borrow_id)"`
	WillMoney         string `orm:"column(will_money)"`
	ReplaymentMoney   string `orm:"column(replayment_money)"`
	NoreplaymentMoney string `orm:"column(noreplayment_money)"`
}

func GetRepaymentStatisticsList(rsr *RepaymentStatisticsRequest) ([]RepaymentStatisticsDetails, error) {
	Logger.Debugf("GetRepaymentStatisticsList input param:%v", rsr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("borrow_id,SUM(will_money) AS will_money,SUM(replayment_money) AS replayment_money,SUM(will_money-replayment_money) AS noreplayment_money").
		From("jl_borrow_repayment").
		Where(fmt.Sprintf("user_id=%d", rsr.UserID)).
		OrderBy("borrow_id").
		GroupBy(fmt.Sprintf("borrow_id,user_id"))

	sql := qb.String()
	Logger.Debugf("GetRepaymentStatisticsList sql:%v", sql)

	var rsd []RepaymentStatisticsDetails
	_, err := o.Raw(sql).QueryRows(&rsd)
	if err != nil {
		Logger.Error("GetRepaymentStatisticsList query failed:", err)
		return nil, err
	}
	Logger.Debugf("GetRepaymentStatisticsList res :%v", rsd)
	return rsd, nil
}

func GetTotalReplaymentMoney(rsr *RepaymentStatisticsRequest) (string, error) {
	Logger.Debugf("GetTotalReplaymentMoney input param:%v", rsr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("SUM(replayment_money) as total").
		From("jl_borrow_repayment").
		Where("1=1")

	if rsr.Status != -1 {
		qb.And(fmt.Sprintf("status=%d", rsr.Status))
	}

	sql := qb.String()
	Logger.Debugf("GetTotalReplaymentMoney sql:%v", sql)
	var totalMoney string
	err := o.Raw(sql).QueryRow(&totalMoney)
	if err != nil {
		Logger.Errorf("GetTotalReplaymentMoney query failed:%v", err)
		return "0", err
	}
	Logger.Debugf("GetTotalReplaymentMoney return value:%v", totalMoney)
	return totalMoney, nil
}
