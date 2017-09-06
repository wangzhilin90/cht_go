package gettenderredbag

import (
	"bytes"
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type TenderRedbagRequest struct {
	UserId               int32
	TenderId             int32  //标ID
	RedId                int32  //红包ID
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
func getBorrowType(trr *TenderRedbagRequest) (int32, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("getBorrowType input param:", trr)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("borrow_type").
		From("jl_borrow").
		Where(fmt.Sprintf("id=%d", trr.TenderId))

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

type MoneyStruct struct {
	Money string `orm:"column(money)"`
}

/**
 * [GetRedBagMoney 获取红包金额]
 * @param    trr *TenderRedbagRequest请求入参
 * @return   int32 返回红包金额
 * @DateTime 2017-09-06T11:05:55+0800
 */
func GetRedBagMoney(trr *TenderRedbagRequest) (string, error) {
	Logger.Debug("GetRedBagMoney input param:", trr)
	borrowType, err := getBorrowType(trr)
	if err != nil {
		return "", err
	}

	o := orm.NewOrm()
	o.Using("default")

	sql := bytes.Buffer{}
	sql.WriteString("SELECT money FROM jl_three_redbag WHERE ")
	sql.WriteString("user_id=? and id=? and status=1 ")
	sql.WriteString("and min_tender<=? and (max_tender=0 OR max_tender>=?) ")
	sql.WriteString(" and FIND_IN_SET(?, time_limit) ")
	sql.WriteString(" and FIND_IN_SET(?, borrow_type) ")
	sql.WriteString(" LIMIT 1")
	Logger.Debug("GetRedBagMoney sql:", sql.String())
	Logger.Debug("user_id", trr.UserId)
	Logger.Debug("red_id", trr.RedId)

	var ms MoneyStruct
	err = o.Raw(sql.String(),
		trr.UserId,
		trr.RedId,
		trr.TenderMoney,
		trr.TenderMoney,
		trr.TimeLimit,
		borrowType,
	).QueryRow(&ms)
	if err != nil {
		Logger.Errorf("GetRedBagMoney query failed ", err)
		return "", err
	}
	Logger.Debugf("GetBorrowType res ", ms.Money)
	return ms.Money, nil
}
