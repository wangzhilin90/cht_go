package userdetailsbynamepassword

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type UserDetailsByNamePasswordRequest struct {
	Name                 string
	Password             string
	ChengHuiTongTraceLog string
}

type SysUserDetailsStruct struct {
	ID            int32  `orm:column(id)`
	RoleID        int32  `orm:column(role_id)`
	Account       string `orm:column(account)`
	Realname      string `orm:column(realname)`
	Password      string `orm:column(password)`
	Mobile        string `orm:column(mobile)`
	Qq            string `orm:column(qq)`
	Lastloginip   string `orm:column(lastloginip)`
	Lastlogintime int32  `orm:column(lastlogintime)`
	CreateTime    int32  `orm:column(create_time)`
	Status        int32  `orm:column(status)`
	Views         int32  `orm:column(views)`
	CustomerType  int32  `orm:column(customer_type)`
}

func GetUseDetailsrByNamePassword(udbpr *UserDetailsByNamePasswordRequest) (*SysUserDetailsStruct, error) {
	Logger.Debugf("GetUseDetailsrByNamePassword input param:%v", *udbpr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_sys_user").
		Where(fmt.Sprintf("account=\"%s\"", udbpr.Name)).
		And(fmt.Sprintf("password=\"%s\"", udbpr.Password))

	sql := qb.String()
	Logger.Debugf("GetUseDetailsrByNamePassword sql", sql)
	var surs SysUserDetailsStruct
	err := o.Raw(sql).QueryRow(&surs)
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err != nil {
		Logger.Errorf("GetUseDetailsrByNamePassword query failed:%v", err)
		return nil, err
	}
	Logger.Debugf("GetUseDetailsrByNamePassword res:%v", surs)
	return &surs, nil
}
