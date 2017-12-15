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
	if err == orm.ErrNoRows {
		return 0, nil
	} else if err != nil {
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
 * @return   string 返回红包金额
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

	buf := bytes.Buffer{}
	buf.WriteString("SELECT money FROM jl_three_redbag WHERE ")
	buf.WriteString("user_id=? and id=? and status=1 ")
	buf.WriteString("and min_tender<=? and (max_tender=0 OR max_tender>=?) ")
	buf.WriteString(" and FIND_IN_SET(?, time_limit) ")
	buf.WriteString(" and FIND_IN_SET(?, borrow_type) ")
	buf.WriteString(" LIMIT 1")

	sql := buf.String()
	Logger.Debug("GetRedBagMoney sql:", sql)

	//用户投资时间大于18个月，当18个月时间投资
	if trr.TimeLimit > 18 {
		trr.TimeLimit = 18
	}

	var ms MoneyStruct
	err = o.Raw(sql,
		trr.UserId,
		trr.RedId,
		trr.TenderMoney,
		trr.TenderMoney,
		trr.TimeLimit,
		borrowType,
	).QueryRow(&ms)
	if err == orm.ErrNoRows {
		return "", err
	} else if err != nil {
		Logger.Errorf("GetRedBagMoney query failed:%v", err)
		return "", err
	}
	Logger.Debugf("GetBorrowType res ", ms.Money)
	return ms.Money, nil
}
