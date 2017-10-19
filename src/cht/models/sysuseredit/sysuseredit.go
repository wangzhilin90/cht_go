package sysuseredit

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type SysUserEditRequest struct {
	Account              string `orm:"column(account)"`
	Password             string `orm:"column(password)"`
	Realname             string `orm:"column(realname)"`
	Mobile               string `orm:"column(mobile)"`
	Qq                   string `orm:"column(qq)"`
	Status               int32  `orm:"column(status)"`
	RoleID               int32  `orm:"column(role_id)"`
	CustomerType         int32  `orm:"column(customer_type)"`
	CreateTime           string `orm:"column(create_time)"`
	UserID               int32  `orm:"column(user_id)"`
	ChengHuiTongTraceLog string `orm:"column(chengHuiTongTraceLog)"`
}

/**
 * [EditSysUser 编辑后台管理用户]
 * @param    suer *SysUserEditRequest 请求入参
 * @return   bool  true:编辑成功 false:编辑失败
 * @DateTime 2017-10-18T16:31:39+0800
 */
func EditSysUser(suer *SysUserEditRequest) bool {
	Logger.Debugf("EditSysUser input param:%v", suer)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Update("jl_sys_user")

	var str string
	if suer.Account != "" {
		str += fmt.Sprintf("account=\"%s\",", suer.Account)
	}

	if suer.Password != "" {
		str += fmt.Sprintf("password=\"%s\",", suer.Password)
	}

	if suer.Realname != "" {
		str += fmt.Sprintf("realname=\"%s\",", suer.Realname)
	}

	if suer.Mobile != "" {
		str += fmt.Sprintf("mobile=\"%s\",", suer.Mobile)
	}

	if suer.Qq != "" {
		str += fmt.Sprintf("qq=\"%s\",", suer.Qq)
	}

	if suer.Status != 0 {
		str += fmt.Sprintf("status=%d,", suer.Status)
	}

	if suer.RoleID != 0 {
		str += fmt.Sprintf("role_id=%d,", suer.RoleID)
	}

	if suer.CustomerType != 0 {
		str += fmt.Sprintf("customer_type=%d,", suer.CustomerType)
	}

	str = strings.TrimSuffix(str, ",")
	qb.Set(str)
	qb.Where(fmt.Sprintf("id=%d", suer.UserID))
	sql := qb.String()

	Logger.Debug("EditSysUser sql:", sql)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		Logger.Debugf("EditSysUser update failed :%v", err)
		return false
	}
	num, _ := res.RowsAffected()
	Logger.Debugf("EditSysUser change num :%v", num)
	if num == 0 {
		return false
	}
	return true
}
