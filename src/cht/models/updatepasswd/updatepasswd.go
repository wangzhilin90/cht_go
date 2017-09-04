package updatepasswd

import (
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type UpdatePasswdRequest struct {
	ID       int32  `orm:"column(id)"`
	Password string `orm:"column(password)"`
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
	res, err := o.Raw("update jl_user set password=? where id=?", upr.Password, upr.ID).Exec()
	if err != nil {
		Logger.Error("UpdatePasswd update failed")
		return false
	}
	num, _ := res.RowsAffected()
	Logger.Debug("rows effect", num)
	return true
}
