package borrowrepaymentstatistics

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type RepaymentStatisticsRequest struct {
	UserID               int32
	ChengHuiTongTraceLog string
}

type RepaymentStatisticsDetails struct {
	BorrowID          int32  `orm:"column(borrow_id)"`
	WillMoney         string `orm:"column(will_money)"`
	ReplaymentMoney   string `orm:"column(replayment_money)"`
	NoreplaymentMoney string `orm:"column(noreplayment_money)"`
}

func GetRepaymentStatisticsDetails(rsr *RepaymentStatisticsRequest) ([]RepaymentStatisticsDetails, error) {
	Logger.Debugf("GetRepaymentStatisticsDetails input param:%v", rsr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("borrow_id,SUM(will_money) AS will_money,SUM(replayment_money) AS replayment_money,SUM(will_money-replayment_money) AS noreplayment_money").
		From("jl_borrow_repayment").
		Where(fmt.Sprintf("user_id=%d", rsr.UserID)).
		OrderBy("borrow_id").
		GroupBy(fmt.Sprintf("borrow_id,user_id"))

	sql := qb.String()
	Logger.Debugf("GetRepaymentStatisticsDetails sql:%v", sql)

	var rsd []RepaymentStatisticsDetails
	_, err := o.Raw(sql).QueryRows(&rsd)
	if err != nil {
		Logger.Error("GetRepaymentStatisticsDetails query failed:", err)
		return nil, err
	}
	Logger.Debugf("GetRepaymentStatisticsDetails res :%v", rsd)
	return rsd, nil
}
