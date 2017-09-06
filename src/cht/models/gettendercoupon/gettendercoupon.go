package gettendercoupon

import (
	"bytes"
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type TenderCouponRequest struct {
	UserId               int32
	TenderId             int32  //标ID
	CouponId             int32  //加息券ID
	TenderMoney          string //用户投资金额
	TimeLimit            int32
	ChengHuiTongTraceLog string
}

type borrowStruct struct {
	BorrowType int32 `orm:"column(borrow_type)"`
}

/**
 * [getBorrowType 获取标类型]
 * @param    trr *TenderRedbagRequest 请求入参
 * @return   int32 返回标类型
 * @DateTime 2017-09-06T11:00:19+0800
 */
func getBorrowType(tcr *TenderCouponRequest) (int32, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("getBorrowType input param:", tcr)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("borrow_type").
		From("jl_borrow").
		Where(fmt.Sprintf("id=%d", tcr.TenderId))

	sql := qb.String()
	Logger.Debug("getBorrowType sql:", sql)
	var bs borrowStruct
	err := o.Raw(sql).QueryRow(&bs)
	if err != nil {
		Logger.Error("GetBorrowType query failed:", err)
		return 0, err
	}
	Logger.Debugf("GetBorrowType res ", bs.BorrowType)
	return bs.BorrowType, nil
}

type CouponStruct struct {
	Apr string `orm:"column(apr)"` //加息值
}

/**
 * [GetRedBagMoney 获取加息值]
 * @param    tcr *TenderCouponRequest 请求入参
 * @return   string 返回加息值
 * @DateTime 2017-09-06T11:05:55+0800
 */
func GetTenderCoupon(tcr *TenderCouponRequest) (string, error) {
	Logger.Debug("GetTenderCoupon input param:", tcr)
	borrowType, err := getBorrowType(tcr)
	if err != nil {
		return "", err
	}

	o := orm.NewOrm()
	o.Using("default")

	sql := bytes.Buffer{}
	sql.WriteString("SELECT apr FROM jl_apr_coupon WHERE ")
	sql.WriteString("user_id=? and id=? and status=1 ") //status=1为可用红包
	sql.WriteString("and min_tender<=? and (max_tender=0 OR max_tender>=?) ")
	sql.WriteString(" and FIND_IN_SET(?, time_limit) ")
	sql.WriteString(" and FIND_IN_SET(?, borrow_type) ")
	sql.WriteString(" LIMIT 1")
	Logger.Debug("GetTenderCoupon sql:", sql.String())
	Logger.Debug("user_id ", tcr.UserId)
	Logger.Debug("coupon_id ", tcr.CouponId)

	//用户投资时间大于18个月，当18个月时间投资
	if tcr.TimeLimit > 18 {
		tcr.TimeLimit = 18
	}

	var cs CouponStruct
	err = o.Raw(sql.String(),
		tcr.UserId,
		tcr.CouponId,
		tcr.TenderMoney,
		tcr.TenderMoney,
		tcr.TimeLimit,
		borrowType,
	).QueryRow(&cs)
	if err != nil {
		Logger.Errorf("GetTenderCoupon query failed ", err)
		return "", err
	}
	Logger.Debugf("GetBorrowType res ", cs.Apr)
	return cs.Apr, nil
}
