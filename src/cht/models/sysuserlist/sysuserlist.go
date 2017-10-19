package sysuserlist

import (
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type SysUserListRequest struct {
	ChengHuiTongTraceLog string
}

type SysUserDetails struct {
	ID            int32  `orm:"column(id)"`
	RoleID        int32  `orm:"column(role_id)"`
	Account       string `orm:"column(account)"`
	Realname      string `orm:"column(realname)"`
	Password      string `orm:"column(password)"`
	Mobile        string `orm:"column(mobile)"`
	Qq            string `orm:"column(qq)"`
	Lastloginip   string `orm:"column(lastloginip)"`
	Lastlogintime int32  `orm:"column(lastlogintime)"`
	CreateTime    int32  `orm:"column(create_time)"`
	Status        int32  `orm:"column(status)"`
	Views         int32  `orm:"column(views)"`
	CustomerType  int32  `orm:"column(customer_type)"`
}

func GetSysUserList(sulr *SysUserListRequest) ([]SysUserDetails, error) {
	Logger.Debugf("GetSysUserList input param:%v", sulr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_sys_user")

	sql := qb.String()
	Logger.Debug("GetSysUserList sql:", sql)
	var sud []SysUserDetails
	_, err := o.Raw(sql).QueryRows(&sud)
	if err != nil {
		Logger.Debugf("GetSysUserList query failed %v", err)
		return nil, err
	}
	Logger.Debugf("GetSysUserList res %v", sud)
	return sud, nil
}
