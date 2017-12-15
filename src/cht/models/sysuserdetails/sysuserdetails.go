package sysuserdetails

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type SysUserDetailsRequest struct {
	UserID               int32
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

/**
 * [GetSysUserDetails 后台管理员详情]
 * @param    sudr *SysUserDetailsRequest 请求入参
 * @return   SysUserDetails 详情结果
 * @DateTime 2017-10-18T10:07:58+0800
 */
func GetSysUserDetails(sudr *SysUserDetailsRequest) (*SysUserDetails, error) {
	Logger.Debugf("GetSysUserDetails input param:%v", sudr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_sys_user").
		Where(fmt.Sprintf("id = %d", sudr.UserID))

	sql := qb.String()
	Logger.Debug("GetSysUserDetails sql:", sql)
	var sud SysUserDetails
	err := o.Raw(sql).QueryRow(&sud)
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err != nil {
		Logger.Errorf("GetSysUserDetails query failed %v", err)
		return nil, err
	}
	Logger.Debugf("GetSysUserDetails res %v", sud)
	return &sud, nil
}
