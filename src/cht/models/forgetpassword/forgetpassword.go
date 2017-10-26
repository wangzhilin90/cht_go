package forgetpassword

import (
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type ForgetPasswordRequest struct {
	ID       int32
	Password string
}

//忘记密码服务
func ForgetPassword(fpw *ForgetPasswordRequest) bool {
	o := orm.NewOrm()
	o.Using("default")

	Logger.Debug("ForgetPassword input param:", fpw)
	res, err := o.Raw("update jl_user set password=? where id=?", fpw.Password, fpw.ID).Exec()
	if err != nil {
		Logger.Error("ForgetPassword update failed")
		return false
	}
	num, _ := res.RowsAffected()
	if num == 0 {
		return false
	}
	Logger.Debug("rows effect", num)
	return true
}
