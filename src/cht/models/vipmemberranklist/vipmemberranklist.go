package vipmemberranklist

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type VipMemberRankListRequest struct {
	Type                 int32  `thrift:"type,1" db:"type" json:"type"`
	Keywords             string `thrift:"keywords,2" db:"keywords" json:"keywords"`
	LimitOffset          int32  `thrift:"limitOffset,3" db:"limitOffset" json:"limitOffset"`
	LimitNum             int32  `thrift:"limitNum,4" db:"limitNum" json:"limitNum"`
	ChengHuiTongTraceLog string `thrift:"chengHuiTongTraceLog,5" db:"chengHuiTongTraceLog" json:"chengHuiTongTraceLog"`
}

type VipMemberRankDetails struct {
	UserID        int32  `orm:"column(user_id)"`
	Money         string `orm:"column(money)"`
	MoneyFreeze   string `orm:"column(money_freeze)"`
	MoneyUsable   string `orm:"column(money_usable)"`
	MoneyWait     string `orm:"column(money_wait)"`
	CashTime      int32  `orm:"column(cash_time)"`
	ReturnMoney   string `orm:"column(return_money)"`
	VipLevel      int32  `orm:"column(vip_level)"`
	VipWaitMoney  string `orm:"column(vip_wait_money)"`
	HsreturnMoney string `orm:"column(hsreturn_money)"`
	HsmoneyFreeze string `orm:"column(hsmoney_freeze)"`
	HsmoneyWait   string `orm:"column(hsmoney_wait)"`
	Username      string `orm:"column(username)"`
	Realname      string `orm:"column(realname)"`
	Addtime       int32  `orm:"column(addtime)"`
}

/**
 * [GetVipMemberRankListTotalNum VIP会员等级总条目数]
 * @param    vmrlr *VipMemberRankListRequest 请求入参
 * @return   int32 总数
 * @DateTime 2017-10-27T11:43:43+0800
 */
func GetVipMemberRankListTotalNum(vmrlr *VipMemberRankListRequest) (int32, error) {
	Logger.Debugf("GetVipMemberRankListTotalNum input param:%v", vmrlr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("COUNT(1)").
		From("jl_user U").
		LeftJoin("jl_account A").On("A.user_id=U.id").
		Where("1=1")

	if vmrlr.Type != 0 && vmrlr.Keywords != "" {
		if vmrlr.Type == 1 {
			qb.And(fmt.Sprintf("U.username=\"%s\"", vmrlr.Keywords))
		} else {
			qb.And(fmt.Sprintf("U.realname=\"%s\"", vmrlr.Keywords))
		}
	}

	sql := qb.String()
	Logger.Debugf("GetVipMemberRankListTotalNum sql:%v", sql)

	var totalNum int32
	err := o.Raw(sql).QueryRow(&totalNum)
	if err != nil {
		Logger.Errorf("GetVipMemberRankListTotalNum query failed :%v", err)
		return 0, err
	}
	Logger.Debugf("GetVipMemberRankListTotalNum return num:%v", totalNum)
	return totalNum, nil
}

func GetVipMemberRankList(vmrlr *VipMemberRankListRequest) ([]VipMemberRankDetails, error) {
	Logger.Debugf("GetVipMemberRankList input param:%v", vmrlr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("A.*,U.username,U.realname,U.addtime").
		From("jl_user U").
		LeftJoin("jl_account A").On("A.user_id=U.id").
		Where("1=1")

	if vmrlr.Type != 0 && vmrlr.Keywords != "" {
		if vmrlr.Type == 1 {
			qb.And(fmt.Sprintf("U.username=\"%s\"", vmrlr.Keywords))
		} else {
			qb.And(fmt.Sprintf("U.realname=\"%s\"", vmrlr.Keywords))
		}
	}

	qb.OrderBy("A.vip_level DESC,A.vip_wait_money DESC")

	if vmrlr.LimitNum != 0 {
		qb.Limit(int(vmrlr.LimitNum))
	}
	if vmrlr.LimitOffset != 0 {
		qb.Offset(int(vmrlr.LimitOffset))
	}

	sql := qb.String()
	Logger.Debugf("GetVipMemberRankList sql:%v", sql)

	var vmrd []VipMemberRankDetails
	_, err := o.Raw(sql).QueryRows(&vmrd)
	if err != nil {
		Logger.Errorf("GetVipMemberRankList query failed :%v", err)
		return nil, err
	}
	Logger.Debugf("GetVipMemberRankList return value:%v", vmrd)
	return vmrd, nil
}
