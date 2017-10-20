package userattestionbaseinfosave

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type UserAttestionCardInfoSaveRequest struct {
	UserID               int32  `orm:"column(user_id)"`
	CardType             int32  `orm:"column(card_type)"`
	CardID               string `orm:"column(card_id)"`
	ChengHuiTongTraceLog string `orm:"column(chengHuiTongTraceLog)"`
}

/**
 * [SaveUserAttestionCardInfo 保存用户认证ID信息]
 * @param    {[type]}                 uacisr *UserAttestionCardInfoSaveRequest)(bool [description]
 * @DateTime 2017-10-20T09:19:33+0800
 */
func SaveUserAttestionCardInfo(uacisr *UserAttestionCardInfoSaveRequest) bool {
	Logger.Debugf("SaveUserAttestionCardInfo input param:%v", uacisr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Update("jl_user_attestation")

	var str string
	if strings.TrimSpace(uacisr.CardID) != "" {
		str += fmt.Sprintf("card_id=\"%s\",", strings.TrimSpace(uacisr.CardID))
	}

	if uacisr.CardType != 0 {
		str += fmt.Sprintf("card_type=%d,", uacisr.CardType)
	}

	str = strings.TrimSuffix(str, ",")
	qb.Set(str)
	qb.Where(fmt.Sprintf("user_id=%d", uacisr.UserID))
	sql := qb.String()

	Logger.Debug("SaveUserAttestionCardInfo sql:", sql)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		Logger.Debugf("SaveUserAttestionCardInfo update failed :%v", err)
		return false
	}
	num, _ := res.RowsAffected()
	Logger.Debugf("SaveUserAttestionCardInfo change num :%v", num)
	if num == 0 {
		return false
	}
	return true
}
