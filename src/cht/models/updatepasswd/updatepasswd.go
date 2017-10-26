package updatepasswd

import (
	"bytes"
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type UpdatePasswdRequest struct {
	ID           int32
	NewPassword_ string
	OldPassword  string
}

func GetDBPasswd(upr *UpdatePasswdRequest) (string, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("GetDBPasswd input param:", upr)

	buf := bytes.Buffer{}
	buf.WriteString("SELECT password FROM jl_user WHERE id=? limit 1")
	sql := buf.String()
	Logger.Debugf("GetDBPasswd sql %v", sql)

	var dbPassword string
	err := o.Raw(sql, upr.ID).QueryRow(&dbPassword)
	if err != nil {
		Logger.Errorf("GetDBPasswd query failed :%v", err)
		return "", err
	}

	Logger.Debugf("GetDBPasswd res :%v", dbPassword)
	return dbPassword, nil
}

/**
 * [UpdatePasswd 更新密码]
 * @param   upr *UpdatePasswdRequest 请求入参
 * @return  bool true：更新成功，false更新失败
 * @DateTime 2017-09-04T11:19:05+0800
 */
func UpdatePasswd(upr *UpdatePasswdRequest) bool {
	o := orm.NewOrm()
	o.Using("default")

	Logger.Debug("UpdatePasswd input param:", upr)
	res, err := o.Raw("update jl_user set password=? where id=?", upr.NewPassword_, upr.ID).Exec()
	if err != nil {
		Logger.Error("UpdatePasswd update failed")
		return false
	}
	num, _ := res.RowsAffected()
	if num == 0 {
		return false
	}
	Logger.Debug("rows effect", num)
	return true
}
