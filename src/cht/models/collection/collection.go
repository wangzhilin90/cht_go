package collection

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

const (
	EVEN_DAYS_QUANTUM = int64(24 * 3600 * 7)
	ONE_MONTH_QUANTUM = int64(24 * 3600 * 30)
	TWO_MONTH_QUANTUM = int64(24 * 3600 * 60)
)

type CollectionRequest struct {
	UserID               int32
	Starttime            int32
	Endtime              int32
	SearchTime           int32
	State                int32
	LimitOffset          int32
	LimitNum             int32
	Borrowid             string
	ChengHuiTongTraceLog string
}

type CollectionInfoStruct struct {
	Username        string `orm:"column(username)"`
	Title           string `orm:"column(title)"`
	IsDatetype      int32  `orm:"column(is_datetype)"`
	TimeLimit       int32  `orm:"column(time_limit)"`
	Zhuanrangren    string `orm:"column(zhuanrangren)"`
	RepayTime       int32  `orm:"column(repay_time)"`
	BorrowID        int32  `orm:"column(borrow_id)"`
	Periods         int32  `orm:"column(periods)"`
	RepayYestime    int32  `orm:"column(repay_yestime)"`
	RepayYesaccount string `orm:"column(repay_yesaccount)"`
	RepayAccount    string `orm:"column(repay_account)"`
	Capital         string `orm:"column(capital)"`
	Interest        string `orm:"column(interest)"`
	LateInterest    string `orm:"column(late_interest)"`
	LateDays        int32  `orm:"column(late_days)"`
	Status          int32  `orm:"column(status)"`
	InterestAdd     string `orm:"column(interest_add)"`
	OldUserID       int32  `orm:"column(old_user_id)"`
	Style           int32  `orm:"column(style)"`
}

/**
 * [GetCollectionInfo 得到回款明细信息]
 * @param    trr *TenderRedbagRequest 请求入参
 * @return   []CollectionInfoStruct   返回回款信息
 * @return   int32 返回没有limit的查询总数
 * @DateTime 2017-09-08T11:37:23+0800
 */
func GetCollectionInfo(trr *CollectionRequest) ([]CollectionInfoStruct, int32, error) {
	Logger.Debugf("GetCollectionInfo input param: %v", trr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("U.username,B.title,B.is_datetype,B.time_limit,B.zhuanrangren,BC.repay_time,BT.borrow_id,BC.periods,BC.repay_yestime,BC.repay_yesaccount,BC.repay_account,BC.capital,BC.interest,BC.late_interest,BC.late_days,BC.status,BC.interest_add,BC.old_user_id,B.style").
		From("jl_borrow_collection BC").
		LeftJoin("jl_borrow B").On("BC.borrow_id=B.id").
		LeftJoin("jl_user U").On("B.user_id=U.id").
		LeftJoin("jl_borrow_tender BT").On("BC.tender_id=BT.id").
		Where(fmt.Sprintf("BC.user_id=%d", trr.UserID))

	//0：查全部,1:近7天，2:1个月，3:2个月
	switch {
	/*时间查全部，开始时间和结束时间才有效*/
	case trr.SearchTime == 0:
		if trr.Starttime != 0 {
			qb.And(fmt.Sprintf("BC.repay_time >= %d", trr.Starttime))
		}
		if trr.Endtime != 0 {
			qb.And(fmt.Sprintf("BC.repay_time <= %d", trr.Endtime))
		}
	case trr.SearchTime == 1:
		/*查最近七天充值记录*/
		qb.And(fmt.Sprintf("BC.repay_time >=%d", time.Now().Unix()-EVEN_DAYS_QUANTUM))
	case trr.SearchTime == 2:
		/*查最近一个月充值记录*/
		qb.And(fmt.Sprintf("BC.repay_time >=%d", time.Now().Unix()-ONE_MONTH_QUANTUM))
	case trr.SearchTime == 3:
		/*查最近两个月充值记录*/
		qb.And(fmt.Sprintf("BC.repay_time >=%d", time.Now().Unix()-TWO_MONTH_QUANTUM))
	}

	//0:全部，1: 还款中2：已回款
	switch {
	case trr.State == 1:
		/*查还款中记录*/
		qb.And(fmt.Sprintf("BC.status=%d", 0))
	case trr.State == 2:
		/*查已回款记录*/
		qb.And(fmt.Sprintf("BC.status=%d", 1)).OrderBy("BC.status ASC , BC.repay_time DESC , BC.id ASC")
	}

	if trr.Borrowid != "" {
		Logger.Debugf("GetCollectionInfo Borrowid:", trr.Borrowid)
		Logger.Debugf("GetCollectionInfo TrimPrefix Borrowid:", strings.TrimPrefix(trr.Borrowid, "CHT"))
		qb.And(fmt.Sprintf("BC.borrow_id=%s", strings.TrimPrefix(trr.Borrowid, "CHT")))
	}

	sql := qb.String()
	Logger.Debugf("GetCollectionInfo origin sql:", sql)

	var cis1 []CollectionInfoStruct
	totalnum, err := o.Raw(sql).QueryRows(&cis1)
	if err != nil {
		Logger.Debug("GetCollectionInfo query failed:", err)
		return nil, 0, err
	}
	/*得到总的查询数*/
	Logger.Debug("GetCollectionInfo query totalnum:", totalnum)

	/*带limit查询得到提现记录数据*/
	if trr.LimitNum != 0 {
		qb.Limit(int(trr.LimitNum))
	}
	if trr.LimitOffset != 0 {
		qb.Offset(int(trr.LimitOffset))
	}
	sql = qb.String()
	Logger.Debugf("GetCollectionInfo sql with limit:", sql)

	var cis []CollectionInfoStruct
	_, err = o.Raw(sql).QueryRows(&cis)
	if err != nil {
		Logger.Debug("GetCollectionInfo queryrows failed")
		return nil, 0, err
	}
	Logger.Debugf("GetCollectionInfo res:%v %d", cis, totalnum)
	return cis, int32(totalnum), nil
}
