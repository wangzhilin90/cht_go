package memberhelperlist

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type MemberHelperListRequest struct {
	Type                 int32
	Keywords             string
	LimitOffset          int32
	LimitNum             int32
	ChengHuiTongTraceLog string
}

type MemberHelperDetails struct {
	Linkman      string `orm:"column(linkman)"`
	Linkrelation string `orm:"column(linkrelation)"`
	Linkphone    string `orm:"column(linkphone)"`
	Updatetime   int32  `orm:"column(updatetime)"`
	ID           int32  `orm:"column(id)"`
	Username     string `orm:"column(username)"`
	Realname     string `orm:"column(realname)"`
	Phone        string `orm:"column(phone)"`
	Customer     int32  `orm:"column(customer)"`
}

/**
 * [GetMemberHelperListTotalNum 客户管理---会员紧急联系人总记录数]
 * @param    mhlr *MemberHelperListRequest 请求入参
 * @return   int32 总记录数
 * @DateTime 2017-10-23T13:59:55+0800
 */
func GetMemberHelperListTotalNum(mhlr *MemberHelperListRequest) (int32, error) {
	Logger.Debugf("GetMemberHelperListTotalNum input param:%v", mhlr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("COUNT(1) FROM jl_user_field UF LEFT JOIN jl_user U ON UF.user_id=U.id WHERE UF.linkman <>''")

	if mhlr.Type != 0 && mhlr.Keywords != "" {
		if mhlr.Type == 1 {
			qb.And(fmt.Sprintf("U.username like \"%%%s%%\"", mhlr.Keywords))
		} else if mhlr.Type == 2 {
			qb.And(fmt.Sprintf("U.realname like \"%%%s%%\"", mhlr.Keywords))
		} else if mhlr.Type == 3 {
			qb.And(fmt.Sprintf("U.phone=\"%s\"", mhlr.Keywords))
		} else {
			qb.And(fmt.Sprintf("id=\"%s\"", mhlr.Keywords))
		}
	}

	sql := qb.String()
	Logger.Debugf("GetMemberHelperListTotalNum sql:%v", sql)
	var totalNum int32
	err := o.Raw(sql).QueryRow(&totalNum)
	if err != nil {
		Logger.Debugf("GetMemberHelperListTotalNum query failed :%v", err)
		return 0, err
	}
	Logger.Debugf("GetMemberHelperListTotalNum return num:%v", totalNum)
	return totalNum, nil
}

/**
 * [GetMemberHelperList 客户管理---会员紧急联系人列表]
 * @param    mhlr *MemberHelperListRequest 请求入参
 * @return   []MemberHelperDetails 列表详情
 * @DateTime 2017-10-23T14:02:00+0800
 */
func GetMemberHelperList(mhlr *MemberHelperListRequest) ([]MemberHelperDetails, error) {
	Logger.Debugf("GetMemberHelperList input param:%v", mhlr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(`UF.linkman,UF.linkrelation,UF.linkphone,UF.updatetime, 
	 U.id,U.username,U.realname,U.phone,U.customer FROM jl_user_field UF LEFT JOIN jl_user U ON UF.user_id=U.id`).
		Where("UF.linkman <> ''")

	if mhlr.Type != 0 && mhlr.Keywords != "" {
		if mhlr.Type == 1 {
			qb.And(fmt.Sprintf("U.username like \"%%%s%%\"", mhlr.Keywords))
		} else if mhlr.Type == 2 {
			qb.And(fmt.Sprintf("U.realname like \"%%%s%%\"", mhlr.Keywords))
		} else if mhlr.Type == 3 {
			qb.And(fmt.Sprintf("U.phone=\"%s\"", mhlr.Keywords))
		} else {
			qb.And(fmt.Sprintf("id=\"%s\"", mhlr.Keywords))
		}
	}

	qb.OrderBy("id").Desc()

	//0:默认不导出,此时limit和offset生效

	if mhlr.LimitNum != 0 {
		qb.Limit(int(mhlr.LimitNum))
	}
	if mhlr.LimitOffset != 0 {
		qb.Offset(int(mhlr.LimitOffset))
	}

	sql := qb.String()
	Logger.Debugf("GetMemberHelperList sql:%v", sql)
	var mhd []MemberHelperDetails
	_, err := o.Raw(sql).QueryRows(&mhd)
	if err != nil {
		Logger.Errorf("GetMemberHelperList query failed :%v", err)
		return nil, err
	}
	Logger.Debugf("GetMemberHelperList return value:%v", mhd)
	return mhd, nil
}
