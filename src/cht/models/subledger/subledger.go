package subledger

import (
	"bytes"
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type SubledgerRequest struct {
	HsZhuanrangrenStr    string
	ChengHuiTongTraceLog string
}

type SubledgerInfo struct {
	UserID   int32  `orm:"column(user_id)"`
	Realname string `orm:"column(realname)"`
	CardID   int32  `orm:"column(card_id)"`
}

/**
 * [GetSubledgerList 获取分账人的用户信息]
 * @param    sr *SubledgerRequest 请求入参
 * @DateTime 2017-09-12T15:17:37+0800
 */
func GetSubledgerList(sr *SubledgerRequest) ([]SubledgerInfo, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("GetSubledgerList input param:", sr)

	userArr := strings.Split(sr.HsZhuanrangrenStr, ",")
	Logger.Debugf("GetSubledgerList array %v", userArr)

	buf := bytes.Buffer{}

	var str string
	for range userArr {
		str += fmt.Sprintf(",%v", "?")
	}
	str = strings.TrimPrefix(str, ",")
	buf.WriteString("SELECT  UA.user_id,U.realname,UA.card_id  FROM jl_user_attestation UA LEFT JOIN  jl_user U ON UA.user_id=U.id  where UA.user_id IN (")
	buf.WriteString(str)
	buf.WriteString(")")
	sql := buf.String()
	Logger.Debugf("GetSubledgerList sql %v", sql)

	var si []SubledgerInfo
	_, err := o.Raw(sql, userArr).QueryRows(&si)
	if err != nil {
		Logger.Errorf("GetSubledgerList query failed %v", err)
		return nil, err
	}

	Logger.Debugf("GetSubledgerList res ", si)
	return si, nil
}
